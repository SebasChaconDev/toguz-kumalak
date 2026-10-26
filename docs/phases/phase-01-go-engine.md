# PH-01 — Motor Go puro y corpus reglamentario

> **Aprobación posterior — 2026-10-26: APROBADO documentalmente.** El propietario declaró «APRUEBO LOTE 5 como documentación técnica.» sobre HEAD 9ff1a2857d0b91b9cf7925a4c36de30e377cef3d; PR #5 integrado por squash 6579ca930a9081f773b390ce61f6bc80cc6b1ae7. El [recibo de aprobación](../../PLANS.md#cierre-lote-5) prevalece sobre los estados de revisión históricos. L4-P01, la política de retención y demás propuestas mantienen su aprobación específica pendiente; sin permiso de implementación, workflows, proveedores o lote6.

ID: PH-01 · Versión: 0.2 · Fecha: 2026-10-26 · Estado documental: APROBADO · Owner: C — Coordinación · Lote: 5.

Plan de implementación futura PROPUESTA; reglas y decisiones fuente APROBADO, ejecución POR_VERIFICAR. No se crea código o test en este lote documental.

## Ficha Milestone/1

| Campo requerido | Contenido del hito |
|---|---|
| Identificador | PH-01 |
| Objetivo | Construir y demostrar un motor Go puro, determinista y versionado antes de tratarlo como autoridad implementada. |
| Dependencias | PH-00; 03.01–03.04, 01.01, 02.01, 09.01/09.03/09.04 y 11.03; AD-NORM-001/TK-N1-1.0.0/14 ACC. |
| Responsable | M implementa; Q coordina verificación y revisor independiente contrasta oráculo; C coordina alcance, propietario acepta. |
| Entrada | Revisión documental de entrada 11.03 y permiso de código concreto; corpus normativo congelado, contexto/contratos versionados y rutas propias asignadas. |
| Entregables | Tipos/validadores, fábrica inicial, transición pura, legales/preview, contexto de repetición, eventos, replay y suites reproducibles. |
| Pruebas positivas/negativas |14 ACC completos, propiedades/fuzzing/regresión/replay, matrices abajo y rutas críticas aun con cobertura suficiente. |
| Riesgos | Mutación por aliasing, índices/cámara, precedencia, hash como igualdad, historia incompleta o corpus sintético presentado como alcanzable. |
| Evidencia obligatoria | QA-Evidence/1 con corpus/procedencia/semillas, esperado/observado, versiones/digest motor, comandos/salidas y cobertura por paquete. |
| Salida | DoD 09.04 observado, objetivo90% sentencias motor y críticos completos; revisión independiente sin P0/P1 o críticos pendientes. |
| Aprobación humana | GateEvaluation/1 y dossier del commit exacto; aceptación no autoriza PH-02 o recursos por inferencia. |
| Corrección/recuperación | Conservar contraejemplo y run fallido; corregir implementación contra fuente y ejecutar regresión afectada. Divergencia normativa real se registra sin cambiar perfil unilateralmente. |

## Fuentes y contratos

Leer [03.01](../specs/03-game-engine/03.01-normative-rules-and-corpus.md), [03.02](../specs/03-game-engine/03.02-state-indices-and-invariants.md), [03.03](../specs/03-game-engine/03.03-moves-and-semantic-events.md) y [03.04](../specs/03-game-engine/03.04-outcomes-repetition-and-replay.md) completos, incluida AD-NORM-001 y ACC-N1-01–14. [01.01](../specs/01-infrastructure/01.01-stack-monorepo.md), [02.01](../specs/02-product/02.01-scope-and-modes.md) y [ADR-0002](../adr/ADR-0002-single-versioned-go-engine.md) preservan módulo único y alcance. Fuentes PRE-03 y PRE-09.A1–A4/B1/B4 se consultan en [REG](../decision-register.md), sin mover dueños.

Consumir [09.01](../specs/09-quality/09.01-engine-data-contract-and-ai-verification.md) F-Q1-N01–N14/F-Q1-01/02, [09.03](../specs/09-quality/09.03-ci-security-and-release-controls.md) para consumidores Game/AI y [09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md) DoD. Aplican [11.03](../specs/11-orchestration/11.03-milestones-gates-and-evidence.md), [PLANS](../../PLANS.md#ejecucion-lote-5), [CONF](../conflicts.md) y [PH-00](phase-00-feasibility-spike.md). Las propuestas internas de las specs conservan su naturaleza; fijar nombres/versiones dentro del alcance autorizado antes de consumidores.

Contratos consumidos: State/PitRef/Move/Transition/RuleError/LegalMovesView/PreviewView, SemanticEvent, Terminal/PositionKey/RepetitionContext y ReplayHeader/MoveRecord/LifecycleRecord/Checkpoint. No mezclar versión de reglas, formato, comando, revisión, ordinal de movimiento o historial. Game autenticará/confirmará después; el motor no consulta identidad/SQL/reloj/red.

## Componentes y tickets ordenados

Rutas candidatas de 01.01: modules/engine y su corpus/suites dentro del módulo; nombres y archivos concretos se fijan en TaskEnvelope/1 posterior. No son rutas creadas aquí. Sin motor TypeScript, Gin/pgx en el núcleo o RNG/reloj ocultos.

| Ticket / depende de | Pasos y razón | Entrega pequeña / aceptación |
|---|---|---|
| PH-01-T01 / entrada | Congelar CaseSpec/1/CorpusEntry/1 con fuente, alcance y esperados de14 ACC; separar SOURCE_DIRECT/DERIVED_MANUAL/SYNTHETIC_CONTRACT y alcanzabilidad. | Corpus revisado; no usar suma 162 como prueba de historia legal. |
| PH-01-T02 / T01 | TDD de Player/PitRef/State/versiones, validación previa a índices/conversiones y fábrica inicial; copiar profundamente contexto/slices/maps. | Inicial18×9, blanco, kazanes0; inválidos/versiones desconocidas sin mutación. |
| PH-01-T03 / T02 | Implementar siembra piedra única/N−1 y captura par rival con eventos desde efectos. | R-01–R-03, varias vueltas y retorno al origen; F-Q1-N01/02. |
| PH-01-T04 / T03 | Implementar tuzdyk automático, prohibiciones9/homólogo/segundo y depósitos al dueño. | R-04/F-Q1-N03–N05; prohibición no vuelve ilegal la jugada. |
| PH-01-T05 / T04 | Completar todos los depósitos y aplicar target→equal→turno rival→atsyrau; conservar remanentes. | F-Q1-N06–N09/N14; sin barrido ni movimiento adicional. |
| PH-01-T06 / T05 | Construir igualdad exacta/ledger candidato de repetición y consulta de reclamabilidad; no adjudicar draw automático. | F-Q1-N10/N12/N13; inicial una, solo salida activa candidata. |
| PH-01-T07 / T06 | Legales/preview usan el mismo núcleo; replay/checkpoint reproducen reglas/contexto/eventos versionados y detectan gaps. | F-Q1-02 y eventos exactos; incompatibilidad sin reparación silenciosa. |
| PH-01-T08 / T07 | Ejecutar propiedades, fuzzing con semilla/corpus persistible autorizado, race cuando aplique y cobertura; conservar mínimos contraejemplos y revisar. | F-Q1-01; 90% motor más críticos y dossier, sin resultado inventado. |

T01 precede código para fijar oráculo; T03–T06 siguen efectos y precedencia. Separar tareas de integración Game de núcleo puro evita fingir que una prueba unitaria demuestra una carrera SQL.

## Matriz14 ACC y frontera con Game

Cada caso canónico F-Q1-Nnn de 09.01 contiene positivo, negativo y evidencia; estas filas fijan dónde se consumen sin duplicar sus esperados.

| ACC / caso 09.01 | Tickets / comportamiento que debe observarse |
|---|---|
| ACC-N1-01 / F-Q1-N01 | T02/T03: inicial, única piedra, origen inválido sin efectos. |
| ACC-N1-02 / F-Q1-N02 | T03: N−1/captura par rival; no captura propia o impar. |
| ACC-N1-03 / F-Q1-N03 | T04: creación automática al3 permitido. |
| ACC-N1-04 / F-Q1-N04 | T04: tres prohibiciones independientes dejan3 normal. |
| ACC-N1-05 / F-Q1-N05 | T04: cada depósito a tuz va al dueño, cero lógico/162. |
| ACC-N1-06 / F-Q1-N06 | T05: siguiente sin piedras K60 pierde sin barrido. |
| ACC-N1-07 / F-Q1-N07 | T05: siguiente K81 empata atsyrau conservando remanente. |
| ACC-N1-08 / F-Q1-N08 | T05: actor vacío/rival con jugadas sigue activo. |
| ACC-N1-09 / F-Q1-N09 | T05: siembra completa y prioridad target/equal antes de atsyrau. |
| ACC-N1-10 / F-Q1-N10 | T06:≥3 habilita derecho, no draw automático. |
| ACC-N1-11 / F-Q1-N11 | T06 define predicado; integración PH-02/04 debe demostrar ambos reclamantes, revisión y carrera real. |
| ACC-N1-12 / F-Q1-N12 | T06: igualdad incluye reglas/pozos/kazanes/tuz/turno; colisión hash compara exacto. |
| ACC-N1-13 / F-Q1-N13 | T06/T07: preview/retry/cierre no suma apariciones; integración posterior conserva lifecycle/ordinal del cierre una vez. |
| ACC-N1-14 / F-Q1-N14 | T05/T06: terminal prevalece, sin detector/aparición posterior. |

PH-01 prueba toda la lógica normativa que corresponde al motor. No marca como observado el lock/receipt/autorización de ACC11/13: registra esos subcasos de integración requeridos en [PH-02](phase-02-contracts-persistence-guests.md)/[PH-04](phase-04-private-multiplayer.md). La salida del motor no equivale al cierre de todas las pruebas de producto de14 ACC ni a autoridad oficial Game implementada.

## Casos adicionales, evidencias y seguridad

| Caso local | Positivo | Negativo | Evidencia futura |
|---|---|---|---|
| PH01-C01 | Secuencias legales de semilla fija conservan162/determinismo. | Entrada/contexto compartido o rechazado no cambia, incluso con aliasing inducido. | F-Q1-01, snapshots/copia, semilla y minimización. |
| PH01-C02 | Mismo perfil y contexto producen mismo estado/eventos. | Cambiar historial relevante no se ignora por igual tablero/count. | Contextos completos y diff; 03.04. |
| PH01-C03 | Replay desde inicial/checkpoint reproduce primer ordinal a último. | Gap/versiones/desorden o hash colisionado no produce verified. | F-Q1-02, checkpoints y primer punto divergente. |
| PH01-C04 | Fuzz genera estructuras y secuencias alcanzables en clases separadas. | Sintético estructural válido no se presenta como DERIVED_PATH_REVIEWED. | CorpusEntry/1, cobertura y contraejemplos. |

QA-Evidence/1 registra versión/digest motor, corpus y herramientas exactas, comandos sanitizados, expected/observed y artefactos accesibles/hash; AcceptanceDossier/1 consolida críticos pendientes y frontera Game. 90% es objetivo aprobado de sentencias, sin sustituir casos negativos. Benchmarks se hacen después de corrección y sin umbral inventado; no se altera reglamento por rendimiento.

Recursos locales/CI/pruebas deben estar incluidos en el permiso de implementación y política externa aplicable; no instalar/arrancar contenedores o servicios por este documento. Datos sintéticos, tamaño/trabajo acotados, sin secretos ni dependencias de proveedores en el motor. RISK-01 no reabre reglas adoptadas: una discrepancia de implementación se corrige; una nueva disputa normativa se comunica a C/propietario y bloquea solo lo afectado.

## DoD, gate y recuperación

DoD 09.04: requisito/caso trazado, motor implementado y resultados observados en commit exacto, cobertura90% y críticos, docs/versiones actualizadas, revisión independiente y evidencia suficiente. FAIL/BLOCKED/NOT_RUN requerido impide salida; exit0 solo no la prueba. P0/P1 y críticos no verificados bloquean; P2 diferido exige owner/plazo/aceptación humana y no exime normas.

C entrega GateEvaluation/1, Q verifica y propietario decide el hito/continuación. Conservar fallo→corrección→nuevo runId enlazado; no cambiar esperado a output ni borrar el contraejemplo. A coordina incidentes operativos, C gobierno y Q recuperación según 10.04. No reparar una partida futura cambiando motor/replay histórico silenciosamente. [PH-02](phase-02-contracts-persistence-guests.md) requiere su alcance de código autorizado; este plan no lo concede.
