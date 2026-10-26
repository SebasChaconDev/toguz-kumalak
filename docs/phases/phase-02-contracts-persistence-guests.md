# PH-02 — Contratos, persistencia e invitados

ID: PH-02 · Versión: 0.1 · Fecha: 2026-10-26 · Estado documental: EN_REVISIÓN · Owner: C — Coordinación · Lote: 5.

Plan futuro PROPUESTA; fuentes APROBADO, producto POR_VERIFICAR. El permiso actual del lote 5 sigue limitado a Markdown/publicación documental.

## Ficha Milestone/1

| Campo requerido | Contenido del hito |
|---|---|
| Identificador | PH-02 |
| Objetivo | Exponer motor/estado mediante servicios y contratos verificables, con Postgre SQL durable, invitado auténtico e idempotencia/autorización. |
| Dependencias | PH-01; 01.01–01.03; 03.01–03.04; 05.01–05.04; 07.01/07.02; 09.01/09.03/09.04; 10.01–10.04; 11.03. |
| Responsable | S desarrolla contratos/servicios/identidad/datos; M revisa integración del motor, A entornos y Q pruebas; C coordina, propietario acepta. |
| Entrada | Motor aceptado en su alcance y frontera Game pendiente explícita; fuentes/revisión documental 11.03, permiso de código y recursos locales/CI/DB aplicables. |
| Entregables | Esqueletos hexagonales de tres servicios, contratos generados Go/TS, migraciones/roles, guest/sesión/JWKS/CSRF, S2S y comandos/lecturas Game atómicos. |
| Pruebas positivas/negativas | Casos PH02-C01–C07, matrices F-Q1-03–14 y subcasos normativos de integración requeridos. |
| Riesgos | RISK-05/07/09; commit incierto, locks incompletos, SQL cruzado, guest recuperado por alias, privilegio alternativo o sesión vencida tras espera. |
| Evidencia obligatoria | QA-Evidence/1 de Postgre SQL real, generación/compatibilidad, transacciones/concurrencia y seguridad con artefactos sanitizados. |
| Salida | HTTP/errores/DB/sesión/S2S seguros y observados, 80% lógica de servicios más críticos, sin adelantar OIDC/ranked/contratos L4-P01. |
| Aprobación humana | AcceptanceDossier/1 y GateEvaluation/1 por commit/combinación; permiso de despliegue/configuración externo separado. |
| Corrección/recuperación | Preservar recibo/intención ante unknown; corregir bajo regresión y revisión. Migración/restore sigue 10.03/10.04, sin down destructivo automático. |

## Entradas y alcance técnico

Fuentes de producto: [01.01](../specs/01-infrastructure/01.01-stack-monorepo.md)–[01.03](../specs/01-infrastructure/01.03-cost-capacity-evolution.md), [03.01](../specs/03-game-engine/03.01-normative-rules-and-corpus.md)–[03.04](../specs/03-game-engine/03.04-outcomes-repetition-and-replay.md), [05.01](../specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md), [05.02](../specs/05-backend/05.02-http-openapi-and-errors.md), [05.03](../specs/05-backend/05.03-transactions-data-and-migrations.md), [05.04](../specs/05-backend/05.04-event-schemas-and-postcommit-publication.md), [07.01](../specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md) y [07.02](../specs/07-identity-data/07.02-tokens-authorization-and-service-identity.md). Leer ADR-0001/0003/0004/0005 en el [índice](../index.md#inventario), PRE-05/PRE-07 y PRE-09.A2/B3/B4/C1/C2 de [REG](../decision-register.md).

Consumir [09.01](../specs/09-quality/09.01-engine-data-contract-and-ai-verification.md), [09.03](../specs/09-quality/09.03-ci-security-and-release-controls.md), [09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md) y [10.01](../specs/10-operations/10.01-deployments-origins-regions-and-secrets.md)–[10.04](../specs/10-operations/10.04-releases-incidents-and-portability.md) para evidencia/checks/entornos/recuperación. Gobierno: [11.03](../specs/11-orchestration/11.03-milestones-gates-and-evidence.md), [PLANS](../../PLANS.md#ejecucion-lote-5), [CONF](../conflicts.md).

Esqueleto de AI aquí significa transporte protegido y puerto/contrato probado con doble declarado; búsqueda real se implementa enPH-05. OIDC/cuentas, perfiles completos, Elo/estadísticas y borrado distribuido completo se terminan enPH-06. Definir barreras/interfaces necesarias ahora no declara esas funciones implementadas. Ningún endpoint nuevo de L4-P01 se introduce para resolver una falta de contrato.

## Contratos y componentes futuros

Rutas candidatas de 01.01: services/identity, services/game, services/ai, modules/engine, contracts/http, contracts/events, modules/contracts y packages/contracts; asignar archivos exactos por TaskEnvelope/1 antes de escribir. Identity/Game tienen esquemas/roles/migraciones propios; AI sin DB. Transporte Gin→aplicación→dominio, puertos junto al consumidor y adaptadores explícitos. No DI global ni HTTP/SQL en motor.

Consumir OpenAPI 3.1, JSON Schema y Problem/1 RFC 9457; CommandEnvelope/CommandReceipt/GameView/SimulationRequest/RecoveryView de 05.02; fingerprints/binding de 07.02 y transacciones de 05.03. Versiones/uint64 decimal/JCS deben ser interoperables Go/TS. No confiar en tablero/actor/outcome enviados por navegador. QA-Evidence/1 y AcceptanceDossier/1 se toman de 09.01/09.04; TaskEnvelope/1 y permisos de 11.03.

## Tickets y secuencia

| Ticket / dependencia | Pasos y propósito | Entrega revisable |
|---|---|---|
| PH-02-T01 / entrada | Fijar releases/generadores compatibles y mapa de módulos; contratos primero, pruebas de forma/error/compatibilidad antes de consumidores. | OpenAPI/esquemas y Go/TS generados sin deriva; C01. |
| PH-02-T02 / T01 | Crear composición/transportes mínimos de tres servicios y adaptadores con arranque local; declarar motor compartido de Game/AI sin depender solo de go.work. | Builds por raíz y fronteras de imports; C01/C06. |
| PH-02-T03 / T02 | Preparar Postgre SQL aislado autorizado, propietarios/roles, Goose expand/contract, sqlc/pgxpool y restricciones reales. Migración por proceso controlado. | Inicial/upgrade/rollback de ensayo no destructivo y denegación SQL cruzada; C02. |
| PH-02-T04 / T03 | TDD de guest/session/JWKS/refresh/cookies/Origin/CSRF; probar recuperación de prepare/commit/receipt según 07.01. | Una sesión local controla ambos colores; UUID no alias como autoridad; C03. |
| PH-02-T05 / T04 | Implementar autenticación de recurso/S2S vinculada a caller/aud/scope/op/método/path/query/body y cuotas durables. | Permiso mínimo y falla cerrada; AI sin privilegio de mover/DB; C04. |
| PH-02-T06 / T05+PH01 | Implementar comando oficial: descubrir guardas, jerarquía locks/releer, validar autoridad temporal, receipt antes de stale, motor, escrituras y commit. | Estado/move/contexto/eventos/lifecycle/receipt unidos; C05. |
| PH-02-T07 / T06 | Lecturas autorizadas/no-store, receipt/retry original y SimulationRequest separado; publicación vía puerto postcommit con doble/fallo declarado. | No304 por gameRevision, cero efecto oficial en simulación, unknown preservado; C05/C06. |
| PH-02-T08 / T07 | Ejecutar carreras/reinicio/fallos PG reales y subcasos ACC11/13; medir cobertura/compatibilidad y revisar seguridad. | Dossier/C07 y frontera con PH04/06 explícita. |

T03/T04 preceden comandos; no construir consumidores sobre sesión o transacción imaginadas. Locks siguen 05.03: guardas ordenadas→sala/partida→ratings aplicables, releer y reiniciar si cambia el conjunto; no red dentro de tx. Ratings aún no se aplican en este incremento.

## Pruebas positivas, negativas y evidencia

| Caso / fuente | Positivo esperado | Negativo esperado | Artefacto adicional a QA-Evidence/1 |
|---|---|---|---|
| PH02-C01 / F-Q1-13, F-DEP-01/02 | Contratos generados/validados, builds limpios y errores seguros. | Payload/enum/eventversion/duplicados/JCS/uint64 inválidos o drift bloquean; Identity no importa reglas. | Fuentes+hashes, diff de regeneración, matriz generador/consumidor y logs por raíz. |
| PH02-C02 / F-Q1-08 | Goose/sqlc/restricciones e índices reales compatibles; app usa propiedad autorizada. | CONNECT/SQL cruzado/DDL app, AI con DB, constraint/overflow o upgrade fallido no se toleran. | Catálogo/roles/SQLSTATE, consultas independientes y migraciones antes/después. |
| PH02-C03 / F-Q1-09/10 | Firma/iss/aud/exp/JWKS correctos y refresh perdido recuperan solo autorización original. | none/HMAC/unknownkid/cache vencida, Origin/CSRF erróneos, igualdad de expiración o alias ajeno rechazan. | F-AUT de 07.02, browser/DB sintéticos, contadores fetch y cronología. |
| PH02-C04 / F-Q1-11, F-OBS-01 | S2S/grant/binding/claim y cuota persistida funcionan entre instancias. | Otro cuerpo/handler/entorno/scope, callback incierto o cuota DB caída no abre privilegio ni segundo cálculo. | Hashes Go/TS Unicode/enteros, estado claim/receipt, carreras/barreras. |
| PH02-C05 / F-Q1-03/04/05/07 | Mismo actor/ID/hash devuelve receipt antes de stale; fallo precommit revierte completo. | ID reutilizado distinto→conflicto; commit sin certeza mantiene unknown, sin nuevo ID ni aviso; deadlock/retry acotado completo. | Postgre SQL real varias conexiones, puntos de fallo y SQL/REST antes/después. |
| PH02-C06 / F-Q1-14/13 | Lectura autorizada y simulación Go sin efectos oficiales; commit precede intento de aviso. | Proveedor caído no revierte commit; Web/AI no oficializan, sin goroutine tardía/exactly-once supuesto. | Eventos/receipts y llamada al puerto, doble vs integración real identificados. |
| PH02-C07 / F-Q1-N11/N13 | Predicado de ambos reclamantes y race real de comando/claim preservan revisión/efecto único. | Cierre no suma aparición; retry no duplica lifecycle; terminal/stale/tercero no cambia tablero. | Ledger repetición y history_entries/lifecycle_records: movimiento+cierre en misma tx tienen dos history_ordinal distintos aunque compartan gameRevision. |

Para ejecutar ACC11 con dos sujetos usar harness de integración que establece participantes válidos bajo contrato, sin presentar como completado el flujo UI/sala PH04. Cada fixture indica derivación/alcanzabilidad y límites. No fabricar una historia deportiva para forzar el umbral.

## Seguridad, recursos y recuperación

Separar local/CI/integración/Preview/producción con datos sintéticos; no copiar prod ni iniciar proveedor por este plan. Cuotas compartidas persistidas, pgxpool acotado, conexión directa para migración y no migrar al arrancar Functions. Secretos solo en entorno autorizado, logs sanitizados sin cookies/DSN/JWT/proofs. Coste objetivo0USD y revisiones50/75/90, sin cambio automático a Redis/cola/servicio de pago. Push/CI/Preview requieren comprobar permiso/automatizaciones.

Ante timeout conservar misma intención y resolver receipt/autorización actual; no inferir rollback ni publicar si commit incierto. Revocación/barrera se revalida tras espera de locks con tiempo DB. Restauración futura debe preservar receipts/tombstones/namespace y no resucitar refresh/grants; backup sin ensayo no acredita seguridad. Ante P0 A coordina operación, C escalamiento/gobierno y Q verifica recuperación, sin reparación destructiva improvisada.

## DoD y gate

DoD 09.04: fuentes vigentes, 80% lógica de servicios más críticos, Postgre SQL real/contratos/sesiones/S2S y negativos observados, documentación/compatibilidad y revisión independiente; subcasos de producto posteriores visibles como pendientes. Un mock no reemplaza integración PG. FAIL/BLOCKED/NOT_RUN requerido o P0/P1/crítico sin verificar bloquea salida; P2 requiere mitigación/owner/plazo/aceptación específica.

C consolida AcceptanceDossier/1 y GateEvaluation/1 del commit exacto; propietario acepta el incremento y siguiente acción. Corregir con nuevo runId enlazado y regresión, mantener fallos previos. No tocar 05.02/05.03/07.02 fuera de un cambio autorizado; L4-P01 permanece PROPUESTA. [PH-03](phase-03-accessible-client-and-3d.md) requiere permiso propio de código, sin implicar despliegue.
