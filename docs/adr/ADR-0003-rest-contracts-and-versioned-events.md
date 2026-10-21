# ADR-0003 — Contratos REST y eventos versionados

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: ADR-0003 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner funcional: S (servicios y seguridad).

Dominio: 05/06, contratos · Lote: 1. Formato y cambios: [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md); propiedad de escritura: [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md).

Decisión base: APROBADO EN FUENTE. Validación técnica: POR_VERIFICAR. La ADR aún no es una especificación aprobada ni prueba de que existan contratos generados.

## Contexto y problema

Go, TypeScript y proveedores deben intercambiar comandos y notificaciones sin depender de estructuras SQL ni de instrucciones de renderizado. Las carreras, reintentos, eventos fuera de orden y despliegues independientes requieren separar versión de API, versión de esquema de evento y versión de estado de partida. La existencia de tipos compilables no demuestra autorización, consistencia ni compatibilidad semántica.

## Fuentes y trazabilidad

- [PRE-05.B1](../decision-register.md#pre-05-b1) define REST `/api/v1`; [PRE-05.B2](../decision-register.md#pre-05-b2) fija OpenAPI 3.1 y modelos/clientes Go y TypeScript con versiones fijadas.
- [PRE-05.B3](../decision-register.md#pre-05-b3) establece Problem Details RFC 9457; [PRE-05.B4](../decision-register.md#pre-05-b4) exige recursos/dominios y paginación sin exponer tablas.
- [PRE-05.E1](../decision-register.md#pre-05-e1) delimita eventos Ably ligeros; [PRE-05.E4](../decision-register.md#pre-05-e4) pide JSON schemas versionados y validación en ambos lenguajes.
- [PRE-03.B3](../decision-register.md#pre-03-b3) exige casilla origen, versión esperada e identificador idempotente; [PRE-06.B2](../decision-register.md#pre-06-b2) exige ID y versión para detectar desorden/pérdidas.
- [PRE-09.C1](../decision-register.md#pre-09-c1) requiere compatibilidad de contratos y consumidores; [PRE-11.E4](../decision-register.md#pre-11-e4) somete cambios relevantes a aprobación.

El [maestro archivado](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) conserva la fuente y la [trazabilidad](../traceability.md) registra especificaciones y evidencia pendiente.

## Decisión delimitada

HTTP se expresa como REST versionado bajo `/api/v1`, con consultas de recursos y comandos explícitos. OpenAPI 3.1 es la fuente de verdad HTTP para validación y generación Go/TypeScript. Se fijarán las versiones de herramientas en una fase autorizada; esta ADR no selecciona generadores ni crea manifiestos.

Los errores usan Problem Details RFC 9457 con código de dominio, requestId y mensajes seguros. El contrato deberá distinguir entrada inválida, falta de permiso, versión obsoleta y conflictos de comando sin filtrar secretos ni detalles de persistencia. Los códigos HTTP exactos y campos finales se resolverán en las especificaciones de servicio, sin dar por aprobada una lista inventada aquí.

Los eventos tienen JSON schemas versionados y validación en Go/TypeScript. Expresan identificador, partida, versión, tipo y datos mínimos necesarios. No transportan autoridad para aceptar un movimiento del cliente ni instrucciones 3D persistentes. REST devuelve o recupera el estado autoritativo de Game; el evento avisa al consumidor y activa reconciliación cuando corresponde.

## Alternativas y motivación

**Histórica descartada en fuente:** Socket.IO fue sustituido por Ably Free en [PRE-01.B1](../decision-register.md#pre-01-b1). Es una sustitución del transporte de notificaciones, no una aprobación de usar eventos como comandos autoritativos.

**PROPUESTA, no adoptadas:** GraphQL, RPC como contrato público, compartir structs internos de base como API, o generar tipos sin validar entradas en runtime. No consta que el propietario haya comparado estas opciones. Adoptarlas requeriría cambio explícito y análisis de compatibilidad/operación.

La motivación trazable está en PRE-05.B2/E4: una definición contractual compartida para dos lenguajes. La separación de recursos y tablas de PRE-05.B4 evita convertir una migración interna en contrato público de manera accidental. El versionado permite planificar compatibilidad; no la garantiza sin pruebas de consumidores.

## Contratos e invariantes

| Concepto | Función | No debe confundirse con |
| --- | --- | --- |
| Versión API `/api/v1` | Familia del contrato HTTP | Número de movimiento o build desplegado |
| Versión de esquema de evento | Forma/semántica que valida el consumidor | Estado vigente de una partida |
| Versión de partida | Concurrencia y detección de actualización | Política de compatibilidad del API |
| Versiones motor/estado/registros | Reproducción y reglas históricas | Orden de entrega de Ably |
| Identificador de comando/evento | Deducción de duplicados y correlación según su contrato | Prueba de autorización |

Un comando de movimiento aporta origen, versión esperada e identificador idempotente dentro de una sesión autorizada. No aporta un tablero que sustituya la verdad de Game. El alcance de la clave idempotente, su retención, el tratamiento de una misma clave con otro contenido y la respuesta a un reintento quedan pendientes de diseño detallado; deben impedir doble efecto y ambigüedad.

El consumidor compara ID/versión conforme al contrato; ante pérdida, duplicado o desorden no aplica ciegamente transiciones. Reconciliará por REST como describe [ADR-0004](ADR-0004-postgres-authority-and-postcommit-ably.md). La política concreta de compatibilidad, cambios aditivos, retirada y versiones desconocidas sigue POR_VERIFICAR; no se presupone tolerancia ilimitada a campos o tipos nuevos.

## Consecuencias y dependencias

Se requiere mantener HTTP y esquemas de eventos, comprobar consumidores y coordinar despliegues compatibles. Los modelos generados ayudan con forma y tipado, pero no reemplazan autorización por recurso, invariantes del motor ni transacciones. La observabilidad correlaciona peticiones sin exponer datos sensibles.

Depende de [03.02, versiones de estado](../index.md#spec-03-02), PENDIENTE_LOTE_2; [05.02, HTTP/OpenAPI](../index.md#spec-05-02), [05.04, esquemas de eventos](../index.md#spec-05-04), [06.02, reconciliación](../index.md#spec-06-02) y [07.02, autorización](../index.md#spec-07-02), PENDIENTE_LOTE_3. [ADR-0001](ADR-0001-three-go-services-and-data-ownership.md) fija propietarios; [ADR-0005](ADR-0005-session-and-service-trust-boundaries.md) fija autoridades de credenciales. [09.01, verificación de contratos](../index.md#spec-09-01), PENDIENTE_LOTE_5, deberá aportar evidencias.

## Riesgos y validación futura

Quedan POR_VERIFICAR el soporte real de OpenAPI 3.1 de los generadores elegidos, validación JSON Schema, semántica de evolución y compatibilidad de despliegues. En una sesión autorizada se contrastarán documentación oficial vigente y contratos generados; se probarán entradas inválidas, consumidores antiguos, tipos desconocidos, duplicados y claves idempotentes reutilizadas con otro contenido. Aquí solo se comprueba coherencia documental.

## Aceptación futura

- **AC-01, positiva:** un contrato HTTP acordado genera modelos Go/TypeScript compatibles y valida una petición/respuesta de referencia; se conserva evidencia de versiones y validación.
- **AC-02, positiva:** un evento de referencia valida en ambos lenguajes y permite identificar partida, esquema y versión sin conocer SQL ni Three.js.
- **AC-03, negativa:** petición mal formada o no autorizada produce error seguro y trazable sin modificar el estado.
- **AC-04, negativa:** evento desordenado, duplicado o con versión incompatible no corrompe la vista; el cliente recupera estado confirmado mediante su política documentada.
- **AC-05, negativa:** un cambio incompatible no se despliega como si mantuviera el contrato anterior sin plan de compatibilidad y autorización correspondiente.

Trazabilidad de aceptación: AC-01/AC-03 → [PRE-05.B2](../decision-register.md#pre-05-b2) y [PRE-05.B3](../decision-register.md#pre-05-b3), dueño [05.02](../index.md#spec-05-02); AC-02 → [PRE-05.E4](../decision-register.md#pre-05-e4), dueño [05.04](../index.md#spec-05-04); AC-04 → [PRE-06.B2](../decision-register.md#pre-06-b2), dueño [06.02](../index.md#spec-06-02); todos PENDIENTE_LOTE_3. AC-05 → [PRE-09.C1](../decision-register.md#pre-09-c1), dueño [09.01](../index.md#spec-09-01), PENDIENTE_LOTE_5.

## Cambio y reversión documental

Modificar autoridad, protocolo público o política de compatibilidad requiere propuesta e impacto en productores/consumidores, aprobación y actualización de [conflictos](../conflicts.md) y trazabilidad. Ante rechazo, restaurar la última redacción aprobada conservando la propuesta identificada. El rollback futuro de despliegues deberá respetar compatibilidad de contratos y datos; no se autoriza por esta ADR ni se ha ensayado.
