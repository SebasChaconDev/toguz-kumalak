# Gate 3.5 — plan de validación remota controlada

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
