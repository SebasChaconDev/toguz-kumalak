# REG — Registro canónico de decisiones

ID: REG · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner: C.

## Procedencia y autoridad

Las **246 decisiones** siguientes reproducen literalmente las formulaciones del Anexo A del [maestro archivado](sources/TOGUZ_PROMPT_MAESTRO_CODEX.md), consolidado el 2026-10-21. La fuente declara su aprobación en conversación. Este registro generado está pendiente de revisión humana; no acredita software, pruebas ni despliegue. La asignación canónica procede del inventario aprobado por el propietario con «APRUEBO INVENTARIO DOCUMENTAL» y se desarrolla en [trazabilidad](traceability.md).

La copia archivada conserva SHA-256 `28C24C498829FCFAD410F07AAA6BD036D5A33382294401E11E86D486E4ECBBEF`. No se encontró registro separado para contrastar. Ante discrepancia con la copia, preservar ambas formulaciones, abrir [conflicto](conflicts.md) y solicitar resolución; nunca corregir requisitos silenciosamente.

## Estados independientes

- Requisito fuente: `APROBADO`.
- Factibilidad no demostrada: `POR_VERIFICAR`.
- Alternativa o detalle nuevo: `PROPUESTA`, sin autorización para implementarlo.
- Documento generado: `EN_REVISIÓN`, hasta aceptación del propietario.

## Sustituciones vigentes

Go/Gin sustituye Node/Fastify; Vercel sustituye Railway; Ably sustituye Socket.IO y WebSocket propio obligatorio; sqlc/pgx/v5/Goose sustituyen Drizzle; Identity/Game/AI sustituyen el monolito; objetivo 0 USD sustituye techo 20 USD. PRE-03 amplía tablas y práctica sin habilitar undo ordinario ni Elo retroactivo. El realismo estilizado de PRE-04.A1 matiza el fotorrealismo general. Local/HTML no implica offline ni otro motor.

Los IDs son los originales `PRE-NN.Xn`; no se reconstruyen opciones de cuestionarios ni votos fuera del anexo. Las ADR del lote1 documentan consecuencias de decisiones existentes y permanecen EN_REVISIÓN. Ninguna spec futura se considera redactada por tener una asignación aquí.

## Resumen

| PRE | Decisiones | Vista contextual |
|---|---:|---|
| PRE-01 | 17 | [PRE-01](preliminares/PRE-01.md) |
| PRE-02 | 18 | [PRE-02](preliminares/PRE-02.md) |
| PRE-03 | 22 | [PRE-03](preliminares/PRE-03.md) |
| PRE-04 | 25 | [PRE-04](preliminares/PRE-04.md) |
| PRE-05 | 24 | [PRE-05](preliminares/PRE-05.md) |
| PRE-06 | 20 | [PRE-06](preliminares/PRE-06.md) |
| PRE-07 | 24 | [PRE-07](preliminares/PRE-07.md) |
| PRE-08 | 24 | [PRE-08](preliminares/PRE-08.md) |
| PRE-09 | 24 | [PRE-09](preliminares/PRE-09.md) |
| PRE-10 | 24 | [PRE-10](preliminares/PRE-10.md) |
| PRE-11 | 24 | [PRE-11](preliminares/PRE-11.md) |

<a id="pre-01"></a>

## PRE-01 — 17 decisiones

<a id="pre-01-a1"></a>

- **PRE-01.A1** — TypeScript para frontend y Go para backend; generar y verificar contratos tipados entre ambos lenguajes.

  Destino canónico: [01.01](index.md#spec-01-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-a2"></a>

- **PRE-01.A2** — Next.js, React y TypeScript para la aplicación web.

  Destino canónico: [01.01](index.md#spec-01-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-a3"></a>

- **PRE-01.A3** — Gin como framework HTTP principal de los microservicios Go.

  Destino canónico: [01.01](index.md#spec-01-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-a4"></a>

- **PRE-01.A4** — Monorepositorio: pnpm para proyectos TypeScript y go.mod/go.work para proyectos Go; módulos compartidos con dependencias controladas.

  Destino canónico: [01.01](index.md#spec-01-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-b1"></a>

- **PRE-01.B1** — Sustituida la elección inicial de Socket.IO por Ably Free para notificaciones distribuidas; no crear gateway Node.js ni exigir servidores WebSocket persistentes propios.

  Destino canónico: [06.02](index.md#spec-06-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-b2"></a>

- **PRE-01.B2** — PostgreSQL como base de datos relacional; Neon Free es proveedor principal previsto.

  Destino canónico: [01.02](index.md#spec-01-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-b3"></a>

- **PRE-01.B3** — Sustituir Drizzle por SQL tipado con sqlc + pgx/v5 y migraciones SQL con Goose, nativos de Go.

  Destino canónico: [05.03](index.md#spec-05-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-b4"></a>

- **PRE-01.B4** — Guardar movimientos aceptados y estado actualizado de forma atómica y transaccional.

  Destino canónico: [05.03](index.md#spec-05-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-b5"></a>

- **PRE-01.B5** — Permitir invitados que juegan de inmediato y cuentas opcionales para identidad persistente y clasificación.

  Destino canónico: [07.01](index.md#spec-07-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-c1"></a>

- **PRE-01.C1** — Sustituir Railway por Vercel como proveedor principal; cuatro proyectos: frontend, Identity, Game y AI; factibilidad pendiente de pruebas de despliegue.

  Destino canónico: [01.02](index.md#spec-01-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-c2"></a>

- **PRE-01.C2** — Sustituir el presupuesto máximo inicial de 20 USD por objetivo operativo de 0 USD/mes; prohibido habilitar recursos de pago sin permiso explícito.

  Destino canónico: [01.03](index.md#spec-01-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-c3"></a>

- **PRE-01.C3** — Desarrollo local con servicios Go/Next.js nativos cuando sea práctico y PostgreSQL/infraestructura auxiliar en Docker Compose.

  Destino canónico: [01.01](index.md#spec-01-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-c4"></a>

- **PRE-01.C4** — Separación entre desarrollo, previews/entorno de integración y producción; no compartir datos ni credenciales productivas.

  Destino canónico: [01.02](index.md#spec-01-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-d1"></a>

- **PRE-01.D1** — GitHub Actions para tipado, formato/lint, tests, build, integración y controles de calidad.

  Destino canónico: [09.03](index.md#spec-09-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-d2"></a>

- **PRE-01.D2** — Observabilidad avanzada con OpenTelemetry, Prometheus/Grafana principalmente localmente y telemetría productiva según cuotas verificadas.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-d3"></a>

- **PRE-01.D3** — Crecer desde una instancia por microservicio con medidas reales; Redis es opcional y se incorporará solo si se justifica por capacidad/coordinación, nunca preventivamente obligatorio.

  Destino canónico: [01.03](index.md#spec-01-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-01-d4"></a>

- **PRE-01.D4** — Repositorio de portafolio público; elegir licencia expresamente antes de publicar derechos de reutilización.

  Destino canónico: [02.03](index.md#spec-02-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02"></a>

## PRE-02 — 18 decisiones

<a id="pre-02-a1"></a>

- **PRE-02.A1** — Jugar localmente con dos personas en un mismo dispositivo y contra IA básica; incluir modo práctica separado.

  Destino canónico: [02.01](index.md#spec-02-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-a2"></a>

- **PRE-02.A2** — Partidas online en salas privadas por enlace o código de invitación; sin emparejamiento abierto obligatorio.

  Destino canónico: [02.01](index.md#spec-02-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-a3"></a>

- **PRE-02.A3** — Partidas online en tiempo real con ambos participantes conectados durante el juego ordinario y política de recuperación al desconectar.

  Destino canónico: [02.01](index.md#spec-02-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-a4"></a>

- **PRE-02.A4** — Sin reloj competitivo ni control de tiempo por partida en el MVP.

  Destino canónico: [02.01](index.md#spec-02-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-b1"></a>

- **PRE-02.B1** — Los invitados pueden jugar local/online; la cuenta registrada habilita historial personal persistente y clasificación.

  Destino canónico: [02.03](index.md#spec-02-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-b2"></a>

- **PRE-02.B2** — Alias personalizable y un identificador temporal verificable para invitados.

  Destino canónico: [02.03](index.md#spec-02-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-b3"></a>

- **PRE-02.B3** — Interfaz en español e inglés desde el MVP.

  Destino canónico: [02.02](index.md#spec-02-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-b4"></a>

- **PRE-02.B4** — Tutorial interactivo opcional y consulta completa del reglamento.

  Destino canónico: [08.04](index.md#spec-08-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-c1"></a>

- **PRE-02.C1** — Resaltar casillas legales, turno activo y vista previa básica de movimientos con datos autorizados por Go.

  Destino canónico: [03.03](index.md#spec-03-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-c2"></a>

- **PRE-02.C2** — En partidas ordinarias no existe deshacer; se añadió expresamente un modo práctica reversible con deshacer/rehacer.

  Destino canónico: [08.03](index.md#spec-08-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-c3"></a>

- **PRE-02.C3** — Reconexión automática y partida online recuperable inicialmente durante 24 h; vencido el plazo, se archiva como inconclusa.

  Destino canónico: [06.03](index.md#spec-06-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-c4"></a>

- **PRE-02.C4** — Abandono voluntario y revancha consensuada; AMPLIADA en PRE-03 para incluir propuestas de tablas por mutuo acuerdo y reclamación reglamentaria por repetición, además de tablas por resultado.

  Destino canónico: [03.04](index.md#spec-03-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-d1"></a>

- **PRE-02.D1** — Historial de partidas terminadas y reproducción básica de sus movimientos.

  Destino canónico: [03.04](index.md#spec-03-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-d2"></a>

- **PRE-02.D2** — Perfil con alias, partidas, victorias, derrotas, empates e historial.

  Destino canónico: [02.03](index.md#spec-02-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-d3"></a>

- **PRE-02.D3** — Guardado automático de la última partida local en el navegador y posibilidad de reanudarla; no se exige funcionamiento offline.

  Destino canónico: [02.02](index.md#spec-02-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-e1"></a>

- **PRE-02.E1** — Three.js como motor visual principal del tablero tridimensional.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-e2"></a>

- **PRE-02.E2** — Estética de tablero físico premium/fotorrealista inspirada en rdomino.com, sin copiar archivos, diseños distintivos, modelos, sonidos ni código.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-02-e3"></a>

- **PRE-02.E3** — MVP ampliado: IA básica, clasificación competitiva y estadísticas avanzadas; torneos, chat, espectadores, IA avanzada, pagos y funciones sociales extensas quedan pospuestos.

  Destino canónico: [02.01](index.md#spec-02-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03"></a>

## PRE-03 — 22 decisiones

<a id="pre-03-a1"></a>

- **PRE-03.A1** — Detectar triple repetición de una misma posición reglamentaria y permitir reclamación de tablas; no empate automático por tercera aparición salvo que la regla formal lo exija.

  Destino canónico: [03.04](index.md#spec-03-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-a2"></a>

- **PRE-03.A2** — Permitir propuestas de tablas por mutuo acuerdo; modifica PRE-02 C4.

  Destino canónico: [03.04](index.md#spec-03-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-a3"></a>

- **PRE-03.A3** — Tuzdyk se crea automáticamente al cumplir las reglas, con señalización, explicación y animación contextual.

  Destino canónico: [03.01](index.md#spec-03-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-a4"></a>

- **PRE-03.A4** — Blancas empiezan; asignación online aleatoria de colores y elección de bando en local/IA.

  Destino canónico: [03.02](index.md#spec-03-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-b1"></a>

- **PRE-03.B1** — Estado lógico con tablero fijo [2][9], kazanes, posición/propiedad de tuzdyks, turno y versión, con convenio documentado de índices y recorrido circular.

  Destino canónico: [03.02](index.md#spec-03-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-b2"></a>

- **PRE-03.B2** — Motor puro determinista que recibe estado + movimiento y devuelve nuevo estado + eventos + error, sin mutar el original ni tocar persistencia/red.

  Destino canónico: [03.03](index.md#spec-03-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-b3"></a>

- **PRE-03.B3** — Comandos autorizados por identidad de sesión con casilla origen, versión esperada e identificador idempotente; no recibir ni confiar en tablero completo de cliente.

  Destino canónico: [03.03](index.md#spec-03-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-b4"></a>

- **PRE-03.B4** — Versionar formato de estado, reglamento y registros de movimientos para reproducción estable.

  Destino canónico: [03.02](index.md#spec-03-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-b5"></a>

- **PRE-03.B5** — Go determina movimientos legales y previsualizaciones; frontend no duplica reglamento.

  Destino canónico: [03.03](index.md#spec-03-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-c1"></a>

- **PRE-03.C1** — Representar tuzdyks y sus propietarios explícitamente; las piedras que caen allí incrementan inmediatamente el kazán correspondiente.

  Destino canónico: [03.01](index.md#spec-03-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-c2"></a>

- **PRE-03.C2** — Generar eventos semánticos de siembra, captura, tuzdyk, turnos y fin para que Three.js pueda animarlos sin conocer reglas.

  Destino canónico: [03.03](index.md#spec-03-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-c3"></a>

- **PRE-03.C3** — Rechazar movimientos ilegales con errores tipados y sin alterar el estado original.

  Destino canónico: [03.03](index.md#spec-03-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-c4"></a>

- **PRE-03.C4** — Guardar estado de resultado separado de su causa (reglas, abandono, empate, desconexión/inconclusa, etc.).

  Destino canónico: [03.04](index.md#spec-03-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-d1"></a>

- **PRE-03.D1** — Modo práctica con control libre y escenarios educativos; sin efectos competitivos.

  Destino canónico: [08.03](index.md#spec-08-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-d2"></a>

- **PRE-03.D2** — Historial de práctica con deshacer/rehacer por estados; al jugar después de retroceder, descartar la rama posterior.

  Destino canónico: [08.03](index.md#spec-08-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-d3"></a>

- **PRE-03.D3** — IA básica con Negamax/Minimax y poda alfa-beta, heurística explícita.

  Destino canónico: [08.01](index.md#spec-08-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-d4"></a>

- **PRE-03.D4** — AI Service ejecuta Go en Vercel con límites estrictos de cálculo, cancelación y respuesta legal.

  Destino canónico: [08.01](index.md#spec-08-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-d5"></a>

- **PRE-03.D5** — Dos dificultades: fácil y normal, diferenciadas y evaluadas.

  Destino canónico: [08.01](index.md#spec-08-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-e1"></a>

- **PRE-03.E1** — Corpus de posiciones del reglamento formal y contraste independiente con implementaciones de referencia.

  Destino canónico: [03.01](index.md#spec-03-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-e2"></a>

- **PRE-03.E2** — Pruebas unitarias, de propiedades, fuzzing, simulaciones e invariantes; en particular total de 162 piedras conservado.

  Destino canónico: [03.02](index.md#spec-03-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-e3"></a>

- **PRE-03.E3** — Reproducción basada en movimientos/estados versionados y recálculo por motor; evitar persistir instrucciones 3D como fuente de verdad.

  Destino canónico: [03.04](index.md#spec-03-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-03-e4"></a>

- **PRE-03.E4** — Benchmarks y objetivos cuantitativos después de obtener corrección funcional y mediciones reproducibles.

  Destino canónico: [09.04](index.md#spec-09-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04"></a>

## PRE-04 — 25 decisiones

<a id="pre-04-a1"></a>

- **PRE-04.A1** — Realismo físico estilizado de alta calidad, con acabados creíbles y estética premium, sin fotorrealismo extremo obligatorio.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-a2"></a>

- **PRE-04.A2** — Maderas nobles cálidas/oscuras, cavidades talladas, piedras naturales pulidas y detalles metálicos discretos.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-a3"></a>

- **PRE-04.A3** — Referencias culturales kazajas verificadas, aplicadas con respeto y sin ornamentos inventados.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-a4"></a>

- **PRE-04.A4** — Three.js representa tablero; menús, perfiles, HUD, estadísticas y controles accesibles en React/HTML/CSS.

  Destino canónico: [04.01](index.md#spec-04-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-b1"></a>

- **PRE-04.B1** — Three.js + React Three Fiber + Drei.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-b2"></a>

- **PRE-04.B2** — Geometría procedural junto con Blender y exportación GLB/glTF para recursos detallados.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-b3"></a>

- **PRE-04.B3** — InstancedMesh para las 162 piedras, con disposición determinista dependiente del estado de juego.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-b4"></a>

- **PRE-04.B4** — Cámara oblicua ajustable con zoom y rotación controlados, cambios de orientación disponibles.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-b5"></a>

- **PRE-04.B5** — Materiales PBR, iluminación ambiental cuidada, sombras suaves y posprocesado moderado adaptativo.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-c1"></a>

- **PRE-04.C1** — Selección y confirmación separadas para evitar jugadas accidentales.

  Destino canónico: [04.03](index.md#spec-04-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-c2"></a>

- **PRE-04.C2** — Animaciones de siembra secuenciales, acelerables y con opción de saltar sin alterar estado lógico.

  Destino canónico: [04.03](index.md#spec-04-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-c3"></a>

- **PRE-04.C3** — Animaciones de Three.js consumen eventos semánticos del motor Go, no infieren reglas.

  Destino canónico: [04.03](index.md#spec-04-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-c4"></a>

- **PRE-04.C4** — Sonidos de materiales y eventos opcionales, discretos, configurables, sin reproducción invasiva.

  Destino canónico: [04.03](index.md#spec-04-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-d1"></a>

- **PRE-04.D1** — Portada, elección de modo, sala, tablero, resultados, tutorial, perfil, historial y clasificación como navegación clara por pantallas.

  Destino canónico: [04.01](index.md#spec-04-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-d2"></a>

- **PRE-04.D2** — Móvil vertical y horizontal, adaptación automática de cámara/controles y sugerencia de horizontal cuando beneficie la legibilidad.

  Destino canónico: [04.01](index.md#spec-04-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-d3"></a>

- **PRE-04.D3** — Orientación de cámara adaptada para juego local, giro opcional entre turnos y opción de cámara fija.

  Destino canónico: [04.01](index.md#spec-04-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-d4"></a>

- **PRE-04.D4** — Accesibilidad de teclado, datos de casillas textuales, indicadores no basados solo en color y alternativa HTML interactiva.

  Destino canónico: [04.01](index.md#spec-04-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-d5"></a>

- **PRE-04.D5** — Tutorial opcional con resaltado contextual de casillas y escenarios dentro del tablero 3D.

  Destino canónico: [08.04](index.md#spec-08-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-e1"></a>

- **PRE-04.E1** — Objetivo 60 FPS escritorio y 30 FPS móvil compatible, contrastados con dispositivos de referencia; no afirmarlos como garantía.

  Destino canónico: [04.04](index.md#spec-04-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-e2"></a>

- **PRE-04.E2** — Calidades automática, baja y alta con ajuste de resolución, sombras y efectos.

  Destino canónico: [04.04](index.md#spec-04-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-e3"></a>

- **PRE-04.E3** — Carga progresiva/diferida de 3D, materiales reutilizados y recursos optimizados/comprimidos.

  Destino canónico: [04.04](index.md#spec-04-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-e4"></a>

- **PRE-04.E4** — Vista alternativa HTML plenamente jugable si WebGL está ausente o falla.

  Destino canónico: [04.04](index.md#spec-04-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-f1"></a>

- **PRE-04.F1** — Pruebas visuales automatizadas, referencia de escenas, revisiones multiplataforma y medición de rendimiento.

  Destino canónico: [09.02](index.md#spec-09-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-f2"></a>

- **PRE-04.F2** — Recursos originales o debidamente licenciados, con inventario de procedencia; no copiar recursos de RDomino.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-04-f3"></a>

- **PRE-04.F3** — Aprobación visual incremental: geometría → materiales/cámara → interacción → animaciones → móvil → acabado.

  Destino canónico: [04.02](index.md#spec-04-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05"></a>

## PRE-05 — 24 decisiones

<a id="pre-05-a1"></a>

- **PRE-05.A1** — SUSTITUIDO el monolito: exactamente tres microservicios Go/Gin independientes (Identity, Game, AI) y frontend Next.js; autonomía por dominio.

  Destino canónico: [05.01](index.md#spec-05-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-a2"></a>

- **PRE-05.A2** — Arquitectura hexagonal pragmática dentro de cada microservicio: transporte, aplicación, dominio y adaptadores.

  Destino canónico: [05.01](index.md#spec-05-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-a3"></a>

- **PRE-05.A3** — Dependencias inyectadas explícitamente mediante constructores e interfaces pequeñas cercanas al consumidor.

  Destino canónico: [05.01](index.md#spec-05-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-a4"></a>

- **PRE-05.A4** — Cuatro proyectos Vercel, despliegues independientes desde el monorepositorio, módulo Go autónomo por servicio.

  Destino canónico: [05.01](index.md#spec-05-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-b1"></a>

- **PRE-05.B1** — REST versionado bajo /api/v1, consultas de recursos y comandos explícitos.

  Destino canónico: [05.02](index.md#spec-05-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-b2"></a>

- **PRE-05.B2** — OpenAPI 3.1 como fuente de verdad de HTTP, clientes/modelos generados en Go y TypeScript, versiones fijadas.

  Destino canónico: [05.02](index.md#spec-05-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-b3"></a>

- **PRE-05.B3** — Errores estructurados compatibles con RFC 9457 (Problem Details), códigos de dominio, requestId y mensajes seguros.

  Destino canónico: [05.02](index.md#spec-05-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-b4"></a>

- **PRE-05.B4** — APIs por recursos/dominios (games, players, practice, leaderboard), paginación y sin exponer tablas directamente.

  Destino canónico: [05.02](index.md#spec-05-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-c1"></a>

- **PRE-05.C1** — Pipeline de jugada: autenticar → autorizar → versión → motor Go → persistencia → respuesta autoritativa.

  Destino canónico: [05.03](index.md#spec-05-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-c2"></a>

- **PRE-05.C2** — Bloqueo transaccional por partida, comprobación de versión e identificador único idempotente.

  Destino canónico: [05.03](index.md#spec-05-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-c3"></a>

- **PRE-05.C3** — Comando aceptado, movimiento, efectos y estado confirmados dentro de la misma transacción.

  Destino canónico: [05.03](index.md#spec-05-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-c4"></a>

- **PRE-05.C4** — Simulación local/práctica mediante endpoints sin autoridad competitiva; sin motor paralelo TypeScript.

  Destino canónico: [02.02](index.md#spec-02-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-d1"></a>

- **PRE-05.D1** — Propiedad de datos separada: Identity DB, Game DB, AI stateless; no consultas SQL cruzadas entre servicios.

  Destino canónico: [05.01](index.md#spec-05-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-d2"></a>

- **PRE-05.D2** — Estado actual compacto versionado en JSONB validado por Go; movimientos normalizados y columnas de consulta relacionales.

  Destino canónico: [05.03](index.md#spec-05-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-d3"></a>

- **PRE-05.D3** — pgxpool limitado, Neon pooled para aplicación cuando proceda y conexión directa para migraciones.

  Destino canónico: [05.03](index.md#spec-05-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-d4"></a>

- **PRE-05.D4** — Goose con migraciones SQL versionadas/controladas, sin migrar automáticamente al arrancar Functions.

  Destino canónico: [05.03](index.md#spec-05-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-e1"></a>

- **PRE-05.E1** — Eventos Ably ligeros con gameId, versión, tipo y datos mínimos; sin instrucciones 3D.

  Destino canónico: [05.04](index.md#spec-05-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-e2"></a>

- **PRE-05.E2** — Confirmar PostgreSQL antes de publicar en Ably; reconciliar por REST tras reconexión/foco y verificaciones acotadas; documentar riesgo de aviso perdido sin outbox.

  Destino canónico: [05.04](index.md#spec-05-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-e3"></a>

- **PRE-05.E3** — Tokens Ably temporales de suscripción limitados a canales de partida; publicación oficial desde Game.

  Destino canónico: [05.04](index.md#spec-05-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-e4"></a>

- **PRE-05.E4** — JSON schemas versionados para eventos y validación en Go/TypeScript.

  Destino canónico: [05.04](index.md#spec-05-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-f1"></a>

- **PRE-05.F1** — OpenTelemetry y trazas de API, servicios, SQL y Ably con referencias correlacionadas.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-f2"></a>

- **PRE-05.F2** — Pruebas con PostgreSQL real, HTTP/Gin, contratos, carreras/consistencia y doble de Ably.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-f3"></a>

- **PRE-05.F3** — Autorización server-side, límites de tasa, validación y antiabuso en salas, invitados e IA.

  Destino canónico: [07.02](index.md#spec-07-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-05-f4"></a>

- **PRE-05.F4** — Prueba temprana obligatoria de Vercel Go/Gin + Neon + Ably + transacción antes de iniciar implementación completa; aún no ejecutada.

  Destino canónico: [01.02](index.md#spec-01-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06"></a>

## PRE-06 — 20 decisiones

<a id="pre-06-a1"></a>

- **PRE-06.A1** — Invitaciones por token de alta entropía y código alternativo protegido con expiración y rate limits.

  Destino canónico: [06.01](index.md#spec-06-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-a2"></a>

- **PRE-06.A2** — Salas pendientes caducan tras 30 min de inactividad (configurable), sin afectar política de partidas iniciadas.

  Destino canónico: [06.01](index.md#spec-06-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-a3"></a>

- **PRE-06.A3** — Sala con máximo dos jugadores y sin espectadores en MVP.

  Destino canónico: [06.01](index.md#spec-06-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-a4"></a>

- **PRE-06.A4** — Revancha genera nueva partida e identifica su relación con anterior; colores reasignados.

  Destino canónico: [06.01](index.md#spec-06-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-b1"></a>

- **PRE-06.B1** — Selección visual inmediata, pero ejecución/animación definitiva solo tras confirmación autoritativa del servidor.

  Destino canónico: [06.02](index.md#spec-06-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-b2"></a>

- **PRE-06.B2** — Eventos con ID y versión; detectar eventos perdidos/desordenados y volver a sincronizar por REST.

  Destino canónico: [06.02](index.md#spec-06-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-b3"></a>

- **PRE-06.B3** — Estados visibles conectado/sincronizando/desconectado/recuperado, bloqueando acciones sin confirmar.

  Destino canónico: [06.02](index.md#spec-06-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-b4"></a>

- **PRE-06.B4** — Un canal privado Ably por partida; cliente suscribe con token de capacidad restringida.

  Destino canónico: [06.02](index.md#spec-06-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-c1"></a>

- **PRE-06.C1** — Reanudar exclusivamente con identidad del participante original, de invitado o registrada, y estado confirmado desde Game.

  Destino canónico: [06.03](index.md#spec-06-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-c2"></a>

- **PRE-06.C2** — Persistir vencimiento de 24 horas y verificarlo en acciones/consultas; limpieza de datos en segundo plano no determina validez.

  Destino canónico: [06.03](index.md#spec-06-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-c3"></a>

- **PRE-06.C3** — Varias pestañas pueden observar; transacciones e idempotencia impiden duplicar movimientos.

  Destino canónico: [06.03](index.md#spec-06-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-c4"></a>

- **PRE-06.C4** — Tras reconexión, reconstruir tablero y mostrar último movimiento confirmado.

  Destino canónico: [06.03](index.md#spec-06-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-d1"></a>

- **PRE-06.D1** — Solo partidas expresamente clasificatorias entre dos registrados afectan Elo; condición fijada al comienzo.

  Destino canónico: [06.01](index.md#spec-06-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-d2"></a>

- **PRE-06.D2** — Elo básico documentado, independiente de estadísticas generales.

  Destino canónico: [07.04](index.md#spec-07-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-d3"></a>

- **PRE-06.D3** — Medidas contra abuso de enfrentamientos repetidos/manipulación de Elo, con límites auditables.

  Destino canónico: [07.04](index.md#spec-07-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-d4"></a>

- **PRE-06.D4** — Abandono voluntario es derrota; partida archivada por desconexión como inconclusa no modifica Elo.

  Destino canónico: [06.01](index.md#spec-06-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-e1"></a>

- **PRE-06.E1** — Ante fallo Ably no perder movimientos PostgreSQL; informar interrupción y reconciliar por REST.

  Destino canónico: [06.02](index.md#spec-06-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-e2"></a>

- **PRE-06.E2** — Game valida tokens firmados sin consultar Identity en cada jugada; caídas de Identity limitan altas/renovaciones.

  Destino canónico: [06.03](index.md#spec-06-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-e3"></a>

- **PRE-06.E3** — E2E con dos clientes, desconexiones, duplicados, solicitudes simultáneas, múltiples pestañas y fallos parciales.

  Destino canónico: [09.02](index.md#spec-09-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-06-e4"></a>

- **PRE-06.E4** — Medir p50/p95 de confirmación y resincronización antes de fijar SLOs definitivos.

  Destino canónico: [09.04](index.md#spec-09-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07"></a>

## PRE-07 — 24 decisiones

<a id="pre-07-a1"></a>

- **PRE-07.A1** — Registro con Google OpenID Connect / Authorization Code + PKCE, sin contraseñas propias en MVP.

  Destino canónico: [07.01](index.md#spec-07-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-a2"></a>

- **PRE-07.A2** — Identity crea invitados con UUID, alias temporal y sesión verificable y recuperable.

  Destino canónico: [07.01](index.md#spec-07-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-a3"></a>

- **PRE-07.A3** — Invitado puede vincular cuenta; partidas elegibles pasan a historial, no se convierten retroactivamente en Elo.

  Destino canónico: [07.01](index.md#spec-07-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-a4"></a>

- **PRE-07.A4** — UUID inmutable para usuario, alias público único de registrados y alias temporales no necesariamente únicos para invitados.

  Destino canónico: [07.01](index.md#spec-07-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-b1"></a>

- **PRE-07.B1** — Access tokens asimétricos cortos (JWKS público), audiencias/emisor validados en Game; Identity no consultado por movimiento.

  Destino canónico: [07.02](index.md#spec-07-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-b2"></a>

- **PRE-07.B2** — Acceso aproximadamente 15 min, sesiones renovables con inactividad configurable y duración de invitados compatible con 24 h.

  Destino canónico: [07.01](index.md#spec-07-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-b3"></a>

- **PRE-07.B3** — Refresh rotatorio, revocación al salir, registro de sesiones; tokens emitidos aún válidos hasta expiración corta salvo mecanismos extraordinarios.

  Destino canónico: [07.01](index.md#spec-07-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-b4"></a>

- **PRE-07.B4** — Cookies HttpOnly/Secure/SameSite adecuadas, mismo origen, CSRF y verificación de origen; no guardar tokens sensibles en localStorage.

  Destino canónico: [07.02](index.md#spec-07-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-c1"></a>

- **PRE-07.C1** — Roles mínimos guest/user/admin con autorización adicional por recurso.

  Destino canónico: [07.02](index.md#spec-07-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-c2"></a>

- **PRE-07.C2** — Game verifica identidad, participante, turno, estado y permiso para cada operación.

  Destino canónico: [07.02](index.md#spec-07-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-c3"></a>

- **PRE-07.C3** — Credenciales servicio-servicio de vida corta, audiencia y permisos propios, HTTPS y rotación.

  Destino canónico: [07.02](index.md#spec-07-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-c4"></a>

- **PRE-07.C4** — Game referencia UUID y obtiene perfiles públicos por API de Identity en lotes; no accede a Identity DB.

  Destino canónico: [07.03](index.md#spec-07-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-d1"></a>

- **PRE-07.D1** — Perfil público limita a alias, Elo y estadísticas; ocultar correo y datos privados.

  Destino canónico: [07.03](index.md#spec-07-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-d2"></a>

- **PRE-07.D2** — Alias editable y avatar generado, sin subidas arbitrarias de imágenes en MVP.

  Destino canónico: [07.03](index.md#spec-07-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-d3"></a>

- **PRE-07.D3** — Validación/moderación básica de alias, reservados y abuso.

  Destino canónico: [07.03](index.md#spec-07-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-d4"></a>

- **PRE-07.D4** — Estadísticas/Elo públicos; historial detallado propio y acceso controlado a reproducción compartida.

  Destino canónico: [07.03](index.md#spec-07-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-e1"></a>

- **PRE-07.E1** — Elo empieza 1200, K=32, primeras diez partidas clasificatorias provisionales; parámetros a simular/validar.

  Destino canónico: [07.04](index.md#spec-07-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-e2"></a>

- **PRE-07.E2** — Estadísticas avanzadas: partidas, resultados, tasa victoria, evolución Elo, tiempo registrado por turno (sin reloj competitivo), capturas, tuzdyks, rendimiento por color/modalidad.

  Destino canónico: [07.04](index.md#spec-07-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-e3"></a>

- **PRE-07.E3** — Ranking global paginado, búsqueda por alias, posición personal e identificación de provisionales.

  Destino canónico: [07.04](index.md#spec-07-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-e4"></a>

- **PRE-07.E4** — Game actualiza resultado, Elo y agregados en transacción atómica al finalizar partida; reintentos idempotentes.

  Destino canónico: [07.04](index.md#spec-07-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-f1"></a>

- **PRE-07.F1** — Identidades invitadas con caducidad y retención limitada, compatibles con recuperación 24 h, limpieza posterior documentada.

  Destino canónico: [07.01](index.md#spec-07-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-f2"></a>

- **PRE-07.F2** — Exportación del propio perfil, historial y estadísticas a JSON estructurado.

  Destino canónico: [07.03](index.md#spec-07-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-f3"></a>

- **PRE-07.F3** — Eliminación de cuenta distribuida e idempotente: bloquear, revocar, solicitar borrado/anonimización en Game, verificar y finalizar según privacidad aplicable.

  Destino canónico: [07.03](index.md#spec-07-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-07-f4"></a>

- **PRE-07.F4** — Pruebas de tokens, cuentas, invitados, vinculación, accesos cruzados, borrado y manipulación de rankings.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08"></a>

## PRE-08 — 24 decisiones

<a id="pre-08-a1"></a>

- **PRE-08.A1** — Turno IA se calcula automáticamente tras confirmar movimiento humano cuando toca a IA.

  Destino canónico: [08.02](index.md#spec-08-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-a2"></a>

- **PRE-08.A2** — Dos dificultades por presupuesto/profundidad/heurística, demostrablemente diferenciadas.

  Destino canónico: [08.01](index.md#spec-08-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-a3"></a>

- **PRE-08.A3** — Mostrar indicador calculando, dificultad y última jugada; no mostrar análisis profundo en partidas normales.

  Destino canónico: [08.02](index.md#spec-08-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-a4"></a>

- **PRE-08.A4** — IA sigue reglas y juega hasta final reglamentario, sin rendiciones automáticas heurísticas.

  Destino canónico: [08.01](index.md#spec-08-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-b1"></a>

- **PRE-08.B1** — Negamax alfa-beta con ordenación y profundización iterativa.

  Destino canónico: [08.01](index.md#spec-08-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-b2"></a>

- **PRE-08.B2** — Evaluación explicable: diferencia de kazanes, tuzdyk, vulnerabilidad, movilidad, fase.

  Destino canónico: [08.01](index.md#spec-08-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-b3"></a>

- **PRE-08.B3** — Tabla de transposiciones temporal y acotada, con tratamiento correcto del historial en triple repetición.

  Destino canónico: [08.01](index.md#spec-08-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-b4"></a>

- **PRE-08.B4** — Límites de tiempo, profundidad, nodos y memoria; context.Context y mejor movimiento legal disponible.

  Destino canónico: [08.01](index.md#spec-08-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-c1"></a>

- **PRE-08.C1** — Práctica controla ambos bandos, IA opcional y escenarios educativos verificados.

  Destino canónico: [08.03](index.md#spec-08-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-c2"></a>

- **PRE-08.C2** — Deshacer/rehacer por movimiento; atajo «volver a mi turno» ante IA.

  Destino canónico: [08.03](index.md#spec-08-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-c3"></a>

- **PRE-08.C3** — Catálogo inicial pequeño de ejercicios con captura, tuzdyk, defensa y final, validado.

  Destino canónico: [08.03](index.md#spec-08-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-c4"></a>

- **PRE-08.C4** — Última sesión de práctica y historial reversible en almacenamiento local versionado.

  Destino canónico: [08.03](index.md#spec-08-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-d1"></a>

- **PRE-08.D1** — Tutorial de lecciones progresivas: turnos, siembra, captura, tuzdyk, final y conceptos básicos.

  Destino canónico: [08.04](index.md#spec-08-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-d2"></a>

- **PRE-08.D2** — Sugerencias opt-in con motivos verificables; deshabilitadas durante partidas clasificatorias.

  Destino canónico: [08.04](index.md#spec-08-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-d3"></a>

- **PRE-08.D3** — Ejercicios con objetivos específicos, retroalimentación y distinción entre jugada legal y objetivo educativo.

  Destino canónico: [08.04](index.md#spec-08-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-d4"></a>

- **PRE-08.D4** — Progreso local, reiniciable/repetible y traducido español/inglés.

  Destino canónico: [08.04](index.md#spec-08-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-e1"></a>

- **PRE-08.E1** — Panel de práctica para deshacer/rehacer/reiniciar, consultar jugadas y escenarios, diferenciando partidas normales.

  Destino canónico: [08.03](index.md#spec-08-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-e2"></a>

- **PRE-08.E2** — Previsualizaciones educativas 3D temporales sin alterar estado confirmado.

  Destino canónico: [08.04](index.md#spec-08-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-e3"></a>

- **PRE-08.E3** — Historial paso a paso con anterior/siguiente/pausa/velocidad y reutilización de eventos semánticos.

  Destino canónico: [08.04](index.md#spec-08-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-e4"></a>

- **PRE-08.E4** — Alternativa HTML sin WebGL permite mismos flujos de tutorial y práctica con un motor único Go.

  Destino canónico: [08.04](index.md#spec-08-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-f1"></a>

- **PRE-08.F1** — Corpus reproducible apertura/medio/final, legalidad y condiciones terminales, cancelaciones y límites.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-f2"></a>

- **PRE-08.F2** — Torneos automatizados versus jugador aleatorio legal y referencias; medir fuerza relativa, tiempo y nodos sin Elo no verificado.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-f3"></a>

- **PRE-08.F3** — Si AI Service falla, Game verifica y puede escoger respaldo legal barato con motor compartido, sin corromper estado.

  Destino canónico: [08.02](index.md#spec-08-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-08-f4"></a>

- **PRE-08.F4** — Medir duración, profundidad, nodos, recursos, cancelaciones y fallos sin datos privados innecesarios.

  Destino canónico: [08.02](index.md#spec-08-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09"></a>

## PRE-09 — 24 decisiones

<a id="pre-09-a1"></a>

- **PRE-09.A1** — Pirámide basada en riesgos: unitarias motor/servicios, integración SQL/contratos y E2E críticos.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-a2"></a>

- **PRE-09.A2** — TDD para reglas, transacciones, autorización y contratos; aceptación y pruebas visuales para UI/3D.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-a3"></a>

- **PRE-09.A3** — Cobertura objetivo 90 % sentencias motor Go, 80 % lógica servicios; rutas críticas probadas aun con cobertura alta.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-a4"></a>

- **PRE-09.A4** — Fixtures y semillas reproducibles, casos reglamentarios y factorías aisladas para usuarios/partidas.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-b1"></a>

- **PRE-09.B1** — Corpus oficial, propiedades, fuzzing y simulación de partidas con conservación de 162 piedras.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-b2"></a>

- **PRE-09.B2** — IA: jugada legal, límites, cancelación, semillas fijas, estados terminales y diferencias medibles.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-b3"></a>

- **PRE-09.B3** — PostgreSQL real efímero en CI, Goose desde cero, consultas sqlc, índices, restricciones y transacciones.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-b4"></a>

- **PRE-09.B4** — Carreras de movimientos duplicados/concurrentes, versiones obsoletas y go test -race donde aplique.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-c1"></a>

- **PRE-09.C1** — Compatibilidad OpenAPI, modelos generados y validación de esquemas de eventos/consumidores.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-c2"></a>

- **PRE-09.C2** — Pruebas JWKS, caducidad, audiencias, rotación, privilegios de servicio y acceso cruzado.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-c3"></a>

- **PRE-09.C3** — Simulación de fallos de Identity, Game, AI, Ably y PostgreSQL; reintentos y degradación segura.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-c4"></a>

- **PRE-09.C4** — Docker Compose para integración aislada; proveedores simulados en la mayoría de tests, integración real aparte.

  Destino canónico: [09.01](index.md#spec-09-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-d1"></a>

- **PRE-09.D1** — Vitest + React Testing Library para componentes y adaptadores visuales.

  Destino canónico: [09.02](index.md#spec-09-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-d2"></a>

- **PRE-09.D2** — Playwright Chromium/Firefox/WebKit, móvil y dos participantes en contextos aislados.

  Destino canónico: [09.02](index.md#spec-09-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-d3"></a>

- **PRE-09.D3** — Screenshots 3D con entorno fijo + revisión manual por variabilidad GPU.

  Destino canónico: [09.02](index.md#spec-09-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-d4"></a>

- **PRE-09.D4** — Objetivo WCAG 2.2 AA aplicable, teclado/lector, movimiento reducido y alternativa HTML jugable.

  Destino canónico: [09.02](index.md#spec-09-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-e1"></a>

- **PRE-09.E1** — Feature branches y PR, protecciones de main, check obligatorios; no fusionar tests rotos.

  Destino canónico: [09.03](index.md#spec-09-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-e2"></a>

- **PRE-09.E2** — CI rápido por PR, completo antes de merge, pruebas intensivas escalonadas/pre-release; cambios compartidos activan consumidores.

  Destino canónico: [09.03](index.md#spec-09-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-e3"></a>

- **PRE-09.E3** — Local/CI aislados, preview Vercel, integración sintética separada y producción independiente.

  Destino canónico: [09.03](index.md#spec-09-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-e4"></a>

- **PRE-09.E4** — Releases trazables, smoke tests, rollback de código, migración expand/contract y recuperación de datos.

  Destino canónico: [09.03](index.md#spec-09-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-f1"></a>

- **PRE-09.F1** — Incidencias P0-P3, P0/P1 bloquean releases, especialmente corrupción, reglas e identidad.

  Destino canónico: [09.04](index.md#spec-09-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-f2"></a>

- **PRE-09.F2** — Medir Web Vitals, p95, IA, FPS, descarga 3D, reconexión y cuotas; umbrales con benchmarks.

  Destino canónico: [09.04](index.md#spec-09-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-f3"></a>

- **PRE-09.F3** — Escáneres de secretos/dependencias, análisis estático, govulncheck y equivalente multimódulo.

  Destino canónico: [09.03](index.md#spec-09-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-09-f4"></a>

- **PRE-09.F4** — Definition of Done: requisito + prueba observada + sin bloqueantes + docs + evidencia; confirmación de agente no basta.

  Destino canónico: [09.04](index.md#spec-09-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10"></a>

## PRE-10 — 24 decisiones

<a id="pre-10-a1"></a>

- **PRE-10.A1** — Cuatro proyectos Vercel vinculados al monorepo con builds independientes; CI verifica Go compartido y dependencias por proyecto.

  Destino canónico: [10.01](index.md#spec-10-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-a2"></a>

- **PRE-10.A2** — Adaptador pequeño de http.Handler para Gin en Go runtime de Vercel, conservando arranque local.

  Destino canónico: [10.01](index.md#spec-10-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-a3"></a>

- **PRE-10.A3** — Frontend y APIs Identity/Game bajo mismo origen vía rewrites; AI expone endpoint HTTPS protegido para llamadas servicio-servicio.

  Destino canónico: [10.01](index.md#spec-10-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-a4"></a>

- **PRE-10.A4** — Regiones Vercel y Neon próximas, elegidas con mediciones, disponibilidad y privacidad.

  Destino canónico: [10.01](index.md#spec-10-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-b1"></a>

- **PRE-10.B1** — Local, integración y producción con credenciales OAuth, Ably y bases independientes.

  Destino canónico: [10.01](index.md#spec-10-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-b2"></a>

- **PRE-10.B2** — Secretos en Vercel por entorno, rotación, no exponer secretos/logs ni versionarlos.

  Destino canónico: [10.01](index.md#spec-10-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-b3"></a>

- **PRE-10.B3** — Previews protegidas dentro de Hobby, datos sintéticos y OAuth de prueba aislado.

  Destino canónico: [10.01](index.md#spec-10-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-b4"></a>

- **PRE-10.B4** — Límites compartidos/centralizados compatibles con serverless por IP/identidad/recurso, especial invitados e IA; no solo RAM de Function.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-c1"></a>

- **PRE-10.C1** — pgxpool con límites serverless, conexión agrupada donde corresponda, directas para migraciones.

  Destino canónico: [10.03](index.md#spec-10-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-c2"></a>

- **PRE-10.C2** — PITR limitado del plan Neon Free + backups PostgreSQL externos cifrados y restauraciones de prueba; servicio externo/coste pendiente de verificación.

  Destino canónico: [10.03](index.md#spec-10-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-c3"></a>

- **PRE-10.C3** — Goose migrations aprobadas, controladas y expand/contract, por esquema de su servicio.

  Destino canónico: [10.03](index.md#spec-10-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-c4"></a>

- **PRE-10.C4** — Expiración lógica en solicitudes más limpieza diaria idempotente compatible con límites de cron Hobby.

  Destino canónico: [10.03](index.md#spec-10-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-d1"></a>

- **PRE-10.D1** — OpenTelemetry Go/Next.js; collector + Prometheus/Grafana locales; producción ajustada a planes gratuitos verificados.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-d2"></a>

- **PRE-10.D2** — Propagar traceId/requestId/gameId entre microservicios, sin datos sensibles en trazas.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-d3"></a>

- **PRE-10.D3** — Métricas p50/p95, errores, invocaciones, login fallido, jugadas inválidas, reconexión, CPU IA y cuotas.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-d4"></a>

- **PRE-10.D4** — Alertas cuando sean posibles y chequeos sintéticos proporcionados, sin prometer 24/7 real no contratado.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-e1"></a>

- **PRE-10.E1** — Coste objetivo 0 USD/mes, ningún recurso de pago sin autorización expresa; Hobby uso personal no comercial.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-e2"></a>

- **PRE-10.E2** — Umbrales internos de revisión 50/75/90 % de cuotas, mediante datos que los proveedores ofrezcan.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-e3"></a>

- **PRE-10.E3** — Límites de CPU/búsquedas IA por identidad, cache segura donde se pruebe y medición de consumo.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-e4"></a>

- **PRE-10.E4** — Escalar solo tras pruebas; Redis, colas o réplicas se evalúan como cambios sujetos a aprobación.

  Destino canónico: [10.02](index.md#spec-10-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-f1"></a>

- **PRE-10.F1** — Despliegues independientes compatibles, smoke tests, migraciones expand/contract y feature flags para riesgos.

  Destino canónico: [10.04](index.md#spec-10-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-f2"></a>

- **PRE-10.F2** — Runbooks P0-P3, degradación controlada ante fallos de reglas, auth, datos, Ably, Elo, WebGL e IA.

  Destino canónico: [10.04](index.md#spec-10-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-f3"></a>

- **PRE-10.F3** — RPO/RTO fundamentados en restauraciones ensayadas y capacidad real, no garantizar cero pérdida.

  Destino canónico: [10.03](index.md#spec-10-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-10-f4"></a>

- **PRE-10.F4** — Arquitectura Go portable a contenedores mediante adaptadores, variables y Dockerfiles documentados (pero NO crearlos en la fase documental).

  Destino canónico: [10.04](index.md#spec-10-04). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11"></a>

## PRE-11 — 24 decisiones

<a id="pre-11-a1"></a>

- **PRE-11.A1** — Specs Markdown por dominios con numeración incremental (01.01, 01.02, etc.), índices y referencias entre documentos.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-a2"></a>

- **PRE-11.A2** — Jerarquía: registro de decisiones aprobadas + ADR, specs detalladas, AGENTS.md operativo y planes por fase; conflictos explícitos.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-a3"></a>

- **PRE-11.A3** — AGENTS.md raíz conciso y específicos por servicio solo si necesarios.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-a4"></a>

- **PRE-11.A4** — ADR numerados con decisión, alternativas, motivo, consecuencias y estado.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-b1"></a>

- **PRE-11.B1** — Plantilla: ID, objetivo, requisitos, dependencias, pasos, porqués, ejemplos, contratos, pruebas y aceptación.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-b2"></a>

- **PRE-11.B2** — Detalle suficiente para implementar sin inventar arquitectura, sin describir código línea por línea.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-b3"></a>

- **PRE-11.B3** — Al menos un ejemplo relevante de Go/TS/SQL/JSON/pseudocódigo por especificación cuando aporte claridad, explicado.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-b4"></a>

- **PRE-11.B4** — Aceptación objetiva verificable y rastreable a casos de uso y fallos.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-c1"></a>

- **PRE-11.C1** — Coordinador de coherencia y especialistas delegados solo cuando haya trabajo independiente.

  Destino canónico: [11.02](index.md#spec-11-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-c2"></a>

- **PRE-11.C2** — Especializaciones: arquitectura, motor Go, microservicios, Three.js, seguridad, QA y operaciones según necesidad.

  Destino canónico: [11.02](index.md#spec-11-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-c3"></a>

- **PRE-11.C3** — Escritura paralela en ficheros independientes, única propiedad por fichero, integración/índices del coordinador.

  Destino canónico: [11.02](index.md#spec-11-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-c4"></a>

- **PRE-11.C4** — Contexto acotado por agente: instrucciones globales, alcance, decisiones y specs dependientes relevantes.

  Destino canónico: [11.02](index.md#spec-11-02). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-d1"></a>

- **PRE-11.D1** — Generación escalonada: auditoría, matriz decisiones, grafo dependencias, lotes documentales y revisión de coherencia.

  Destino canónico: [11.03](index.md#spec-11-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-d2"></a>

- **PRE-11.D2** — Priorizar según dependencias técnicas, aunque los IDs clasifiquen por dominios.

  Destino canónico: [11.03](index.md#spec-11-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-d3"></a>

- **PRE-11.D3** — Aprobación humana por lotes coherentes; informes de cambios/riesgos antes de seguir.

  Destino canónico: [11.03](index.md#spec-11-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-d4"></a>

- **PRE-11.D4** — Auditoría automática de enlaces, IDs, Markdown, secciones, referencias y revisión semántica.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-e1"></a>

- **PRE-11.E1** — Primera ejecución redacta SOLO documentación .md; no código, tests ejecutables, dependencias, infra real ni despliegue.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-e2"></a>

- **PRE-11.E2** — Plan futuro de entregas verticales y spike temprano de Go/Vercel/Neon/Ably, documentado ahora, ejecutado solo más tarde y con aprobación.

  Destino canónico: [11.03](index.md#spec-11-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-e3"></a>

- **PRE-11.E3** — Cierre de fase futuro condicionado a tests observados, aceptación, revisión, docs y evidencias.

  Destino canónico: [11.03](index.md#spec-11-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-e4"></a>

- **PRE-11.E4** — Cambios relevantes requieren propuesta/impacto/ADR y aprobación; agentes no reescriben unilateralmente decisiones.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-f1"></a>

- **PRE-11.F1** — README profesional orientado a portafolio, arquitectura, desafíos, pruebas y demo prevista, sin fingir funcionalidades implementadas.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-f2"></a>

- **PRE-11.F2** — Matriz requisito → spec → fase → componentes → pruebas → evidencia → estado.

  Destino canónico: [11.03](index.md#spec-11-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-f3"></a>

- **PRE-11.F3** — Codex en primera fase solo puede leer y generar .md autorizados; cambios en código, secretos, deps, despliegues y Git requieren autorización independiente.

  Destino canónico: [11.01](index.md#spec-11-01). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

<a id="pre-11-f4"></a>

- **PRE-11.F4** — Un prompt maestro orquesta inventario, lotes, verificaciones y gates de aprobación, con informe final de documentación creada/pending/bloqueos.

  Destino canónico: [11.03](index.md#spec-11-03). Naturaleza de la decisión: `APROBADO`. Evidencia de implementación: `PENDIENTE`.

## Actualización y revisión

Toda modificación semántica exige propuesta de cambio, impacto sobre consumidores, ADR cuando corresponda y aprobación humana. No reemplazar el texto histórico: añadir resolución aprobada y su relación de sustitución. [11.01](https://github.com/SebasChaconDev/toguz-kumalak/blob/docs/specs-lote-01/docs/specs/11-orchestration/11.01-document-standard-and-change-control.md) define el procedimiento.

Aceptación de este registro: 246 IDs únicos; conteos PRE iguales a la fuente; 246 textos iguales al anexo; un destino canónico válido por ID; sustituciones conservadas; ninguna afirmación de implementación. Las comprobaciones y pendientes del lote se registran en [PLANS](../PLANS.md#informe-lote-1).
