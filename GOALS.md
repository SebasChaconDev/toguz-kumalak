# GOALS — Objetivos y alcance comprometido

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: GOALS · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner: C.

## Objetivo de la misión actual

Transformar las 246 decisiones de PRE-01 a PRE-11 en especificaciones Markdown trazables y revisables por lotes. El propietario aprobó el inventario de 81 archivos y el lote 1 de 31 Markdown, ya integrado. Autorizó diez specs del lote 2 y su PR sin fusionar; las nuevas specs quedan EN_REVISIÓN y el producto no se implementa.

Fuentes: [maestro íntegro](docs/sources/TOGUZ_PROMPT_MAESTRO_CODEX.md), [REG](docs/decision-register.md), [inventario aprobado](docs/index.md#inventario), [PLANS](PLANS.md).

## Resultado de producto aprobado

| Objetivo | Resultado verificable futuro | Fuente principal / spec canónica |
|---|---|---|
| G-01 Reglas tradicionales |18 pozos,9 por jugador,9 piedras iniciales por pozo,162 en total; victoria 82, empate 81–81; siembra/captura/tuzdyk/final según norma contrastada | [PRE-03](docs/preliminares/PRE-03.md), [03.01–03.04](docs/index.md#spec-03-01) |
| G-02 Modos de juego | Local dos personas, IA fácil/normal, práctica separada, privado online; selección de bando local/IA y colores aleatorios online | [PRE-02](docs/preliminares/PRE-02.md), [02.01](docs/index.md#spec-02-01) |
| G-03 Recuperación y continuidad | Identidad invitada verificable, última sesión local reanudable y online recuperable 24 h; inconclusa al expirar | [PRE-06](docs/preliminares/PRE-06.md), [06.03](docs/index.md#spec-06-03) |
| G-04 Aprendizaje accesible | Tutorial opcional ES/EN, reglamento completo, escenarios verificados, undo/redo solo práctica, hints fuera de ranked | [PRE-08](docs/preliminares/PRE-08.md), [08.03–08.04](docs/index.md#spec-08-03) |
| G-05 Experiencia premium | Three.js/R3F/Drei, tablero propio, cámara/materiales/animaciones y HTML plenamente jugable sin WebGL | [PRE-04](docs/preliminares/PRE-04.md), [04.01–04.04](docs/index.md#spec-04-01) |
| G-06 Identidad y datos | Google OIDC/PKCE, invitados vinculables, perfiles, historial/replay con acceso controlado, exportación y borrado distribuido | [PRE-07](docs/preliminares/PRE-07.md), [07.01–07.03](docs/index.md#spec-07-01) |
| G-07 Competición íntegra | Solo ranked entre registrados altera Elo;1200/K32/diez provisionales por validar; clasificación y estadísticas avanzadas | [PRE-07.E1–E4](docs/decision-register.md#pre-07-e1), [07.04](docs/index.md#spec-07-04) |
| G-08 Arquitectura de portafolio | Tres microservicios Go/Gin independientes, motor compartido y contratos tipados; propiedad de datos separada | [PRE-05](docs/preliminares/PRE-05.md), [05.01](docs/index.md#spec-05-01) |
| G-09 Calidad observable | Pruebas críticas y negativas, CI y evidencia de aceptación; ninguna afirmación de agente sustituye resultados | [PRE-09](docs/preliminares/PRE-09.md), [09.04](docs/index.md#spec-09-04) |
| G-10 Operación viable | Cuatro proyectos Vercel, Neon/Ably, objetivo 0 USD, backups y recuperación ensayada, portabilidad documentada | [PRE-10](docs/preliminares/PRE-10.md), [10.01–10.04](docs/index.md#spec-10-01) |

## MVP base y ampliado

Ambos forman el alcance comprometido de lanzamiento. Se propone una progresión vertical desde motor y cliente mínimo hasta online, aprendizaje, cuentas y competición; la distribución detallada se cerrará en 02.01 y 11.03 sin retirar ninguna capacidad. La IA básica aparece tanto en modos iniciales como en alcance ampliado: eso permite introducción incremental, no posponerla fuera del lanzamiento.

Permanecen incluidos perfiles, historial/reproducción, Elo, clasificación, estadísticas avanzadas (resultados, tasa de victoria, evoluciónElo, tiempo por turno sin reloj competitivo, capturas, tuzdyks y rendimiento por color/modalidad), tutorial y práctica reversible.

## Límites aprobados

- Sin reloj competitivo MVP. Medir tiempo por turno para estadísticas no crea un límite de juego.
- Salas privadas de dos participantes; sin matchmaking abierto obligatorio ni espectadores.
- Sin undo en partidas ordinarias; práctica reversible no afecta Elo.
- Sin Elo en IA, invitados, práctica o inconclusas. Vincular cuenta no reclasifica partidas pasadas.
- Torneos, chat, espectadores, IA avanzada, pagos y funciones sociales extensas quedan pospuestos según PRE-02.E3.
- Juego local no es requisito offline. Go calcula reglas también para HTML y simulación local/práctica.
- Sin cuarto microservicio ni Redis/colas/réplicas preventivas. Escalado y pagos requieren aprobación.
- Activos originales/licenciados y referencias culturales contrastadas; no copiar RDomino.
- No elegir unilateralmente licencia, región, retenciones o parámetros técnicos aún pendientes.

## Criterios de éxito documental

| ID | Entrada → resultado → evidencia exigida |
|---|---|
| AC-G-01 | Anexo completo →246 decisiones conservadas sin huérfanas → comparación de IDs/textos y matriz |
| AC-G-02 | Inventario autorizado →solo rutas del lote en curso → manifiesto de archivos y estado Git |
| AC-G-03 | Nueva spec →requisitos, diseño, pasos/porqués, fallos y AC suficientes → revisión 17 apartados y enlaces canónicos |
| AC-G-04 | Dato no probado →POR_VERIFICAR con validación futura → registro de conflictos y ausencia de claims falsos |
| AC-G-05 | Cambio de alcance →propuesta e impacto antes de modificarlo → aprobación humana y ADR cuando corresponda |
| AC-G-06 | Fin de lote →reporte, estadoEN_REVISIÓN y parada → autorización explícita del propietario para avanzar |

## Riesgos y éxito futuro del producto

Los 14 casos de [CONF](docs/conflicts.md) impiden convertir aspiraciones en garantías. La aceptación de implementación requerirá reglas contrastadas, resultados de pruebas observadas, ausencia de bloqueantes, revisión humana y evidencia conforme a la futura DoD. No hay fechas, FPS, SLO, fuerzaIA, RPO/RTO o coste real garantizados en esta entrega.
