# Requirements, implementation and planned acceptance checks

Scope: the current diploma prototype. File paths demonstrate implementation location; they do not establish that every acceptance condition has been executed. Requirements are defined in [REPORT.md](REPORT.md).

| Requirement | Implementation / artifact | Existing test family or planned check |
| --- | --- | --- |
| Demo roles | `frontend/src/components/RoleSwitcher.vue`, `backend/internal/tasks/access.go` | Guest/owner/team access matrix; server and integration tests |
| Private drafts | `frontend/src/pages/TaskEditor.vue`, `backend/internal/tasks/store.go` | Save–reload; non-owner private draft rejection; DB integration |
| Contextual clarification | `frontend/src/services/ai-api.ts`, `backend/internal/ai/` | AI validator tests; live description → 3–5 questions; provider failure |
| Traceable suggestions | `backend/internal/ai/`, `backend/internal/rating/rating.go` | Fabricated quote rejection; negative/contradictory input corpus |
| Editable brief | `frontend/src/pages/TaskEditor.vue` | Manual full workflow with inference unavailable; no draft loss |
| Confirmation and rating | `backend/internal/tasks/`, `backend/internal/rating/rating_test.go` | Exact score fixtures; confirmed/public snapshot separate from draft |
| Public catalog | `frontend/src/pages/Catalog.vue`, `backend/internal/tasks/http.go` | Published-only listing; filters, empty state and pagination/limits if applicable |
| Team proposals | `frontend/src/components/ProposalForm.vue`, `backend/internal/offers/proposals.go` | Proposal validation and reload; HTTP and DB tests |
| Human selection | `frontend/src/pages/BusinessTask.vue`, `backend/internal/offers/` | Owner-only decisions; multiple accepted proposals permitted |
| Milestone awards | `backend/internal/offers/milestones.go`, `backend/internal/offers/integration_test.go` | Repeated confirmation grants exactly one award, including competing requests |
| AI chat | `frontend/src/components/AiAssistant.vue`, `backend/internal/ai/assistant.go` | Inaccessible context excluded; live answer and error state |
| Voice input | `frontend/src/components/VoiceRecorder.vue`, `backend/internal/ai/` | Permission denial, 60-second limit, 15 MiB rejection, editable transcript |
| Revision conflicts | `backend/internal/tasks/store.go`, frontend services | Two clients save same revision; stale save yields 409 and recovery UI |
| Verified authentication | Future verified identity design | No secure authentication claim; acceptance required before public deployment |
| Performance | Planned local load run | 1,000 synthetic tasks; 20 concurrent clients; 200 requests; p95 <1 s target |
| Reliability | AI services and UI error states | Timeout/quota/malformed output without lost drafts or fabricated responses |
| Usability; Explainability | Planned pilot and observation form | Task completion, score comprehension; independent quality assessment |
| Confidentiality | Config, access checks, `.gitignore` | Secret exclusion, browser bundle inspection, authorized context tests |
| Reproducibility | README and verification log | Reproduce from checkout, record tools and configuration without secrets |

## Readiness formula

The current seven groups have maxima 20, 20, 15, 15, 10, 10 and 10. Fields contribute only when eligible and confirmed for the confirmed rating. Success target and acceptance method depend on a confirmed success metric. Bands are priority ≥90, ready ≥70, working ≥40, draft <40. This is a heuristic for confirmed information coverage; it is not a probability of success or an empirically validated quality score.

## Evidence chain to complete

`real participant finding → requirement → real YouTrack issue → task branch/PR → other member review → test evidence → submission tag`.

Existing PLAN IDs are local planning labels. Do not claim the chain exists until the real URLs, author, review and results are recorded. Historical commits retain their original authors and timestamps.
