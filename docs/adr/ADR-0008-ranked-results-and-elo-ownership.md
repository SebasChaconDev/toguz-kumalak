# ADR-0008 — Resultados clasificatorios y propiedad de Elo en Game

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: ADR-0008 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner funcional: S (servicios y seguridad).

Dominio: 06/07, resultados y clasificación · Lote: 1. Formato y cambios: [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md); propiedad de escritura: [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md).

Decisión base: APROBADO EN FUENTE. Validación técnica: POR_VERIFICAR. La fórmula detallada, concurrencia y controles de abuso aún necesitan especificación y pruebas; esta ADR no afirma un ranking operativo.

## Contexto y problema

Una partida puede terminar por reglas, abandono, acuerdo o desconexión inconclusa. No todos los resultados pertenecen al ranking, y una vinculación posterior de invitado no debe reescribir la elegibilidad competitiva. Resultado, Elo y agregados necesitan una única autoridad y efectos atómicos, incluso con reintentos o varias partidas terminando simultáneamente para el mismo jugador.

## Fuentes y trazabilidad

- [PRE-06.D1](../decision-register.md#pre-06-d1) limita Elo a partidas expresamente clasificatorias entre registrados y fija la condición al comienzo; [PRE-06.D2](../decision-register.md#pre-06-d2) separa Elo de estadísticas generales.
- [PRE-06.D3](../decision-register.md#pre-06-d3) exige medidas auditables contra enfrentamientos repetidos/manipulación; [PRE-06.D4](../decision-register.md#pre-06-d4) distingue abandono y desconexión inconclusa.
- [PRE-07.A3](../decision-register.md#pre-07-a3) impide Elo retroactivo al vincular invitado; [PRE-07.E1](../decision-register.md#pre-07-e1) fija 1200, K=32 y diez primeras clasificatorias provisionales, sujetos a simulación/validación.
- [PRE-07.E4](../decision-register.md#pre-07-e4) exige que Game actualice resultado, Elo y agregados en transacción atómica e idempotente; [PRE-03.C4](../decision-register.md#pre-03-c4) separa resultado y causa.
- [PRE-07.E2](../decision-register.md#pre-07-e2) y [PRE-07.E3](../decision-register.md#pre-07-e3) cubren estadísticas y ranking; [PRE-07.F4](../decision-register.md#pre-07-f4) exige pruebas de manipulación.

El [maestro y sus controles de ranking](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) conserva la fuente y la [trazabilidad](../traceability.md) identifica validaciones pendientes.

## Decisión delimitada

Game es propietario de partidas, resultados, Elo y agregados. Finalizar una partida y aplicar todos sus efectos elegibles ocurre dentro de una transacción de Game con reintentos idempotentes. Identity mantiene identidad y perfiles, sin escribir ratings de Game ni compartir una transacción entre bases.

Solo una partida expresamente clasificatoria entre dos registrados puede afectar Elo, con elegibilidad fijada al comienzo. Invitados, IA, práctica y partidas casuales quedan fuera de Elo. Vincular después una cuenta puede trasladar partidas elegibles al historial según la fuente, pero no las convierte retroactivamente en clasificatorias.

El abandono voluntario es derrota. Una partida archivada como inconclusa por desconexión no modifica Elo. Se almacena resultado separado de causa, sin convertir ausencia temporal en abandono ni equiparar inconclusa a empate. La recuperación de 24 h y su vencimiento persistido siguen [PRE-06.C1](../decision-register.md#pre-06-c1) y [PRE-06.C2](../decision-register.md#pre-06-c2).

Se conservan parámetros aprobados: Elo inicial 1200, K=32 y primeras diez partidas clasificatorias provisionales. La simulación debe validar su comportamiento; no autoriza alterarlos sin aprobación. Redondeo, precisión, cómputo de provisionales, orden de actualización y casos administrativos no quedan fijados por esta ADR y son POR_VERIFICAR.

## Alternativas y motivación

**Histórica descartada en fuente:** no consta una implementación previa de Elo en Identity ni otra fórmula sustituida. No se inventa una migración histórica. La frontera de elegibilidad procede directamente de PRE-06.D1 y PRE-07.A3.

**PROPUESTA, no adoptadas:** servicio separado de ratings, actualización asíncrona de Elo fuera de la transacción de resultado, clasificación retroactiva o modificar K para ciertos usuarios. Exigirían propuesta y aprobación; las dos primeras además revisarían propiedad/atomicidad aprobadas, y la retroactividad contradice PRE-07.A3.

La motivación trazable es la integridad competitiva de PRE-07.E4 y la separación de estadísticas generales de PRE-06.D2. Tener un historial de una partida no la hace clasificatoria. Hacer públicos Elo/estadísticas no autoriza publicar correos o historiales detallados ajenos.

## Contratos e invariantes

1. Elegibilidad, participantes y modalidad fijados al inicio permiten decidir si corresponde actualización competitiva al cierre; un cliente no decide su rating ni cambia retrospectivamente esa condición.
2. Cada finalización aceptada produce una sola aplicación de resultado, Elo y agregados. Repetir el mismo comando o evento no añade otra victoria, derrota, empate o provisional.
3. Un fallo en la actualización de cualquier efecto obligatorio revierte la transacción; no queda resultado confirmado con un rating parcialmente aplicado dentro del modelo aprobado.
4. Las estadísticas generales pueden incluir modalidades no competitivas según su especificación, conservando etiqueta de modalidad y causa. No se inventa aquí qué agregado incluye cada modalidad.
5. El ranking público paginado, búsqueda por alias, posición personal y provisionales consume datos de Game y perfiles públicos por API de Identity. No accede a Identity DB, conforme a [ADR-0001](ADR-0001-three-go-services-and-data-ownership.md).
6. Ably se publica después del commit y no decide resultado/Elo, conforme a [ADR-0004](ADR-0004-postgres-authority-and-postcommit-ably.md). La pérdida del aviso no repite el efecto competitivo.

## Concurrencia y abuso pendientes de diseño

El bloqueo de una partida por sí solo no resuelve dos partidas distintas que actualizan simultáneamente el rating del mismo jugador. La especificación futura debe analizar lecturas obsoletas, actualizaciones perdidas, bloqueo de ambos jugadores, interbloqueos, reintentos y orden observable de finalizaciones. **PROPUESTA, no adoptadas:** bloqueo ordenado de registros de ratings o estrategia equivalente de aislamiento/reintento. La elección debe demostrar atomicidad sin dar por existente una transacción distribuida ni imponer un mecanismo no aprobado.

PRE-06.D3 exige controles auditables de enfrentamientos repetidos y manipulación. Quedan pendientes los límites concretos, señales, ventanas, tratamiento de falsos positivos y eventual intervención administrativa. No se inventa un número de partidas permitido ni una política automática de sanción. Las medidas requieren evaluación de privacidad y evidencia antes de presentarlas como eficaces.

Las carreras entre finalización, abandono, reclamación de tablas, vencimiento o eliminación de cuenta también necesitan orden y respuesta definidos. Una repetición reglamentaria dudosa se registra en [conflictos](../conflicts.md) y bloquea el cierre del algoritmo afectado; no se resuelve ajustando Elo.

## Consecuencias, dependencias y validación futura

Los datos competitivos viven junto al resultado de Game y requieren más disciplina transaccional y evidencia. No hay garantía actual de ausencia de abuso, capacidad, equidad o consistencia bajo carga.

Depende de [03.04, resultados](../index.md#spec-03-04), PENDIENTE_LOTE_2; [05.03, transacciones](../index.md#spec-05-03), [06.01, ciclo clasificatorio](../index.md#spec-06-01), [07.01, vinculación](../index.md#spec-07-01) y [07.04, Elo/estadísticas](../index.md#spec-07-04), PENDIENTE_LOTE_3; [09.01, verificación de datos](../index.md#spec-09-01) y [10.04, incidencias](../index.md#spec-10-04), PENDIENTE_LOTE_5. Deberán simularse parámetros y probarse carreras con PostgreSQL real, duplicados, rollback parcial, vinculación invitada y borrado distribuido. No se ejecutan pruebas de producto en este lote.

## Aceptación futura

- **AC-01, positiva:** al terminar una clasificatoria entre registrados, Game confirma resultado, cambios de ambos ratings y agregados en un único commit; existe evidencia de valores antes/después y parámetros aplicados.
- **AC-02, positiva:** abandono válido produce derrota y los efectos elegibles correspondientes; una inconclusa por vencimiento conserva causa y no cambia Elo.
- **AC-03, negativa:** invitado, IA, práctica, casual o vínculo posterior nunca obtiene actualización retroactiva de Elo.
- **AC-04, negativa:** repetición del cierre, pérdida de respuesta o aviso duplicado no repite resultado, estadísticas ni provisional.
- **AC-05, negativa:** dos finales concurrentes con jugador compartido no pierden actualizaciones ni confirman efectos parciales; la evidencia describe el orden y los reintentos de la estrategia elegida.
- **AC-06, negativa:** un fallo inducido entre efectos revierte todos los cambios obligatorios de esa finalización.
- **AC-07, negativa:** una regla de abuso aún propuesta no se presenta como sanción definitiva ni como protección probada; requiere política aprobada y evaluación auditable.

Trazabilidad de aceptación: AC-01/AC-04/AC-05/AC-06 → [PRE-07.E4](../decision-register.md#pre-07-e4), dueño [07.04](../index.md#spec-07-04); AC-02 → [PRE-06.D4](../decision-register.md#pre-06-d4), dueño [06.01](../index.md#spec-06-01); AC-03 → [PRE-06.D1](../decision-register.md#pre-06-d1) y [PRE-07.A3](../decision-register.md#pre-07-a3), dueños [06.01](../index.md#spec-06-01)/[07.01](../index.md#spec-07-01); AC-07 → [PRE-06.D3](../decision-register.md#pre-06-d3), dueño [07.04](../index.md#spec-07-04); todos PENDIENTE_LOTE_3. Riesgo asociado: [RISK-08](../conflicts.md#risk-08).

## Cambio y reversión documental

Cambiar elegibilidad, parámetros, propietario o atomicidad exige propuesta, simulaciones/evidencias, impacto sobre historiales y aprobación expresa. Registrar conflictos y actualizar trazabilidad. Una revisión rechazada se revierte documentalmente sin reescribir las fuentes. Recalcular ratings históricos o reparar datos no forma parte de este lote y necesita plan específico, semántica versionada y autorización; rollback de código no implica deshacer resultados válidos.
