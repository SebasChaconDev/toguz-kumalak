# PRE-04 — Experiencia visual, UI y Three.js

ID: PRE-04 · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner: V — UX.

## Objetivo y contexto

Consolidar la dirección visual y los flujos de interacción de las 25 decisiones aprobadas. El tablero físico premium debe resultar legible, accesible y compatible con la autoridad del motor Go. No se producen diseños, modelos, sonidos ni assets en este lote.

## Fuentes y decisiones

La fuente íntegra es el [maestro archivado](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md); las formulaciones exactas se conservan en el [registro](../decision-register.md). Implicaciones de naturaleza `APROBADO`:

- [PRE-04.A1](../decision-register.md#pre-04-a1): el acabado será realismo físico estilizado premium, sin exigir fotorrealismo extremo.
- [PRE-04.A2](../decision-register.md#pre-04-a2): maderas cálidas/oscuras, cavidades, piedras pulidas y metal discreto orientan materiales.
- [PRE-04.A3](../decision-register.md#pre-04-a3): referencias culturales kazajas requerirán verificación y uso respetuoso.
- [PRE-04.A4](../decision-register.md#pre-04-a4): Three.js dibuja tablero y React/HTML/CSS contiene HUD, menús y controles accesibles.
- [PRE-04.B1](../decision-register.md#pre-04-b1): el stack visual será Three.js, React Three Fiber y Drei.
- [PRE-04.B2](../decision-register.md#pre-04-b2): geometría procedural y Blender/GLB/glTF se complementarán según detalle necesario.
- [PRE-04.B3](../decision-register.md#pre-04-b3): las 162 piedras usarán InstancedMesh y colocación determinista ligada al estado.
- [PRE-04.B4](../decision-register.md#pre-04-b4): la cámara oblicua tendrá zoom, rotación controlada y cambios de orientación.
- [PRE-04.B5](../decision-register.md#pre-04-b5): PBR, ambiente, sombras suaves y efectos moderados se adaptarán a capacidad.
- [PRE-04.C1](../decision-register.md#pre-04-c1): seleccionar y confirmar serán acciones separadas para prevenir accidentes.
- [PRE-04.C2](../decision-register.md#pre-04-c2): la siembra se animará por secuencia acelerable/saltable sin efecto lógico.
- [PRE-04.C3](../decision-register.md#pre-04-c3): la animación consume eventos Go y no calcula reglas.
- [PRE-04.C4](../decision-register.md#pre-04-c4): el audio será opcional, discreto y configurable, sin reproducción invasiva.
- [PRE-04.D1](../decision-register.md#pre-04-d1): la navegación cubre portada, modos, sala, partida, resultado, aprendizaje, perfil, historial y ranking.
- [PRE-04.D2](../decision-register.md#pre-04-d2): móvil vertical/horizontal adaptará cámara y controles y podrá sugerir horizontal por legibilidad.
- [PRE-04.D3](../decision-register.md#pre-04-d3): local permitirá giro entre turnos opcional y cámara fija.
- [PRE-04.D4](../decision-register.md#pre-04-d4): teclado, descripción textual, señales adicionales al color y HTML interactivo son requisitos de acceso.
- [PRE-04.D5](../decision-register.md#pre-04-d5): tutorial opcional usará resaltado contextual y escenarios en el tablero 3D.
- [PRE-04.E1](../decision-register.md#pre-04-e1): 60 FPS escritorio y 30 móvil son objetivos sujetos a dispositivos/medición, no garantías.
- [PRE-04.E2](../decision-register.md#pre-04-e2): calidades automática, baja y alta regularán resolución, sombras y efectos.
- [PRE-04.E3](../decision-register.md#pre-04-e3): carga diferida/progresiva, reutilización y compresión reducirán recursos 3D.
- [PRE-04.E4](../decision-register.md#pre-04-e4): ausencia o fallo WebGL activará una vista HTML plenamente jugable.
- [PRE-04.F1](../decision-register.md#pre-04-f1): se requieren referencias visuales, automatización, revisión multiplataforma y medición.
- [PRE-04.F2](../decision-register.md#pre-04-f2): la procedencia/licencia de recursos debe registrarse; no copiar RDomino.
- [PRE-04.F3](../decision-register.md#pre-04-f3): la aprobación visual avanzará de geometría a acabado por etapas explícitas.

## Agrupaciones de diseño y consecuencias

**Dirección artística y procedencia.** La materialidad premium se valida por legibilidad y coherencia además de apariencia. RDomino es referencia de calidad, no biblioteca de recursos ni autorización para copiar diseños distintivos. La revisión cultural y el inventario de licencias preceden a aceptar recursos definitivos.

**Interacción y animación.** Selección, previsualización, confirmación pendiente y movimiento confirmado necesitan estados distinguibles. La animación solo empieza de forma definitiva con confirmación autoritativa en online. Saltar, acelerar o interrumpir una secuencia debe reconstruir la misma posición confirmada. Disposición determinista de piedras facilita reconciliación y referencias visuales.

**Acceso y rendimiento.** La interfaz HTML/CSS conserva lectura, controles y datos de casillas; la vista alternativa permite jugar, aprender y practicar con el mismo motor Go. No es una captura estática ni un segundo reglamento. Calidad adaptativa y carga progresiva deben preservar interacción ante GPU limitada; los objetivos de FPS requieren escenas, dispositivos y método de medición definidos.

## Sustituciones y matices

[PRE-02.E2](../decision-register.md#pre-02-e2) queda precisada por el realismo estilizado aprobado en A1. La sugerencia horizontal no impide juego vertical. Audio y giro son configurables, no condiciones para jugar. El rendimiento objetivo no prueba cumplimiento; la alternativa HTML tampoco significa offline. No hay permiso para copiar archivos, código, sonidos o modelos de la referencia visual.

## Dependencias y fuentes canónicas

El [dominio 04 del índice](../index.md#spec-04-01) es `PENDIENTE_LOTE_4`. Consume estados/eventos de [PRE-03](PRE-03.md), confirmación de [PRE-06](PRE-06.md), flujos de [PRE-02](PRE-02.md) y aprendizaje de [PRE-08](PRE-08.md). [PRE-09](PRE-09.md) fija verificación visual/accesible. Aplican [11.01](https://github.com/SebasChaconDev/toguz-kumalak/blob/docs/specs-lote-01/docs/specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](https://github.com/SebasChaconDev/toguz-kumalak/blob/docs/specs-lote-01/docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [trazabilidad](../traceability.md).

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: referencias culturales, licencias, variabilidad GPU, dispositivos de referencia, presupuesto de descarga y mediciones FPS. Véase [conflictos](../conflicts.md). No afirmar pruebas visuales, conformidad WCAG o rendimiento observado. La producción de geometrías, GLB, texturas, capturas y código queda fuera del lote.

## Criterios futuros de aceptación y verificación

- **AC-01:** el recorrido visual documenta todas las pantallas, estados de confirmación y fallos con adaptación móvil y control por teclado.
- **AC-02:** el mismo caso confirmado/reconectado produce posición equivalente al saltar animaciones y al cambiar a HTML.
- **AC-03:** una matriz de accesibilidad exige texto de casillas, foco, señales no solo cromáticas y movimiento reducido; la conformidad requiere pruebas observadas futuras.
- **AC-04:** escenas/dispositivos de referencia fijan método de medición y comparaciones; cada etapa visual y recurso tiene revisión y procedencia registradas.

Los controles actuales son documentales, sin ejecución de producto ni medición GPU.

## Revisión humana

Revisar dirección artística, límites de copia, equivalencia HTML y secuencia de aprobación visual. El documento sigue `EN_REVISIÓN`; «APRUEBO LOTE 1» autoriza únicamente el cierre documental de este lote y exige detenerse antes del siguiente.
