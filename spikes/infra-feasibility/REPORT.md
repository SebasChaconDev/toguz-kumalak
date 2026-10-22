# Gate 3.5 — informe de viabilidad previo a recursos externos

**Estado vigente: GO_LOCAL**, verificación integral del **2026-10-22**. La [adenda §10](#integral-20261022) y el [informe integral](INTEGRAL_REPORT.md) sustituyen los bloqueos locales y limitaciones de cobertura históricos de este informe. POC-01–11 pasan en el alcance local solicitado; todos siguen BLOCKED para proveedor remoto. Publicación retenida por automatización externa no descartada. Presupuesto: **0 USD adicionales**; ningún recurso externo creado o desplegado. El siguiente gate es **APRUEBO SPIKE REMOTO**.

**Entrega anterior, conservada como historia:** Docker fue reparado y actualizado; PostgreSQL local y los cinco casos de integración SQL pasaron. Los cuatro builds locales funcionan y comparten una única fuente del sustituto Go. La sección 9 sustituyó el bloqueo Docker inicial. Sus resultados anteriores no se contabilizan como ejecuciones de la verificación integral nueva.

## 1. Base, alcance y autoridad

Base `main` vigente verificada por GitHub: `aaeeace21860305091c0351bedb0aca0df386816`; rama nueva `spike/infra-feasibility-01`. PR #3 **OPEN**, HEAD `9db0f3e9d4bf1fb642697103d9232bcfaef42205`, sin fusión ni cambio de estados. No existía rama experimental homónima remota/local. Todo el cambio se limita a `spikes/infra-feasibility/`.

Fuentes consultadas: AGENTS, GOALS, PLANS, REG y CONF; 01.01–01.03; ADR-0001/0002/0004/0005/0006; 03.01–03.04 y AD-NORM-001. Los contratos 05.01–05.04 y 07.02 se consultaron en `origin/docs/specs-lote-03` como propuestas **no aprobadas**. La autorización posterior del Gate 3.5 permite el experimento aislado y no cambia decisiones de producto, PRE-09, TK-N1-1.0.0 o los estados documentales anteriores.

Fecha de revisión local 2026-10-21 Caracas; ejecución UTC 2026-10-22. No se consultaron reglas ni motores externos para reconsiderar el perfil normativo.

## 2. Arquitectura realmente probada y resultados

Next independiente y tres módulos Go/Gin independientes, con arranque loopback y entradas `api/index.go` adaptadas a `net/http`. Game y AI importan el mismo `engine v0.1.0`; Identity no lo importa. `engine` solo calcula un SHA-256 determinista, sin reglas del juego. `modules/probe` es biblioteca de soporte, no otro microservicio. No se probó el MVP.

Comandos, herramientas y reproducción: [README](README.md), [script Go](tests/Verify-Local.ps1), [web](apps/web/README.md). `go test -count=1 ./...` se ejecutó por las cinco raíces Go; las entradas HTTP se probaron con Handler real y servidores `httptest`. Inicialmente se compiló la suite `-tags=integration -run '^$'` sin ejecutar pruebas DB y después se intentó ejecutarla, con exit 1 por DB inaccesible. Ese bloqueo es histórico: la ejecución posterior a reparar Docker terminó **exit 0**, según §9. npm ci/build/test/smoke se ejecutaron desde la raíz web.

| Build local | Contexto/cache | Muestra de tiempo total | Resultado |
|---|---|---:|---|
| Identity | Fuentes de Identity + probe; GOCACHE/GOMODCACHE nuevos | 27.550 s | PASS |
| Game | Fuentes de Game + probe + única fuente engine; caches nuevos | 27.322 s | PASS |
| AI | Fuentes de AI + probe + misma engine; caches nuevos | 26.643 s | PASS |
| Web | Su propio manifiesto/lock y fuentes; `.next` ausente | 11.090 s | PASS |

`GOWORK=off`, `GOENV=off`, `GOTOOLCHAIN=local`, `CGO_ENABLED=0`; builds desde **cada raíz**, no desde un workspace global. Evidencia: [local-verification.json](evidence/local-verification.json), inventarios de módulos y metadatos de ejecutables. Versiones Go 1.26.2/Gin 1.12.0/pgx 5.11.0; web Next 16.4.2/React 19.3.1/TS 5.9.3, Node 24.13.0/npm 11.6.2. La compilación local no prueba la versión soportada ni el comportamiento efectivo del builder remoto.

Hash SHA-256 de `engine.go` en ambos contextos: `BD0D2F40296E32B6C832DEA9B22C0CB1B957E428D2DC01E20456C303AB14C973`. Ambos manifests requieren `v0.1.0` y resuelven `../../modules/engine`; es una versión **local experimental sin release/tag publicado**. Su fuente canónica permanece en una sola ruta. Los contextos de prueba copian la misma fuente para representar builds aislados; no se crean motores divergentes ni copias canónicas en los servicios.

## 3. PostgreSQL, transacciones y avisos: bloqueo inicial, resuelto el 2026-10-22

Al cierre inicial, Docker CLI 29.2.1 y Compose 5.1.0 existían, pero el backend fallaba antes del motor: no pudo acceder/eliminar el socket de inferencia `Docker/run/dockerInference`; el log registra parada de motores a `2026-10-22T02:19:05Z`. No se reseteó Docker/WSL ni se borró ese socket. Entonces no había contenedor, volumen, imagen PostgreSQL descargada ni migración aplicada por este spike. `config --quiet` validó Compose, sin imprimir variables. Ese estado es histórico; consultar la resolución y evidencia posterior en §9.

La configuración local tiene dos bases con propietarios distintos, CONNECT público revocado, pools pgx máximos 2 y migraciones up/down. En la primera entrega, cinco tests se detuvieron en el fixture: [salida histórica](evidence/sql-blocked.txt). Después de reparar Docker se ejecutaron sus aserciones contra PostgreSQL real, con PASS: [salida posterior](evidence/sql-local-20261022.txt), [casos actualizados](TEST_MATRIX.md). No se sustituyó PostgreSQL por memoria o SQLite para producir PASS.

El código experimental persiste recibo+estado+efecto en una transacción, serializa por agregado y consulta recibo antes de rechazar revisión vieja; el ensayo confirmó que un reintento idéntico recupera un commit cuya respuesta se perdió. Un doble local intenta aviso después de commit y REST vuelve a autorizar la consulta. La ejecución SQL está verificada para estos fixtures; no es validación del motor de juego ni de Neon/Ably reales. No hay promesa de outbox, entrega exactamente una vez o recuperación mientras REST esté indisponible.

## 4. Seguridad y cancelación observadas

PASS local de los subcasos de credencial ausente/malformada, firma inválida, audiencia, caducidad, duración excesiva, scope, tipo usuario/servicio y principal interno. Los handlers internos solo admiten servicio Game para la audiencia del destino; credenciales usuario no sirven como servicio. El 2026-10-22 también se verificaron acceso ajeno denegado y AI sin permiso de commit contra la DB real. Falta el flujo saliente real Game→AI/Identity. Los clientes de `httptest` demuestran fronteras HTTP de los handlers, no esa integración completa.

Autenticación HMAC **experimental**, claves separadas por kind compartidas con verificadores de prueba; no JWT asimétrico/JWKS, Google OIDC, sesiones ni seguridad contra un servicio comprometido. Una clave de verificación simétrica no demuestra que AI sea incapaz de falsificar una credencial si se compromete. No hay emisor HTTP ni bypass permisivo. Configuración ausente falla cerrada y los binarios locales escuchan solo en loopback. El remoto necesita revisión de controles y credenciales dedicadas antes de exponerlos.

AI rechaza entorno DB; su test ahora establece una configuración válida 200 antes de comprobar los tres nombres de DSN prohibidos. Trabajo sintético: 1–200 unidades, contexto de petición, deadline HTTP de 40 ms y cancelación. Pruebas observadas de resultado, límite, deadline y cancelación; guardia aislada de revisión obsoleta en Game. No se probó una carrera AI→transacción Game ni capacidad de Negamax.

La web comprobó cabeceras, rutas/query, Authorization y 401 a través de rewrites con dobles Node; en la repetición del 2026-10-22 se midieron ocho peticiones y 69 intentos de preparación. Los ocho tiempos observados van de **1.32 a 50.64 ms** en loopback, durante la actualización de Docker; no representan latencia Go, DB, Ably o regiones. Se revisaron HTML/logs/nueve archivos JS contra fixtures. Cookies transportadas manualmente no prueban enforcement en navegador ni autenticación completa. OIDC **NOT_RUN**.

El escaneo adicional de exposición compara los secretos sintéticos generados con archivos publicables y tres ejecutables; ver `evidence/exposure-check.json`. No se imprimen tokens, DSN o passwords. npm audit quedó **BLOCKED** por fallo de conexión del endpoint de advisories; instalación/build no acreditan ausencia de vulnerabilidades.

## 5. Auditoría de proveedores — solo lectura

Documentación oficial consultada el **2026-10-21 Caracas**, no reutilizada de cifras históricas del repositorio. Planes públicos y cuotas efectivamente disponibles son evidencias distintas.

| Proveedor | Conexión de cuenta observada | Cuenta/plan/cuota residual |
|---|---|---|
| Vercel | CLI existente; credencial caducada 2026-09-11. GET de usuario y proyectos con credencial existente devolvió 403; no login ni refresh. | **BLOCKED — credenciales o conexión no disponibles**. Identidad, plan, proyectos y consumo no verificados. |
| Neon | CLI existente; credencial caducada 2026-05-04. GET de usuario con credencial existente devolvió 401; no login ni refresh. | **BLOCKED — credenciales o conexión no disponibles**. Cuenta, plan y consumo desconocidos. |
| Ably | CLI/conexión no disponibles en rutas/variables inspeccionadas. | **BLOCKED — credenciales o conexión no disponibles**. Cuenta, plan y consumo desconocidos. |

### Condiciones públicas verificadas, sin afirmar cuota disponible

- [Vercel Hobby](https://vercel.com/docs/plans/hobby) (actualizado 2026-09-14) y [límites](https://vercel.com/docs/limits) (2026-09-16): uso personal/no comercial; 200 proyectos, 25 por repositorio, 100 deployments/día, un deployment concurrente, build máximo 45 minutos, upload CLI Hobby 100 MB. Bolsas mensuales anunciadas: 1 M invocaciones, 4 CPU-h activas, 360 GB-h memoria provisionada, 100 GB Fast Data Transfer y 10 GB Fast Origin Transfer. Cuatro proyectos no multiplican las bolsas compartidas.
- [Protección](https://vercel.com/docs/deployment-protection) y [uso/precios](https://vercel.com/docs/deployment-protection/usage-and-pricing) (2026-09-15): Hobby ofrece Vercel Authentication, Standard/All Deployments y automation bypass; Password Protection no está en Hobby. Standard deja producción pública: propuesta del spike, **All Deployments** y previews dedicados, sin `--prod`. Interacción entre protección y S2S pendiente.
- [Go Runtime](https://vercel.com/docs/functions/runtimes/go) (2026-08-11): Go beta, Gin, módulo raíz, preset Go con entrypoint main y opción legada `api/*.go` Handler; soporta vendoring y build personalizado. Se conserva el adaptador HTTP previsto. [Fluid Compute](https://vercel.com/docs/fluid-compute) (2026-08-24) no enumera Go: no asignar el límite general Hobby de 300 s a estos servicios sin verificar runtime/configuración efectiva.
- [Neon planes](https://neon.com/docs/introduction/plans) y [precios](https://neon.com/pricing) (sin fecha editorial visible): Free $0 sin tarjeta; 100 proyectos, 10 ramas/proyecto, 100 CU-h/proyecto/mes, hasta 2 CU, 1 GB/proyecto y 20 GB agregados, 5 GB egress/proyecto/mes; suspensión tras cinco minutos inactivo. Exceso de cómputo/egress suspende; storage lleno restringe crecimiento. Son límites publicados, no saldo de cuenta.
- [Neon pooling](https://neon.com/docs/connect/connection-pooling): PgBouncer transaccional para app/serverless y conexión directa para migración/administración. Admisión de 10.000 clientes no equivale a 10.000 consultas simultáneas; documentación presenta 104 conexiones a 0.25 CU, siete reservadas. No depender de estado de sesión; pgx/pool/roles requieren prueba efectiva.
- [Ably Free](https://ably.com/docs/platform/pricing/free) y [límites](https://ably.com/docs/platform/pricing/limits) (sin fecha editorial visible): $0 sin tarjeta para experimentación/PoC, 6 M mensajes/mes, 500 mensajes/s, 200 canales y conexiones; 50 publicaciones/s/canal, 25.000 HTTP/h, 50 HTTP/s. [Conteo](https://ably.com/docs/platform/pricing/message-counting) incluye publicación y entregas; no presupuestar solo publicaciones.
- [Capacidades Ably](https://ably.com/docs/auth/capabilities): `message-subscribe` restringe a mensajes; `subscribe` agrega otras familias de datos. Probar token efímero para canal exacto, sin publish/presencia/comodines. La clave emisora también limita capacidades. Ningún permiso se ha ensayado contra Ably real.

## 6. Matriz previa al despliegue — propuesta condicionada al gate

| Campo obligatorio | Vercel | Neon | Ably |
|---|---|---|---|
| Cuenta y plan, sin datos sensibles | No verificados, **BLOCKED** | No verificados, **BLOCKED** | No conectada, **BLOCKED** |
| Recursos necesarios | Cuatro proyectos dedicados: web/Identity/Game/AI; previews protegidos | Un proyecto Free dedicado; propuesta experimental de dos DB y roles distintos, sin decidir layout de producto | Una app dedicada; canal exacto por partida sintética y claves/capacidades restringidas |
| Coste mensual previsto y condiciones | $0 solo si Hobby real, uso permitido y cuota residual suficiente | $0 solo si Free real y recursos dentro del cupo | $0 solo si Free real y carga dentro del cupo |
| Cuotas disponibles verificadas | **Ninguna**; solo tabla pública | **Ninguna**; solo tabla pública | **Ninguna**; solo tabla pública |
| Riesgo de cargos/upgrades | Rechazar Pro, trials, addons y tarjeta; no operar sobre proyectos de pago existentes | Rechazar Launch/Scale/tarjeta; no reutilizar recursos facturables | Rechazar planes de pago; no contar con tolerancia al exceso |
| Datos sintéticos | Build IDs, IDs/revisiones/payloads de sonda | Identidades y efectos sintéticos, sin datos reales | Avisos ID/revisión/tipo, sin credenciales o estado completo |
| Seguridad y visibilidad | All Deployments + auth propia; bypass solo servidor y header, nunca URL/JS; ninguna producción | TLS, roles separados, denegaciones cruzadas, pool acotado; AI sin DSN | Tokens temporales, `message-subscribe` en recurso exacto; cliente sin publicación oficial |
| Método de creación previsto | Dashboard/API/CLI después del gate y de identificar cuenta/plan/configuración | Console/API después del gate, proyecto exclusivo | Dashboard/Control API después del gate, app exclusiva |
| Método de eliminación previsto | Solo proyectos/deployments/secretos dedicados inventariados, conforme al plan de retirada autorizado | Solo proyecto dedicado, después de capturar evidencia | Revocar credenciales temporales y eliminar app dedicada |
| Evidencias esperadas | Roots/configuración efectivos, builds/runtime/logs sanitizados, cuatro endpoints, negativos, consumo antes/después | Pooled/direct, migraciones, roles negativos, commit/rollback/concurrencia, conexiones y consumo | Recepción permitida, otro canal denegado, publish cliente denegado, fallo/REST y estadísticas |

Estimación experimental: **hasta 100 solicitudes externas y cinco concurrentes como techo total de prueba**, distribuidas entre los casos, sin saturación. Cuatro proyectos pueden consumir cuota compartida por builds/previews/invocaciones; no sumar cuatro veces el cupo. Reintentos, arranque DB, publicaciones y cada entrega también consumen recursos. No hay extrapolación a tráfico mensual real, p95 o coste cero indefinido. Gasto adicional observado en recursos contratados: $0, porque no se contrató ni provisionó ningún recurso facturable externo.

## 7. Problema específico del módulo fuera del Root Directory

[Configuración de builds](https://vercel.com/docs/builds/configure-a-build) (2026-08-28) limita el contexto al root; [FAQ de monorepos](https://vercel.com/docs/monorepos/monorepo-faq) (2026-08-11) documenta **Include source files outside of the Root Directory in the Build Step**. Los roots propuestos son los cuatro subdirectorios del spike. Para Game/AI, `replace ../../modules/engine` queda fuera de su root; probe también queda fuera en los tres servicios.

La prueba local reconstruyó el contexto necesario y compiló sin workspace/caches compartidos. **No demuestra que el builder Go de Vercel preserve esas rutas**. Hay una ruta documental razonable: incluir fuentes exteriores, mantener la topología del checkout, root `go.mod` por servicio y Handler `api/index.go`; configurar cada proyecto por separado y verificar el contexto efectivo.

Después del gate, experimento exacto: registrar cuenta/Hobby/cuota y protección; configurar roots independientes y opción exterior; construir Game/AI sin reutilizar caché de proyecto; capturar `go version`, `GOWORK`, `go list -m`, ruta resuelta y hash de engine; desplegar solo previews y consultar versión autenticada de ambos. Ambos deben reportar el mismo motor y funcionar de forma independiente. Si falla la resolución, guardar error e impacto; no duplicar motor, unir servicios, cambiar stack o pagar para ocultarlo. Identificar herramienta/versiones y si realmente se usa el adaptador legado o preset; no confundir compilación net/http con runtime validado.

## 8. Publicación, comprobaciones y próximos gates

GitHub indicó cero workflows, webhooks, deployments y environments en este repositorio. Sin embargo, consulta de instalaciones GitHub App devolvió 403 y el endpoint de instalación de repositorio exigió credencial de app; la lista Vercel de proyectos quedó en 403. [Monorepos Vercel](https://vercel.com/docs/monorepos) documenta despliegues disparados por push en proyectos conectados. **No se ha descartado una integración externa**.

Los `vercel.json` del experimento deshabilitan Git deployments para sus propios contextos como defensa, sin aplicar nada al proveedor; eso no prueba qué raíz/configuración usa una integración preexistente. Por la condición expresa del propietario, se conserva commit **local**, sin push ni Draft PR hasta acreditar seguridad o recibir autorización que cubra ese riesgo. Título previsto del Draft PR: `spike: verify Go microservices and serverless infrastructure feasibility`. No hay enlace porque **no se publicó**. PR #3 permanece separado y abierto.

Comprobaciones: builds por raíz/contexto limpio; tests locales y Handler exportados; suite SQL real inicialmente bloqueada y posteriormente **PASS**; reinicio normal de Docker actualizado; validación Compose sin secretos; revisión independiente de código/oráculos y límites de claims; inventario/enlaces/fences/JSON de este árbol, límites de rutas Git y escaneo de exposición. La matriz [POC-01–11](TEST_MATRIX.md) distingue PASS/FAIL/BLOCKED/NOT_RUN y LOCAL/REMOTO. Los controles documentales no sustituyen las pruebas ejecutadas ni validan los proveedores remotos.

Bloqueos restantes: credenciales/conexiones Vercel/Neon/Ably y cuotas, autorización remota, contexto Go fuera de root, protección/S2S/cookies efectivos, capacidades Ably y regiones/métricas/consumo remotos. Docker y los casos SQL/idempotencia/REST locales están resueltos según §9. Ninguno es una nueva elección reglamentaria; no se altera TK-N1-1.0.0. Una discrepancia arquitectónica demostrada futura se registrará como propuesta aparte, sin editar ADR aprobadas unilateralmente.

**Recursos externos creados: ninguno. Retirada externa: no aplica.** No se modificaron recursos preexistentes de proveedores externos ni planes/costes. Los cambios posteriores autorizados en Docker son locales. El trabajo queda detenido antes de fase C, lote 4, implementación del juego o cualquier merge.

Para autorizar la prueba externa acotada descrita, el gate requerido es **APRUEBO SPIKE REMOTO**. Esa aprobación no sustituye la identificación de cuentas, cuota residual y seguridad efectiva antes de cada creación.

<a id="docker-repair-20261022"></a>

## 9. Reparación de Docker y validación local — 2026-10-22

El propietario autorizó expresamente los cambios necesarios en Docker para ejecutar las pruebas. Causa observada: Docker Desktop **4.65.0** abortaba al inicializar sockets AF_UNIX inaccesibles en `Docker/run/dockerInference` y después en `docker-secrets-engine/engine.sock`; Windows devolvía **ERROR_CANT_ACCESS_FILE (1920)**. El backend abortaba antes de que el motor Linux estuviera disponible; no era un fallo de PostgreSQL o de las transacciones Go. No se ha demostrado qué originó el primer socket averiado ni corrupción de datos de contenedores.

Primero se preservaron los directorios de runtime bajo nombres de cuarentena y se recrearon vacíos, con Docker detenido y rutas verificadas. Eso permitió un arranque y la primera suite SQL PASS, pero un stop/start controlado reprodujo el fallo. **La limpieza manual no era una solución estable**. [Docker reconoce esta clase de incidencia y señala corrección desde 4.92](https://github.com/docker/desktop-feedback/issues/460); [el seguimiento documenta el mismo patrón](https://github.com/docker/desktop-feedback/issues/554).

El actualizador oficial (`docker desktop update --check-only`, seguido de `docker desktop update --quiet`) ofreció e instaló **4.96.3 (244472)**. La primera invocación de update con el backend caído se quedó esperando; se canceló ese helper y se recuperó el runtime para ejecutar el actualizador. La invocación que descargó/instaló devolvió exit 1 cuando se cerró el CLI durante el relevo al instalador; no se tomó ese código como prueba de éxito. Se confirmó después la versión instalada, log del instalador, firma Authenticode **Valid**, engine disponible y ciclo normal de parada/arranque.

| Verificación posterior | Observación |
|---|---|
| Docker Desktop / CLI / Engine / Compose | **4.96.3.244472 / 29.9.2 / 29.9.2 / 5.6.1** |
| Stop normal | Exit **0**, **13.656 s** |
| Start normal | Exit **0**, **7.333 s**, sin renombrar/borrar sockets entre ambas operaciones |
| Recursos persistentes | Contenedor sintético conservado durante el reinicio; al aplicar el pin de imagen, Compose lo recreó conservando su volumen y el estado `revision=1,effects=1` |
| PostgreSQL | **18.6**, estado **healthy**, `127.0.0.1:15435`, límite 512 MiB/una CPU |
| Imagen fijada | `postgres:18.6@sha256:edb0c08dc33befd0262ab3b5cf1c4f3f8cfd8b09e88c05b45adee29d6761ddb6` |
| Suite real posterior al reinicio | **Exit 0**, cinco casos de integración PASS; dos tests previos de la app también PASS |
| Regresión local Go y web | Cinco raíces Go PASS; `npm test` y `npm run smoke` PASS; los builds previos siguen siendo evidencia de fuentes Go/web sin cambios |

Evidencia: [entorno](evidence/docker-local-environment-20261022.json), [reinicio](evidence/docker-restart-20261022.json), [SQL final](evidence/sql-local-20261022.txt), [regresión Go](evidence/unit-verification.json), [smoke web](apps/web/evidence/web-smoke.json). Los tiempos SQL del último run fueron 0.08/0.03/0.02/0.03/0.03 s por caso; son muestras locales individuales, no percentiles ni capacidad de producción.

Se comprobaron ambos accesos SQL cruzados denegados (42501), up/down de migraciones, pool máximo 2, primer commit, retry seguro, ID con contenido diferente rechazado, carrera sobre misma revisión, rollback antes de commit, pérdida de respuesta después de commit y acceso ajeno/AI denegados. El doble local consultó la DB al recibir aviso para demostrar el orden postcommit; su fallo no invalidó el estado y GET REST recuperó la revisión vigente con avisos perdidos/duplicados/desordenados.

No se ejecutó factory reset, prune, borrado de datos, unregister/shutdown general de WSL ni desactivación de seguridad. No se tocaron imágenes/volúmenes de otros proyectos; los directorios de cuarentena se conservaron. No se instaló un workaround automático de arranque ni se cambiaron manualmente ajustes globales de Docker. Se actualizó únicamente la aplicación con su mecanismo oficial. Un reinicio normal observado no garantiza ausencia de cualquier fallo futuro de Windows/Docker.

El contenedor/red/volumen **locales** del spike quedan disponibles para continuar la revisión; sus credenciales siguen ignoradas. No se crearon recursos en proveedores externos ni se hizo push, PR o merge. POC-03/04/05 pasan en su alcance local; el GO global y la fase remota siguen **BLOCKED** hasta el gate y las evidencias de proveedor requeridas. No se inicia desarrollo del juego ni un lote documental nuevo.

<a id="integral-20261022"></a>

## 10. Adenda de verificación integral — 2026-10-22

Conclusión única del gate local: **GO_LOCAL**. El [informe integral](INTEGRAL_REPORT.md) documenta entorno, comandos, correcciones, DB-01–08, TX-01–07, seguridad y límites; la [matriz vigente](TEST_MATRIX.md) separa los once PASS locales de los once BLOCKED remotos. La [evidencia del gate](evidence/integral-local-gate.json) se registró antes de preparar el plan remoto.

Se repitieron los builds de la misma web y cinco módulos Go; test/vet/build Windows, seis builds Go aislados con cachés nuevas, test/build/smoke Next y race Linux terminaron satisfactoriamente. Race Windows nativo no se ejecutó por falta de compilador; lint no está definido. [Logs y hashes](evidence/integral-builds-summary.md) conservan también los fallos iniciales y su repetición corregida.

La prueba nueva reinició el ejecutable Game (PID 24004→7992) y después PostgreSQL conservando el volumen esperado: revisión/efectos/recibos permanecieron en `3/3/3`, incluidos reintentos de la intención original. La sonda saliente Game→Identity/AI funcionó entre procesos reales; detener AI produjo error controlado y restaurarlo recuperó la operación. El runner separa claves de usuario por audiencia, limita el entorno de cada hijo y comprueba decisiones HTTP, filas y recibos. Sustituye la carencia de S2S saliente y las pruebas de reinicio limitadas a objetos descritas anteriormente.

Quince escenarios de notificación/capacidad pasaron contra Game/PostgreSQL reales con un doble HTTP y cliente Go de prueba. No validan Ably ni la reconciliación de Next. Cierre/reapertura del listener no demuestra reinicio de RAM del broker. HMAC/S2S siguen siendo provisionales; no hay certificación de seguridad de producción, OIDC o navegador. El rol local de app coincide con el propietario sintético; la separación DDL/DML remota es una condición futura explícita.

La [auditoría nueva de cuentas y fuentes oficiales](REMOTE_AUDIT.md) reemplaza §5/§8 para el estado vigente: Vercel no tenía token utilizable y no recibió petición HTTP; Neon respondió 401; Ably no estaba conectado. Las cuotas disponibles y regiones de cuenta siguen sin comprobar. GitHub no mostró workflows/hooks/deployments, pero las instalaciones de Apps y conexiones Vercel no se pudieron descartar. Por esa condición expresa de publicación, los cambios se conservan **solo localmente**, sin push ni Draft PR. Título reservado: `spike: local infrastructure verification and remote feasibility plan`.

El [plan remoto concreto](REMOTE_PLAN.md) propone cuatro proyectos Vercel, dos propiedades PostgreSQL aisladas en un proyecto Neon y una app Ably, nombres dedicados, secretos, criterios POC, presupuesto y retirada. No se ejecutó. Se requiere **APRUEBO SPIKE REMOTO**, más cuentas/planes/cuotas y protección verificables, antes de cualquier escritura externa. No se fusionó PR #3 ni se inició lote 4 o motor completo.
