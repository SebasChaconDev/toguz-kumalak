# PH-04 — Multijugador privado, reconciliación y recuperación

ID: PH-04 · Versión: 0.1 · Fecha: 2026-10-26 · Estado documental: EN_REVISIÓN · Owner: C — Coordinación · Lote: 5.

Plan futuro PROPUESTA; fuentes APROBADO, producto POR_VERIFICAR. Este incremento entrega salas privadas casuales e invitados. Las clasificatorias/Elo permanecen comprometidas para PH-06 y el lanzamiento completo; no se ofrecen antes ni se reclasifica una partida ya iniciada.

## Ficha Milestone/1

| Campo requerido | Contenido del hito |
|---|---|
| Identificador | PH-04 |
| Objetivo | Jugar entre dos participantes privados con autoridad Postgre SQL y recuperación segura aun si Ably no entrega. |
| Dependencias | PH-03; 05.01–05.04; 06.01–06.03; 07.01/07.02 y contratos 07.03/07.04; 03.03/03.04; 04.01/04.03/04.04; 01.03; 09.01–09.04; 10.01/10.02/10.04; 11.03. |
| Responsable | S implementa salas/permisos/reconciliación; V integra ambas vistas, A recursos/observación y Q verifica; C coordina, propietario acepta. |
| Entrada | Cliente local/contratos/sesión/PG aceptados, permisos de código y recursos Ably/entornos si se ensayan realmente, política de cuotas y fuentes/gate 11.03 verificados. |
| Entregables | Crear/invitar/unir/ready/start, pertenencia/cupo/turnos, postcommit, scheduler REST, reconexión24h, multitab y resultados casuales. |
| Pruebas positivas/negativas | PH04-C01–C07 y matrices06/09; dos contextos reales separados y tercero ajeno, carreras y fallos controlados. |
| Riesgos | Fugas de invitación/estado, doble admisión, confiar en Ably, cronómetro cliente, deadlock/commit incierto y vinculación que cambia sujeto efectivo. |
| Evidencia obligatoria | QA-Evidence/1 con PG/REST/receipts, cronología y barreras, permisos/capacidades/consumo, DOM/foco y snapshots de dos participantes. |
| Salida | Partida casual y cierres correctos, recuperación y fallos observados sin P0/P1/críticos pendientes, sin afirmar exactly-once ni Elo operativo. |
| Aprobación humana | AcceptanceDossier/1 y GateEvaluation/1 del conjunto compatible; recursos/deploy y siguiente fase requieren permisos aplicables independientes. |
| Corrección/recuperación | Conservar intención/receipt y converger por REST; contener funciones afectadas sin alterar reglas/datos; A opera incidente, C gobierna y Q verifica reapertura. |

## Fuentes y contratos consumidos

Leer [06.01](../specs/06-multiplayer/06.01-rooms-invitations-and-ranked-lifecycle.md), [06.02](../specs/06-multiplayer/06.02-ably-notifications-and-reconciliation.md), [06.03](../specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md), [05.01](../specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md)–[05.04](../specs/05-backend/05.04-event-schemas-and-postcommit-publication.md) y [07.01](../specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md)/[07.02](../specs/07-identity-data/07.02-tokens-authorization-and-service-identity.md). [07.03](../specs/07-identity-data/07.03-profiles-history-and-data-rights.md)/[07.04](../specs/07-identity-data/07.04-elo-leaderboard-and-statistics.md) se consumen para no romper barreras/elegibilidad aunque sus funciones completas llegan después.

Reglas/representación: [03.03](../specs/03-game-engine/03.03-moves-and-semantic-events.md)/[03.04](../specs/03-game-engine/03.04-outcomes-repetition-and-replay.md), [04.01](../specs/04-visual-experience/04.01-navigation-and-accessibility.md)/[04.03](../specs/04-visual-experience/04.03-interaction-animation-and-audio.md)/[04.04](../specs/04-visual-experience/04.04-performance-and-html-fallback.md). Calidad/operación: [09.01](../specs/09-quality/09.01-engine-data-contract-and-ai-verification.md)–[09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md), [01.03](../specs/01-infrastructure/01.03-cost-capacity-evolution.md), [10.01](../specs/10-operations/10.01-deployments-origins-regions-and-secrets.md)/[10.02](../specs/10-operations/10.02-observability-quotas-and-rate-limits.md)/[10.04](../specs/10-operations/10.04-releases-incidents-and-portability.md) y [11.03](../specs/11-orchestration/11.03-milestones-gates-and-evidence.md). Consultar PRE-06/PRE-09.C3/D2, [REG](../decision-register.md), [CONF](../conflicts.md) y [PLANS](../../PLANS.md#ejecucion-lote-5).

Contratos: RoomView/GameView/ReadyConsent/RecoveryView/RematchView, CommandEnvelope/Receipt, NoticeV1/SemanticEventV1, capacidades temporales 07.02 y metadatos de actividad 06.03. Cualquier parámetro digital nuevo conserva PROPUESTA. El reloj DB decide vigencia; GET/capacidad/aviso no prolongan actividad. Misma gameRevision puede traer metadatos nuevos, por eso no-store/sin304 basado solo en esa revisión. Revancha es agregado nuevo, sin reabrir ni cambiar elegibilidad del anterior.

## Componentes futuros y tickets

Rutas candidatas: services/game para casos de uso/adaptadores PG/Ably, services/identity para capacidades de sesión existentes, apps/web para controlador/sala/scheduler compartido y contratos generados. Asignar archivos por TaskEnvelope/1, un escritor activo; ninguna fuente aprobada cambia por comodidad del consumidor.

| Ticket / depende de | Pasos y motivo | Entrega pequeña |
|---|---|---|
| PH-04-T01 / entrada | Crear/leer sala privada casual, invitación/código alternativo, expiración/cupo2 y permisos sin autojoin al montar UI. | Admisión y denegación bajo lock; C01. |
| PH-04-T02 / T01 | Join/ready/start con consentimiento bilateral, horizonte de sesión y colores CSPRNG persistidos; revalidar guardas/miembros. | Inicio único recuperable y modalidad/elegibilidad fija; C01/C02. |
| PH-04-T03 / T02 | Conectar turnos/move/claim/renuncia/oferta-aceptación y estados accesibles HTML/3D. | Dos participantes completan partida; claim cualquiera de ambos, no turno obligatorio; C02/C03. |
| PH-04-T04 / T03 | Publicador Game postcommit y capacidades de suscripción mínimas por partida; pérdida/error/timeout separado del receipt. | Un aviso lógico puede entregarse0/1/N; no worker/outbox inventado; C04. |
| PH-04-T05 / T04 | Scheduler 06.02: single-flight/coalescer mayor revisión, epochs, snapshot inicial/reconexión/foco y sondeo acotado compartido. | Convergencia a REST sin polling por panel; C04/C05. |
| PH-04-T06 / T05 | Actividad/resume y ventana24h; multitab/unknown conservan intención/sujeto; probar antes/en/después de deadline y espera de lock. | Estado no se confirma por timer local ni GET; C05/C06. |
| PH-04-T07 / T06 | Terminal/revancha consentida en agregado nuevo y archivo inconcluso persistido sin Elo; distinguir metadata rematch de revisión deportiva. | Cierre una vez y lectura terminal actual no-store; C03/C06. |
| PH-04-T08 / T07 | E2E tres motores/ES/EN/HTML/3D, fallos frontera y consumo por pestaña/entorno; revisión independiente del dossier. | C07, compatibilidad y DoD 09.04 observados. |

Implementar solo rama casual/invitados de 06.01 en este incremento; endpoints/control UI ranked no anuncian disponibilidad y servidor rechaza creación/admisión ranked hasta PH06 habilitado bajo contrato aprobado. No cambiar 06.01 ni retirar ranked del alcance final.

## Pruebas y evidencias

| Caso / fuentes QA | Positivo esperado | Negativo esperado | Artefacto adicional |
|---|---|---|---|
| PH04-C01 / F-Q2-03, F-Q1-07 | A crea/invita, B entra/ready/start en dos contextos con cookies distintas. | C ajeno/código inválido/expirado/cupo lleno o dos joins no filtran ni añaden tercero; no autojoin por render. | Red/PG/roles/locks, ventanas y estado antes/después. |
| PH04-C02 / F-Q1-03/04/05, F-Q2-03 | Move confirmado/retry con mismo ID tiene único efecto y ambos ven snapshot actual. | Carrera/stale/otro actor/body, pérdida de respuesta/commit incierto no genera nuevo ID o rollback supuesto. | Receipt/history/contexto/SQL independiente y barreras. |
| PH04-C03 / F-Q1-N11/N13, F-Q2-04/08 | Claim válido de cualquiera, acuerdo bilateral/renuncia y terminal producen resultado/causa correctos. | Oferta unilateral, tercero, claim tardío/terminal o movimiento concurrente rechazan; cierre no suma repetición. | Ambas órdenes de race, lifecycle una vez; movimiento+cierre dos history_ordinal distintos en misma tx/gameRevision, retry sin duplicado. |
| PH04-C04 / F-Q1-14, F-Q2-06 | Ably caído/último aviso perdido conserva commit y REST converge. | Precommit no publica; duplicado/desorden/gap/versión desconocida no retrocede/aplica delta parcial ni promete exactly-once. | Cronología commit/publicación/REST, eventos/IDs y capacidades. |
| PH04-C05 / F-Q2-06/07, F-Q4-07 | Dos pestañas y vuelta de A tras movimiento deB mantienen snapshot/metadata/foco/intención. | Montar panel/cambiar HTML/3D no duplica read/heartbeat/POST; aviso no prueba deadline. | Contadores scheduler/red, epochs y DOM/foco. |
| PH04-C06 / F-Q1-10/14, F-Q2-06 | Sesión/horizonte válidos recuperan antes del deadline; vencimiento DB archiva inconclusa una vez. | Igualdad/postdeadline/lock tardío/refresh perdido/link G→U no revive permiso ni cambia Elo; desconexión no renuncia automática. | Reloj DB/barreras/sesiones/receipts y cronología. |
| PH04-C07 / F-Q2-10/11, F-OBS-01/07 | Escenarios ES/EN/teclado/lector tres motores y cuotas durables por ámbito permiten revisar consumo. | Invitado/casual/inconclusa/revancha previa sin ranked nunca cambia ratings; cuota desconocida/50/75/90 no activa pago o degradación nueva. | Review acceso, ratings antes/después, consumo/decisión humana. |

QA-Evidence/1 identifica casos/AC/commit/build/motor/contratos/herramientas/DB/navegador/entorno/semillas y diferencia doble, fallo inyectado y proveedor real. Ensayar Ably real requiere permiso externo; un doble no acredita SDK/capacidad/región. Medir retorno→snapshot y→acciones con errores/censura/frío-caliente 09.04, sin SLO inventado.

## Seguridad, recursos y gate

Tokens temporales solo canal/recurso autorizado y suscripción mínima, navegador sin publish oficial/presence/history indebidos; no claves en bundle. Invitación no es identidad ni permiso permanente. Ocultar códigos, cookies y payload privado en artefactos. Compartir presupuesto entre pestañas/entornos/cuentas; contar publicación y cada entrega, 50/75/90 revisan sin pago/escala automática. No convertir Redis/Socket. IO/cola en requisito.

DoD 09.04 y compatibilidad 10.04: flujos/negativos/críticos observados, cobertura aplicable80% servicios más rutas críticas, revisión independiente/docs y artefactos accesibles. P0/P1 o requerido FAIL/BLOCKED/NOT_RUN impide salida. P2 diferido requiere mitigación/owner/plazo/aceptación específica. C prepara GateEvaluation/1 y propietario acepta el alcance exacto; aceptación no crea proveedor/despliega.

Corregir sobre misma autoridad e IDs, conservar intentos y ejecutar regresión afectada. A coordina operación del incidente, C gobierno/escalamiento y Q verifica recuperación antes de reapertura 10.04; código revertido no restaura datos. [PH-05](phase-05-ai-practice-tutorial.md) requiere autorización de implementación propia y L4-P01 separado para sus tickets afectados.
