# H-CTX — Paquetes mínimos de contexto verificable

> **Aprobación posterior — 2026-10-26: APROBADO documentalmente.** El propietario declaró «APRUEBO LOTE 5 como documentación técnica.» sobre HEAD 9ff1a2857d0b91b9cf7925a4c36de30e377cef3d; PR #5 integrado por squash 6579ca930a9081f773b390ce61f6bc80cc6b1ae7. El [recibo de aprobación](../../PLANS.md#cierre-lote-5) prevalece sobre los estados de revisión históricos. L4-P01, la política de retención y demás propuestas mantienen su aprobación específica pendiente; sin permiso de implementación, workflows, proveedores o lote6.

ID: H-CTX · Versión: 0.2 · Fecha: 2026-10-26 · Estado documental: APROBADO · Owner: C — Coordinación · Lote: 5.

Diseño documental PROPUESTA de contexto para trabajo futuro; fuentes APROBADO y comportamiento POR_VERIFICAR. No genera paquetes/archivos auxiliares, no duplica repositorios y no autoriza implementación. Un brief/resumen orienta lectura, pero no sustituye una fuente ni transmite permisos.

## Núcleo común y selección por ticket

Todo paquete lee [AGENTS](../../AGENTS.md), [PLANS vigente](../../PLANS.md#ejecucion-lote-5), secciones pertinentes de [REG](../../docs/decision-register.md)/[SRC](../../docs/sources/TOGUZ_PROMPT_MAESTRO_CODEX.md)/[CONF](../../docs/conflicts.md), [11.01](../../docs/specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](../../docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [11.03](../../docs/specs/11-orchestration/11.03-milestones-gates-and-evidence.md). Incluir el PH/ticket exacto de [fases](../../docs/index.md#phase-00), contrato directo productor/consumidor y aceptación QA aplicable. Las 40 specs no se vuelcan en cada brief; la revisión documental global de entrada 11.03 es un gate distinto de la lectura acotada del implementador.

Consumir TaskEnvelope/1/AuthorizationReference/1 de 11.03, QA-Evidence/1 de [09.01](../../docs/specs/09-quality/09.01-engine-data-contract-and-ai-verification.md) y AcceptanceDossier/1 de [09.04](../../docs/specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md). Este documento define una vista de selección/procedencia del contexto, sin cambiar esos campos o crear otra API/schema obligatoria de producto.

| Parte del paquete | Datos que el autor debe resolver |
|---|---|
| Tarea/alcance | ticketId/phaseId, resultado/AC, owner/revisor, componentes, rutas propias/solo lectura, dependencias/salidas aceptadas y límites de parada. |
| Revisión documental | Repo/remoto, docCommit efectivo que contiene fuentes, estado/aprobación de cada documento y versión, fecha de lectura y hash SHA256 de bytes leídos por ruta. |
| Revisión de código | codeBaseCommit/branch/HEAD y diff autorizado separados del commit documental; motor/contratos/builds/migraciones por consumidor si ya existen. |
| Fuentes | PRE/spec/ADR/AC y secciones exactas con enlaces; contratos/versiones/dueño/consumidores y conflictos aplicables. |
| Pruebas | CaseSpec/caseIds, positivos/negativos/oráculos/fixtures/semillas/procedencia, evidencia requerida y recursos/aislamiento. |
| Permisos | Tres AuthorizationReference/1, fuente humana accesible y alcance; aclarar cuál acción requiere cuál permiso y siguiente acción permitida. |
| Procedencia/limitaciones | Fuente confirmada o working-tree/hash, contenido ausente/inaccesible, borrador/propuesta, dependencia parcial y revisión invalidada. |

Mantener el paquete en conversación o ruta expresamente autorizada; este lote no crea JSON/manifiestos/scripts/briefs nuevos en disco. Adjuntar solo excerpts mínimos útiles con referencia, nunca secretos ni cuerpos privados de evidencia.

## Verificar commit, versión y hash antes de usarlo

1. Resolver repositorio/base/branch reales y fuente humana de permiso; consultar ledger H-LOG/PLANS para conservar trabajo ya hecho. No inferir estado de una sesión anterior ni del nombre de rama.
2. Identificar docCommit publicado que contiene cada fuente y su aprobación aplicable. La base b91ba98a380ea12b8b6ffc24155edd437e958ed0 es la base histórica del lote 5; no contiene las nueve nuevas specs 09/10/11.03. Los documentos de este lote están EN_REVISIÓN hasta gate humano; un hash no los aprueba.
3. Leer fuente directa completa en las secciones que gobiernan ticket/contrato; calcular/registrar SHA256 sobre bytes exactos y versión/metadatos, no sobre resumen o texto reformateado. Si está sin commit, declarar working-tree+hash y base; no fabricar un SHA que contenga el contenido. Para futura implementación resolver autorización de esa entrada y la aprobación documental requerida 11.03.
4. Comparar paquetes recibidos con fuentes Git/diff actuales: commit, versión, hash, estado y permiso. Si cambia cualquiera que afecte contrato, congelar consumidor, releer delta y reevaluar compatibilidad/pruebas/revisión. Un SHA de motor igual no cubre configuración distinta.
5. Validar enlaces/acceso a artefactos y fuentes; si revisor no accede, registrar límite/BLOCKED en el caso requerido, sin convertir falta de datos en valor cero o PASS.
6. Al entregar, incluir revisión usada, archivos leídos/modificados y fuentes cambiadas. Antes de reasignar confirmar cese y nuevo autor; conservar paquetes previos como referencias, no duplicar todo el repositorio.

Ejemplo sintético de referencias futuras, NO EJECUTADO ni autorización:

```text
ticketId: <PH-NN-Tnn autorizado>
docCommit: <SHA publicado que realmente contenga fuentes aprobadas requeridas>
docSource: <ruta> / version <v> / estado <estado real> / sha256 <bytes leídos>
codeBaseCommit: <SHA real de base de implementación>
branch: <rama autorizada>
authorizationRefs: <referencias humanas verificables documental/código/externo>
si docCommit o aprobación faltan: entrada no validada; no sustituir con b91ba98
```

## Nueve paquetes por dominio

Cada fila añade al núcleo común las fuentes directas indicadas; sus dependencias exactas están en§4/IDX y se incorporan cuando el comportamiento del ticket las consume. Reducir contexto no permite omitir una frontera de seguridad/estado. No copiar contratos completos al brief como fuente canónica alternativa.

| Dominio / PH | Fuentes mínimas directas | Decisiones/contratos y dependencias que deben retenerse | Restricciones y casos/evidencias pertinentes |
|---|---|---|---|
| Motor / PH01 | [03.01](../../docs/specs/03-game-engine/03.01-normative-rules-and-corpus.md), [03.02](../../docs/specs/03-game-engine/03.02-state-indices-and-invariants.md), [03.03](../../docs/specs/03-game-engine/03.03-moves-and-semantic-events.md), [03.04](../../docs/specs/03-game-engine/03.04-outcomes-repetition-and-replay.md), [ADR-0002](../../docs/adr/ADR-0002-single-versioned-go-engine.md), 09.01 | PRE-03 y PRE-09.B1; AD-NORM-001/TK-N1-1.0.0/14 ACC; State/PitRef/Transition/PositionKey/contexto/eventos/replay. 01.01 y 02.01 delimitan integración/alcance. | Puro/determinista/162; sin SQL/HTTP/reloj/motor TS; corpus por clase/alcanzabilidad. F-Q1-N01–N14/01/02; carreras Game no se acreditan con unitarias. |
| Backend / PH02 | [05.01](../../docs/specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md), [05.02](../../docs/specs/05-backend/05.02-http-openapi-and-errors.md), [05.03](../../docs/specs/05-backend/05.03-transactions-data-and-migrations.md), [05.04](../../docs/specs/05-backend/05.04-event-schemas-and-postcommit-publication.md), [01.01](../../docs/specs/01-infrastructure/01.01-stack-monorepo.md), 07.02 | PRE-05/PRE-09.B3/B4/C1; OpenAPI/Problem/Receipt/JCS/locks/PG/postcommit; 03, motor, 07.01, sesión, 07.04 si afecta Elo. | Tres servicios/hexagonal/dos dueños SQL/AI sin DB; F-Q1-03–08/11/13/14. PG real, history ordinal separado de aparición y commit unknown sin rollback supuesto. |
| Identity / PH02/06 | [07.01](../../docs/specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md), [07.02](../../docs/specs/07-identity-data/07.02-tokens-authorization-and-service-identity.md), [07.03](../../docs/specs/07-identity-data/07.03-profiles-history-and-data-rights.md), [02.03](../../docs/specs/02-product/02.03-accounts-history-and-licensing.md), 05.02/05.03 | PRE-07.A–D/F y contratos guest/OIDC/PKCE/refresh/proof/link/sesión/JWKS/CSRF/recurso/export/borrado; 06.03, horizonte y 10.03, retención. Leer IDs individuales exactos en REG según tarea. | Alias/email no prueban identidad; link G→U no transfiere activas/receipts ni Elo. F-Q1-09/10/12; RISK-10 bloquea acciones afectadas, no código independiente autorizado. |
| Multijugador / PH04/06 | [06.01](../../docs/specs/06-multiplayer/06.01-rooms-invitations-and-ranked-lifecycle.md), [06.02](../../docs/specs/06-multiplayer/06.02-ably-notifications-and-reconciliation.md), [06.03](../../docs/specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md), 05.02–05.04, 07.01/07.02/07.04 | PRE-06/PRE-09.C3/D2; Room/Game/Recovery/RematchView, ready/participant, caps/Notice, receipt/scheduler/actividad/24h; 03.04, claims y 01.03, cuotas. | PH04 casual/guest, ranked en PH06, Ably no autoridad/exactly-once; F-Q2-03/04/06/07/08; dos contextos y tercero, PG real/locks/deadline. |
| Frontend / PH03–07 | [04.01](../../docs/specs/04-visual-experience/04.01-navigation-and-accessibility.md), [04.02](../../docs/specs/04-visual-experience/04.02-board-assets-camera-and-materials.md), [04.03](../../docs/specs/04-visual-experience/04.03-interaction-animation-and-audio.md), [04.04](../../docs/specs/04-visual-experience/04.04-performance-and-html-fallback.md), [02.02](../../docs/specs/02-product/02.02-local-flows-and-i18n.md), 05.02/05.04, 06.02/06.03 | PRE-04/PRE-02.B3/D3/PRE-09.D1–D4; PitRef/Simulation/Envelope/eventos/renderer/intención/foco, scheduler; 03.02–04 y 08 si aprendizaje. | HTML primero/equivalencia ES/EN/sin motor TS/offline inferido; licencias/cultura/FPS como metas. F-Q2-01–12 según modo; manual de teclado/lector además de automático. |
| IA / PH05 | [08.01](../../docs/specs/08-ai-practice/08.01-search-heuristics-and-compute-limits.md), [08.02](../../docs/specs/08-ai-practice/08.02-ai-orchestration-fallback-and-metrics.md), [08.03](../../docs/specs/08-ai-practice/08.03-reversible-practice-and-scenarios.md), [08.04](../../docs/specs/08-ai-practice/08.04-tutorial-hints-and-educational-replay.md), 03.04, 05.02/05.03, 07.02 | PRE-08/PRE-09.B2; SearchRequest/Stats/Result, TT solo orden/contexto histórico, grant/callback/efecto, rama/generación/intención/hint. L4-P01 y 11.03, cambio antes de funciones afectadas. | AI sin DB/commit; buscar independiente con permiso; hint cero POST move/progreso, sin antitrampa de copias synthetic. F-Q1-15–18/F-Q2-09, budgets/overshoot/arena/CPU/torneos con incertidumbre. |
| Seguridad / transversal | 07.01–07.03, 05.02/05.03, [09.03](../../docs/specs/09-quality/09.03-ci-security-and-release-controls.md), [10.01](../../docs/specs/10-operations/10.01-deployments-origins-regions-and-secrets.md), [10.02](../../docs/specs/10-operations/10.02-observability-quotas-and-rate-limits.md), [ADR-0005](../../docs/adr/ADR-0005-session-and-service-trust-boundaries.md) | PRE-07/PRE-09.C2/F3; user/S2S/Google/Ably separados, iss/aud/sign/CSRF/Origin/binding/claims/caps/roles/quotas; contrato del recurso afectado y 10.03, restore. | UUID/rol admin/invitación no permiso sobre recurso ajeno, PR externo sin secretos, fallo cerrado; F-Q1-09–12/F-DEP-03/05/06/07, F-OBS-01/03, runids/negativos/DB sin PII. |
| Operaciones / PH00/02/07 | [01.02](../../docs/specs/01-infrastructure/01.02-platform-feasibility.md), [01.03](../../docs/specs/01-infrastructure/01.03-cost-capacity-evolution.md), 10.01/10.02, [10.03](../../docs/specs/10-operations/10.03-backups-restoration-and-retention.md), [10.04](../../docs/specs/10-operations/10.04-releases-incidents-and-portability.md), 09.04 | PRE-10; Deployment/Region/Secret/Observation/Budget y Release/Compatibility/Incident/RecoveryDecision/restore; 05.03/07.02/07.03, integridad/retención. | Cuatro proyectos, tres Go, permisos externos/cuotas 50/75/90/no autopago; backup no restaurable sin ensayo, rollback de código≠datos. F-DEP/F-OBS/F-REL/F-INC y casos 10.03. Tiempos medidos, no RPO/RTO garantizados. |
| QA / transversal | 09.01, [09.02](../../docs/specs/09-quality/09.02-ui-e2e-accessibility-and-visual-verification.md), 09.03, 09.04, 10.04 y 11.03; fuente de producto directa de cada oráculo | PRE-09; CaseSpec/Corpus/QA-Evidence/AcceptanceDossier/Measurement/RiskAcceptance; matrices 14 ACC/pirámide/CI/severidad/gates. | 90/80 más críticos, PG real/acceso manual, datos reproducibles; PASS observado y artefactos accesibles, omitido no PASS. P0/P1/críticos bloquean, P2 aceptado y acotado; Q verifica recuperación, A operación, C gobierno. |

Las abreviaturas de IDs en esta tabla remiten al documento enlazado y al inventario, no permiten omitir lectura del archivo correspondiente; el TaskEnvelope enumera rutas resueltas/IDs PRE-NN.Xn exactos. Referencias por grupo no reasignan propiedad canónica de decisiones. Hitos posteriores consumen también fuentes de derechos/Elo/IA aunque una tarea temprana solo prepare una interfaz.

## Cambios de contexto y transferencia

Si un productor está parcial/incompatible, no cerrar consumidor sobre un contrato inventado. Coordinación indica dependencia/estado/restricción y permite solo partes independientes. Recibir una entrega estable no la vuelve APROBADO; cotejar autorización/estado y cambiar paquete solo tras revisar contenido.

Una transferencia contiene ticket/resultado/diff, fuentes/hash/versiones/commits, comandos/resultados observados/artefactos, pendientes/bloqueos, autor/revisor, permisos/siguiente acción y cese de escritura. Receptor verifica contra Git/archivos reales y [H-LOG](../build-log.md); no confía exclusivamente en resumen. Recuperar contexto no implica repetir spike/pruebas externas, ni dar por ejecutadas instrucciones previstas. El agente que recibe no escribe hasta tener propiedad única y permiso.

Para derechos de datos, conservar expresamente la separación de 05.03/07.03: petición de borrado autenticada/autorizada→revocación y barrier S2S sin policyVersion→ACK_REVOKED durable; sin política, raíz/barrier quedan BLOCKED_PENDING_DATA_POLICY. Cleanup exige policyVersion aprobada no nula e inmutable; null no se envía, cambio de política en el mismo paso produce conflicto. La barrera no significa supresión completa ni requiere esperar la política de cleanup.

## Verificación del paquete

| Caso de proceso | Positivo | Negativo y salida |
|---|---|---|
| HCTX-C01 / F-GOV-03 | Fuentes directas/versión/hash/docCommit coinciden y permisos cubren ticket. | SHA/hash/estado cambió o fuente inaccesible→bloquear uso dependiente y releer delta. |
| HCTX-C02 / F-GOV-02/06 | Tres permisos y ContractChange específico vigentes cubren acciones. | GateDOC5 o resumen de agente no aprueba L4-P01/código/externos; detener acción afectada. |
| HCTX-C03 / F-GOV-07/09 | Paquete acotado conserva decisiones/fronteras/negativos e identifica NOT_RUN. | Brief sin contrato directo, synthetic como alcanzable o pruebas de cada SHA mezcladas no permite cerrar. |
| HCTX-C04 / F-GOV-08/11/12 | Receptor contrasta evidencia/cese y conserva límites del spike/restore. | Mensaje«terminado», hash sin oráculo, backup sin ensayo o nuevo escritor sin cese bloquea transferencia. |

Estas verificaciones futuras usan QA-Evidence/1 para observar el proceso cuando esté autorizado. Actualmente solo se revisa documentación. Fallos/pedidosdecambio se devuelven a coordinación yautor; no cambiar silenciosamente fuente aprobada ni crearun cuarto servicio de contexto.
