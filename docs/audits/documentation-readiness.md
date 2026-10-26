# Auditoría integral de preparación documental para implementación — Toguz Kumalak

> **Estado vigente — 2026-10-26: LOTE_6_APROBADO_DOCUMENTALMENTE.** El propietario aprobó AUD-01 como auditoría documental final sobre HEAD 6052689ff9b15a19214267d7cc6a614d17f8fcdd; PR #6 integrado por e81f01a0e94c3037e3b922894202b6eee80534e8. [Recibo de cierre](../../PLANS.md#cierre-lote-6) prevalece sobre estados y paradas históricos. AUD6-F01, AUD6-F02 y AUD6-C01–C06 permanecen pendientes para sus fases. No se autoriza implementación, ejecución de PH00/PH01, workflows o despliegues.

## 1. Identificación

ID: AUD-01 · Versión: 0.2 · Fecha: 2026-10-26 · Estado documental: APROBADO · Owner: C — coordinación.

| Campo | Base examinada / alcance |
|---|---|
| Repositorio | SebasChaconDev/toguz-kumalak |
| Commit auditado | 00c8e121f96e990beb052f72d19132f282d08cfa, main local/remoto verificados |
| Rama de entrega | docs/audit-lote-06; no fusión autorizada |
| PR5 / integración | MERGED, HEAD aprobado 9ff1a2857d0b91b9cf7925a4c36de30e377cef3d; squash 6579ca930a9081f773b390ce61f6bc80cc6b1ae7 |
| Evidencia experimental | Spike 504caa5d8d4f538e131e1beffc8843cb8938772a, consultado mediante Git sin ejecución |
| Fuente de autorización | Encargo del propietario de 759 líneas / 22.618 bytes, SHA256 88554893b57fbefbc5a93bb410b1e4d0601eec1236201d2fac1fd016aa03544b; «AUTORIZO LA EJECUCIÓN Y PUBLICACIÓN DOCUMENTAL DEL LOTE 6.» |
| Escritura permitida | Este archivo y cambios mínimos en IDX/TRZ/CONF/PLANS; ninguna spec, PRE, ADR, código o proveedor |
| Naturaleza | Fuentes APROBADO; recomendaciones/correcciones PROPUESTA; comportamiento del producto POR_VERIFICAR. Evidencia documental observada lleva procedencia, sin nuevo estado documental |

El [ledger y autorización](../../PLANS.md#ejecucion-lote-6) registra base, responsables, controles y publicación. Los recibos posteriores del propietario prevalecen sobre cortes históricos. Este documento no cambia decisiones, aprueba propuestas técnicas específicas o autoriza programar.

## 2. Resumen ejecutivo

La base permite solicitar autorización para un primer incremento puro y delimitado de PH-01 sin inventar reglas, identidad, persistencia o autoridad deportiva. No se demostró P0/P1 que bloquee PH-01 o el cierre documental. La decisión formal está en§14; sus condiciones delimitan trabajo independiente y tareas posteriores.

Se examinó la coherencia de las 40 specs originales por dominios, los ocho PH y los tres harness. Los 246 requisitos canónicos conservan texto y dueño; las 14 ACC conservan identidad/significado. La revisión semántica no se sustituye por validar Markdown o por aceptar informes de agentes anteriores.

Se registran dos observaciones nuevas: una precisión de perspectiva lógica del oráculo de IA(P2/NON_BLOCKING) y un ejemplo de ruta inconsistente(P3/NON_BLOCKING). Se conservan condiciones conocidas de L4-P01, política de retención, licencias, recursos/medición y restauración, cada una ligada a sus acciones afectadas. Ninguna requiere reabrir TK-N1-1.0.0 o impide preparar corpus/tipos Go.

Primero: PH-01-T01, corpus/oráculos. Primer código Go: PH-01-T02 después de T01 aceptado. No se han iniciado esas actividades. Aprobar esta auditoría será un gate documental; código, pruebas, instalaciones y operaciones externas requieren alcance humano posterior identificable.

## 3. Inventario y cobertura

Recuento directo de archivos y filas, no reproducción de un reporte anterior: base80 Markdown versionados y .gitignore,40 specs aprobadas, ochoPH y tresharness. Esta entrega añade únicamente AUD-01:81 Markdown/40 specs/41 documentos de soporte. El manual sin versionar de la carpeta original no pertenece a ese inventario ni se incorpora.

La tabla reproduce todos los nodos, versiones, propietarios, asignación PRE y prerrequisitos extraídos de IDX. También es la lista de adyacencia completa del DAG; las referencias de consulta no se convierten en aristas inversas.

| Spec y fuente | Versión | Owner | PRE propias | Prerrequisitos canónicos | Estado |
|---|---|---|---:|---|---|
| [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md) | 0.2 | C | 13 | REG/SRC | APROBADO |
| [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md) | 0.2 | C | 4 | 11.01 | APROBADO |
| [01.01](../specs/01-infrastructure/01.01-stack-monorepo.md) | 0.2 | A | 5 | REG/SRC | APROBADO |
| [01.02](../specs/01-infrastructure/01.02-platform-feasibility.md) | 0.2 | A | 4 | 01.01 | APROBADO |
| [01.03](../specs/01-infrastructure/01.03-cost-capacity-evolution.md) | 0.2 | A | 2 | 01.02 | APROBADO |
| [02.01](../specs/02-product/02.01-scope-and-modes.md) | 0.2 | C | 5 | REG/SRC | APROBADO |
| [02.02](../specs/02-product/02.02-local-flows-and-i18n.md) | 0.2 | C | 3 | 02.01 | APROBADO |
| [02.03](../specs/02-product/02.03-accounts-history-and-licensing.md) | 0.2 | C | 4 | 02.01 | APROBADO |
| [03.01](../specs/03-game-engine/03.01-normative-rules-and-corpus.md) | 0.4 | M | 3 | 02.01 | APROBADO |
| [03.02](../specs/03-game-engine/03.02-state-indices-and-invariants.md) | 0.4 | M | 4 | 01.01, 03.01 | APROBADO |
| [03.03](../specs/03-game-engine/03.03-moves-and-semantic-events.md) | 0.4 | M | 6 | 03.02 | APROBADO |
| [03.04](../specs/03-game-engine/03.04-outcomes-repetition-and-replay.md) | 0.4 | M | 6 | 03.03 | APROBADO |
| [05.01](../specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md) | 0.2 | S | 5 | 01.02, 03.04 | APROBADO |
| [05.02](../specs/05-backend/05.02-http-openapi-and-errors.md) | 0.2 | S | 4 | 05.01, 07.01, 07.02 | APROBADO |
| [05.03](../specs/05-backend/05.03-transactions-data-and-migrations.md) | 0.2 | S | 8 | 05.02, 03.04 | APROBADO |
| [05.04](../specs/05-backend/05.04-event-schemas-and-postcommit-publication.md) | 0.2 | S | 4 | 05.03, 07.02 | APROBADO |
| [06.01](../specs/06-multiplayer/06.01-rooms-invitations-and-ranked-lifecycle.md) | 0.2 | S | 6 | 05.04, 07.04 | APROBADO |
| [06.02](../specs/06-multiplayer/06.02-ably-notifications-and-reconciliation.md) | 0.2 | S | 6 | 06.01, 05.04 | APROBADO |
| [06.03](../specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md) | 0.2 | S | 6 | 06.02, 07.01, 07.02 | APROBADO |
| [07.01](../specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md) | 0.2 | S | 8 | 05.01, 02.03 | APROBADO |
| [07.02](../specs/07-identity-data/07.02-tokens-authorization-and-service-identity.md) | 0.2 | S | 6 | 07.01 | APROBADO |
| [07.03](../specs/07-identity-data/07.03-profiles-history-and-data-rights.md) | 0.2 | S | 7 | 05.04, 07.02 | APROBADO |
| [07.04](../specs/07-identity-data/07.04-elo-leaderboard-and-statistics.md) | 0.2 | S | 6 | 05.03, 03.04, 07.02 | APROBADO |
| [04.01](../specs/04-visual-experience/04.01-navigation-and-accessibility.md) | 0.2 | V | 5 | 02.02, 02.03, 03.03, 05.02, 06.01, 06.03 | APROBADO |
| [04.02](../specs/04-visual-experience/04.02-board-assets-camera-and-materials.md) | 0.2 | V | 12 | 04.01, 03.02 | APROBADO |
| [04.03](../specs/04-visual-experience/04.03-interaction-animation-and-audio.md) | 0.2 | V | 4 | 04.01, 03.03, 06.02 | APROBADO |
| [04.04](../specs/04-visual-experience/04.04-performance-and-html-fallback.md) | 0.2 | V | 4 | 04.02, 04.03 | APROBADO |
| [08.01](../specs/08-ai-practice/08.01-search-heuristics-and-compute-limits.md) | 0.2 | I | 9 | 03.04, 01.02 | APROBADO |
| [08.02](../specs/08-ai-practice/08.02-ai-orchestration-fallback-and-metrics.md) | 0.2 | I | 4 | 08.01, 05.02, 05.03, 07.02 | APROBADO |
| [08.03](../specs/08-ai-practice/08.03-reversible-practice-and-scenarios.md) | 0.2 | I | 8 | 02.02, 03.04, 05.02 | APROBADO |
| [08.04](../specs/08-ai-practice/08.04-tutorial-hints-and-educational-replay.md) | 0.2 | I | 9 | 08.03, 03.03 | APROBADO |
| [09.01](../specs/09-quality/09.01-engine-data-contract-and-ai-verification.md) | 0.2 | Q | 16 | 05.04, 07.04, 08.02 | APROBADO |
| [09.02](../specs/09-quality/09.02-ui-e2e-accessibility-and-visual-verification.md) | 0.2 | Q | 6 | 04.04, 06.03, 08.04 | APROBADO |
| [09.03](../specs/09-quality/09.03-ci-security-and-release-controls.md) | 0.2 | Q | 6 | 09.01, 09.02 | APROBADO |
| [09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md) | 0.2 | Q | 5 | 09.03 | APROBADO |
| [10.01](../specs/10-operations/10.01-deployments-origins-regions-and-secrets.md) | 0.2 | A | 7 | 01.03, 05.04, 07.02 | APROBADO |
| [10.02](../specs/10-operations/10.02-observability-quotas-and-rate-limits.md) | 0.2 | A | 11 | 10.01, 06.03, 08.02 | APROBADO |
| [10.03](../specs/10-operations/10.03-backups-restoration-and-retention.md) | 0.2 | A | 5 | 10.01, 05.03, 07.03 | APROBADO |
| [10.04](../specs/10-operations/10.04-releases-incidents-and-portability.md) | 0.2 | A | 3 | 10.02, 10.03 | APROBADO |
| [11.03](../specs/11-orchestration/11.03-milestones-gates-and-evidence.md) | 0.2 | C | 7 | 09.04, 10.04 | APROBADO |

Totales:246 PRE asignadas una sola vez,40 nodos y80 aristas. Las40 specs tienen17 secciones numeradas; todas mantienen APROBADO. Los20 documentos del lote5 conservan v0.2/APROBADO. Ningún ticket implica código ya escrito.

[REG](../decision-register.md) y [SRC](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) conservan246 textos literales iguales. SHA256 del maestro:28C24C498829FCFAD410F07AAA6BD036D5A33382294401E11E86D486E4ECBBEF. Las decisiones documentales y el producto probado son ejes distintos.

## 4. Integridad y trazabilidad

| Control observado sobre la base | Resultado / método |
|---|---|
| Archivos/specs/estados |80 Markdown;40 specs APROBADO y40×17 secciones; nueve nuevas del lote5 incluidas |
| PRE/propiedad/TRZ |246 textos/246 dueños únicos/246 filas SPEC_APROBADO; sin referencias PRE desconocidas o decisiones huérfanas |
| Versiones/IDs |39 IDs de soporte en metadatos inline y40 IDs de specs en fichas, sin colisión; SRC se conserva como fuente íntegra sin forzar metadatos |
| DAG documental |40 nodos/80 aristas extraídos de las filas IDX; algoritmo topológico visita40/40, cero ciclos y mismo mapa de dueños/dependencias que la base |
| Tablas |495 tablas de los80 Markdown; cero diferencias de columnas fuera de fences |
| Enlaces/anclas |4.531 destinos locales y2.953 fragmentos contrastados contra archivos/IDs/slugs, cero ausentes |
| Fences/JSON |Fences cerrados;18 bloques JSON completos parseados; IDs explícitos únicos por archivo |
| Aceptación normativa |14 filas ACC-N1 de REG y03.01 literalmente iguales; consumidores09.01/PH01 contrastados semánticamente en§5 |
| PH/harness |8 planes / 3 harness;65 tickets y54 familias de casos;96 campos Milestone(12×8), detalles§9 |
| Fuente/autoridad |REG/maestro/ADR/PRE/specs protegidos sin cambios en esta rama; no nuevas reglas o arquitectura |

Los números de enlaces/tablas anteriores identifican el corte base auditado; el cierre de esta entrega añade enlaces de AUD/registros y comunica su recuento final en PLANS. Una herramienta no demuestra semántica, seguridad, alcanzabilidad o rendimiento por analizar texto.

La jerarquía vigente es instrucción humana posterior→REG/ADR→specs aprobadas→AGENTS→planes/briefs. Las menciones históricas a NORM sin resolver, L5 EN_REVISIÓN, gates pendientes o spike BLOCKED están sustituidas por AD-NORM-001, AD-GATE-005 y los recibos específicos. Las propuestas actuales conservan su etiqueta aunque su contenedor Markdown esté aprobado.

## 5. Revisión semántica por dominio

### 5.1. Motor Go, normativa e historia

[03.01](../specs/03-game-engine/03.01-normative-rules-and-corpus.md)–[03.04](../specs/03-game-engine/03.04-outcomes-repetition-and-replay.md) fijan un motor puro/determinista/versionado:18 otau / 162 piedras, estado validado antes de indexar/convertir, PitRef por color/índice y copia profunda. Siembra completa, captura rival par, tuzdyk automático y prohibiciones no convierten una jugada en ilegal ni requieren un motor frontend.

R-05/R-06 completan efectos, prueban target≥82/equal81–81 y solo después transfieren conceptualmente turno al rival para atsyrau. No hay barrido ni movimiento extra. Target/equal pueden conservar el turno del actor; atsyrau representa al siguiente agotado. Ganador/cause/outcome no se infieren del turno.

R-07 compara exactamente reglas,18 pozos, 2 kazanes, tuzdyks y turno. Inicial cuenta una vez; solo un movimiento confirmado/completo con salida activa añade aparición. Tres habilitan reclamación de cualquiera de ambos participantes, no empate automático. Simulación, preview, retry y cierres no añaden apariciones.

La separación clave está coherente: [05.03](../specs/05-backend/05.03-transactions-data-and-migrations.md)§7 L113–139 conserva historyOrdinal para movimientos y cierres; movimiento+cierre tienen dos ordinales aunque compartan revisión, con eventos asignados una sola vez por [05.04](../specs/05-backend/05.04-event-schemas-and-postcommit-publication.md)L259. Cero nueva aparición terminal no significa borrar el registro Lifecycle. [09.01](../specs/09-quality/09.01-engine-data-contract-and-ai-verification.md)L60/L168 y PH01 preservan esa diferencia; L5-R01 ya corregido no se reabre.

| ACC canónico | Conducta y negativo contrastados | Suite / tickets previstos |
|---|---|---|
|ACC-N1-01|Una piedra sale al siguiente; origen inválido no muta|F-Q1-N01; PH01T02/T03|
|ACC-N1-02|Última siembra rival par captura; propia/impar no|F-Q1-N02; T03|
|ACC-N1-03|Tres rival permitido crea tuzdyk automáticamente|F-Q1-N03; T04|
|ACC-N1-04|Pozo9/homólogo/segundo prohíben crear, no mover|F-Q1-N04; T04|
|ACC-N1-05|Depósito en tuzdyk al kazán del dueño, pozo0/162|F-Q1-N05; T04|
|ACC-N1-06|Siguiente vacío K60 pierde atsyrau, sin liquidación|F-Q1-N06; T05|
|ACC-N1-07|Siguiente vacío K81 empata conservando remanentes|F-Q1-N07; T05|
|ACC-N1-08|Actor vacío y rival con jugadas continúa|F-Q1-N08; T05|
|ACC-N1-09|Kazán82 prima tras completar siembra|F-Q1-N09; T05|
|ACC-N1-10|Tercera aparición habilita derecho, no draw automático|F-Q1-N10; T06|
|ACC-N1-11|Ambos participantes reclaman con revisión antes de próxima jugada válida|F-Q1-N11; T06 predicado, PH02/04 integración|
|ACC-N1-12|Turno/tuz/kazán distintos no son misma posición; stale rechaza|F-Q1-N12; T06|
|ACC-N1-13|Preview/retry/cierre sin aparición nueva; lifecycle durable una vez|F-Q1-N13; T06/T07, PH02/04 receipt/locks|
|ACC-N1-14|Terminal definitivo prima sobre repetición y no crea derecho posterior|F-Q1-N14; T05/T06|

Son pruebas futuras. PH01 no acredita mediante unitarias la autorización, lock o receipt de ACC11/13; esos subcasos permanecen en PH02/04. Corpus distingue SOURCE_DIRECT/INITIAL, DERIVED_MANUAL/DERIVED_PATH_REVIEWED y SYNTHETIC_CONTRACT/UNPROVEN. Suma 162 y nombre DER no demuestran alcanzabilidad. Replay exige versiones, contexto y prefijo/ordinales completos; un gap no se repara silenciosamente.

### 5.2. Backend, identidad y datos

[05.01](../specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md)–[05.04](../specs/05-backend/05.04-event-schemas-and-postcommit-publication.md) y [07.01](../specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md)–[07.04](../specs/07-identity-data/07.04-elo-leaderboard-and-statistics.md) mantienen tres servicios:Identity posee sus datos, Game sus partidas/recibos/resultados/Elo, AI no tiene DB ni autoridad de movimiento. No hay SQL cruzado o permiso por conocerUUID. Game/AI consumen un motor Go; Identity no necesita importar reglas.

Comandos ligan actor original,commandId,operación/recurso y fingerprint. Autorización actual precede al receipt y receipt precede a stale. Un timeout de COMMIT sin confirmación permanece unknown, con la misma intención; no demuestra rollback. Cierre/estado/eventos/contexto/ratings/ledger/agregados/receipt son atómicos.

La jerarquía común de guardas humanas UUID ASC→rooms→games→ratings→derivados y reinicio completo al cambiar el conjunto evita adquirir locks inferiores tarde; no promete ausencia universal de deadlocks. Dos partidas de A usan su rating actual serializado, sin duplicar Elo por retries o usar rating inicial congelado.

Credenciales de producto, Google, S2S, protección hosting y Ably se separan. Firma/issuer/audience/typ/env/frescura/JWKS, Origin/CSRF, bindingJCS y recurso son defensas distintas. Vínculo G→U no transfiere activas o receipts. Refresh perdido recupera solo la respuesta original acreditada por pruebas y dentro de su horizonte; no emite otro derecho por retry.

### 5.3. Multijugador, visual y accesibilidad

[06.01](../specs/06-multiplayer/06.01-rooms-invitations-and-ranked-lifecycle.md)–[06.03](../specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md) y [04.01](../specs/04-visual-experience/04.01-navigation-and-accessibility.md)–[04.04](../specs/04-visual-experience/04.04-performance-and-html-fallback.md) separan Room/Game, consentimiento ready, elegibilidad y autorización. Claim, renuncia/oferta no se restringen por turno del movimiento. Actividad tiene versión por participante; GET/Ably/duplicados no renuevan24h ni un reloj cliente cierra la partida.

Ably NoticeV1 es aviso postcommit; duplicado, pérdida/desorden o SDK reconectado no confirma estado. Un controlador de reconciliación instala proyección REST autorizada; receipt histórico no retrocede snapshot actual. Eventos/replay completos y versionados animan después de confirmación, con skip/gaps/unknown seguros.

HTML/3D comparten controlador, PitRef, foco, selección e intención; cambiar renderer no envía comandos. Mantienen funciones esenciales, ES/EN, teclado/lector/móvil y aprendizaje. Three.js/R3F/Drei representan estado de Go, no calculan reglas. Recursos premium/procedencia/cultura y objetivos60/30FPS requieren activos aprobados/medición real; axe, emulación o una escena idle no acreditan WCAG o rendimiento.

### 5.4. IA, práctica y aprendizaje

[08.01](../specs/08-ai-practice/08.01-search-heuristics-and-compute-limits.md)–[08.04](../specs/08-ai-practice/08.04-tutorial-hints-and-educational-replay.md) especifican Negamax/alfa-beta/profundización iterativa, candidato legal barato y última iteración completa, no parcial. Evaluación terminal usa outcome/winner y perspectiva lógica de búsqueda, que puede diferir de State.turn retenido. La mejora AUD6-F01 hace explícito ese riesgo en el oráculo, sin nuevo reglamento.

TT move-seed-only/1 orienta orden y revalida legalidad; no guarda score,bound,terminal o elegibilidad, ni siquiera dentro de una búsqueda. Igual tablero/count actual no reemplaza historias relevantes de otras posiciones. Cada rama conserva/copía contexto; colisión hash compara exacto.

Tiempo/nodos/profundidad/arena y trabajo de features cuentan; cancelación por invalidación no permite aplicar candidato. AI calcula fuera de locks, tras claim confirmado, y Game revalida/commitea un efecto. Callback incierto, respuestas tardías y carrera AI/fallback tienen negativos, sin prometer exactly-once del cómputo.

Práctica usa rama/cursor/generación/intención, no solo revisión reutilizable. Undo/redo/restauración y bifurcación no producen Elo o escrituras deportivas. Hint solo presenta: no selecciona intención, envía Move, cambia timeline o progreso; una jugada posterior exige acción humana separada. Se rechaza ayuda ligada a recurso ranked conocido, sin prometer detectar una copia synthetic sin procedencia ni prohibir toda práctica concurrente.

L4-P01 sigue pendiente: catálogo05.02, persistencia05.03 y autorización07.02 deben cambiar con aprobación específica antes de integrar operaciones de rama/PracticeRecord/ObjectiveEvaluation/cancelación/recovery nuevos. No bloquea motor o búsqueda pura; delimitación§12.

### 5.5. Calidad y operaciones

[09.01](../specs/09-quality/09.01-engine-data-contract-and-ai-verification.md)–[09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md) exigen requisito/owner/caso±/método/entorno/esperado/evidencia/AC.90% sentencias motor y80% lógica de servicios no sustituyen rutas críticas,14 ACC, integración PostgreSQL real, fuzz/race, contratos, JWT/JWKS, Elo o accesibilidad manual.

[09.03](../specs/09-quality/09.03-ci-security-and-release-controls.md)L84–115 define impacto por consumidor, PR/main/pre-release/smoke y PR externos sin secretos. Cambiar motor dispara Game/AI, OpenAPI genera ambos consumidores y migración exige PostgreSQL real. Check requerido cancelado/ausente o artefacto inaccesible no pasa. No hay CI instalada por aprobar la spec.

[09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md) y [10.04](../specs/10-operations/10.04-releases-incidents-and-portability.md) comparten bloqueos:P0/P1/crítico no validado/fallo requerido; P2 diferido exige owner/plazo/mitigación y aceptación humana específica, sin eximir PRE o volver verdeFAIL. Release exige compatibilidad, evidencia, revisión y permiso externo.

[10.01](../specs/10-operations/10.01-deployments-origins-regions-and-secrets.md)–[10.04](../specs/10-operations/10.04-releases-incidents-and-portability.md) cubren secretos/regiones/orígenes, observación/cuotas persistidas, backup→restore aislado, expand/backfill/contract, incidentes/rollback/portabilidad. Backup no demuestra restaurabilidad; rollback de código no revierte datos. Sin verdad posterior suficiente para reconciliar auth/recibos/Elo, se mantiene cuarentena. A coordina operación,C gobierno,Q verifica recuperación.

## 6. Revisión de contratos entre dominios

Los contratos mantienen su naturaleza fuente, incluidos formatos PROPUESTA. Esta matriz no los aprueba por inferencia. Pruebas y resultados son exigencias futuras, no observaciones de producto.

| Control / owner | Productor → consumidor; versión y permiso | Positivo / negativo / resultado esperado | Método, entorno y evidencia futura |
|---|---|---|---|
| X01 M/Q | 03.*→Game/AI; State/PitRef/Transition y versiones de reglas/estado/movimiento | Inicial y14ACC conservan162; inválidos y aliasing no mutan entrada | Corpus, unidad, propiedades y fuzz; Go local sin red; copia/semilla, F-Q1-N01–14/01/02 y AC03.02/03.03 |
| X02 S/Q | 05.03→receipt/snapshot; actor,commandId,fingerprint y autorización actual | Retry recupera decisión única; hash distinto409; unknown conserva intención | PG real, conexiones/barreras y reinicio, SQL antes/después; F-Q1-03–05/F-TX-03–06/15,05.03 AC03/05 |
| X03 S/Q | Game→Elo/ledger; elegibilidad inicial y ratings actuales | G1(A,B)/G2(A,C) ambos órdenes; doble cierre o fallo ledger no produce parcial/rating congelado | Locks/SQLSTATE/ledger/secuencias en PG real; F-Q1-06/07,F-ELO-05–09/14/15;05.03 AC04 |
| X04 S/V/Q | 07.*→browser/Game; JWT/JWKS,cookie,CSRF,guest/refresh | Credenciales/pruebas válidas recuperan derecho original; key/env/Origin/actor cruzados rechazan | Fixtures cripto/reloj/JWKS304/Age,PG y browsersHTTPS sintéticos; F-AUT-01–14/22,F-SES;07.02 AC01–06/11 |
| X05 S/I/Q | Game→AI→claim; ServiceClaims/JCS/grant/attempt y método/path/query/body | Claim confirmado habilita un handler; otro attempt/callback incierto no calcula y AI nunca escribe partida | Fixtures Go/TS,HTTP/CAS PG/race; F-AUT-15–19/F-Q1-11,07.02 AC07–09 |
| X06 S/V/Q | Commit→NoticeV1→REST→UI; versión/rango y capability exacta | Aviso mayor coalescido y REST autorizado; pérdida/duplicado/gap no adjudica y HTTP perdido recupera | Game/PG real; broker doble separado de Ably autorizado; trazas DB/red; F-EVT/F-SYN/F-Q2-03/04/06–08 |
| X07 M/S/V/Q | Historia→ReadReplay→renderer; historyOrdinal/moveOrdinal/revisión/sequence | Move+cierre conserva dos ordinales y cada evento una vez; gap/página incompatible no se repara | Replay desde inicial/checkpoint y cotejo; F-Q1-02/N13,F-INT-04/06,F-Q2-08 |
| X08 V/Q | Controlador→HTML/3D; PresentationContext/PitRef/foco/intención | Renderer conserva funciones/foco; WebGL caído/unknown no envía Move/Start/Resume | Vitest/RTL/Playwright,tres browsers,ES/EN,móvil y teclado/lector manual; DOM/red/baseline; F-Q2-01–12 |
| X09 I/M/Q | Contexto Go→SearchInput/Result/Stats /1; salida no autoritativa | Exhaustivo/Negamax coinciden a profundidad completa; State.turn terminal/hash/count no reemplaza contexto | Árboles/corpus/colores/terminales,TT on/off sin valores; Go local,nodos/memoria/semilla; F-BUS-01–08/F-Q1-15–17 |
| X10 I/S/V/Q | Rama/Hint/Objective→simulación/UI; deltaL4-P01 pendiente,tuple/generación/intento | Hint solo expone y jugada exige otra intención; respuesta vieja/replay/CAS no mueve ni progresa | Dobles UI declarados; integración tras cambio aprobado; PG/red/timeline antes/después; F-Q1-18,F-ORQ-06–09,PH05-C05–07 |
| X11 S/A/Q | Identity delete→Game barrier/cleanup; raíz/step/ACK/policyVersion | Barrier revoca sin política y ACK_REVOKED durable; cleanup null/cambiado/tercero no completa | PG real,fallos entre commits/lotes,reinicio,manifest/acuse; F-DAT-10–18,F-TX-12,PH06-C07;07.03 AC08–11 |
| X12 Q/A/C | QA→CI→AcceptanceDossier/ReleaseManifest; HEAD exacto y permisos separados | Trigger cubre consumidores; required omitido,PR externo con secretos o P2 sin aceptación bloquean | Runner autorizado,matriz impacto/checks/filtros/scans/artefactos; F-Q3-01–08,F-Q4-01–03/09,F-REL-01/02 |
| X13 A/S/Q | Instancias→QuotaObservation/RatePolicy; contador durable y cap global | Instancias comparten límite; reinicio RAM/cuota desconocida/cardinalidad o PII ilimitadas no pasan | PG,canarios,fallo exporter,series crudas/ledger,carga autorizada; F-OBS-01–11,10.02 AC01–11 |
| X14 A/S/Q | Backup→RestoreTrial→RecoveryDecision; versión,cortes,auth y datos | Restore aislado reconcilia; backup truncado,postbackup perdido o refresh revivido mantienen cuarentena | PG real,replay Go sin red,manifests/hash/tiempos,expand/contract; F-BAK-01–07,F-Q4-08,F-REL-03/05,PH06-C08/PH07-C06/07 |
| X15 V/A/Q | Assets/mediciones→DoD/release; procedencia,dispositivo,versión y permiso | Datos/regresión manual coherentes; baseline que oculta error,axe/emulación/muestra única no certifican | OS/browser/GPU/DPR/fuentes/locale,muestras/percentiles/errores/censura; F-SCN/F-Q2-10–12/F-Q4,PH03-C05/06 |

La lectura comprueba compatibilidad textual y controles previstos; falta demostrar su implementación. Un doble de broker/IA no acredita Ably/Negamax o interoperabilidad real.

## 7. Evidencia experimental

Fuentes inmutables: [INTEGRAL_REPORT](https://github.com/SebasChaconDev/toguz-kumalak/blob/504caa5d8d4f538e131e1beffc8843cb8938772a/spikes/infra-feasibility/INTEGRAL_REPORT.md), [TEST_MATRIX](https://github.com/SebasChaconDev/toguz-kumalak/blob/504caa5d8d4f538e131e1beffc8843cb8938772a/spikes/infra-feasibility/TEST_MATRIX.md), [REMOTE_AUDIT§13](https://github.com/SebasChaconDev/toguz-kumalak/blob/504caa5d8d4f538e131e1beffc8843cb8938772a/spikes/infra-feasibility/REMOTE_AUDIT.md#remote-close-20261025) y REMOTE_PLAN§11 del mismo commit. Se consultaron informes/cierres y artefactos granulares mediante Git, sin ejecutar el experimento o consultar cuentas.

| Artefacto contrastado | Observación histórica | Límite |
|---|---|---|
| integral-local-gate.json y summary del run integral-7d08fb85f2ed4c62b505c985d76e8d5b | GO_LOCAL; PostgreSQL real; Game PID24004→7992; DB reiniciada conserva revisión/efectos/receipts3/3/3 | Procesos locales/PGDocker; Ably doble, HMAC provisional y motor/AI sintéticos |
| remote-functional-20261025.json |108 registros de casos; fallos CAP-03-other-channel,DB-CROSS-DENIED-identity,COLD-INSTANCE-game y posteriores correcciones preservados | No son108casos independientes todos verdes. SSE HTTP200 puede contener error40160; se verifica el evento |
| TX-FIRST/RETRY/CONFLICT/CONCURRENT/ROLLBACK y TX-FINAL-SQL | Commit/relectura antes del aviso; retry único, carrera200/409, rollback conserva estado y SQL/REST final4/4/4 | Dos contendientes y datos sintéticos; respuesta perdida modelada, sin corte TCP físico |
| POC01/02 y source-manifest/equivalence |6builds READY sin caché,4roots,8intentos con2cancelados;42entradas fuente; Game/AI comparten motor e Identity no lo importa | Go remoto1.26.0/local1.26.2 y VercelCLI65.1.0 son versiones experimentales, no lockfile de producto |
| POC03/05/06/07/08 | NeonPG18, roles/CONNECT/DDL42501; Ably real postcommit/capabilities; OIDC de hosting y cookie Chrome protegidos | Reconciliador sintético; no JWT/JWKS,GoogleOIDC o seguridad integral del producto |
| remote-latency-followup-20261025.json |12SQL instrumentadas y una muestra por métrica: pooled519,170ms,TX287,993ms,publish→receive258,356ms; commit visible antes de publicar | Operador→proveedores, no Game/Vercel ni p95/SLO; sin reactivación Neon o colas PgBouncer |
| remote-gate-20261025.json y REMOTE_AUDIT L224/L251 | SQL80 previsto superado:117reconstruidas conservadoras,49directas instrumentadas,exactWireCount=null;0USD adicional contratado | No cumplimiento80, conteo exhaustivo global o garantía de coste futuro |
| Retirada§13 y gateJSON | Último cierre:4proyectosVercel/6deployments;2generaciones secuenciales Neon/Ably;8clavesapp revocadas;remainingOwn=0 | Otros4proyectos/2deployments cancelados históricos ya retirados; no inventario actual de cuentas |

La excepción de bootstrap production y ocho builds perteneció exclusivamente al experimento protegido y desechable. El propietario retiró el techo100HTTP sin borrar consumo/fallos. La desviaciónSQL permanece explícita; no se inventa excepción literal.

GO_LOCAL/GO_REMOTE acredita factibilidad experimental. Motor TK,Negamax,clientes reales,restore,rendimiento/coste de producto y fallos físicos necesitan pruebas propias PRE-09. La capacidad/cuenta/cuota/runtime se revalida al implementar el consumidor afectado, sin bloquear por ello todo el motor puro.

## 8. Seguridad y datos

El diseño mantiene mínimos por propósito/entorno:app/DDL,Identity/Game,Game/AI,usuario/S2S/hosting/Ably. Que un upstream esté protegido no reemplaza autorización de recurso; conocer una invitación, alias,UUID o ser administrador no concede lectura deportiva de un tercero.

07.02 exige JWKS confiable/fresco, firma/tipo/audiencia/emisor/expiración, rutas permitidas y binding de bytes; prueba de caché/rotación/unknownkid y PRexternos sin secretos siguen pendientes. Logs/evidencia no contienen JWT/cookies/proofs/DSN/PII; hashes y referencias de artefactos no sustituyen acceso y observación del revisor.

Derechos de datos separan confirmación legítima, raízdurable,revocación ybarrierS2S sinpolicyVersion,ACK_REVOKED persistido. Cleanup requierepolicyVersionaprobada/no nula/inmutable; sin política raíz/barrier quedanBLOCKED_PENDING_DATA_POLICY. ACK_REVOKED no es COMPLETE y falta política no aplaza una barrera autorizada. No se modifica resultado deportivo por delete ni se impiden acciones legales del rival.

La política de retención/plazos,temporales,telemetría,copias ylegitimidad de conservación siguen pendientes del propietario. No hay cumplimiento jurídico o anonimización irreversible acreditados. Restore a cuarentena debe cotejar auth/grants/recibos/Elo yrevocaciones posteriores; si falta verdad de decisiones posteriores, no reabre. Un protocolo/epoch nuevo exige control de cambios; no se agrega en esta auditoría.

## 9. Planes de implementación

Lectura completa de PH00–07 y contraste con contratos. Recuento independiente:65tickets únicos,54familias y ocho fichas con12campos exactos. Expandir rangos como T03–T08 produce71aristas internas:65/65visitas topológicas, cero ausentes/ciclos. La cadena PH00→01→02→03→04→05→06 y PH07 dependiente de00–06 suma13aristas entre fases, sin ciclos.

| Plan | Tickets/casos | Entrada→entrega y frontera/gate |
|---|---:|---|
| [PH00](../phases/phase-00-feasibility-spike.md) |5/5| Recibo experimental, versiones y límites→dossier de entrada; no repetir el spike automáticamente |
| [PH01](../phases/phase-01-go-engine.md) |8/4| Corpus→motor puro/contexto/eventos/replay;14ACC,90% y críticos; sin cuentas,SQL,red,WebGL,AI o Ably |
| [PH02](../phases/phase-02-contracts-persistence-guests.md) |8/7| Contratos/PGreal/guest/S2S/receipts; transporte AI puede usar doble declarado sin afirmar búsqueda completa |
| [PH03](../phases/phase-03-accessible-client-and-3d.md) |8/6| HTML jugable primero y3D progresivo; contratos de salas/cuentas mediante interfaces/dobles con integración pendiente |
| [PH04](../phases/phase-04-private-multiplayer.md) |8/7| Privado casual/invitados,notificación/recovery; ranked/Elo espera PH06 sin retirarlo del lanzamiento |
| [PH05](../phases/phase-05-ai-practice-tutorial.md) |9/8| Búsqueda pura independiente; delta de ramas/hints condicionado aC01; oráculo con perspectiva explícita |
| [PH06](../phases/phase-06-accounts-elo-statistics.md) |9/8| OIDC/cuentas/Elo concurrente/estadísticas/derechos; cleanup y retenciones afectados esperan política |
| [PH07](../phases/phase-07-hardening-and-portfolio.md) |10/9| DoD integral,restore,medición,licencias/cultura/portafolio; release/demo solo con permiso y evidencia |

Estado de los65tickets: planificados; ninguna implementación o prueba de producto acreditada. Cada Milestone define responsable,riesgos,fuentes y gate; cada ticket añade dependencia,pasos/entrega y aceptación. TaskEnvelope posterior fija archivos concretos antes de escribir.

No se encontró ciclo por consumir contratos posteriores: PH02 declara doble deAI; PH03 separa UI de integración de salas/cuentas; PH04 no ofrece ranked antes dePH06. Referencias PH02T06→PH01,PH03T02→PH02 yPH06T04→PH04 apuntan a ancestros. Leer un contrato no demuestra un componente implementado.

## 10. Orquestación de agentes

[H-BUILD](../../harness/build/README.md),[H-CTX](../../harness/context/README.md) y[H-LOG](../../harness/build-log.md) consumen11.03,QA-Evidence/1 yAcceptanceDossier/1; no crean un servicio de permisos. H-CTX contiene nueve paquetes con docCommit/hash/versión separados de codeBaseCommit. H-LOG solo contiene plantilla y ejemplo SYNTHETIC/NOT_RUN, sin agente/comando/permiso real inventado.

Recorrido conceptual de una tarea posterior autorizada, **no ejecutado**:

| Paso | Positivo exigido | Negativo/parada y evidencia futura |
|---|---|---|
|1 Alcance| Fuente humana,ticket,rutas,operaciones/coste resueltos | Plantilla,resumen o CIverde no autoriza; AuthorizationReference comprobable |
|2 Contexto| Fuentes directas en docCommit que las contiene | b91ba98 no contiene09/10/11.03; hash cambiado exige releer/invalidar revisión |
|3 Aislamiento| Git/base/HEAD y trabajo previo contrastados | Sin stash/reset/limpieza ajena; diff/base/propiedad trazados |
|4 Delegación| Rutas disjuntas,salidas estables y escritor único | Dos autores activos detienen la ruta; cese/transferencia verificables |
|5 Código| TaskEnvelope y permiso cubren TDD y contratos | Editar fuente aprobada para pasar exige detener y preparar ContractChange |
|6 Prueba| Casos±,datos/semilla,toolchain y entorno reales | Mock no acredita PGreal; requerido omitido no es PASS; comandos/salidas exactos |
|7 Evidencia| Nuevo runId,expected/observed y artefactos/hash/acceso | Exit0 o “agente terminó” no es oráculo; FAIL/BLOCKED/NOT_RUN se conservan |
|8 Revisión| Revisor independiente coteja código/caso/fuente/evidencia | P0/P1 o crítico faltante bloquea; contraejemplo no se borra |
|9 Corrección| Fallo→fix→rerun enlazados | No cambiar esperado al output ni sustituir el intento fallido |
|10 Publicación| Permiso Git y automatizaciones comprobados | Push con deploy fuera de permiso se detiene antes; sin bypass o secretos en PR externos |
|11 Gate| Dossier técnico separado de aceptación humana | Build/checkverde/Markdown aprobado no acepta hito o proveedor |
|12 Continuidad| Releer PLANS/HLOG/Git/fuentes/permisos efectivos | Interrupción no repite spike/proveedor automáticamente; conserva unknown y pendientes |

Un agente nuevo contrasta commit/hash/diff/casos/artefactos y cese antes de continuar. La eficacia real del proceso sigue POR_VERIFICAR: el harness define guardas,revisión y evidencia; su lectura no demuestra que los agentes las aplicarán correctamente.

## 11. Hallazgos P0–P3

No se demostró P0/P1 ni bloqueo de cierre documental o PH01. Las observaciones siguientes son propuestas, no correcciones aplicadas.

### AUD6-F01 — Perspectiva lógica del oráculo IA

| Campo | Registro |
|---|---|
| Severidad/alcance | P2 / NON_BLOCKING. Precisión acotada antes de validar oráculo/búsqueda PH05T01–T02; sin bloqueo PH01 |
| Archivo/sección/líneas |09.01§7 L91 yF-Q1-15 L187;08.01§7 L74;03.04§7 L81–88;ACC09-IN/OUT03.01 L205–210 |
| Decisión/contrato |PRE-03.D3/PRE-08.B1/PRE-09.B2; perspectiva Negamax,State.turn terminal youtcome/winner;F-BUS-01/02/F-Q1-15 |
| Escenario/impacto | White gana inmediatamente; target conserva turn=White. Evaluar hijo+M desde State.turn y negar en padre da−M; perspectiva lógica Black da hijo−M/raíz+M. Es un escenario de contrato, no ejecución observada |
| Evidencia/juicio |09.01 dice “jugador al turno de hoja”;08.01 prohíbe usar el turno retenido y09.01L187 exige signos correctos. Existe lectura compatible; no contradicción normativa ni implementación defectuosa demostrada |
| Corrección mínima propuesta | Hacer explícita la perspectiva recursiva en el oráculo; fixtures espejo White/Black,target/equal/atsyrau y profundidad impar/par. Editar09.01 solo con autorización de fuente; no cambiar reglas |
| Owner/dependencias |I/Q,consultaM;03.04/08.01,F-BUS-01/02 yPH05T01–T02 |
| Verificación/gate | Valor raíz correcto contra exhaustivo,sinTT yconTT de orden; no deducir perspectiva del turno terminal. Pruebas PH05,08.01AC01/07 y09.01AC09 antes de aceptar búsqueda |
| Estado/disenso |PROPUESTA_DE_PRECISIÓN. A propuso bloqueo acotado PH05; C descartó contradicción; D confirmó lectura compatible. Coordinación deduplica NON_BLOCKING porque08.01 ya fija la regla explícita; no exige otra decisión fundamental |

### AUD6-F02 — Ejemplo de ruta singular

| Campo | Registro |
|---|---|
| Severidad/alcance |P3 / NON_BLOCKING |
| Archivo/sección/líneas |10.01§11 L129 frente05.02§7 L85/L132/L139 |
| Fuente/contrato |PRE-10.A3; catálogo05.02 sirve/api/v1/games y conserva método/ruta |
| Escenario/impacto | Copiar /api/v1/game/... en un rewrite apunta a ruta inexistente; el catálogo completo resuelve la intención. No concede permisos ni cambia autoridad |
| Corrección mínima propuesta | Cambiar solo el ejemplo a /api/v1/games/...; requiere autorización para editar10.01 aprobada, no cambiar catálogo/contratos |
| Owner/dependencias |A,consultaS y05.02; consumidor PH02/operación |
| Verificación/gate | Ejemplo coincide con operación catalogada y conserva auth completa; revisión editorial. Sin bloqueo de cierre,PH01 o release por sí solo |
| Estado |PROPUESTA;10.01 intacta |

### Contrastes descartados y límites

Historia/apariciones ybarrier/cleanup están separados expresamente; no se reabren L5-R01/R04. Estados históricos BLOCKED del spike o EN_REVISIÓN no revocan recibos posteriores. Runtime/cuota desconocida,assets sin licencia yRPO/RTO sin medir no son P0/P1 del motor.

El disenso sobre F01 se conserva con su escenario y juicio coordinado. No se lo convierte en decisión normativa o modificación de fuentes.

## 12. Condiciones pendientes

Son condiciones conocidas confirmadas, no nuevos requisitos P0/P1. Cada acción conserva su permiso; un doble no acredita una fase completada.

| ID/severidad/alcance | Evidencia,escenario e impacto | Owner/dependencias | Tratamiento,verificación y gate |
|---|---|---|---|
|AUD6-C01 P2 BLOCKS_LATER_PHASE|L4-P01;08.02§8L89–129,08.03§8L108–114,08.04§8L102–120. Usar ramaAI como gameId oficial o camposai/recovery nuevos rompe catálogo/permiso|S/I/C,V/Q;05.02/05.03/07.02 y04/08/09|ContractChange→aprobación específica→versionado autorizado→compatibilidad/negativos→código autorizado. GatePH05T05/T06 y consecuencias |
|AUD6-C02 P2 BLOCKS_LATER_PHASE|RISK10;07.03§7L151–171 yPH06L46–50. Cleanup sinpolicy oACK tratado como COMPLETE inventa retención/supresión|Propietario/S/A;02.03/05.03/07.03/10.03|Política aprobada,versión no nula/inmutable y ensayos recuperables. Bloquea cleanup/retenciones afectados yPH06 completo; no barrier legítima niPH01 |
|AUD6-C03 P2 BLOCKS_RELEASE_ONLY|RISK13,04.02§8/12 y08.04§12; activos/textos/cultura sin permiso no pueden publicarse|Propietario/V/Q;02.03/04.02/08.04|Permisos,atribución,licencia/revisión cultural y activos finales antes de publicar lo afectado,PH03/07. Fixtures propios sintéticos son independientes |
|AUD6-C04 P2 BLOCKS_RELEASE_ONLY|10.03§9L83–91/§11L104–111. Backup sin verdad posterior puede revivir auth/efectos; RPO/RTO/SLO sin medición|A/S/Q;05.03/07.02/07.03/09.04/10.04|Restore aislado,cortes/receipts/revocaciones/compatibilidad y muestras; cuarentena hasta demostrar reapertura. GatePH06T08/PH07T06 yrelease,no núcleoGo |
|AUD6-C05 P2 BLOCKS_LATER_PHASE|01.02/01.03/10.01§9/10.02§7. Capacidad,protección,SDK/runtime/cuotas del spike no siguen vigentes por defecto|A/Q/S; consumidor y permisos externos|Revalidar versión/plan/saldo/protección/coste/retirada antes de operación externa o claim medido. Sin upgrade automático; no bloquea PH01 local |
|AUD6-C06 P2 NON_BLOCKING|11.03§7–8/PH01 entrada. Formatos propuestos y rutas concretas requieren alcance posterior; plantilla no autoriza|Propietario/C/M/Q;TaskEnvelope/CaseSpec/CorpusEntry/QA-Evidence|Autorizar ticket,rutas,pruebas/operaciones y parámetros propuestos consumidos; fuente/versión comprobables. Gate de entradaPH01,decisión humana prevista |

Delimitación PH05/C01:

- T01–T03: búsqueda pura independiente. T04: camino oficial vigente sin campos nuevos.
- T05: prepara el cambio; no lo aprueba. T06: esperaT03/T04 yT05 aprobado.
- T07: puede preparar modelos puros de timeline/cursor/generación/fixtures; PracticeRecord/restauraciónHTTP/AI espera el delta.
- T08: catálogoES/EN,predicadosGo,progreso/UI pueden prepararse con dobles declarados; ObjectiveEvaluation,hintsIA/metadatos nuevos esperan contrato.
- T09: medición de búsqueda pura puede ser independiente; dossier/E2E completo esperaT06–T08 integrados.

## 13. Preparación de PH-01

PH01 puede aislar corpus/tipos/motor de cuentas,PostgreSQL,WebGL,Ably,IAcompleta ydespliegues. La entrada documental existe; no hay código ejecutado que aceptar.

**Primer ticket recomendado: PH-01-T01.** Congelar CaseSpec/CorpusEntry para14ACC con input/action/expected/contexto/versiones/clase/alcanzabilidad y referencia N1/AD-NORM-001. Un revisor independiente contrasta oráculos. Los sintéticos siguen UNPROVEN; una trayectoria legal necesita prueba aparte. No ejecutar un motor inexistente para fabricar esperados.

**Primer Go: PH-01-T02, tras T01 aceptado.** Tipos/versiones,validación estructural,fábrica inicial,copia profunda y unitariasTDD. Siembra/captura/tuz llegan enT03/T04; no endpoints,Gin,SQL,tokens,AI o frontend.

| Parte del permiso futuro | Alcance candidato que debe aprobarse |
|---|---|
| Fuentes/entrada |AUD aceptado,reciboPH00 sin repetir spike,03.01–04/ADR0002/01.01/02.01/09.01/09.03/09.04/11.03,commit documental publicado y formatos propuestos aceptados para el ticket |
| RutasT01 |modules/engine/testdata/tk-n1-1.0.0/corpus.json yREADME.md del corpus; formato/nombres exactos en TaskEnvelope. No creados aquí |
| RutasT02 |modules/engine/go.mod,state.go,state_test.go y pruebas del mismo módulo; nombres candidatos a autorizar. Sin segundo motor ni modificación de specs |
| Herramientas |Fijar versiónGo compatible y dependencias antes de usarlas; instalación/download solo si el permiso los cubre. No heredar lockfiles del spike |
| Pruebas |Inicial18×9/162,blancas,kazanes0,sintuz;18PitRef únicos/ciclo; dimensiones/rangos/versiones/total/tuz inválidos rechazan antes de indexar/convertir; revisión máxima sinwrap; copia profunda |
| Evidencia/cierre |RED observado→GREEN/refactor/regresión local; comandos/salidas/versión/semilla/fixtures,QA-Evidence yrevisión.90% es objetivo del hito completo, no sustituto de críticos del incremento |
| Operaciones |T01:fixtures y comprobaciones locales JSON/aritmética/lectura. T02:código/tests locales acotados. Git/publicación según permiso; no migración,contenedor,proveedor,workflow,pago odeploy por inferencia |

Negativos iniciales:163/161piedras,índices extremos,turno/reglas desconocidos yaliasing. El oráculo usa fuente/copia, no la implementación bajo prueba. Fixtures/formatos candidatos se resuelven dentro del alcance autorizado, sin reabrir el deporte.

## 14. Decisión formal

**READY_WITH_SCOPED_CONDITIONS**

Base suficiente para solicitar autorización de PH01T01 y primer GoT02 tras aceptar el corpus. No se demostró P0/P1 que bloquee PH01 o el cierre documental. C01–C06 yF01/F02 tienen propietarios,consumidores,pruebas y gates; no impiden ese primer trabajo independiente.

A yD consideran suficiente solicitar PH01 delimitado; B recomienda conservar explícitamente condiciones posteriores. Coordinación aplica el veredicto acotado al conjunto de fases, sin desacuerdo sobre la viabilidad documental del primer incremento.

Estar preparado para autorizar implementación no significa estar preparado para producción. El informe no certifica seguridad,normativa federativa,disponibilidad,coste,rendimiento orestauración.

## 15. Acciones recomendadas

1. Revisar y,si procede,aceptar AUD-01 sobre el HEAD publicado delPR. Mantener separadas las propuestas correctivas.
2. Autorizar PH01T01 y,si se incluye expresamente,T02 después de aceptar corpus: fuentes/commit,rutas,formatos,pruebas,toolchain,operacionesGit ylímites. No pedir confirmación repetida para pasos ya cubiertos.
3. Congelar oráculos y demostrar tipos,fábrica,validación ycopias porTDD local/revisión. CompletarT02 no autoriza PH02 oexternos.
4. Continuar PH01T03–T08 solo dentro de autorización que los incluya:14ACC,propiedades/fuzz/race aplicable,replay/cobertura yfronteraGame.
5. Antes de PH05 afectado,tratarC01 yF01; antes de cleanup,C02; antes de publicar activos,C03; antes de reapertura/release,C04 yC05. F02 es corrección editorial propuesta con permiso de fuente, sin reiniciar planificación.

No se recomienda rehacer normativa,arquitectura,spike oplanificación completa. Las correcciones de fuentes aprobadas se presentan por contenido/consumidores; esta auditoría no las ejecuta.

## 16. Evidencias

| Participación real / método | Cobertura original y límites |
|---|---|
|A:l6_architecture_audit|13specs primarias íntegras:01.01–03/02.01/03.01–04/05.01–04/09.01;11.01/02,PH00/01,ADR0001/2/3/4/8,PRE pertinentes yconsumidores. F01 propuesto; sinP0/P1 |
|B:l6_security_audit|11specs primarias íntegras:02.03/07.01–04/09.03–04/10.01–04; también05.02/03/11.03,PH01/06/07,PRE07/09/10 yADRseguridad. F02 observado; sinP0/P1 |
|C:l6_experience_ai_audit|13specs primarias íntegras:02.02/04.01–04/06.01–03/08.01–04/09.02;PRE04/08,11.01/02,PH03/05 ycontrastes03/05/07. L4-P01 acotado; sinP0/P1 |
|D:l6_agentic_audit|3specs primarias íntegras:11.01–03;8PH/3harness/encargo/PRE11;recuento ydependencias independientes,contraste03/08/09/10. Proceso/permisos/oráculo revisados |
|Coordinación|Integridad80Markdown/246PRE/40specs,DAG40/80,495tablas,14ACC idénticos; spike informes/JSON granulares,PH01/11/harness yjuicio deduplicado. Único escritor deAUD/registros |
|Revisión de esta entrega|B leyó AUD íntegro y los cuatro deltas de registros: CONFORME; dos referencias editoriales P3 corregidas y revalidadas ADDRESSED. Resultado en PLANS; no sustituye aceptación humana |

Cobertura primaria completa colectiva:13+11+13+3=40specs; lecturas cruzadas no aumentan ese total. Los auditores recuperaron rangos truncados antes de declarar documentos íntegros. Las lecturas selectivas deREG/maestro/spike se declaran como tales, sin atribuir a todos lectura integral de todo el repositorio.

Métodos:Gitstatus/log/fetch/ls-remote/show/diff,listas de archivos versionados y análisis de texto/JSON en memoria. Ningún verificador en disco, test de producto,instalación,cuenta,credencial,recurso odespliegue. Fallos históricos del spike conservados, no nuevas ejecuciones.

Fuentes protegidas sin diferencias contra00c8e12. En la carpeta original,AGENTS modificado ymanual sin versionar conservan SHA256 antes/después, sin stash/copia/publicación. Se reutiliza la worktree aislada existente. [PLANS](../../PLANS.md#ejecucion-lote-6) conserva reciboPR/HEAD ycontroles finales.

## 17. Gate humano

AUD-01 permanece **EN_REVISIÓN**. Gate solicitado: **APRUEBO LOTE 6**.

No fusionar el PR en esta tarea. La autorización de auditoría termina al publicar/verificar la entrega. Ni el veredicto,ni aprobarAUD,ni un checkCI autoriza programar.

La autorización posterior identificará al menos fase/ticket,archivos exactos,fuentes/contratos/versiones,pruebas,operaciones ylímites; herramientas,Git orecursos externos cuando corresponda. L4-P01,retención ydemás propuestas conservan aprobación específica pendiente. No iniciar PH00/PH01 oproducto automáticamente.
