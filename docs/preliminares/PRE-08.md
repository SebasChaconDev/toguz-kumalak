# PRE-08 — IA, práctica y tutorial

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: PRE-08 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner: I — IA.

## Objetivo y contexto

Consolidar 24 decisiones de rival automático y aprendizaje con un único motor Go. La IA debe producir propuestas legales dentro de recursos limitados; Game confirma los movimientos oficiales. No se declara fuerza, tiempo de respuesta ni dificultad demostrada en este lote.

## Fuentes y decisiones

Consultar el [maestro íntegro](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) y el [registro canónico](../decision-register.md). La síntesis siguiente tiene naturaleza `APROBADO`:

- [PRE-08.A1](../decision-register.md#pre-08-a1): tras confirmar jugada humana se calculará automáticamente el turno IA cuando corresponda.
- [PRE-08.A2](../decision-register.md#pre-08-a2): dos dificultades se distinguirán por presupuesto, profundidad o heurística con evidencia.
- [PRE-08.A3](../decision-register.md#pre-08-a3): UI mostrará cálculo, dificultad y última jugada, sin análisis profundo en juego normal.
- [PRE-08.A4](../decision-register.md#pre-08-a4): la IA sigue reglamento hasta fin y no se rinde por valoración heurística.
- [PRE-08.B1](../decision-register.md#pre-08-b1): búsqueda Negamax alfa-beta usará ordenación y profundización iterativa.
- [PRE-08.B2](../decision-register.md#pre-08-b2): heurística explicará kazanes, tuzdyk, vulnerabilidad, movilidad y fase.
- [PRE-08.B3](../decision-register.md#pre-08-b3): transposiciones tendrán memoria temporal acotada y tratamiento correcto de repetición histórica.
- [PRE-08.B4](../decision-register.md#pre-08-b4): tiempo, profundidad, nodos y memoria se limitan con context.Context y mejor jugada legal disponible.
- [PRE-08.C1](../decision-register.md#pre-08-c1): práctica permite ambos bandos, IA opcional y escenarios educativos verificados.
- [PRE-08.C2](../decision-register.md#pre-08-c2): deshacer/rehacer opera por movimiento y habrá atajo para volver al turno humano ante IA.
- [PRE-08.C3](../decision-register.md#pre-08-c3): catálogo inicial pequeño cubre captura, tuzdyk, defensa y final con validación.
- [PRE-08.C4](../decision-register.md#pre-08-c4): última práctica e historial reversible se guardarán localmente con versión.
- [PRE-08.D1](../decision-register.md#pre-08-d1): lecciones progresivas enseñarán turnos, siembra, captura, tuzdyk, final y fundamentos.
- [PRE-08.D2](../decision-register.md#pre-08-d2): sugerencias explicadas son opt-in y están deshabilitadas en ranked.
- [PRE-08.D3](../decision-register.md#pre-08-d3): cada ejercicio distingue legalidad de cumplimiento del objetivo pedagógico y ofrece respuesta.
- [PRE-08.D4](../decision-register.md#pre-08-d4): progreso local será reiniciable/repetible en español e inglés.
- [PRE-08.E1](../decision-register.md#pre-08-e1): panel de práctica agrupa control reversible, reinicio, jugadas y escenarios separado del juego normal.
- [PRE-08.E2](../decision-register.md#pre-08-e2): vistas previas 3D educativas son temporales y no cambian estado confirmado.
- [PRE-08.E3](../decision-register.md#pre-08-e3): reproducción paso a paso reutiliza eventos con anterior/siguiente/pausa/velocidad.
- [PRE-08.E4](../decision-register.md#pre-08-e4): HTML sin WebGL conserva tutorial/práctica mediante el mismo motor Go.
- [PRE-08.F1](../decision-register.md#pre-08-f1): corpus reproducible cubre fases, legalidad, terminales, cancelaciones y límites.
- [PRE-08.F2](../decision-register.md#pre-08-f2): torneos frente a aleatorio legal y referencias medirán fuerza relativa, duración y nodos sin atribuir Elo no validado.
- [PRE-08.F3](../decision-register.md#pre-08-f3): ante fallo AI, Game podrá validar/elegir respaldo legal barato con el motor compartido.
- [PRE-08.F4](../decision-register.md#pre-08-f4): telemetría medirá búsqueda, recursos, cancelación y errores minimizando datos privados.

## Agrupaciones de diseño y consecuencias

**Búsqueda acotada y legalidad.** La profundización iterativa permite conservar una mejor jugada disponible al agotar presupuesto. El contrato debe distinguir estado terminal, cancelación y fallo, sin inventar una jugada cuando no existe movimiento legal. La caché de transposiciones debe considerar la historia reglamentaria relevante: dos tableros iguales pueden tener distinto contexto de reclamación. AI no confirma resultados; Game valida la propuesta contra el estado vigente.

**Práctica reversible.** Ambos bandos y la IA opcional comparten reglas, pero no autoridad competitiva. El historial por estados permite retroceder; jugar desde un punto anterior descarta la rama posterior según [PRE-03.D2](../decision-register.md#pre-03-d2). La última práctica guardada debe versionarse y distinguirse de la última partida local ordinaria. El atajo «volver a mi turno» necesita expresar cuántos movimientos revierte.

**Pedagogía y presentación.** Un movimiento legal puede fallar el objetivo educativo sin ser error reglamentario. Ejercicios necesitan posición inicial, objetivo, respuesta y explicación verificables. Previsualización y reproducción presentan eventos sin mutar la partida. HTML y 3D comparten flujos y contratos, con textos bilingües y progreso reiniciable.

## Sustituciones y matices

Negamax concreta la familia Negamax/Minimax aprobada en [PRE-03.D3](../decision-register.md#pre-03-d3); no introduce motor distinto. Fácil/normal requieren diferencia demostrable, sin prometer Elo, profundidad fija o latencia antes de medir. Las ayudas no se habilitan en ranked. La práctica reversible no extiende deshacer al juego ordinario. Almacenamiento local no aprueba offline ni guardar tokens sensibles en navegador.

## Dependencias y fuentes canónicas

El [dominio 08 del índice](../index.md#spec-08-01) es `PENDIENTE_LOTE_4`. Depende del reglamento/versionado de [PRE-03](PRE-03.md), contratos de [PRE-05](PRE-05.md), UI/accesibilidad de [PRE-04](PRE-04.md), pruebas de [PRE-09](PRE-09.md) y límites de [PRE-10](PRE-10.md). Aplicar [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [trazabilidad](../traceability.md).

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: reglamento/corpus, efecto de historial en transposiciones, presupuesto real Vercel, fuerza relativa, recuperación ante respuesta obsoleta/cancelada y validez educativa. Véase [conflictos](../conflicts.md). No se ejecutan búsqueda, torneos, benchmarks, modelos, servicios ni tests; todos sus resultados siguen pendientes.

## Criterios futuros de aceptación y verificación

- **AC-01:** contratos distinguen jugada válida, terminal, límite, cancelación y fallback; ninguna propuesta obsoleta puede confirmar estado oficial.
- **AC-02:** corpus con semillas y presupuestos reproducibles permite comparar fácil/normal y registrar legalidad, tiempo, nodos y fuerza relativa sin Elo inventado.
- **AC-03:** práctica define reversión, bifurcación, reinicio, persistencia versionada y recuperación de datos incompatibles sin efecto competitivo.
- **AC-04:** cada lección/ejercicio identifica fuente, objetivo y explicación; el flujo bilingüe está disponible también sin WebGL y bloquea sugerencias en ranked.

La aceptación futura requiere evidencia observada; esta consolidación no verifica IA ni aprendizaje.

## Revisión humana

Revisar fidelidad pedagógica, fronteras de ayudas y definición de dificultad antes de especificarlas en detalle. El texto permanece `EN_REVISIÓN`; el coordinador pedirá «APRUEBO LOTE 1» y detendrá el trabajo autorizado al cerrar el lote.
