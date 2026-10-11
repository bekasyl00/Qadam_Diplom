# Qadam Backend

The Go API handles task drafts, confirmed readiness scores, published briefs, team proposals, manual decisions, and one verified milestone per accepted proposal. PostgreSQL stores application state. AI clarification, chat, and transcription are available through OpenAI or Groq.

The server uses Go 1.23+, `net/http`, `database/sql`, `lib/pq`, and `godotenv`, without an ORM. The Vue frontend connects through the Vite development proxy.

## Local startup

Start Docker first. From `backend`, using Bash or Zsh:

```bash
docker compose up -d --wait postgres
[ -f .env ] || cp .env.example .env
go mod download
go run ./cmd/seed
go run ./cmd/api
```

PowerShell equivalent:

```powershell
docker compose up -d --wait postgres
if (!(Test-Path .env)) { Copy-Item .env.example .env }
go mod download
go run ./cmd/seed
go run ./cmd/api
```

The API listens on `http://127.0.0.1:8080`. Check `/api/health` in a browser or run:

```bash
curl http://127.0.0.1:8080/api/health
```

Commands load `.env` from the current directory. Existing environment variables take precedence. Change `PORT` if needed and restart the API after configuration changes.

Startup applies embedded, ordered SQL migrations transactionally. The seed command adds missing sample records and preserves existing edits.

## PostgreSQL and optional administration UI

The supplied Compose service exposes PostgreSQL on `127.0.0.1:5433`. Keep `DATABASE_URL` from [.env.example](.env.example) when using this setup. Local database credentials and the database name are defined in [docker-compose.yml](docker-compose.yml).

To also start pgAdmin:

```bash
docker compose up -d pgadmin
```

Open `http://localhost:5050` and use the local pgAdmin credentials in the Compose file. Register PostgreSQL with host `postgres` and port `5432` inside the Compose network, using the database credentials from the same file. The Go process running on the host instead uses port `5433`.

## AI configuration

Edit `.env` locally and select one provider. Do not commit credentials or send them to the frontend or Postman.

### OpenAI

```dotenv
AI_PROVIDER=openai
OPENAI_API_KEY=replace_with_your_key
OPENAI_MODEL=gpt-4.1-mini
OPENAI_TRANSCRIBE_MODEL=gpt-4o-mini-transcribe
AI_MODE=live
```

### Groq

```dotenv
AI_PROVIDER=groq
GROQ_API_KEY=replace_with_your_key
AI_MODEL=openai/gpt-oss-20b
AI_TRANSCRIBE_MODEL=whisper-large-v3-turbo
AI_MODE=live
```

`AI_PROVIDER` defaults to `openai`. The listed models are configurable defaults in `internal/config/config.go`. Access depends on the provider account and quota. Groq's base URL can be overridden with `AI_BASE_URL`.

To use manual task workflows without a provider key, set `AI_MODE=fallback`. The public clarification endpoint reports AI unavailability rather than returning prepared questions as model output.

Check live provider access from `backend`:

```bash
go run ./cmd/check-ai
```

The command sends a small real analysis request and exits unsuccessfully if live generation fails. It may consume provider quota. Its initial setup message mentions OpenAI, but configuration also accepts Groq through the variables above.

The optional `scripts/set-openai.ps1` helper provides masked key entry on Windows for OpenAI configuration. `.env` is ignored by Git but remains a plaintext local file.

## AI operations

- **Analysis:** generates clarification questions and proposed fields with bounded output, source checks, retry handling, and a total analysis timeout. OpenAI requests use structured output and `store=false`.
- **Chat:** `POST /api/ai/chat` loads an accessible public snapshot or the selected business actor's own draft and sends contextual messages to the configured provider.
- **Transcription:** `POST /api/ai/transcribe` accepts WebM, MP4, Ogg, or WAV audio up to 15 MiB, validates the container signature, and returns editable text. The application does not retain the audio.

Provider failures are returned as bounded application errors without exposing provider response bodies. The internal fallback analyzer supports diagnostics; it is not presented to users as live AI.

## API contracts and profiles

- [Frontend routes and payloads](docs/api-request-list.md)
- [Frontend integration contract](../frontend/API-CONTRACT.md)
- [Chat and voice contract](../frontend/AI-CONTRACT.md)
- [Postman walkthrough](docs/postman/README.md)

Select a prepared profile with `X-Demo-Actor: business:UUID` or `team:UUID`. `GET /api/teams` returns team profiles and IDs. The seeded business ID is `10000000-0000-4000-8000-000000000001`.

Task IDs, draft IDs, proposal IDs, and revisions returned by the API should be used for subsequent requests. Seed data maps its fixture identifiers to stable UUIDs.

A typical flow is: create a draft, request clarification, save answers and fields, confirm information, publish, submit a proposal, accept it manually, submit a milestone, and confirm the result. Repeated confirmation does not award points twice.

## Verification

From `backend`:

```bash
go test ./...
go vet ./...
go build ./...
```

For PostgreSQL integration checks, set `TEST_DATABASE_URL` to a local PostgreSQL connection string. The supplied connection is documented in `.env.example`.

Bash or Zsh:

```bash
TEST_DATABASE_URL='<local_postgresql_connection_string>' go test ./... -count=1 -v -timeout=90s
```

PowerShell:

```powershell
$env:TEST_DATABASE_URL = '<local_postgresql_connection_string>'
go test ./... -count=1 -v -timeout=90s
```

Database tests create and drop randomly named schemas and require permission to create schemas. They skip without `TEST_DATABASE_URL`. They do not reset application data.

Coverage includes scoring and confirmation, public/private snapshots, seed idempotence, revision conflicts, team selection, milestone permissions, concurrent one-time awards, transaction rollback, persistence, AI source validation, malformed output, retries, and timeouts. Provider tests use a local HTTP fake; real provider access is checked separately.

## Current limits

- Prepared actors are not authentication. Keep the loopback-bound API local until production identity and authorization are implemented.
- `DEMO_MODE=false` disables application routes; it does not activate another login system.
- The development frontend proxies `/api` to Go. Production hosting needs separate routing; wildcard CORS is not enabled.
- Team selection and information confirmation remain human decisions.
- Only confirmed milestones contribute progress points. Older stored acceptance awards do not count toward progress.
- Existing `/offers` aliases remain for compatibility; new clients should use `/proposals` and the documented `{data: ...}` envelopes.
- There are no payments, file storage, automatic recommendations, or destructive public reset endpoints.
