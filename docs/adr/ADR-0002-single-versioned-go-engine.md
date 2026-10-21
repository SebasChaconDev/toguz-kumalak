# ADR-0002 — Motor Go único, puro y versionado

ID: ADR-0002 · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner funcional: M (motor).

Dominio: 03/08, motor compartido · Lote: 1. Formato y cambios: [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md); propiedad de escritura: [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md).

Decisión base: APROBADO EN FUENTE. Validación técnica: POR_VERIFICAR. El documento queda sujeto a revisión; no afirma que exista un motor o un reglamento implementado y validado.

## Contexto y problema

Partidas oficiales, búsqueda IA, práctica, tutorial, reproducción y dos vistas necesitan interpretar las mismas reglas. Duplicarlas en el cliente produciría autoridades incompatibles y haría ambigua una reproducción histórica. A la vez, el motor no debe asumir persistencia, identidad ni redes: esas responsabilidades pertenecen a los servicios.

La definición de índices, recorrido, siembra, tuzdyk, atsyrau, repetición y final necesita contraste reglamentario. Esta ADR fija la arquitectura del motor, sin resolver silenciosamente posibles discrepancias normativas.

## Fuentes y trazabilidad

- [PRE-03.B2](../decision-register.md#pre-03-b2) define transformación pura y determinista sin mutación, persistencia ni red; [PRE-03.B1](../decision-register.md#pre-03-b1) fija el estado lógico y su versión.
- [PRE-03.B4](../decision-register.md#pre-03-b4) exige versiones de estado, reglamento y movimientos; [PRE-03.B5](../decision-register.md#pre-03-b5) reserva legalidad y previsualizaciones a Go.
- [PRE-03.C2](../decision-register.md#pre-03-c2) define eventos semánticos; [PRE-03.E3](../decision-register.md#pre-03-e3) basa la reproducción en movimientos/estados, sin instrucciones 3D como verdad.
- [PRE-05.C4](../decision-register.md#pre-05-c4) mantiene endpoints de simulación sin motor TypeScript; [PRE-08.E4](../decision-register.md#pre-08-e4) extiende ese motor único a la alternativa HTML.
- [PRE-03.E1](../decision-register.md#pre-03-e1), [PRE-03.E2](../decision-register.md#pre-03-e2) y [PRE-08.F3](../decision-register.md#pre-08-f3) definen corpus, invariantes y respaldo legal de AI.

Consultar también el [maestro y sus controles reglamentarios](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md), la [trazabilidad](../traceability.md) y el registro de [conflictos](../conflicts.md).

## Decisión delimitada

Se documenta un único motor de reglas Go puro y versionado, compartido por Game y AI. Recibe un estado lógico y un movimiento; produce nuevo estado, eventos semánticos y error, sin modificar la entrada. No consulta sesiones, bases, relojes de infraestructura ni proveedores.

Game aplica autenticación, autorización, versiones, idempotencia y persistencia alrededor de ese motor. AI usa el mismo reglamento para buscar, pero sus propuestas no confirman una partida. El cliente recibe legalidad, previsualizaciones, estados y eventos por contratos; no calcula una segunda versión del reglamento en TypeScript.

El estado `[2][9]`, kazanes, tuzdyks con propiedad, turno y versión siguen PRE-03.B1. La convención exacta de orientación/índices queda pendiente de la especificación del motor y del contraste normativo. No se establece una codificación nueva en esta ADR.

Práctica y tutorial comparten legalidad; sus controles de deshacer, rehacer o escenarios son flujos explícitos fuera de la autoridad competitiva. Una sesión local significa dos jugadores en un dispositivo, no funcionamiento offline. La vista HTML tampoco introduce un motor navegador.

## Alternativas y motivación

**Histórica descartada en fuente:** no consta una sustitución histórica de un motor ya elegido. El registro excluye expresamente un motor paralelo TypeScript; no se inventa una evaluación previa entre implementaciones.

**PROPUESTA, no adoptadas:** reglas duplicadas en cliente, reglas propias en AI, o portar el motor a WebAssembly para operar offline. Cualquier cambio requiere revisar autoridad, distribución, seguridad, compatibilidad y presupuesto con el propietario; no está autorizado por el modo local ni por la accesibilidad.

La motivación trazable es el determinismo de PRE-03.B2 y la reproducción estable de PRE-03.B4/E3. La ausencia de efectos laterales permite razonar sobre una transición sin confundirla con el commit de Game. Esta consecuencia no demuestra por sí sola corrección reglamentaria ni rendimiento.

## Contratos e invariantes

1. El estado original se conserva al aceptar o rechazar una transición. Los errores ilegales son tipados conforme a [PRE-03.C3](../decision-register.md#pre-03-c3).
2. Una entrada equivalente bajo las mismas versiones y contexto reglamentario necesario produce el mismo resultado lógico y eventos. El historial relevante para repetición no puede omitirse de esa definición.
3. Se conserva el total de 162 piedras en transiciones reglamentarias válidas; la ubicación entre tablero y kazanes se verifica mediante corpus y propiedades futuros.
4. Un evento de captura o siembra expresa significado de dominio. La animación puede acelerarse, omitirse o fallar sin cambiar la transición, según [ADR-0007](ADR-0007-3d-and-accessible-html-view.md).
5. Versiones de formato de estado, reglamento y registro de movimientos son conceptos distintos. Su política de compatibilidad, recuperación de partidas antiguas y retención de motores queda POR_VERIFICAR; no se presupone que un motor nuevo interprete sin cambios todos los registros anteriores.
6. El tablero completo enviado por un cliente no reemplaza el estado confirmado de Game, según [PRE-03.B3](../decision-register.md#pre-03-b3).

## Consecuencias y dependencias

Se evita repartir autoridad de reglas entre vistas y servicios. A cambio, cambios del motor requieren verificar consumidores Game/AI, contratos y reproducción de registros existentes. Las simulaciones locales/práctica dependen de endpoints y su disponibilidad; no se promete offline ni una latencia medida.

Depende de [03.01, reglamento y corpus](../index.md#spec-03-01), [03.02, estado e invariantes](../index.md#spec-03-02), [03.03, movimientos/eventos](../index.md#spec-03-03) y [03.04, resultados/reproducción](../index.md#spec-03-04), PENDIENTE_LOTE_2; y [05.02, contratos HTTP](../index.md#spec-05-02), PENDIENTE_LOTE_3. [04.03, animación](../index.md#spec-04-03) y [08.04, tutorial](../index.md#spec-08-04), PENDIENTE_LOTE_4, consumen sus eventos. [ADR-0001](ADR-0001-three-go-services-and-data-ownership.md) y [ADR-0003](ADR-0003-rest-contracts-and-versioned-events.md) fijan los límites de autoridad e intercambio.

## Riesgos y validación futura

Una divergencia de siembra, declaración automática de tuzdyk, triple repetición o atsyrau se registra como conflicto con fuente, impacto y alternativas. Se preserva PRE-03 y se bloquea el cierre del algoritmo afectado hasta resolución humana. No se usa una ilustración como reglamento.

En una sesión autorizada se ejecutarán corpus normativo, propiedades, fuzzing, inmutabilidad y reproducción entre versiones. Primero se obtiene corrección; los benchmarks y objetivos cuantitativos siguen [PRE-03.E4](../decision-register.md#pre-03-e4). La compilación del módulo compartido en Vercel también continúa POR_VERIFICAR.

## Aceptación futura

- **AC-01, positiva:** para una posición normativa y movimiento válido, Game y AI obtienen idéntica transición con las mismas versiones; se conservan 162 piedras y la entrada intacta, con evidencia reproducible.
- **AC-02, positiva:** un registro histórico con versiones identificadas se reproduce con resultado esperado mediante su política compatible documentada.
- **AC-03, negativa:** una jugada ilegal devuelve error tipado sin mutar estado ni generar un movimiento confirmado.
- **AC-04, negativa:** ante discrepancia normativa sin resolver, el algoritmo correspondiente no se declara cerrado ni se sustituye la decisión fuente.
- **AC-05, negativa:** desactivar WebGL o saltar animaciones no cambia movimientos legales ni resultado.

Ninguna de estas pruebas de producto se ha ejecutado en el lote documental.

Trazabilidad de aceptación: AC-01/AC-03 → [PRE-03.B2](../decision-register.md#pre-03-b2) y [PRE-03.C3](../decision-register.md#pre-03-c3), dueño [03.03](../index.md#spec-03-03); AC-02 → [PRE-03.E3](../decision-register.md#pre-03-e3), dueño [03.04](../index.md#spec-03-04); AC-04 → [PRE-03.E1](../decision-register.md#pre-03-e1), dueño [03.01](../index.md#spec-03-01); los tres son PENDIENTE_LOTE_2. AC-05 → [PRE-04.C2](../decision-register.md#pre-04-c2), dueño [04.03](../index.md#spec-04-03), PENDIENTE_LOTE_4. Bloqueo normativo: [RISK-01](../conflicts.md#risk-01).

## Cambio y reversión documental

Cambiar lenguaje, duplicar reglas o alterar una decisión reglamentaria exige propuesta, impacto sobre partidas/versiones/corpus, entrada en conflictos y aprobación del propietario. La reversión documental recupera la última redacción aprobada preservando fuentes y decisiones posteriores rechazadas como historial. Una reversión de código o datos necesitará su propio plan futuro; no debe reinterpretar registros con una versión distinta en silencio.
