# PH-00 — Factibilidad y entrada a implementación

ID: PH-00 · Versión: 0.1 · Fecha: 2026-10-26 · Estado documental: EN_REVISIÓN · Owner: C — Coordinación · Lote: 5.

Plan operativo futuro; decisiones fuente APROBADO, organización de tareas PROPUESTA y comportamiento de producto POR_VERIFICAR. No es una orden de repetir el spike ni permiso para implementar o crear recursos.

## Ficha Milestone/1

| Campo requerido | Contenido del hito |
|---|---|
| Identificador | PH-00 |
| Objetivo | Establecer qué demuestra la factibilidad experimental existente y qué falta antes de iniciar producto. |
| Dependencias | 01.01–01.03; contratos 05.01–05.04 y 07.02; 09.01/09.03/09.04; 10.01–10.04; 11.01–11.03; evidencia experimental identificada abajo. |
| Responsable | C reúne el expediente; A interpreta plataforma/recursos, S seguridad/datos y Q suficiencia de evidencia; propietario decide el alcance posterior. |
| Entrada | Fuentes y autorización vigentes verificadas; lectura del antecedente sin ejecutar su código, revisión documental global prevista por 11.03 y del alcance del siguiente ticket. |
| Entregables | Matriz hipótesis→evidencia→límite→verificación pendiente, inventario de cambios materiales y TaskEnvelope/1 de la siguiente tarea autorizable. |
| Pruebas positivas/negativas | Matriz PH00-C01–C05 y casos QA/operación pertinentes; resultados de producto todavía pendientes. |
| Riesgos | RISK-02/03/04/06/07/09/14; contexto experimental obsoleto, falso permiso externo y atribución de seguridad productiva a una sonda. |
| Evidencia obligatoria | Fuentes históricas por commit y artefacto, condiciones/versiones, retirada y desviaciones; QA-Evidence/1 nuevo únicamente si se autoriza un ensayo nuevo. |
| Salida | Diferencias y límites revisados, ningún antecedente presentado como producto y GateEvaluation/1 con siguiente acción delimitada. |
| Aprobación humana | Revisión/aceptación del dossier; permiso posterior de código y de recursos cuando la acción los requiera, independientes de la aceptación documental. |
| Corrección/recuperación | Corregir interpretación o referencia; si aparece cambio material, proponer ensayo acotado con autorización aplicable y retirada antes de ejecutarlo. No resucitar recursos retirados. |

## Fuentes y contratos consumidos

Leer [AGENTS](../../AGENTS.md), [PLANS vigente](../../PLANS.md#ejecucion-lote-5), [REG](../decision-register.md), [CONF](../conflicts.md) y [SRC](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md). PRE-01.C1/C2/C4, PRE-05.F4, PRE-10.A1–A4 y PRE-11.E2 gobiernan factibilidad; sus textos y dueños no cambian.

| Fuente | Uso operativo |
|---|---|
| [01.01](../specs/01-infrastructure/01.01-stack-monorepo.md), [01.02](../specs/01-infrastructure/01.02-platform-feasibility.md), [01.03](../specs/01-infrastructure/01.03-cost-capacity-evolution.md) | Tres módulos Go autónomos, motor compartido Game/AI, cuatro proyectos previstos y presupuesto por ámbito. |
| [05.01](../specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md)–[05.04](../specs/05-backend/05.04-event-schemas-and-postcommit-publication.md), [07.02](../specs/07-identity-data/07.02-tokens-authorization-and-service-identity.md) | Comparar experimento con contratos de producto, transacción, autorización y publicación. |
| [09.01](../specs/09-quality/09.01-engine-data-contract-and-ai-verification.md), [09.03](../specs/09-quality/09.03-ci-security-and-release-controls.md), [09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md) | QA-Evidence/1, checks por consumidor, AcceptanceDossier/1 y DoD. |
| [10.01](../specs/10-operations/10.01-deployments-origins-regions-and-secrets.md)–[10.04](../specs/10-operations/10.04-releases-incidents-and-portability.md) | Plataforma, observación/cuotas, restauración y releases; no prestaciones gratuitas presumidas. |
| [11.03](../specs/11-orchestration/11.03-milestones-gates-and-evidence.md) | Milestone/1, TaskEnvelope/1, AuthorizationReference/1 y GateEvaluation/1 sin redefinirlos. |

Las specs nuevas del lote 5 están EN_REVISIÓN y sus protocolos son PROPUESTA. Su existencia no satisface la revisión documental global de las 40 specs, aceptación de lotes y [AUD-01](../index.md#aud-01), PENDIENTE_LOTE_6, prevista antes de implementar. Una excepción requiere instrucción posterior expresa; este plan no inicia lote 6.

## Antecedente observado y límites

Procedencia fija: rama experimental spike/infra-feasibility-01, commit 504caa5d8d4f538e131e1beffc8843cb8938772a. Consultar [cierre remoto](https://github.com/SebasChaconDev/toguz-kumalak/blob/504caa5d8d4f538e131e1beffc8843cb8938772a/spikes/infra-feasibility/REMOTE_AUDIT.md#remote-close-20261025) y [gate local](https://github.com/SebasChaconDev/toguz-kumalak/blob/504caa5d8d4f538e131e1beffc8843cb8938772a/spikes/infra-feasibility/evidence/integral-local-gate.json). Son evidencias históricas del experimento, sin trasladar sus permisos.

| Antecedente | Alcance observado | Pendiente de producto |
|---|---|---|
| GO_LOCAL | Gate de las sondas locales conservado en el artefacto citado. | Motor TK-N1, contratos finales, SDK/generadores y pruebas PRE-09 de producto. |
| GO_REMOTE | Seis builds READY sin caché, ocho intentos acumulados incluidos dos cancelados; raíces independientes y fuente compartida Game/AI. Go remoto1.26.0/local1.26.2, CLI65.1.0. | Volver a comprobar compatibilidad efectiva si cambian runtime, releases, contexto de build o dependencias; no inferir soporte futuro. |
| Motor de sonda | spike-engine/v0.1.0 y hash de fuentes equivalente. | Ese hash no acredita implementación de reglas TK-N1-1.0.0, Negamax ni14 ACC. |
| Neon/Ably y confianza | Postgre SQL 18 real con dos bases/roles, commit/receipt/reintento/carrera; aviso postcommit y REST; protección/OIDC de despliegue y cookie Chrome por rewrite. | JWT/JWKS/CSRF/sesiones de producto completos, seguridad integral, privacidad, cuotas a escala y restore. Sonda HMAC no equivale al protocolo final. |
| Fallos/medición | Fallo Ably conserva commit; descarte de respuesta por harness; rechazo OIDC; muestras C6 únicas del operador. | No corte TCP real instrumentado, caída física S2S, reactivación Neon ni colas PgBouncer; no p95/SLO. |
| Consumo/desviación | 117 operaciones SQL conservadoras frente al objetivo técnico80, sin traza wire exacta; coste adicional contratado0USD. | No afirmar cumplimiento de80, consumo nulo por API retrasada ni coste productivo garantizado. |
| Retirada | Proyectos/deployments propios retirados y claves propias revocadas; recursos previos preservados según ledger. | Ningún permiso para volver a crearlos o usar recursos anteriores. |

## Tickets pequeños y orden

Componentes futuros candidatos: manifiesto de contexto por ticket, matriz de compatibilidad y adaptadores mínimos dentro de las rutas propuestas de 01.01. Aquí no se crean directorios de producto.

| Ticket / dependencia | Pasos y razón | Resultado revisable y casos |
|---|---|---|
| PH-00-T01 / fuentes | Identificar commit documental futuro, versiones y permisos; leer evidencia histórica y sus correcciones. | Inventario de fuentes con límites; C01. |
| PH-00-T02 / T01 | Contrastar hipótesis de 01.02 con artefactos504caa5; separar sonda de producto y resultados ausentes. | Matriz de reutilización, sin ejecutar nada; C02/C03. |
| PH-00-T03 / T02 | Comparar toolchains, runtime, módulos, contratos y capacidad necesaria del ticket candidato. | Lista de cambios materiales y verificaciones mínimas; no autoactualizar ni instalar; C04. |
| PH-00-T04 / T03 | Preparar TaskEnvelope/1 para PH-01 y, solo si se justifica, propuesta de ensayo adicional independiente. Fijar objetivo, límite de carga, datos sintéticos, retirada y parada. | Propuesta concreta para decisión humana; C05. |
| PH-00-T05 / T04 | Q/revisor contrastan fuentes, C integra dossier y propietario evalúa siguiente alcance. | GateEvaluation/1; no recursos ni código por cambiar etiqueta. |

T01–T05 son tareas de preparación; cualquier prueba nueva es una tarea separada con permiso específico. No ejecutar el antiguo spike como comprobación rutinaria de este plan.

## Pruebas y evidencias futuras

| Caso | Positivo | Negativo | Evidencia / fuente de aceptación |
|---|---|---|---|
| PH00-C01 | Revisión fija recupera artefactos y condiciones correctas. | Referencia a otra revisión/resumen sin fuente no pasa revisión. | URI/hash/acceso y lectura; F-GOV-03/12 de 11.03. |
| PH00-C02 | Builds de sonda equivalentes acreditan esa hipótesis. | Hash sintético presentado como motor TK/14 ACC se rechaza. | Manifiestos, scope observado y matriz de pendientes; 01.02/09.01. |
| PH00-C03 | Desviaciones SQL y fallos históricos quedan visibles. | Muestra única, cuota desconocida o descarte del harness no se convierte en p95/cero/corte TCP. | Ledger/C6 y límites; 09.04 AC-04/07. |
| PH00-C04 | Cambio material produce plan de compatibilidad acotado. | Runtime cambiado no hereda un PASS histórico; tampoco fuerza repetición íntegra sin motivo. | Diff fuente/config propuesta y plan; F-DEP-01/02. |
| PH00-C05 | Acción nueva solo procede con permiso que cubra código/externo según corresponda. | Gate documental o GO_REMOTE no crea recursos/paga/despliega. | AuthorizationReference/1 y ausencia de acción no cubierta; F-GOV-02/10. |

Si se autoriza producto, sus casos usarán QA-Evidence/1 y AcceptanceDossier/1; este antecedente no se reescribe como una ejecución de esos formatos. Identificar documento/versión/hash y limitación si faltan campos históricos.

## Seguridad, recursos y riesgos

Lectura documental no renueva sesiones ni consulta cuentas. Un ensayo futuro distingue local/CI/integración/Preview/producción, datos y credenciales separados, límite por proveedor y permisos humanos; 50/75/90 activan revisión, sin pago automático. Verificar automatizaciones antes de push que pudiera desplegar. No registrar secretos ni copiar evidencias privadas al repositorio público.

## DoD, gate y corrección

DoD de este hito de preparación: matriz de hipótesis completa, procedencia verificable, límites/desviaciones explícitos, fuentes actuales revisadas, revisión independiente y siguiente TaskEnvelope/1 delimitado. No afirma producto construido ni solicita repetir pruebas completadas.

C prepara GateEvaluation/1 con AcceptanceDossier/1; Q verifica suficiencia y el propietario acepta el contenido/acción concreta. P0/P1, crítico no validado o autoridad contradictoria bloquean la salida afectada. P2 diferible requiere aceptación humana con owner/plazo; no permite omitir un crítico. Corregir referencias o ampliar ensayo propuesto sin cambiar decisiones fuente. En incidente futuro A coordina operación, C gobierno y Q recuperación según 09.04/10.04. La salida posible es preparar [PH-01](phase-01-go-engine.md); solo un permiso posterior aplicable inicia código.
