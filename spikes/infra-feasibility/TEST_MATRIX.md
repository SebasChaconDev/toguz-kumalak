# Matriz oficial POC-01–POC-11 — verificación integral


> **Estado vigente — 2026-10-25 (Caracas): GO_REMOTE, limitado al spike desechable.** Seis builds remotos READY, ocho intentos acumulados contando dos cancelados históricos; pruebas reales Neon/Ably/OIDC/cookie y seguimiento C6 completados. Todos los recursos propios retirados y recursos anteriores preservados. [Cierre, evidencias y límites](REMOTE_AUDIT.md#remote-close-20261025). Las notas anteriores de bloqueo que siguen son cortes históricos, sustituidos por esta adenda. No acredita producto, aprobación automática de propuestas ni generación del lote 4.

> **Cierre actual: BLOCKED_REMOTE por destino Vercel, recursos propios retirados.** Cuotas y protección gratuita verificadas; techo100 retirado por el propietario. Dos builds fueron cancelados porque Vercel devolvió `production` incluso solicitando staging; cero secretos y alias sin asignar. Los cuatro proyectos y ambos deployments se retiraron: inventarios finales vacíos.41 API directas más navegación UI; sin lote4. [Evidencia y decisión pendiente](REMOTE_AUDIT.md#12-prueba-real-de-protección-y-destino-inicial-bloqueo-específico-recursos-retirados). Notas anteriores conservadas como historia.

> **Estado vigente: cuotas verificadas; techo de100 solicitudes retirado por el propietario.** Se conservan18 lecturas API directas y errores históricos; las lecturas UI posteriores no equivalen a un total HTTP exacto. Sigue el spike remoto previamente autorizado con0USD y máximo6builds. Aún no hay PASS remoto ni permiso de lote4. [Auditoría actual](REMOTE_AUDIT.md#11-decisión-posterior-y-comprobación-de-cuotas--2026-10-22). Las notas siguientes son cortes históricos.

> **Estado vigente: Control API Ably autenticada; POC-01–POC-11 siguen BLOCKED/NOT_RUN.** Los HTTP 200 de cuenta, inventario y estadísticas son prechecks; no validan Pub/Sub, cuotas completas ni una prueba funcional. Se contabilizan **16/100**, incluidas la lectura de estadísticas no procesada y su repetición explícita. Reserva diez, disponibles 74; cero builds o escrituras externas. [Registro](evidence/remote-preconditions-budget-20261022.json). Las notas de PAT ausente siguientes son históricas.

> **Última actualización: `BLOCKED_REMOTE`, 12/100 solicitudes operativas.** GET de estadísticas Ably devolvió **401 / 40160** con la clave de aplicación aportada; es un precheck de permisos rechazado, no un PASS/FAIL funcional POC. Las once filas remotas siguen BLOCKED/NOT_RUN. Cero builds o escrituras externas; reserva 10, disponibles 78. [Evidencia](evidence/remote-preconditions-budget-20261022.json). Los cortes anteriores se conservan como historia.

> **Continuación vigente: `BLOCKED_REMOTE`; POC-01–POC-11 siguen BLOCKED/NOT_RUN.** La sesión renovada Neon ya pasó; Vercel confirmó Hobby. Hay **11 solicitudes operativas acumuladas**, cero SQL, cero builds y cero escrituras externas. La definición de contador está resuelta por el propietario; el plan completo de 100 aún no se valida. [Evidencia actual](evidence/remote-preconditions-budget-20261022.json), [revisión presupuestaria](REMOTE_PLAN.md#8-continuación-tras-renovar-neon-unidad-y-presupuesto-revisados). Se conservan los estados anteriores como historia.

> **Estado remoto vigente: `BLOCKED_REMOTE`, 2026-10-22, 20:27 UTC.** El spike ya fue autorizado. La sesión CLI de Neon devolvió 401 en la precondición inicial y se detuvo el intento antes de crear recursos. Todas las filas remotas conservan **BLOCKED**; su ejecución en proveedor es **NOT_RUN**. [Registro y contabilidad](evidence/remote-execution-20261022.json). Los PASS locales no se convierten en PASS remotos.

Actualización: **2026-10-22, America/Caracas**. Conclusión: **GO_LOCAL**, limitada al spike. Base inicial `cf5d3973adc977f0ec09c738cbe6cfc078883593`; hashes y commit de entrega identifican las fuentes corregidas. [Informe](INTEGRAL_REPORT.md), [registro observado](evidence/integral-local-gate.json). Sustituye estados locales previos; historia conservada en Git/evidencias.

Entorno: Windows 11 amd64, PowerShell7.6.5; Go1.26.2 Windows/Linux; Gin1.12.0/pgx5.11.0; Docker29.9.2/Compose5.6.1/PostgreSQL18.6; Node24.13.0/npm11.6.2/Next16.4.2/React19.3.1/TS5.9.3. [Detalle](evidence/integral-environment.json). REMOTO BLOCKED significa gate/evidencia de proveedor pendientes, sin FAIL de arquitectura demostrado.

| ID / objetivo | LOCAL | REMOTO | Procedimiento y resultado observado | Evidencia y límites |
|---|---|---|---|---|
| POC-01 — Cuatro builds independientes y motor compartido | **PASS** | **BLOCKED; 2 intentos CANCELED** | Tres contextos Go/cachés nuevos, servidor/API, GOWORK=off; Next instalación nueva/lock intacto. Game/AI mismo engine versión/hash; Identity sin engine. | [Builds](evidence/integral-builds-summary.md). Vercel inició2 builds con raíz Game y fuentes exteriores; ambos cancelados por targetproduction. Sin compilación ni runtime PASS; ver auditoría§12. |
| POC-02 — Gin y adaptador HTTP local/remoto | **PASS** | **BLOCKED** | Handlers exportados API probados; tres ejecutables reales y conectividad autenticada200/204; test/vet/build y raceLinux correctos. | [Runtime](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/summary.json), [builds](evidence/integral-builds-summary.md). Sin runtime remoto ejecutado. |
| POC-03 — Neon, pooling y aislamiento SQL | **PASS** (PostgreSQL local) | **BLOCKED** | DB-01–08: bases/roles propios, CONNECT cruzado42501, pool máximo2, up/down, persistencia tras reiniciar PostgreSQL conservando volumen; AI sin DB. | [Seguridad/DB](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/security.json), [reinicio](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/summary.json). No prueba Neon/PgBouncer. |
| POC-04 — Idempotencia y concurrencia transaccional | **PASS** | **BLOCKED** | TX-01–07: primer commit, retry, payload incompatible, dos HTTP concurrentes, rollback, pérdida de respuesta, reinicio Game/PID distinto. Estado/efectos/recibos3/3/3 sin duplicación. | [TX](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/transactions.json), [retry](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/retry.json). Dos contendientes, sin claim de escala productiva. |
| POC-05 — Commit, fallo Ably y recuperación REST | **PASS** (doble local) | **BLOCKED** | NOT-01–07: commit auditado SQL; entrega, pérdida, duplicación, orden invertido, respuesta perdida, listener reiniciado, timeout; REST conserva autoridad. | [Casos](evidence/integral-notifications-cases.json), [race](evidence/integral-builds-summary.md). Broker y cliente de prueba, no Ably/UI productiva; sin exactly-once. |
| POC-06 — Mismo origen, cookies y protección de entornos | **PASS** (controles locales) | **BLOCKED** | Orígenes fijos y rechazo de destinos maliciosos; Next conserva auth/401/cabecerascookie; fixtures ausentes de HTML/logs/nueve JS. | [Web](evidence/integral-builds/20261022T182305Z-31b332eb56cf4ddfbab35a1d970a9b7b/web-smoke.json). Sin navegador/OIDC/cookieGo/TLS/previews reales. |
| POC-07 — Comunicación entre servicios y permisos | **PASS** (sonda local) | **BLOCKED** | Game proceso→Identity/AI204; ausente/inválida/audiencia401, recurso ajeno403, AI-write401, firma Game-user con claveAI401. AI caído→503 controlado, recuperado→204. | [Seguridad](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/security.json), [fallo](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/internal-down.json). HMAC provisional, claveS2S compartida; no aislamiento productivo frente a compromiso. |
| POC-08 — Tokens Ably limitados por recurso | **PASS** (simulación) | **BLOCKED** | CAP-01–08: suscripción al canal exacto; publish cliente/otro recurso/expiración/firma/missing denegados; publicador distinto. | [Capabilities](evidence/integral-notifications-cases.json), [límites](tests/notificationprobe/README.md). Sin SDK/token/cuenta Ably. |
| POC-09 — Regiones y mediciones de latencia | **PASS** (mediciones locales) | **BLOCKED** | Tiempos buildfrío/HTTP/SQL/recuperación y muestra puntual de CPU/memoria; sin región elegida o latencia remota inferida. | [Muestra](evidence/integral-environment.json), [runtime](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/summary.json), [builds](evidence/integral-builds-summary.md). Sin p95/SLO. |
| POC-10 — Cancelación y presupuesto de AI | **PASS** (sonda sintética) | **BLOCKED** | Context cancelado/deadline/work limit; proceso work2→200/work201→422/work100→408; base2 rechazada frente a SQLrev3 sin efecto. | [AI](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/security.json), tests/raceAI. Guardia aislada, sin Negamax ni comando de jugada AI. |
| POC-11 — Cuotas, consumo y objetivo de coste cero | **PASS** (inventario/modelo local) | **BLOCKED** | DB1CPU/512MiB; verificadores efímeros2CPU/4GiB; gasto contratado adicional0. Modelo remoto acotado, cuotas reales sujetas a cuenta/gate. | [Consumo](evidence/integral-environment.json), [plan](REMOTE_PLAN.md), [auditoría](REMOTE_AUDIT.md). No saldo ni coste mensual de cuenta demostrados. |

Los PASS siguen las expectativas **locales** de la nueva solicitud. No existe GO global ni proveedor real validado. OIDC/navegador permanecen fuera de ese alcance; lint no definido y race Windows no disponible están declarados, mientras race Linux sí se ejecutó. DB-01–08, TX-01–07, comandos exactos, fallos y correcciones se documentan en [INTEGRAL_REPORT](INTEGRAL_REPORT.md).

La ejecución externa requiere **APRUEBO SPIKE REMOTO**, cuentas válidas, cuotas y seguridad efectiva. No hay provisión, pago, despliegue o merge autorizado por GO_LOCAL.

## Adenda del intento autorizado

La autorización posterior satisface el gate de permiso. `REMOTE-001` fue rechazada por Neon con **401** a las 20:27:50 UTC, sin retry ni refresh. La condición de parada del propietario impidió comenzar etapa A y las pruebas B–E. No se reclasificó ningún PASS histórico local.

| Conjunto | Estado remoto | Ejecución de este intento | Evidencia pendiente |
|---|---|---|---|
| POC-01 / POC-02 | BLOCKED | NOT_RUN | Cuatro builds, módulo compartido, Gin y HTTP protegido reales |
| POC-03 / POC-04 | BLOCKED | NOT_RUN | Neon, roles negativos, transacciones, idempotencia y concurrencia |
| POC-05 / POC-06 / POC-07 / POC-08 | BLOCKED | NOT_RUN | Ably real, REST, cookies, protección, S2S y capacidades negativas |
| POC-09 / POC-10 / POC-11 | BLOCKED | NOT_RUN | Región, latencia, cancelación AI, cuotas y consumo remoto |

Consumo observado: **una solicitud API a proveedores**, **cero builds**, **cero recursos/credenciales creados** y **ninguna retirada necesaria**. Se realizaron además consultas públicas de documentación cuyo tráfico HTTP interno no expone la herramienta; por ello no se certifica un total global exacto ni saldo disponible del techo de 100. Véase [contabilidad y límites](REMOTE_PLAN.md#7-adenda-de-ejecución-y-contabilidad). No se consumió reserva de retirada mediante pruebas funcionales, porque no se ejecutaron.

## Continuación de precondiciones

Las diez lecturas nuevas satisfactorias verifican sesiones/planes Neon y Vercel, inventarios de proyectos y automatizaciones del alcance inspeccionado; **no sustituyen pruebas POC**. El 401 anterior permanece en la suma: **11 consumidas + 79 disponibles para otras operaciones + 10 reservadas = 100**. SQL y documentación se excluyen de este cupo por confirmación expresa, sin ocultar sus contadores.

La secuencia detallada F01–F51 y sus operaciones administrativas suma **108 más prechecks/alta PAT pendientes** con los ahorros ya acreditados; requiere otra revisión antes de crear. No se declara que sea un mínimo universal ni se elimina cobertura para forzar 100. Control API Ably, cuotas actuales, instrumentación de sus mecanismos y retirada completa siguen pendientes. Ningún PASS/FAIL remoto se infiere de las consultas de cuenta; todas las once filas remotas permanecen BLOCKED, ejecución NOT_RUN.

### Cierre remoto de esta continuación

Las cuotas y la configuración All Deployments tienen evidencia real de cuenta. POC-01 registra dos intentos **CANCELED** al detectar `target=production`; POC-02–POC-10 funcionales permanecen **NOT_RUN/BLOCKED**. POC-11 tiene cuotas previas verificadas, pero no consumo de una ejecución funcional completa: permanece **BLOCKED**. No hay PASS remoto. La retirada está **PASS**: cuatro proyectos y dos deployments propios ausentes; no se creó nada en Neon/Ably. El límite de solicitudes fue retirado por el propietario y no se usa como bloqueo vigente. El bloqueo actual es la restricción de destino; propuesta de excepción en REMOTE_PLAN §10. Once filas, enlaces, fences y JSON comprobados sin errores; GO_LOCAL conserva su objeto original.
<a id="remote-matrix-20261025"></a>

## Matriz remota vigente — 2026-10-25

**GO_REMOTE limitado al spike.** Historia anterior conservada. [Pruebas, fallos corregidos, desviaciones y límites](REMOTE_AUDIT.md#remote-close-20261025).

| ID | REMOTO | Evidencia y alcance |
|---|---|---|
| POC-01 | PASS | Seis READY sin caché, cuatro roots; motor único Game/AI, Identity sin engine. Ocho intentos acumulados. Logs/manifiesto42 archivos. |
| POC-02 | PASS | Gin/HTTP Go1.26.0 real protegido: runtime-* y CONNECT-*, sin autorización denegado. |
| POC-03 | PASS limitado | Neon18/pgx pooled, up/down/up directo, roles/CONNECT/DDL42501, pool2/rol4 y dos backends observados. Sin medir colas/clientes PgBouncer. |
| POC-04 | PASS | TX primer commit/retry/conflicto/carrera200-409/rollback/nueva instancia, SQL/REST4-4-4. Respuesta perdida modelada, no TCP. |
| POC-05 | PASS | Commit/relectura antes de Ably, entrega real, fallo401 conserva commit, avisos2/1/2 y reconexión REST. Consumidor sintético. |
| POC-06 | PASS | All Deployments/OIDC, Web→Go, cookie Chrome HttpOnly real y HTML/seis chunks sin claves. No sesiones de producto. |
| POC-07 | PASS | S2S válido, negativos401/403, OIDC inválido503/upstream401 y recuperación204. No corte físico ni JWKS productivo. |
| POC-08 | PASS | JWT message-subscribe exacto300s, entrega y CAP01–10 negativos; claves separadas. Error SSE leído dentro del stream. |
| POC-09 | PASS limitado | iad1/aws-us-east-1, HTTP/procesos nuevos y cronómetros C6 SQL/Ably separados. Muestras únicas operador→proveedores; sin p95 ni reactivación Neon probada. |
| POC-10 | PASS | AI sintética200, work201→422, deadline40ms→408 y stale409. Sin Negamax/fuerza medida. |
| POC-11 | PASS limitado | Hobby/Free, uso leído,0USD contratado, sin upgrades y retirada completa. SQL80 previsto superado; conteo reconstruido declarado, métricas con retardo, sin garantía de coste futuro. |

Revisión independiente cerró el hallazgo C6 con [seguimiento](evidence/remote-latency-followup-20261025.json). Retirada PASS: cuatro Vercel/seis deployments y ambas generaciones Neon/Ably ausentes; ocho claves de aplicación revocadas. Recursos anteriores preservados. Ningún PASS acepta funcionalmente producto ni sustituye PRE-09.
