# Qadam Frontend

The Qadam web interface supports business task formulation, AI clarification, published briefs, team proposals, and milestone review.

Built with Vue 3 Composition API, TypeScript, Vite, Pinia, and Vue Router. The interface uses custom CSS, Lucide icons, and a locally bundled Manrope font.

## Start locally

Start PostgreSQL and the Go API using the [project setup guide](../README.md#installation-and-startup). Then run from `frontend`:

```bash
npm ci
npm run dev
```

Use Node.js 22.12+ and npm. Open `http://127.0.0.1:5173`. The development server requires this port to be available rather than silently selecting another port.

## Backend connection

The default data mode is `api`. To override configuration, create `frontend/.env.local`:

```dotenv
VITE_DATA_MODE=api
VITE_API_BASE_URL=/api
API_PROXY_TARGET=http://127.0.0.1:8080
```

Restart Vite after changing environment variables. Its development proxy forwards `/api` to Go, keeping browser requests on the same origin. `API_PROXY_TARGET` only configures the development server; production hosting requires its own API routing.

Never put AI provider credentials in this directory or in `VITE_*` variables. Those values are included in client-side JavaScript. AI requests go through the Go backend.

The HTTP adapter reports backend errors. It does not silently replace unavailable API responses with local sample data. When AI is unavailable, users can continue editing task fields manually.

## Workflows

- **Catalog:** public browsing, search, category and readiness filters, sorting, and empty states.
- **Task builder:** initial description, generated clarification questions, and 16 editable fields arranged in seven groups.
- **Confirmation:** a draft score preview, explicit information confirmation, and a separately stored public snapshot.
- **Publication:** no minimum readiness score; editing a draft does not automatically replace its confirmed snapshot.
- **Proposals:** teams submit an idea, plan, timeframe, and prototype link; businesses accept or reject proposals manually.
- **Milestones:** an accepted team submits a result; business confirmation awards 10 points once.
- **Workspaces:** business tasks and team proposals, selected through demonstration profiles.
- **Interaction support:** validation, visible loading and error states, unsaved-change prompts, keyboard focus styles, reduced motion, and responsive layouts.

The current interface uses Russian text. Documentation is written in English.

## AI assistant and voice input

The task builder requests generated questions through `/api/ai/analyze`. Qadam AI opens an assistant panel and sends chat messages to `/api/ai/chat` with the relevant page and task context.

Voice controls record descriptions or answers using MediaRecorder. Supported browser audio is sent to `/api/ai/transcribe`; the returned text remains editable before use. Voice input requires microphone permission and a secure browser context such as localhost or HTTPS.

Clarification, chat, and transcription require a running backend and a configured live AI provider, including when browser-only data mode is selected. Prepared local questions are not substituted for model output.

See the [AI contract](AI-CONTRACT.md) for payloads and error handling.

## Browser-only data mode

Set `VITE_DATA_MODE=demo` in `frontend/.env.local` and restart Vite to use local sample data without PostgreSQL. The interface explicitly identifies this mode.

Sample data comes from `../docs/fixtures/seed.json`. Local state is stored under `qadam:demo:v1`; the selected profile is stored under `qadam:actor`. A separate browser profile or private window provides a fresh local walkthrough. Removing these keys deletes the corresponding browser-only state.

Example-fill controls are available only in this mode. They do not automatically confirm or publish information. `/demo/prototype.html` is a static sample artifact with synthetic figures, not a result from a real business.

## Example walkthrough

1. Browse the catalog, then select the business profile.
2. Create a task, provide its description, and request clarification.
3. Answer the questions, review the brief, confirm the information, and publish.
4. Switch to a team profile and submit a proposal for the published task.
5. Return to the business workspace and accept or reject the proposal.
6. For an accepted proposal, submit a milestone as the team and confirm it as the business.
7. Check the resulting team points and reload the page to check persistence.
8. Open Qadam AI to ask a question about the current task.

## Source layout

```text
src/pages/          Routes and user workflows
src/components/     Reusable forms, cards, scoring, icons and assistant controls
src/domain/         TypeScript DTOs, field groups and readiness preview
src/services/       Gateway, REST adapter, local data adapter and tests
src/stores/         Selected actor, catalog, teams and notifications
src/composables/    Shared behavior and task-loading coordination
src/data/           Synthetic fixtures and the example brief
```

## Development checks

Run from `frontend`:

```bash
npm test
npm run typecheck
npm run build
npm run format:check
```

The test suite covers publication and scoring, confirmation snapshots, actor rules, proposal decisions, one-time milestone awards, revision conflicts, persistence, AI field handling, and HTTP errors.

Use `npm run format` to apply formatting. After building, `npm run preview` serves the bundle at `http://127.0.0.1:4173`; configure access to the API separately because the development proxy is not a production deployment setup.

## Identity and deployment limits

`/start` and the header select the business or a prepared team profile. Requests use `X-Demo-Actor`; this is a local demonstration identity mechanism, not production authentication. Registration, passwords, JWT sessions, and production access control are not implemented.

A production web server must serve `index.html` for frontend routes such as `/tasks/:id` and `/business`, and route `/api/*` to Go.

See the [API contract](API-CONTRACT.md) and [backend guide](../backend/README.md) for integration details.
