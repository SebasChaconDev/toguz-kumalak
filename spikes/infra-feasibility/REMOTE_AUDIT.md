# Gate 3.5 — Auditoría previa al experimento remoto


> **Estado vigente — 2026-10-25 (Caracas): GO_REMOTE, limitado al spike desechable.** Seis builds remotos READY, ocho intentos acumulados contando dos cancelados históricos; pruebas reales Neon/Ably/OIDC/cookie y seguimiento C6 completados. Todos los recursos propios retirados y recursos anteriores preservados. [Cierre, evidencias y límites](REMOTE_AUDIT.md#remote-close-20261025). Las notas anteriores de bloqueo que siguen son cortes históricos, sustituidos por esta adenda. No acredita producto, aprobación automática de propuestas ni generación del lote 4.

> **Cierre actual: BLOCKED_REMOTE por destino Vercel, recursos propios retirados.** Cuotas y protección gratuita verificadas; techo100 retirado por el propietario. Dos builds fueron cancelados porque Vercel devolvió `production` incluso solicitando staging; cero secretos y alias sin asignar. Los cuatro proyectos y ambos deployments se retiraron: inventarios finales vacíos.41 API directas más navegación UI; sin lote4. [Evidencia y decisión pendiente](REMOTE_AUDIT.md#12-prueba-real-de-protección-y-destino-inicial-bloqueo-específico-recursos-retirados). Notas anteriores conservadas como historia.

> **Estado vigente: cuotas verificadas; techo de100 solicitudes retirado por el propietario.** Se conservan18 lecturas API directas y errores históricos; las lecturas UI posteriores no equivalen a un total HTTP exacto. Sigue el spike remoto previamente autorizado con0USD y máximo6builds. Aún no hay PASS remoto ni permiso de lote4. [Auditoría actual](REMOTE_AUDIT.md#11-decisión-posterior-y-comprobación-de-cuotas--2026-10-22). Las notas siguientes son cortes históricos.

> **Estado vigente: Control API Ably autenticada; preparación `BLOCKED_REMOTE` por presupuesto y cuotas todavía pendientes.** REMOTE-013–016 devolvieron HTTP 200. Se verificaron cuenta, permisos, dos apps preexistentes y acceso a estadísticas. Contador: **16/100**, reserva **10**, disponibles para otras operaciones **74**. [Evidencia actual](evidence/remote-preconditions-budget-20261022.json). Los estados de PAT ausente o 12 solicitudes siguientes son históricos.

> **Última lectura Ably: `BLOCKED_REMOTE`, 12 solicitudes operativas acumuladas.** La clave de aplicación suministrada por el propietario se utilizó una sola vez, en memoria, para GET `/stats?unit=month&limit=1`: **HTTP 401 / Ably 40160**, capacidad insuficiente para leer estadísticas. No hubo retry, publicación, emisión de token o modificación de la app. No se guardó la credencial en archivos nuevos, comandos ni evidencias. Los estados de 11 solicitudes siguientes son el corte anterior.

> **Estado vigente de la continuación: `BLOCKED_REMOTE` por precondiciones y presupuesto pendientes; sesión Neon ya resuelta.** Diez consultas nuevas de cuenta fueron satisfactorias, con **11 solicitudes operativas acumuladas** al incluir el 401 anterior. SQL y documentación están fuera del cupo HTTP por confirmación expresa del propietario. No se creó ningún recurso. [Evidencia actual](evidence/remote-preconditions-budget-20261022.json) y [presupuesto revisado](REMOTE_PLAN.md#8-continuación-tras-renovar-neon-unidad-y-presupuesto-revisados). Los apartados anteriores de caducidad y contador global desconocido son históricos.

> **Adenda vigente — 2026-10-22, 20:27 UTC: `BLOCKED_REMOTE`.** Se recibió la autorización explícita. Antes de cualquier escritura externa, GET Neon `/api/v2/users/me` devolvió **401** con la credencial CLI existente; se activó la parada obligatoria. [Registro](evidence/remote-execution-20261022.json). Los apartados originales siguientes son históricos y no acreditan una revalidación actual de cuentas, cuotas o automatizaciones.

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

## 7. Adenda posterior a la autorización

Después del informe inicial, el preflight comunicado al propietario observó Vercel `sebaschacondev`, equipo UltraSéneca / `ultra-seneca`, Owner y Hobby, sin proyectos/deployments/integraciones; Neon UltraSéneca Free; Ably Free y Owner. GitHub y sus instalaciones permitieron descartar automatización activa dentro del alcance entonces inspeccionado. La preparación se comunicó como `READY_FOR_REMOTE_APPROVAL`. Son antecedentes de esa auditoría, **no nuevas consultas de este intento**. Las credenciales y riesgos desconocidos de los apartados 1 y 5 corresponden a la fase inicial, anterior a aquellas comprobaciones.

Inventario de referencia anterior, pendiente de revalidación antes de crear:

| Proveedor | Recursos/identificadores conocidos | Restricción |
|---|---|---|
| Vercel | Equipo `ultra-seneca`, Hobby; cero proyectos observados anteriormente | No asumir que el inventario continúe vacío; comprobar protección y automatizaciones |
| Neon | Organización `org-royal-mountain-44404235`; **9 Stones**, proyecto `summer-thunder-19865667`; rama `br-solitary-credit-b5lwxosv`; compute `ep-plain-tree-b5zuzsvn` | Conservar íntegros; no reutilizar para el spike |
| Ably | Cuenta Free `Mkf2Mg` / `107133`; **Player chat**, aplicación `rfrH2Q` | Conservar íntegra; no reutilizar credenciales |
| GitHub | `SebasChaconDev/toguz-kumalak`; PR #3 abierto, HEAD observado `9db0f3e9d4bf1fb642697103d9232bcfaef42205` | No modificar ni fusionar; repetir comprobación de automatizaciones antes de publicar |

La única solicitud API nueva a una cuenta fue `REMOTE-001`, GET `https://console.neon.tech/api/v2/users/me`, **20:27:50.143–20:27:50.489 UTC**, HTTP **401**. La caducidad local de la credencial era **2026-10-22T20:15:13.583Z**. Se conservaron solo metadatos sanitizados; no se guardó la respuesta privada ni se imprimieron credenciales. Sin renovación, redirect o retry. No se revalidó la sesión web de Neon ni se concluye que esté cerrada.

| Precondición | Resultado del intento |
|---|---|
| Commit `9409462`, árbol inicialmente limpio y GO_LOCAL intacto | PASS de integridad local; no nueva ejecución de las POC |
| Sesión Neon utilizable por el ejecutor API | BLOCKED: credencial CLI rechazada 401 |
| Nueva comprobación completa de planes, cuotas y sesiones de los tres proveedores | NOT_RUN completa, por parada previa |
| Nueva comprobación de automatizaciones GitHub/Vercel | NOT_RUN; publicación retenida |
| Inventario previo actual y presupuesto global demostrable | NOT_RUN completo; presupuesto no certificado |
| Aprovisionamiento, configuración y pruebas remotas | NOT_RUN |

Ninguna escritura externa, pago, upgrade, push, creación/borrado o emisión de credencial fue realizada. **9 Stones** y **Player chat** no fueron modificados por el agente; su estado externo no se volvió a comprobar después de parar. Ningún coste adicional fue contratado; no hay una nueva conciliación de facturación. No existen recursos propios de este intento que retirar. Resultado **BLOCKED_REMOTE**, sin fallo de arquitectura demostrado. La sesión CLI puede restablecerse mediante el flujo oficial `neonctl auth`, sin compartir tokens en el chat; luego siguen siendo obligatorios los demás prechecks y la contabilidad descrita en [el plan](REMOTE_PLAN.md#7-adenda-de-ejecución-y-contabilidad).

## 8. Revalidación tras el auth del propietario

| Evidencia operativa | Resultado fresco |
|---|---|
| REMOTE-002: Neon `/users/me` | 200, `sebaschacondev`, plan Free; credencial renovada con caducidad 21:34:47.676 UTC |
| REMOTE-003: organización Neon | 200, `org-royal-mountain-44404235`, UltraSéneca, Free |
| REMOTE-004: proyectos Neon | 200, un proyecto: **9 Stones**, `summer-thunder-19865667`, PG18, `aws-us-east-2`; sin colisión con el nombre del spike |
| REMOTE-005: equipos Vercel | 200, UltraSéneca / `ultra-seneca`, `team_xzyyrjATBH1J8bl5U89AiC59`, OWNER, Hobby activo; página final |
| REMOTE-006/007/008: Vercel | 200, **cero proyectos, deployments e integraciones** en el equipo inspeccionado |
| REMOTE-009: consulta GraphQL GitHub | 200, `SebasChaconDev`, ADMIN; main `aaeeace21860305091c0351bedb0aca0df386816`, cuatro ramas, spike ausente, cero deployments/environments; PR #3 OPEN con HEAD intacto |
| REMOTE-010/011: GitHub workflows/hooks | 200, **cero workflows y webhooks** |

El cursor de proyectos Neon no se interpreta por sí solo como otra página: el CLI oficial 2.22.0 termina cuando la respuesta contiene menos elementos que el límite pedido (uno frente a 100); no apareció el campo `unavailable` en la respuesta. Vercel indica `next=null`. GitHub indica `hasNextPage=false`; sus listas vacías no requieren páginas adicionales. No se hicieron peticiones de paginación especulativas.

Se conservan HEAD `9409462` y el blob original `GO_LOCAL`; el árbol tiene las actualizaciones documentales autorizadas, **no está limpio**. No hay cambios de código. La revisión local no halló workflows, hooks Git activos u otros proveedores configurados; los cuatro `vercel.json` deshabilitan despliegue Git. Las lecturas actuales no hallan triggers en workflows/webhooks ni proyectos Vercel conectados. Las instalaciones de GitHub Apps no se revalidaron en este paso; no se declara ausencia universal de automatizaciones. El candidato de archivos inline evita publicar una rama, pero no autoriza un push no auditado.

Ably: se leyó el inventario **ya renderizado** de la sesión del propietario, sin navegar ni recargar: cuenta Free `Mkf2Mg`, una app **Player chat**, `rfrH2Q`. Es observación de UI, **no una consulta actual al servidor ni prueba de cuota residual o Control API**. El propietario confirmó que solo dispone de sesión web, sin token Control API guardado localmente. No se solicitó que publique un token ni se creó uno. No se utilizó una clave de Player chat.

Los planes Free/Hobby están confirmados para Neon/Vercel, pero las **cuotas residuales y consumo actual de los tres proveedores siguen sin comprobarse por completo**. La API de consumo Neon descrita por el cliente oficial instalado se limita a planes superiores; no se hizo una llamada Free no soportada ni se tomó un cero legado como saldo. Cero proyectos Vercel no demuestra consumo cero durante todo el ciclo. No se crean recursos, activan trials ni contratan costes para suplir estas carencias.

El propietario excluyó expresamente documentación y SQL del contador de 100 operativas. Saldo verificado: **11 gastadas, 89 restantes, mínimo 10 retiradas reservadas, máximo 79 para el resto**. Una secuencia completa concreta, después de ahorros API acreditados, suma **108 más operaciones pendientes**; no es un mínimo universal ni un presupuesto aprobado. El [apartado 8 del plan](REMOTE_PLAN.md#8-continuación-tras-renovar-neon-unidad-y-presupuesto-revisados) detalla los costes y la falta de garantía de retirada con diez. Por tanto, la precondición presupuestaria no se declara satisfecha. **BLOCKED_REMOTE; cero escrituras externas, cero builds, POC remotos NOT_RUN.**

## 9. Clave de aplicación proporcionada por el propietario

`REMOTE-012` verifica únicamente lectura de estadísticas de una aplicación mediante Basic Auth sobre TLS. Ably respondió **401 / 40160**. La [tabla oficial](https://ably.com/docs/platform/errors/codes) define 40160 como `capability_denied`; no se interpreta como prueba de clave inexistente, ni se certifican otros permisos o identidad de aplicación con esta respuesta. No se solicitó historial de mensajes, presencia, suscripción o publicación.

El [REST de aplicación](https://ably.com/docs/api/rest-api) admite claves de aplicación en Basic Auth; la [Control API](https://ably.com/docs/platform/account/control-api) exige un token personal Bearer diferente. No se envió deliberadamente la clave de aplicación a la Control API con un esquema incompatible. Para el preflight siguen faltando un PAT Control API utilizable con lectura de apps/estadísticas y las cuotas completas. El secreto se tomó de la entrada del propietario en el historial de este mismo chat, exclusivamente en memoria, para evitar repetirlo en argumentos de herramientas o crear otro archivo de credenciales. El registro contiene solo endpoint, HTTP, código de error y contadores.

Saldo nuevo: **12 consumidas, 88 restantes, 10 reservadas, 78 para otras operaciones**. La propuesta que antes sumaba 108 pasa a **109 más pendientes** al incorporar esta lectura; no es un mínimo universal ni un presupuesto cerrado. Ningún recurso, clave o PAT fue creado, revocado o modificado por el agente. Los POC remotos permanecen BLOCKED/NOT_RUN.

## 10. Control API autenticada con el token posterior del propietario

`REMOTE-013`: GET `/v1/me`, **200**, cuenta `Mkf2Mg` / `Free account`; vigencia devuelta hasta **2026-11-21T21:12:01Z**. El token permite `read:app` y `read:stats`, además de permisos de escritura y otras capacidades que no se utilizaron. El nombre de la cuenta no sustituye verificar el paquete de facturación. No se guardó el token en archivos nuevos, argumentos de comandos ni evidencias; solo se usó en memoria a partir de la entrada del propietario en este mismo chat.

`REMOTE-014`: GET `/v1/accounts/Mkf2Mg/apps`, **200**, dos aplicaciones, sin enlace de página siguiente:

| App preexistente al GET | ID | Estado | TLS |
|---|---|---|---|
| Player chat | `rfrH2Q` | enabled | obligatorio |
| 9 stones | `KLXkxg` | enabled | obligatorio |

Ninguna fue creada por el agente. Se preservan ambas: su presencia no autoriza reutilizarlas, configurarlas ni retirarlas como recursos del spike. No hay colisión con `tk-spike01-realtime`.

`REMOTE-015`: estadísticas de cuenta del mes UTC actual, **200**, pero el recolector no conservó las métricas: PowerShell enumeró automáticamente la respuesta JSON de un solo elemento. Se reprodujo este comportamiento en memoria con un fixture local y se corrigió únicamente la lectura con `-NoEnumerate`. `REMOTE-016` es una **repetición explícita adicional, contada**, del mismo intervalo; **200**, una fila mensual. Los únicos nombres de métricas presentes eran `apiRequests.all.refused`, `apiRequests.other.refused` y `peakRates.apiRequests`; los contadores de mensajes/conexiones/canales no se infieren como cero. No se contabilizó la primera lectura como fallo de Ably ni se ocultó su consumo.

Estas lecturas resuelven el acceso e inventario Control API. No prueban todos los saldos residuales, paquete Free actual ni compatibilidad funcional del experimento. Los POC remotos permanecen BLOCKED/NOT_RUN. Saldo: **16 consumidas, 84 restantes, 10 reservadas, 74 para otras operaciones**. La misma propuesta completa incorpora cuatro lecturas nuevas y pasa de 109 a **113 más pendientes**; sigue sin ser un mínimo universal o un presupuesto aprobado. Cero escrituras externas, credenciales emitidas por el agente, recursos creados, builds, pushes o cambios al PR #3.

## 11. Decisión posterior y comprobación de cuotas — 2026-10-22

El propietario indicó: **«Olvida el presupuesto. No importa. Pero por favor termina de hacer las comprobaciones que necesito para proceder con el desarrollo del lote 4»**. Se retira el techo de **100 solicitudes**; no se presenta una falsa demostración de que la propuesta anterior cabía. Se conserva el historial de las **18 solicitudes API directas** y sus errores. La navegación autenticada posterior incluye peticiones internas del dashboard no instrumentadas individualmente; 18 no se declara total combinado exacto. Permanecen **0 USD**, máximo seis builds, datos sintéticos, protección, retirada y prohibición de modificar PR #3 o iniciar lote 4.

`REMOTE-017`: facturación FOCUS Vercel devolvió **404 / costs_not_found**; no prueba consumo cero. `REMOTE-018`: GET del proyecto Neon protegido devuelve **200**, `free_v3`, 77 CPU-segundos (aprox. 0,02139 CU-h), 304 segundos activo, límite de rama **1.073.741.824 bytes** y contadores de almacenamiento/egress informados como cero. La API advierte retraso; el cero no acredita ausencia física de datos. Periodo **2026-10-22T19:10:16Z–2026-11-01T00:00:00Z**. No se ejecutó SQL ni se modificó 9 Stones. La consulta de proyecto ofrece consumo Free, a diferencia de `consumption_history` de planes superiores; esta precisión sustituye la limitación general anterior.

El navegador renovó su conexión. Tras retirar el techo, se consultaron las páginas de cuenta sin escrituras de configuración:

| Proveedor | Evidencia observada | Conclusión previa al experimento |
|---|---|---|
| Vercel Hobby, UltraSéneca | Usage mensual y hoy: invocaciones 0/1.000.000; duración clásica 0/100 GB-h; Fluid 0/360 GB-h y 0/4 CPU-h; transferencia 0/100 GB y origen 0/10 GB; build clásico 0/100 h | Capacidad para carga mínima disponible; las métricas Fluid no se atribuyen automáticamente a Go |
| Neon Free | Organización e inventario anteriores, GET proyecto con consumo y límite real; una de 100 plazas públicas ocupada | Nuevo proyecto aislado recibe cupos por proyecto; no se consume ni altera el existente |
| Ably Free | Package actual **Free/$0**; Limits actual 0/6.000.000 mensajes, 0/28,6 GiB, 0/200 conexiones y canales, reset en 10 días; 500 mensajes/s, 64 KiB, historia24h | Saldo suficiente para ≤20 publicaciones pequeñas, dos clientes y tres canales; no es medición de máximos históricos |

Vercel mostró ventanas **Sep22 15:00–Oct22 15:00** y **Oct21 21:00–Oct22 20:59**; se conservan las etiquetas de su UI sin asumir que usa America/Caracas. Los valores fueron explícitos del dashboard, no inferidos de su inventario vacío. `Build Minutes 0s/0s` es una métrica facturable distinta del límite clásico observado100h, sin activar máquinas de pago.

Fuentes públicas actuales: [Neon Free](https://neon.com/docs/introduction/plans), [Vercel Hobby](https://vercel.com/docs/plans/hobby), [protección incluida](https://vercel.com/docs/deployment-protection/usage-and-pricing). La cuota observada habilita únicamente las pruebas mínimas aprobadas; no certifica producción ni coste indefinido. La evidencia sanitizada está en el [ledger](evidence/remote-preconditions-budget-20261022.json).

La precondición de cuota se considera satisfecha para esta carga; el límite de solicitudes fue sustituido por decisión posterior. Continúa el spike ya autorizado, sin push: archivos inline, proyectos sin gitRepository, secretos únicamente después de verificar `ssoProtection=all`. Primero configuraciones Vercel; Neon/Ably se crean solo al superar la etapa de empaquetado, evitando recursos innecesarios si falla. Antes de cada operación se registra el intento. Eliminar exclusivamente recursos propios, revocar sus credenciales y confirmar inventarios al terminar. No se necesita publicar la rama para esta ruta; el permiso de push no se amplía.

## 12. Prueba real de protección y destino inicial: bloqueo específico, recursos retirados

Después de retirar el techo100 y comprobar cuotas, se ejecutó la etapa A ya autorizada. `REMOTE-019` confirmó inventario Vercel vacío. Se crearon exclusivamente estas cuatro configuraciones, sin gitRepository, secretos o builds; las respuestas efectivas confirmaron **`ssoProtection.deploymentType=all`**, cero variables y fuentes exteriores activadas:

| Recurso creado exclusivamente por el agente | ID registrado | Alta | Retirada |
|---|---|---|---|
| tk-spike01-web | `prj_nGwijRre5BVDKgkC3qc3MFmM6zGu` | REMOTE-020 | REMOTE-036,204 |
| tk-spike01-identity | `prj_kKzjE2mt9F9SS3iqmKlutYIy3Kw6` | REMOTE-022 | REMOTE-037,204 |
| tk-spike01-game | `prj_d0oDXqAVkCsGVzM48lDoe7wxzo9h` | REMOTE-023 | REMOTE-038,204 |
| tk-spike01-ai | `prj_U7d8JYBpeAiV3y5gf0oBmf7anbEy` | REMOTE-024 | REMOTE-039,204 |

El recolector se detuvo tras REMOTE-020 porque trató el array histórico `resourcesCreated` como número. Además, `@($null).Count` devuelve1 en PowerShell: el contador inicial de variables1 era un artefacto local. Se reprodujeron ambos en memoria; se conservó el POST200, se registró propiedad del recurso y REMOTE-021 verificó por GET cero variables/deployments. No se duplicó su creación ni se imputó el fallo al proveedor.

**Incidente de destino, conservado sin ocultarlo:**

- `REMOTE-025`: POST deployment con18 archivos versionados inline,36.037bytes, raíz Game y módulos exteriores preservados; sin Git, `target` omitido y cero secretos. Se esperaba Preview por el contrato REST, pero la respuesta devolvió **`target=production`**. No se aceptó como Preview ni como POC PASS. Cancelación REMOTE-026200: `CANCELED`, `aliasAssigned=false`; metadatos REMOTE-027200. Logs REMOTE-028200:54 eventos, builderCLI65.1.0, iad1, máquina2cores/8GB, sin caché, descargaGo1.26.0 y dependencias. No hay compilación terminada ni operación funcional demostrada. Borrado REMOTE-029200 y ausencia REMOTE-030404.
- `REMOTE-031`: después de investigar el contrato, se envió **`target=staging` explícito**, opción no productiva documentada, con idénticos archivos y cero secretos. Vercel volvió a devolver **`target=production`**. Cancelación inmediata REMOTE-032200 (`CANCELED`, alias no asignado), logs REMOTE-033200 (54 eventos), borrado REMOTE-034200 y ausencia REMOTE-035404. No habrá más intentos para tantear parámetros sin resolver esta restricción.

La [documentación REST](https://vercel.com/docs/rest-api/deployments/create-a-new-deployment) dice que omitir target genera Preview y reconoce staging; la [documentación de dominios](https://vercel.com/docs/domains/working-with-domains/deploying-and-redirecting) y el [anuncio oficial](https://vercel.com/blog/default-production-domain) indican que el primer deployment se promueve automáticamente a producción. Los dos resultados de cuenta verifican esta excepción en el contexto ensayado. No se concluye que todo Preview o Go en Vercel sea inviable. Tampoco basta `--skip-domain`: la documentación lo presenta para producción sin alias; no demuestra un target no productivo.

**Retirada observada:** cuatro DELETE204, GET de inventario REMOTE-040200 con **cero proyectos** y página final; GET deployments REMOTE-041200 con **cero deployments** y página final. Ambos deployments tienen GET404 individual. Ningún secreto, token de aplicación, DSN o dato sintético se creó/configuró. Los únicos recursos nuevos fueron las cuatro configuraciones Vercel y los dos deployments cancelados. No se creó nada en Neon/Ably ni se emitieron credenciales. No se modificó ninguno de los recursos previos ni el PR#3, ni se hizo push.

**Cierre actual:** **41 solicitudes API directas acumuladas**, más navegación de dashboards posterior a retirar el techo100, sin pretender un total HTTP exacto. **Dos intentos de build cancelados**, ambos cuentan en el máximo6. Cuotas y protección gratuita están verificadas; las POC dependientes permanecen **BLOCKED/NOT_RUN**, POC01 tiene intento remoto **CANCELED**. El objeto de GO_LOCAL y HEAD9409462 siguen intactos.

**Resultado: BLOCKED_REMOTE por restricción de destino**, no por presupuesto100, autenticación Ably, Docker ni una incompatibilidad Go demostrada. Sigue prohibido iniciar lote4. La decisión pendiente es permitir exclusivamente el primer destino production automático en los cuatro proyectos desechables protegidos, sin producto/dominios adicionales/datos reales, o conservar la restricción actual. La secuencia completa original de seis builds útiles requeriría **ocho acumulados** incluyendo los dos cancelados; no se reinicia el contador ni se eleva ese techo unilateralmente. [Propuesta concreta](REMOTE_PLAN.md#10-excepción-operativa-propuesta-no-aprobada).
<!-- El cierre vigente se agrega al final; los cortes históricos se conservan. -->
<a id="remote-close-20261025"></a>

## 13. Cierre remoto — 2026-10-25 Caracas / 2026-10-26 UTC

**GO_REMOTE acotado al spike desechable.** El propietario autorizó: «Autorizo los pasos y requerimientos que sean necesarios para poder empezar a desarrollar lote 4. Aplica lo necesario y cuando todo esté listo para lote 4 avisa». Se aplicaron la excepción de bootstrap production exclusivamente a cuatro proyectos protegidos y ocho builds acumulados, incluidos dos cancelados históricos. No hubo --prod, producto, pago, upgrade, trial, addon ni dominio adicional. Las repeticiones tienen target=null/Preview y no promovieron los alias iniciales.

El propietario había retirado100 HTTP. El ledger continúa acumulado. **El objetivo técnico SQL80 del plan se superó**:105 operaciones reconstruidas conservadoramente en fase principal y12 instrumentadas en seguimiento =117; no es una traza exacta. No se ocultan las27 DDL/semillas ni se atribuye al propietario una excepción SQL literal. Se registra la desviación al completar su autorización amplia, manteniendo0USD, carga acotada y datos sintéticos; no se afirma cumplimiento de80.

### Pruebas observadas y reproducción

[Ledger](evidence/remote-preconditions-budget-20261022.json): método, endpoint, hora, resultado y errores sanitizados. [Manifiesto42 archivos](evidence/remote-source-manifest-20261025.json) y [equivalencia](evidence/remote-source-equivalence-20261025.json): mismos bytes/motor único. Hashtable produjo fingerprints distintos por orden de propiedades; ambos se reproducen exactamente desde las mismas fuentes. Operador corregido a serialización ordenada, sin otro build.

Logs: [Identity](evidence/remote-build-identity-20261025.txt), [Game](evidence/remote-build-game-20261025.txt), [AI](evidence/remote-build-ai-20261025.txt), [Web](evidence/remote-build-web-20261025.txt), [Game repetido](evidence/remote-build-game-repeat-20261025.txt), [AI repetido](evidence/remote-build-ai-repeat-20261025.txt). Seis READY sin caché, roots independientes y fuentes compartidas exteriores. Go remoto1.26.0, local1.26.2; Vercel CLI65.1.0. Game/AI usan spike-engine/v0.1.0 y el mismo hash de fuentes; Identity no importa el motor.

[Ejecución funcional](evidence/remote-functional-20261025.json):70 peticiones HTTP instrumentadas, fallos históricos conservados y repeticiones corregidas:
- Neon PostgreSQL18 real, dos bases y app roles separados, CONNECT cruzado y DDL42501. [Setup27 DDL/semillas](evidence/remote-neon-setup-20261025.json), up/down/up; pgx pool2 por instancia, CONNECTION LIMIT4 por rol, dos backends observados. AI sin DSN ni clave Game.
- Commit/receipt antes de stale; intención distinta409; carrera de dos200/409; precommit503 conserva revisión2. Fallo Ably401 conserva commit. Estado final SQL/REST4 revisiones/4 efectos/4 receipts; nueva instancia recupera receipt1 y estado vigente4.
- Game confirma/relee SQL antes de publicar Ably201. Entrega real; cliente sintético conserva3 frente a avisos2/1/2 y recupera4 vía REST tras desconexión. Sin exactly-once ni outbox obligatorio.
- Claves separadas publish/emisión, canal exacto. Control API rechazó message-subscribe en la clave; emisor privado usa subscribe y JWT cliente restringe message-subscribe/TTL300s. Entrega y denegaciones de publish, otro canal, expiración, firma, missing, suscripción del publicador, presence e history verificadas. El emisor no llega al navegador.
- All Deployments permanece activo. Trusted Sources/OIDC: Web→Identity/Game, Game→Identity/AI. Sin protección401/redirección; sin aplicación401; ajeno403. OIDC inválido503 con upstream401 y recuperación204. No se desactivó protección ni se compartieron bypass secrets.
- [Cookie Chrome](evidence/remote-browser-cookie-20261025.json): navegación real protegida, __Host- Secure/HttpOnly/SameSite=Lax, Identity recibe por rewrite y JS no lee. HTML/seis chunks remotos sin claves.
- AI sintética válida200, work201→422, deadline40ms→408 y base2 contra revisión4→409; no es Negamax ni fuerza medida.

SSE puede devolver HTTP200 y luego evento error40160. La primera expectativa HTTP se conserva como fallo del colector; se corrigió leyendo el evento. pgx CONNECT devuelve42501 envuelto: se corrigió errors.As, sin transformar errores de red en denegaciones SQL.

### C6, consumo y límites

La revisión independiente detectó cronómetros SQL/Ably ausentes. [Seguimiento C6](evidence/remote-latency-followup-20261025.json) los aporta: conexión directa577,476ms; adquisición pooled519,170ms; primera SELECT144,507ms; SELECT caliente71,654ms; BEGIN/UPDATE/COMMIT287,993ms; publish ACK258,356ms; publicación→recepción258,356ms y commit→recepción621,974ms. Revisión SQL1 confirmada antes de publicar, una publicación93bytes, un suscriptor y12 SQL instrumentadas.

Son muestras únicas del operador local hacia proveedores, no latencias interiores de Game/Vercel ni p95/SLO. Región observada iad1/aws-us-east-1; procesos nuevos tras builds repetidos, sin garantía de ausencia de precalentamiento. No se probó suspensión/reactivación Neon ni se midieron clientes/colas PgBouncer. El fallo S2S es rechazo OIDC, no corte físico; la pérdida de respuesta es descarte del harness, no caída TCP instrumentada.

[Uso/cuotas](evidence/remote-usage-observation-20261025.json): Hobby/Free, cero upgrades; Vercel mostraba67 invocaciones/1M,656,64kB/100GB,10min/100h de build y1s/4h CPU Fluid, con redondeo/retardo. Ably principal siete publicaciones/seis entregas =13 mensajes/585bytes y pico1 conexión/canal; seguimiento añade una publicación y entrega. Neon API seguía mostrando0 pese a SQL real: no equivale a consumo nulo. Coste adicional contratado0USD; sin garantía de coste productivo indefinido.

Las78 operaciones funcionales principales son reconstrucción desde rutas/resultados: diez SQL directas del operador y68 de servicios, incluyendo dos pings explícitos. Con setup27 suman105. Health checks pgxpool usan consulta comentario -- ping; sus roundtrips no están medidos. Seguimiento añade12 instrumentadas. HTTP UI/Git/procesos internos no tiene traza exhaustiva: no se presenta suma global exacta.

### Retirada y revisión

Cuatro proyectos Vercel/seis deployments READY retirados. Primer conjunto Neon cold-king-37625726 y Ably MEQwvA; seguimiento de reemplazo Neon winter-wind-87809024 y Ably ujZ-YA después de retirar el anterior. Nunca más de un Neon y un Ably propios activos simultáneamente; ningún build adicional para métricas.

Ocho claves de aplicación revocadas, incluidas cuatro predeterminadas amplias. Ably mantiene registros status1=revocado; REMOTE-118 conserva etiqueta incorrecta activeKeys4 y REMOTE-119/120 la corrigen según [schema oficial](https://ably.com/docs/open-specs/control-v1.yaml). No quedaron claves habilitadas.

Borrar proyectos Vercel dejó seis registros de deployments; DELETE explícitos127–132 y lectura133 verificaron cero. Lecturas finales151/152 preservan solo Neon9 Stones summer-thunder-19865667, Player chat rfrH2Q y9 stones KLXkxg. Cero escrituras sobre ellos. Las sesiones del propietario se reutilizaron, no son credenciales creadas por el spike.

remote_evidence_review reprodujo42 hashes, auditó código/resultados, detectó C6 incompleto y revisó el seguimiento: dictamen posterior GO_REMOTE acotado, sin Critical/Important pendientes para preparar documentos04/08. lot3_readiness_review leyó las once specs y halló dos precisiones menores05.02: código K01- y rango F-HTTP-01–16. Su corrección e integración se registran en PLANS del cierre documental.

Referencias: [OIDC](https://vercel.com/docs/oidc), [Trusted Sources](https://vercel.com/docs/deployment-protection/methods-to-bypass-deployment-protection/trusted-sources), [JWT Ably](https://ably.com/docs/auth/token/jwt), [SSE](https://ably.com/docs/api/sse). PRE-09,246 PRE y TK-N1-1.0.0 permanecen intactos. No hay pruebas de producto ni lote4 generado.
