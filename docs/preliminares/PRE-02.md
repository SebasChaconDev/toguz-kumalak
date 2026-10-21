# PRE-02 — Producto, funcionalidades y alcance

ID: PRE-02 · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner: C — producto.

## Objetivo y contexto

Delimitar los modos de juego y el lanzamiento comprometido sin convertir hitos internos en recortes del producto. Las 18 decisiones aprobadas describen lo que deberá ofrecerse; este documento no afirma que ninguna pantalla o funcionalidad esté implementada.

## Fuentes y decisiones

El [maestro archivado](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md), Anexo A, y el [registro canónico](../decision-register.md) son las fuentes. La siguiente síntesis tiene naturaleza `APROBADO`:

- [PRE-02.A1](../decision-register.md#pre-02-a1): habrá juego de dos personas en un dispositivo, rival IA básica y práctica separada.
- [PRE-02.A2](../decision-register.md#pre-02-a2): el online se organiza en salas privadas por enlace o código; matchmaking abierto no es obligatorio.
- [PRE-02.A3](../decision-register.md#pre-02-a3): la experiencia ordinaria es simultánea en tiempo real, con recuperación ante desconexión.
- [PRE-02.A4](../decision-register.md#pre-02-a4): el MVP no incorpora reloj competitivo ni límite de tiempo por partida.
- [PRE-02.B1](../decision-register.md#pre-02-b1): invitados acceden a local y online; registrados obtienen historial persistente y clasificación.
- [PRE-02.B2](../decision-register.md#pre-02-b2): el invitado tendrá alias personalizable e identidad temporal comprobable.
- [PRE-02.B3](../decision-register.md#pre-02-b3): español e inglés forman parte del MVP.
- [PRE-02.B4](../decision-register.md#pre-02-b4): tutorial opcional y consulta completa de reglas acompañarán al juego.
- [PRE-02.C1](../decision-register.md#pre-02-c1): turno, opciones legales y vista previa se mostrarán a partir de datos autorizados por Go.
- [PRE-02.C2](../decision-register.md#pre-02-c2): deshacer queda reservado a práctica reversible, sin extenderse a partidas ordinarias.
- [PRE-02.C3](../decision-register.md#pre-02-c3): se prevé reconexión automática y recuperación de 24 h; después se archiva como inconclusa.
- [PRE-02.C4](../decision-register.md#pre-02-c4): abandono y revancha consensuada se complementan con las tablas precisadas en PRE-03.
- [PRE-02.D1](../decision-register.md#pre-02-d1): las partidas terminadas tendrán historial y reproducción básica de movimientos.
- [PRE-02.D2](../decision-register.md#pre-02-d2): el perfil reúne alias, partidas, victorias, derrotas, empates e historial.
- [PRE-02.D3](../decision-register.md#pre-02-d3): la última partida local se guardará en navegador para reanudar; no se exige offline.
- [PRE-02.E1](../decision-register.md#pre-02-e1): Three.js será el motor visual del tablero 3D.
- [PRE-02.E2](../decision-register.md#pre-02-e2): la referencia premium de RDomino inspira calidad física sin copiar recursos o elementos distintivos.
- [PRE-02.E3](../decision-register.md#pre-02-e3): IA básica, ranking competitivo y estadísticas avanzadas integran MVP ampliado; las otras extensiones citadas quedan pospuestas.

## Agrupaciones de diseño y consecuencias

**Modos y derechos.** Local, IA, online y práctica necesitan límites visibles. Una identidad verificable permite recuperar una partida privada; el alias por sí solo no autoriza acceso. La cuenta aporta persistencia personal y elegibilidad competitiva, sin impedir jugar como invitado.

**Experiencia y aprendizaje.** Las previsualizaciones y ayudas consumen el motor Go; la UI no deriva reglas de su animación. El tutorial se puede omitir y volver a consultar. La última partida local y el historial de práctica son datos de navegador diferentes del historial oficial de Game. Guardado local no elimina la dependencia del motor Go ni promete uso sin internet.

**Lanzamiento.** MVP base y MVP ampliado organizan hitos del mismo compromiso. La IA básica aparece tanto en el modo de juego aprobado como en la ampliación: su mención en dos lugares no permite excluirla del lanzamiento. Torneos, chat, espectadores, IA avanzada, pagos y funciones sociales extensas permanecen pospuestos.

## Sustituciones y matices

PRE-02.C2 amplía el alcance con práctica reversible y mantiene la prohibición de deshacer en juego ordinario. [PRE-03.A1](../decision-register.md#pre-03-a1) y [PRE-03.A2](../decision-register.md#pre-03-a2) amplían C4 con reclamación reglamentaria y tablas acordadas; una repetición no equivale automáticamente a empate. El tiempo registrado por turno para estadísticas no introduce reloj competitivo. La estética se precisa en [PRE-04](PRE-04.md) como realismo físico estilizado sin obligación de fotorrealismo extremo.

## Dependencias y fuentes canónicas

El [dominio 02 del índice](../index.md#spec-02-01) es `PENDIENTE_LOTE_2`. [PRE-03](PRE-03.md) es canónico para reglas/estados; [PRE-06](PRE-06.md), para recuperación y elegibilidad ranked; [PRE-07](PRE-07.md), para perfiles/estadísticas; [PRE-08](PRE-08.md), para enseñanza y práctica. Usar [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [trazabilidad](../traceability.md) al distribuir requisitos.

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: consistencia de la recuperación de invitados, alcance exacto de estados de guardado local, adecuación cultural/licencias y validación normativa del reglamento consultable. [Conflictos](../conflicts.md) conserva discrepancias. No se añaden matchmaking, reloj, espectadores ni operación offline. Los umbrales de usabilidad o rendimiento requieren evidencia futura y no se consideran medidos aquí.

## Criterios futuros de aceptación y verificación

- **AC-01:** una matriz modo × identidad × persistencia × ayudas × Elo distingue todas las combinaciones; solo dos registrados en partida fijada como ranked pueden afectar Elo.
- **AC-02:** los recorridos de inicio, invitación, recuperación, final y revancha incluyen estados negativos y distinguen abandono de archivo inconcluso.
- **AC-03:** tutorial, reglamento, práctica reversible y flujos de juego están definidos en español/inglés y en alternativa HTML.
- **AC-04:** la tabla de hitos conserva todo el MVP ampliado en lanzamiento y marca explícitamente las funcionalidades pospuestas.

La verificación de este lote es textual; la aceptación funcional se demostrará en sesiones posteriores autorizadas.

## Revisión humana

Validar modos, ampliaciones C2/C4 y compromiso completo de lanzamiento. El texto permanece `EN_REVISIÓN`; cualquier recorte requiere cambio aprobado. El coordinador solicitará «APRUEBO LOTE 1» y detendrá la generación posterior.
