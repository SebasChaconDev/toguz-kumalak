# ADR-0004 — PostgreSQL autoritativo y Ably después del commit

ID: ADR-0004 · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner funcional: S (servicios y seguridad).

Dominio: 05/06, consistencia y notificaciones · Lote: 1. Formato y cambios: [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md); propiedad de escritura: [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md).

Decisión base: APROBADO EN FUENTE. Validación técnica: POR_VERIFICAR. La ADR no afirma que exista transacción, conexión a Ably ni garantía de entrega observada.

## Contexto y problema

Un movimiento online cambia datos de la partida y necesita avisar a participantes. PostgreSQL y Ably son sistemas diferentes: no se presupone un commit atómico entre ambos. Notificar antes de confirmar podría mostrar una jugada rechazada; confirmar primero abre una ventana en la que el dato existe y el aviso puede perderse. El diseño aprobado preserva la verdad en Game/PostgreSQL y obliga a reconciliar.

## Fuentes y trazabilidad

- [PRE-01.B4](../decision-register.md#pre-01-b4) exige atomicidad de movimiento y estado; [PRE-05.C1](../decision-register.md#pre-05-c1) define el pipeline autoritativo.
- [PRE-05.C2](../decision-register.md#pre-05-c2) exige bloqueo por partida, versión e idempotencia; [PRE-05.C3](../decision-register.md#pre-05-c3) incluye comando, movimiento, efectos y estado en una transacción.
- [PRE-05.E2](../decision-register.md#pre-05-e2) ordena confirmar PostgreSQL antes de publicar y reconoce aviso perdido sin outbox; [PRE-05.E1](../decision-register.md#pre-05-e1) delimita el evento ligero.
- [PRE-06.B2](../decision-register.md#pre-06-b2) requiere detectar eventos perdidos/desordenados; [PRE-06.E1](../decision-register.md#pre-06-e1) conserva movimientos ante fallo Ably.
- [PRE-06.C3](../decision-register.md#pre-06-c3) cubre varias pestañas; [PRE-06.E4](../decision-register.md#pre-06-e4) exige mediciones antes de SLOs definitivos.

La fuente íntegra está en el [maestro](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md); la [trazabilidad](../traceability.md) mantiene las futuras comprobaciones.

## Decisión delimitada

Game/PostgreSQL es fuente de verdad. Se autentica y autoriza el comando, se controla versión/idempotencia, se ejecuta el motor y se persisten todos los efectos aceptados en una transacción de Game. Solo después de commit confirmado se publica la notificación oficial en Ably. El cliente confirma ejecución/animación definitiva desde la respuesta autoritativa y reconcilia cuando corresponda.

No se adopta un outbox duradero en MVP ni se impone Redis, una cola o un gateway adicional. Sin ese mecanismo no existe garantía de entrega eventual del aviso. La reconciliación REST es parte del flujo aprobado y debe funcionar aunque la publicación no ocurra. Su disponibilidad también depende de Game y PostgreSQL; no se promete continuidad durante una caída de esos sistemas.

## Secuencia y ventanas de fallo

| Momento del fallo | Verdad persistida | Recuperación que debe especificarse |
| --- | --- | --- |
| Antes del commit o transacción revertida | No hay movimiento nuevo confirmado | Respuesta segura o reintento idempotente; ninguna publicación oficial de esa transición |
| Commit confirmado; publicación falla o proceso termina | El movimiento está confirmado | Recuperar por REST; visibilidad de interrupción/reconciliación, sin deshacer el commit |
| Commit confirmado; respuesta HTTP se pierde | El cliente puede desconocer el resultado | Reintento del mismo comando y/o consulta autoritativa sin doble efecto |
| Aviso duplicado, tardío o desordenado | La versión en DB decide | Comparar ID/versión y reconciliar sin aplicar una segunda jugada |
| Resultado de commit desconocido por corte de conexión | No se puede inferir éxito o fracaso | Consultar comando/estado por el mecanismo idempotente; no ejecutar como una nueva orden a ciegas |

La última fila es un caso de análisis técnico, no una nueva garantía de infraestructura. El mecanismo exacto de consulta y las respuestas serán definidos en los contratos futuros.

## Alternativas y motivación

**Histórica descartada en fuente:** Socket.IO fue sustituido por Ably Free, conforme a [PRE-01.B1](../decision-register.md#pre-01-b1). Redis no es infraestructura preventiva obligatoria según [PRE-01.D3](../decision-register.md#pre-01-d3).

**PROPUESTA, no adoptadas:** outbox duradero, reintentos persistentes con cola o un coordinador adicional. Podrían cambiar la ventana de aviso perdido, pero exigen diseño, operación, cuotas y aprobación. Publicar antes de commit contradice directamente PRE-05.E2 y no es una alternativa autorizada.

La motivación trazable es no perder la consistencia de movimientos por un fallo de notificaciones. La consecuencia aceptada en la fuente es una demora o ausencia de aviso hasta reconciliación; no una promesa de que Ably y SQL confirmen juntos.

## Contratos e invariantes

1. Ably notifica; no convierte un movimiento no persistido en oficial. Game publica; los clientes solo suscriben con capacidades restringidas, según [ADR-0005](ADR-0005-session-and-service-trust-boundaries.md).
2. Un reintento o varias pestañas no multiplican efectos. El bloqueo y la versión se comprueban sobre la partida vigente; resultado/Elo siguen [ADR-0008](ADR-0008-ranked-results-and-elo-ownership.md).
3. Un aviso contiene ID, gameId, versión, tipo y datos mínimos conforme a [ADR-0003](ADR-0003-rest-contracts-and-versioned-events.md). Su llegada no reemplaza la autorización para leer una partida por REST.
4. Tras reconexión/foco y verificaciones acotadas se recupera el estado confirmado, mostrando el último movimiento. La frecuencia, límites y estrategia de sondeo quedan POR_VERIFICAR; no se prescribe una cifra ni sondeo ilimitado.
5. El cliente distingue conectado, sincronizando, desconectado y recuperado; bloquea acciones sin confirmar conforme a [PRE-06.B3](../decision-register.md#pre-06-b3). No confunde selección visual con movimiento ejecutado.

## Consecuencias, dependencias y riesgos

La degradación segura prioriza el dato confirmado. El cliente necesita estado explícito de sincronización, lectura REST y tratamiento de duplicados. Son costes de complejidad y consumo de cuotas cuyo impacto no está medido. No se garantizan entrega exactamente una vez, latencia ni cero pérdida de datos; la recuperación de PostgreSQL depende de la política operativa futura.

Depende de [05.03, transacciones](../index.md#spec-05-03), [05.04, publicación posterior al commit](../index.md#spec-05-04), [06.02, reconciliación](../index.md#spec-06-02) y [06.03, recuperación/fallos parciales](../index.md#spec-06-03), PENDIENTE_LOTE_3; [09.01, verificación de datos](../index.md#spec-09-01) y [10.02, observabilidad/cuotas](../index.md#spec-10-02), PENDIENTE_LOTE_5. La validez de partida recuperable 24 horas se comprueba en solicitudes, no depende de la limpieza, según [PRE-06.C2](../decision-register.md#pre-06-c2).

## Validación y aceptación futura

En una sesión autorizada, con PostgreSQL real y doble de Ably, se comprobarán carreras, commit incierto, corte entre commit y publicación y pérdida de respuesta. Integración real con proveedor será independiente. No se ejecutan estas pruebas en este lote.

- **AC-01, positiva:** movimiento válido confirma comando, movimiento, efectos y versión en la misma transacción; el registro de evidencia muestra publicación posterior al commit.
- **AC-02, positiva:** si Ably falla después del commit, REST devuelve el movimiento confirmado y ambos participantes pueden reconstruirlo sin doble efecto.
- **AC-03, negativa:** rollback SQL no produce aviso oficial ni cambio de estado visible como confirmado.
- **AC-04, negativa:** dos comandos concurrentes con versión anterior o un comando repetido no generan dos movimientos por la misma intención idempotente.
- **AC-05, negativa:** un aviso perdido o desordenado no autoriza al cliente a inventar el siguiente estado ni a declarar perdido el dato confirmado.

Trazabilidad de aceptación: AC-01/AC-03/AC-04 → [PRE-05.C2](../decision-register.md#pre-05-c2), [PRE-05.C3](../decision-register.md#pre-05-c3) y [PRE-05.E2](../decision-register.md#pre-05-e2), dueños [05.03](../index.md#spec-05-03)/[05.04](../index.md#spec-05-04); AC-02/AC-05 → [PRE-06.B2](../decision-register.md#pre-06-b2) y [PRE-06.E1](../decision-register.md#pre-06-e1), dueño [06.02](../index.md#spec-06-02); todos PENDIENTE_LOTE_3. Riesgo asociado: [RISK-07](../conflicts.md#risk-07).

## Cambio y reversión documental

Añadir outbox, colas, Redis o garantías nuevas de entrega requiere propuesta, impacto en costes/consistencia y aprobación expresa, registrada en [conflictos](../conflicts.md). Una revisión rechazada restaura la última decisión aprobada sin borrar su fuente. El rollback futuro no debe revertir un movimiento confirmado únicamente porque Ably falló; cualquier reparación de datos requiere su procedimiento y autorización específicos.
