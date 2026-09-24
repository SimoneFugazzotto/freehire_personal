Freehire Personal - Product Vision
1. Missione
Freehire Personal deve diventare un motore di ricerca lavoro personale capace di trovare il maggior numero possibile di opportunità reali e potenzialmente interessanti, senza dipendere da un singolo job board e senza nascondere prematuramente offerte che potrebbero essere valide.
La prima versione è progettata per Simone, inizialmente focalizzata soprattutto su:
- Energy
- Nuclear
- Thermal / thermo-fluid engineering
- CFD
- Heat transfer
- Process engineering
- Plant engineering
- Systems engineering
- Mechanical engineering
- Aerospace
- EPC
- Industrial engineering
- R&D tecnico
In futuro il profilo candidato deve essere sostituibile e il prodotto deve poter funzionare per altri utenti e altri settori.
2. Principio fondamentale
La pipeline deve seguire questo principio:
COLLECT FIRST
NORMALIZE
DEDUPLICATE
UNDERSTAND
RANK
FILTER FOR THE USER
Non:
decidi prima cosa sembra interessante
→ raccogli soltanto quello
Una vacancy poco interessante per Simone oggi può essere utile per un altro profilo domani.
Il catalogo raw deve quindi essere il più possibile domain-agnostic.
3. Due obiettivi distinti
Il sistema deve risolvere due problemi diversi.
A. Company coverage
Non perdere offerte delle aziende che abbiamo deciso di seguire.
Si parte da una lista curata di circa 100-300 aziende target.
Per ciascuna azienda il sistema deve identificare automaticamente:
azienda
→ career site
→ ATS / career platform
→ board
→ vacancy
Se il sito utilizza Workday, Oracle, SuccessFactors, Greenhouse, Lever, Ashby, iCIMS ecc., devono essere riutilizzati gli adapter disponibili.
Se utilizza un career portal proprietario, il sistema deve poterlo trattare separatamente.
B. Market discovery
Scoprire opportunità e aziende che non erano nella lista iniziale.
Le fonti secondarie e il web possono far emergere:
nuova vacancy
→ nuova azienda
→ identificazione dell'azienda
→ ricerca del career site ufficiale
→ discovery della sorgente
→ eventuale ingresso nel catalogo monitorato
Questa parte deve espandere progressivamente l'universo delle aziende.
4. Company Registry
Deve esistere una registry centrale delle aziende.
Una azienda può trovarsi almeno negli stati:
candidate
resolved
monitored
temporarily unavailable
retired
Esempio concettuale:
Company

name
aliases
domains
industry
country

official_website
careers_url

source_provider
board_id

discovery_method
verification_status

last_checked
last_successful_ingest
active_jobs
source_health
Una nuova azienda scoperta automaticamente non deve diventare immediatamente trusted.
Prima deve essere verificato il collegamento con la sorgente corretta.
5. Source priority
Quando possibile, usare questo ordine di preferenza:
1. ATS / career page ufficiale dell'azienda
2. Career portal first-party
3. Fonte ufficiale alternativa
4. LinkedIn
5. HiringCafe
6. Indeed / altri job board
7. Altri aggregatori
Questo ordine riguarda la fonte canonica, non significa che le fonti secondarie debbano essere ignorate.
6. Multi-source jobs
Una vacancy può comparire contemporaneamente su:
career site ufficiale
LinkedIn
HiringCafe
Indeed
altri aggregatori
Non devono diventare cinque lavori separati.
Il sistema deve distinguere:
CanonicalJob
da:
JobSourceRecord
Esempio:
Canonical Job
Thermal Engineer - Company X

Sources:
✓ Company career site
✓ LinkedIn
✓ HiringCafe
La vacancy rimane una sola nell'interfaccia.
La fonte ufficiale deve essere quella preferita per:
Apply
descrizione
status open/closed
canonical URL
ma deve essere conservato il fatto che la vacancy è stata trovata anche altrove.
7. Provenance
Ogni job deve essere tracciabile.
L'utente deve poter sapere:
Official source: Workday
Also found on: LinkedIn, HiringCafe
First discovered through: LinkedIn
Questa informazione è importante sia per trasparenza sia per debugging della coverage.
8. Raccolta geografica
Quando interroghiamo una board aziendale, non filtriamo per geografia durante l'ingestion.
Esempio:
Airbus Workday
→ raccogli tutte le vacancy disponibili
Solo successivamente il motore può mostrare:
Italia
Europa
Torino
Milano
Remote
ecc.
Questo evita di dover modificare la pipeline quando cambiano le preferenze dell'utente.
9. Job data model
Quando disponibile, conservare almeno:
title
company
location raw

city
region
country

remote / hybrid / onsite

description

employment type
seniority

source
source job id
board id

canonical URL
apply URL

published_at
first_seen
last_seen
closed_at

raw source data
Non tutti i provider offriranno tutti i campi.
I dati mancanti non devono impedire l'ingestion.
10. History
Il sistema non deve rappresentare soltanto lo stato corrente.
Deve distinguere:
new
existing
updated
closed
reopened
reposted
Il timestamp pubblicato dal sito non deve essere confuso con:
first_seen
Sono informazioni diverse.
11. Search modes
Devono esistere almeno due esperienze principali.
All Jobs / Exhaustive
Deve permettere di esplorare tutto il catalogo.
Filtri possibili:
company
industry
country
city
location
remote
role family
seniority
publication date
first seen
source
keywords
Una vacancy non deve sparire soltanto perché il matching la considera debole.
Recommended for Simone
Mostra lo stesso catalogo ordinato secondo la compatibilità con Simone.
Recommended deve essere principalmente:
RANKING
non:
HARD FILTER
12. Profilo Simone
Inizialmente esiste un unico profilo candidato.
Non deve essere rappresentato soltanto dal testo completo del CV.
Il sistema deve strutturarlo in dimensioni separate.
Esempio:
Education
Experience
Domains
Technical capabilities
Methods
Software
Programming
Industries
Languages
Seniority
Role preferences
Location preferences
Relocation
Esempio per Simone:
Energy Engineering
Nuclear Engineering

Thermal-hydraulics
CFD
Heat transfer
Fluid systems
Thermo-fluid analysis

Experimental validation
Numerical modelling
Pressure drop analysis

ANSYS CFX
STAR-CCM+
Python

Liquid metals
Nuclear systems
Energy systems

MSc
Junior / Graduate
Il sistema deve quindi essere in grado di capire che un ruolo può essere pertinente anche senza contenere esattamente le parole del CV.
13. Matching
Il matching finale dovrebbe combinare progressivamente:
hard constraints
lexical search
semantic retrieval
profile similarity
skill match
domain match
seniority compatibility
location compatibility
recency
Non deve essere necessario implementare tutto immediatamente.
Si parte dalla soluzione più semplice che funziona sui dati reali e si aggiunge complessità solo quando serve.
14. Explainability
Il ranking non deve ridursi a un numero opaco.
Meglio:
High relevance because:

+ CFD
+ heat transfer
+ thermal-fluid systems
+ ANSYS
+ MSc Engineering

Potential gaps:

- Fluent requested
- 2+ years experience preferred
Uno score numerico può esistere internamente, ma non deve sostituire la spiegazione.
15. Seniority
Non usare regole semplicistiche come:
"3 years requested"
→ elimina automaticamente
Bisogna cercare di distinguere:
required
preferred
nice to have
e considerare l'intero annuncio.
Questo è particolarmente importante per graduate/junior roles.
16. Autonomous discovery
L'espansione autonoma deve essere controllata.
Possibili canali:
secondary job boards
existing vacancy URLs
ATS board discovery
company career links
open web
industry/company directories
companies repeatedly appearing in relevant results
Esempio:
LinkedIn mostra:
Junior Thermal Engineer
XYZ Energy

XYZ Energy non è nella registry

→ candidate company XYZ Energy
→ trova sito ufficiale
→ trova Careers
→ identifica ATS
→ valida board
→ XYZ Energy diventa monitored
Questo meccanismo permette di uscire progressivamente dalla lista iniziale senza trasformare il catalogo in un insieme incontrollato di aziende.
17. Coverage
La coverage deve essere una caratteristica visibile e misurabile.
Per ogni azienda:
source resolved?
source healthy?
jobs active
last crawl
last successful crawl
errors
Per il catalogo:
target companies
resolved companies
monitored companies
unresolved companies

active jobs
canonical jobs

jobs official-source
jobs secondary-only
Non devono essere mostrati numeri che confondono:
numero di adapter supportati
con:
numero di aziende realmente monitorate
18. Freshness
Il sistema deve periodicamente verificare le fonti.
La frequenza potrà dipendere dal provider.
Obiettivo:
nuova vacancy → compare rapidamente
vacancy rimossa → viene marcata closed
Non serve ottenere realtime assoluto.
Affidabilità e sostenibilità sono più importanti di richieste continue.
19. Secondary sources
LinkedIn, HiringCafe, Indeed e altri possono avere tre funzioni:
Discovery
Trovare vacancy che il nostro crawler non ha ancora visto.
Company discovery
Far emergere aziende non presenti nella registry.
Fallback source
Se una vacancy è disponibile solo su una fonte secondaria, può comunque essere mostrata.
Deve però essere chiaramente indicato:
Official source not currently available
Source: LinkedIn
Se successivamente compare la versione ufficiale:
secondary record
+
official record
→ stesso canonical job
20. Application workflow
La prima versione deve concentrarsi su:
DISCOVER
SEARCH
RANK
OPEN JOB
APPLY
Solo successivamente potranno essere aggiunti:
Saved
Interested
Applied
Interview
Rejected
Offer
Notes
Contacts
Follow-up
L'architettura non deve rendere queste funzioni impossibili, ma non sono una priorità della prima fase.
21. Development philosophy
Seguire sempre:
real data before theoretical complexity

reuse existing adapter
before creating custom scraper

simple working pipeline
before perfect architecture

measure the problem
before fixing the problem

coverage
before sophisticated ranking
Non fare refactor grandi senza una necessità concreta.
Roadmap
Phase 1 - End-to-end POC
Circa 10 aziende.
Dimostrare:
Company
→ source discovery
→ board validation
→ ingestion
→ raw DB
→ canonical jobs
→ search
Inclusi ruoli engineering/non-tech.
Phase 2 - Target-company expansion
Preparare una lista iniziale di circa:
100-300 aziende
rilevanti per Simone.
Il sistema deve automaticamente tentare per ciascuna:
website
→ careers
→ ATS
→ board
→ validation
→ ingestion
Le non risolte restano in una coda.
Questa fase è una delle priorità principali del progetto.
Phase 3 - Coverage improvement
Per le aziende irrisolte:
career portal proprietary
dynamic websites
secondary-source clues
search-based discovery
custom adapters when justified
Phase 4 - Autonomous market discovery
Utilizzare le vacancy provenienti dal mercato per scoprire:
nuovi job
nuove aziende
nuove board
e far crescere automaticamente la registry.
Phase 5 - Personalized search
Costruire il profilo strutturato di Simone.
Aggiungere:
Recommended
All Jobs
hybrid search
matching explanations
Phase 6 - Workflow
Successivamente:
saved jobs
application tracking
alerts
notifications
CV variants
application assistance
Success criteria
Il successo non è:
“il sistema mi mostra 20 job molto simili al CV.”

Il successo è:
“il sistema osserva una parte molto ampia del mercato rilevante, perde il minor numero possibile di opportunità e mi aiuta a capire rapidamente quali meritano attenzione.”
