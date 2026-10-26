# PH-07 — Hardening, recuperación, portafolio y preparación de lanzamiento

ID: PH-07 · Versión: 0.1 · Fecha: 2026-10-26 · Estado documental: EN_REVISIÓN · Owner: C — Coordinación · Lote: 5.

Plan futuro PROPUESTA; fuentes APROBADO, resultados POR_VERIFICAR. No hay demo, release o fecha de lanzamiento acreditadas por redactarlo. Conserva MVP base y ampliado completos.

## Ficha Milestone/1

| Campo requerido | Contenido del hito |
|---|---|
| Identificador | PH-07 |
| Objetivo | Demostrar DoD integral, seguridad/acceso/rendimiento/recuperación y preparar publicación profesional y demo condicionadas a aprobación. |
| Dependencias | PH-00–PH-06, todas las specs01–08 aplicables, 09.01–09.04, 10.01–10.04 y 11.01–11.03; licencia/procedencia/cultura/política pendientes resueltas antes de acciones afectadas. |
| Responsable | Q coordina verificación; A operación/releases/recuperación, M/S/I/V corrigen dominios, C gobierno/dossier y propietario aceptación/publicación. |
| Entrada | Incrementos completos con evidencia/revisiones, revisión documental 11.03, permisos de hardening/código/ensayos y alcance externo aplicable; combinación candidata congelada. |
| Entregables | Dossier integral, matrices de compatibilidad/CI/seguridad/acceso, benchmarks, restore ensayado, runbooks, procedencia/licencias y material de portafolio/demo preparado. |
| Pruebas positivas/negativas | PH07-C01–C09 y suites09/10 requeridas del conjunto exacto, sin omitir críticos o convertir unavailable en PASS. |
| Riesgos | Seguridad/corrupción ocultas por checks verdes, recursos/coste, restore incompleto, rollback incompatible, licencia/cultura/política o demo publicada sin permiso. |
| Evidencia obligatoria | QA-Evidence/1/AcceptanceDossier/1 completos, manifest compatible, mediciones crudas, ensayo restore y aprobaciones de recursos/contenido/release. |
| Salida | Cero P0/P1/críticos sin validar, DoD integral y gate humano del candidato; release/demo solo si permiso específico la cubre y smoke/observación demuestran estado. |
| Aprobación humana | Tres permisos independientes documental/código/externo, más aceptación concreta de candidato/publicación/merge cuando corresponda; no inferencia por compilar o cerrar lote 5. |
| Corrección/recuperación | Contener/revertir código compatible o forward-fix autorizado, tratar datos/seguridad con 10.03, preservar evidencia y validar reapertura antes de afirmarla. |

## Fuentes y mapa de consumo

Leer [inventario40specs](../index.md#inventario) y fuentes directas de cada prueba; no usar un resumen de agentes como sustituto. [REG](../decision-register.md), [CONF](../conflicts.md), [SRC](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md), [AGENTS](../../AGENTS.md) y [PLANS](../../PLANS.md#ejecucion-lote-5) fijan autoridad. PRE-09.F1–F4, PRE-10.C2/C3/F1–F4, PRE-01.D4 y PRE-04.F2 son especialmente aplicables.

| Grupo de fuentes | Consumo del cierre |
|---|---|
| [01.01](../specs/01-infrastructure/01.01-stack-monorepo.md)–[01.03](../specs/01-infrastructure/01.03-cost-capacity-evolution.md) | Módulos/builds/recursos/capacidad y coste sin garantía por spike. |
| [02.01](../specs/02-product/02.01-scope-and-modes.md)–[02.03](../specs/02-product/02.03-accounts-history-and-licensing.md) | Base/ampliado, ES/EN, cuentas/derechos/licencia/portafolio. |
| [03.01](../specs/03-game-engine/03.01-normative-rules-and-corpus.md)–[03.04](../specs/03-game-engine/03.04-outcomes-repetition-and-replay.md) | Perfil14 ACC, historia/replay/terminal/contexto/pureza. |
| [04.01](../specs/04-visual-experience/04.01-navigation-and-accessibility.md)–[04.04](../specs/04-visual-experience/04.04-performance-and-html-fallback.md) | Equivalencia HTML/3D, acceso, visual/recursos/cultura/rendimiento. |
| [05.01](../specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md)–[05.04](../specs/05-backend/05.04-event-schemas-and-postcommit-publication.md) | Fronteras/contratos/transacciones/recibos/postcommit. |
| [06.01](../specs/06-multiplayer/06.01-rooms-invitations-and-ranked-lifecycle.md)–[06.03](../specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md) | Salas/ranked/reconciliación/24h/multitab y errores. |
| [07.01](../specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md)–[07.04](../specs/07-identity-data/07.04-elo-leaderboard-and-statistics.md) | Sesiones/trust/derechos/Elo/concurrencia/estadísticas. |
| [08.01](../specs/08-ai-practice/08.01-search-heuristics-and-compute-limits.md)–[08.04](../specs/08-ai-practice/08.04-tutorial-hints-and-educational-replay.md) | Búsqueda/recursos/fallback/ramas/tutorial/hints y delta autorizado. |
| [09.01](../specs/09-quality/09.01-engine-data-contract-and-ai-verification.md)–[09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md) | Matrices por riesgo, CI, 14 ACC, seguridad, DoD y mediciones. |
| [10.01](../specs/10-operations/10.01-deployments-origins-regions-and-secrets.md)–[10.04](../specs/10-operations/10.04-releases-incidents-and-portability.md) | Entornos/secretos/cuotas/backup-restore/releases/runbooks/portabilidad. |
| [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md)–[11.03](../specs/11-orchestration/11.03-milestones-gates-and-evidence.md) | Gobierno, cambios/aprobaciones, propiedad, dossier y continuación. |

Consumir QA-Evidence/1, MeasurementPlan/Result/1, AcceptanceDossier/1 y RiskAcceptance/1 sin redefinir criterios; ReleaseManifest/1, CompatibilityMatrix/1, IncidentRecord/1 y RecoveryDecision/1 de 10.04; protocolo de restauración 10.03; GateEvaluation/1/AuthorizationReference/1/ContractChange/1 de 11.03. Las propuestas nuevas no se convierten enAPROBADO por incorporarlas a este plan.

## Componentes y tickets futuros

Se revisa el conjunto real de apps/web, services/identity/game/ai, modules/engine y contratos según 01.01; tooling/documentación de implementación solo dentro de rutas/permiso futuro. Cada ticket tiene un escritor por archivo/revisor independiente; coordinación integra evidencias y docs.

| Ticket / depende de | Pasos y propósito | Entrega concreta |
|---|---|---|
| PH-07-T01 / entrada | Congelar cuatro commits/builds/config/motor/contratos/migraciones/keyring y matriz de coexistencia actual/candidata/anterior. | ReleaseManifest/CompatibilityMatrix y lista de casos requeridos, C01. |
| PH-07-T02 / T01 | Ejecutar suites completas 09.01/02 y controles 09.03 por PR/main/pre-release segúntrigger, consumidores Game/AI y PG real. | C01/C02; no omisiones críticas para ahorrar tiempo. |
| PH-07-T03 / T02 | Verificar secreto/dependencia/estático/govulncheck, permisosActions/PR externos y auth/recurso/trust/CSRF/JWKS/quotas. | C02/C03 y revisión de amenazas/contingencias. |
| PH-07-T04 / T02 | Completar revisión manual HTML/3D/ES/EN/teclado/lector/móvil y visualGPU con fixtures/control de baseline. | C04 y lista de hallazgos trazada. |
| PH-07-T05 / T02 | Medir HTTP/SQL/reconexión/IA/WebVitals/frame times/descarga/CPU/memoria/cuotas con 09.04 y 10.02. | C05, crudos/errores/censura y naturaleza de umbrales. |
| PH-07-T06 / T01+permiso externo | Backup cifrado identificado→restore aislado→integridad/seguridad/revocación/efectos→decisión de reapertura; ensayar compatibilidad migración/rollback. | C06/C07, tiemposRPO/RTO observados o límite explícito. |
| PH-07-T07 / T03+T06 | Ensayar runbooks 10.04 para diez incidentes, roles/contención/recuperación/postmortem y portabilidad si está autorizada. | C07/C08, sin crear proveedor/infra alternativa automáticamente. |
| PH-07-T08 / T04 | Resolver licencia elegida por propietario, fichas autor/origen/licencia/permisos/atribución y revisión cultural; prepararREADME/arquitectura/limitaciones/demo profesional. | C09; material local revisable y datos sintéticos, sin afirmar demo ya publicada. |
| PH-07-T09 / T03–T08 | Corregir P0/P1/criticos, revisarP2 con owner/plazo/mitigación/aceptación, repetir controles afectados y reunir dossier independiente. | GateEvaluation/1 ready_for_owner o blocked con razón. |
| PH-07-T10 / T09+permiso publicación | Publicar/desplegar solo candidato/entorno aprobados, verificar target/protección/automatizaciones, smoke y observar; detener/contener si falla. | Release/demo observada solo tras evidencia; sin permiso, entrega queda preparada y acción externa pendiente. |

T03/T04/T05 pueden repartir archivos/mediciones independientes con presupuesto común trasT02. T10 no es consecuencia automática deT09, de código autorizado o de repositorio público. No fecha de lanzamiento inventada.

## Casos y evidencias futuras

| Caso / fuente | Positivo esperado | Negativo esperado | Evidencia adicional |
|---|---|---|---|
| PH07-C01 / F-REL-01, F-Q1-13 | Matriz compatible y todos los consumidores/checks requeridos ejercitados en builds exactos. | Motor compartido cambia y faltaGame/AI, API/event/schema/keyring incompatible o SHA distinto bloquea. | Manifests/digests/regeneración/logs y smoke por combinación. |
| PH07-C02 / F-Q4-02/09, 09.03 | PR/main/pre-release/despliegue autorizado ejecutan conjuntos definidos con resultados observados. | SKIPPED/UNAVAILABLE/cancelled/NOT_RUN requerido o artefacto inaccesible no equivale a PASS; mensaje/compile no cierra. | CI trigger/matriz/check/runId/artefacto/review y normalización 09.01. |
| PH07-C03 / F-Q1-09–12, F-DEP-05/06/07 | Trust/recurso/rotación/entorno y mínimo privilegio con datos sintéticos. | PR externo no recibe secretos; credencial enbundle/log, acceso ajeno oPreview desprotegida impide promover. | Escaneos acotados, respuestas negativas, permisos y evidencias sanitizadas. |
| PH07-C04 / F-Q2-07–12 | Tareas de todos los modos completas HTML/3D/ES/EN/teclado/lector/reducedMotion. | WebGL perdido enunknown, contraste/foco/lectura esencial ausente o baseline autoaceptada bloquea. | AccessibilityReview, recorridos manuales, DOM/foco/red y visualdiff por entorno. |
| PH07-C05 / F-Q4-04–07, F-OBS-04/07 | Plan reproducible obtiene muestras/errores/censura/colas/consumo por ámbito y comparación honesta. | Muestra única, idleGPU, promedioFPS engañoso o cuota desconocida no acreditanSLO/60/30/coste 0; 50/75/90 no pagan. | Plan/crudos/hash/método/hardware/región/versiones y límites instrumentales. |
| PH07-C06 / F-Q4-08, 10.03 | Restore aislado prueba162/reglas/historia/contexto/replay/outcomes/ratings/ledger/receipts/roles y reconciliación de revocación/supresión. | Backup presente/hash correcto sin ensayo no es restaurable; refresh/grant/PII/ID consumido no resucita; duda bloquea reapertura. | Backup identificado/corte/SQL/replay/seguridad/RecoveryDecision yRPO/RTO observado, con pérdidas porregistro. |
| PH07-C07 / F-REL-03/05 | Rollback código compatible o forward-fix autorizado conserva datos actuales y recupera servicio seguro. | Down destructivo, restaurar clave revocada/receipts perdidos ohealth200 sin integridad no reabre. | Matriz antes/después, migración/backup validado, smoke y decisión humana. |
| PH07-C08 / F-INC-01–10, F-REL-04 | Diez incidentes tienen detección/severidad/owner/contención/recuperación/postmortem probados; portabilidad conserva contratos si se autoriza. | Ably caído no deshace commit, falloIdentity no omite firma, IA no duplica efecto, PG caído no autorizaRAM/pago/infra nueva. | IncidentRecord/runbook/cronología/evidencia; trialportable separado con permisos/retirada. |
| PH07-C09 / F-GOV-02/10, F-REL-02 | Licencia/procedencia/cultura/política revisadas y tres permisos válidos permiten solo publicación exacta autorizada. | Demo inexistente no se anuncia; permiso documental/código no crea despliegue/pago; recurso sin derechos no se publica. | Decisiones humanas/fichas/contenido sanitizado/permiso/target/smoke observado, oacciónpendiente. |

Evidencia por QA-Evidence/1 y dossier por 09.04; no sumar PASS de otros commits como si fueran el candidato. El hash permite contrastar artefacto, no sustituye oráculo/observación. Rerun conserva fallo y enlaza nuevo intento. Cobertura90/80 más críticos y pruebas positivas/negativas/fallos; no aceptar crítico sin medir bajo P2.

## Seguridad, recursos, DoD y gate

Presupuestar CI/artefactos/Preview/DB/Ably/telemetría/backup por cuenta/entorno/ventana, aislamiento y retirada. Ningún plan gratuito se presume vigente o suficiente; verificar capacidades bajo permiso y detener antes de pago/upgrade. Telemetría minimizada/cardinalidad acotada sin PII/tokens/DSN; no prometer guardia 24/7. RISK-10/13 se resuelve para acciones afectadas, sin licencia o plazo jurídico inventados.

DoD 09.04 completo: requisitos/AC/casos observados del conjunto exacto, críticos validados, cero P0/P1, P2 diferidos solo con aceptación explícita/owner/plazo/mitigación, revisión independiente, docs/compatibilidad/operación/licencias/privacidad actualizadas y evidencia accesible. C reúne GateEvaluation/1 y propietario decide aceptación; implementación y recursos/publicación son permisos separados, sin pedir de nuevo pasos ya cubiertos inequívocamente.

Ante fallo A coordina operación/contención y cronología, C gobierno/escalamiento/dossier, dueños M/S/I/V corrigen y Q verifica recuperación. Preservar intentos/datos/efectos inciertos; rollback de código no revierte datos. El candidato permanece bloqueado hasta nueva evidencia/revisión; no escribir un anuncio optimista. Una demo externa requiere revisión documental/aceptación, permiso de código y permiso externo aplicables; si falta uno, entregar material preparado y detener solo esa acción. Este lote 5 no implementa PH07, no publica demo ni inicia lote 6.
