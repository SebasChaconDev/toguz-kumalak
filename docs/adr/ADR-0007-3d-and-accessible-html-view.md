# ADR-0007 — Tablero 3D y alternativa HTML accesible

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: ADR-0007 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner funcional: V (visual y UX).

Dominio: 04/08, experiencia visual y accesible · Lote: 1. Formato y cambios: [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md); propiedad de escritura: [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md).

Decisión base: APROBADO EN FUENTE. Validación técnica: POR_VERIFICAR. No existen evidencias de escenas, assets, accesibilidad, FPS o dispositivos probados por esta ADR.

## Contexto y problema

La experiencia visual principal debe mostrar un tablero físico de calidad premium y, simultáneamente, permitir jugar cuando WebGL falla o no está disponible. Dos presentaciones no pueden convertirse en dos reglamentos, ni la animación puede decidir el resultado de una partida. La accesibilidad exige interacción y datos útiles, además de una representación gráfica.

## Fuentes y trazabilidad

- [PRE-02.E1](../decision-register.md#pre-02-e1) elige Three.js y [PRE-02.E2](../decision-register.md#pre-02-e2) define inspiración sin copia.
- [PRE-04.A4](../decision-register.md#pre-04-a4) separa tablero Three.js de interfaz React/HTML/CSS; [PRE-04.B1](../decision-register.md#pre-04-b1) fija React Three Fiber + Drei.
- [PRE-04.C3](../decision-register.md#pre-04-c3) consume eventos del motor; [PRE-04.D4](../decision-register.md#pre-04-d4) exige teclado, texto e indicadores no solo por color.
- [PRE-04.E4](../decision-register.md#pre-04-e4) exige alternativa HTML plenamente jugable; [PRE-08.E4](../decision-register.md#pre-08-e4) incluye tutorial/práctica con motor Go único.
- [PRE-04.E1](../decision-register.md#pre-04-e1) fija objetivos de FPS sujetos a contraste; [PRE-09.D4](../decision-register.md#pre-09-d4) fija WCAG 2.2 AA como objetivo aplicable.
- [PRE-04.F2](../decision-register.md#pre-04-f2) exige procedencia/licencias y [PRE-04.F3](../decision-register.md#pre-04-f3) revisión visual incremental.

El [maestro](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) conserva decisiones y límites; la [trazabilidad](../traceability.md) registra especificaciones y evidencia futura.

## Decisión delimitada

La vista principal usa Three.js + React Three Fiber + Drei para tablero, piedras y animaciones. Menús, perfiles, HUD, estadísticas y controles accesibles usan React/HTML/CSS. Se conserva una alternativa HTML interactiva plenamente jugable ante ausencia/fallo WebGL, incluidos tutorial y práctica.

Ambas vistas consumen los mismos contratos, estado y reglas del motor Go. La vista HTML no añade un motor TypeScript ni implica funcionamiento offline. El modo local es juego en un dispositivo y sigue los contratos Go previstos.

La estética es realismo físico estilizado de alta calidad, con maderas, cavidades, piedras pulidas y detalles discretos según fuentes. RDomino inspira calidad perceptiva sin copiar archivos, diseños distintivos, modelos, sonidos ni código. Referencias culturales requieren verificación; no se inventan símbolos como si fueran auténticos.

## Alternativas y motivación

**Histórica descartada en fuente:** no consta un motor visual anterior sustituido. Three.js es la elección aprobada; no se fabrica un historial de comparación con otros motores.

**PROPUESTA, no adoptadas:** sustituir Three.js por otro motor, una experiencia exclusivamente WebGL o HTML con reglas propias. Las dos últimas contradicen decisiones aprobadas. Una alternativa visual requeriría revisión del propietario y evaluación de accesibilidad, compatibilidad y coste; no se adopta por preferencia del autor.

La motivación trazable es la combinación de calidad visual de PRE-02.E1/E2 y jugabilidad alternativa de PRE-04.E4. La separación entre representación y motor evita que una animación omitida determine una captura o el final. La accesibilidad no se da por conseguida con la mera presencia de HTML.

## Contratos e invariantes

1. La selección y la confirmación son acciones distintas. Una selección inmediata no confirma movimiento; ejecución/animación definitiva espera la autoridad del servidor conforme a [PRE-04.C1](../decision-register.md#pre-04-c1) y [PRE-06.B1](../decision-register.md#pre-06-b1).
2. Three.js anima eventos semánticos de Go. No deduce siembra/captura/tuzdyk desde posiciones de mallas. Saltar o acelerar una animación no cambia el estado confirmado.
3. Los datos textuales expresan casillas, turno y acciones disponibles; los indicadores no dependen solo de color. Teclado, foco y lector de pantalla se evaluarán en los flujos completos.
4. La alternativa HTML conserva las acciones y resultados de dominio de los modos aprobados; su diseño no necesita imitar una escena 3D para ser jugable. El detalle de controles y transición entre vistas queda POR_VERIFICAR.
5. El historial reproduce estados/movimientos versionados, no instrucciones 3D persistidas como verdad, según [ADR-0002](ADR-0002-single-versioned-go-engine.md).
6. El render previsto de 162 piedras mediante InstancedMesh y disposición determinista sigue [PRE-04.B3](../decision-register.md#pre-04-b3). Calidad automática/baja/alta, carga progresiva y efectos adaptativos no fijan aquí presupuestos técnicos nuevos.

## Consecuencias, costes y dependencias

Se mantienen dos presentaciones y una autoridad de reglas. Esto exige comprobar equivalencia funcional, accesibilidad, sincronización, reconexión y estados de error en ambas. La carga de assets, variabilidad GPU y dispositivos móviles añaden coste de optimización y revisión, sin evidencia actual de rendimiento.

Los objetivos 60 FPS escritorio y 30 FPS móvil compatible son metas de PRE-04.E1, no mediciones ni garantías. La matriz de dispositivos, escenarios y umbrales de calidad se definirá con benchmarks. El objetivo WCAG 2.2 AA requiere evaluación técnica y manual; no se afirma conformidad.

Depende de [03.03, motor/eventos](../index.md#spec-03-03), PENDIENTE_LOTE_2, y [05.02, contratos](../index.md#spec-05-02), PENDIENTE_LOTE_3. [04.01, accesibilidad](../index.md#spec-04-01), [04.02, tablero/assets](../index.md#spec-04-02), [04.03, interacción](../index.md#spec-04-03), [04.04, alternativa HTML](../index.md#spec-04-04) y [08.04, tutorial](../index.md#spec-08-04) son PENDIENTE_LOTE_4; [09.02, verificación visual/accesible](../index.md#spec-09-02) es PENDIENTE_LOTE_5. [ADR-0003](ADR-0003-rest-contracts-and-versioned-events.md) fija eventos y [ADR-0004](ADR-0004-postgres-authority-and-postcommit-ably.md) la confirmación autoritativa. No se crean assets en este lote.

## Riesgos y validación futura

Riesgos abiertos: divergencia de interacción entre vistas, foco perdido al cambiar de modo, animación atrasada respecto al estado confirmado, legibilidad móvil, movimiento no reducido y procedencia insuficiente de assets. La futura aprobación visual seguirá geometría, materiales/cámara, interacción, animaciones, móvil y acabado. Se medirán escenas con entorno fijo y revisión manual por variabilidad GPU.

## Aceptación futura

- **AC-01, positiva:** una partida de referencia termina con iguales estados, movimientos legales y resultado en 3D y HTML, usando el mismo contrato Go y evidencia de ambas rutas.
- **AC-02, positiva:** tutorial/práctica y sus controles aprobados son utilizables por teclado en HTML, con datos textuales y foco verificables; se documenta revisión de lector de pantalla y movimiento reducido.
- **AC-03, negativa:** desactivar o perder WebGL mantiene una vía HTML jugable; una selección no confirmada no se convierte en movimiento oficial durante el cambio.
- **AC-04, negativa:** saltar animaciones o recibir una captura mientras la escena está atrasada no cambia el resultado ni duplica transiciones.
- **AC-05, negativa:** un recurso sin procedencia/licencia válida o una referencia cultural no verificada no se presenta como asset autorizado/autenticidad acreditada.
- **AC-06, negativa:** FPS objetivo o HTML presente por sí solos no se reportan como rendimiento medido o conformidad WCAG demostrada.

Los criterios requieren implementación futura y no han sido ejecutados aquí.

Trazabilidad de aceptación: AC-01/AC-03 → [PRE-04.E4](../decision-register.md#pre-04-e4), dueño [04.04](../index.md#spec-04-04); AC-02 → [PRE-04.D4](../decision-register.md#pre-04-d4) y [PRE-08.E4](../decision-register.md#pre-08-e4), dueños [04.01](../index.md#spec-04-01)/[08.04](../index.md#spec-08-04); AC-04 → [PRE-04.C2](../decision-register.md#pre-04-c2), dueño [04.03](../index.md#spec-04-03); AC-05 → [PRE-04.F2](../decision-register.md#pre-04-f2), dueño [04.02](../index.md#spec-04-02); todos PENDIENTE_LOTE_4. AC-06 → [PRE-04.E1](../decision-register.md#pre-04-e1) y [PRE-09.D4](../decision-register.md#pre-09-d4), dueños [04.04](../index.md#spec-04-04), PENDIENTE_LOTE_4, y [09.02](../index.md#spec-09-02), PENDIENTE_LOTE_5. Riesgos: [RISK-13](../conflicts.md#risk-13) y [RISK-14](../conflicts.md#risk-14).

## Cambio y reversión documental

Eliminar la alternativa, cambiar motor visual o duplicar reglas requiere propuesta, impacto y aprobación expresa. Registrar controversias de accesibilidad/IP en [conflictos](../conflicts.md). La reversión documental restaura la última decisión aprobada y conserva fuentes; una futura degradación visual puede reducir efectos según calidades aprobadas, pero no autoriza eliminar modos o sustituir reglas.
