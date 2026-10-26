# Gate 3.5 — plan de validación remota controlada


> **Estado vigente — 2026-10-25 (Caracas): GO_REMOTE, limitado al spike desechable.** Seis builds remotos READY, ocho intentos acumulados contando dos cancelados históricos; pruebas reales Neon/Ably/OIDC/cookie y seguimiento C6 completados. Todos los recursos propios retirados y recursos anteriores preservados. [Cierre, evidencias y límites](REMOTE_AUDIT.md#remote-close-20261025). Las notas anteriores de bloqueo que siguen son cortes históricos, sustituidos por esta adenda. No acredita producto, aprobación automática de propuestas ni generación del lote 4.

> **Cierre actual: BLOCKED_REMOTE por destino Vercel, recursos propios retirados.** Cuotas y protección gratuita verificadas; techo100 retirado por el propietario. Dos builds fueron cancelados porque Vercel devolvió `production` incluso solicitando staging; cero secretos y alias sin asignar. Los cuatro proyectos y ambos deployments se retiraron: inventarios finales vacíos.41 API directas más navegación UI; sin lote4. [Evidencia y decisión pendiente](REMOTE_AUDIT.md#12-prueba-real-de-protección-y-destino-inicial-bloqueo-específico-recursos-retirados). Notas anteriores conservadas como historia.

> **Estado vigente: cuotas verificadas; techo de100 solicitudes retirado por el propietario.** Se conservan18 lecturas API directas y errores históricos; las lecturas UI posteriores no equivalen a un total HTTP exacto. Sigue el spike remoto previamente autorizado con0USD y máximo6builds. Aún no hay PASS remoto ni permiso de lote4. [Auditoría actual](REMOTE_AUDIT.md#11-decisión-posterior-y-comprobación-de-cuotas--2026-10-22). Las notas siguientes son cortes históricos.

> **Actualización vigente: acceso Control API Ably resuelto; presupuesto completo sin validar.** Cuenta/aplicaciones/estadísticas respondieron 200. El contador es **16 consumidas + 74 para otras operaciones + 10 reservadas = 100**. La propuesta completa anterior incorpora estas cuatro lecturas y suma **113 más pendientes**. No se crean recursos con ese presupuesto sin cerrar. [Evidencia](evidence/remote-preconditions-budget-20261022.json), [detalle Ably](REMOTE_AUDIT.md#10-control-api-autenticada-con-el-token-posterior-del-propietario). Los cortes de 12/109 y PAT ausente siguientes son históricos.

> **Actualización Ably: `BLOCKED_REMOTE`, contador 12/100.** La única lectura nueva de estadísticas fue rechazada con HTTP 401 / código 40160 por capacidad insuficiente. Se mantienen **10 reservadas y 78 disponibles para otras operaciones**. El PAT Control API y el presupuesto completo siguen pendientes; la propuesta detallada incorpora la lectura nueva y suma ahora **109 más pendientes**. [Evidencia](evidence/remote-preconditions-budget-20261022.json), [interpretación](REMOTE_AUDIT.md#9-clave-de-aplicación-proporcionada-por-el-propietario). Los cortes de 11/108 siguientes son históricos.

> **Estado vigente de la continuación: `BLOCKED_REMOTE`, antes de crear recursos.** Neon renovado responde 200; ya no bloquea su sesión CLI. El propietario fijó el cupo en **100 solicitudes operativas acumuladas**, incluido el 401 anterior; SQL y documentación se cuentan aparte. Hay **11 consumidas, 89 restantes y 10 reservadas para retirada**: quedan como máximo **79** para otras operaciones. La propuesta completa contabilizada suma **108 más operaciones pendientes**, por lo que **no se valida como presupuesto ejecutable de 100**. [Evidencia de la continuación](evidence/remote-preconditions-budget-20261022.json). Esta nota y el apartado 8 prevalecen sobre los estados históricos siguientes.

> **Adenda vigente — autorización recibida; intento `BLOCKED_REMOTE`, 2026-10-22, 20:27 UTC.** Neon respondió 401 en la primera comprobación operativa y se detuvo el intento antes de aprovisionar. [Registro](evidence/remote-execution-20261022.json). El reparto original siguiente es histórico: **no demuestra una cota ejecutable de 100 solicitudes**. La autorización posterior exige al menos 10 reservadas para retirada y seis builds como techo absoluto.

Preparado el **2026-10-22**, después de [GO_LOCAL](evidence/integral-local-gate.json). **No ejecutado.** Requiere **APRUEBO SPIKE REMOTO**. No autoriza producción, pagos, cambios a recursos preexistentes ni merge. [Auditoría vigente](REMOTE_AUDIT.md), [evidencia de cuentas](evidence/integral-provider-audit.json).

## 1. Condiciones previas y presupuesto

Las cuentas/planes/cuotas Vercel, Neon y Ably no están verificadas: no se puede provisionar todavía ni afirmar coste mensual disponible. Primero identificar la cuenta autorizada y Free/Hobby efectivos, saldo y límites; detener cualquier trial/upgrade/tarjeta/addon. Si una capacidad mínima exige pago, registrar incompatibilidad y presentar impacto, sin cambiar stack o alcance unilateralmente.

Presupuesto: **0 USD adicionales**. Máximo **100 solicitudes externas iniciadas por la ejecución instrumentada**, con máximo cinco en vuelo y concurrencia normal dos: 70 asignadas a casos, 20 a control y 10 reservadas, seis exclusivamente para retirada y cuatro para errores/reintentos. Las llamadas de emisión/publish/recepción, redirects y reintentos explícitos consumen presupuesto; retries automáticos deshabilitados cuando sea posible. Los requests de infra internos del proveedor no son carga generada por el harness ni se extrapolan como cuota disponible. Si el cliente/CLI no permite contabilizar, usar API instrumentada o detener el caso; no ejecutar un despliegue opaco que consuma sin control.

SQL: hasta **80 sentencias** sintéticas durante pruebas, dos conexiones de app **por pool/instancia**, cero cargas de saturación. Varias Functions pueden abrir varios pools: no equivale a dos conexiones globales por servicio. Techo experimental previsto: ocho conexiones PostgreSQL backend de app en total (cuatro por rol) y una directa de migración; comprobar límites efectivos y observar también clientes/colas del pooler. Si el proveedor no permite controlar/observar ese agregado, mantener BLOCKED la prueba de presupuesto de conexiones. Ably: hasta **20 publicaciones pequeñas** (solo las necesarias, descontadas del presupuesto HTTP cuando proceda), dos clientes, tres canales como máximo, payload hasta 1 KiB, TTL de token hasta cinco minutos. Dataset máximo previsto: 20 identidades, 20 agregados, 100 receipts; objetivo menos de 1 MiB lógico. Máximo **seis builds remotos** —cuatro iniciales y repetición sin caché de Game/AI—; no suponer cuota suficiente.

Contabilidad propuesta de llamadas, sin contar dos veces una observación compartida:

| Bolsa | Clases y máximos | Total |
|---|---|---:|
| POC-01 | Seis envíos de build/deployment, una consulta terminal por cada uno, dos versiones/hash autenticadas de Game/AI | 14 |
| Otros POC | Requests funcionales/negativos/mediciones y cuentas/consumo según §5; pueden reutilizar respuestas anteriores | 56 |
| Control | Seis creaciones de recursos, cuatro configuraciones Vercel agrupadas, tres lecturas de inventario/preflight, una configuración Ably, seis recuperaciones de logs | 20 |
| Reserva de retirada | Un borrado por cada uno de los seis recursos dedicados; revocación por eliminación solo si la API lo garantiza | 6 |
| Reserva de incidentes | Poll adicional, fallo, redirect, revocación separada, confirmación de retirada o reintento, sin comprometer retirada | 4 |
| Techo | Todas las solicitudes contadas, incluidas las negativas | **100** |

Este reparto **no demuestra todavía** que las APIs disponibles puedan completar el plan en 100 llamadas. Antes de crear, comprobar endpoints, paginación, configuración agrupada, obtención de logs, carga de fuentes y eliminación/revocación. Una sola consulta terminal por build es un límite, no una promesa de que el build terminará a tiempo: esperar localmente y consultar una vez; si sigue pendiente, consume reserva o queda BLOCKED. La fuente prevista es gitSource de una revisión publicada con seguridad; uploads fragmentados o nuevas consultas GitHub también consumen presupuesto si hacen falta. Si el ledger concreto excede el techo, falta reserva segura de retirada o la API no garantiza revocación al borrar, **detenerse antes de aprovisionar** y presentar el presupuesto revisado para aprobación. No reducir cobertura ni ocultar llamadas para aparentar que cabe.

Registrar consumo anterior/posterior y reservar capacidad de retirada antes de crear. Si un cupo, coste o visibilidad no se puede demostrar, mantener ese caso BLOCKED. El objetivo $0 mensual es condicional al plan/consumo, no resultado observado de un mes.

## 2. Recursos propuestos — nombres todavía no creados

| Nombre propuesto | Proveedor/plan previsto | Finalidad y seguridad | Datos permitidos | Creación/retirada prevista | Prueba y evidencia para PASS |
|---|---|---|---|---|---|
| `tk-spike01-web` | Vercel Hobby, por verificar | Proyecto Next separado; Preview protegido, destinos HTTPS fijos de Identity/Game, sin secretos públicos | Marcadores e IDs sintéticos | API/dashboard bajo cuenta identificada; eliminar solo proyecto/deployments propios inventariados | POC-01/06: build y origen real, 401/403, cookies/cabeceras; browser/TLS según caso |
| `tk-spike01-identity` | Vercel Hobby, por verificar | Root propio Go; Handler API existente; autenticación propia y protección Preview | Identidades sintéticas de la sonda | API con gitSource/revisión publicada segura; eliminar solo proyecto propio | POC-01/02/03/07: build independiente, conectividad de su rol, rechazo cruzado |
| `tk-spike01-game` | Vercel Hobby, por verificar | Root Go separado; engine único; autoridad SQL; publicación oficial restringida | Agregados/revisiones/efectos/receipts sintéticos | Igual, sin `--prod`; retirada de proyecto/secretos propios tras evidencia | POC-01–05/07: módulo, atomicidad, duplicación/carrera, recuperación y permisos reales |
| `tk-spike01-ai` | Vercel Hobby, por verificar | Root Go separado, mismo engine; sin DB ni privilegio de comando Game; cálculo acotado | Contexto sintético efímero | Igual, nunca credenciales DB; retirar proyecto/secretos propios | POC-01/02/07/10: módulo exacto, auth negativa, cancelación y revisión obsoleta |
| `tk-spike01-db` | Neon Free, por verificar | Proyecto dedicado; propuesta de dos DB `tk_identity_probe` y `tk_game_probe`, propietarios separados y app roles sin privilegios administrativos | Solo fixtures anteriores | Console/API; SQL directo para bases/roles/migraciones; eliminar solo proyecto dedicado conforme al inventario | POC-03/04: consultas con pgx pooled, directa para migración, roles negativos, commit/rollback y límites |
| `tk-spike01-realtime` | Ably Free, por verificar | App dedicada, clave publicadora de Game y emisor restringido; canal exacto por agregado; cliente solo `message-subscribe` | Avisos ID/revisión/tipo; nunca DSN, sesión o estado privado completo | Dashboard/Control API; revocar claves/tokens y eliminar solo app propia | POC-05/08: entrega real y denegación de canal/publish, fallo controlado y REST |

Los cuatro proyectos no multiplican cuotas de cuenta/team. Dos propiedades PostgreSQL significan aquí dos bases con propiedad/credenciales aisladas dentro del proyecto exclusivo; es una propuesta experimental, no elección física definitiva para producto. Si el proveedor no soporta el aislamiento con sus roles/cuotas, registrar FAIL/BLOCKED según evidencia, sin compensarlo con privilegios globales o SQL cruzado.

## 3. Matriz de secretos y exposición

| Unidad | Material exclusivo de integración | Material que no debe recibir |
|---|---|---|
| Web | Configuración de rewrites y, si se necesita, bypass de protección exclusivamente del lado servidor | Claves de firma, DSN, clave Ably publicadora, secretos en NEXT_PUBLIC/URL/HTML/JS |
| Identity | DSN app Identity, clave de verificación/sonda para su audiencia | DSN Game y credencial de publicación Game |
| Game | DSN app Game; permiso limitado de publicación Ably; credencial S2S hacia AI/Identity; verificadores necesarios | Credenciales de producción o administrativas globales |
| AI | Verificación de su audiencia y permiso exclusivamente para cálculo sintético | Cualquier DSN, clave de usuario Game o permiso de escritura oficial |
| Operador de migración | Roles propietarios/migradores distintos de app, conexión directa, guardados fuera del repositorio | Exportarlos al runtime/web o migrar al arrancar Functions |
| Cliente de prueba | Credencial sintética corta y token Ably temporal para canal exacto | Clave Ably maestra/publicadora, credenciales de otro agregado |

Usar roles PostgreSQL sin heredar privilegios administrativos globales, revocar CONNECT ajeno y comprobar SQLSTATE de denegación en ambas direcciones. App role solo DML necesario; migrador DDL de su propiedad. No convertir una denegación debida a red en PASS de aislamiento.

Propuesta de protección: **All Deployments** con Vercel Authentication, sin dominios/target de producción, y autenticación propia obligatoria. Validar disponibilidad efectiva de Hobby y S2S/rewrites a través de la protección. Automation bypass solo por header servidor, con secreto dedicado; nunca querystring/client JS. No desactivar protección para conseguir PASS. Si este modelo no permite el experimento con exposición aceptable, detener y presentar alternativas autorizables.

El HMAC local no certifica seguridad productiva. Mantener credenciales de usuario por audiencia y no entregar a AI material capaz de firmar usuario Game. La clave S2S compartida local necesita revisión de mínimo privilegio para el remoto; no extrapolar resistencia a un servicio comprometido. Google OIDC completo no se implementará: seguirá NOT_RUN si no se prueba contra un proveedor autorizado.

## 4. Raíces y empaquetado Go: experimento decisivo

Roots propuestos: `spikes/infra-feasibility/apps/web` y `spikes/infra-feasibility/services/{identity,game,ai}`. Cada Go root contiene go.mod/go.sum y el adaptador legado `api/index.go`. Game/AI requieren exactamente engine `v0.1.0` con replace `../../modules/engine`; los tres usan probe fuera de root. No usar go.work ni duplicar fuente del motor. Versión nominal local sin release/tag publicado: identificar también revisión y hash, no presentar v0.1.0 como descarga pública existente.

La opción **Include source files outside of the Root Directory in the Build Step** ofrece un camino documental, sin probar el contexto del builder Go. Activarla solo en proyectos dedicados tras el gate y registrar el valor efectivo. Mantener la topología del checkout, desplegar cada proyecto por separado y repetir Game/AI con caché de proyecto vacía. Capturar toolchain/config/entrada usados, módulo/version/replace resuelto y hash del engine, además de ejecutar operaciones autenticadas. El hash local de referencia es `BD0D2F40296E32B6C832DEA9B22C0CB1B957E428D2DC01E20456C303AB14C973`.

No basta healthcheck ni un build desde raíz común. Si falla resolución externa, conservar logs y tratar impacto; no unir microservicios, vendorizar dos fuentes divergentes, añadir motor TypeScript, cambiar stack o pagar. No asumir que un límite Fluid publicado para otros runtimes se aplique a Go.

**Adaptaciones experimentales previas al primer deploy, después del gate:** reemplazar los destinos loopback locales de Next/Game por una allowlist cerrada de las URLs HTTPS dedicadas; conservar validación estricta, timeouts y auth. El CLI actual escucha loopback/puertos fijos: no se está eligiendo automáticamente el preset Go basado en PORT. Mantener el Handler API, verificar que esa entrada real es la ejecutada y no afirmar que el main local es un preset alojado ya compatible. Conectar la notificación a Ably real con credenciales/capacidades restringidas, sin confundir el hook del harness con integración ya implementada. Repetir controles locales afectados antes de exponerlos.

## 5. Casos remotos y límites de aceptación

Todos están **BLOCKED / NOT_RUN en proveedor**. Las 70 solicitudes siguientes son asignaciones máximas, incluidas negativas y muestras; si se necesitan más pruebas para resolver un fallo, parar al agotar el presupuesto y pedir revisión, no relajar criterios.

| ID | Criterio objetivo en entorno real | Solicitudes máximas |
|---|---|---:|
| POC-01 | Cuatro builds/deployments independientes más dos repeticiones frías Game/AI; mismo módulo/versión/hash, Identity sin engine; versiones autenticadas corresponden a la revisión | 14 |
| POC-02 | Handler Gin real responde con semántica/rutas/status previstos; entorno/runtime identificados | 2 |
| POC-03 | Pooled pgx de cada app y directa de migrador; roles propios operan y ajenos fallan por permiso; AI sin DB; conexiones dentro del límite observado | 4 + SQL acotado |
| POC-04 | Primera intención, retry/cambio de payload, carrera de dos, fallo precommit y pérdida de respuesta/retry conservan decisiones/contadores/receipts; nuevas invocaciones no dependen de RAM | 12 |
| POC-05 | Aviso real posterior a commit; fallo de publicación no revierte; cliente recupera estado REST tras pérdida/reintento; ninguna notificación confirma un comando | 6 |
| POC-06 | Same-origin rewrites y credenciales atraviesan protección correctamente; negativos de acceso/origen; cookies de prueba con navegador/TLS cuando se mida enforcement; sin secretos en JS/URL/logs | 6 |
| POC-07 | S2S válida funciona y falta/firma/audiencia/recurso/AI-write son rechazados; caída interna produce error controlado; no se omite protección Preview | 8 |
| POC-08 | Token efímero canal exacto recibe; otro canal no; cliente no publica; Game sí; permisos/TTL auditados en clave y token reales | 8 |
| POC-09 | Registrar muestras frías/calientes/DB inactivity y regiones efectivas; sin p95 ni SLO con muestra insuficiente; reutilizar mediciones de otros casos | 0 adicionales |
| POC-10 | Deadline/cancelación/work limit y revisión obsoleta se cumplen dentro del runtime real; sin efecto oficial tardío | 4 |
| POC-11 | Consumo anterior/posterior y cuota disponible de tres proveedores; cero upgrade/gasto; carga dentro del cupo observado | 6 |
| Total de casos | No sumar otra vez las llamadas instrumentadas compartidas entre casos | **70** |

Selección de región: aún pendiente. Antes de crear, verificar Vercel y Neon en una región próxima permitida para esa cuenta, privacidad y plan; documentar disponibilidad y decisión experimental. Ably es red gestionada: registrar ubicación del cliente/camino observado, sin atribuir al experimento una región privada elegida. No asumir región por defecto como óptima.

Si protección consume más requests de autenticación que los presupuestados, el contador global prevalece. No hacer load test ni agotar cuota para probar un rechazo; simular límites localmente y observar solo la pequeña carga autorizada.

## 6. Orden, evidencia y retirada

1. Tras aprobación y conexión válida: verificar cuentas/planes/cuotas/protección e inventario previo. Resolver riesgo de auto-deploy antes de push o asociación Git.
2. Revisar adaptaciones locales, configuración de Root Directory y secretos dedicados; fijar revisión. Crear solo los nombres anteriores, abortando si colisionan con un recurso preexistente.
3. Probar builds/módulos, aislamiento/conexiones, transacciones, fallos/REST, Ably y seguridad; después métricas/consumo. Guardar versiones, inputs sintéticos, status, snapshots SQL y logs sanitizados por ID de caso.
4. Comparar consumo y cuotas, emitir PASS/FAIL/BLOCKED por caso y recomendación global justificada. GO exige toda condición obligatoria real; falta de permiso/credencial/evidencia es BLOCKED, no NO-GO.
5. Plan de retirada exclusivo del spike: capturar evidencia, revocar sus tokens/keys/bypass, eliminar deployments/proyectos Vercel dedicados, app Ably y proyecto Neon propios. Confirmar inexistencia y consumo final. No eliminar recursos preexistentes ni de otra tarea; cualquier duda requiere autorización adicional.

Los cuatro vercel.json locales tienen Git deployments deshabilitados, pero eso no descarta otra integración con distinto root/configuración. [La auditoría](REMOTE_AUDIT.md) no pudo demostrar push seguro. Por tanto, no se publicó la rama ni se abrió Draft PR en esta ejecución. No se autoriza cambiar ajustes externos para desbloquear la publicación antes del gate.

**Gate solicitado: APRUEBO SPIKE REMOTO.** El plan está preparado; crear/desplegar sigue prohibido hasta permiso expreso y condiciones previas verificables.

## 7. Adenda de ejecución y contabilidad

Se recibió la aprobación explícita del alcance; no se pide aprobarlo otra vez. Se mantiene la secuencia protección → secretos → builds → pruebas → retirada y la obligación de parar antes de crear si falla una precondición. La primera lectura de Neon fue rechazada 401; no se ejecutó ninguna etapa A–E, push ni creación de recursos. Para reanudar se requieren sesión Neon válida y repetición de todos los prechecks, con inventario actual que preserve **9 Stones** y **Player chat**.

| Contador observado del intento | Valor y alcance |
|---|---|
| HTTP directo a cuentas/proveedores | **1**: `REMOTE-001`, GET Neon, 401; cero retry/redirect/refresh |
| Solicitudes funcionales / SQL / GitHub del repositorio Toguz / despliegues | **0 / 0 / 0 / 0** |
| Builds remotos | **0 de 6** |
| Recursos o credenciales creados / retirados | **0 / 0**; retirada no necesaria |
| Investigación pública delegada | **13 llamadas a la herramienta web**: 8 búsquedas, 18 búsquedas dentro de páginas y 31 aperturas/clics, de estas 14 satisfactorias y 17 fallidas |
| Tráfico HTTP interno de esa herramienta | **No observable**; sus operaciones no se equiparan a un número exacto de solicitudes HTTP |
| Total global y saldo frente a 100 | **No certificables**; no se presenta una solicitud de proveedor como el total de todo el tráfico |
| Reserva de retirada | Regla mínima de **10**, no utilizada por pruebas funcionales; no se afirma saldo global exacto |
| Coste adicional contratado | **Ninguno**; no se crearon recursos, planes, trials o add-ons; facturación no reconsultada |

La lectura de documentación pública se realizó en paralelo para revisar protección y presupuesto. Al no pasar por un contador HTTP observable, no permite demostrar el límite total exigido. Esto constituye una limitación adicional de contabilidad y debe resolverse antes de toda futura operación, sin excluir silenciosamente consultas o reintentos. La revisión local delegada realizó cero solicitudes externas y cero escrituras. La única evidencia nueva de cuenta es el 401; las demás cuotas no se revalidaron.

La revisión encontró observaciones compartibles sin reducir cobertura: POC-02 puede usar versiones/conectividad/operaciones de POC-01/03/04; POC-09 utiliza sus mediciones; POC-11 comparte lecturas iniciales/finales con prechecks y retirada; POC-05/08 comparten publicación real y recepción autorizada; el retry de POC-04 puede seguir a la repetición fría Game de POC-01. No se demuestra por ello que todo quepa en 100. Deben contarse también Web→API, Game→Identity/AI, Game→Ably, autenticación, cargas Git, verificaciones de protección, configuración Neon y emisión/revocación Control API. Mantener contadores separados HTTP y SQL, sin ocultar ninguno; la unidad del techo debe quedar explícita antes de ejecutar.

El ledger futuro debe reservar como mínimo diez solicitudes de retirada y demostrar que bastan para revocar credenciales, borrar únicamente los seis recursos dedicados y comprobar la conservación de los preexistentes. Debe incluir las precondiciones nuevas, paginación, cargas, polling/logs, errores y las llamadas descendentes de la aplicación. Ninguna bolsa propuesta de 66/24/10 o 70/20/10 constituye una garantía. Si no puede demostrarse la cota restante, continuar **BLOCKED_REMOTE** antes de crear recursos; no retirar cobertura ni debilitar seguridad para encajar.

La documentación pública consultada indica que [Vercel Authentication y All Deployments](https://vercel.com/docs/deployment-protection/usage-and-pricing) y [Trusted Sources](https://vercel.com/docs/deployment-protection/methods-to-bypass-deployment-protection/trusted-sources) están disponibles en Hobby. Trusted Sources/OIDC es el candidato preferente: autorizar únicamente proyectos y entornos necesarios, sin permiso CI implícito del mismo repositorio ni proveedores externos. Son condiciones documentales, **no configuración de cuenta comprobada ni POC PASS**. Crear sin `gitRepository`, comprobar protección antes de secretos y conservar fuentes fuera de Root Directory requieren verificar las respuestas API efectivas. No se configuró nada en este intento.

## 8. Continuación tras renovar Neon: unidad y presupuesto revisados

El propietario confirmó expresamente: **«Sí: 100 operativas acumuladas; SQL y documentación aparte»**. Se resuelve la ambigüedad del apartado 7 sin reiniciar el contador: `REMOTE-001` sigue costando una solicitud. Las diez nuevas, `REMOTE-002`–`REMOTE-011`, están identificadas por endpoint, hora, HTTP y observación sanitizada. No hubo renovación automática, retry de aplicación, redirect, SQL, pruebas funcionales, builds ni escritura externa. El POST GraphQL contiene exclusivamente una consulta de lectura.

**Control del cupo:** `100 = 11 observadas + 79 disponibles antes de retirada + 10 reservadas`. Esto es un saldo cerrado y verificable del contador, **no prueba que la ejecución completa quepa en ese saldo**. Las consultas públicas de documentación se excluyen por decisión expresa y se conservan en un contador separado; el tráfico interno del buscador ya no bloquea la definición de esta unidad operativa.

Se incorporaron dos ahorros acreditados: validar `ssoProtection.deploymentType=all` en la respuesta real de [creación Vercel](https://vercel.com/docs/rest-api/projects/create-a-new-project), sin un GET adicional; y enviar `env` junto con `build.env` dentro del POST de deployment posterior, como hace el CLI oficial instalado. **Los proyectos se crean primero sin secretos** y solo se despliegan con credenciales después de verificar protección. No se introducen secretos en el POST inicial para ahorrar llamadas. [Crear deployment](https://vercel.com/docs/rest-api/deployments/create-a-new-deployment) admite archivos pequeños inline, lo que permite proponer este experimento sin push ni uploads fragmentados; aún no se ha ejecutado esa ruta.

La siguiente es una secuencia concreta revisable, **no un mínimo universal ni una ejecución realizada**:

| Grupo | Operaciones contabilizadas | HTTP |
|---|---|---:|
| P | Precondiciones acumuladas ya observadas, incluido el 401 | 11 |
| C | Crear cuatro proyectos Vercel, un proyecto Neon y una app Ably | 6 |
| V | Cuatro PATCH con Trusted Sources y fuentes exteriores cuando corresponda | 4 |
| K | Dos claves Ably separadas: publicador y emisor de suscripción | 2 |
| B | Seis deployments, seis consultas terminales, seis recuperaciones de logs | 18 |
| F | F01–F51: pruebas, negativos y llamadas Web→API/Game→servicios/Ably | 51 |
| U | Consumo final de cada proveedor, mecanismo Free todavía no acreditado | 3 |
| T | Retirada conservadora explícita y verificación de credenciales | 13 |
| **Secuencia contabilizada** | **Todavía excluye prechecks pendientes y alta de PAT** | **108** |

La evidencia enumera **F01–F51 individualmente**. Incluye ocho pruebas de capacidad Ably, seguridad de aplicación/protección, cookies, S2S, rollback, retry, carrera acotada y recuperación. El presupuesto no se reduce omitiendo llamadas descendentes. Reutilizar F03/F12 como la misma firma inválida y alcance de clave ahorraría una petición. Compartir recuperación S2S con la operación AI válida y su metadata podría ahorrar otras tres, pero exige adaptar y verificar el harness: aun admitiéndolas resultarían **104 más pendientes**, con esta secuencia. Seis builds son el plan actual y el techo; no se afirma que se hayan ejecutado ni que toda alternativa requiera seis.

No se presume que el EOF de [deployment events](https://vercel.com/docs/rest-api/deployments/get-deployment-events) demuestre READY. Sustituir consultas terminales por respuestas funcionales de la URL inmutable, logs y procedencia sería otra optimización a demostrar; el `/version` actual informa `local-unpublished` y no basta para acreditar el SHA. Tampoco se supone que crear una app Ably devuelva dos claves: el contrato oficial `app_response` no las incluye. La [Control API oficial](https://ably.com/docs/open-specs/control-v1.yaml) garantiza DELETE de app/204, sin una promesa explícita de revocación en cascada suficiente para este gate.

Desglose conservador de retirada: revocar dos claves Ably (2), borrar solo los seis recursos propios (6), consultar los tres inventarios para comprobar ausencia/preservación (3), revocar el PAT si se creó exclusivamente para el spike (1), y comprobar que ese PAT ya no autoriza GET `/me` (1): **13**. El mecanismo de revocación del PAT y la instrumentación de su alta aún no están acreditados. Credenciales de operador preexistentes ajenas al spike no deben revocarse. Se conserva la reserva mínima autorizada de diez; **no se afirma que diez basten** para este diseño. Ningún recurso se crea mientras no exista un plan completo de retirada que encaje en las 100.

SQL permanece separado, con techo original de 80 sentencias. La revisión estática del harness actual cuenta 82 en preparación/transacciones/retry, incluyendo BEGIN/COMMIT/ROLLBACK, antes de roles, seguridad y notificaciones; **no son sentencias ejecutadas en esta continuación**. Es necesario otro ledger y compartir snapshots sin perder oráculos. Los adaptadores remotos de cookies, HTTPS, OIDC, procedencia y Notify real tampoco están implementados; esta continuación no modifica código para ocultar sus costes.

**Gate de creación: BLOCKED.** Sesiones Neon/Vercel válidas, pero faltan credencial Control API Ably disponible para consultas contables, saldo actual de cuotas de los tres proveedores, revalidación completa de instalaciones externas si se propone publicar y un presupuesto completo demostrado dentro de 100. El propietario confirmó **«No; solo tengo la sesión web»** respecto del PAT local. No se solicitó su valor ni se creó un PAT para suplirlo mientras el presupuesto permanece sin validar. No se hizo push, emisión de credenciales, aprovisionamiento, despliegue ni cambio al PR #3. No se aumenta el cupo ni se recorta cobertura por decisión unilateral.

## 9. Plan vigente tras retirar el techo de solicitudes

La decisión posterior del propietario y las cuotas actuales constan en [la auditoría §11](REMOTE_AUDIT.md#11-decisión-posterior-y-comprobación-de-cuotas--2026-10-22). Sustituye exclusivamente el límite de100 solicitudes; no hay nuevo permiso de pago, lote4, producción o merge. No se recorta cobertura para sostener una cifra histórica. API directas y observaciones UI se registran por separado, reconociendo el tráfico interno no instrumentado del dashboard.

1. **A:** verificar HEAD9409462/blobGO_LOCAL; crear cuatro configuraciones Vercel sin Git, builds o secretos; exigir respuesta efectiva All Deployments. Si no se conserva protección o requiere pago, retirar lo propio y parar.
2. **B:** incluir archivos versionados pequeños inline conservando raíces; habilitar fuentes exteriores; ejecutar cuatro builds independientes y hasta dos repeticiones frías Game/AI, máximo absoluto6. Logs y estado terminal se consultan con backoff acotado, sin retries de build automáticos. Comprobar módulo compartido exacto y ejecución Handler; fallos de empaquetado detienen dependientes.
3. **C:** solo tras B, crear un Neon dedicado; dos propiedades/roles y pruebas SQL positivas/negativas dentro del techo separado80, con snapshots compartidos y un fixture; nada en9Stones. No ejecutar el harness local82 sin adaptación de su ledger.
4. **D:** solo tras C, crear appAbly dedicada y claves separadas mínimas; pruebas realesCAP01–08, eventos posteriores a commit y recuperaciónREST. Conservar Player chat y9stones. JWT firmado localmente es una opción documentada sin canje `requestToken`, si se comprueba su uso real; no se cuenta un TokenRequest como token emitido.
5. **E:** cookies/rewrites/S2S bajo protección, cancelaciónAI, muestras y consumo final. Cada POC exige evidencia; no sustituir con afirmaciones ni resultados locales.
6. **Retirada:** revocar claves propias, borrar únicamente seis recursos si llegaron a crearse, verificar ausencia y recursos previos. La revocación del PAT dedicado está documentada en el dashboard; no se inventa un endpoint ControlAPI ni se presume una sola petición UI. Su coste ya no bloquea una cota100 retirada por el propietario. Credenciales de operador ajenas al experimento no se revocan.

Se mantienen concurrencia normal2, máximo5, sin saturación, hasta20 publicaciones≤1KiB, dos clientes y tres canales. Los límites del plan gratuito se comprueban antes/después; cualquier exigencia de upgrade detiene la prueba. [JWT Ably](https://ably.com/docs/auth/token/jwt), [revocación de PAT](https://ably.com/docs/platform/account/access-tokens).

## 10. Excepción operativa propuesta, no aprobada

**Estado: NO EJECUTAR.** El [incidente y la retirada](REMOTE_AUDIT.md#12-prueba-real-de-protección-y-destino-inicial-bloqueo-específico-recursos-retirados) acreditan promoción inicial production incluso con staging explícito. No se pretende resolverlo cambiando arquitecturas, pagando, desactivando protección o llamando Preview a production. Hace falta una decisión expresa sobre la restricción del propietario **«Sin despliegue de producción»**, conservada en el Gate3.5-B, y el máximo6 builds del adjunto remoto.

Alternativa revisable para continuar: permitir el destino production automático **solo al inicializar** cada uno de los cuatro proyectos nuevos `tk-spike01-{web,identity,game,ai}`; todos pertenecen exclusivamente al spike y se retirarían. Ningún proyecto real, dato real, dominio adicional, conexión Git, trial o coste se autoriza. Exigir nuevamente All Deployments antes de secretos y comprobar acceso anónimo denegado. El nombre técnico del entorno se registra literalmente: nunca se lo presenta como Preview ni como producto publicado.

Proponer máximo **8 builds acumulados**: los2 cancelados observados +4 builds independientes +2 repeticiones frías Game/AI. Si un fallo exige otro build, parar antes del noveno. Preservar cuotas gratuitas, concurrencia máxima5, SQL80 con ledger propio y ausencia de cargas de saturación. Preparar Neon solo tras protección y antes de los builds que necesiten DSN para comprobar el Handler real; ninguna prueba SQL positiva/negativa se ejecuta sobre9Stones. Ably dedicado y secretos mínimos, sin tocar apps existentes. Las adaptaciones experimentales pendientes se prueban localmente antes de exponerlas; no se implementa producto ni lote4.

Retirar todos los recursos y credenciales propios al acabar, conciliar uso y demostrar ausencia. El PAT ControlAPI de operador se revoca solo si se confirma exclusivo del spike; no se revocan sesiones generales de Vercel/Neon ni credenciales ajenas. Esta propuesta **no cambia por sí misma la autorización** y no permite seguir desplegando hasta respuesta expresa.
<a id="remote-plan-close-20261025"></a>

## 11. Ejecución posterior y desviaciones — 2026-10-25

Plan cerrado bajo autorización amplia de preparar lote4: protección primero, credenciales exclusivas después, cuatro builds iniciales y dos repeticiones sin caché; ocho acumulados contando dos cancelados históricos. Bootstrap production solo en cuatro proyectos propios protegidos; repeticiones Preview, cero --prod, pagos o producto.

100 HTTP fue retirado por el propietario; ledger no reiniciado. El objetivo técnico SQL80 **se superó**:105 operaciones reconstruidas principales, incluidas27 de preparación, más12 instrumentadas en seguimiento. No afirmar cumplimiento80 ni conteo wire exacto. Se preservaron0USD, carga sintética acotada, dos contendientes y planes gratuitos.

Se ejecutaron permisos, transacciones, Ably, Chrome y OIDC. Después de retirar el primer conjunto, revisión independiente detectó cronómetros C6 incompletos: un reemplazo Neon y Ably permitió medirlos sin otro build; ambos se retiraron de inmediato. [Auditoría y límites](REMOTE_AUDIT.md#remote-close-20261025), [matriz](TEST_MATRIX.md#remote-matrix-20261025), [C6](evidence/remote-latency-followup-20261025.json).

Reproducción futura: tests/Remote-Control.ps1 usa sesiones oficiales y configuración DPAPI; tests/Remote-Build.ps1 preserva roots/allowlists. cmd/remotesetup, cmd/remotecheck y cmd/remotemetrics reciben configuración por entorno, nunca argumentos con secretos. Leer este plan no autoriza recreación automática: IDs están retirados y otro experimento requiere alcance/cuotas vigentes.

Quedan para pruebas específicas antes de producto: contabilidad SQL exhaustiva, clientes/colas PgBouncer, suspensión Neon, JWT/JWKS de producto, caída física de proveedores, pérdida TCP real, SLO y fuerza IA. No bloquean redactar las ocho specs04/08.
