# Project Status

## Current state

Phase 1 pilot is complete and operational for the active boards. Thales's Phenom board was crawled
successfully on 2026-09-24: its live listing reported 2,590 rows and the ingest persisted
2,458 distinct Thales vacancies, all with descriptions. Of those, 2,430 are searchable
canons; the other 28 are correctly collapsed reposts. The listing contains repeated
external IDs, so the database count is the meaningful catalogue count.

The current local checkpoint is 7,566 open jobs, including 7,431 canonical, non-private
jobs in PostgreSQL. Earlier counts (including 7,186) are historical Phase 1 snapshots and
are superseded by the Phase 2 recovery measurement.

Phase 2 has completed the Italy-priority-1 cohort from the immutable 269-company research
seed. Its append-only phase2_ledger.jsonl records each processed target and evidence
without turning the untrusted seed into a second registry. The cohort contains 59
organizations: seven were already covered by existing parent or active boards, and two
new official SuccessFactors sources passed real adapter validation and were onboarded:
De Nora (jobs.denora.com, five jobs) and RINA (careers.rina.org, 233 jobs).

The local checkpoint is now 7,566 open jobs and 7,431 canonical, non-private jobs in
PostgreSQL. The search outbox is empty. A targeted recovery pass added a general
Inrecruiting adapter after verifying its public listing, pagination, and JSON-LD detail
contract across Ansaldo Energia and Bonatti. Bonatti passed live validation, was onboarded,
and added nine searchable jobs. Ansaldo Energia exposes the same public tenant but currently
needs a tenant-level mapping follow-up because the adapter probe did not yield its visible
postings. Ten companies with confirmed official careers pages but no live public vacancies
are represented in the ledger as resolved_no_openings rather than discovery failures.

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
- Investigated Thales end to end. `careers.thalesgroup.com` uses Phenom's current
  `POST /widgets` `refineSearch` and `jobDetail` payloads; the adapter's endpoint,
  locale/path handling, pagination, and response mapping all match the live API.
  The initial zero was an unrun board, not an empty listing or parser failure.
- Hardened every Phenom board against transient detail failures: a failed or malformed
  detail is now recorded as an unreadable marker, while a 404/410 remains a confirmed
  removal. The regression tests cover both outcomes.
- Drained the Thales search outbox and verified 2,430 searchable Thales documents in
  Meilisearch, including real cybersecurity and engineering roles.
- Expanded location normalization from observed Phenom data: native Italian spellings
  `Roma` and `Milano` now resolve to Italy, and `MI` is treated as Michigan only when a
  recognized US/Canadian city corroborates it. This corrected 47 Thales rows previously
  unrecognized or misclassified as US, plus two `Marcallo` rows; 82 Thales vacancies are
  now normalized as Italy and searchable through that facet. Regression tests preserve
  `Detroit, MI` as US.

## Open issues

- Thales's career-site aggregate reports 93 Italian openings, while the 2026-09-24 API
  snapshot now yields 82 normalized `{it}` rows. The sampled loss was caused by location
  normalization and is fixed; the remaining aggregate mismatch needs comparison against a
  fresh Phenom listing snapshot, not a DB/search repair.
- Assystem has no validated SmartRecruiters board. The previous `oneclick-ui` match was
  a resolved false positive; discovering its official careers source remains pending.

## Next macro step

Continue Phase 2 in seed order with Italy priority 2. The seed is intentionally untrusted
and has no assumed career URLs, ATS providers, or board identifiers. Continue bounded
batches: discover the official careers page, detect the ATS, validate the live board,
then add only confirmed boards to the catalog.
