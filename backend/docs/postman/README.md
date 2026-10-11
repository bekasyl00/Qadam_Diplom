# Qadam API Collection

This directory contains a Postman v3 folder-based collection for the Qadam API. Open `backend/docs` in Postman's Local View. The workspace configuration in `.postman/resources.yaml` registers the collection under [collections](collections/). These YAML files are not JSON import files.

## Prerequisites

Start PostgreSQL, run `go run ./cmd/seed`, and start `go run ./cmd/api` from `backend`. See the [backend setup guide](../../README.md).

The collection's `base_url` defaults to `http://127.0.0.1:8080`. Provider keys belong only in the Go server environment, never in Postman. Live clarification requires a configured AI provider.

## Request sequence

Run the folders in order:

1. Check API health.
2. Inspect business tasks and team profiles; capture the prepared team IDs.
3. Create a draft, request clarification, save answers, assemble and edit the brief, confirm information, and publish a 40-point card.
4. Read the catalog and task details, then submit a team proposal.
5. Accept the first proposal; create and reject a separate second proposal.
6. Submit a milestone result, confirm it as the business, and inspect the team's point award.
7. Complete additional brief fields, reconfirm the card, and inspect its updated score and catalog position.

Response scripts store task, proposal, and milestone IDs and revisions for later requests. Manual field values make the task workflow reproducible without relying on generated content; AI-specific requests still require live provider access. In the application, users must review suggestions before saving or confirming them.

Proposal decisions are final in the current workflow. Repeated milestone confirmation does not award points again. Repeating the full collection creates additional sample tasks and milestones, so compare each milestone's award rather than expecting the team's total to remain exactly 10.

See the [request reference](../api-request-list.md) for payloads and errors.

## Optional collection check

With Node.js and Python with PyYAML installed, run from `backend`:

```bash
node scripts/check-postman.cjs
```

The script sends the YAML requests to the running local API and executes the assertion subset used by the response scripts. It does not automate the Postman interface. It creates sample tasks, proposals, and a milestone, so use a local development instance.

Set `BASE_URL` to use another local port. Remote hosts are rejected by the script. Node.js and Python are needed for this optional check, not for the Go API itself.
