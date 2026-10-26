# Gate 3.5 — verificación integral local, 2026-10-22


> **Estado vigente — 2026-10-25 (Caracas): GO_REMOTE, limitado al spike desechable.** Seis builds remotos READY, ocho intentos acumulados contando dos cancelados históricos; pruebas reales Neon/Ably/OIDC/cookie y seguimiento C6 completados. Todos los recursos propios retirados y recursos anteriores preservados. [Cierre, evidencias y límites](REMOTE_AUDIT.md#remote-close-20261025). Las notas anteriores de bloqueo que siguen son cortes históricos, sustituidos por esta adenda. No acredita producto, aprobación automática de propuestas ni generación del lote 4.

> **Cierre actual: BLOCKED_REMOTE por destino Vercel, recursos propios retirados.** Cuotas y protección gratuita verificadas; techo100 retirado por el propietario. Dos builds fueron cancelados porque Vercel devolvió `production` incluso solicitando staging; cero secretos y alias sin asignar. Los cuatro proyectos y ambos deployments se retiraron: inventarios finales vacíos.41 API directas más navegación UI; sin lote4. [Evidencia y decisión pendiente](REMOTE_AUDIT.md#12-prueba-real-de-protección-y-destino-inicial-bloqueo-específico-recursos-retirados). Notas anteriores conservadas como historia.

> **Estado vigente: cuotas verificadas; techo de100 solicitudes retirado por el propietario.** Se conservan18 lecturas API directas y errores históricos; las lecturas UI posteriores no equivalen a un total HTTP exacto. Sigue el spike remoto previamente autorizado con0USD y máximo6builds. Aún no hay PASS remoto ni permiso de lote4. [Auditoría actual](REMOTE_AUDIT.md#11-decisión-posterior-y-comprobación-de-cuotas--2026-10-22). Las notas siguientes son cortes históricos.

> **Última continuación: Control API Ably verificada; `BLOCKED_REMOTE` por presupuesto/cuotas pendientes.** Se observaron HTTP 200, cuenta esperada y apps preexistentes Player chat/9 stones, sin modificarlas. Contador **16/100**, diez reservadas, 74 para otras operaciones; cero builds o recursos creados. [Detalle y repetición explícita contabilizada](REMOTE_AUDIT.md#10-control-api-autenticada-con-el-token-posterior-del-propietario). El bloqueo histórico de credencial Ably ya está resuelto; GO_LOCAL permanece intacto.

> **Última continuación: `BLOCKED_REMOTE`, 12 solicitudes operativas, cero builds y cero escrituras externas.** La clave de aplicación aportada permitió intentar una lectura Ably: GET estadísticas rechazado **401 / 40160**, sin retry. No sustituye al PAT exigido por Control API. Se conservan 10 solicitudes para retirada y 78 para otras operaciones. [Detalle sanitizado](REMOTE_AUDIT.md#9-clave-de-aplicación-proporcionada-por-el-propietario); los contadores anteriores son históricos. No se incluyó la credencial en los archivos del repositorio.

> **Continuación vigente: `BLOCKED_REMOTE`, sin aprovisionamiento.** El propietario renovó Neon y confirmó que el techo incluye 100 solicitudes operativas acumuladas, con SQL/documentación aparte. Neon/Vercel ya responden con Free/Hobby; el contador es **11**, incluidas las diez lecturas nuevas y el 401 histórico. Quedan presupuesto completo, Ably y cuotas por cerrar. [Evidencia](evidence/remote-preconditions-budget-20261022.json). Esta actualización prevalece sobre el bloqueo histórico de la sesión Neon descrito a continuación.

> **Adenda vigente — intento remoto, 2026-10-22, 20:27 UTC: `BLOCKED_REMOTE`.** Se recibió `APRUEBO SPIKE REMOTO`. La primera comprobación operativa, de la sesión CLI de Neon, devolvió **401**; se aplicó la parada obligatoria anterior al aprovisionamiento. Hubo **una solicitud API a proveedores, cero builds y cero recursos o credenciales creados**. La investigación pública se contabiliza por separado y no permite certificar el total HTTP global. [Registro sanitizado](evidence/remote-execution-20261022.json). Los resultados locales siguientes se conservan como historia; sus solicitudes de permiso y estados antiguos de cuenta quedan subordinados a esta adenda.

**Conclusión de salida: GO_LOCAL.** Se ejecutaron las pruebas locales obligatorias del spike existente con PostgreSQL real. Es un gate del experimento local, sin GO de plataforma, MVP o seguridad de producción. [Registro observado](evidence/integral-local-gate.json). La validación remota requiere autorización y comprobaciones de cuenta.

La ejecución partió de `cf5d3973adc977f0ec09c738cbe6cfc078883593`, rama `spike/infra-feasibility-01`, sin reconstruir otro experimento. Los cambios posteriores están en el diff de entrega y los hashes de fuentes de builds. Los resultados previos no se reutilizaron como PASS de esta ejecución.

## A. Entorno comprobado

| Elemento | Observación nueva |
|---|---|
| Host/workspace | Windows 11 Pro 10.0.26200 amd64, PowerShell 7.6.5; `C:/Users/swmfi/Documents/ChatGPT/9 Stones` |
| Git | Remoto `https://github.com/SebasChaconDev/toguz-kumalak.git`, rama local existente, árbol inicialmente limpio |
| Docker/Compose | Client y Server 29.9.2; Server Linux amd64; Compose 5.6.1 |
| Contexto | `desktop-linux`, endpoint local `npipe:////./pipe/dockerDesktopLinuxEngine`; sin cambiar contexto, WSL o permisos globales |
| Go/Node/npm | Go 1.26.2 Windows/Linux; Node 24.13.0; npm 11.6.2 |
| PostgreSQL | 18.6 con digest fijado en Compose; healthy; solo loopback `127.0.0.1:15435` |
| Unidades | Next.js; tres servidores Go/Gin independientes; motor sustituto Go; biblioteca probe; PostgreSQL Compose |

[Entorno y consumo puntual](evidence/integral-environment.json). Se inspeccionaron puertos, etiquetas y montajes antes de operar. El Compose existente contiene solo PostgreSQL con red/volumen propios, healthcheck, una CPU y 512 MiB. No había conflicto del puerto 15435 con otro proyecto. No existen Dockerfiles: las apps Go se ejecutaron como procesos locales y Next desde su build real. No se alteraron datos o contenedores ajenos.

## B. Ejecución reproducible

Orden desde la raíz del repositorio. No solapar fixtures, reinicio DB ni usuarios de puertos 18100–18103:

```powershell
docker compose --env-file ./spikes/infra-feasibility/.local/compose.env -p toguz-infra-spike -f ./spikes/infra-feasibility/infra/compose.yaml config --quiet
docker compose --env-file ./spikes/infra-feasibility/.local/compose.env -p toguz-infra-spike -f ./spikes/infra-feasibility/infra/compose.yaml up --build -d --wait --wait-timeout 120
./spikes/infra-feasibility/tests/Run-Integration.ps1
./spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase Web
./spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase Go
./spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase LinuxRace
./spikes/infra-feasibility/tests/notificationprobe/Run.ps1
./spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase NotificationRace
./spikes/infra-feasibility/tests/Verify-IntegralRuntime.ps1
```

También se ejecutaron inspecciones no destructivas Git, `docker version`, `docker info` con campos seleccionados sin proxies/secretos, `docker context show`, `docker compose version`, `docker ps` y verificación del listener. `Run-Integration.ps1` restablece exclusivamente las tablas sintéticas existentes; por eso debe preceder a los otros harnesses. Estos insertan IDs nuevos sin borrar tablas. No se ejecutó down con volúmenes, prune, limpieza global o reinicio global Docker.

| Control | Resultado nuevo |
|---|---|
| Go Windows, cinco módulos | Test/vet/build: quince comandos exit 0 |
| Builds aislados | Tres pares servidor/API, seis builds exit 0, cachés GOCACHE/GOMODCACHE nuevas, GOWORK=off; Identity sin copiar engine |
| Race Linux | Cinco módulos y notificationprobe exit 0, Go 1.26.2 + GCC; no carreras detectadas |
| Race Windows | No ejecutado nativamente: gcc ausente; no se atribuye el resultado Linux a Windows |
| Next | Misma app/lockfile, npm ci con instalación/caché nuevas, test/build/smoke exit 0; lock intacto; lint no definido |
| SQL existente | Cinco tests de integración y dos tests previos de la app exit 0 en ejecución nueva |
| Notificación/capacidad | Quince escenarios explícitos PASS con Game/PostgreSQL reales y broker/capacidades simulados |
| Procesos/reinicio DB | Game PID 24004→7992; PostgreSQL reiniciado conservó mismo volumen y `3/3/3` |

[Comandos, versiones, tiempos y logs de builds](evidence/integral-builds-summary.md). Race utilizó `golang:1.26.2-bookworm`, digest `sha256:47ce5636e9936b2c5cbf708925578ef386b4f8872aec74a67bd13a627d242b19`, contenedores de dos CPU/4 GiB, fuentes de solo lectura y retirada `--rm`. No hubo instalación global de compilador Windows.

### DB-01–DB-08

DB-01/02/03: Identity y Game conectaron con bases/propietarios distintos y respondieron conectividad 200/204. Ambas conexiones cruzadas fueron rechazadas con SQLSTATE **42501**, no por un fallo de red. DB-04: AI se inició con entorno permitido explícito, sin DSN ni PG*, y respondió `database=none` autenticado.

DB-05/06: estado, efecto y recibo se confirmaron juntos; fallo anterior al commit no dejó escrituras parciales. DB-07: [antes](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/snapshot-before.json) y [después](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/snapshot-after.json) del reinicio muestran revisión 3, efectos 3, recibos 3, volumen esperado `toguz-infra-spike_spike_pg` conservado.

DB-08: up/down y reaplicación con propietarios sintéticos, nunca migración al arrancar una Function. El mínimo local utiliza propietario también como rol de app; no acredita segregación DDL/DML productiva. El plan remoto propone roles de migración distintos. [Ejecución final](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/summary.json).

### TX-01–TX-07

Formato: revisión/efectos/recibos. [Evidencia por caso](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/transactions.json) registra entrada, respuesta, snapshots y tiempo; el driver decodifica las decisiones HTTP.

| Caso | Antes → entrada/fallo → respuesta → PostgreSQL posterior |
|---|---|
| TX-01 | `0/0/0` → first/revisión0 → 200, decisión1/1 → `1/1/1` |
| TX-02 | `1/1/1` → intención idéntica → 200/misma decisión → `1/1/1`; retry tardío devuelve decisión original1/1 y conserva estado3/3/3 |
| TX-03 | `1/1/1` → mismo ID con payload distinto → 409 → `1/1/1` |
| TX-04 | `1/1/1` → dos HTTP sobre revisión1 → una200/una409 → `2/2/2` |
| TX-05 | `2/2/2` → fallo inyectado en Store antes de commit, sin endpoint HTTP de inyección de fallos → error → `2/2/2` |
| TX-06 | `2/2/2` → proxy recibe commit200 de Game y corta respuesta → cliente pierde respuesta; retry200 → `3/3/3` |
| TX-07 | `3/3/3` → reiniciar ejecutable Game, PID distinto → retry200 → `3/3/3`; repetido después del reinicio DB |

No se deduce integridad de un 200 aislado: se verifican filas y recibos. La deduplicación no depende de RAM de Game.

### Notificación y seguridad

NOT-01–07 comprueban entrega, pérdida, duplicados, orden invertido, respuesta HTTP perdida, cierre/reapertura del listener y falta de respuesta/cancelación. La publicación consulta DB por otra conexión y exige commit ya visible. El cliente vuelve a REST y recupera estado autoritativo sin incrementar efectos por avisos. CAP-01–08 restringen recurso/acción/tiempo/firma y separan suscriptor/publicador. [Quince casos](evidence/integral-notifications-cases.json).

El broker es un **doble HTTP**, la reconciliación un **cliente Go de prueba**, y la conexión Notify se inyecta por el hook de Store. Reiniciar el listener conserva el objeto del doble: no demuestra recuperación de su RAM/proceso ni de Ably. El CLI Game no está conectado a un proveedor. Sin REST no se garantiza recuperación; no hay exactly-once.

El proceso Game llamó realmente a Identity y AI mediante `/internal/check/:service`: solo principal Game, destinos loopback fijos, timeout250ms. Cerrar AI produjo 503 controlado; restaurarlo produjo204. No admite una URL arbitraria. Faltas de credencial/firma/audiencia se rechazaron; recurso ajeno403 y AI→comando Game401. Cada audiencia del runner recibe una clave de usuario distinta; intentar firmar un usuario Game con la clave AI produjo401 y SQL intacto. [Seguridad](evidence/integral-runtime/integral-7d08fb85f2ed4c62b505c985d76e8d5b/security.json).

El mecanismo HMAC y su clave S2S compartida son provisionales; no prueban aislamiento completo ante compromiso de servicio, JWT/JWKS, CSRF, Google OIDC ni seguridad productiva. Los hijos reciben solo variables de sistema permitidas, claves sintéticas y su DSN cuando corresponde. Se revisaron respuestas, logs y JavaScript contra fixtures, incluidas contraseñas/keys por audiencia; logs finales de procesos: cero bytes. Los escaneos siguen siendo limitados a valores/patrones conocidos.

Se probaron contexto cancelado/deadline/work limit AI y rechazo de base2 frente a revisión SQL3 sin efecto DB. La guardia es una función experimental aislada, no un comando para confirmar jugadas AI ni Negamax.

## C. Correcciones y regresión

- Se añadió la sonda S2S saliente que faltaba. `TestInternalProbe` se observó RED antes del cliente; se conservaron ambas pasadas fallidas coincidentes y se repitieron test/vet/build/race y contextos aislados con implementación corregida.
- Se sustituyó la inferencia de reinicio de Store por una prueba de ejecutable/PID; se añadió reinicio DB con snapshots separados y comprobación del volumen esperado.
- Se añadieron los escenarios de entrega/capacidad que faltaban, con oráculos RED para admisión insegura y entrega ausente. El broker no es una nueva dependencia obligatoria del producto.
- La revisión corrigió cleanup que podía abortar antes de restaurar entorno, herencia de secretos por AI, escaneo solo de DSN completos y comparación insuficiente de decisiones HTTP. Las regresiones finales pasaron sin hallazgos bloqueantes.

No se alteraron specs aprobadas, decisiones, ADR, TK-N1-1.0.0 o engine. Los módulos nuevos `tests/processprobe` y `tests/notificationprobe` son harnesses sin despliegue propio. Los textos de fallos no se eliminaron para conseguir PASS.

Comprobaciones de entrega: 93 enlaces relativos, nueve Markdown, 67 JSON y once filas POC sin errores; ningún cambio fuera del spike. Los 45 hashes de fuentes observadas en las pasadas Go/race coinciden con las fuentes de entrega. El escaneo comparó diez valores sintéticos conocidos (incluidas tres claves derivadas por audiencia) contra 256 archivos publicables y 22 ejecutables Windows originales/integrales, sin coincidencias ni patrones de credenciales detectados. [Documentos](evidence/document-check.json) y [exposición](evidence/exposure-check.json) registran alcance y límites. Se analizó la sintaxis PowerShell de los runners sin errores y un revisor independiente contrastó oráculos y claims.

## D. Evaluación y límites

Comprobados cuatro builds separados; engine único v0.1.0 con replace local y hash coincidente; Identity sin engine; SQL separado; atomicidad, retry, concurrencia, persistencia entre procesos/contenedor y recuperación REST. Los paquetes API compilan y se ejercitan localmente. No demuestra Vercel, módulos fuera de Root Directory, S2S protegido remoto o cookies reales.

Muestra puntual de PostgreSQL: **25.36 MiB/512 MiB, 6.14 % CPU, nueve procesos**, después de las pruebas. No es pico/capacidad/SLO. Tiempos por operación/build: muestras locales bajo carga compartida, sin p95 ni extrapolación mensual. El coste contratado adicional es cero; cuotas remotas disponibles todavía desconocidas.

Quedan fuera del PASS local: OIDC/navegador/cookies reales, Ably, Neon y Vercel; race Windows nativo no disponible, compensado por race Linux ejecutado; lint no definido. npm ci no acredita ausencia de vulnerabilidades: npm audit anterior fue inconcluso. Los límites son del alcance explícito local, no pruebas omitidas presentadas como ejecutadas.

## E. Siguiente gate

[Matriz POC](TEST_MATRIX.md), [auditoría de proveedores](REMOTE_AUDIT.md), [plan remoto](REMOTE_PLAN.md). La preparación comenzó después del registro GO_LOCAL. No se creó ni desplegó recurso externo, activó pago, fusionó PR o inició lote4/MVP.

Publicación retenida: la auditoría no pudo descartar automatización GitHub/Vercel. Se conserva el resultado en la rama local; no se hizo push ni se abrió Draft PR. Título reservado: `spike: local infrastructure verification and remote feasibility plan`.

Se solicita **APRUEBO SPIKE REMOTO** para ejecutar el plan acotado, además de cuenta/plan/cuota y seguridad efectiva verificados. GO_LOCAL no sustituye ese permiso.

## F. Intento remoto autorizado — parada anterior a etapa A

El propietario autorizó únicamente cuatro proyectos Vercel protegidos, un proyecto Neon nuevo y una aplicación Ably nueva, con coste adicional autorizado **0 USD**, máximo **100 solicitudes instrumentadas totales**, al menos **10 reservadas para retirada** y máximo **seis builds**. Esta autorización sustituye la petición del apartado E, sin eliminar las precondiciones.

Al comenzar se observó HEAD `94094620637b6705193ce984e8e5d54e39a3e45e`, rama `spike/infra-feasibility-01` y árbol limpio. El registro `GO_LOCAL` permanece intacto; no se repitieron pruebas de producto ni builds. `REMOTE-001`, GET Neon `/api/v2/users/me`, respondió **401** el **2026-10-22 a las 20:27:50 UTC**. La credencial CLI utilizada tenía caducidad **20:15:13.583 UTC**. No hubo refresh, retry ni redirect. Esto demuestra el rechazo de esa credencial, sin demostrar que la sesión del navegador esté cerrada o que la cuenta carezca de permisos.

Se aplicó «Si falla una precondición, detente sin aprovisionar». No se completó una nueva comprobación de los tres planes, cuotas, inventarios o automatizaciones. Las lecturas del preflight anterior no se presentan como actuales. No se hizo push, configuración externa, contratación, despliegue, emisión de credenciales ni modificación de **9 Stones** o **Player chat**. No se volvió a verificar su estado externo después de parar.

Los once POC remotos quedan **BLOCKED**, con ejecución **NOT_RUN**. No hay incompatibilidad arquitectónica nueva demostrada que justifique `NO_GO_REMOTE`, ni evidencia para `GO_REMOTE`. La retirada es **no necesaria en este intento**, pues no hay recursos o credenciales propios. Se conserva la autorización recibida: para reanudar hacen falta una sesión Neon válida, repetir las precondiciones y demostrar un presupuesto ejecutable. La contabilidad y sus límites figuran en [el plan actualizado](REMOTE_PLAN.md#7-adenda-de-ejecución-y-contabilidad). PR #3, lote 4 y producto no fueron modificados.

## G. Continuación de precondiciones y revisión del cupo

Se comprobó la sesión renovada Neon (200), su organización Free y el único proyecto 9 Stones; Vercel confirmó OWNER/Hobby activo y cero proyectos/deployments/integraciones. GitHub confirmó main y PR #3 intactos, cero workflows/webhooks/deployments/environments. Son **prechecks de lectura**, no POC de producto. GO_LOCAL mantiene su blob original `1c22815c323b61635a334d3008f2cbbcf1c5b361`; no se repitieron pruebas locales ni cambió el código.

El contador mantiene el 401 y registra **11/100**, con **10 reservadas** y **79 como máximo para operaciones distintas de retirada**. La decisión expresa del propietario excluye documentación y SQL del cupo operativo; sus contadores se conservan separados. La propuesta completa revisada suma **108 más pendientes**, incluso al validar protección en el POST de proyecto y enviar env en el despliegue posterior. No se presenta esa propuesta como un mínimo universal ni como ejecución que cabe en 100. [Detalle y retirada](REMOTE_PLAN.md#8-continuación-tras-renovar-neon-unidad-y-presupuesto-revisados).

Faltan acceso Control API de Ably disponible para consultas instrumentadas, cuotas residuales actuales y prueba de un plan íntegro de operación/retirada dentro del cupo. Leer la UI Ably ya renderizada no acredita esas condiciones. Resultado **BLOCKED_REMOTE** antes de crear recursos: ninguna credencial emitida por el agente, recurso creado, build, push o cambio al PR #3. Todos los POC remotos siguen BLOCKED/NOT_RUN. No hay retirada que ejecutar en esta continuación.

## J. Cierre de esta continuación y comprobación documental

Resultado actual **BLOCKED_REMOTE** por la promoción inicial `production` de Vercel, observada dos veces y retirada. Las cuotas y All Deployments gratuitos quedaron comprobados; los casos de integración no pasaron por ello. Dos intentos de build cancelados, sin compilación completada, cero secretos y cero SQL remoto. Las cuatro configuraciones creadas y los dos deployments se eliminaron y su ausencia quedó verificada. No se creó ningún recurso Neon o Ably; no hubo push ni cambio al PR #3 o inicio del lote 4. El techo de 100 solicitudes fue retirado por el propietario, sin borrar las 41 llamadas API directas ni sus fallos; tráfico UI posterior registrado con su limitación de medida.

Comprobaciones documentales: HEAD y blob/SHA256 GO_LOCAL intactos; exactamente seis rutas autorizadas cambiadas, sin código. Nueve Markdown, 69 JSON válidos, 129 destinos de archivo locales y 24 anclas sin ausencias, fences cerrados, once filas POC únicas y `git diff --check` correcto. Escaneo en memoria contra credenciales conocidas: cero coincidencias en los seis archivos. Ledger de 40 llamadas nuevas más la histórica conserva IDs REMOTE-001–041, dos cancelaciones y retirada completa. La primera comprobación de status falló por la coexistencia histórica `status`/`httpStatus`; se corrigió el lector, preservando las entradas originales. Estos son controles documentales y de contabilidad, no pruebas de producto.

La propuesta que necesita decisión humana está en REMOTE_PLAN §10: excepción solo para la inicialización automática de cuatro proyectos aislados protegidos y máximo ocho builds acumulados, incluidos los dos cancelados. No se ejecuta hasta aprobación expresa. La credencial Neon consultada expiró después de sus lecturas válidas; debe renovarse/revalidarse por el mecanismo oficial antes de futuras operaciones, sin compartir secretos en el chat.
<a id="integral-remote-close-20261025"></a>

## Cierre de factibilidad remota — 2026-10-25 Caracas

**GO_REMOTE acotado al spike.** [Once filas vigentes](TEST_MATRIX.md#remote-matrix-20261025), [auditoría/evidencias/límites](REMOTE_AUDIT.md#remote-close-20261025). GO_LOCAL conserva su objeto y hash originales.

Se verificaron Vercel/Go/Gin, motor compartido, Neon aislado, transacciones/concurrencia, Ably real posterior al commit, permisos, cookie Chrome y OIDC entre proyectos protegidos. Seis builds READY sin caché y dos cancelados históricos suman ocho. Código nuevo solo experimental bajo spikes/infra-feasibility; producto sin implementar.

Revisión independiente detectó cronómetros C6 ausentes. El seguimiento de reemplazo aportó519,170ms pooled,287,993ms transacción y258,356ms publicación→recepción; una muestra operador→proveedores, sin atribuirla a Game en Vercel. Otro dictamen cerró ese hallazgo.

Desviación explícita: SQL80 previsto superado;105 operaciones principales reconstruidas conservadoramente más12 instrumentadas. Sin conteo wire exhaustivo ni excepción SQL literal inventada. Se mantuvieron0USD, planes gratuitos y carga acotada.

Retirada final verificada: cuatro proyectos Vercel/seis deployments, dos generaciones secuenciales Neon/Ably y ocho claves de aplicación revocadas. Recursos previos del propietario preservados. Sin claims de p95, seguridad completa, fuerza IA, coste productivo indefinido, reactivación Neon, clientes/colas PgBouncer, corte físico S2S o caída TCP real.

La revisión documental independiente de las once specs del PR3 solo encontró dos precisiones menores05.02. Corregirlas y registrar la autorización efectiva del propietario sin inventar la frase «APRUEBO LOTE3». Lote4 mantiene ocho specs04.01–04.04/08.01–08.04 y55 decisiones, todavía sin generar; pruebas de producto PRE-09 pendientes.
