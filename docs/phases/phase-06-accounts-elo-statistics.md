# PH-06 — Cuentas, historial, Elo, estadísticas y derechos de datos

> **Aprobación posterior — 2026-10-26: APROBADO documentalmente.** El propietario declaró «APRUEBO LOTE 5 como documentación técnica.» sobre HEAD 9ff1a2857d0b91b9cf7925a4c36de30e377cef3d; PR #5 integrado por squash 6579ca930a9081f773b390ce61f6bc80cc6b1ae7. El [recibo de aprobación](../../PLANS.md#cierre-lote-5) prevalece sobre los estados de revisión históricos. L4-P01, la política de retención y demás propuestas mantienen su aprobación específica pendiente; sin permiso de implementación, workflows, proveedores o lote6.

ID: PH-06 · Versión: 0.2 · Fecha: 2026-10-26 · Estado documental: APROBADO · Owner: C — Coordinación · Lote: 5.

Plan futuro PROPUESTA; fuentes APROBADO, integración/seguridad/política de datos POR_VERIFICAR. RISK-10 condiciona acciones afectadas de retención/supresión y política, sin detener tareas independientes que sí tengan contratos/permisos suficientes.

## Ficha Milestone/1

| Campo requerido | Contenido del hito |
|---|---|
| Identificador | PH-06 |
| Objetivo | Completar cuentas OIDC, perfiles/historial/replay y ranked/Elo/estadísticas atómicos, con exportación y proceso de supresión honestos. |
| Dependencias | PH-05; 02.03; 03.04; 05.01–05.04; 06.01–06.03; 07.01–07.04; 04.01/04.04; 09.01–09.04; 10.01–10.04; 11.03. |
| Responsable | S implementa identidad/datos/ratings, V pantallas y Q pruebas; A entorno/retención/restore, C gobierno y propietario política/aceptación. |
| Entrada | Gates/documentos 11.03 revisados, permiso de código y de proveedor OIDC/recursos cuando corresponda; política humana resuelta antes de acciones de datos que la requieren. |
| Entregables | OIDC/PKCE y vínculo, perfil/alias/historial/replay autorizado, ranked/Elo/ledger/agregados, ranking/estadísticas, exportación y borrado distribuido. |
| Pruebas positivas/negativas | PH06-C01–C08, F-Q1-06/07/09–13, F-Q2-05/10/11 y recuperación 10.03. |
| Riesgos | Apropiación de guest, fuga de perfil/export, rating perdido/doble, barrera y borrado parcial, retención/backups incompatibles y falsa anonimización. |
| Evidencia obligatoria | QA-Evidence/1 con browser/PG/locks/ratings/ledger/sequence, operaciones/acuse por dueño, política aplicada y datos sanitizados. |
| Salida | Cuentas/ranked/estadísticas/derechos del alcance observados, cero P0/P1/críticos pendientes; no cierre completo si faltan decisiones/pruebas de funciones comprometidas. |
| Aprobación humana | GateEvaluation/1/dossier exacto, decisiones de política y cualquier operación irreversible por su alcance; no aprobación jurídica implícita. |
| Corrección/recuperación | Bloquear/revocar según contrato, conservar procesos/recibos, reparar solo con plan aprobado y reconciliación; no recalcular historia/rating ni borrar pruebas automáticamente. |

## Fuentes y contratos

Leer [07.01](../specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md), [07.02](../specs/07-identity-data/07.02-tokens-authorization-and-service-identity.md), [07.03](../specs/07-identity-data/07.03-profiles-history-and-data-rights.md), [07.04](../specs/07-identity-data/07.04-elo-leaderboard-and-statistics.md), [02.03](../specs/02-product/02.03-accounts-history-and-licensing.md), [03.04](../specs/03-game-engine/03.04-outcomes-repetition-and-replay.md), [05.01](../specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md)–[05.04](../specs/05-backend/05.04-event-schemas-and-postcommit-publication.md), [06.01](../specs/06-multiplayer/06.01-rooms-invitations-and-ranked-lifecycle.md)–[06.03](../specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md) y [ADR-0008](../adr/ADR-0008-ranked-results-and-elo-ownership.md). PRE-07/PRE-06.D1–D4 y PRE-09.C2/F1 del [REG](../decision-register.md) conservan sus dueños/parámetros.

Presentación/QA/operación: [04.01](../specs/04-visual-experience/04.01-navigation-and-accessibility.md)/[04.04](../specs/04-visual-experience/04.04-performance-and-html-fallback.md), [09.01](../specs/09-quality/09.01-engine-data-contract-and-ai-verification.md)–[09.04](../specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md), [10.01](../specs/10-operations/10.01-deployments-origins-regions-and-secrets.md)–[10.04](../specs/10-operations/10.04-releases-incidents-and-portability.md), [11.03](../specs/11-orchestration/11.03-milestones-gates-and-evidence.md), [CONF RISK-10](../conflicts.md#risk-10) y [PLANS](../../PLANS.md#ejecucion-lote-5).

Consumir OIDC transaction/session/refresh/link 07.01, cookies/JWT/JWKS/CSRF/S2S 07.02, PublicProfile/HistoryEntry/ReplayAccess/DataOperationView y handshakes 07.03, Rating/ledger/contribuciones/posición/ranking 07.04 y jerarquía 05.03. Identity posee perfiles/identidades, Game resultados/Elo/agregados; composición por API, sin SQL cruzado o ACID distribuido ficticio. Parámetros fuente 1200/K32/diez provisionales permanecen; redondeo/fórmula/admisión de detalle conservan naturaleza de 07.04.

## Componentes futuros y tickets

Rutas candidatas: services/identity para OIDC/perfiles/proceso propio, services/game para asociación/ratings/historia/derechos propios, apps/web para pantallasES/EN/HTML/3D y contratos generados. Decidir archivos exactos por TaskEnvelope/1 y permisos, sin crear un servicio de clasificación.

| Ticket / depende de | Pasos y justificación | Entrega pequeña |
|---|---|---|
| PH-06-T01 / entrada | Configurar transacciones OIDC/Code+PKCE/nonce/state/callback de un solo uso y retorno local permitido; proveedor real solo con permiso externo. | Login/cancelación/error seguros y sesión propia; C01. |
| PH-06-T02 / T01 | Vincularguest con prueba de ambas identidades, colisión/reintento y conservación contextoG; cambio aU exige intención separada. | Sin apropiación por alias/email ni transferencia de activas/receipts; C02. |
| PH-06-T03 / T02 | Perfil/alias/avatar generado y APIs públicas minimizadas, historial propio/replay compartido por recurso y paginación. | C03; correo/sesión/historial privado no públicos. |
| PH-06-T04 / T02+PH04 | Habilitar admisión ranked explícita de dos registrados con elegibilidad fijada al inicio; probar límites antiabuso del contrato sin sanción nueva. | C04; nunca migrar casual ya iniciada a ranked. |
| PH-06-T05 / T04 | Finalización transaccional con guardas/ratings actuales/ledger/sequences/contribuciones/agregados e idempotencia. | C04/C05, ambos ratings o ninguno. |
| PH-06-T06 / T05+T03 | Ranking/posición/búsqueda/provisionales y estadísticas avanzadas/evolución con denominadores/modalidad; composición por API. | C06; no confundir estadísticas generales con Elo. |
| PH-06-T07 / T03 | Separar subpasos: exportación JSON propia; petición de borrado autenticada y autorizada con raíz durable, revocación Identity y barrier S2S Game sin policyVersion; después cleanup por manifest solo con política aprobada/versionada. | C07: ACK_REVOKED persistido no es supresión; sin política conservar raíz/barrier y BLOCKED_PENDING_DATA_POLICY. Cleanup requiere policyVersion no nula, inmutable y verificable. |
| PH-06-T08 / T05+T07 | Ensayar pérdida de recibos/restauración aislada y reconciliación revocación/supresión; derechos ante fallos parciales/backups. | C08 con 10.03, sin reapertura si resucita autoridad. |
| PH-06-T09 / T06+T08 | E2EES/EN/teclado/lector HTML/3D, seguridad/concurrencia y revisión independiente; integrar dossier. | DoD de cuentas/competición/derechos completo o bloqueo visible. |

RISK-10 no bloquea automáticamente T01–T06 ni la revocación/barrera de T07 que tengan actor, permiso y contratos válidos. La barrera S2S autorizada no necesita policyVersion y su ACK_REVOKED durable impide acciones del titular sin adjudicar forfeit ni detener los efectos legítimos del rival. Sin política aprobada, conservar raíz/barrier con BLOCKED_PENDING_DATA_POLICY, sin emitir cleanup(null). Solo cleanup exige policyVersion aprobada, no nula, inmutable y verificable; cambiar policyVersion/subject/scope en un paso ya iniciado es conflicto y requiere transición autorizada. No se inventa permiso de aplicar una barrera sin petición de borrado autenticada/autorizada. Mientras los derechos requeridos de T07/T08 sigan pendientes, PH06 completo permanece sin aceptar, sin reducir el alcance del lanzamiento.

## Pruebas y evidencia futura

| Caso / fuente | Positivo esperado | Negativo esperado | Artefacto adicional |
|---|---|---|---|
| PH06-C01 / F-Q1-09/10, F-DEP-03 | OIDC correlacionado/PKCE/nonce/state/sesión válidos; rotaciónJWKS K1/K2 dentro de ventana. | Callback replay/retorno externo/firma/iss/aud/exp/Origin/CSRF/cache inválidos no autentican, sin tokens Google usados como sesión. | Browser/DB/cronología claves y contadores sanitizados. |
| PH06-C02 / F-Q1-10, F-Q2-05 | Vínculo probado y repetido asocia solo historial permitido una vez. | Alias/correo no prueban identidad; link G→U no transfiere activas/receipts ni crea Elo retroactivo. | Sujetos/sesiones/operación/ratings antes-después. |
| PH06-C03 / F-Q1-12/13, F-Q2-05 | Propio/compartido autorizado ve perfil/historial/replay versionado. | Tercero, UUID ajeno, grant vencido/revocado o alias manipulado no leen privados ni ejecutan contenido. | Matriz permisos/JSON/DOM/cursores y acceso revocado. |
| PH06-C04 / F-Q1-06, F-Q2-05 | Ranked inicial dos registrados termina con ambos ratings/agregados/ledger en un commit. | Guest/casual/IA/práctica/replay/inconclusa o vínculo posterior no actualizan Elo; unilateral/rating en el body no decide. | Elegibilidad inicial, fórmula 07.04, valores1200/K32/provisionales y diff transaccional. |
| PH06-C05 / F-Q1-06/07 | G1(A, B) y G2(A, C), ambos órdenes/colores, leen rating A actual serializado y un cierre cada una. | Doble final, guarda ausente, miembros cambiados o fallo parcial no pierden/duplican efectos; deadlock/retry completo acotado. | PG real/barreras/locks UUID/ratings/sequence/ledger/SQLSTATE, sin red bajo lock. |
| PH06-C06 / F-Q2-05, F-Q1-13 | Ranking paginado/posición/búsqueda/provisionales y estadísticas usan denominadores/modalidades versionados. | Snapshot vivo no se vende como ranking congelado; replay/noelegible no añade contribución; perfil cae sin SQLIdentity alternativo. | Dataset/consulta/cursores/contribuciones y errores seguros. |
| PH06-C07 / F-Q1-12, F-Q2-05 | JSON propio sin secretos/rival privado; petición de borrado válida permite barrier S2S sin policyVersion y ACK_REVOKED persistido. Cleanup posterior con política aprobada se reanuda idempotentemente y verifica cada dueño. | Sin política: raíz/barrier en BLOCKED_PENDING_DATA_POLICY y cero cleanup(null); policyVersion cambiada en paso iniciado produce conflicto. Game caído/ACK_REVOKED no equivale a COMPLETE ni supresión; no forfeit por delete ni anonimización irreversible por tombstone. | Actor/permiso y raíz/step/hash/acuse durables, política autorizada no nula/inmutable, manifest/estados/reintentos y retención pendiente. |
| PH06-C08 / F-Q4-08, F-REL-03 | Restore aislado preserva historia/receipts/ledger y aplica revocaciones/supresiones posteriores al backup antes de reabrir. | Receipt postbackup perdido no libera commandId ni repite Elo; refresh/grant consumido o PII suprimida no resucita. | Ensayo 10.03 y RecoveryDecision/1 con integridad/roles/seguridad, no solo bytes/hash. |

Todos usan QA-Evidence/1: commit/build/motor/contratos/DB/navegador/entorno/versiones/datos/semilla/comandos/expected/observed/artefactos y revisión. Cuentas de prueba y datos sintéticos, proveedor simulado identificado y real separado. Exportaciones/logs no se publican como fixtures. Evidencia E2E incluye ES/EN/HTML/3D/teclado/lector; no declarar WCAG solo por axe.

## Seguridad, recursos, DoD y recuperación

Política de datos/licencia no se inventa; contexto/retención/backups/derechos de terceros se decide con propietario/revisión pertinente. Sin plazo jurídico automático ni afirmar cumplimiento legal. Barreras, roles, mínimos de datos y proceso recuperable hacen visibles límites. Backup externo cifrado/destino/retención/coste y proveedor OIDC/Preview necesitan permisos; no pago ni nueva infraestructura por conveniencia.

DoD 09.04 requiere fuentes/contratos aprobados para lo implementado, 80% servicios/90% motor si cambia más críticos, observación PG/seguridad/E2E/restauración aplicables, documentación/revisión independiente/artefactos accesibles. P0/P1, crítico sin validar o requerido FAIL/BLOCKED/NOT_RUN bloquea cierre. P2 diferido con mitigación/owner/plazo/aceptación humana, sin eximir política/permiso/integridad.

C prepara AcceptanceDossier/1 y GateEvaluation/1; propietario acepta contenido/continuación concretos. Reparar Elo/datos requiere plan autorizado, no recalcular todo al arrancar ni down destructivo. A coordina incidente y contención, C gobierno y Q verifica recuperación según 10.04; conservar fallos/intentIDs/recibos y nuevo runId tras corrección. [PH-07](phase-07-hardening-and-portfolio.md) y release/recursos necesitan permisos propios; una cuenta operativa no habilita publicación general.
