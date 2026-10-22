# Gate 3.5 — experimento desechable de infraestructura

**Estado vigente, 2026-10-22: GO_LOCAL.** [Informe integral](INTEGRAL_REPORT.md), [matriz POC](TEST_MATRIX.md), [auditoría vigente](REMOTE_AUDIT.md) y [plan remoto pendiente](REMOTE_PLAN.md). PostgreSQL real, reinicio de procesos/contenedor, builds aislados y race Linux se verificaron de nuevo. Vercel/Neon/Ably reales siguen pendientes; push y Draft PR están retenidos por automatización externa no descartada. Se requiere **APRUEBO SPIKE REMOTO** antes de escribir en proveedores.

Autorización: solicitud del propietario «GATE 3.5: SPIKE DE VIABILIDAD DE INFRAESTRUCTURA», 2026-10-21 (Caracas). Permite código mínimo, dependencias gratuitas, servicios y pruebas **locales**; proveedores externos requieren **APRUEBO SPIKE REMOTO**. Este árbol no implementa el MVP ni las reglas TK-N1-1.0.0 y no constituye código listo para producción.

Actualización autorizada el **2026-10-22**: el propietario solicitó reparar Docker y ejecutar correctamente las pruebas locales. Docker Desktop fue actualizado de 4.65.0 a 4.96.3; los cinco casos de integración PostgreSQL pasaron. El fallo y la salida anteriores se conservan como historia. [Resolución y evidencia](REPORT.md#docker-repair-20261022).

Base `main`: `aaeeace21860305091c0351bedb0aca0df386816`. Rama `spike/infra-feasibility-01`. PR #3 permanece abierto en `9db0f3e9d4bf1fb642697103d9232bcfaef42205`; sus contratos 05.01–05.04 y 07.02 son referencias provisionales. La instrucción posterior autoriza exclusivamente este directorio, sin modificar documentación aprobada, decisiones o ADR.

## Mapa y límites

| Ruta | Propósito |
|---|---|
| `apps/web` | Next.js independiente; página experimental y rewrites locales de destinos fijos. |
| `services/identity` | Módulo Go/Gin propio; health, versión, conectividad DB y operación interna autenticada. Sin motor. |
| `services/game` | Módulo Go/Gin propio; transacción sintética, recibos idempotentes y guardia de revisión. |
| `services/ai` | Módulo Go/Gin propio; trabajo sintético cancelable. Rechaza entorno con credenciales DB. |
| `modules/engine` | **Única** fuente del sustituto puro: SHA-256 determinista; no siembra, captura ni evalúa partidas. |
| `modules/probe` | Biblioteca de apoyo para HTTP, autenticación experimental y pgxpool; no es otro servicio ni otro motor. |
| `infra`, `tests`, `evidence` | Compose, migraciones sintéticas, procedimientos y resultados sanitizados. |

Game y AI requieren el mismo módulo `engine v0.1.0`, resuelto por `replace ../../modules/engine`, y la constante `spike-engine/v0.1.0`. Es una versión experimental local sin tag publicado; su hash identifica los bytes efectivamente compilados. No se usa `go.work`. Los helpers no son contratos de producto; no se introducen Redis, sockets propios ni un cuarto microservicio.

## Ejecución local reproducible

Herramientas iniciales: Windows amd64, PowerShell 7.6.5, Go 1.26.2, Node 24.13.0, npm 11.6.2, Docker CLI 29.2.1 y Compose 5.1.0. Después de la reparación se verificó Docker Desktop 4.96.3/Engine 29.9.2 y PostgreSQL 18.6. Compose fija ahora la imagen al digest observado `sha256:edb0c08dc33befd0262ab3b5cf1c4f3f8cfd8b09e88c05b45adee29d6761ddb6`; los números iniciales se conservan como procedencia de las primeras pruebas.

Desde la raíz del repositorio. Para PostgreSQL, elegir **una** alternativa: primera ejecución con generación de credenciales, o arranque con las credenciales ya existentes (caso actual). `Setup-Local.ps1` rechaza sobrescribirlas.

```powershell
# Tests Go; ColdBuild crea contextos/copias y cachés nuevos por cada servicio.
./spikes/infra-feasibility/tests/Verify-Local.ps1 -ColdBuild

# Elige automáticamente según haya credenciales locales previas:
if (Test-Path ./spikes/infra-feasibility/.local/compose.env) {
  docker compose --env-file ./spikes/infra-feasibility/.local/compose.env -f ./spikes/infra-feasibility/infra/compose.yaml up -d --wait
} else {
  # Genera secretos sintéticos solo en .local/ (ignorado) y arranca Compose.
  ./spikes/infra-feasibility/infra/Setup-Local.ps1
}
./spikes/infra-feasibility/tests/Run-Integration.ps1

Push-Location spikes/infra-feasibility/apps/web
$env:NEXT_TELEMETRY_DISABLED='1'
npm ci --ignore-scripts --no-fund --no-audit
npm test
npm run build
npm run smoke
Pop-Location
```

`Verify-Local.ps1 -ColdBuild` exige contextos inexistentes para no confundir una segunda ejecución con caché limpia. Para repetir, usar un checkout experimental nuevo o conservar/retirar expresamente los contextos anteriores. Sin `-ColdBuild`, repite los tests con `-count=1`. Restaurar variables de entorno previas forma parte de los scripts; no cargar `test.env` globalmente en una terminal ni pasarlo a AI. No imprimir `docker compose config` sin `--quiet`, pues expone variables interpoladas.

Cada build parte de `services/<servicio>` en un contexto que contiene sus fuentes, `go.mod/go.sum`, `modules/probe` y, solo para Game/AI, `modules/engine`. `GOWORK=off`, `GOENV=off`, `GOTOOLCHAIN=local`, `CGO_ENABLED=0`, `GOCACHE` y `GOMODCACHE` nuevos por proyecto, `go build -mod=readonly -trimpath`. No requiere archivos de producto de la raíz. La web usa únicamente su subárbol y lock npm. Salidas/cachés y credenciales quedan ignoradas.

Los comandos `cmd/server` escuchan únicamente `127.0.0.1`, puertos Identity 18101, Game 18102, AI 18103. Next usa 18100. Las entradas `api/index.go` exportan `Handler(http.ResponseWriter,*http.Request)` e invocan el mismo Gin. Se compilan y ejercitan localmente; no se ha ejecutado el builder de Vercel.

## Autenticación experimental

Las credenciales de la sonda son un sobre JSON firmado con HMAC-SHA256, claves distintas para usuario y servicio, `kind`, `sub`, `aud`, `scope=probe`, caducidad máxima de cinco minutos. **No son JWT/OIDC ni una implementación de Identity**. No existe ruta HTTP de emisión ni bypass anónimo. Las claves aleatorias locales se reciben solo por entorno; ausencia, igualdad o longitud insuficiente falla de forma cerrada. Todos los endpoints diagnósticos Go también requieren credencial. La operación interna admite únicamente principal de servicio `game` y audiencia del destino; eso no habilita a AI para escribir Game.

`testkit` firma fixtures únicamente desde tests; los ejecutables no importan ese paquete. La web no recibe claves de servicio/DB y su smoke utiliza dobles controlados con credenciales efímeras en memoria. Secure/HttpOnly/SameSite se verifican como cabeceras transportadas, sin afirmar enforcement del navegador o sesión real. No publicar estos endpoints como producto.

El runner integral deriva claves de usuario distintas por audiencia y entrega a cada proceso solo su clave. AI no recibe la clave de usuario Game ni DSN/PG*. La clave S2S continúa compartida en este experimento y no acredita aislamiento frente al compromiso de un servicio. `/internal/check/:service` exige principal Game, usa destinos loopback cerrados y timeout; permitió probar el trayecto saliente y una caída real de AI.

## Repetición de la verificación integral

Después del fixture SQL anterior, ejecutar secuencialmente, sin otro harness que reinicie la DB o reconstruya sus tablas:

```powershell
./spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase Web
./spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase Go
./spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase LinuxRace
./spikes/infra-feasibility/tests/notificationprobe/Run.ps1
./spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase NotificationRace
./spikes/infra-feasibility/tests/Verify-IntegralRuntime.ps1
./spikes/infra-feasibility/tests/Verify-Documents.ps1
./spikes/infra-feasibility/tests/Verify-Exposure.ps1
```

Cada build/runtime crea un directorio de ejecución nuevo. El runner runtime compila procesos, inserta IDs sintéticos nuevos, mata/reinicia solo su Game y AI, y reinicia solo PostgreSQL del proyecto/volumen esperado. `-SkipDatabaseRestart` es diagnóstico parcial y no sirve para aceptar DB-07. Al terminar cierra sus procesos y restaura el entorno; deja PostgreSQL y su volumen disponibles. No ejecutar junto con usuarios de 18100–18103 o fixtures destructivos.

Las fases race usan una imagen Go/GCC identificada y contenedores efímeros, sin instalar compilador global. Windows race nativo permanece no ejecutado. [Índice de builds](evidence/integral-builds-summary.md), [notificaciones simuladas](tests/notificationprobe/README.md) e [informe integral](INTEGRAL_REPORT.md) detallan comandos, evidencia final y fallos históricos conservados. Los controles documentales/de exposición no validan comportamiento del proveedor ni ausencia universal de vulnerabilidades.

## PostgreSQL y notificaciones

Compose propone una instancia exclusivamente local, memoria 512 MiB/una CPU y puerto loopback 15435. Dos bases con propietarios/contraseñas distintos; `CONNECT` público revocado. AI no dispone de DSN. Pools de aplicación: máximo dos conexiones, mínimo cero. Migraciones reversibles se ejecutan explícitamente en el fixture de integración, nunca al arrancar una Function.

El fixture reconstruye tablas **sintéticas del spike**: no ejecutarlo contra bases preexistentes ni de producto. `Store.Apply` bloquea el agregado, comprueba propietario, busca recibo antes de validar revisión, confirma recibo+revisión+efecto en una transacción y luego intenta la notificación con límite de tiempo. El doble de notificación es un adaptador local, no el SDK ni una app Ably. No hay outbox, entrega exactamente una vez ni recuperación garantizada sin REST disponible. Ver casos locales verificados en [TEST_MATRIX](TEST_MATRIX.md).

## Recursos y retirada

No se creó ningún recurso externo. Después de la reparación autorizada se crearon únicamente el contenedor PostgreSQL, red y volumen locales del Compose `toguz-infra-spike`. Los directorios con sockets averiados se preservaron con nombres de cuarentena y Docker fue actualizado mediante su actualizador oficial; no hubo factory reset, prune ni eliminación de imágenes/volúmenes existentes. Los procesos HTTP efímeros de tests se cierran automáticamente. `.local/` conserva credenciales sintéticas/cachés/builds locales ignorados; no son cuentas reales.

El Compose dedicado ya está creado y se conserva para continuar las pruebas. Su retirada explícita sería `docker compose --env-file <spike>/.local/compose.env -f <spike>/infra/compose.yaml down --volumes`, tras comprobar nombre `toguz-infra-spike` y recursos propios. No se ejecutó la retirada ni debe aplicarse contra otros proyectos. `.local/` puede retirarse solo después de conservar la evidencia necesaria y comprobar su ruta absoluta dentro de este spike.

Los cuatro `vercel.json` deshabilitan despliegues Git para estos contextos; no se ha aplicado configuración externa. Eso no prueba ausencia de integraciones preexistentes con otra raíz. **Push y Draft PR retenidos** hasta demostrar que no activan despliegues no autorizados. Ver [REPORT](REPORT.md) para auditoría, costes, bloqueos y gate remoto.
