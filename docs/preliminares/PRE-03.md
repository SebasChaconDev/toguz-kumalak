# PRE-03 — Reglas, estados y motor Go

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: PRE-03 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner: M — motor.

## Objetivo y contexto

Separar reglamento, función pura, autoridad de Game y representación visual. Las 22 decisiones aprobadas fijan responsabilidades; no constituyen todavía un algoritmo reglamentario cerrado. La base del maestro es tablero de 18 casillas, nueve por jugador, 162 piedras iniciales, objetivo 82, empate 81–81 y kazanes. Siembra, captura par rival, tuzdyk, atsyrau y finales requieren contraste normativo antes de cerrar su especificación.

## Fuentes y decisiones

Consultar el [maestro íntegro](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) y el [registro canónico](../decision-register.md). Síntesis de naturaleza `APROBADO`:

- [PRE-03.A1](../decision-register.md#pre-03-a1): se detectará triple repetición reglamentaria para reclamación; la automaticidad depende de la regla formal.
- [PRE-03.A2](../decision-register.md#pre-03-a2): se admiten propuestas de tablas acordadas, ampliando PRE-02.C4.
- [PRE-03.A3](../decision-register.md#pre-03-a3): crear tuzdyk será automático cuando sea reglamentario, con explicación visual contextual.
- [PRE-03.A4](../decision-register.md#pre-03-a4): empiezan blancas; online asigna colores al azar y local/IA permite elegir bando.
- [PRE-03.B1](../decision-register.md#pre-03-b1): el estado incluirá [2][9], kazanes, tuzdyks, turno y versión con índices/recorrido explícitos.
- [PRE-03.B2](../decision-register.md#pre-03-b2): estado y movimiento producen estado nuevo, eventos o error sin mutación, red ni persistencia.
- [PRE-03.B3](../decision-register.md#pre-03-b3): los comandos aportan origen, versión e idempotencia y se autorizan por sesión, nunca por tablero cliente.
- [PRE-03.B4](../decision-register.md#pre-03-b4): formatos de estado, reglas y movimientos tendrán versiones para reproducción estable.
- [PRE-03.B5](../decision-register.md#pre-03-b5): legales y previsualizaciones proceden de Go, evitando un reglamento paralelo en frontend.
- [PRE-03.C1](../decision-register.md#pre-03-c1): tuzdyks tienen posición/propietario explícitos y transfieren inmediatamente las piedras recibidas al kazán.
- [PRE-03.C2](../decision-register.md#pre-03-c2): eventos semánticos describen siembra, captura, tuzdyk, turno y fin para la animación.
- [PRE-03.C3](../decision-register.md#pre-03-c3): una jugada ilegal genera error tipado sin modificar el estado de entrada.
- [PRE-03.C4](../decision-register.md#pre-03-c4): resultado y causa de terminación se modelan por separado.
- [PRE-03.D1](../decision-register.md#pre-03-d1): práctica admite escenarios/control libre sin efectos competitivos.
- [PRE-03.D2](../decision-register.md#pre-03-d2): deshacer/rehacer usa estados y una nueva jugada descarta la rama futura.
- [PRE-03.D3](../decision-register.md#pre-03-d3): la IA básica empleará Negamax/Minimax con alfa-beta y heurística explícita.
- [PRE-03.D4](../decision-register.md#pre-03-d4): AI ejecutará Go en Vercel con cancelación, cómputo limitado y respuesta legal.
- [PRE-03.D5](../decision-register.md#pre-03-d5): fácil y normal deberán diferenciarse y evaluarse.
- [PRE-03.E1](../decision-register.md#pre-03-e1): posiciones del reglamento y referencias independientes formarán el corpus.
- [PRE-03.E2](../decision-register.md#pre-03-e2): unitarias, propiedades, fuzzing y simulaciones comprobarán especialmente conservación de 162 piedras.
- [PRE-03.E3](../decision-register.md#pre-03-e3): reproducción usa movimientos/estados versionados recalculados; instrucciones 3D no son autoridad.
- [PRE-03.E4](../decision-register.md#pre-03-e4): los objetivos de rendimiento se fijarán después de corrección y benchmarks reproducibles.

## Agrupaciones de diseño y consecuencias

**Reglamento y representación.** El convenio de índices debe relacionar vista de ambos jugadores, casillas numeradas y recorrido circular. La equivalencia usada por repetición debe considerar los componentes reglamentarios pertinentes y su historial; no basta comparar una captura de pantalla o un hash sin semántica. Las restricciones de tuzdyk y las condiciones completas de final no se infieren del aspecto del tablero.

**Motor y autoridad.** La función pura compartida resuelve legalidad; Game autentica, autoriza, aplica control de versión y confirma transaccionalmente. AI propone movimientos y el cliente presenta eventos. Ninguno sustituye la autoridad de Game sobre resultados oficiales. La inmutabilidad permite probar entrada intacta ante errores y reutilizar estados en práctica/reproducción.

**Estado, eventos y causas.** Las animaciones pueden saltarse sin alterar una partida. La persistencia debe distinguir victoria/empate/inconclusa de su causa, para no computar Elo por desconexión archivada. El versionado necesita una política futura de compatibilidad de reproducciones, no una promesa de recalcular cualquier registro con cualquier versión.

## Sustituciones y matices

Las tablas acordadas y la reclamación por repetición amplían [PRE-02.C4](../decision-register.md#pre-02-c4). Tuzdyk automático conserva su decisión fuente: una diferencia con normativa formal se registra y bloquea el algoritmo afectado, no se cambia unilateralmente. El motor único será Go también para local, práctica y vista HTML; no se aprueba motor TypeScript ni funcionamiento offline.

## Dependencias y fuentes canónicas

El [dominio 03 del índice](../index.md#spec-03-01), `PENDIENTE_LOTE_2`, debe resolver primero fuente normativa, convención y corpus. [PRE-05](PRE-05.md) delimita transacción, [PRE-04](PRE-04.md) consume eventos, [PRE-08](PRE-08.md) detalla IA/práctica y [PRE-09](PRE-09.md) define evidencia. Seguir [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [trazabilidad](../traceability.md).

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: fuente normativa oficial, recorrido de siembra, atsyrau, restricciones/declaración de tuzdyk, equivalencia de repetición y prioridad entre condiciones terminales. El [registro de conflictos](../conflicts.md) gobierna el bloqueo de cierre y la aprobación requerida. Tampoco están medidas fuerza IA, coste o duración. No se escriben algoritmos ejecutables ni corpus de producto en disco en este lote.

## Criterios futuros de aceptación y verificación

- **AC-01:** cada regla tiene fuente, interpretación aprobada y posiciones entrada → salida, incluidos casos negativos y terminales; ninguna discrepancia normativa permanece ocultada.
- **AC-02:** contrato puro demuestra determinismo, ausencia de mutación, errores tipados y conservación del total entre tablero/kazanes; eventos explican el cambio sin gobernarlo.
- **AC-03:** comandos duplicados, ilegales o con versión antigua no generan un segundo movimiento ni efecto oficial.
- **AC-04:** reproducción y práctica usan versiones e historial definido; bifurcar tras deshacer invalida el futuro anterior y no modifica estadísticas competitivas.

Estos criterios describen evidencia futura. La revisión documental no valida la corrección del juego.

## Revisión humana

El propietario debe resolver cualquier discrepancia normativa antes del cierre del algoritmo afectado. Mantener `EN_REVISIÓN` y las decisiones fuente intactas. La aprobación del lote será «APRUEBO LOTE 1»; no autoriza implementación.
