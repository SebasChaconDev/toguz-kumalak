# H-LOG — Registro factual para trabajo futuro

> **Aprobación posterior — 2026-10-26: APROBADO documentalmente.** El propietario declaró «APRUEBO LOTE 5 como documentación técnica.» sobre HEAD 9ff1a2857d0b91b9cf7925a4c36de30e377cef3d; PR #5 integrado por squash 6579ca930a9081f773b390ce61f6bc80cc6b1ae7. El [recibo de aprobación](../PLANS.md#cierre-lote-5) prevalece sobre los estados de revisión históricos. L4-P01, la política de retención y demás propuestas mantienen su aprobación específica pendiente; sin permiso de implementación, workflows, proveedores o lote6.

ID: H-LOG · Versión: 0.2 · Fecha: 2026-10-26 · Estado documental: APROBADO · Owner: C — Coordinación · Lote: 5.

Estructura documental PROPUESTA; no contiene entradas de implementación real. No se ha ejecutado código, test, CI, restore o despliegue por crear este archivo. El ejemplo final es SYNTHETIC/NOT_RUN y no representa un trabajo efectuado, un permiso o una aceptación.

## Autoridad y relación con otros registros

El ledger de ejecución documental vigente está en [PLANS](../PLANS.md#ejecucion-lote-5); no se duplica aquí como si fueran builds de producto. [AGENTS](../AGENTS.md), [11.01](../docs/specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](../docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [11.03](../docs/specs/11-orchestration/11.03-milestones-gates-and-evidence.md) gobiernan futuras entradas. [H-BUILD](build/README.md) describe ejecución autorizada y [H-CTX](context/README.md) procedencia del contexto.

Consumir TaskEnvelope/1, AuthorizationReference/1 y GateEvaluation/1 de 11.03; QA-Evidence/1 de [09.01§8](../docs/specs/09-quality/09.01-engine-data-contract-and-ai-verification.md#8-contratos-de-casos-y-evidencia), AcceptanceDossier/1 de [09.04§8](../docs/specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md#8-contratos-de-aceptación-y-medición) e IncidentRecord/RecoveryDecision/ReleaseManifest de [10.04](../docs/specs/10-operations/10.04-releases-incidents-and-portability.md). Este log enlaza esos registros; no redefine sus campos/resultados ni reemplaza artefactos de prueba.

## Reglas de escritura futura

Añadir una entrada solo bajo autorización que incluya actualizar esta ruta y con un único escritor activo. Mantener cronología UTC, autor/agente real y commit/contexto exactos; corregir una entrada con nota trazable, sin borrar intento fallido ni reescribir retrospectivamente una aprobación. Si varios agentes ejecutan tareas, coordinación reúne entregas y escribe el log; los autores no coeditan la ruta.

Separar observado/esperado/previsto y distinguir tres estados: documental BORRADOR/EN_REVISIÓN/APROBADO; naturaleza APROBADO/PROPUESTA/POR_VERIFICAR; ejecución PASS/FAIL/BLOCKED/NOT_RUN según 09.01. Estados propuestos de tarea (preparada/en curso/entregada/bloqueada/aceptada) son organización y no aprobación de producto. Una entrada «entregada» no cierra hito; la decisión humana y GateEvaluation se registran por referencia verificable.

No guardar JWT/cookies/proofs/claves/DSN/invitaciones/PII ni outputs privados completos. Comandos sanitizados deben conservar suficiente estructura para reproducir, con valores sensibles sustituidos por referencias restringidas; señalar redacción y acceso del revisor. No publicar exportaciones o datos reales como fixtures. Retención de log/artefactos requiere política aplicable; no inventar plazo jurídico.

## Estructura de cada entrada

Los campos siguientes son vista factual de los contratos citados; resolverlos con datos reales o marcar explícitamente no ejecutado/no aplica/desconocido y razón. Desconocido no es cero. Si falta evidencia requerida, el gate queda bloqueado, sin rellenar un PASS.

| Campo | Contenido requerido y validación |
|---|---|
| Identidad/tiempo | entryId único, fecha UTC inicio/fin, autor/agente real, owner/revisor y clase REAL o SYNTHETIC. No una fecha futura presentada como observada. |
| Fase/tarea | phaseId/ticketId, objetivo/resultado, TaskEnvelope/1 y dependencias/salidas consumidas. |
| Git/contexto | Repo/base/rama/HEAD y codeCommit(s)/builds; docCommit/versión/hash por fuente leída, working-tree/diff cuando exista; no SHA fabricado. |
| Especificaciones | PRE/spec/AC/caseIds, contratos/motor/migración/versiones y conflictos/cambios aplicados con aprobación identificada. |
| Autorizaciones | AuthorizationReference/1 documental/código/externo, fuente humana/fecha/alcance/rutas/entorno/operaciones/coste/restricciones. Una referencia null no habilita acciones. |
| Archivos | Rutas propias modificadas/nuevas/eliminadas realmente, diff/alcance y trabajo ajeno preservado; rutas solo lectura separadas. |
| Comandos | Comando exacto sanitizado, directorio/contexto, hora, herramienta/versión/entorno, salida/exitcode observados y razón si no se ejecutó. Separar comandos previstos. |
| Pruebas/resultados | QA-Evidence/1 runId nuevo por intento, expected/observed separados, status por caso, datos/semilla/procedencia/aislamiento y límites del doble/fallo. |
| Evidencia | Artefactos URI/tipo/SHA256/acceso/redacción, cobertura/muestra/duración/unidad, outputs y revisión independiente. Hash no sustituye observación/oráculo. |
| Fallos/correcciones | Hallazgo/severidad/fuente/impacto, fallo/contraejemplo/attempt previo, cambio aplicado y runId posterior enlazado. No borrar el primer fallo. |
| Riesgos/operación | Pendientes/bloqueantes/P2 con owner-plazo-aceptación, cuotas/consumo previo/posterior y política, incidentes/recovery/compatibilidad; no RPO/RTO sin restore. |
| Estado/decisión | Estado real de tarea, resultados/gate técnico, AcceptanceDossier/1/GateEvaluation/1 y decisión humana separada con contenido/commit/alcance. |
| Continuidad | Siguiente acción permitida y requisitos aún pendientes, dependencias bloqueadas, cese de escritura/transferencia/verificación de contexto. |

Para un test omitido por plataforma, normalizar SKIPPED/UNAVAILABLE/cancelled a BLOCKED o NOT_RUN con razón de 09.01 y conservar origen; requerido bloquea salida. Exit0 sin artefacto/oráculo o «agente terminó» no es prueba. Cada rerun enlaza el run fallido anterior, sin reutilizar runId o mezclar commits como si fueran un conjunto.

## Plantilla de entrada vacía conceptual

La plantilla permanece en Markdown; no es una entrada REAL, un archivo nuevo a crear ni un registro ya ejecutado:

```text
entryId: <único>
recordClass: <REAL solo si observado; SYNTHETIC si ejemplo>
startedAtUTC / finishedAtUTC / author / reviewer: <hechos o pendiente explicado>
phaseId / ticketId / taskEnvelopeRef: <resueltos>
repo / base / branch / HEAD / codeCommitsBuilds: <observados>
docSources: <ruta, versión, estado, docCommit o working-tree, sha256 leído>
specsPREAC / contractsMotorMigrations / dependencies: <referencias exactas>
authorizationRefs: <documental, código, externo con alcance verificable>
filesActuallyChanged / readOnlyFiles / preservedWork: <rutas reales>
commandsActuallyExecuted: <comando sanitizado, contexto, hora, versión, exit, outputURI>
commandsPlannedButNotRun: <lista separada y razón>
testAttempts: <runId, caseIds, expected, observed, status, artifactsURI/hash/access>
findings / corrections / previousRunRefs / newRunRefs: <hechos>
risks / acceptedP2 / resourcesConsumption / incidentRecoveryRefs: <hechos y pendientes>
dossierRef / technicalGate / humanDecisionRef: <separados, alcance exacto>
taskState / permittedNextAction / stopReason / writingCeased: <estado real>
```

## Ejemplo inequívocamente sintético — SYNTHETIC / NOT_RUN

Ejemplo de tarea hipotética bloqueada antes de código. Todo el bloque es ilustrativo: no hubo agente implementador, rama creada, comando ejecutado o prueba observada. Los símbolos entre ángulos son placeholders, no SHA/permiso válido. No usarlo para aceptación o cobertura.

```text
entryId: SYNTHETIC-HLOG-EXAMPLE-01
recordClass: SYNTHETIC
startedAtUTC: <fecha futura no establecida>
finishedAtUTC: <no aplica, ejemplo sin ejecución>
author: <agente hipotético>
reviewer: <revisor hipotético, revisión no realizada>
phaseId: PH-05
ticketId: PH-05-T06
taskEnvelopeRef: <no existe una asignación operativa para este ejemplo>
repo / base / branch / HEAD: <no inspeccionados por este ejemplo>
codeCommitsBuilds: []
docSources:
  - 11.03, versión 0.1 EN_REVISIÓN: contrato de cambio y permisos
  - 08.02/08.03/08.04 y 05.02/05.03/07.02: fuentes a revisar en tarea real
  - docCommit: <SHA futuro que realmente contenga fuentes requeridas>
  - sha256: <sin calcular en este ejemplo>
  - b91ba98: base histórica, no SHA que contiene nuevas specs lote 5
specsPREAC: <resolver fuentes/AC/caseIds exactos antes de una ejecución real>
contractsMotorMigrations: L4-P01 PROPUESTA, delta aún no aprobado en este ejemplo
dependencies: aprobación/versionado de contratos antes de PH-05-T06
authorizationRefs:
  documental: <ningún permiso adicional emitido por el ejemplo>
  código: null
  externo: null
  contractChangeApproval: null
filesActuallyChanged: []
readOnlyFilesActuallyReadByExample: []
commandsActuallyExecuted: []
commandsPlannedButNotRun:
  - <definir comando real tras contrato/permisos; no se ha ejecutado>
testAttempts:
  - runId: <ninguno, no hubo intento real>
    caseIds: [PH05-C05, F-GOV-06]
    expected: sin aprobación específica, detener código de rama afectado
    observed: NO_OBSERVADO; es un ejemplo, no un resultado empírico
    status: NOT_RUN
    reason: SYNTHETIC, sin ejecución/permisos/artefactos
    artifacts: []
findings: <hipótesis de falta de aprobación, no incidente observado>
corrections / previousRunRefs / newRunRefs: []
risks: no inferir delta por aprobación documental general
resourcesConsumption: NO_MEDIDO; ningún recurso creado por el ejemplo
dossierRef: <no existe dossier de implementación>
technicalGate: <evaluación hipotética blocked / CONTRACT_APPROVAL_MISSING>
humanDecisionRef: null
taskState: ejemplo preparado, ninguna tarea real iniciada
permittedNextAction: en una tarea real, preparar ContractChange y verificar permiso
stopReason: no implementar consecuencias de L4-P01 sin aprobación y permiso de código
writingCeased: <no aplica al ejemplo, no hay autor de producto activo>
```

El ejemplo tampoco bloquea una búsqueda pura independiente si una autorización real posterior la cubre; su TaskEnvelope y log tendrían alcance/casos propios. L4-P01 no se aprueba por cambiar null a texto sin fuente humana verificable.

## Auditoría y continuidad

Revisor coteja entrada con Git/diff/rutas, comandos/salidas/artefactos y fuentes directas del contexto. Positivo: ejecución real con permisos/versiones/resultados trazables permite revisar el siguiente alcance; negativo: mezclar ejemplo con REAL, PASS sin observación, commit documental que no contiene fuente nueva, comando previsto como ejecutado o aceptación sin fuente bloquean gate. Casos F-GOV-03/04/08/09/11/12 de 11.03 cubren esos fallos.

Al recuperar sesión, partir de PLANS/H-LOG y estado Git real, resolver fuentes/permisos/propiedad y acciones pendientes; no repetir spike, proveedor, restore o test externo automáticamente. Si se pierde recibo/efectos inciertos, usar contrato de producto y 10.03/10.04, no este log como autoridad deportiva. A coordina operación, C gobierno y Q verifica recuperación. El cierre de lote 5 permanece documental, sin implementación/recursos/demo/lote 6/fusión por este archivo.
