# Qadam first pre-defense materials

**Development of the Qadam Platform for Business Task Formulation Using Artificial Intelligence**

Team: Bekasyl Zhaksylyk, Mereke Myrzakhan, Amantai Zere. Supervisor: Altynbek Seitenov.

This package contains the project proposal, literature review, requirements, architecture, evaluation plan and team work plan for the first pre-defense.

## Materials

- [Assignment criteria and remaining evidence](CRITERIA_CHECKLIST.md).

- [Word report](Qadam_Predefense_Report.docx) and [editable text](REPORT.md).
- [Diagrams](diagrams/): context, use cases, components, ERD, activity, sequence and interface workflow; SVG sources and PNG exports.
- [Reading preparation](LITERATURE_NOTES.md).
- [Contribution and work plan](YOUTRACK_TASKS.md) and [CSV snapshot](BACKLOG.csv).
- [Evidence templates](EVIDENCE_FORMS.md), [traceability](TRACEABILITY.md), [verification](VERIFICATION.md).
- [Presentation and rehearsal guide](PRESENTATION.md).

## Team tracking

[Qadam project](https://bekasyl00.youtrack.cloud/projects/QADAM) · [Kanban board](https://bekasyl00.youtrack.cloud/agiles/204-1/218-1)

Use one accountable owner, a reviewer, target date and acceptance evidence per task. The administrator is currently the only project member; named ownership is recorded in descriptions until accounts are added. A prepared report awaits peer review; study results, interviews and submission receipts must come from actual work.

## Preparation

Read the assigned papers, review the diagrams, confirm individual contributions and agree the evaluation protocol.

## Rebuild

The editable sources are `REPORT.md`, `tools/build_report.py` and `tools/architecture_diagrams.py`. The builder requires Python, python-docx, Pillow and Arial or a compatible configured font.

```bash
python3 -m pip install python-docx pillow
python3 docs/diploma/tools/build_report.py
```

Render and inspect the Word output after editing. Application source and Git history are preserved in this documentation revision.
