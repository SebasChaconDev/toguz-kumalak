# PH-05 — IA, práctica reversible, tutorial y pistas

ID: PH-05 · Versión: 0.1 · Fecha: 2026-10-26 · Estado documental: EN_REVISIÓN · Owner: C — Coordinación · Lote: 5.

Plan futuro PROPUESTA; fuentes APROBADO, resultados POR_VERIFICAR. L4-P01 permanece PROPUESTA pendiente de aprobación de contratos; la búsqueda pura independiente puede avanzar bajo su permiso auténtico sin dar por aprobadas las funcionalidades afectadas.

## Ficha Milestone/1

| Campo requerido | Contenido del hito |
|---|---|
| Identificador | PH-05 |
| Objetivo | Entregar rival fácil/normal legal y acotado, práctica reversible, tutorial/replay y ayuda voluntaria sin efecto deportivo automático. |
| Dependencias | PH-04; 08.01–08.04 y sus entradas 03.02–03.04/01.02/02.02/05.02/05.03/07.02; 04.01–04.04; 05.04/06.02/06.03/07.01; 09.01–09.04; 10.02/10.04; 11.03. |
| Responsable | I búsqueda/aprendizaje, S orquestación/seguridad, V acceso y Q verificación; C coordina, propietario acepta y decide ContractChange separado. |
| Entrada | Fuentes/gate 11.03 y permiso de código por ticket; para delta L4-P01 aprobación expresa, actualización/versionado autorizado de contratos y pruebas de compatibilidad antes de implementar sus consecuencias. |
| Entregables | Negamax/heurística/budgets/TT segura, orquestación/fallback Game, rama/undo/redo/guardado, escenarios/tutorial/hints/replay equivalentes HTML/3D. |
| Pruebas positivas/negativas | PH05-C01–C08, F-Q1-15–18/F-Q2-02/09 y casos08/07.02 de cada contrato. |
| Riesgos | Signo/perspectiva, TT sin historia, memoria no contada, cálculo/callback inciertos, generación vieja, hint convertido en move y contrato pendiente. |
| Evidencia obligatoria | QA-Evidence/1, árboles de búsqueda y semillas, CPU/nodos/arena/overshoot, grants/receipts, red/timeline/progreso y revisión de cambio identificada. |
| Salida | Funciones completas con límites/negativos observados, delta aprobado cuando necesario, sin P0/P1/críticos pendientes ni fuerza/Elo declarados sin medir. |
| Aprobación humana | AcceptanceDossier/1 y GateEvaluation/1 del alcance; aceptación de lote 5 o permiso genérico IA no aprueba L4-P01 por inferencia. |
| Corrección/recuperación | Cancelar/incrementar generación, conservar último estado válido, revalidar en Game; no repetir cálculo o movimiento incierto a ciegas, corregir y enlazar run nuevo. |

## Fuentes y contratos consumidos

Leer completos [08.01](../specs/08-ai-practice/08.01-search-heuristics-and-compute-limits.md), [08.02](../specs/08-ai-practice/08.02-ai-orchestration-fallback-and-metrics.md), [08.03](../specs/08-ai-practice/08.03-reversible-practice-and-scenarios.md), [08.04](../specs/08-ai-practice/08.04-tutorial-hints-and-educational-replay.md) y [CONF L4-P01](../conflicts.md#lote-4-riesgos). Fuentes PRE-08 y PRE-09.B2 del [REG](../decision-register.md) preservan legalidad/límites/acceso sin motor duplicado.

Entradas canónicas: [03.02](../specs/03-game-engine/03.02-state-indices-and-invariants.md)–[03.04](../specs/03-game-engine/03.04-outcomes-repetition-and-replay.md), [01.02](../specs/01-infrastructure/01.02-platform-feasibility.md), [02.02](../specs/02-product/02.02-local-flows-and-i18n.md), [05.02](../specs/05-backend/05.02-http-openapi-and-errors.md)/[05.03](../specs/05-backend/05.03-transactions-data-and-migrations.md)/[05.04](../specs/05-backend/05.04-event-schemas-and-postcommit-publication.md), [07.01](../specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md)/[07.02](../specs/07-identity-data/07.02-tokens-authorization-and-service-identity.md), [06.02](../specs/06-multiplayer/06.02-ably-notifications-and-reconciliation.md)/[06.03](../specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md) y [04.01](../specs/04-visual-experience/04.01-navigation-and-accessibility.md)–[04.04](../specs/04-visual-experience/04.04-performance-and-html-fallback.md).

QA/operación/gobierno: [09.01](../specs/09-quality/09.01-engine-data-contract-and-ai-verification.md)–[09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md), [10.02](../specs/10-operations/10.02-observability-quotas-and-rate-limits.md)/[10.04](../specs/10-operations/10.04-releases-incidents-and-portability.md), [11.03](../specs/11-orchestration/11.03-milestones-gates-and-evidence.md) y [PLANS](../../PLANS.md#ejecucion-lote-5).

Consumir SearchRequest/Result/Stats y perfiles 08.01, AIRequest/grant/claim/callback de 07.02/08.02, PracticeSaveEnvelope/rama/generación/contexto 08.03 y contenido/progreso/replay 08.04. RequestBranchAI/ReadBranchAI/CancelBranchAI y recovery=fallback son delta propuesto, no rutas aprobadas por este plan. ContractChange/1 debe identificar fuente/hash, propuesta completa, alternativas, binding/recurso/intención/cuotas/persistencia/metadatos, consumidores/seguridad/datos/coste, compatibilidad, reversibilidad y aprobación del contenido/rutas; actualizar 05.02/05.03/07.02 solo después bajo autorización de edición. Luego obtener/aplicar permisos de código/externos que cubran la implementación.

## Componentes futuros y tickets ordenados

Rutas candidatas: services/ai para búsqueda sin DB, modules/engine compartido, services/game para admisión/orquestación/efecto y apps/web para rama/UI; contratos generados según cambio autorizado. No cuarto servicio, worker/cola/Redis preventivo ni estado durable solo en RAM.

| Ticket / depende de | Pasos y razón | Entrega pequeña |
|---|---|---|
| PH-05-T01 / entrada | Fijar árboles/corpus pequeño y exhaustivo de referencia; implementar evaluación versionada con perspectiva al turno y terminal separado. | Oráculo de signos/features antes de optimizar; C01. |
| PH-05-T02 / T01 | Negamax alfa-beta/profundización/orden estable, última iteración completa y candidato Go legal barato. | Equivalencia a profundidad completa, sin rendición ni draw por count3; C01/C02. |
| PH-05-T03 / T02 | Límites cooperativos tiempo/nodos/depth/arena total y cancelación; TT solo orden con origen revalidado, nunca scores/bounds/terminales. | C02/C03 y contextos por rama, sin truncar historia. |
| PH-05-T04 / T03 | Integrar modo IA oficial según contratos vigentes: tras commit humano/inicio IA, Game solicita cálculo fuera de lock y revalida al aplicar. | Legalidad/efecto único del camino ya autorizado; C04. Recovery nuevo espera T05. |
| PH-05-T05 / entrada contractual | Preparar/revisar ContractChange L4-P01 con S/I/V/Q/A y propietario; tras aprobación identificada actualizar/versionar fuentes/consumidores y probar compatibilidad. | Dossier de cambio; sin aprobación este ticket queda preparación y tickets afectados bloqueados, C05. |
| PH-05-T06 / T03+T04+T05 aprobado | Implementar rama/generación/intención/claim/lectura/cancelación y recuperación/fallback afectados por delta aprobado. La recuperación oficial consume la orquestación y el efecto Game entregados en T04. | C04/C05/C06, hint/opponent_move distintos. |
| PH-05-T07 / T06 | Undo/redo un movimiento completo, volver a mi turno, bifurcar solo al éxito, reset/guardado separado/restauración Go y escenarios con objetivo/procedencia. | C06/C07; sin historia/ratings oficiales. |
| PH-05-T08 / T07 | Tutorial opcional seis lecciones, legal-vsobjetivo, reglamento completo, hints voluntarios y replay educativo HTML/3D/ES/EN. | C07/C08; ayuda no selecciona ni mueve. |
| PH-05-T09 / T04+T08 | Torneos reproducibles/colores/semillas, límites y E2E/manuales; revisar dossier con diferencia fácil/normal y límites honestos. | Evidencia de legalidad/recursos/diferencia, sin Elo/fuerza inventada. |

T01–T03 no dependen del delta ni autorizan por sí mismos otras tareas. T05 puede prepararse en paralelo solo con rutas/entregas independientes y un autor por archivo; no consume aprobación imaginada. Para cada ticket separar preparación de contrato, código y proveedor. El hito completo no cierra con las funciones afectadas sin implementar.

## Casos y evidencia futura

| Caso / QA | Positivo esperado | Negativo esperado | Artefacto adicional |
|---|---|---|---|
| PH05-C01 / F-Q1-15 | Exhaustivo pequeño y Negamax completo coinciden con colores/hojas alternos y terminales. | Ganador no se infiere del turno, signos invertidos o iteración parcial fallan. | Árbol/valores/features/versiones y comparación sin TT. |
| PH05-C02 / F-Q1-16 | H1/H2 misma raíz/count con historia futura distinta mantienen valor usando TT solo orden. | Colisión/origen ilegal no reutiliza score/bound/terminal ni contexto de otra rama. | Contenido TT inspeccionado, igualdad exacta/contextos, con/sin TT. |
| PH05-C03 / F-Q1-17, F-Q4-06 | Easy/normal respetan presupuesto admitido y cancelación libera arena/auxiliares. | Tiempo/nodos/memoria/depth excedidos/tardíos no mueven; arena no se presenta como RSS y overshoot no se omite. | SearchStats, CPU/wall/RSS/arena/TT/contexto, contadores auxiliares y cancelación. |
| PH05-C04 / F-Q1-11/17 | Game revalida snapshot/turno/base y respuesta AI/fallback en ambos órdenes producen máximo un efecto. | Grant consumido/callback incierto/crash/stale/ilegal no inicia doble cálculo ni confirma segundo movimiento; unknown permanece honesto. | Estado PG/grant/effect receipt/barreras y red, fallo real vs doble declarado. |
| PH05-C05 / F-GOV-06 | Delta con aprobación específica y contratos/versiones actualizadas permite revisar código afectado. | Aprobación documental, mensaje de agente o referencia null bloquea editar fuentes/implementar endpoints nuevos. | ContractChange/1 y AuthorizationReference/1 accesibles, diff/review/compatibilidad. |
| PH05-C06 / F-Q2-09 | Undo/redo/restauración reconstruyen estado+contexto/cursor/generación; nueva jugada aceptada elimina futuro. | Respuesta vieja aunque revisión coincida, cancelación/CAS o move rechazado no revive rama ni elimina futuro válido. | Tuplas/estado/almacenamiento/cursor y tráfico antes/después. |
| PH05-C07 / F-Q1-18, F-Q2-09 | Hint vigente ofrece motivo/preview voluntaria; opponent_move válido requiere la intención correspondiente. | Recibir/releer hint genera cero POST move y cero cambios timeline/progreso; ranked oficial rechaza ayuda. | Red/estado/progreso antes/después y permisos por recurso real; no garantía de detectar tablero copiado como escenario synthetic libre. |
| PH05-C08 / F-Q2-09/10/11 | Tutorial/replay/escenarios completos por teclado/lector/tacto HTML/3D ES/EN con progreso versionado. | Storage corrupto/lleno, contenido incompatible o WebGL ausente no bloquea reglamento/funciones ni aplica reglas TS. | AccessibilityReview/1, corpus/objetivos revisados, eventos/cursor y errores. |

09.01/09.04 fijan QA-Evidence/1/AcceptanceDossier/1. Torneos informan W/D/L, muestra, semillas/colores/corpus y referencia, duración/nodos/violaciones e incertidumbre; no certifican Elo o superioridad universal. Perfiles150/700ms, 12k/100knodos, 4/10plies, 16/32MiB son PROPUESTA de 08.01, conservando techos 07.02; no son latencia end-to-end ni presupuesto adicional autorizado. Memoria incluye TT/contexto/auxiliares antes de crecer.

## Seguridad, recursos, DoD y recuperación

Game autoriza recurso/intención/cuota; AI no recibe permiso de mover ni DB. Hint no ejecuta simulación move automáticamente ni cambia objetivo/progreso; acción humana posterior es intención separada. UI ranked carece de ayudas y Game rechaza ayuda ligada a recurso oficial ranked. Un tablero libre sin procedencia no permite prometer antitrampa contra copias o ayuda externa; nueva política requeriría decisión/contrato propios.

Trabajo/test/CI/recursos externos solo con alcance autorizado; límites persistidos compartidos, datos/semillas sintéticos, logs sin tokens/posición privada innecesaria. Revisiones50/75/90 sin upgrade automático. Cancelación por invalidación descarta salida; fallback legal no evade cuota/permiso ni revive claim consumido. Si no puede demostrarse un resultado conservar estado recuperable según contrato.

DoD 09.04: todos los modos y contratos afectados aprobados/versionados, pruebas positivas/negativas observadas, 80% lógica de servicios/90% motor si cambia más críticos, revisión independiente/docs/evidencia. P0/P1/crítico o requerido FAIL/BLOCKED/NOT_RUN impide cierre; P2 con owner/plazo/mitigación y aceptación acotada, sin eximir contrato/ayuda segura. C reúne GateEvaluation/1 y propietario acepta el commit exacto. A coordina operación, C gobierno y Q verifica recuperación 10.04; conservar run fallido→corrección→nuevo runId y no borrar datos/contexto para ocultarlo. [PH-06](phase-06-accounts-elo-statistics.md) necesita su autorización de código/recursos aplicable.
