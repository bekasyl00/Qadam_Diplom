# Qadam

**Development of the Qadam Platform for Business Task Formulation Using Artificial Intelligence**

Qadam helps businesses turn an informal request into a structured task that a student team can understand and work on. AI asks clarifying questions and suggests task details; the business reviews the information, publishes the task, and chooses a team.

## Problem and users

A business request often lacks a clear outcome, available data, constraints, or acceptance criteria. Teams need these details before they can propose a realistic solution. Qadam brings task formulation and the first collaboration steps into one application.

- **Business representatives** describe a need, confirm the task details, review proposals, and approve a submitted result.
- **Student teams** explore published tasks, propose an approach, and submit a milestone result.
- **Visitors** browse the public catalog without selecting a profile.

## Features

- Public task catalog with search, category and readiness filters, and sorting.
- Task builder with an initial description, AI clarification, and an editable brief with 16 fields.
- AI-generated questions and field suggestions through OpenAI or Groq.
- Voice recording and transcription for descriptions and answers.
- Qadam AI chat with the current page and accessible task context.
- Separate draft completeness preview and confirmed readiness score from 0 to 100.
- Private drafts, confirmed public snapshots, and PostgreSQL persistence.
- Team proposals with an idea, plan, timeframe, and prototype link.
- Manual acceptance or rejection of proposals, including selection of multiple teams.
- Milestone submission and a one-time award of 10 team points after business confirmation.
- Validation of actor permissions, revision conflicts, and sources for AI suggestions.

The readiness score measures the completeness of confirmed information, not the quality of the idea. A low score does not block publication or proposals. Team points are awarded for a confirmed result, not for submitting a proposal.

## User workflow

1. Select the business profile and describe a task using text or voice.
2. Request AI clarification and answer the generated questions.
3. Review the proposed brief, edit its fields, and confirm the information.
4. Publish the task. The backend calculates the confirmed readiness score.
5. Switch to a student team, find the task, and submit a proposal.
6. As the business, review proposals and choose a team manually.
7. The selected team submits a milestone result; the business confirms it and the team receives points.

Saving draft edits does not automatically replace the confirmed public snapshot. Updated information must be confirmed again.

## Technology

| Component | Implementation |
| --- | --- |
| Frontend | Vue 3, Composition API, TypeScript, Vite, Pinia, Vue Router |
| Interface | Custom CSS, Lucide icons, locally bundled Manrope font |
| Audio capture | Browser MediaRecorder API |
| Backend | Go 1.23+, `net/http`, `database/sql` |
| Database | PostgreSQL 16, `lib/pq`, embedded SQL migrations |
| Configuration | Environment variables and `godotenv` |
| AI questions and chat | OpenAI Responses API or Groq's OpenAI-compatible API |
| Default text models | OpenAI `gpt-4.1-mini`; Groq `openai/gpt-oss-20b` |
| Default transcription models | OpenAI `gpt-4o-mini-transcribe`; Groq `whisper-large-v3-turbo` |
| Local infrastructure | Docker Compose, optional pgAdmin |
| Verification tools | Vitest, vue-tsc, Prettier, Go tests and PostgreSQL integration tests |

Models are configurable. Live AI requires network access and a provider account with access to the selected models and sufficient quota. Provider credentials stay on the backend.

## Architecture

```text
Browser: Vue + Pinia
        |
        | REST /api + selected demonstration profile
        v
Vite development proxy --> Go API
                              |-- PostgreSQL: drafts, snapshots, proposals, milestones, points
                              |-- OpenAI / Groq: clarification, brief suggestions, chat
                              `-- Transcription API: audio to editable text
```

```text
frontend/
  src/pages/                  Catalog, task builder and workspaces
  src/components/             Forms, scoring, assistant and audio controls
  src/services/               REST adapter and browser-only data adapter
  src/stores/                 Shared application state
backend/
  cmd/api/                    API server
  cmd/seed/                   Additive sample-data loader
  cmd/check-ai/               Live AI configuration check
  internal/ai/                Provider requests and output validation
  internal/tasks/             Task and draft workflows
  internal/offers/            Proposals and milestone workflows
  internal/rating/            Confirmed readiness calculation
  internal/teams/             Team profiles and progress
  internal/database/          PostgreSQL connection and migrations
docs/                         Specifications, fixtures and technical records
```

The frontend uses the Go API by default. An explicit browser-only mode stores sample data in localStorage. API failures do not silently switch the application to local data.

## Installation and startup

### Prerequisites

- Go 1.23 or newer.
- Node.js 22.12+ and npm.
- Docker with Compose; start Docker Desktop first on macOS or Windows.
- An OpenAI or Groq API key for live AI features.

The commands below use Bash or Zsh. Windows users can use the PowerShell instructions in the [backend guide](backend/README.md).

### 1. Clone the repository

```bash
git clone https://github.com/bekasyl00/Qadam_Diplom.git
cd Qadam_Diplom
```

### 2. Start PostgreSQL and configure the backend

```bash
cd backend
docker compose up -d --wait postgres
[ -f .env ] || cp .env.example .env
go mod download
```

Open `backend/.env` in an editor. Keep the database connection from `.env.example` for the supplied Compose setup. The API uses port `8080`; PostgreSQL is exposed locally on port `5433`.

For **Groq**, add or update:

```dotenv
AI_PROVIDER=groq
GROQ_API_KEY=replace_with_your_key
AI_MODEL=openai/gpt-oss-20b
AI_TRANSCRIBE_MODEL=whisper-large-v3-turbo
AI_MODE=live
DEMO_MODE=true
```

Alternatively, for **OpenAI**:

```dotenv
AI_PROVIDER=openai
OPENAI_API_KEY=replace_with_your_key
OPENAI_MODEL=gpt-4.1-mini
OPENAI_TRANSCRIBE_MODEL=gpt-4o-mini-transcribe
AI_MODE=live
DEMO_MODE=true
```

Choose one provider. `.env` is excluded from Git. Never put provider keys in frontend files or `VITE_*` variables, which are exposed to the browser.

To run without an AI key, leave `AI_MODE=fallback`. Manual task editing, publication, proposals, and milestones remain available. AI endpoints report unavailability instead of presenting prepared questions as generated output.

### 3. Seed data and start the API

From `backend`:

```bash
go run ./cmd/seed
go run ./cmd/api
```

Migrations run automatically. Seeding adds missing sample records and preserves existing edits.

Check [http://127.0.0.1:8080/api/health](http://127.0.0.1:8080/api/health). To check live AI separately, run the following from another terminal in `backend`:

```bash
go run ./cmd/check-ai
```

This sends a real provider request and may consume quota. Restart the API after changing `.env`. Existing shell environment variables take precedence over `.env` values.

### 4. Start the frontend

From the repository root in a second terminal:

```bash
cd frontend
npm ci
npm run dev
```

Open **[http://127.0.0.1:5173](http://127.0.0.1:5173)**. Vite forwards `/api` to the backend on port `8080`.

For a different backend address, create `frontend/.env.local` using `.env.example`, update `API_PROXY_TARGET`, and restart Vite. For browser-only sample data, explicitly set `VITE_DATA_MODE=demo`; AI still requires the Go API and a configured provider.

## Example walkthrough

The current interface uses Russian labels. Follow the corresponding controls for each action below.

1. Browse the catalog as a visitor and try the search and filters.
2. Select the business profile and create a task such as **Store purchase planning**.
3. Enter: “Our store keeps sales and stock levels in CSV files. The manager spends one hour preparing a purchase list. We need a web prototype in 14 days that prepares the list in 10 minutes.”
4. Request AI questions, answer them, and review the suggested brief. Use manual editing if AI is unavailable.
5. Optionally record a short description, check the transcription, and edit the text before continuing.
6. Confirm the information and publish the task. Inspect its confirmed score and catalog entry.
7. Switch to a student team and submit a proposal. The included `/demo/prototype.html` page can serve as an explicitly synthetic example link.
8. Switch back to the business profile and accept the proposal.
9. Submit a milestone as the team, then confirm it as the business. Check that the team receives 10 points and that the result survives a page reload.
10. Open Qadam AI and ask which information the team needs to start this task.

## Checks

Frontend, from `frontend`:

```bash
npm test
npm run typecheck
npm run build
npm run format:check
```

Backend, from `backend`:

```bash
go test ./...
go vet ./...
go build ./...
```

Database integration tests require `TEST_DATABASE_URL` pointing to a local test database. Use the connection settings from `backend/.env.example` when running against the supplied local PostgreSQL service:

```bash
TEST_DATABASE_URL='<local_postgresql_connection_string>' go test ./... -count=1 -timeout=90s
```

These tests create temporary schemas rather than resetting application data. Without the variable, database integration tests skip. Provider unit tests use local HTTP fakes; `check-ai` tests real provider access separately.

## Data and integrations

- `docs/fixtures/seed.json` contains synthetic drafts, published briefs, teams, and proposals.
- PostgreSQL stores application state across restarts.
- OpenAI or Groq receives the submitted task text, answers, accessible chat context, or audio needed for the requested AI operation.
- OpenAI Responses requests use `store=false`. The application does not retain the original uploaded audio; provider-side retention is governed by the provider's own settings and terms.
- CSV is an example of business input described in a task. The platform does not implement business CSV processing, CRM connections, or payment integrations.

## Current limitations

- The current prototype uses selectable demonstration profiles. It does not implement production authentication, registration, password recovery, or JWT sessions. `X-Demo-Actor` is not proof of identity.
- The API binds to loopback for local use. Setting `DEMO_MODE=false` disables application routes; it does not enable an alternative authentication system.
- AI suggestions require human review. AI does not confirm business facts, select teams, or award points.
- Voice input is record-and-transcribe, not a continuous voice conversation. Browser recordings are limited to 60 seconds and server uploads to 15 MiB. Microphone access requires permission and localhost or HTTPS.
- Each accepted proposal supports one confirmable milestone. Payments, file storage, and a complete project-management lifecycle are not implemented.
- AI availability and latency depend on connectivity, model access, and quota.
- The proposed diploma evaluation has not yet established improvements in brief quality or completion time.

## Diploma project

The project studies how AI can assist the transformation of informal business needs into structured, human-confirmed tasks. The planned evaluation compares AI-assisted task formulation with manual completion of the same brief.

**Team members**

- Bekasyl Zhaksylyk — frontend, UX, integration and demonstration.
- Mereke Myrzakhan — ongoing AI development, prompts, provider behavior and AI evaluation.
- Amantai Zere — backend, database integrity and API tests.

**Supervisor:** Altynbek Seitenov

**Current status:** runnable local prototype and first pre-defense preparation. The revised proposal, architecture and evaluation protocol are prepared for team review; real-user evidence and comparative experiments are still outstanding. Mereke's current AI changes are in progress outside the repository and must be published and reviewed separately.

**Expected outcomes:** a reproducible platform, a human-confirmed structured brief and an evidence-based comparison of assisted versus manual task formulation. Improvements in time or quality are hypotheses, not established results.

[YouTrack project](https://bekasyl00.youtrack.cloud/projects/QADAM) · [Kanban board](https://bekasyl00.youtrack.cloud/agiles/204-1/218-1) · [Pre-defense documents](docs/diploma/README.md) · [Criteria checklist](docs/diploma/CRITERIA_CHECKLIST.md)

Use individual accounts and issue-linked task branches. Submit a pull request with validation evidence and obtain another team member's review before merging to `main`. The submission form requires [`prob-period-1`](https://github.com/bekasyl00/Qadam_Diplom/tree/prob-period-1); keep this tag fixed after submission. The [consolidated PDF](docs/diploma/SE-2418_Zhaksylyk_ProbPeriod1.pdf) and [Word report](docs/diploma/Qadam_Predefense_Report.docx) describe this checkpoint. Publication does not establish independent review or completed study evidence.

## Deployment

No verified public deployment URL is recorded in the repository. Use the local setup above. `npm run build` generates the frontend production bundle; hosting it requires a separately configured API, `/api` routing, and SPA fallback routing.

## Documentation

- [Frontend setup and workflows](frontend/README.md)
- [Backend setup and configuration](backend/README.md)
- [Postman collection guide](backend/docs/postman/README.md)
- [Frontend API contract](frontend/API-CONTRACT.md)
- [AI, chat, and voice contract](frontend/AI-CONTRACT.md)
