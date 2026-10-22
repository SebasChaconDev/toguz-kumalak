# Verificación independiente de builds — Gate 3.5

Fecha observada: 2026-10-22 UTC. Rama: `spike/infra-feasibility-01`. HEAD base: `cf5d3973adc977f0ec09c738cbe6cfc078883593`. La pasada final incluye los cambios locales de Game realizados durante Gate 3.5; el HEAD por sí solo no identifica todo el código probado. Los [hashes de fuentes Go/SQL y manifiestos](integral-builds/20261022T182553Z-4d2760b3d4354a0b822fe531f4edcdb4/source-hashes.json) identifican el estado observado después de la pasada Linux satisfactoria.

## Resultados observados

| Control | Resultado | Evidencia |
|---|---|---|
| Cinco módulos: test, vet y build en Windows/amd64 | PASS: los quince comandos terminaron con código 0 | [Resultados Go](integral-builds/20261022T182541Z-83d3e9ea66aa4fc5bdd97b83b3f2032c/results.json) |
| Identity, Game y AI: build frío del servidor y del paquete API | PASS: seis comandos con código 0; tres contextos y pares de cachés nuevos | Mismo registro Go; logs `*-isolated-*` |
| Frontera del motor compartido | PASS: Game y AI incluyen el mismo `engine.go`; Identity no incluye ni importa engine | Inventarios `*-isolated-modules.txt`, `*-binary-modules.txt` y `*-dependencies.txt` del registro Go |
| Cinco módulos: detector de carreras Linux/amd64 | PASS: cinco `go test -race` con código 0; contenedor eliminado automáticamente | [Log Linux](integral-builds/20261022T182553Z-4d2760b3d4354a0b822fe531f4edcdb4/linux-race.txt) |
| Detector de carreras Windows nativo | NOT_RUN: `CC=gcc`, ejecutable gcc ausente del PATH | [Compilador Windows](integral-builds/20261022T182541Z-83d3e9ea66aa4fc5bdd97b83b3f2032c/windows-compiler.json) |
| Next: npm ci, test, build y smoke | PASS: cuatro comandos con código 0, caché npm y node_modules nuevos | [Resultados web](integral-builds/20261022T182305Z-31b332eb56cf4ddfbab35a1d970a9b7b/results.json) |
| Next: lint | NOT_RUN: package.json original no declara script lint | Mismo registro web, marcado NOT_DEFINED |
| Lockfile web | PASS: idéntico tras npm ci; SHA-256 `54C056FB7252D33D29CA382DCB1D038C953D64BD61F07162CC302FA2CAD7263F` | Mismo registro web |

Los módulos comprobados fueron `modules/engine`, `modules/probe`, `services/identity`, `services/game` y `services/ai`. Los builds aislados copiaron exclusivamente el servicio correspondiente, probe y engine cuando era necesario, incluyendo los SQL embebidos. No se copió el repositorio raíz ni un go.work. Antes de cada build frío no existían sus directorios GOCACHE y GOMODCACHE. Los inventarios de módulos y del binario registran las dependencias efectivamente utilizadas.

Game y AI requieren engine `v0.1.0` con reemplazo local `../../modules/engine`, sin acreditar un tag publicado. El SHA-256 observado de `engine.go` en ambos contextos es `BD0D2F40296E32B6C832DEA9B22C0CB1B957E428D2DC01E20456C303AB14C973`. Identity se compiló sin copiar engine; su lista de dependencias tampoco lo contiene.

## Comandos reproducibles

Desde la raíz del repositorio:

```powershell
pwsh -NoProfile -File spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase Web
pwsh -NoProfile -File spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase Go
pwsh -NoProfile -File spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase LinuxRace
```

La fase Web necesita libres los puertos loopback 18100, 18101 y 18102. Su smoke finalizó antes de iniciar los servicios Go de integración. El runner crea un UUID nuevo por invocación y conserva logs/resultados en `evidence/integral-builds/<run-id>`, y cachés y artefactos privados ignorados por Git en `.local/integral-builds/<run-id>`. Restaura las variables de entorno que modifica y termina con error cuando un comando devuelve código distinto de cero.

Comandos Go ejecutados por módulo, con `GOWORK=off`, `GOENV=off`, `GOTOOLCHAIN=local` y `CGO_ENABLED=0` en Windows:

```text
go test -mod=readonly -count=1 ./...
go vet -mod=readonly ./...
go build -mod=readonly ./...
```

Cada servicio aislado ejecutó `go build -mod=readonly -trimpath -o <contexto>/<servicio>.exe ./cmd/server`, `go build -mod=readonly ./api`, `go list -m -json all`, `go version -m <binario>` y `go list -deps ./...`. Los argumentos y tiempos exactos figuran en results.json, con la ruta privada del workspace sustituida por `<SPIKE_ROOT>`.

Linux ejecutó `go test -mod=readonly -race -count=1 ./...` en cada uno de los cinco módulos, con CGO_ENABLED=1, GOWORK=off, GOENV=off, GOTOOLCHAIN=local y caches nuevas dentro del contenedor. El comando Docker completo está en su results.json. Se usó `--rm`, máximo 2 CPU y 4 GiB, sin publicar puertos y con montajes de módulos y servicios de solo lectura. No se montó `.local`.

## Versiones y límites

Go observado: `go1.26.2 windows/amd64` y `go1.26.2 linux/amd64`. Node: `v24.13.0`. npm: `11.6.2`. Next: `16.4.2`; React y React DOM declarados `19.3.1`.

Imagen oficial Linux: `golang:1.26.2-bookworm`, digest `sha256:47ce5636e9936b2c5cbf708925578ef386b4f8872aec74a67bd13a627d242b19`. Compilador: `gcc (Debian 12.2.0-14+deb12u1) 12.2.0`. [Identidad de imagen observada](integral-builds/20261022T182553Z-4d2760b3d4354a0b822fe531f4edcdb4/docker-go-image.txt).

El smoke web ejecutó ocho peticiones medidas y nueve intentos de readiness, y revisó nueve archivos JavaScript de cliente; sus dobles HTTP son Node locales. Prueba forwarding de Authorization, conservación de 401, rewrites, atributos Set-Cookie y ausencia de los marcadores sintéticos generados en HTML, logs capturados y JavaScript de cliente. No prueba Identity/OIDC real, enforcement de cookies por navegador, cookies reales de Go ni Vercel. [Detalle y límites del smoke](integral-builds/20261022T182305Z-31b332eb56cf4ddfbab35a1d970a9b7b/web-smoke.json).

Los quince comandos Go ordinarios no usan el tag integration y por sí solos no prueban PostgreSQL. No se atribuye el resultado Linux race a Windows. Estos resultados acreditan POC-01 local y aportan controles locales a POC-02/06/07; no acreditan empaquetado ni despliegue real en Vercel, ni validación remota.

## Pasadas iniciales conservadas

Las pasadas [Go inicial](integral-builds/20261022T182319Z-e9daa927b65e442e942b535505904548/results.json) y [Linux inicial](integral-builds/20261022T182400Z-a2166bbbc0ae48d08a844295cba0eb2c/results.json) fallaron realmente en `TestInternalProbe`, con `valid upstream rejected`. Coincidieron con el stub temporal de la prueba TDD del coordinador: `checkInternal` devolvía siempre `ErrInternalUnavailable`. El coordinador confirmó ese estado RED intencionado y posteriormente implementó la llamada S2S. Se conservaron ambos fallos y se repitieron completas las fases Go y Linux race con cachés nuevas sobre la implementación corregida. No fue un fallo atribuido al entorno ni se eliminó la prueba.

## Regresión adicional: notificación con PostgreSQL y race

Después de la señal de PostgreSQL healthy del coordinador, se ejecutó:

```powershell
pwsh -NoProfile -File spikes/infra-feasibility/tests/Verify-IntegralBuilds.ps1 -Phase NotificationRace
```

Resultado observado: PASS, código 0, sin carreras detectadas; 46.42 s de contenedor y 1.329 s de paquete. [Log completo](integral-builds/20261022T183118Z-e39b51d8d4fe405ea6da88bd40038dae/notification-linux-race.txt), [comandos exactos](integral-builds/20261022T183118Z-e39b51d8d4fe405ea6da88bd40038dae/results.json), [conteo verificable](integral-builds/20261022T183118Z-e39b51d8d4fe405ea6da88bd40038dae/notification-summary.json) y [hashes del harness](integral-builds/20261022T183118Z-e39b51d8d4fe405ea6da88bd40038dae/source-hashes.json).

El comando del módulo `tests/notificationprobe` fue `go test -mod=readonly -tags=integration -race -count=1 -json ./...`. Los eventos JSON registran ocho tests principales y ocho subtests aprobados: siete escenarios de notificación y ocho escenarios de capabilities, quince escenarios explícitos. El evento de aprobación del padre de capabilities explica el decimosexto evento PASS de test.

Se observaron entrega posterior al commit, aviso perdido, duplicados, desorden, respuesta HTTP perdida, reinicio del broker y broker sin respuesta, además de rechazos de publicación por cliente, recurso ajeno, expiración, firma inválida y credencial ausente. Los snapshots antes/después, contadores de efectos y receipts, auditorías de commit anterior a publicación y recuperación REST se incluyen en las líneas `CASE` del log. Se utilizaron PostgreSQL real y Game real (`app.New` + Store); el broker HTTP y sus capabilities siguen siendo un doble local, sin validar Ably real. Aporta regresión Linux a POC-04/05/07/08 locales; no completa por sí sola sus otros criterios.

Se mantuvo la misma imagen oficial y digest. El contenedor tuvo límites de 2 CPU/4 GiB, cachés vacías, GOWORK=off, CGO_ENABLED=1 y tres montajes de solo lectura para módulos, servicios y harness. Se conectó a la red ya existente `toguz-infra-spike_default`. Un envfile privado, ignorado por Git, cambió exclusivamente el destino loopback del PostgreSQL sintético a `postgres:5432` dentro de esa red; las credenciales no se imprimieron ni se montó `.local`. El contenedor se eliminó mediante `--rm`; la inspección posterior no mostró contenedores `toguz-integral-*` en ejecución. No se detuvieron ni eliminaron contenedores o volúmenes existentes.

El escaneo de los valores reales de las credenciales sintéticas contra todos los archivos de `evidence/integral-builds` no encontró coincidencias; no imprimió esos valores. La sintaxis PowerShell del runner se analizó sin errores. No se modificó código de aplicación ni configuración global en esta auditoría de builds. El harness `processprobe` no se sometió a race en Linux: no contiene tests funcionales y esa invocación solo comprobaría compilación; su vet/build Windows corresponde a la verificación separada del coordinador.
