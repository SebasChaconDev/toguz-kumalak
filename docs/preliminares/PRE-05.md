# PRE-05 — Backend, contratos y persistencia

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: PRE-05 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner: S — servicios.

## Objetivo y contexto

Delimitar servicios, contratos y consistencia a partir de 24 decisiones aprobadas. El diseño necesita conservar autonomía e integridad en Functions, sin atribuir a notificaciones o clientes autoridad sobre una partida. Este preliminar no acredita implementación ni una prueba de factibilidad realizada.

## Fuentes y decisiones

El [maestro archivado](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) y el [registro canónico](../decision-register.md) preservan las fuentes. Las implicaciones siguientes tienen naturaleza `APROBADO`:

- [PRE-05.A1](../decision-register.md#pre-05-a1): el monolito queda sustituido por Identity, Game y AI Go/Gin, autónomos, más frontend Next.js.
- [PRE-05.A2](../decision-register.md#pre-05-a2): transporte, aplicación, dominio y adaptadores se separan dentro de cada servicio.
- [PRE-05.A3](../decision-register.md#pre-05-a3): constructores explícitos e interfaces pequeñas junto al consumidor controlarán dependencias.
- [PRE-05.A4](../decision-register.md#pre-05-a4): cada servicio tendrá módulo autónomo y uno de cuatro despliegues Vercel del monorepo.
- [PRE-05.B1](../decision-register.md#pre-05-b1): REST se versionará en /api/v1 con recursos consultables y comandos explícitos.
- [PRE-05.B2](../decision-register.md#pre-05-b2): OpenAPI 3.1 gobierna HTTP y generación Go/TypeScript con versiones fijadas.
- [PRE-05.B3](../decision-register.md#pre-05-b3): Problem Details RFC 9457 incorporará código de dominio, requestId y mensajes seguros.
- [PRE-05.B4](../decision-register.md#pre-05-b4): games, players, practice y leaderboard exponen dominio/paginación, no tablas SQL.
- [PRE-05.C1](../decision-register.md#pre-05-c1): la jugada pasa por autenticación, autorización, versión, motor, persistencia y respuesta autoritativa.
- [PRE-05.C2](../decision-register.md#pre-05-c2): bloqueo por partida, versión e identificador idempotente protegerán concurrencia.
- [PRE-05.C3](../decision-register.md#pre-05-c3): comando, movimiento, efectos y estado se confirman juntos en transacción.
- [PRE-05.C4](../decision-register.md#pre-05-c4): local/práctica usa endpoints sin autoridad competitiva y conserva un único motor Go.
- [PRE-05.D1](../decision-register.md#pre-05-d1): Identity DB y Game DB tienen propiedad separada; AI es stateless y no hay SQL cruzado.
- [PRE-05.D2](../decision-register.md#pre-05-d2): estado JSONB compacto/versionado y validado en Go convive con movimientos y consultas relacionales.
- [PRE-05.D3](../decision-register.md#pre-05-d3): conexiones de aplicación/migración se distinguen y pgxpool se limita para serverless.
- [PRE-05.D4](../decision-register.md#pre-05-d4): Goose ejecutará migraciones SQL controladas, nunca al iniciar automáticamente Functions.
- [PRE-05.E1](../decision-register.md#pre-05-e1): los avisos Ably llevan gameId, versión, tipo y mínimos datos, sin instrucciones 3D.
- [PRE-05.E2](../decision-register.md#pre-05-e2): PostgreSQL confirma antes de publicar; REST reconcilia y el aviso perdido sin outbox queda documentado.
- [PRE-05.E3](../decision-register.md#pre-05-e3): tokens Ably temporales autorizan suscripción limitada; Game publica los eventos oficiales.
- [PRE-05.E4](../decision-register.md#pre-05-e4): eventos usarán esquemas JSON versionados y validación Go/TypeScript.
- [PRE-05.F1](../decision-register.md#pre-05-f1): OTel correlacionará trazas HTTP, servicio, SQL y publicación Ably.
- [PRE-05.F2](../decision-register.md#pre-05-f2): verificación futura cubrirá PostgreSQL real, Gin/HTTP, contratos, carreras y doble Ably.
- [PRE-05.F3](../decision-register.md#pre-05-f3): validación, autorización y límites server-side protegerán salas, invitados e IA.
- [PRE-05.F4](../decision-register.md#pre-05-f4): la prueba Go/Gin/Vercel/Neon/Ably/transacción precederá a implementación completa; sigue sin ejecutarse.

## Agrupaciones de diseño y consecuencias

**Autonomía y contratos.** Compartir motor Go no comparte transportes, DB o autoridad. Identity posee cuentas/sesiones; Game, partidas/resultados/Elo/agregados; AI calcula propuestas sin DB inicial ni confirmación oficial. La arquitectura hexagonal es pragmática y sus interfaces deben responder al dominio, sin añadir un servicio nuevo para cada adaptador.

**Unidad de consistencia.** Cada comando requiere una identidad autorizada, versión esperada y clave idempotente. El contrato deberá precisar reintentos, conflicto de versión y reutilización de clave con contenido distinto. La transacción evita aceptar dos jugadas incompatibles; los movimientos no se aceptan a partir de estados completos del navegador. El cierre competitivo exige aplicar resultado y estadísticas una única vez en Game.

**Persistencia y avisos.** JSONB guarda estado validado, mientras columnas/filas normalizadas soportan consultas y restricciones. Publicar después del commit evita anunciar una jugada inexistente, pero abre una ventana de aviso perdido. No se supone entrega exactamente una vez, outbox duradero ni transacción ACID entre PostgreSQL y Ably. Reconexión, foco y verificaciones acotadas deben recuperar el estado oficial por REST.

## Sustituciones y matices

A1 reemplaza expresamente el monolito. [PRE-01](PRE-01.md) recoge los reemplazos de Drizzle y Socket.IO; no reintroducirlos como dependencias implícitas. Un endpoint de simulación local/práctica no puede modificar Elo ni confirmar un resultado competitivo. Las credenciales Ably no sustituyen sesión propia ni permisos por recurso.

## Dependencias y fuentes canónicas

El [dominio 05 del índice](../index.md#spec-05-01) es `PENDIENTE_LOTE_3`, tras [PRE-03](PRE-03.md) y [PRE-01](PRE-01.md). Los límites de identidad, salas y operación se consumen de [PRE-07](PRE-07.md), [PRE-06](PRE-06.md) y [PRE-10](PRE-10.md). Regir redacción por [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [trazabilidad](../traceability.md).

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: compatibilidad serverless, módulos Go compartidos, pools Neon, límites de conexiones, publicación posterior al commit y reconciliación sin outbox. [Conflictos](../conflicts.md) conserva impactos y aprobaciones necesarias. No están autorizados endpoints, SQL ejecutable, migraciones, conexiones a proveedores o pruebas de integración en este lote.

## Criterios futuros de aceptación y verificación

- **AC-01:** contratos documentan recursos, errores, identidad, versión, idempotencia, eventos y compatibilidad sin exponer tablas ni secretos.
- **AC-02:** escenarios de concurrencia/reintento/fallo antes y después de commit describen exactamente un efecto oficial y recuperación del estado aceptado.
- **AC-03:** propiedad de datos e interfaces demuestra ausencia de SQL cruzado y de autoridad de AI sobre resultados.
- **AC-04:** la prueba futura de plataforma tiene precondiciones, casos negativos, evidencia exigida y gate antes de implementar el conjunto.

Ninguno de estos criterios equivale a una integración probada durante el trabajo documental.

## Revisión humana

Revisar fronteras, ventana de notificación perdida y contratos de idempotencia. El documento permanece `EN_REVISIÓN`; las decisiones conservan autoridad independiente de este borrador. El coordinador pedirá «APRUEBO LOTE 1» y no iniciará otro lote sin instrucción.
