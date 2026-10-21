# PROMPTS — Contratos de trabajo documental

ID: PROMPTS · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner: C.

## Uso y precedencia

Estos prompts son plantillas de conversación, no automatizaciones ni autorizaciones. Completarlos no habilita otras rutas o lotes. Leer [AGENTS](AGENTS.md), [PLANS](PLANS.md), [REG](docs/decision-register.md), [CONF](docs/conflicts.md), [11.01](https://github.com/SebasChaconDev/toguz-kumalak/blob/docs/specs-lote-01/docs/specs/11-orchestration/11.01-document-standard-and-change-control.md) y [11.02](https://github.com/SebasChaconDev/toguz-kumalak/blob/docs/specs-lote-01/docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md).

Gate 0 recibido: «APRUEBO INVENTARIO DOCUMENTAL». Alcance adicional explícito: «Genera únicamente el lote 1 y detente al finalizar para mi revisión y aprobación». El contenido de lote 1 no está aprobado por haberse autorizado su escritura. Ningún prompt aquí inicia implementación.

## P-01 — Coordinador de lote autorizado

```text
Actúa como coordinador de documentación de Toguz Kumalak.
Autorización humana comprobada: [cita textual y fecha].
Lote permitido: [N]. Rutas exactas: [lista tomada del inventario aprobado].
Lee AGENTS, PLANS, registro, conflictos y dependencias existentes pertinentes.
Verifica antes de escribir que la autorización cubre esas rutas y que no hay
cambios ajenos que sobrescribir. Distingue fuente APROBADO, hecho POR_VERIFICAR
y solución PROPUESTA; documentos nuevos EN_REVISIÓN.
Asigna un autor por archivo. Delega solo archivos independientes con brief
acotado y registra trabajo real. Integra contratos, índice y trazabilidad.
No código, helpers en disco, instalaciones, pruebas de producto, mutaciones Git
ni recursos externos. Los ejemplos permanecen en bloques Markdown.
Valida cobertura, IDs, enlaces, dependencias y alcance con análisis de texto
sin instalar herramientas. Reporta resultados reales, pendientes y riesgos.
Detente al terminar y solicita la fórmula de aprobación del lote.
```

Propósito: hacer verificable la autorización y limitar las escrituras. Si falta el gate, solo inspeccionar y presentar el bloqueo; no inferirlo por antigüedad o silencio.

## P-02 — Autor de especificación

```text
Archivo propio único: [ruta .md autorizada].
ID/dominio/owner: [según inventario]. Versión/fecha: [valor real].
Estado documental inicial: EN_REVISIÓN.
Decisiones fuente: [IDs exactos con links al registro].
Dependencias de cierre: [documentos existentes revisados].
Referencias futuras: [entrada del índice y PENDIENTE_LOTE_N].
Consumidores: [IDs y responsabilidades].
Redacta con la plantilla de 11.01: objetivo, alcance, invariantes, contratos,
metodología futura paso a paso y justificaciones, ejemplos no ejecutados,
fallos/seguridad, verificación futura y AC entrada→resultado→evidencia.
No inventes parámetros aprobados; registra lo no definido como PROPUESTA
o POR_VERIFICAR con owner, impacto y prueba futura.
No repitas contratos canónicos de otro archivo. Si falta una dependencia
crítica o contradice la fuente, informa antes de cerrar la spec.
Escribe únicamente el archivo asignado; devuelve fuentes, cambios,
comprobaciones realmente realizadas y pendientes.
```

## P-03 — Revisor sin escritura

```text
Revisa [rutas] sin modificarlas. Usa fuente y dependencias indicadas.
Comprueba preservación del alcance, autoridad de Game, motor único,
propiedad de datos, seguridad, idempotencia y criterios negativos pertinentes.
Devuelve hallazgo con ubicación, fuente PRE/contrato, impacto y corrección
sugerida. Distingue bloqueo de coherencia y mejora editorial.
No apruebes requisitos ni afirmes pruebas de producto no observadas.
El autor conserva propiedad hasta que el coordinador integre la revisión.
```

## P-04 — Auditoría de lote

```text
Entrada: manifiesto aprobado del lote, maestro archivado, registro y matriz.
1. Enumera todos los archivos fuera de .git y compara con rutas autorizadas.
2. Compara IDs/textos de las 246 decisiones fuente con REG; exige unicidad.
3. Compara246 filas de trazabilidad y sus destinos con el inventario.
4. Valida archivos y anchors de links locales; rutas futuras deben apuntar
   al índice existente, nunca a stubs de otros lotes.
5. Revisa estructura, fences, metadata, dependencias acíclicas y semántica.
6. Consulta Git en modo solo lectura y diff --check cuando sea aplicable;
   complementa la consulta con lectura de archivos no rastreados.
7. Reporta resultados observados, método y límites. No ejecutar tests de
   producto ni instalar validadores. No guardar scripts ejecutables.
No confundir encaminamiento de decisiones con specs redactadas ni producto
probado. Si hay fallos, corregir dentro de rutas autorizadas y volver a
comprobar lo afectado antes de afirmar el cierre documental.
```

## P-05 — Cambio de requisito

```text
Presenta una propuesta, sin aplicarla:
ID del cambio; decisión fuente y formulación actual; motivo/evidencia;
alternativas; efecto sobre alcance, arquitectura, seguridad, coste y datos;
specs/ADR/matriz/planes consumidores; validación y reversibilidad;
texto de aprobación solicitado.
Conserva la fuente histórica. No habilites pagos, cuarto servicio,
otro motor o proveedor por una recomendación del agente.
Tras aprobación, registra resolución y actualiza solo las rutas autorizadas.
```

## P-06 — Entrega de lote

```text
Entrega: archivos creados/modificados y links; decisiones referenciadas;
dependencias; autor/revisor efectivos; conflictos y riesgos abiertos;
verificaciones de texto realmente ejecutadas y limitaciones; criterios
pendientes; declaración de que implementación no se ha iniciado.
Deja documentos EN_REVISIÓN y detente.
Para lote 1 solicita: APRUEBO LOTE 1.
No generar rutas del lote siguiente por el mero hecho de listarlas.
```

## Ejemplo de brief seguro

```text
Tarea: revisión documental de ADR-0004.
Owner de escritura: autorADR; revisor: revisorS (solo lectura).
Entrada mínima: AGENTS, PRE-05.C1, PRE-05.C2, PRE-05.C3, PRE-05.E1, PRE-05.E2, PRE-05.E3, PRE-05.E4, PRE-06.B2 y PRE-06.E1,
ADR0004 y entradas de índice05.03/05.04/06.02.
Salida: comentarios sobre commit-first, aviso perdido e idempotencia.
Prohibido: añadir outbox, Redis o tests ejecutables.
Aceptación: cada hallazgo cita fuente y no declara integración probada.
```

Es un ejemplo de contrato de tarea, no una orden de ejecutarla. La trazabilidad de trabajo real se conserva en el reporte del lote; los planes de harness quedan para lote 5.

