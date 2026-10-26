# PH-03 — Cliente HTML jugable y experiencia 3D progresiva

> **Aprobación posterior — 2026-10-26: APROBADO documentalmente.** El propietario declaró «APRUEBO LOTE 5 como documentación técnica.» sobre HEAD 9ff1a2857d0b91b9cf7925a4c36de30e377cef3d; PR #5 integrado por squash 6579ca930a9081f773b390ce61f6bc80cc6b1ae7. El [recibo de aprobación](../../PLANS.md#cierre-lote-5) prevalece sobre los estados de revisión históricos. L4-P01, la política de retención y demás propuestas mantienen su aprobación específica pendiente; sin permiso de implementación, workflows, proveedores o lote6.

ID: PH-03 · Versión: 0.2 · Fecha: 2026-10-26 · Estado documental: APROBADO · Owner: C — Coordinación · Lote: 5.

Plan futuro PROPUESTA; decisiones APROBADO, implementación/accesibilidad/rendimiento POR_VERIFICAR. La primera partida HTML no espera todos los materiales premium; el acabado comprometido sigue dentro del lanzamiento.

## Ficha Milestone/1

| Campo requerido | Contenido del hito |
|---|---|
| Identificador | PH-03 |
| Objetivo | Permitir una partida local completa por HTML/Go y añadir Three.js sobre el mismo modelo, con foco/intención y acceso equivalentes. |
| Dependencias | PH-02; 02.01–02.03; 03.02–03.04; 04.01–04.04; 05.02/05.04; 06.01–06.03 y 07.01–07.04 como contratos consumidos según pantalla; 09.02/09.03/09.04; 10.02; 11.03. |
| Responsable | V implementa frontend/escena/acceso; S/M apoyan integración Go; Q verifica, C coordina y propietario acepta revisiones visuales/hito. |
| Entrada | Simulación/lecturas/sesión PH02 disponibles y aceptadas; fuentes/gate 11.03, permiso de código, contratos estables y recursos/fixtures autorizados. |
| Entregables | Shell ES/EN, partida local HTML, guardado/restauración, controlador de intención, tablero/recursos/cámara/animación/audio y fallback. |
| Pruebas positivas/negativas | PH03-C01–C06, F-Q2-01/02/04/07/08/10/11/12 y mediciones 09.04 aplicables. |
| Riesgos | Motor duplicado, foco perdido, preview vieja, unknown duplicado, reglas deducidas de meshes, GPU y procedencia/licencias no acreditadas. |
| Evidencia obligatoria | QA-Evidence/1, DOM/foco/red/Go, revisión manual teclado/lector, visuales reproducibles y medición por dispositivo/versiones. |
| Salida | Partida local completa HTML/3D, invariantes/controles equivalentes, fallback seguro; sin P0/P1 o críticos pendientes y DoD 09.04. |
| Aprobación humana | Dossier y revisiones incrementales de geometría, cámara/materiales, interacción, animación, móvil/acabado; aceptación no prueba FPS/WCAG por sí sola. |
| Corrección/recuperación | Conservar estado/foco/intención al degradar; abandonar lote incompleto y mostrar snapshot, corregir/regresar/revisar sin regenerar baseline automáticamente. |

## Fuentes, contratos y límites de incremento

Leer [02.01](../specs/02-product/02.01-scope-and-modes.md)–[02.03](../specs/02-product/02.03-accounts-history-and-licensing.md), [03.02](../specs/03-game-engine/03.02-state-indices-and-invariants.md)–[03.04](../specs/03-game-engine/03.04-outcomes-repetition-and-replay.md) y [04.01](../specs/04-visual-experience/04.01-navigation-and-accessibility.md), [04.02](../specs/04-visual-experience/04.02-board-assets-camera-and-materials.md), [04.03](../specs/04-visual-experience/04.03-interaction-animation-and-audio.md), [04.04](../specs/04-visual-experience/04.04-performance-and-html-fallback.md). PRE-04, PRE-02.B3/D3 y PRE-09.D1–D4 del [REG](../decision-register.md) y [ADR-0007](../adr/ADR-0007-3d-and-accessible-html-view.md) gobiernan equivalencia/procedencia.

Consumir [05.02](../specs/05-backend/05.02-http-openapi-and-errors.md)/[05.04](../specs/05-backend/05.04-event-schemas-and-postcommit-publication.md), [06.01](../specs/06-multiplayer/06.01-rooms-invitations-and-ranked-lifecycle.md)–[06.03](../specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md) y [07.01](../specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md)–[07.04](../specs/07-identity-data/07.04-elo-leaderboard-and-statistics.md) al preparar pantallas/estados; su lectura no afirma implementadas salas/cuentas/ranking. [09.02](../specs/09-quality/09.02-ui-e2e-accessibility-and-visual-verification.md), [09.03](../specs/09-quality/09.03-ci-security-and-release-controls.md), [09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md), [10.02](../specs/10-operations/10.02-observability-quotas-and-rate-limits.md) y [11.03](../specs/11-orchestration/11.03-milestones-gates-and-evidence.md) fijan QA/observación/gates. Consultar [PLANS](../../PLANS.md#ejecucion-lote-5)/[CONF](../conflicts.md).

Contratos: PitRef/State/LegalMovesView/PreviewView, SimulationRequest/Reply, LocalSaveEnvelope/PracticeSaveEnvelope, SemanticEvent y réplicas autorizadas. Local es dos personas en un dispositivo mediante Go; no offline ni dos UUID obligatorios. Guardado local no crea historia/Elo Game. Una preview es hipotética; selección≠confirmación. Para futuras partidas oficiales usar controlador pending/unknown/receipt y scheduler 06.02, sin polling por componente; el flujo multijugador real se termina en PH04.

## Componentes futuros y tickets

Rutas candidatas: apps/web, packages/contracts y contratos/puertos ya establecidos PH02. Separar shell React/HTML/CSS, controlador de estado/intención, adaptador Go/almacenamiento, tablero textual, escena R3F/Drei y proyección de eventos. Fijar versiones/peer dependencies exactas antes de código; familias sugeridas en 04.02 no son lockfile ni compatibilidad acreditada.

| Ticket / depende de | Pasos y razón | Resultado pequeño |
|---|---|---|
| PH-03-T01 / entrada | Diseñar shell/landmarks/rutas ES/EN, estados carga/vacío/error/denegación y vocabulario revisado. | Navegación accesible sin canvas; C01. |
| PH-03-T02 / T01+PH02 | Tablero HTML18pozos/2kazanes/turno/tuz/resultado y selección→previewGo→confirmación explícita; teclado/tacto. | Primera partida local legal completa por ambos colores; C01/C02. |
| PH-03-T03 / T02 | Guardar última local explícitamente, validar/restaurar con Go/versiones/contexto y política de dos pestañas; informar quota/red/corrupción. | Reanudación honesta sin credenciales ni reparación silenciosa; C02. |
| PH-03-T04 / T02 | Controlador compartido de foco/intención y adaptador de eventos/revisión; fixture de pending/unknown oficial declarado para no duplicar comandos al cambiar renderer. | Semántica común y un scheduler; C03/C04. |
| PH-03-T05 / T04 | Geometría procedural inicial, slots162/InstancedMesh/PitRef estable y cámara/reset/presets; revisión humana de legibilidad antes de detalle. | Tablero3D funcional con HTML disponible; C04/C05. |
| PH-03-T06 / T05 | Materiales/luz/calidades progresivas, assets originales/autorizados y ficha de procedencia; revisar neutral/cámara/móvil. | Acabado incremental aprobado, sin copia de RDomino ni autenticidad cultural inventada; C05. |
| PH-03-T07 / T04+T05 | Animar solo eventos completos Go, skip/reducedMotion/audio opt-in; fallback tras WebGL ausente/context loss sin cambiar intención/cursor. | Equivalencia inmediata al snapshot; C03/C04. |
| PH-03-T08 / T03+T06+T07 | Vitest/RTL/Playwright tres motores, teclado/lector manual, responsive ES/EN y mediciones 04.04/09.04; revisar artefactos. | Dossier con alcance local y pendientes de modos futuros; C06. |

La UI puede usar dobles etiquetados de contratos de salas/cuentas para desarrollar componentes; no anuncia esas funciones operativas hasta sus fases. T06 no bloquea la primera partida T02; tampoco elimina premium del lanzamiento.

## Pruebas, aceptación y evidencias

| Caso / fuente QA | Positivo esperado | Negativo esperado | Evidencia adicional |
|---|---|---|---|
| PH03-C01 / F-Q2-01/02/10/11 | Navegación/partida local completas en ES/EN por teclado/tacto con un principal. | Sin mouse/hover/audio/color/GPU no se pierde acción; input inválido no se calcula enTS. | DOM, roles/nombres/foco, llamadasGo y tarea manual registrada. |
| PH03-C02 / F-Q2-02/09 | Guardar/restaurar contexto/versiones y continuar legalmente. | Offline, storage ausente/corrupto/lleno, versión desconocida o dos tabs conservan último válido y explican límite. | Envelope/hash no secreto, Go validate, timeline y errores. |
| PH03-C03 / F-Q2-07 | Cambiar a HTML desde selected/pending/unknown/animating preserva foco y operación original. | Renderer nuevo no emite Move/Start/Resume ni adopta preview o resultado incierto. | Red/IDs antes-después, fixture o PG real según etapa; tipo de ensayo declarado. |
| PH03-C04 / F-Q2-04/08 | Eventos completos ordenados, skip/reducedMotion convergen al mismo snapshot162. | Duplicado/desorden/gap/version desconocida abandona animación parcial; no barrido atsyrau ni empate automático. | Eventos/revisiones/IDs, estados exactos y comparación HTML/3D. |
| PH03-C05 / F-Q2-12 | Geometría/cámara/material y slots se comparan en escena fija, procedencia válida. | Baseline no se autoacepta; dato falso/cantidad/tuz/resultado no se oculta con máscara; asset sin permiso no pasa. | Baseline/actual/diff/hash, viewport/GPU/fuentes y revisión humana/ficha. |
| PH03-C06 / F-Q2-10/11, F-Q4-05 | Chromium/Firefox/WebKit, portrait/landscape, reducedMotion, zoom/reflow y lector completan tareas; recursos se liberan. | axe/emulación/idle no acreditan conformidad física/manual o60/30FPS; foco tapado o acción ausente bloquea. | AccessibilityReview/1, recorrido manual, medición frame times/colas/bytes/memoria por dispositivo. |

QA-Evidence/1 liga PRE/spec/AC/commit/build/versiones/corpus/entorno/comando/observado/artefacto/revisor. Capturas fijan escena/cámara/viewport/fuentes y versión, tolerancia calibrada por entorno; revisión manual GPU acompaña automatización. Medir frío/caliente, HTML usable/3D listo, Web Vitals, frames y memoria con 09.04. Metas60/30 son fuente; presupuestos 04.04 y criterios de medición nuevos son PROPUESTA, nunca un resultado prometido.

## Seguridad, recursos y riesgos

No tokens/PII en localStorage, URL, screenshots o logs públicos. Recursos/decoders/licencias/cultura deben revisarse antes de publicación; compras/proveedores requieren permiso separado. Audio solo opt-in sin autoplay. Cargar3D bajo demanda preserva HTML y límites de descarga; liberar recursos al salir sin duplicar conexiones/audio/polling. RISK-13/14 y L4-P05/06/08 siguen abiertos hasta evidencia aplicable.

## DoD, gate y corrección

DoD 09.04: local jugable/reanudable sobre Go, equivalencia HTML/3D y fallback observados, críticos/negativos completos, revisión visual/manual independiente, docs/compatibilidad/evidencias accesibles. No cerrar con P0/P1, crítico no validado o requerido NOT_RUN/BLOCKED/FAIL; P2 solo con aceptación humana acotada/owner/plazo, sin eximir acceso esencial.

Q contrasta dossier, C prepara GateEvaluation/1 y propietario acepta salida/revisiones visuales del commit exacto. Corregir bajo fuente y repetir controles afectados; preservar fallo y baseline previa. A coordina incidente operativo, C gobierno y Q recuperación; fallback puede contener daño sin declarar resuelto el defecto. [PH-04](phase-04-private-multiplayer.md) necesita permiso propio y pruebas reales de dos participantes; este hito no autoriza release/recursos ni anuncia el producto completo.
