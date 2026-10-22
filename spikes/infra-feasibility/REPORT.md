# Gate 3.5 — informe de viabilidad previo a recursos externos

**Recomendación: BLOCKED.** Los cuatro builds locales funcionan y comparten una única fuente del sustituto Go. Faltan PostgreSQL local operativo y toda la validación alojada. No se recomienda GO ni NO-GO con esta evidencia. Presupuesto autorizado: **0 USD adicionales**. No se creó, modificó ni desplegó ningún recurso externo.

## 1. Base, alcance y autoridad

Base `main` vigente verificada por GitHub: `aaeeace21860305091c0351bedb0aca0df386816`; rama nueva `spike/infra-feasibility-01`. PR #3 **OPEN**, HEAD `9db0f3e9d4bf1fb642697103d9232bcfaef42205`, sin fusión ni cambio de estados. No existía rama experimental homónima remota/local. Todo el cambio se limita a `spikes/infra-feasibility/`.

Fuentes consultadas: AGENTS, GOALS, PLANS, REG y CONF; 01.01–01.03; ADR-0001/0002/0004/0005/0006; 03.01–03.04 y AD-NORM-001. Los contratos 05.01–05.04 y 07.02 se consultaron en `origin/docs/specs-lote-03` como propuestas **no aprobadas**. La autorización posterior del Gate 3.5 permite el experimento aislado y no cambia decisiones de producto, PRE-09, TK-N1-1.0.0 o los estados documentales anteriores.

Fecha de revisión local 2026-10-21 Caracas; ejecución UTC 2026-10-22. No se consultaron reglas ni motores externos para reconsiderar el perfil normativo.

## 2. Arquitectura realmente probada y resultados

Next independiente y tres módulos Go/Gin independientes, con arranque loopback y entradas `api/index.go` adaptadas a `net/http`. Game y AI importan el mismo `engine v0.1.0`; Identity no lo importa. `engine` solo calcula un SHA-256 determinista, sin reglas del juego. `modules/probe` es biblioteca de soporte, no otro microservicio. No se probó el MVP.

Comandos, herramientas y reproducción: [README](README.md), [script Go](tests/Verify-Local.ps1), [web](apps/web/README.md). `go test -count=1 ./...` se ejecutó por las cinco raíces Go; las entradas HTTP se probaron con Handler real y servidores `httptest`. Se compiló también la suite `-tags=integration -run '^$'` **sin ejecutar pruebas DB**, seguido de intento real de esa suite, exit 1 por DB inaccesible. npm ci/build/test/smoke se ejecutaron desde la raíz web.

| Build local | Contexto/cache | Muestra de tiempo total | Resultado |
|---|---|---:|---|
| Identity | Fuentes de Identity + probe; GOCACHE/GOMODCACHE nuevos | 27.550 s | PASS |
| Game | Fuentes de Game + probe + única fuente engine; caches nuevos | 27.322 s | PASS |
| AI | Fuentes de AI + probe + misma engine; caches nuevos | 26.643 s | PASS |
| Web | Su propio manifiesto/lock y fuentes; `.next` ausente | 11.090 s | PASS |

`GOWORK=off`, `GOENV=off`, `GOTOOLCHAIN=local`, `CGO_ENABLED=0`; builds desde **cada raíz**, no desde un workspace global. Evidencia: [local-verification.json](evidence/local-verification.json), inventarios de módulos y metadatos de ejecutables. Versiones Go 1.26.2/Gin 1.12.0/pgx 5.11.0; web Next 16.4.2/React 19.3.1/TS 5.9.3, Node 24.13.0/npm 11.6.2. La compilación local no prueba la versión soportada ni el comportamiento efectivo del builder remoto.

Hash SHA-256 de `engine.go` en ambos contextos: `BD0D2F40296E32B6C832DEA9B22C0CB1B957E428D2DC01E20456C303AB14C973`. Ambos manifests requieren `v0.1.0` y resuelven `../../modules/engine`; es una versión **local experimental sin release/tag publicado**. Su fuente canónica permanece en una sola ruta. Los contextos de prueba copian la misma fuente para representar builds aislados; no se crean motores divergentes ni copias canónicas en los servicios.

## 3. PostgreSQL, transacciones y avisos: bloqueo observado

Docker CLI 29.2.1 y Compose 5.1.0 existen. Se inició Docker Desktop oculto. El backend falló antes del motor: no pudo acceder/eliminar el socket de inferencia `Docker/run/dockerInference`; el log registra parada de motores a `2026-10-22T02:19:05Z`. No se reseteó Docker/WSL ni se borró ese socket. No hay contenedor, volumen, imagen PostgreSQL descargada ni migración aplicada por este spike. `config --quiet` validó Compose, sin imprimir variables.

La propuesta local tiene dos bases con propietarios distintos, CONNECT público revocado, pools pgx máximos 2 y migraciones up/down. **Aislamiento SQL, pooling efectivo, commit/rollback, idempotencia, carrera, pérdida de respuesta, postcommit y recuperación REST no están demostrados**. Cinco tests reales se detuvieron en el fixture: [salida sanitizada](evidence/sql-blocked.txt), [casos](TEST_MATRIX.md). No se sustituyó PostgreSQL por memoria o SQLite para producir PASS.

El código previsto persiste recibo+estado+efecto en una transacción, serializa por agregado y consulta recibo antes de rechazar revisión vieja; un reintento idéntico puede recuperar un commit cuya respuesta se perdió. Un doble local intenta aviso después de commit y REST vuelve a autorizar la consulta. Esto es implementación experimental **sin verificación de ejecución SQL**. No hay promesa de outbox, entrega exactamente una vez o recuperación mientras REST esté indisponible.

## 4. Seguridad y cancelación observadas

PASS local de los subcasos de credencial ausente/malformada, firma inválida, audiencia, caducidad, duración excesiva, scope, tipo usuario/servicio y principal interno. Los handlers internos solo admiten servicio Game para la audiencia del destino; credenciales usuario no sirven como servicio. Faltan las denegaciones por recurso sustentadas en DB y el flujo saliente real Game→AI/Identity. Los clientes de `httptest` demuestran fronteras HTTP de los handlers, no esa integración completa.

Autenticación HMAC **experimental**, claves separadas por kind compartidas con verificadores de prueba; no JWT asimétrico/JWKS, Google OIDC, sesiones ni seguridad contra un servicio comprometido. Una clave de verificación simétrica no demuestra que AI sea incapaz de falsificar una credencial si se compromete. No hay emisor HTTP ni bypass permisivo. Configuración ausente falla cerrada y los binarios locales escuchan solo en loopback. El remoto necesita revisión de controles y credenciales dedicadas antes de exponerlos.

AI rechaza entorno DB; su test ahora establece una configuración válida 200 antes de comprobar los tres nombres de DSN prohibidos. Trabajo sintético: 1–200 unidades, contexto de petición, deadline HTTP de 40 ms y cancelación. Pruebas observadas de resultado, límite, deadline y cancelación; guardia aislada de revisión obsoleta en Game. No se probó una carrera AI→transacción Game ni capacidad de Negamax.

La web comprobó cabeceras, rutas/query, Authorization y 401 a través de rewrites con dobles Node; ocho peticiones medidas y cuatro intentos de preparación. Los ocho tiempos observados van de **1.44 a 14.93 ms** en loopback; no representan latencia Go, DB, Ably o regiones. Se revisaron HTML/logs/nueve archivos JS contra fixtures. Cookies transportadas manualmente no prueban enforcement en navegador ni autenticación completa. OIDC **NOT_RUN**.

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

Estimación experimental: **hasta 100 solicitudes externas y cinco concurrentes como techo total de prueba**, distribuidas entre los casos, sin saturación. Cuatro proyectos pueden consumir cuota compartida por builds/previews/invocaciones; no sumar cuatro veces el cupo. Reintentos, arranque DB, publicaciones y cada entrega también consumen recursos. No hay extrapolación a tráfico mensual real, p95 o coste cero indefinido. Gasto adicional observado en recursos contratados: $0, porque no se contrató ni provisionó nada.

## 7. Problema específico del módulo fuera del Root Directory

[Configuración de builds](https://vercel.com/docs/builds/configure-a-build) (2026-08-28) limita el contexto al root; [FAQ de monorepos](https://vercel.com/docs/monorepos/monorepo-faq) (2026-08-11) documenta **Include source files outside of the Root Directory in the Build Step**. Los roots propuestos son los cuatro subdirectorios del spike. Para Game/AI, `replace ../../modules/engine` queda fuera de su root; probe también queda fuera en los tres servicios.

La prueba local reconstruyó el contexto necesario y compiló sin workspace/caches compartidos. **No demuestra que el builder Go de Vercel preserve esas rutas**. Hay una ruta documental razonable: incluir fuentes exteriores, mantener la topología del checkout, root `go.mod` por servicio y Handler `api/index.go`; configurar cada proyecto por separado y verificar el contexto efectivo.

Después del gate, experimento exacto: registrar cuenta/Hobby/cuota y protección; configurar roots independientes y opción exterior; construir Game/AI sin reutilizar caché de proyecto; capturar `go version`, `GOWORK`, `go list -m`, ruta resuelta y hash de engine; desplegar solo previews y consultar versión autenticada de ambos. Ambos deben reportar el mismo motor y funcionar de forma independiente. Si falla la resolución, guardar error e impacto; no duplicar motor, unir servicios, cambiar stack o pagar para ocultarlo. Identificar herramienta/versiones y si realmente se usa el adaptador legado o preset; no confundir compilación net/http con runtime validado.

## 8. Publicación, comprobaciones y próximos gates

GitHub indicó cero workflows, webhooks, deployments y environments en este repositorio. Sin embargo, consulta de instalaciones GitHub App devolvió 403 y el endpoint de instalación de repositorio exigió credencial de app; la lista Vercel de proyectos quedó en 403. [Monorepos Vercel](https://vercel.com/docs/monorepos) documenta despliegues disparados por push en proyectos conectados. **No se ha descartado una integración externa**.

Los `vercel.json` del experimento deshabilitan Git deployments para sus propios contextos como defensa, sin aplicar nada al proveedor; eso no prueba qué raíz/configuración usa una integración preexistente. Por la condición expresa del propietario, se conserva commit **local**, sin push ni Draft PR hasta acreditar seguridad o recibir autorización que cubra ese riesgo. Título previsto del Draft PR: `spike: verify Go microservices and serverless infrastructure feasibility`. No hay enlace porque **no se publicó**. PR #3 permanece separado y abierto.

Comprobaciones: builds por raíz/contexto limpio; tests locales y Handler exportados; compilación de suite SQL y bloqueo real; validación Compose sin secretos; revisión independiente de código/oráculos y límites de claims; inventario/enlaces/fences/JSON de este árbol, límites de rutas Git y escaneo de exposición. La matriz [POC-01–11](TEST_MATRIX.md) distingue PASS/FAIL/BLOCKED/NOT_RUN y LOCAL/REMOTO. Los controles documentales no validan los casos DB/remotos.

Bloqueos restantes: Docker Desktop local, credenciales/conexiones Vercel/Neon/Ably y cuotas, autorización remota, contexto Go fuera de root, protección/S2S/cookies efectivos, SQL/idempotencia/REST real, capacidades Ably, regiones/métricas/consumo. Ninguno es una nueva elección reglamentaria; no se altera TK-N1-1.0.0. Una discrepancia arquitectónica demostrada futura se registrará como propuesta aparte, sin editar ADR aprobadas unilateralmente.

**Recursos externos creados: ninguno. Retirada externa: no aplica.** No se modificaron recursos preexistentes ni planes/costes. El trabajo queda detenido antes de fase C, lote 4, implementación del juego o cualquier merge.

Para autorizar la prueba externa acotada descrita, el gate requerido es **APRUEBO SPIKE REMOTO**. Esa aprobación no sustituye la identificación de cuentas, cuota residual y seguridad efectiva antes de cada creación.
