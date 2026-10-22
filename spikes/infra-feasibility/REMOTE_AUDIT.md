# Gate 3.5 — Auditoría previa al experimento remoto

Fecha de consulta: **2026-10-22, America/Caracas (UTC−04:00)**. Naturaleza: documentación oficial y lecturas de cuenta; ninguna creación, despliegue, cambio de configuración, login ni renovación de autenticación. Base local inspeccionada: `cf5d3973adc977f0ec09c738cbe6cfc078883593`, rama `spike/infra-feasibility-01`, con cambios locales de la validación integral. Esta auditoría comienza después de [GO_LOCAL](evidence/integral-local-gate.json), sin convertirlo en GO remoto.

**Resultado: preparación documental disponible; verificación de cuentas Vercel/Neon/Ably bloqueada. Publicación: `BLOCKED_AUTO_DEPLOY_STATUS_UNKNOWN`.** El siguiente experimento requiere la autorización explícita `APRUEBO SPIKE REMOTO` y resolver sus precondiciones. Este documento no autoriza operaciones remotas.

## 1. Disponibilidad real de cuentas

Las observaciones reproducibles y sus endpoints están en [integral-provider-audit.json](evidence/integral-provider-audit.json). Se guardan metadatos y códigos HTTP; no tokens, respuestas privadas, DSN ni contraseñas. No se invocaron comandos que puedan iniciar login/refresh automáticamente. La búsqueda de herramientas conectadas no encontró un conector Vercel, Neon o Ably.

| Proveedor | Observación fresca | Consecuencia |
|---|---|---|
| Vercel | CLI presente. El archivo de autenticación existente tiene 3 bytes y es un objeto sin claves: no contiene token utilizable. No se hizo petición autenticada a Vercel. Ninguna variable de entorno del proveedor encontrada. | **BLOCKED_NO_USABLE_EXISTING_CREDENTIAL**. Identidad, equipo, plan, pago/trial, saldo de cuotas, proyectos, raíces, Git conectado, protección y regiones de cuenta desconocidos. |
| Neon | CLI y credencial existentes. Caducidad almacenada: `2026-05-04T22:00:44.634Z`. GET `https://console.neon.tech/api/v2/users/me`, con el token existente y sin refresh: **401**. | **BLOCKED_EXISTING_CREDENTIAL_REJECTED**. Plan, uso, proyectos, roles, límites y regiones realmente disponibles desconocidos. |
| Ably | `Get-Command ably` sin resultado; ninguna variable Ably ni conector encontrado en esta sesión. No se realizó consulta de cuenta. | **BLOCKED_NO_CONNECTED_SESSION_FOUND**. Cuenta, plan, aplicaciones, claves/capacidades, consumo y saldo desconocidos. No se concluye que el propietario carezca de cuenta. |
| GitHub | GET `/user`: **200**, usuario `SebasChaconDev`. Repositorio público `SebasChaconDev/toguz-kumalak`, permiso admin, default `main`. El campo de plan no aparece en la respuesta utilizada. | Identidad y acceso al repo verificados. Plan de cuenta, facturación y cuotas de almacenamiento disponibles **no verificados**. |

El informe anterior registraba una credencial Vercel caducada y respuestas 403. Ese dato es histórico: el archivo actualmente inspeccionado está vacío y aquellas solicitudes no se presentan como repetidas hoy. Las cuotas publicadas a continuación tampoco son cuotas residuales de estas cuentas.

## 2. Condiciones oficiales consultadas

Cada enlace se abrió directamente el **2026-10-22**; la columna editorial transcribe la fecha mostrada por la fuente, cuando existe. «Sin fecha visible» no significa que el contenido se haya publicado hoy. Los valores son condiciones públicas, sujetas a confirmación de la cuenta antes de ejecutar el gate.

| Fuente oficial | Fecha editorial visible | Condición pertinente |
|---|---|---|
| [Vercel Go Runtime](https://vercel.com/docs/functions/runtimes/go) | 2026-08-11 | Go está en beta en todos los planes. Admite preset Go con `go.mod` raíz y servidor que escuche `PORT`; también conserva `/api/*.go` con `http.HandlerFunc` exportado. |
| [Vercel Hobby](https://vercel.com/docs/plans/hobby) | 2026-09-14 | Uso personal/no comercial. Bolsas anunciadas: 1 M invocaciones, 4 CPU-h, 360 GB-h memoria, 100 GB transferencia y 10 GB origen. El exceso generalmente pausa uso; no activar Pro/trial/add-ons. |
| [Límites Vercel](https://vercel.com/docs/limits) | 2026-09-16 | Hobby: 200 proyectos, 25 proyectos conectados por repo, 100 deployments/día, uno concurrente, build máximo 45 minutos y upload estático/CLI de 100 MB. Son techos, no cuatro plazas libres comprobadas. |
| [Monorepos Vercel](https://vercel.com/docs/monorepos) y [FAQ](https://vercel.com/docs/monorepos/monorepo-faq) | 2026-08-11 ambas | Un proyecto por directorio; un push puede disparar despliegues de los proyectos conectados. Existe opción para incluir fuentes fuera del Root Directory en el build. |
| [Límites de Functions](https://vercel.com/docs/functions/limitations) | 2026-08-24 | La tabla de duración con Fluid identifica Node.js/Bun/Python. No se extrapola su duración de 300 s como prueba del adaptador Go del spike; comprobar runtime y límites efectivos del deployment. |
| [Regiones Vercel](https://vercel.com/docs/functions/configuring-functions/region) | 2026-08-11 | Hobby usa una sola región por función; `iad1` es el valor por defecto documentado. Región efectiva de los cuatro proyectos pendiente. |
| [Deployment Protection](https://vercel.com/docs/deployment-protection) y [precios](https://vercel.com/docs/deployment-protection/usage-and-pricing) | 2026-09-15 ambas | Vercel Authentication, All Deployments y bypass para automatización incluidos en Hobby. Standard excluye dominios de producción. Password Protection no disponible en Hobby. |
| [Configuración Git Vercel](https://vercel.com/docs/project-configuration/git-configuration) | 2026-08-25 | `git.deploymentEnabled: false` desactiva despliegues Git automáticos del proyecto que lee esa configuración. No prueba la configuración de otros proyectos desconocidos. |
| [Neon planes](https://neon.com/docs/introduction/plans) | Sin fecha visible | Free: $0; 100 proyectos, 10 ramas/proyecto, 100 CU-h/proyecto/mes, hasta 2 CU, 1 GB/proyecto y 20 GB totales de Postgres, 5 GB egress/proyecto/mes; suspensión tras 5 min inactivo. Exceso de CU/egress suspende; storage lleno rechaza operaciones que incrementen almacenamiento. |
| [Neon pooling](https://neon.com/docs/connect/connection-pooling) | Sin fecha visible | PgBouncer transaccional, hostname `-pooler` para app/serverless y conexión directa para migración/administración. 10.000 clientes admitidos no equivalen a igual número de consultas simultáneas. Evitar dependencia de estado de sesión. |
| [Neon regiones](https://neon.com/docs/introduction/regions) | Página sin fecha general; lista NAT: 2026-10-07 | Postgres publicado en AWS, incluido `aws-us-east-1`. Azure está deprecado para nuevos proyectos. La selección disponible en esta cuenta no pudo consultarse. |
| [Ably Free](https://ably.com/docs/platform/pricing/free) | Sin fecha visible | PoC/experimentación, $0 sin tarjeta; 6 M mensajes/mes, 500 mensajes/s, 200 canales y conexiones concurrentes. |
| [Ably límites](https://ably.com/docs/platform/pricing/limits) y [conteo](https://ably.com/docs/platform/pricing/message-counting) | Sin fecha visible | Publicación y cada entrega consumen mensajes; también cuentan persistencia, rewind/historia e integraciones según operación. Los límites no sustituyen medir consumo ni demostrar saldo de cuenta. |
| [Ably Token Auth](https://ably.com/docs/auth/token) y [Capabilities](https://ably.com/docs/auth/capabilities) | Sin fecha visible | Tokens breves emitidos por servidor tras autorizar al cliente. `message-subscribe` limita recepción de mensajes; `subscribe` agrega presencia, objetos y anotaciones. La clave emisora limita los permisos concedibles. |
| [Ably Free y regiones](https://ably.com/docs/platform/pricing/enterprise) | Sin fecha visible | Restricción regional del tráfico/datos es característica Enterprise; no presupuestarla en Free. El PoC utilizaría Pub/Sub global y exclusivamente datos sintéticos. |
| [GitHub Actions billing](https://docs.github.com/en/billing/concepts/product-billing/github-actions) | Sin fecha visible | Minutos de runners estándar son gratuitos para repos públicos; runners grandes son facturables. Artifacts/Packages comparten almacenamiento y el cache tiene cuota propia. No se ha comprobado el presupuesto de esta cuenta. |
| [GitHub ramas protegidas](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/about-protected-branches) | Sin fecha visible | Protección disponible en repos públicos con Free. Puede exigir PR/checks y restringir force push; requiere configurar su aplicación a administradores si se quiere evitar bypass administrativo. |

Cuatro proyectos son aritméticamente inferiores a los techos públicos de 200 y 25; **no está demostrado que la cuenta concreta tenga cuatro plazas, cuota suficiente ni plan Hobby activo**. No se multiplican por cuatro las bolsas compartidas del equipo. Antes de crear nada: verificar plan Free/Hobby sin trial/pago, consumo inicial, límites, proyectos ya presentes y saldo suficiente para el presupuesto completo del [plan remoto](REMOTE_PLAN.md). Si falta ese dato, el caso permanece BLOCKED.

## 3. Empaquetado Go y cuatro raíces

Propuesta para el siguiente gate, todavía **NOT_RUN remoto**:

| Proyecto propuesto | Root Directory real del spike | Dependencias compartidas que deben estar presentes en build |
|---|---|---|
| `tk-spike01-web` | `spikes/infra-feasibility/apps/web` | Lockfile y fuentes Next.js existentes; destinos de rewrites de integración. |
| `tk-spike01-identity` | `spikes/infra-feasibility/services/identity` | `go.mod`, `go.sum`, `api/index.go`, aplicación Gin y `../../modules/probe`. |
| `tk-spike01-game` | `spikes/infra-feasibility/services/game` | Lo anterior más `../../modules/engine`; persistencia pgx real. |
| `tk-spike01-ai` | `spikes/infra-feasibility/services/ai` | `probe` y `engine`, sin credencial PostgreSQL. |

Cada servicio declara `go 1.26.0` y usa `replace` relativo hacia los módulos indicados. Los archivos exteriores tienen que conservar su topología al configurar cada Root Directory. La opción de incluir fuentes exteriores es una **ruta documental candidata**, no una solución ensayada con el builder Go. La prueba local aislada y `GOWORK=off` no establecen qué fuentes subirá o conservará Vercel.

Mantener inicialmente el `api/index.go` existente que exporta Handler y reutiliza Gin. La documentación también describe un preset de servidor Go; los `cmd/server` locales usan listener loopback y puerto de desarrollo, de modo que su build local no demuestra compatibilidad del preset `PORT`. No adoptar ese camino ni consolidar los tres servicios en un proyecto solo para obtener un build favorable.

Para aceptar POC-01/02 remotos se necesitan cuatro deployments separados, logs de build con versiones/módulos y SHA revisado, raíz y configuración verificadas, y llamadas funcionales a cada servicio. Game y AI deben ejecutar el mismo módulo engine compartido y devolver la versión/digest esperado. Identity debe usar su aplicación/DB y probe propios. Un healthcheck o un binario monolítico no cumple el criterio. El engine experimental sigue siendo un digest determinista, no el motor completo TK-N1-1.0.0.

## 4. Seguridad, regiones y coste: condiciones para el gate

- **Vercel:** cuatro proyectos dedicados sin Git automático, previews manuales y All Deployments con Vercel Authentication. Verificar la protección antes de introducir secretos. Probar bypass de automatización exclusivamente en el ejecutor o servidor autorizado: nunca en JavaScript público, URL o logs. Probar web→API y S2S bajo esa protección; no declararlo compatible de antemano. Mantener autenticación de aplicación además de la del deployment. No usar `--prod`, dominios de producto, add-ons ni upgrade.
- **Neon:** dos propiedades dedicadas de datos y roles separados; pooled para Identity/Game, direct y rol migrador para DDL. Ninguna credencial DB en AI/web. TLS, permisos mínimos y aislamiento negativo comprobados en el proveedor real. El candidato geográfico es Vercel `iad1` y Neon `aws-us-east-1`, sujeto a disponibilidad y mediciones. No se garantiza latencia por proximidad nominal ni red privada/IP Allow en Free.
- **Ably:** una app dedicada Free; clave de publicación oficial solo en Game; emisor limitado y separado del cliente. Capability propuesta para cliente: `{"spike:game:<id-sintetico>":["message-subscribe"]}`, TTL propuesto ≤300 s, sin comodines/publish/presence/objects/annotations. El nombre de canal es un fixture, no un ID real. Probar token permitido, canal ajeno, expirado, ausente e inválido y publicación del cliente rechazada. `message-subscribe` puede permitir rewind reciente aun sin `history`: los avisos deben contener solo identificador/revisión sintéticos autorizados. La simulación local no valida estos permisos en Ably.
- **GitHub:** si se propone CI, usar únicamente runner estándar y retención mínima dentro de saldo verificado, sin credenciales de proveedores ni paso deploy. No se creó workflow, protección o presupuesto. La ausencia actual de protección de `main` es un riesgo observado; no se cambió la configuración del repositorio.

Estas son condiciones propuestas, no configuraciones ya aplicadas. El presupuesto operativo y los PASS por POC pertenecen al plan remoto. Coste objetivo $0 requiere comprobar la cuenta, limitar carga/tiempo y detenerse antes del presupuesto; una etiqueta Free pública no acredita consumo cero medido ni bloqueos económicos efectivos en una cuenta desconocida.

## 5. Riesgo de publicación

Quince GET de GitHub frescos confirmaron **0** workflows, runs, artifacts, webhooks, deployments, environments y rulesets. `main` figura `protected: false`; GET de su protección devuelve **404 — Branch not protected**. Actions está habilitado, permite todas las acciones y usa permisos de workflow `read`; no puede aprobar reviews. La rama del spike no aparece entre las cuatro ramas remotas obtenidas.

Sin embargo, GET `/user/installations` devuelve **403** por requerir autorización de GitHub App; GET de instalación del repo devuelve **401** con esta credencial OAuth. No es evidencia de «cero apps». Vercel no dispone de credencial utilizable para consultar proyectos y conexiones Git. No se ha descartado una aplicación externa que despliegue por push/PR.

Los cuatro `vercel.json` locales contienen `git.deploymentEnabled: false`, control documentado para sus proyectos. No demuestran que una integración desconocida use una de esas raíces ni esa configuración; tampoco excluyen otro proveedor. La documentación de monorepos confirma que los proyectos conectados pueden construir tras un push. Por ello no se justifica afirmar SAFE_PUSH ni publicar aún la rama o el Draft PR. Conservar los resultados localmente hasta poder verificar de solo lectura las integraciones y sus triggers, o recibir una instrucción explícita del propietario que resuelva este riesgo concreto. Esto es el gate de publicación solicitado por el propietario, no un rechazo automático de aprobación.

## 6. Estado que debe conservar el informe final

`GO_LOCAL` acredita únicamente la base desechable verificada. Go en Vercel con reemplazos exteriores, PostgreSQL Neon/pooler/roles, tokens Ably, protección/cookies/S2S efectivos, regiones, latencias, consumo y saldo permanecen **NOT_RUN/BLOCKED remotos**. No se ha creado ni desplegado recurso externo. Para ejecutar el plan se necesita **`APRUEBO SPIKE REMOTO`**, resolver sesiones/cuotas y comprobar primero las protecciones; aprobar el plan no convierte los casos remotos en PASS ni autoriza servicios de pago.
