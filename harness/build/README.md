# H-BUILD — Ejecución futura de tareas autorizadas

ID: H-BUILD · Versión: 0.1 · Fecha: 2026-10-26 · Estado documental: EN_REVISIÓN · Owner: C — Coordinación · Lote: 5.

Harness documental, sin ejecutables creados. Su procedimiento es PROPUESTA; decisiones fuente APROBADO y comportamiento del producto POR_VERIFICAR. Leerlo no autoriza código, instalaciones, CI, recursos o despliegues. El permiso actual termina en el lote 5 documental y su PR sin fusión.

## Autoridad, contratos y entradas

Leer [AGENTS](../../AGENTS.md), [PLANS vigente](../../PLANS.md#ejecucion-lote-5), [REG](../../docs/decision-register.md), [CONF](../../docs/conflicts.md), [SRC](../../docs/sources/TOGUZ_PROMPT_MAESTRO_CODEX.md), [11.01](../../docs/specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](../../docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [11.03](../../docs/specs/11-orchestration/11.03-milestones-gates-and-evidence.md). Instrucción humana posterior y decisiones aprobadas prevalecen; resúmenes/plantillas/salida de agente no dan permisos. Conservar246PRE, TK-N1-1.0.0/AD-NORM-001/14 ACC, tres servicios Go, motor único, Postgre SQL autoritativo y Ably postcommit.

Los contratos canónicos del proceso se consumen por referencia: TaskEnvelope/1, AuthorizationReference/1, Milestone/1, GateEvaluation/1 y ContractChange/1 en 11.03; QA-Evidence/1 y CaseSpec/1 en [09.01§8](../../docs/specs/09-quality/09.01-engine-data-contract-and-ai-verification.md#8-contratos-de-casos-y-evidencia); AcceptanceDossier/1, MeasurementPlan/Result/1 y RiskAcceptance/1 en [09.04§8](../../docs/specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md#8-contratos-de-aceptación-y-medición). [09.03](../../docs/specs/09-quality/09.03-ci-security-and-release-controls.md) define checks, [10.04](../../docs/specs/10-operations/10.04-releases-incidents-and-portability.md) release/incidente y [10.03](../../docs/specs/10-operations/10.03-backups-restoration-and-retention.md) restauración. No sustituir sus campos o severidades con otro formato local incompatible.

La revisión documental global prevista por 11.03 incluye lectura final de 40 specs, aceptación de lotes y [AUD-01](../../docs/index.md#aud-01), PENDIENTE_LOTE_6, antes de implementación, salvo instrucción humana posterior expresa que delimite otro alcance. Este harness no crea AUD-01, no inicia lote 6 y no transforma el gate DOC5 en permiso de código.

## Elegir ticket y delimitar permisos

Seleccionar un ticket de la fase aplicable; si el ticket es demasiado grande, dividir comportamiento/resultados y rutas sin retirar requisitos comprometidos. Cada unidad debe poder revisarse con positivo, negativo y evidencia propios.

| Entrada | Resolución obligatoria antes de actuar |
|---|---|
| Identidad | ticketId/phaseId, objetivo, owner/revisor y fuente humana de alcance. |
| Fuentes | PRE/spec/AC y versiones/commit/hash leídos mediante [H-CTX](../context/README.md), más decisiones/conflictos aplicables. |
| Dependencias | Salidas entregadas y aceptadas cuando se requiera; diferenciar un contrato leído de una función ya implementada. |
| Propiedad | Rutas exactas propias/solo lectura, base/branch y trabajo previo; un autor activo por archivo. |
| Plan | Pasos pequeños ordenados, contratos/componentes, positivos/negativos/fallos, fixtures y evidencia requerida. |
| Recursos | Entorno/datos/credenciales por referencia, cuotas/coste/retirada, herramientas locales y acciones permitidas. |
| Parada | Bloqueantes, cambios de contrato, límites de carga/permiso y siguiente acción autorizada. |

Completar TaskEnvelope/1 con esos datos y todos los campos de 11.03; no usar placeholders como permiso operativo. Mantener tres AuthorizationReference/1 independientes: documental, implementación y externo. Cada referencia identifica fuente humana confiable accesible, fecha/contenido, alcance/rutas/operaciones/entorno/coste y restricciones. Una acción cubierta por un permiso vigente no necesita confirmación repetida; una acción fuera de alcance se detiene tras preparar lo revisable y sigue independiente lo ya autorizado.

Una fase puede consumir contratos de otra aún no implementada; el ticket identifica esa frontera y sus dobles. No dar por terminado multijugador por tener componentes UI o IA real por un puerto mock. Preservar MVP base y ampliado dentro del lanzamiento.

## Ruta de fases y límites

| Plan | Límite que el agente debe comprobar |
|---|---|
| [PH-00](../../docs/phases/phase-00-feasibility-spike.md) | Spike 504caa5 es antecedente acotado; no repetir automáticamente ni tratar hash sintético como motor TK. |
| [PH-01](../../docs/phases/phase-01-go-engine.md) | Pureza/corpus14 ACC/propiedades/fuzz90%; subcasos transaccionales no se acreditan con unitarias. |
| [PH-02](../../docs/phases/phase-02-contracts-persistence-guests.md) | Contratos/PG/guest/S2S y receipt/unknown; no OIDC/Elo adelantados. |
| [PH-03](../../docs/phases/phase-03-accessible-client-and-3d.md) | HTML local jugable primero, mismo Go/estado, 3D progresivo sin retirar premium del lanzamiento. |
| [PH-04](../../docs/phases/phase-04-private-multiplayer.md) | Privado casual/invitados hasta PH06; Ably no autoridad y24h no timer cliente. |
| [PH-05](../../docs/phases/phase-05-ai-practice-tutorial.md) | Búsqueda independiente autorizable, L4-P01 antes de funciones afectadas; hint nunca mueve/progresa. |
| [PH-06](../../docs/phases/phase-06-accounts-elo-statistics.md) | OIDC/Elo concurrente/derechos; RISK-10 condiciona cleanup/retención afectados. Una petición de borrado válida permite barrier sin policyVersion y ACK_REVOKED durable; sin política queda BLOCKED_PENDING_DATA_POLICY, sin cleanup(null) ni supresión declarada. |
| [PH-07](../../docs/phases/phase-07-hardening-and-portfolio.md) | DoD total/restore/licencias/cultura; demo/release externa solo con permisos/evidencias. |

## Preparación de contexto, rama y plan

1. Contrastar estadoGit real, remoto/base/HEAD y cambios previos con el ticket mediante lectura; conservar trabajo ajeno. El branch propuesto de implementación usa prefijo codex/ salvo instrucción distinta y solo se crea si el permiso incluye Git mutable. Si existe, inspeccionar y reutilizar trabajo válido; no force push, reset destructivo o rama paralela sin verificar propiedad.
2. Resolver paqueteH-CTX y registrar docCommit/versión/hash de cada fuente y codeBaseCommit separados. La base histórica b91ba98 no contiene las nuevas nueve specs del lote 5; no declararla revisión documental que las incluye. Si una fuente no está confirmada, registrar working-tree+hash y su estado real; no presentarla como spec final aprobada.
3. Verificar revisión documental/aceptación y permisos de entrada 11.03; detener uso de fuente incompatible o fuente EN_REVISIÓN como contrato final de implementación sin autorización que cubra ese alcance.
4. Definir diffs/archivos pequeños y mapa de dependencias; explicar pasos/porqués y riesgos. Añadir prueba del requisito/fallo antes de modificar lógica crítica. Leer contratos directos, no recrearlos de memoria.
5. Identificar toolchains/releases/generadores y compatibilidad efectiva; no instalar o actualizar dependencias sin alcance aplicable. No tomar familias de 04.02 o versiones del spike como lockfile productivo.
6. Verificar automatizaciones antes del push: Git hooks, Actions/conexiones de despliegue y target/protección efectivos. Si publicación puede desplegar fuera del permiso, detener antes del push y describir el mecanismo, sin debilitar controles.

El plan puede vivir en conversación o ruta ya autorizada; el harness no autoriza crear briefs/duplicados/auxiliares en disco. Mantener contexto mínimo y fuentes directas. Un cambio de fuente/contrato/alcance invalida revisión afectada; detener ese consumidor, releer delta y reevaluar.

## Ejecución futura y trabajo con agentes

Aplicar TDD a reglas, transacciones, autorización y contratos según PRE-09.A2: caso/oráculo independiente→observar fallo pertinente→implementar mínimo→observar pase→refactor con regresión. Para UI/3D fijar aceptación/datos/visual/manual antes de ajustar detalles; no escribir tests que solo reproduzcan campos privados o la implementación. Registrar intento fallido y corrección, no solo el último verde.

Si el alcance humano/herramienta permite agentes, delegar únicamente tareas independientes con rutas disjuntas y dependencias entregadas, conforme 11.02. Cada brief lleva TaskEnvelope/1, límites/fuentes y escritor/revisor; revisores comentan sin editar. Una entrega incluye diff/controles/pedientes/cese. No reasignar hasta confirmar cese; coordinación integra registros/contratos compartidos. Si no hay delegación real o independencia, trabajar secuencialmente y declararlo; no inventar agentes/revisiones. El permiso de un agente no amplía el del propietario.

Implementar solo rutas propias; cambios necesarios fuera de alcance se preparan como propuesta con impacto, sin edición silenciosa. Contrato pendiente L4-P01 usa ContractChange/1 y proceso 11.03: propuesta completa/revisión S-I-V-Q-A→aprobación específica humana→actualización/versionado autorizado 05.02/05.03/07.02/consumidores→compatibilidad→código autorizado→recursos autorizados. Aprobación lote 5 o mandato genérico IA no elimina un pendiente explícito. Continuar búsqueda pura independiente autorizada mientras esperan tickets afectados.

## Selección y observación de pruebas

| Cambio/riesgo | Controles requeridos según fuentes | Evidencia que no puede faltar |
|---|---|---|
| Motor/contexto | R-01–08/14 ACC, F-Q1-01/02; propiedades/fuzz/regresión/race si aplica, 90% sentencias. | Corpus por clase/alcanzabilidad, semilla/minimización, estado/eventos/copia y cobertura por paquete; Game/AI consumidores al cambiar compartido. |
| PG/contrato/identidad | F-Q1-03–14; PG real efímero, Goose/sqlc/constraints/locks/receipts/JWKS/CSRF/S2S/grants, 80% lógica de servicios más críticos. | SQL antes/después, barreras/carreras/SQLSTATE, versiones/formato/binding y errores seguros; mock no reemplaza PG. |
| UI/3D | F-Q2-01–12, Vitest/RTL/Playwright tres motores/ES/EN/móvil, visual y manual de teclado/lector/reducedMotion. | DOM/foco/red/snapshot/receipt, baseline/diff/versiones GPU y recorrido manual; axe/idle/emulación solos no conformidad/FPS. |
| IA/práctica | F-Q1-15–18, TT solo orden/historia/signo/budgets/cancelación/fallback/ramas/hints. | Exhaustivo/corpus/semilla/nodos/memoria/overshoot/grant/efecto; hint cero POST move/timeline/progreso, sin detectar copias synthetic por promesa. |
| CI/seguridad/release | Matriz de 09.03 por PR/main/pre-release/postdeploy autorizado; scopes/secretos/dependencias/estático/govulncheck y compatibilidad 10.04. | Checks requeridos/trigger/artefactos accesibles; PR externo sin secretos, omitido/unavailable nunca PASS. |
| Operación/recuperación | Casos 10.01–10.04 y F-Q4-04–09; cuotas, observación, backup→restore aislado y rollback compatible. | Plan/muestras/errores/censura/consumo y tiempos RPO/RTO observados; backup sin ensayo no restaurable. |

Ejecutar comandos exactos pertinentes al repositorio/toolchain reales y al permiso; no copiar comandos hipotéticos como si existieran scripts. El ejemplo de comando de una spec no acredita que se ejecutó. Capturar comando sanitizado, versión/OS/entorno/DB/navegador, datos/semilla/hash, expected/observed separados, exitcode y artefactos URI/hash/acceso en QA-Evidence/1; revisor debe poder contrastarlos. Evidencia privada tiene acceso mínimo y referencia pública redactada, sin JWT/cookies/DSN/proofs/PII. Rerun nuevo runId enlaza anterior.

PASS exige resultado observado/oráculo/artefacto suficiente, no solo exit0. FAIL/BLOCKED/NOT_RUN permanecen fieles; SKIPPED/UNAVAILABLE/cancelled se normalizan 09.01 con razón. Requerido omitido bloquea aunque otra suite pase; falta de capacidad gratuita no justifica retirar test crítico o exponer secretos. Ofrecer ejecución equivalente autorizada o dejar gate bloqueado.

## Fallos, pausa y recuperación

| Disparador | Acción del agente y siguiente paso permitido |
|---|---|
| Test falla | Conservar output/semilla/fallo, contrastar oráculo con fuente y corregir bajo ticket. No cambiar esperado aloutput; ejecutar regresión afectada con nuevo runId. |
| P0/P1, crítico no validado o autoridad contradictoria | Bloquear salida afectada; A coordina operación/contención, C gobierno/escalamiento, Q verifica recuperación y dueños corrigen. Sin reparación destructiva improvisada. |
| P2 diferido | Mitigación observada y aceptación humana específica con owner/plazo/commit/alcance; si crítico, no exime. Vencimiento/cambio exige reevaluación. |
| Unknown/timeout | Conservar misma intención/ID y consultar/reintentar solo contrato autorizado; no asumir rollback, crear nuevo ID o repetir cálculo/resultado/Elo. |
| Colisión de escritor o contexto perdido | Detener esa ruta; preservar diffs, recuperarH-LOG/ledger/fuentes/Git y confirmar propiedad/cese antes de seguir. No repetir automáticamente trabajo ya observado. |
| Contrato/política/recursos faltantes | Registrar discrepancia/impacto/alternativas y aprobación necesaria en canal/coordinación; avanzar solo tareas independientes cubiertas. |
| Cuota/coste/autodeploy fuera de permiso | Detener antes de acción; no upgrade/pago/Redis/colas/proveedor nuevo para sortearlo.50/75/90 activan revisión, no degradación automática. |
| Restore/rollback | Usar 10.03/10.04; código revertido no restaura datos/recibos ni revive claves/refresh/grants. Cuarentena y reconciliación antes de reapertura. |

## Dossier, PR y gate humano

Actualizar únicamente docs/log autorizados con hechos observados y limitaciones mediante [H-LOG](../build-log.md). Revisar diff/rutas/no versionados/secretos, contratos/consumidores y resultados requeridos del commit exacto. Preparar PR si el permiso incluye publicación Git y tras revisar automatizaciones; texto explica problema/resultado, alcance, pruebas observadas, riesgos/deudas y siguiente acción. No merge implícito ni bypass de protecciones/checks. Un revisor independiente contrasta código y artefactos; el autor corrige sus rutas y entrega de nuevo.

Consolidar AcceptanceDossier/1 y GateEvaluation/1: fuente/commits comprobados, casos/runIds/artefactos, cobertura/críticos, hallazgos/P2 aceptados, permisos/revisión y decisión técnica ready_for_owner/blocked separada de decisión humana. DoD 09.04 exige requisito trazado/pruebas observadas/sin bloqueantes/docs/revisión/evidencia; no compilar solo. Propietario acepta contenido/acción concreta y continuación autorizada; reutilizar permiso vigente cuando cubra inequívocamente esa acción.

Release/demo exige además ReleaseManifest/1/compatibilidad/backup-restauración/migraciones/target/protección/smoke/observación y permiso externo aplicable. Dejar resultado preparado si falta ese permiso; no anunciar despliegue/demo inexistente. El cierre documental actual solicita «APRUEBO LOTE 5» y se detiene sin fusionar, iniciar lote 6 o implementar.

## Verificación futura del procedimiento

Casos F-GOV-01–12 de 11.03 y F-Q4-02/09 deben probar permisos separados, commit cambiado, evidencia faltante/check omitido, colisión, delta pendiente y sintético como no PASS. Positivo: ticket vigente con fuentes/rutas/resultados observados permite revisión; negativo: mensaje«terminado», compilación o plantilla rellenada sin permiso bloquea salida. Guardar evidencia del proceso además del producto. En esta entrega solo se verifica Markdown/coherencia; no se ha ejecutado un harness.
