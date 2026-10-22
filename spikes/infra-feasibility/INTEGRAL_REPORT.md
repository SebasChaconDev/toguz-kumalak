# Gate 3.5 — verificación integral local, 2026-10-22

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
