# IDX — Índice e inventario documental

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: IDX · Versión: 0.3 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner: C.

## Estado de la entrega

El inventario aprobado conserva **81 archivos Markdown: 40 especificaciones y 41 documentos de soporte**. El lote 1 (31 Markdown) está APROBADO e integrado en main por PR #1. La entrega del lote 2 añade exactamente diez specs EN_REVISIÓN y actualiza únicamente PLANS, IDX, TRZ y CONF; sus nuevas versiones de integración también están EN_REVISIÓN, sin revocar la aprobación histórica. Quedan 40 rutas futuras de lotes 3–6. La copia SRC conserva literalmente su estado histórico; las 246 decisiones fuente permanecen APROBADO. Esta cobertura documental no acredita implementación.

Inicio recomendado: [README](../README.md), [objetivos](../GOALS.md), [plan y reporte](../PLANS.md), [registro](decision-register.md), [trazabilidad](traceability.md), [conflictos](conflicts.md), [norma 11.01](specs/11-orchestration/11.01-document-standard-and-change-control.md) y [orquestación 11.02](specs/11-orchestration/11.02-ownership-delegation-and-context.md).

## Convenciones

Distribución vigente: `main` contiene los 31 Markdown aprobados de lote 1 y `.gitignore`; `docs/specs-lote-02` añade diez specs y cuatro registros actualizados mediante PR sin fusionar. `.gitignore` es la excepción operativa autorizada, ajena al inventario Markdown. La autorización y el informe actual están en [PLANS](../PLANS.md#informe-lote-2).

B/M/A: complejidad relativa baja/media/alta, sin fechas prometidas. Owners funcionales C coordinación/producto; A arquitectura/operación; M motor; S servicios/identidad/multiplayer; V UX/3D; I IA/aprendizaje; Q calidad. El autor efectivo y las revisiones de esta entrega constan en PLANS.

Los vínculos de lotes 1 y 2 apuntan a archivos entregados. Las rutas futuras se muestran como texto y se enlazan mediante el anchor de esta entrada, nunca mediante un vínculo que aparente un archivo existente. Estado futuro `PENDIENTE_LOTE_N` no equivale a deferimiento aprobado de la funcionalidad.

<a id="inventario"></a>

## Inventario autorizado

### Lote 1 — 31 archivos

| ID | Archivo o ruta prevista | Propósito / entradas y dependencias | Owner | Complejidad | Estado |
|---|---|---|---|---|---|
| <a id="agents"></a>AGENTS | [AGENTS.md](../AGENTS.md) | Instrucciones operativas y límites | C | M | APROBADO |
| <a id="goals"></a>GOALS | [GOALS.md](../GOALS.md) | Objetivos y alcance comprometido | C | M | APROBADO |
| <a id="plans"></a>PLANS | [PLANS.md](../PLANS.md) | Gates, lotes y roadmap documental | C | M | EN_REVISIÓN |
| <a id="prompts"></a>PROMPTS | [PROMPTS.md](../PROMPTS.md) | Briefs y prompts de trabajo autorizable | C | M | APROBADO |
| <a id="readme"></a>README | [README.md](../README.md) | Presentación y estado profesional | C | M | APROBADO |
| <a id="src"></a>SRC | [docs/sources/TOGUZ_PROMPT_MAESTRO_CODEX.md](sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) | Copia íntegra y sin cambios del maestro | C | B | APROBADO |
| <a id="reg"></a>REG | [docs/decision-register.md](decision-register.md) | 246 decisiones literales y destinos canónicos | C | M | APROBADO |
| <a id="conf"></a>CONF | [docs/conflicts.md](conflicts.md) | Riesgos, discrepancias y validaciones pendientes | C | M | EN_REVISIÓN |
| <a id="idx"></a>IDX | [docs/index.md](index.md) | Inventario y dependencias | C | B | EN_REVISIÓN |
| <a id="trz"></a>TRZ | [docs/traceability.md](traceability.md) | Decisión→spec→fase→componente→verificación→evidencia | C | M | EN_REVISIÓN |
| <a id="pre-01"></a>PRE-01 | [docs/preliminares/PRE-01.md](preliminares/PRE-01.md) | Contexto de 17 decisiones de REG; sin reconstruir cuestionarios. | A | M | APROBADO |
| <a id="pre-02"></a>PRE-02 | [docs/preliminares/PRE-02.md](preliminares/PRE-02.md) | Contexto de 18 decisiones de REG; sin reconstruir cuestionarios. | C | M | APROBADO |
| <a id="pre-03"></a>PRE-03 | [docs/preliminares/PRE-03.md](preliminares/PRE-03.md) | Contexto de 22 decisiones de REG; sin reconstruir cuestionarios. | M | M | APROBADO |
| <a id="pre-04"></a>PRE-04 | [docs/preliminares/PRE-04.md](preliminares/PRE-04.md) | Contexto de 25 decisiones de REG; sin reconstruir cuestionarios. | V | M | APROBADO |
| <a id="pre-05"></a>PRE-05 | [docs/preliminares/PRE-05.md](preliminares/PRE-05.md) | Contexto de 24 decisiones de REG; sin reconstruir cuestionarios. | S | M | APROBADO |
| <a id="pre-06"></a>PRE-06 | [docs/preliminares/PRE-06.md](preliminares/PRE-06.md) | Contexto de 20 decisiones de REG; sin reconstruir cuestionarios. | S | M | APROBADO |
| <a id="pre-07"></a>PRE-07 | [docs/preliminares/PRE-07.md](preliminares/PRE-07.md) | Contexto de 24 decisiones de REG; sin reconstruir cuestionarios. | S | M | APROBADO |
| <a id="pre-08"></a>PRE-08 | [docs/preliminares/PRE-08.md](preliminares/PRE-08.md) | Contexto de 24 decisiones de REG; sin reconstruir cuestionarios. | I | M | APROBADO |
| <a id="pre-09"></a>PRE-09 | [docs/preliminares/PRE-09.md](preliminares/PRE-09.md) | Contexto de 24 decisiones de REG; sin reconstruir cuestionarios. | Q | M | APROBADO |
| <a id="pre-10"></a>PRE-10 | [docs/preliminares/PRE-10.md](preliminares/PRE-10.md) | Contexto de 24 decisiones de REG; sin reconstruir cuestionarios. | A | M | APROBADO |
| <a id="pre-11"></a>PRE-11 | [docs/preliminares/PRE-11.md](preliminares/PRE-11.md) | Contexto de 24 decisiones de REG; sin reconstruir cuestionarios. | C | M | APROBADO |
| <a id="adr-0001"></a>ADR-0001 | [docs/adr/ADR-0001-three-go-services-and-data-ownership.md](adr/ADR-0001-three-go-services-and-data-ownership.md) | Decisión arquitectónica existente; REG y PRE pertinentes; alternativas y validación futura. | A | M | APROBADO |
| <a id="adr-0002"></a>ADR-0002 | [docs/adr/ADR-0002-single-versioned-go-engine.md](adr/ADR-0002-single-versioned-go-engine.md) | Decisión arquitectónica existente; REG y PRE pertinentes; alternativas y validación futura. | M | M | APROBADO |
| <a id="adr-0003"></a>ADR-0003 | [docs/adr/ADR-0003-rest-contracts-and-versioned-events.md](adr/ADR-0003-rest-contracts-and-versioned-events.md) | Decisión arquitectónica existente; REG y PRE pertinentes; alternativas y validación futura. | S | M | APROBADO |
| <a id="adr-0004"></a>ADR-0004 | [docs/adr/ADR-0004-postgres-authority-and-postcommit-ably.md](adr/ADR-0004-postgres-authority-and-postcommit-ably.md) | Decisión arquitectónica existente; REG y PRE pertinentes; alternativas y validación futura. | S | M | APROBADO |
| <a id="adr-0005"></a>ADR-0005 | [docs/adr/ADR-0005-session-and-service-trust-boundaries.md](adr/ADR-0005-session-and-service-trust-boundaries.md) | Decisión arquitectónica existente; REG y PRE pertinentes; alternativas y validación futura. | S | M | APROBADO |
| <a id="adr-0006"></a>ADR-0006 | [docs/adr/ADR-0006-vercel-hobby-and-zero-cost-guardrails.md](adr/ADR-0006-vercel-hobby-and-zero-cost-guardrails.md) | Decisión arquitectónica existente; REG y PRE pertinentes; alternativas y validación futura. | A | M | APROBADO |
| <a id="adr-0007"></a>ADR-0007 | [docs/adr/ADR-0007-3d-and-accessible-html-view.md](adr/ADR-0007-3d-and-accessible-html-view.md) | Decisión arquitectónica existente; REG y PRE pertinentes; alternativas y validación futura. | V | M | APROBADO |
| <a id="adr-0008"></a>ADR-0008 | [docs/adr/ADR-0008-ranked-results-and-elo-ownership.md](adr/ADR-0008-ranked-results-and-elo-ownership.md) | Decisión arquitectónica existente; REG y PRE pertinentes; alternativas y validación futura. | S | M | APROBADO |
| <a id="spec-11-01"></a>11.01 | [docs/specs/11-orchestration/11.01-document-standard-and-change-control.md](specs/11-orchestration/11.01-document-standard-and-change-control.md) | Norma documental y control de cambios; entradas PRE-11.A1, PRE-11.A2, PRE-11.A3, PRE-11.A4, PRE-11.B1, PRE-11.B2, PRE-11.B3, PRE-11.B4, PRE-11.D4, PRE-11.E1, PRE-11.E4, PRE-11.F1, PRE-11.F3. Dependencias: REG y maestro. | C | A | APROBADO |
| <a id="spec-11-02"></a>11.02 | [docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md](specs/11-orchestration/11.02-ownership-delegation-and-context.md) | Propiedad, delegación y contexto; entradas PRE-11.C1, PRE-11.C2, PRE-11.C3, PRE-11.C4. Dependencias: [11.01](#spec-11-01). | C | M | APROBADO |

### Lote 2 — 10 archivos

| ID | Archivo o ruta prevista | Propósito / entradas y dependencias | Owner | Complejidad | Estado |
|---|---|---|---|---|---|
| <a id="spec-01-01"></a>01.01 | [docs/specs/01-infrastructure/01.01-stack-monorepo.md](specs/01-infrastructure/01.01-stack-monorepo.md) | Stack, monorepo y desarrollo local; entradas PRE-01.A1, PRE-01.A2, PRE-01.A3, PRE-01.A4, PRE-01.C3. Dependencias: REG y maestro. | A | A | EN_REVISIÓN |
| <a id="spec-01-02"></a>01.02 | [docs/specs/01-infrastructure/01.02-platform-feasibility.md](specs/01-infrastructure/01.02-platform-feasibility.md) | Factibilidad de plataforma y entornos; entradas PRE-01.B2, PRE-01.C1, PRE-01.C4, PRE-05.F4. Dependencias: [01.01](#spec-01-01). | A | A | EN_REVISIÓN |
| <a id="spec-01-03"></a>01.03 | [docs/specs/01-infrastructure/01.03-cost-capacity-evolution.md](specs/01-infrastructure/01.03-cost-capacity-evolution.md) | Coste, capacidad y evolución; entradas PRE-01.C2, PRE-01.D3. Dependencias: [01.02](#spec-01-02). | A | A | EN_REVISIÓN |
| <a id="spec-02-01"></a>02.01 | [docs/specs/02-product/02.01-scope-and-modes.md](specs/02-product/02.01-scope-and-modes.md) | Alcance, modos y MVP; entradas PRE-02.A1, PRE-02.A2, PRE-02.A3, PRE-02.A4, PRE-02.E3. Dependencias: REG y maestro. | C | M | EN_REVISIÓN |
| <a id="spec-02-02"></a>02.02 | [docs/specs/02-product/02.02-local-flows-and-i18n.md](specs/02-product/02.02-local-flows-and-i18n.md) | Flujos locales e internacionalización; entradas PRE-02.B3, PRE-02.D3, PRE-05.C4. Dependencias: [02.01](#spec-02-01). | C | M | EN_REVISIÓN |
| <a id="spec-02-03"></a>02.03 | [docs/specs/02-product/02.03-accounts-history-and-licensing.md](specs/02-product/02.03-accounts-history-and-licensing.md) | Cuentas, historial y licencia; entradas PRE-01.D4, PRE-02.B1, PRE-02.B2, PRE-02.D2. Dependencias: [02.01](#spec-02-01). | C | M | EN_REVISIÓN |
| <a id="spec-03-01"></a>03.01 | [docs/specs/03-game-engine/03.01-normative-rules-and-corpus.md](specs/03-game-engine/03.01-normative-rules-and-corpus.md) | Reglamento normativo y corpus; entradas PRE-03.A3, PRE-03.C1, PRE-03.E1. Dependencias: [02.01](#spec-02-01). | M | A | EN_REVISIÓN |
| <a id="spec-03-02"></a>03.02 | [docs/specs/03-game-engine/03.02-state-indices-and-invariants.md](specs/03-game-engine/03.02-state-indices-and-invariants.md) | Estado, índices e invariantes; entradas PRE-03.A4, PRE-03.B1, PRE-03.B4, PRE-03.E2. Dependencias: [01.01](#spec-01-01), [03.01](#spec-03-01). | M | A | EN_REVISIÓN |
| <a id="spec-03-03"></a>03.03 | [docs/specs/03-game-engine/03.03-moves-and-semantic-events.md](specs/03-game-engine/03.03-moves-and-semantic-events.md) | Movimientos y eventos semánticos; entradas PRE-02.C1, PRE-03.B2, PRE-03.B3, PRE-03.B5, PRE-03.C2, PRE-03.C3. Dependencias: [03.02](#spec-03-02). | M | A | EN_REVISIÓN |
| <a id="spec-03-04"></a>03.04 | [docs/specs/03-game-engine/03.04-outcomes-repetition-and-replay.md](specs/03-game-engine/03.04-outcomes-repetition-and-replay.md) | Resultados, repetición y reproducción; entradas PRE-02.C4, PRE-02.D1, PRE-03.A1, PRE-03.A2, PRE-03.C4, PRE-03.E3. Dependencias: [03.03](#spec-03-03). | M | A | EN_REVISIÓN |

### Lote 3 — 11 archivos

| ID | Archivo o ruta prevista | Propósito / entradas y dependencias | Owner | Complejidad | Estado |
|---|---|---|---|---|---|
| <a id="spec-05-01"></a>05.01 | `specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md` | Límites de servicios y diseño hexagonal; entradas PRE-05.A1, PRE-05.A2, PRE-05.A3, PRE-05.A4, PRE-05.D1. Dependencias: [01.02](#spec-01-02), [03.04](#spec-03-04). | S | A | PENDIENTE_LOTE_3 |
| <a id="spec-05-02"></a>05.02 | `specs/05-backend/05.02-http-openapi-and-errors.md` | HTTP, OpenAPI y errores; entradas PRE-05.B1, PRE-05.B2, PRE-05.B3, PRE-05.B4. Dependencias: [05.01](#spec-05-01), [07.01](#spec-07-01), [07.02](#spec-07-02). | S | A | PENDIENTE_LOTE_3 |
| <a id="spec-05-03"></a>05.03 | `specs/05-backend/05.03-transactions-data-and-migrations.md` | Transacciones, datos y migraciones; entradas PRE-01.B3, PRE-01.B4, PRE-05.C1, PRE-05.C2, PRE-05.C3, PRE-05.D2, PRE-05.D3, PRE-05.D4. Dependencias: [05.02](#spec-05-02), [03.04](#spec-03-04). | S | A | PENDIENTE_LOTE_3 |
| <a id="spec-05-04"></a>05.04 | `specs/05-backend/05.04-event-schemas-and-postcommit-publication.md` | Eventos y publicación posterior al commit; entradas PRE-05.E1, PRE-05.E2, PRE-05.E3, PRE-05.E4. Dependencias: [05.03](#spec-05-03), [07.02](#spec-07-02). | S | A | PENDIENTE_LOTE_3 |
| <a id="spec-06-01"></a>06.01 | `specs/06-multiplayer/06.01-rooms-invitations-and-ranked-lifecycle.md` | Salas, invitaciones y ciclo ranked; entradas PRE-06.A1, PRE-06.A2, PRE-06.A3, PRE-06.A4, PRE-06.D1, PRE-06.D4. Dependencias: [05.04](#spec-05-04), [07.04](#spec-07-04). | S | A | PENDIENTE_LOTE_3 |
| <a id="spec-06-02"></a>06.02 | `specs/06-multiplayer/06.02-ably-notifications-and-reconciliation.md` | Ably y reconciliación; entradas PRE-01.B1, PRE-06.B1, PRE-06.B2, PRE-06.B3, PRE-06.B4, PRE-06.E1. Dependencias: [06.01](#spec-06-01), [05.04](#spec-05-04). | S | A | PENDIENTE_LOTE_3 |
| <a id="spec-06-03"></a>06.03 | `specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md` | Recuperación, pestañas y fallos parciales; entradas PRE-02.C3, PRE-06.C1, PRE-06.C2, PRE-06.C3, PRE-06.C4, PRE-06.E2. Dependencias: [06.02](#spec-06-02), [07.01](#spec-07-01), [07.02](#spec-07-02). | S | A | PENDIENTE_LOTE_3 |
| <a id="spec-07-01"></a>07.01 | `specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md` | OIDC, invitados y ciclo de sesión; entradas PRE-01.B5, PRE-07.A1, PRE-07.A2, PRE-07.A3, PRE-07.A4, PRE-07.B2, PRE-07.B3, PRE-07.F1. Dependencias: [05.01](#spec-05-01), [02.03](#spec-02-03). | S | A | PENDIENTE_LOTE_3 |
| <a id="spec-07-02"></a>07.02 | `specs/07-identity-data/07.02-tokens-authorization-and-service-identity.md` | Tokens, autorización e identidad de servicio; entradas PRE-05.F3, PRE-07.B1, PRE-07.B4, PRE-07.C1, PRE-07.C2, PRE-07.C3. Dependencias: [07.01](#spec-07-01). | S | A | PENDIENTE_LOTE_3 |
| <a id="spec-07-03"></a>07.03 | `specs/07-identity-data/07.03-profiles-history-and-data-rights.md` | Perfiles, historial y derechos de datos; entradas PRE-07.C4, PRE-07.D1, PRE-07.D2, PRE-07.D3, PRE-07.D4, PRE-07.F2, PRE-07.F3. Dependencias: [05.04](#spec-05-04), [07.02](#spec-07-02). | S | A | PENDIENTE_LOTE_3 |
| <a id="spec-07-04"></a>07.04 | `specs/07-identity-data/07.04-elo-leaderboard-and-statistics.md` | Elo, clasificación y estadísticas; entradas PRE-06.D2, PRE-06.D3, PRE-07.E1, PRE-07.E2, PRE-07.E3, PRE-07.E4. Dependencias: [05.03](#spec-05-03), [03.04](#spec-03-04), [07.02](#spec-07-02). | S | A | PENDIENTE_LOTE_3 |

### Lote 4 — 8 archivos

| ID | Archivo o ruta prevista | Propósito / entradas y dependencias | Owner | Complejidad | Estado |
|---|---|---|---|---|---|
| <a id="spec-04-01"></a>04.01 | `specs/04-visual-experience/04.01-navigation-and-accessibility.md` | Navegación y accesibilidad; entradas PRE-04.A4, PRE-04.D1, PRE-04.D2, PRE-04.D3, PRE-04.D4. Dependencias: [02.02](#spec-02-02), [02.03](#spec-02-03), [03.03](#spec-03-03), [05.02](#spec-05-02), [06.01](#spec-06-01), [06.03](#spec-06-03). | V | A | PENDIENTE_LOTE_4 |
| <a id="spec-04-02"></a>04.02 | `specs/04-visual-experience/04.02-board-assets-camera-and-materials.md` | Tablero, assets, cámara y materiales; entradas PRE-02.E1, PRE-02.E2, PRE-04.A1, PRE-04.A2, PRE-04.A3, PRE-04.B1, PRE-04.B2, PRE-04.B3, PRE-04.B4, PRE-04.B5, PRE-04.F2, PRE-04.F3. Dependencias: [04.01](#spec-04-01), [03.02](#spec-03-02). | V | A | PENDIENTE_LOTE_4 |
| <a id="spec-04-03"></a>04.03 | `specs/04-visual-experience/04.03-interaction-animation-and-audio.md` | Interacción, animación y audio; entradas PRE-04.C1, PRE-04.C2, PRE-04.C3, PRE-04.C4. Dependencias: [04.01](#spec-04-01), [03.03](#spec-03-03), [06.02](#spec-06-02). | V | A | PENDIENTE_LOTE_4 |
| <a id="spec-04-04"></a>04.04 | `specs/04-visual-experience/04.04-performance-and-html-fallback.md` | Rendimiento y alternativa HTML; entradas PRE-04.E1, PRE-04.E2, PRE-04.E3, PRE-04.E4. Dependencias: [04.02](#spec-04-02), [04.03](#spec-04-03). | V | A | PENDIENTE_LOTE_4 |
| <a id="spec-08-01"></a>08.01 | `specs/08-ai-practice/08.01-search-heuristics-and-compute-limits.md` | Búsqueda, heurística y límites IA; entradas PRE-03.D3, PRE-03.D4, PRE-03.D5, PRE-08.A2, PRE-08.A4, PRE-08.B1, PRE-08.B2, PRE-08.B3, PRE-08.B4. Dependencias: [03.04](#spec-03-04), [01.02](#spec-01-02). | I | A | PENDIENTE_LOTE_4 |
| <a id="spec-08-02"></a>08.02 | `specs/08-ai-practice/08.02-ai-orchestration-fallback-and-metrics.md` | Orquestación IA, respaldo y métricas; entradas PRE-08.A1, PRE-08.A3, PRE-08.F3, PRE-08.F4. Dependencias: [08.01](#spec-08-01), [05.02](#spec-05-02), [05.03](#spec-05-03), [07.02](#spec-07-02). | I | A | PENDIENTE_LOTE_4 |
| <a id="spec-08-03"></a>08.03 | `specs/08-ai-practice/08.03-reversible-practice-and-scenarios.md` | Práctica reversible y escenarios; entradas PRE-02.C2, PRE-03.D1, PRE-03.D2, PRE-08.C1, PRE-08.C2, PRE-08.C3, PRE-08.C4, PRE-08.E1. Dependencias: [02.02](#spec-02-02), [03.04](#spec-03-04), [05.02](#spec-05-02). | I | A | PENDIENTE_LOTE_4 |
| <a id="spec-08-04"></a>08.04 | `specs/08-ai-practice/08.04-tutorial-hints-and-educational-replay.md` | Tutorial, sugerencias y replay educativo; entradas PRE-02.B4, PRE-04.D5, PRE-08.D1, PRE-08.D2, PRE-08.D3, PRE-08.D4, PRE-08.E2, PRE-08.E3, PRE-08.E4. Dependencias: [08.03](#spec-08-03), [03.03](#spec-03-03). | I | A | PENDIENTE_LOTE_4 |

### Lote 5 — 20 archivos

| ID | Archivo o ruta prevista | Propósito / entradas y dependencias | Owner | Complejidad | Estado |
|---|---|---|---|---|---|
| <a id="spec-09-01"></a>09.01 | `specs/09-quality/09.01-engine-data-contract-and-ai-verification.md` | Verificación de motor, datos, contratos e IA; entradas PRE-05.F2, PRE-07.F4, PRE-08.F1, PRE-08.F2, PRE-09.A1, PRE-09.A2, PRE-09.A3, PRE-09.A4, PRE-09.B1, PRE-09.B2, PRE-09.B3, PRE-09.B4, PRE-09.C1, PRE-09.C2, PRE-09.C3, PRE-09.C4. Dependencias: [05.04](#spec-05-04), [07.04](#spec-07-04), [08.02](#spec-08-02). | Q | A | PENDIENTE_LOTE_5 |
| <a id="spec-09-02"></a>09.02 | `specs/09-quality/09.02-ui-e2e-accessibility-and-visual-verification.md` | Verificación UI, E 2 E, accesibilidad y visual; entradas PRE-04.F1, PRE-06.E3, PRE-09.D1, PRE-09.D2, PRE-09.D3, PRE-09.D4. Dependencias: [04.04](#spec-04-04), [06.03](#spec-06-03), [08.04](#spec-08-04). | Q | A | PENDIENTE_LOTE_5 |
| <a id="spec-09-03"></a>09.03 | `specs/09-quality/09.03-ci-security-and-release-controls.md` | CI, seguridad y controles de release; entradas PRE-01.D1, PRE-09.E1, PRE-09.E2, PRE-09.E3, PRE-09.E4, PRE-09.F3. Dependencias: [09.01](#spec-09-01), [09.02](#spec-09-02). | Q | A | PENDIENTE_LOTE_5 |
| <a id="spec-09-04"></a>09.04 | `specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md` | Aceptación, rendimiento y DoD; entradas PRE-03.E4, PRE-06.E4, PRE-09.F1, PRE-09.F2, PRE-09.F4. Dependencias: [09.03](#spec-09-03). | Q | A | PENDIENTE_LOTE_5 |
| <a id="spec-10-01"></a>10.01 | `specs/10-operations/10.01-deployments-origins-regions-and-secrets.md` | Despliegues, origen, regiones y secretos; entradas PRE-10.A1, PRE-10.A2, PRE-10.A3, PRE-10.A4, PRE-10.B1, PRE-10.B2, PRE-10.B3. Dependencias: [01.03](#spec-01-03), [05.04](#spec-05-04), [07.02](#spec-07-02). | A | A | PENDIENTE_LOTE_5 |
| <a id="spec-10-02"></a>10.02 | `specs/10-operations/10.02-observability-quotas-and-rate-limits.md` | Observabilidad, cuotas y límites de tasa; entradas PRE-10.B4, PRE-10.D1, PRE-10.D2, PRE-10.D3, PRE-10.D4, PRE-10.E1, PRE-10.E2, PRE-10.E3, PRE-10.E4, PRE-01.D2, PRE-05.F1. Dependencias: [10.01](#spec-10-01), [06.03](#spec-06-03), [08.02](#spec-08-02). | A | A | PENDIENTE_LOTE_5 |
| <a id="spec-10-03"></a>10.03 | `specs/10-operations/10.03-backups-restoration-and-retention.md` | Backups, restauración y retención; entradas PRE-10.C1, PRE-10.C2, PRE-10.C3, PRE-10.C4, PRE-10.F3. Dependencias: [10.01](#spec-10-01), [05.03](#spec-05-03), [07.03](#spec-07-03). | A | A | PENDIENTE_LOTE_5 |
| <a id="spec-10-04"></a>10.04 | `specs/10-operations/10.04-releases-incidents-and-portability.md` | Releases, incidentes y portabilidad; entradas PRE-10.F1, PRE-10.F2, PRE-10.F4. Dependencias: [10.02](#spec-10-02), [10.03](#spec-10-03). | A | A | PENDIENTE_LOTE_5 |
| <a id="spec-11-03"></a>11.03 | `specs/11-orchestration/11.03-milestones-gates-and-evidence.md` | Hitos, gates y evidencias; entradas PRE-11.D1, PRE-11.D2, PRE-11.D3, PRE-11.E2, PRE-11.E3, PRE-11.F2, PRE-11.F4. Dependencias: [09.04](#spec-09-04), [10.04](#spec-10-04). | C | A | PENDIENTE_LOTE_5 |
| <a id="phase-00"></a>PH-00 | `phases/phase-00-feasibility-spike.md` | Plan futuro PH-00; specs y gate de 11.03. | C | M | PENDIENTE_LOTE_5 |
| <a id="phase-01"></a>PH-01 | `phases/phase-01-go-engine.md` | Plan futuro PH-01; specs y gate de 11.03. | C | M | PENDIENTE_LOTE_5 |
| <a id="phase-02"></a>PH-02 | `phases/phase-02-contracts-persistence-guests.md` | Plan futuro PH-02; specs y gate de 11.03. | C | M | PENDIENTE_LOTE_5 |
| <a id="phase-03"></a>PH-03 | `phases/phase-03-accessible-client-and-3d.md` | Plan futuro PH-03; specs y gate de 11.03. | C | M | PENDIENTE_LOTE_5 |
| <a id="phase-04"></a>PH-04 | `phases/phase-04-private-multiplayer.md` | Plan futuro PH-04; specs y gate de 11.03. | C | M | PENDIENTE_LOTE_5 |
| <a id="phase-05"></a>PH-05 | `phases/phase-05-ai-practice-tutorial.md` | Plan futuro PH-05; specs y gate de 11.03. | C | M | PENDIENTE_LOTE_5 |
| <a id="phase-06"></a>PH-06 | `phases/phase-06-accounts-elo-statistics.md` | Plan futuro PH-06; specs y gate de 11.03. | C | M | PENDIENTE_LOTE_5 |
| <a id="phase-07"></a>PH-07 | `phases/phase-07-hardening-and-portfolio.md` | Plan futuro PH-07; specs y gate de 11.03. | C | M | PENDIENTE_LOTE_5 |
| <a id="h-build"></a>H-BUILD | `../harness/build/README.md` | Consumo futuro de planes por agentes | C | M | PENDIENTE_LOTE_5 |
| <a id="h-ctx"></a>H-CTX | `../harness/context/README.md` | Paquetes mínimos de contexto | C | M | PENDIENTE_LOTE_5 |
| <a id="h-log"></a>H-LOG | `../harness/build-log.md` | Registro factual de trabajo documental | C | B | PENDIENTE_LOTE_5 |

### Lote 6 — 1 archivos

| ID | Archivo o ruta prevista | Propósito / entradas y dependencias | Owner | Complejidad | Estado |
|---|---|---|---|---|---|
| <a id="aud-01"></a>AUD-01 | `audits/documentation-readiness.md` | Auditoría documental integral | C | M | PENDIENTE_LOTE_6 |

## Dependencias de redacción

Las flechas indican prerrequisito hacia consumidor. Las referencias de consulta pueden mirar a una entrada futura del índice sin ser un requisito de cierre ni crear un ciclo. Las specs iniciales incorporan calidad aprobada PRE-09; el dominio 09 la consolida después.

```mermaid
flowchart TD
  R["Registro y maestro"] --> G["11.01 Gobierno"]
  G --> H["11.02 Delegación"]
  G --> INF["01 Infraestructura"]
  G --> P["02 Producto"]
  INF --> ENG["03 Motor"]
  P --> ENG
  INF --> S["05.01 Límites de servicios"]
  ENG --> S
  S --> AU["07.01–07.02 Identidad y autorización"]
  P --> AU
  AU --> API["05.02–05.04 HTTP, datos y eventos"]
  ENG --> API
  API --> DA["07.03–07.04 Perfiles y Elo"]
  AU --> DA
  DA --> MP["06 Multiplayer"]
  API --> MP
  AU --> MP
  MP --> UX["04 Visual y accesibilidad"]
  API --> UX
  P --> UX
  ENG --> AI["08 IA y aprendizaje"]
  API --> AI
  UX --> QA["09 Calidad"]
  AI --> QA
  MP --> QA
  DA --> QA
  INF --> OP["10 Operaciones"]
  API --> OP
  MP --> OP
  DA --> OP
  AI --> OP
  QA --> CL["11.03 Hitos y evidencias"]
  OP --> CL
  CL --> PH["Planes y harness"]
  PH --> AUD["Auditoría lote6"]
```

### Orden de las bases del lote 1

SRC → REG → inventario/trazabilidad inicial → 11.01 → AGENTS y 11.02 → PROMPTS. PRE y ADR consumen decisiones fuente y pueden redactarse en paralelo tras fijar el formato de referencias. Los índices se integran al final; no son una aprobación semántica del contenido que indexan. GOALS y README expresan el alcance del registro; PLANS documenta estados y autorizaciones.

### Secuencias posteriores

- Lote 2: 01.01 → 01.02 → 01.03; 02.01 → 02.02/02.03; 03.01 → 03.02 → 03.03 → 03.04. Arquitectura y producto admiten trabajo paralelo.
- Lote 3: 05.01 → 07.01 → 07.02 → 05.02 → 05.03 → 05.04 → 07.03/07.04 → 06.01 → 06.02 → 06.03.
- Lote 4: V e I trabajan en archivos independientes; revisión conjunta de eventos, práctica y equivalencia HTML/3D.
- Lote 5: Q y A trabajan en paralelo; 11.03, fases y harness se integran tras consolidar calidad y operación.
- Lote 6: auditoría independiente; los cambios semánticos vuelven al propietario.

El lote 2 está autorizado y entregado para revisión. No hay autorización de redacción para lotes 3–6. [PLANS](../PLANS.md) conserva el gate vigente.

## Estado posterior a aprobación del lote 1

Lote 1 APROBADO por el propietario e integrado en main mediante squash 329b7c1059e8a74f191c256870734e8ae8259966. La autorización actual comprende diez specs del lote 2 y registros necesarios, sin ejecución de producto. Las referencias anteriores a distribución 29+2 describen la publicación previa al merge; main contiene ahora los 31 Markdown. Lotes 3–6 siguen pendientes.


## Integración de lote 2 en revisión

Diez specs nuevas, 17 apartados cada una; 42 decisiones canónicas con spec EN_REVISIÓN. La distribución total es 17 decisiones con spec APROBADA, 42 con spec EN_REVISIÓN y 187 con spec futura. Se conservan los 40 nodos y las dependencias del inventario original, sin ciclos. El informe, las contribuciones, los controles y los bloqueos se documentan en [PLANS](../PLANS.md#informe-lote-2); el detalle normativo abierto permanece en [RISK-01](conflicts.md#risk-01). El gate «APRUEBO LOTE 2» está pendiente.
