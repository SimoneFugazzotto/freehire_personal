# Project Status

## Current state

Phase 1 pilot is operational for the active boards. As of 2026-09-24, the local
catalogue contains 4,756 open canonical jobs from Airbus, Baker Hughes, GE Aerospace,
GE Vernova, Leonardo, Eni, newcleo, and Technip Energies. Thales is active but currently
returns no jobs.

The Meilisearch `jobs` index contains 4,756 documents, matching the open canonical
Postgres scope exactly.

## Important changes

- Added and crawled the validated Workday pilot boards for Airbus, Baker Hughes, GE
  Aerospace, and GE Vernova, plus the Oracle board for Technip Energies; all are active.
- Changed the search scope so every open, canonical, non-private vacancy is indexed.
  An unresolved category or a temporarily missing description is now a ranking and
  presentation signal, never an exclusion from All Jobs.
- Aligned incremental search drain, full rebuild, link import, and company job counts
  to that same scope.
- Rebuilt the local index and verified real Airbus engineering vacancies without a
  resolved category are searchable.

## Open issues

- Thales's Phenom board is active but has no currently stored jobs; recheck its live
  listing before treating it as a failed source.
- The validated SmartRecruiters Assystem seed has not entered the board catalog because
  the prior validation run did not retain a live board with jobs.

## Next macro step

Prepare the Phase 2 target registry
(100-300 relevant energy, nuclear, engineering, EPC, aerospace, and industrial companies)
with official career URLs and ATS discovery candidates.
