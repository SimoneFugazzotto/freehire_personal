# Project Status

## Current state

Phase 1 pilot is complete and operational for the active boards. Thales's Phenom board was crawled
successfully on 2026-09-24: its live listing reported 2,590 rows and the ingest persisted
2,458 distinct Thales vacancies, all with descriptions. Of those, 2,430 are searchable
canons; the other 28 are correctly collapsed reposts. The listing contains repeated
external IDs, so the database count is the meaningful catalogue count.

The current local checkpoint is 7,186 open, canonical, non-private jobs in PostgreSQL
and 7,186 documents in the Meilisearch `jobs` index. The earlier 4,756 index count was
recorded before the Thales crawl and is superseded by this measurement.

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

Start Phase 2 in a new session from the curated 269-company discovery seed. It is intentionally
untrusted and has no assumed career URLs, ATS providers, or board identifiers. Process the
priority-1 cohort in bounded batches: discover the official careers page, detect the ATS,
validate the live board, then add only confirmed boards to the catalog.
