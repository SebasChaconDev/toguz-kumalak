# CONF — Conflictos, riesgos y validaciones pendientes

> **Estado vigente — 2026-10-26: LOTE_6_APROBADO_DOCUMENTALMENTE.** El propietario aprobó AUD-01 como auditoría documental final sobre HEAD 6052689ff9b15a19214267d7cc6a614d17f8fcdd; PR #6 integrado por e81f01a0e94c3037e3b922894202b6eee80534e8. [Recibo de cierre](../PLANS.md#cierre-lote-6) prevalece sobre estados y paradas históricos. AUD6-F01, AUD6-F02 y AUD6-C01–C06 permanecen pendientes para sus fases. No se autoriza implementación, ejecución de PH00/PH01, workflows o despliegues.

> **Estado de trabajo — 2026-10-26: LOTE_6_EN_REVISIÓN.** Autorizada auditoría y publicación documental en docs/audit-lote-06. [AUD-01](audits/documentation-readiness.md) contrasta la base00c8e12; las40specs y246decisiones mantienen aprobación documental. Este permiso posterior sustituye la parada histórica únicamente para esta auditoría; sin cambios a specs, código, workflows, recursos, despliegues o fusión. [Alcance y ledger](../PLANS.md#ejecucion-lote-6).

> **Estado vigente — 2026-10-26: LOTE_5_APROBADO_DOCUMENTALMENTE.** PR #5 integrado en main por 6579ca930a9081f773b390ce61f6bc80cc6b1ae7, desde HEAD aprobado 9ff1a2857d0b91b9cf7925a4c36de30e377cef3d. [Recibo de aprobación e integración](../PLANS.md#cierre-lote-5) prevalece sobre cortes históricos: veinte documentos del lote5 aprobados y246decisiones con especificación redactada y aprobada documentalmente. L4-P01, retención y demás propuestas siguen pendientes de aprobación específica. Detenerse: sin producto, workflows, proveedores/despliegues o lote6.

> **Estado vigente — 2026-10-26: LOTE_5_EN_REVISIÓN.** Veinte documentos del lote5 autorizados para redacción/revisión y PR sin fusión; [alcance vigente](../PLANS.md#ejecucion-lote-5) prevalece sobre cortes históricos. Lotes1–4 APROBADOS/integrados;180decisiones con spec aprobada y66 con spec EN_REVISIÓN. L4-P01 sigue PROPUESTA pendiente de aprobación de contratos; sin implementación, recursos/despliegues o lote6.

> **Estado vigente — 2026-10-26: LOTE_4_APROBADO_DOCUMENTALMENTE.** PR #4 integrado en main por `d5dccfbe0a55c78ddc7c2c0396ead9fd2bfad910`, desde el HEAD aprobado `3ef3f7cbfd3bfc4aea837d9f1b0b414be064e8ee`. [Recibo de aprobación e integración](../PLANS.md#cierre-lote-4) prevalece sobre los cortes históricos. L4-P01 sigue PROPUESTA pendiente de aprobación de contratos antes de implementación. Detenerse: no producto, despliegues o lote 5.

> **Estado vigente — 2026-10-26: LOTE_4_EN_REVISIÓN.** Ocho specs documentales autorizadas para revisión y PR sin fusión; [alcance posterior](../PLANS.md#ejecucion-lote-4). Los cortes previos son históricos. Ninguna propuesta nueva está implementada o aprobada por el mero hecho de redactarla.

> **Estado vigente — 2026-10-25: READY_FOR_LOTE_4_DOCUMENTATION.** PR #3 integrado en main mediante c50ba9f; GO_REMOTE experimental cerrado y recursos remotos propios retirados. [Recibo verificado](../PLANS.md#lote-3-integrado-verificado) prevalece sobre los cortes históricos siguientes. El lote 4 aún no se ha generado; el producto no está implementado.

> Aprobación posterior del lote 2: el propietario declaró «APRUEBO LOTE 2» y autorizó fusionar PR #2 en main. Contenido aprobado: HEAD 598ac9a4fc6ce14cf512785fb1c4160fafb4eb9e; integrado por squash 1f401af8d76794981bbcfdf9a3aa1831cf2621f5 (2026-10-21T21:10:14Z). Este registro sustituye las menciones históricas a revisión pendiente o prohibición de esa fusión. No autoriza lote 3 ni implementación; el producto sigue sin verificar.

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: CONF · Versión: 0.14 · Fecha: 2026-10-26 · Estado documental: APROBADO · Owner: C; responsables técnicos por caso.

## Estado y reglas de uso

No existían documentos de producto que contradijeran el maestro: el árbol de trabajo estaba vacío en ETAPA 0 y al iniciar lote 1. La fuente aprobada y el inventario se preservan. Esta lista distingue riesgo de factibilidad, detalle de diseño pendiente y discrepancia textual ya matizada por el propio anexo. No acredita incidentes de un producto existente.

Un `POR_VERIFICAR` nunca se transforma en hecho probado por tener una ADR o una página oficial. Una `PROPUESTA` necesita aprobación si altera alcance, arquitectura, parámetros aprobados o rutas autorizadas. El lote 2 añade investigación documental y propuestas; no resuelve por preferencia las interpretaciones reglamentarias ni ejecuta el spike futuro. NORM-01/02/03 y el alcance reglamentario de RISK-01 quedan resueltos para TK-N1-1.0.0 por la adenda vigente. Las verificaciones técnicas y los demás riesgos continúan sin acreditar.

Referencias: [registro](decision-register.md), [trazabilidad](traceability.md), [inventario](index.md), [gates](../PLANS.md).

## Procedimiento de resolución

1. Abrir un ID estable con fuente PRE exacta, formulación original y evidencia disponible.
2. Clasificar: contradicción de requisitos, riesgo técnico o detalle pendiente.
3. Registrar impacto, propietario, documentos consumidores y evidencia futura necesaria.
4. Investigar en modo permitido. Una prueba de producto exige autorización independiente; la aprobación documental no la permite.
5. Si el resultado contradice una decisión, presentar alternativas e impacto al propietario. No cambiar el registro ni la pila por preferencia del agente.
6. Conservar aprobación humana textual y fecha; añadir ADR/cambio de requisitos cuando corresponda y actualizar matriz/consumidores.
7. Separar resolución documental de validación técnica. Una decisión puede seguir APROBADO y su implementación POR_VERIFICAR.

## Estado de riesgos

| Caso | Naturaleza | Owner | Cierre documental / técnico |
|---|---|---|---|
| RISK-01 | RESUELTO_POR_DECISIÓN_DE_PRODUCTO para TK-N1-1.0.0 | Propietario / M | Reglas cerradas; pruebas del motor pendientes, sin certificación federativa |
| RISK-02 | POR_VERIFICAR | A | Plan de factibilidad / spike autorizado |
| RISK-03 | POR_VERIFICAR | A | Modelo de consumo / cuotas y mediciones |
| RISK-04 | POR_VERIFICAR | S | Política Ably / capacidad medida |
| RISK-05 | POR_VERIFICAR | S | Sesiones / recuperación probada |
| RISK-06 | POR_VERIFICAR | A + S | Origen/previews / integración probada |
| RISK-07 | POR_VERIFICAR | S | Consistencia y reconciliación / fallos observados |
| RISK-08 | POR_VERIFICAR | S | Elo y concurrencia / simulaciones e integración |
| RISK-09 | PROPUESTA pendiente | S + A | Aislamiento físico y contratos / prueba de permisos |
| RISK-10 | POR_VERIFICAR | S + propietario | Política de privacidad aplicable / proceso verificado |
| RISK-11 | POR_VERIFICAR | A | Backup y restauración / ensayo medido |
| RISK-12 | POR_VERIFICAR | S + A | Expiración lógica / pruebas temporales |
| RISK-13 | POR_VERIFICAR | V + propietario | Procedencia y licencia / autorización de publicación |
| RISK-14 | POR_VERIFICAR | Q + V + I | Rendimiento y fuerza / benchmarks reproducibles |

<a id="risk-01"></a>

### RISK-01 — Reglamento y edición normativa

<a id="adenda-tk-n1-100"></a>

#### Adenda vigente AD-NORM-001 — 2026-10-21

El propietario adopta expresamente **TK-N1-1.0.0** definido en [resolución íntegra de REG](decision-register.md#ad-norm-001). Fuente contractual única N1, sin afirmar vigencia federativa actual. Las convenciones digitales se atribuyen al propietario, no al texto ministerial.

| Caso | Estado vigente | Resolución aplicada |
|---|---|---|
| NORM-01 | RESUELTO_POR_DECISIÓN_DE_PRODUCTO | N1 único, adopción contractual2026-10-21; no se exige autenticar vigencia federativa para aplicar este perfil |
| NORM-02 | RESUELTO_POR_DECISIÓN_DE_PRODUCTO | R-05/R-06: target→equal→turno del siguiente; atsyrau si ese jugador no tiene piedras, resultado por su kazán; sin adicional/barrido/transferencia de remanentes |
| NORM-03 | RESUELTO_POR_DECISIÓN_DE_PRODUCTO | R-07: igualdad exacta regla/tablero/kazanes/tuz/turno; inicial+confirmados activos;≥3 habilita a ambos, ventana hasta siguiente movimiento/cierre, revisión y transacción atómica |
| NORM-04, detalle de ofertas de otras fuentes | NO_APLICA_AL_PERFIL | Se descarta importar límite1/2 o reloj de fuentes excluidas. N1+R-05 exige consentimiento bilateral y no agrega aquellas restricciones |

No queda ninguna elección reglamentaria pendiente en este perfil. Se retiran peticiones de edición/interpretaciones nuevas y comparaciones con fuentes/motores externos para decidir conducta. Una variante futura exige otro perfil y autorización expresa; no es trabajo del lote3. PRE-01–PRE-11 permanecen literales con adenda posterior; tres servicios y funcionalidades no cambian.

Las ocho reglas y14 casos se integran en [03.01](specs/03-game-engine/03.01-normative-rules-and-corpus.md#matriz-normativa) y [03.04](specs/03-game-engine/03.04-outcomes-repetition-and-replay.md#atsyrau-y-prioridad). PRE-09 conserva pruebas automatizadas de integridad/corrección, historia e integración. Datos/historia corruptos o perfil desconocido siguen siendo errores técnicos: **no son bloqueos normativos de TK-N1-1.0.0**. La adopción no afirma motor implementado, seguro o probado, y el gate documental del lote2 sigue pendiente.

<details>
<summary>Historial anterior a AD-NORM-001, sustituido para TK-N1-1.0.0</summary>

Los estados, alternativas y solicitudes que siguen describen la investigación previa; no son vigentes para decidir las reglas de este perfil.

Fuentes canónicas: [PRE-03.A1](decision-register.md#pre-03-a1), [A3](decision-register.md#pre-03-a3), [E1](decision-register.md#pre-03-e1); no se reescriben por investigación. Consulta documental: 2026-10-21.

| Evidencia identificada | Acceso y límite |
|---|---|
| N1: [Laws of Togyzkumalak Game, PDF inglés, dos páginas](https://mindsportsolympiad.com/wp-content/uploads/2021/07/Rules-toguz-en-official.pdf) | Lectura íntegra; encabezado Ministerio de Turismo y Deportes de Kazajistán/Comité de Deportes. No constan edición, fecha, acto firmado o vigencia; 2021/07 es ruta de alojamiento. |
| N2: [Asociación, Astana 2013, artículo 375](https://9kumalak.com/index.php/zhyldy-zhospar/375-ojyn-a-idasy-men-erezheleri) | Encabezado, cláusulas y cierre recuperados por búsqueda; apertura directa falló. Campos de aprobación sin completar; sin autenticación íntegra ni acto vigente verificado. |
| AUX-509: [artículo 509](https://9kumalak.com/index.php/zhyldy-zhospar/509-to-yz-mala-a-idasy-men-erezhesi) | Solo información indexada; timeout directo. No se presenta como fuente íntegra ni edición vigente adoptada. |

N1 §1.5.5 describe repetición reclamada y §1.4 la declaración obligatoria de tuzdyk; no se ha demostrado contradicción con PRE-03.A1/A3. La automatización digital del tuzdyk conserva la decisión aprobada, sin afirmar certificación federativa. N1 §1.5.3 evalúa el kazán del jugador sin movimientos y no ordena expresamente barrer; N2 §1.5.1 menciona un movimiento adicional y recogida. La relación precisa entre esas formulaciones sigue pendiente; no se convierte una inferencia en regla.

| Bloqueo de aprobación normativa | Impacto / resolución exigida |
|---|---|
| NORM-01 — edición y vigencia | Todas las reglas de 03.01–03.04: identificar edición adoptable, autoridad y procedencia; revisión humana. |
| NORM-02 — atsyrau y prioridad terminal | Precisar instante de detección, significado del movimiento adicional, recogida/remanentes y precedencia con umbral de victoria; no implementar barrido por intuición. |
| NORM-03 — repetición reclamada | Resolver equivalencia de posición, actor/ventana de reclamación y rechazo inválido. PositionKey propuesto no constituye regla oficial. |

Owner M + propietario. Estado de los tres: **BLOQUEADO_PARA_APROBACIÓN_NORMATIVA**; la entrega Markdown EN_REVISIÓN es revisable, no acepta los algoritmos afectados. [03.01](specs/03-game-engine/03.01-normative-rules-and-corpus.md) registra cláusulas y corpus directo/derivado/pendiente; [03.04](specs/03-game-engine/03.04-outcomes-repetition-and-replay.md) conserva los bloqueos terminales. Se exige fuente→cláusula→lectura→posición→oráculo aprobado y resolución humana identificable. Un estado sintético que suma 162 no demuestra alcanzabilidad. El contraste con implementación independiente identificada/versionada de PRE-03.E1 sigue PENDIENTE. No adoptar variantes por similitud de nombre ni declarar victoria por una limpieza de sala.

<a id="revision-normativa-lote-2"></a>

#### Corrección normativa del PR #2 — 2026-10-21

El propietario **NO APRUEBA TODAVÍA LOTE 2** y autoriza corregir únicamente03.01–03.04 y registros necesarios, en la misma rama/PR, sin producto, lote3 ni merge. La [matriz NM-01–10](specs/03-game-engine/03.01-normative-rules-and-corpus.md#matriz-normativa) relaciona cláusula, interpretación, algoritmo candidato y caso; [03.04](specs/03-game-engine/03.04-outcomes-repetition-and-replay.md#atsyrau-y-prioridad) distingue detección/adicional/repoblación/liquidación/adjudicación.

| Bloqueo | Avance demostrado / lo que sigue sin demostrarse |
|---|---|
| NORM-01 — edición | N1 íntegro sigue sin acto/fecha/vigencia; N2/AUX-509 con acceso parcial no son una misma edición: ofertas1 vs2 en§2.6.9/§2.6.8. Reglamento federativo de torneo Almaty2023§VII y noticia gubernamental del torneo2023-11-13 remiten a reglas aprobadas, sin identificar edición. No se consiguió una edición completa verificable para adoptar. |
| NORM-02 — atsyrau | N1§1.5.3 compara kazán afectado; N2/AUX-509§1.5.1–2 mencionan adicional/recogida. Artículo didáctico KazAST2023 p.20 punto9 contempla continuidad si se devuelven jugadas, sin autoridad para adoptarla. ATS-DRAIN/DRAW/REFEED concretan impactos; salida aceptada bloqueada. |
| NORM-03 — repetición | N1§1.5.5 respalda ≥3 y declaración por un participante; no exclusividad de turno ni ventana/igualdad detalladas. REP-* son ledgers sintéticos sin ciclo legal comprobado. Contador/hashes no acreditan historia. Actor participante aclarado; habilitación por turno/ventana y definición formal continúan bloqueadas. |

Los tres permanecen **BLOQUEADO_PARA_APROBACIÓN_NORMATIVA**. La divergencia de límites de ofertas abre detalle NORM-04: confirmar texto aplicable o adaptación expresa; no imponer límite1/2 ni trasladar reloj físico al MVP. No se declara contradicción de resultados en atsyrau donde solo hay diferencias de formulación.

[REF-01](specs/03-game-engine/03.01-normative-rules-and-corpus.md#referencia-independiente) queda identificada y fijada a commit: selección documental completada; contraste de resultados **PENDIENTE**. Desactiva repetición y no muestra fase adicional explícita, de modo que no cierra ninguno de los tres bloqueos. Su marca de tuz1 se adapta a cero piedras lógicas. El corpus incorpora una ruta publicada de seis jugadas con revisión manual completa; no se ejecutó motor, replay ni prueba.

**Decisión humana solicitada, todavía sin respuesta:** (A) aportar/identificar edición completa con aprobación y aclaraciones verificables; o (B) autorizar expresamente una adaptación del producto, versionada y diferenciada de norma federativa, resolviendo frontera/adicional/continuidad/recogida/prioridad de atsyrau, igualdad/historia/reclamante/ventana y política de ofertas. B no se considera autorizada por pedir esta investigación. Fuente y PRE permanecen intactos. Aprobación documental general no se presume resolución de cada punto: exigir decisión identificable por alcance.

</details>

<a id="risk-02"></a>

### RISK-02 — Go/Gin, Vercel beta y módulos compartidos

Fuentes: [PRE-01.A4](decision-register.md#pre-01-a4), [C1](decision-register.md#pre-01-c1), [PRE-05.F4](decision-register.md#pre-05-f4), [PRE-10.A1–A2](decision-register.md#pre-10-a1). La [documentación oficial Go/Vercel](https://vercel.com/docs/functions/runtimes/go) consultada en ETAPA 0, el 2026-10-21, describía Go beta, Gin y entradas HTTP. Eso no demuestra empaquetado de módulos compartidos fuera de una raíz de proyecto ni compatibilidad del monorepo previsto.

Impacto: [01.02](index.md#spec-01-02), 05.01 y 10.01. Evidencia futura: builds y despliegues independientes de los tres servicios, importación de la misma versión del motor, transacción Neon, aviso Ably y trazas correlacionadas. El spike de PH-00 no está ejecutado ni autorizado por este lote. El adaptador aprobado de http.Handler se conserva; otras formas de empaquetado solo se evaluarán como propuestas.

<a id="risk-03"></a>

### RISK-03 — Coste objetivo cero y capacidad de cuatro proyectos

Fuentes: [PRE-01.C2](decision-register.md#pre-01-c2), [PRE-10.E1–E4](decision-register.md#pre-10-e1). Cuatro proyectos y separación de entornos pueden competir por CPU, invocaciones, conexiones y transferencia. [Vercel Hobby](https://vercel.com/docs/plans/hobby) declara uso personal no comercial; los límites deben reconfirmarse con fecha antes de cada elección operativa. No se han inspeccionado cuentas ni consumos reales.

Impacto: [01.03](index.md#spec-01-03) y 10.02. Cierre: tabla de cuotas por proveedor, ámbito de cuenta/proyecto/entorno, modelo de carga incluyendo IA, umbrales 50/75/90 y degradación al agotarse. Prohibido activar pago; ninguna alternativa reduce unilateralmente los tres servicios o el alcance de lanzamiento.

<a id="risk-04"></a>

### RISK-04 — Ably Free y presupuesto de canales

Fuentes: [PRE-01.B1](decision-register.md#pre-01-b1), [PRE-06.B4](decision-register.md#pre-06-b4), [PRE-05.E3](decision-register.md#pre-05-e3). [Ably Free](https://ably.com/docs/platform/pricing/free), consultado en ETAPA 0 el 2026-10-21, describe el paquete para experimentación/PoC y establece límites de mensajes, conexiones y canales. Su adecuación al lanzamiento no está probada.

Impacto: [06.02](index.md#spec-06-02), 10.02. Cierre futuro: consumo por partida/pestaña/entorno, renovaciones de tokens, presencia y reconexiones, capacidad negativa y comportamiento al alcanzar límites. No adoptar otro proveedor sin aprobación.

<a id="risk-05"></a>

### RISK-05 — Sesiones, vinculación y recuperación

Fuentes: [PRE-07.A2–A4](decision-register.md#pre-07-a2), [B1–B4](decision-register.md#pre-07-b1), [F1](decision-register.md#pre-07-f1), [PRE-06.C1](decision-register.md#pre-06-c1). Deben concretarse inactividad, duración máxima de invitado y sesión, caducidad de refresh, rotación/reutilización, pérdida de cookie y vinculación de cuentas. Solo está aprobado un acceso aproximadamente 15 min y compatibilidad con recuperación 24 h; no fijar los demás plazos por intuición.

Impacto: [07.01](index.md#spec-07-01), 07.02 y 06.03. Cierre: línea temporal de sesiones y partida, casos de expiración, logout con access aún vigente, rotación JWKS y vinculación sin Elo retroactivo. No recuperar una partida con un alias o invitación como sustituto de identidad verificable.

<a id="risk-06"></a>

### RISK-06 — Mismo origen, cookies, OAuth y previews

Fuentes: [PRE-07.B4](decision-register.md#pre-07-b4), [PRE-10.A3](decision-register.md#pre-10-a3), [B1–B3](decision-register.md#pre-10-b1). Rewrites entre cuatro proyectos, callback OIDC y protección de previews deben funcionar con audiencias, cookies y credenciales separadas. La [protección de despliegues Vercel](https://vercel.com/docs/deployment-protection), consultada en ETAPA 0 el 2026-10-21, exige considerar autenticación de solicitudes protegidas.

Impacto: [10.01](index.md#spec-10-01) y 07.02. Evidencia futura: rutas del mismo origen y callbacks en entorno sintético, invocaciones servicio-servicio protegidas y ausencia de datos/credenciales productivos en previews. No emitir credenciales ni configurar proveedores ahora.

<a id="risk-07"></a>

### RISK-07 — Commit confirmado y notificación perdida

Fuentes: [PRE-05.C2–C3](decision-register.md#pre-05-c2), [E2](decision-register.md#pre-05-e2), [PRE-06.B2](decision-register.md#pre-06-b2). Un commit puede completarse y la publicación posterior fallar. Sin outbox duradero no hay garantía de publicación por cada commit; una respuesta o mensaje perdido tampoco revierte el movimiento confirmado.

Impacto: [05.03](index.md#spec-05-03), 05.04 y 06.02–06.03. Cierre: máquina de reconciliación REST, IDs/versiones, política acotada de comprobaciones, duplicados y reintentos. Probar commit fallido, aviso perdido, mensajes desordenados y retorno de foco. No prometer entrega exactamente una vez ni introducir colas/Redis preventivamente.

<a id="risk-08"></a>

### RISK-08 — Resultado, Elo y concurrencia entre partidas

Fuentes: [PRE-06.D1–D4](decision-register.md#pre-06-d1), [PRE-07.E1–E4](decision-register.md#pre-07-e1), [A3](decision-register.md#pre-07-a3). Game posee resultado, Elo y agregados. Además del lock por partida, deben concretarse finalizaciones simultáneas que afectan a un mismo jugador, orden de locks, reintentos, redondeo y abuso de enfrentamientos.

Impacto: [07.04](index.md#spec-07-04), 05.03 y 06.01. Cierre: casos de 1200/K32 y diez provisionales, doble finalización y carreras; invitados/IA/práctica/inconclusas no modifican Elo; vinculación no reclasifica partidas retrospectivamente. Los umbrales antiabuso y redondeo no se eligen en este lote.

<a id="risk-09"></a>

### RISK-09 — Propiedad separada y distribución física pendiente

Fuentes: [PRE-05.D1](decision-register.md#pre-05-d1), [PRE-07.C4](decision-register.md#pre-07-c4), [F3](decision-register.md#pre-07-f3). Está aprobado separar propiedad de Identity/Game y mantener AI stateless. Falta elegir la disposición física exacta de bases/esquemas/proyectos y sus roles de acceso, sin compartir propiedad ni permitir SQL cruzado.

Impacto: [05.01](index.md#spec-05-01), 05.03 y 07.03. Cierre documental: propuesta de aislamiento con privilegios mínimos y coste; aprobación si cambia una decisión. Cierre técnico futuro: credenciales de un servicio no leen/escriben propiedad del otro; perfiles por API; borrado distribuido con acuses/reintentos, sin ACID entre servicios.

<a id="risk-10"></a>

### RISK-10 — Privacidad y aplicabilidad normativa

Fuentes: [PRE-07.D1/D4](decision-register.md#pre-07-d1), [F1–F3](decision-register.md#pre-07-f1). Faltan ámbito territorial, jurisdicción, retención concreta, encargados y procedimiento aplicable a eliminación/anonimización. No se declara cumplimiento GDPR ni conclusión jurídica en esta misión.

Impacto: [07.03](index.md#spec-07-03), 10.01 y 10.03. Cierre: decisión del propietario sobre contexto de uso y política, revisión normativa pertinente y escenarios de exportación propia/borrado parcial/reintento/verificación. Datos privados y correo no pasan a perfiles públicos, trazas o ejemplos.

<a id="risk-11"></a>

### RISK-11 — Backup externo y restauración no acreditados

Fuentes: [PRE-10.C2](decision-register.md#pre-10-c2), [F3](decision-register.md#pre-10-f3). [Neon](https://neon.com/docs/introduction/plans), consultado en ETAPA 0 el 2026-10-21, documenta limitaciones de restauración del plan Free. No se ha seleccionado almacenamiento externo cifrado ni demostrado su gratuidad, retención o recuperación.

Impacto: [10.03](index.md#spec-10-03) y 10.04. Cierre: plan de cifrado/claves/almacenamiento/retención, recuperación de ambas propiedades de datos y ensayo autorizado en entorno aislado; medir pérdida recuperable y tiempo real. No garantizar RPO cero ni RTO sin ensayo.

<a id="risk-12"></a>

### RISK-12 — Caducidades distintas y limpieza eventual

Fuentes: [PRE-06.A2](decision-register.md#pre-06-a2), [C2](decision-register.md#pre-06-c2), [PRE-10.C4](decision-register.md#pre-10-c4). Sala pendiente de 30 min e iniciada recuperable 24 h son políticas distintas. Faltan definición exacta del instante de inicio/reinicio de cada vencimiento y frontera temporal ante solicitudes concurrentes.

Impacto: [06.01](index.md#spec-06-01), 06.03 y 10.03. Cierre: cronograma persistido y comparaciones en consulta/comando, con pruebas antes/en/después del límite y limpieza demorada. La limpieza diaria no concede validez a recursos expirados ni cambia inconclusa por derrota.

<a id="risk-13"></a>

### RISK-13 — Licencia, recursos y referencias culturales

Fuentes: [PRE-01.D4](decision-register.md#pre-01-d4), [PRE-04.A3](decision-register.md#pre-04-a3), [F2](decision-register.md#pre-04-f2). Publicación de portafolio no elige licencia. No existen assets originales ni inventario de procedencia; tampoco está contrastado el repertorio cultural.

Impacto: [02.03](index.md#spec-02-03) y 04.02. Cierre: elección expresa de licencia por el propietario y ficha por recurso con autor/origen/licencia/atribución/permisos. RDomino es inspiración de experiencia, no autorización para copiar recursos, diseños distintivos o código.

<a id="risk-14"></a>

### RISK-14 — Rendimiento, accesibilidad y fuerza IA sin mediciones

Fuentes: [PRE-03.E4](decision-register.md#pre-03-e4), [PRE-04.E1/F1](decision-register.md#pre-04-e1), [PRE-08.F1–F4](decision-register.md#pre-08-f1), [PRE-09.D4/F2](decision-register.md#pre-09-d4). FPS 60/30 y WCAG 2.2 AA son objetivos; no existen dispositivos de referencia, latencias, benchmarks o fuerza IA observados.

Impacto: [04.04](index.md#spec-04-04), 08.01–08.02 y dominio 09. Cierre: entorno/dispositivos/corpus/semillas reproducibles, resultados y evaluación manual de accesibilidad/3D; fijar SLO definitivos después de medir. Usar alternativa HTML y degradación prevista sin inventar un motor JS u offline.

## Discrepancias textuales y matices de la fuente

<a id="discrepancy-01"></a>

### DISCREPANCY-01 — Ejemplo de ID en la plantilla

El §6 del maestro ejemplifica `PRE-xx.Cy`; los identificadores reales son `PRE-NN.Xn` con gruposA–F. Tratamiento editorial propuesto: citar siempre los IDs originales y anchors estables; no renumerar las 246 decisiones. Evidencia: REG y Anexo A. No cambia ningún requisito.

<a id="discrepancy-02"></a>

### DISCREPANCY-02 — Fotorrealismo y realismo estilizado

[PRE-02.E2](decision-register.md#pre-02-e2) usa premium/fotorrealista; [PRE-04.A1](decision-register.md#pre-04-a1) matiza con realismo físico estilizado sin extremo obligatorio. Aplicar el matiz posterior conservando calidad premium. Estado: interpretación fundada en fuente, documentación EN_REVISIÓN; no aceptar visuales sin revisión incremental.

<a id="discrepancy-03"></a>

### DISCREPANCY-03 — Undo y tablas ampliadas

[PRE-02.C2](decision-register.md#pre-02-c2) prohíbe undo ordinario y añade práctica reversible, desarrollada en [PRE-03.D1–D2](decision-register.md#pre-03-d1). [PRE-02.C4](decision-register.md#pre-02-c4) se amplía con [PRE-03.A1–A2](decision-register.md#pre-03-a1). Son matices explícitos, no permisos para undo ranked, Elo en práctica o empate automático no normativo. Falta cotejo reglamentario RISK-01; este lote no lo resuelve.

## Criterios de aceptación del registro de conflictos

- AC-CONF-01: cada riesgo tiene fuente, impacto, owner, acción y evidencia futura; un riesgo sin prueba permanece abierto.
- AC-CONF-02: una preferencia técnica no sustituye una decisión; se exige aprobación y trazabilidad antes de alterar registro o consumidores.
- AC-CONF-03: una validación documental satisfactoria no cambia factibilidad a probada.
- AC-CONF-04: cerrar lote 1 no cierra automáticamente ninguno de los 14 riesgos técnicos ni aprueba las tres interpretaciones documentales.


## Avance documental del lote 2, sin cierre técnico

[01.01](specs/01-infrastructure/01.01-stack-monorepo.md) conserva tres Go/Gin, único motor Go y propiedad separada. [01.02](specs/01-infrastructure/01.02-platform-feasibility.md) desarrolla POC-01–11, todos PENDIENTE: Go/Vercel beta, empaquetado compartido, Neon transaccional, avisos Ably y entornos no han sido probados. [01.03](specs/01-infrastructure/01.03-cost-capacity-evolution.md) identifica cuotas oficiales fechadas y presupuesto por cuenta/proyecto/entorno, sin consumos reales ni gasto autorizado. RISK-02/03/04/06/09/11 no se cierran por esas lecturas.

[02.01](specs/02-product/02.01-scope-and-modes.md) conserva base y ampliado para lanzamiento; [02.02](specs/02-product/02.02-local-flows-and-i18n.md) distingue persistencia de navegador y autoridad Go, sin offline. [02.03](specs/02-product/02.03-accounts-history-and-licensing.md) mantiene pendientes jurisdicción/política y elección de licencia; RISK-10/13 siguen abiertos. Las propuestas de estructura, campos, migración local, concurrencia y PositionKey necesitan revisión de diseño; no alteran PRE. La evidencia y los hallazgos de integración se registran en [reporte L2](../PLANS.md#informe-lote-2).


## Prevalencia de la adenda normativa

Las menciones anteriores de investigación o validación de RISK-01, incluida DISCREPANCY-03 y avances previos, quedan sustituidas en su alcance reglamentario por [AD-NORM-001](#adenda-tk-n1-100). Las fuentes ajenas a N1 se conservan solo como historia, no como alternativas normativas disponibles. La resolución por decisión no cierra pruebas, rendimiento, seguridad, licencias o factibilidad; esos controles siguen siendo POR_VERIFICAR. No existe una nueva solicitud de pronunciamiento normativo para TK-N1-1.0.0.

## Adenda de estado — lote 2 aprobado e integrado

La instrucción expresa «APRUEBO LOTE 2» satisface el gate documental y autoriza el cierre de PR #2, integrado por squash `1f401af8d76794981bbcfdf9a3aa1831cf2621f5` desde HEAD aprobado `598ac9a4fc6ce14cf512785fb1c4160fafb4eb9e`. CONF y las especificaciones del lote están APROBADOS documentalmente. No se cierran por ello riesgos técnicos ni pruebas POR_VERIFICAR. NORM-01/02/03 conservan RESUELTO_POR_DECISIÓN_DE_PRODUCTO para TK-N1-1.0.0; no hay bloqueos normativos pendientes para ese perfil. Los informes de gate pendiente son históricos. No se inicia lote 3 ni implementación.

<a id="lote-3-riesgos"></a>

## Adenda de lote 3 — propuestas técnicas y límites de evidencia

La solicitud posterior del propietario autoriza once specs y sus registros para revisión, sin fusionar el PR ni ejecutar producto. Esta adenda sustituye las paradas históricas únicamente en ese alcance. Los estados APROBADO de lotes 1 y 2 y AD-NORM-001 no se revocan; las nuevas versiones de integración están EN_REVISIÓN.

TK-N1-1.0.0 y sus 14 ACC son entrada vinculante: no se consultan otros reglamentos/motores ni se reabren NORM-01/02/03, que siguen RESUELTO_POR_DECISIÓN_DE_PRODUCTO. Las incertidumbres siguientes son técnicas o de política de datos, nunca una elección normativa pendiente.

| Propuesta / riesgo relacionado | Alcance para revisión y dueño | Evidencia que no existe todavía |
|---|---|---|
| L3-P01, RISK-05/06 | 07.01: guest sliding48h, reserva firmada de horizonte recuperable, refresh prepare/commit y recibo cifrado acotado; vínculo existente solo a historial terminado, con UUID y claves idempotentes originales intactos | Carreras, cookies, criptografía, Google web/PKCE y recuperación de respuesta perdida; límites de cookie perdida/tardía explícitos |
| L3-P02, RISK-05/06 | 07.02: ES256/JWKS acotado; cookies/CSRF persistentes con retención de transporte separada de vigencia lógica; S2S ligado a operación, método, body y parámetros canónicos JCS; cuotas persistidas, grant IA de un intento y capacidades Ably restringidas | Revisión documental de fronteras realizada; bibliotecas, aislamiento, rotación, cookies y fallos de proveedor POR_VERIFICAR. Ninguna credencial emitida |
| L3-P03, RISK-07/08/09 | 05.03: guardas humanas UUIDASC→roomsASC→gamesASC→ratingsASC→derivados; idempotencia persistida, history_entries común y finalización atómica entre partidas de un mismo jugador | Diseño revisado documentalmente; PostgreSQL real, índices, interbloqueos, reintentos, commit incierto y pools siguen POR_VERIFICAR |
| L3-P04, RISK-08/10 | 07.03/07.04: separar permiso para actuar de registros mínimos que reciben efectos; borrar no equivale a rendirse ni cambia elegibilidad inicial o derechos del rival | Política aplicable de retención/supresión, anonimización o pseudonimización, revisión humana y pruebas distribuidas |
| L3-P05, RISK-04/07/12 | 06.01–06.03: pendiente30min frente a recuperación24h por participante; propia actividad nueva, nunca sondeo/rival/Ably/duplicado; límite temporal tras locks, en igualdad vence y terminal no se reevalúa | Contratos y cronologías revisados documentalmente; reloj real, regresión, carreras, cadencias, consumo y reconexión requieren pruebas futuras. Outage no pausa automáticamente y browser no adjudica resultado |
| L3-P06, RISK-02/03/04/09 | Runtime, módulos compartidos, disposición física Neon, SDK y cuotas permanecen POR_VERIFICAR según 01.02/01.03 | Sin PoC, cuentas de proveedores, builds, capacidad, coste o despliegue acreditados |
| L3-P07, RISK-05/07/09/10 | 05.02: catálogo REST v1/OpenAPI3.1, límites por operación, uint64 como string, JCS y binding URI, cursor15min/ranking vivo sin snapshot, no-store de estado/actividad y separación barrier/cleanup por paso inmutable | Revisión documental de 61 operaciones y ejemplos; releases de generadores/validadores, transporte efectivo y comportamiento concurrente POR_VERIFICAR |
| L3-P08, RISK-08/10 | 07.04: fórmula Elo binary64 orientada UUID y redondeo único; entero int32 comprobado, secuencias por jugador, métricas por owner sin doble conteo, rebuild de una política. Admisión propuesta de tres inicios por pareja/24 h antes de Start, sin veto retroactivo ni máximo de una activa | Seis ejemplos W/D/L y dos órdenes concurrentes revisados aritméticamente en memoria; Go, locks, overflow, equivalencia numérica, eficacia antiabuso y retención no probados. 1200/K32/diez provisionales no cambian |
| L3-P09, RISK-05/06/09/10 | 07.03: alias ASCII 3–24, tres cambios/día y reserva30d; avatar generado, replay terminado guest/user TTL24h/máximo7d, export per_owner temporal24h, cache pública60s. Delete prepare/confirm con proof/nonce previos, primera confirmación10min y seguimiento30d solo de la misma raíz | Parámetros PROPUESTA; revisión de contratos y fallos realizada. Navegadores, scopes, pérdida de pruebas, retención y alcance jurídico requieren implementación/pruebas/política autorizadas; ACK_REVOKED nunca prueba supresión completa |
| L3-P10, RISK-03/04/07/12 | 06.01/06.02: ReadyConsent sin JWT y horizonte de ambos antes de Start; revancha TTL10min sin Notice anterior. Reconciliador con activityVersion por UUID, epoch/single-flight y caps compartidos: check60s, fallback5s hasta10min,30 lecturas/min y120/10min; terminal15s máximo40/10min. Presupuesto Ably cuenta entregas/chunks/pestañas, con reservas de planificación propuestas | Contratos y fuentes públicas revisados, sin SDK/cuenta/carga. Cadencias/caps/reservas requieren aprobación y pruebas, no garantía de capacidad 0USD ni degradación automática por umbrales50/75/90. La fábrica pura r0 se mantiene y Game alinea metadata al Start oficial, con aparición inicial1 |
| L3-P11, RISK-05/07/10/12 | 06.03: ActivityDecision inmutable separada de RecoveryView actual; heartbeat60s±20% visible y consentido, cuotas servidor4/min sujeto+partida y20/min sujeto acumuladas30/100; cliente4/min y20/10min. Journal sessionStorage cerrado no secreto hasta24h, sin invitación/código, y límites explícitos por pruebas/Set-Cookie tardío | Regresión temporal contra ambos participantes falla con DATA_UNAVAILABLE; expiry_pending vacía allowedActions sin afirmar cierre. Controles documentales no prueban cookies, storage, reloj o recuperación universal. Retención/política y pruebas reales pendientes |

RISK-10 requiere decisión humana sobre contexto/política antes de completar una implementación de derechos de datos: el UUID/tombstone puede seguir siendo pseudónimo. No se afirma cumplimiento legal ni se inventa un plazo jurídico. Las propuestas técnicas pueden ser revisadas con este lote; publicar documentación no autoriza aplicarlas a datos reales.

Las frases históricas «no fijar en este lote» describen su alcance original. La solicitud actual permite concretar propuestas técnicas para revisión, sin convertirlas en decisiones aprobadas ni cambiar fuentes. La auditoría final registrará contradicciones reales si aparecen; una sugerencia preliminar descartada antes de integrarse no se presenta como cambio de requisito autorizado.
<a id="conf-gate-35-close"></a>

## Adenda Gate3.5 y continuidad documental — 2026-10-25

[GO_REMOTE y límites publicados](https://github.com/SebasChaconDev/toguz-kumalak/blob/504caa5d8d4f538e131e1beffc8843cb8938772a/spikes/infra-feasibility/REMOTE_AUDIT.md#remote-close-20261025) mitigan parcialmente RISK-02/03/04/06/07/09: runtime/paquetización, cuotas gratuitas observadas, capacidades Ably, origen/cookie/OIDC, transacciones/reconciliación y aislamiento SQL. No se cierran globalmente: faltan producto, SDK/generadores, seguridad completa, métricas/SLO, restauración y demás pruebas de cada requisito.

RISK-05 sesiones, RISK-08 Elo y RISK-10 política/privacidad conservan pendientes propios. No se aprobó una política jurídica mediante el spike ni se equipara una sonda HMAC a JWT/JWKS productivo. Tampoco se atribuyen mediciones de colas PgBouncer, reactivación Neon, caída física S2S o pérdida TCP real.

Desviación de ejecución explícita: SQL80 técnico previsto superado;117 operaciones conservadoras declaradas, sin traza wire exacta. Recursos propios retirados,0USD contratado y cero modificaciones a recursos previos. Esta desviación no se oculta como cumplimiento presupuestario ni altera requisitos de producto.

La revisión independiente del lote3 encontró únicamente dos Minor05.02, corregidas. No quedan conflictos normativos TK-N1 ni hallazgos técnicos Critical/Important que impidan redactar04/08; los riesgos de implementación anteriores permanecen POR_VERIFICAR.

<a id="lote-4-riesgos"></a>

## Lote 4 — condiciones técnicas y control de cambios, 2026-10-26

Autoridad posterior: «AUTORIZO LA GENERACIÓN Y PUBLICACIÓN DOCUMENTAL DEL LOTE 4». La redacción de las ocho specs y este registro no altera las 246 PRE, dueños, reglas, decisiones de arquitectura o archivos aprobados de lotes anteriores. NORM-01/02/03 siguen RESUELTO_POR_DECISIÓN_DE_PRODUCTO para TK-N1-1.0.0. Las adendas normativas no se reabren ni se atribuyen convenciones digitales al texto literal del reglamento.

La auditoría Gate4.0 concluyó GO_LOTE_4_CON_CONDICIONES: ninguna contradicción P0/P1 demostrada bloquea la redacción. La clasificación siguiente distingue diseño propuesto para revisar de evidencia empírica que todavía falta; no declara cerrado globalmente ningún riesgo por un PASS sintético o experimental.

| ID / prioridad / naturaleza | Fuente, discrepancia o incertidumbre | Tratamiento documental y responsable | Aprobación / evidencia pendiente |
|---|---|---|---|
| L4-P01 / P2 / NON_BLOCKING | 05.02 RequestAIMove requiere gameId; SimulationRequest solo move/legal/preview/validate. No hay binding aprobado para IA de una rama restaurada/escenario ni campo recovery para solicitar respaldo | I,08.02/08.03: propuesta explícita de solicitud hipotética ligada a sujeto/rama/generación/contexto/intención/cuota, con autoridad Go y respuesta no deportiva; recuperación del efecto IA oficial mediante extensión propuesta. No pasar tableros cliente a SubmitMove ni fingir que existe ya otra operación autorizada | Antes de implementar el delta se exige aprobación humana del cambio y revisión de contratos 05.02/05.03/07.02; este lote no modifica esos archivos. Revisar transporte, grant/recurso, scopes, metadatos mínimos y negativos sin crear cuarto servicio/DB AI |
| L4-P02 / P2 / NON_BLOCKING | 03.04 define PositionKey e historia; mismo tablero y count actual no garantizan valores de búsqueda iguales si difieren ocurrencias relevantes futuras | I,08.01: política segura de TT por contexto histórico o exclusión de scores/bounds incompatibles; contexto separado por rama, igualdad exacta además de hash. La tercera aparición no produce draw automático | Pruebas con posición idéntica/historias distintas, ramas con reclamación y TT activada/desactivada; conservar 14 ACC |
| L4-P03 / P2 / NON_BLOCKING | Undo/redo/reset puede reutilizar revisión local y recibir AI/preview de otra rama | I,08.03 y V,04.01/04.03: generación e intención además de cursor/revisión, reconstrucción histórica y descarte explícito de futuro; local no significa offline | Roundtrip de almacenamiento versionado, rama abandonada, respuesta tardía y cuota/storage corrupto. Cero GameDB deportiva, Elo o ranking por práctica |
| L4-P04 / P2 / NON_BLOCKING | 05.02/07.02 autorizan orquestación y grants, pero crash/timeout no garantizan resultado IA o entrega | I,08.02: recuperación acotada por consulta/reentrada autorizada, respaldo Go barato y legal, revalidación transaccional de un efecto; resultado tardío descartado. Ningún worker/cola/Redis añadido | Integración futura con PostgreSQL, callback incierto, crash tras commit humano y carreras AI/respaldo; claim no equivale a movimiento ni exactly-once de cómputo |
| L4-P05 / P2 / NON_BLOCKING | Estado Go y eventos completos son autoridad; cámara, instancias y fotogramas no son posiciones de reglas | V,04.02/04.03: PitRef estable, tuzdyk por ubicación/propietario, selección/confirmación separadas y animación posterior a confirmación; no barrido atsyrau ni claim restringido por turno | Reconstrucción por snapshot/eventos, skip/movimiento reducido/rotación y orden/huecos; invariantes en estados completos y no en fotogramas intermedios |
| L4-P06 / P2 / NON_BLOCKING | HTML, acceso, reconciliación y presupuesto compartido deben cubrir todos los modos, no solo mostrar el tablero | V,04.01/04.04 e I,08.03/08.04: equivalencia esencial HTML/3D, foco/intención pendientes preservados, ES/EN, teclado/móvil, scheduler único de lecturas y actividad separada según06.02/06.03 | Playwright, lector/teclado, pérdida WebGL/red, recuperación REST y prueba negativa ranked; sin bucles de polling por componente ni motor frontend |
| L4-P07 / P2 / BLOCKS_IMPLEMENTATION_ONLY | RISK-14: FPS, dificultad/fuerza, memoria, descarga y CPU no están medidos; SDK/runtime de producto tampoco fijados por el spike | V,04.04 e I,08.01/08.02: presupuestos y perfiles de medición propuestos con corpus/dispositivos/versiones/duración; cancelación cooperativa y memoria de TT/contexto incluidas | 60/30 son metas; validar matriz de versiones y dispositivos, memoria, nodos, tiempo, legalidad y dificultad antes de afirmar cumplimiento. No Elo de IA inventado |
| L4-P08 / P2 / BLOCKS_IMPLEMENTATION_ONLY | RISK-10/13: derechos de uso, licencia, revisión cultural y política de datos no acreditados | V,04.02 e I,08.04: procedencia verificable, permisos/atribución, no copia de RDomino, texto/contenido externo tratado como no confiable; minimización en telemetría/localstorage | Decisión humana de política/licencia cuando corresponda y autorización de recursos definitivos antes de publicar lo afectado. La redacción no atribuye certificación o cumplimiento jurídico |
| L4-P09 / P2 / NON_BLOCKING | Spike: módulo de hash/AI sintética, HMAC, muestras limitadas, SQL80 superado y falta de traza exhaustiva | Coordinación: conservar GO_LOCAL/GO_REMOTE acotados en504caa5 como evidencia experimental y enlaces exactos, no prueba de funcionalidad o seguridad productiva | PRE-09 y verificaciones de cada dueño siguen pendientes; sin nuevos builds/proveedores en este lote |
| L4-P10 / P2 / NON_BLOCKING | Un escenario synthetic libre no acredita si su tablero fue copiado de una partida ranked; purpose declarado no demuestra procedencia | I/S,08.02/08.04: ayudas no se ofrecen en UI ranked y Game rechaza solicitudes ligadas a recurso oficial ranked; no se promete detectar una copia reenviada como práctica independiente ni impedir ayuda externa. Negativos de permisos verifican el recurso real, no una procedencia inexistente | Una política antitrampa más fuerte o prohibición de práctica concurrente requeriría decisión humana posterior y contrato/evidencia propios; no se introduce en este lote |

Las propuestas técnicas se someten a la revisión del lote, sin transformar su naturaleza en APROBADO por actualizar el estado de un Markdown. El delta L4-P01 tiene revisión de cambio separada antes de implementación: aprobar ocho documentos no altera silenciosamente el catálogo o scopes de las specs previas. No hay una elección reglamentaria pendiente para este perfil.

Los rótulos históricos de README/PROMPTS/AGENTS y las specs aprobadas se conservan por estar fuera del alcance de edición; el permiso explícito posterior y los encabezados de estos cuatro registros identifican el estado vigente. Una limpieza editorial futura no autoriza cambiar contenido canónico.

Revisión independiente final: L4-R01 exige discriminar intent=hint/opponent_move en08.02/08.03 y negativos08.02/08.04: recibir una pista no selecciona, envía movimiento ni cambia timeline/progreso; posterior jugada exige intención humana separada. L4-R02 exige acotar la garantía ranked comoL4-P10. Son correcciones de coherencia de borradores dentro de PRE-08.D2, no cambios de reglas o funcionalidades aprobadas; su verificación final consta enPLANS. El P3 visual de sumas de bytes se corrigió en04.04 sin convertir presupuesto propuesto en medición.

## Adenda de aprobación documental del lote 4 — 2026-10-26

El propietario aprobó el lote4 como documentación técnica sobre HEAD `3ef3f7cbfd3bfc4aea837d9f1b0b414be064e8ee`. PR #4 integrado por squash `d5dccfbe0a55c78ddc7c2c0396ead9fd2bfad910`; [recibo y condiciones](../PLANS.md#cierre-lote-4). Los estados de revisión anteriores son cortes históricos; L4-R01/L4-R02 y P3 visual ya fueron corregidos/verificados antes de esa aprobación.

**L4-P01 se conserva PROPUESTA técnica pendiente de aprobación de contratos antes de implementar**. La aprobación del lote no modifica el catálogo, persistencia o permisos de05.02/05.03/07.02. Tampoco convierte presupuestos, SDK, rendimiento, fuerza, licencias, privacidad o seguridad productiva POR_VERIFICAR en resultados demostrados. No se cierra globalmente ningún riesgo por cambiar el estado de una spec. TK-N1-1.0.0 y sus reglas/14ACC permanecen intactos; no se implementa producto, despliega o inicia lote5.

<a id="lote-5-riesgos"></a>

## Lote 5 — propuestas y verificaciones posteriores, 2026-10-26

La autorización posterior permite20documentos y sus registros para revisión/PR, sin código, CI instalada, recursos o despliegue. Lotes1–4 y su aprobación permanecen intactos; el cierre administrativo del lote4 ya está en b91ba98. Las66decisiones restantes se desarrollan en nueve specs EN_REVISIÓN, sin modificar PRE/ADR/especificaciones aprobadas ni reabrir TK-N1-1.0.0.

| ID / fuente / owner | Propuesta o pendiente | Gate y evidencia futura exigida |
|---|---|---|
| L5-P01 / L4-P01, S/I/C | Procedimiento11.03 para IA de rama/hints/cancelación/recuperación y delta de05.02/05.03/07.02 | Sigue PROPUESTA; aprobar lote5 no aprueba contratos. Cambio concreto con versión/commit, impacto y consumidores, aceptación humana de edición, actualización de fuentes y pruebas antes de implementar consecuencias. No nuevo servicio/DB AI/Redis |
| L5-P02 / PRE09/11, Q/C | QA-Evidence/1, AcceptanceDossier/1, TaskEnvelope/1, GateEvaluation/1 y formatos de casos/tickets | Son contratos documentales de ejecución futura, no APIs de producto o automatizaciones instaladas. Revisión de formatos y campos; requerido FAIL/BLOCKED/NOT_RUN impide salida, no se sustituye por declaración de agente/exit0 |
| L5-P03 / RISK02/03/06, Q/A | Gates QA-FAST/INTEGRATION/FULL/RELEASE/SMOKE, main protegida y verificaciones por trigger/consumidor | Antes de configurar/ejecutar, permisos explícitos y capacidades reales gratuitas verificadas; PR externo sin secretos, motor dispara Game/AI, PostgreSQL real y generadores fijados. Hoy cero workflows/checks activos no demuestra CI o protección implementadas |
| L5-P04 / RISK14, Q/V/I/A | Protocolos de muestras/percentiles/entornos y presupuestos de medición; objetivos60/30FPS y cobertura90/80 se conservan en su naturaleza fuente | Benchmarks observados y revisión de umbrales no aprobados. Sin SLO/fuerza/Elo/latencia/capacidad o conformidadWCAG declarados por planificación, HTML o herramienta automática |
| L5-P05 / RISK10/11, A/S/propietario | BackupManifest/RestoreTrial/MigrationPlan/RetentionPolicy y respaldo externo cifrado, proveedor/retención/RPO/RTO por decidir/verificar | Restauración aislada y tiempos medidos; cortesIdentity/Game, datos/recibos/ratings y efectos posteriores comparados. No resucitar refresh/grants ni reejecutar un efecto incierto. Sin verdad posterior, mantener cuarentena. Epoch/protocolo nuevo requiere control de cambios antes de fuentes/código |
| L5-P06 / RISK03/04/10, A/S | TelemetryPolicy/MeasurementSeries/QuotaObservation/RatePolicy, sampling/cardinalidad/retención y alertas candidatas | Política de minimización/privacidad y medición de coste/capacidad; correlación sin tokens/cookies/proofs/DSN/PII. Contadores globales persistidos, no RAM;50/75/90 activan revisión, no pago o cambio de alcance automático |
| L5-P07 / PRE10.F1/F2/F4, A/Q | ReleaseManifest/CompatibilityMatrix/IncidentRecord/RecoveryDecision/PortabilityTrial, flags y runbooks | Compatibilidad demostrada, smoke tras despliegue autorizado, restauración y aprobación de reapertura. Rollback de código no revierte datos; no bypass de integridad, protección, permisos o requisito crítico |
| L5-P08 / RISK13/14, V/Q/propietario | Assets/identidad cultural/licencias, dispositivos y revisión visual/accesible reales | Falta de recursos finales o permisos no impide documentar, pero bloquea publicar esos recursos/atribuir autenticidad. Baselines no se regeneran para esconder diferencias; revisión manual y prueba de lector/teclado requeridas |
| L5-P09 / PRE11.D2/E3, C/Q | SecuenciaPH00–07 y harness de tareas/versiones/autoridad/evidencias | Aprobación documental, código y externos independientes. PH04 inicialmente casual/invitados no adelanta EloPH06 ni elimina ranked del lanzamiento; L4-P01 condiciona solo tareas dependientes. Ejemplo sintético no acredita implementación |

El spike504caa5 permanece antecedente experimental: runtime/builds/SQL/Ably/cookie/negativos acotados, recursos retirados y desviaciónSQL80→117 conservadora reconocida. No confirma seguridadJWT/JWKS, sesionesGoogle, motorTK, Negamax, p95, RPO/RTO, suspensiónNeon/PgBouncer o coste de producto. PH00 no ordena repetir automáticamente el experimento.

Hallazgos documentales corregidos antes de integrar: L5-R01/P2,09.01 separa apariciones de repetición del lifecycle/history_ordinal de un cierre; L5-R02/P3,09.04 diferencia A/coordinación operativa de C/gobierno y Q/verificación. Revalidación independiente ADDRESSED; sin alterar las fuentes aprobadas03.01/03.04/05.03 ni otras anteriores. El borrador10.04 conserva su misma asignación operativa. Las pruebas y escenarios nuevos son futuros, no ejecución observada.

Revalidación final dePH/harness: L5-R03/P2 yL5-R04/P2 ADDRESSED. PH05-T06 consumeT04 además deT03/T05; PH06-T07/C07 separa revocación/barrier sinpolicyVersion yACK_REVOKED durable de cleanup con política aprobada/no-null/inmutable. Sin política, raíz/barrier BLOCKED_PENDING_DATA_POLICY; ACK_REVOKED no es COMPLETE. H-BUILD/H-CTX propagan esos contratos aprobados sin cambiarlos. Revisor independiente:20documentos CONFORMES documentalmente, sin hallazgos restantes; gatehumanoL5 y propuestas L5-P01–09 siguen pendientes, sin pruebas de producto.

## Adenda de aprobación documental del lote 5 — 2026-10-26

El propietario aprobó HEAD 9ff1a2857d0b91b9cf7925a4c36de30e377cef3d; PR #5 integrado por squash 6579ca930a9081f773b390ce61f6bc80cc6b1ae7. [Recibo](../PLANS.md#cierre-lote-5). L5-R01–R04 siguen ADDRESSED y gate documentalL5 satisfecho; las menciones anteriores de gate pendiente son cortes históricos.

**L4-P01, política de retención y demás propuestas L5-P01–L5-P09 permanecen pendientes de aprobación específica.** No se cierra RISK-10 ni se fija policyVersion/plazo/backup/RPO/RTO/SLO/cuota/licencia por aprobar Markdown. Las condiciones aprobadas de barrera/revocación sin política y cleanup con política aprobada permanecen intactas. APROBADO documental no promueve afirmaciones PROPUESTA/POR_VERIFICAR. No se añaden elecciones reglamentarias aTK-N1-1.0.0, pruebas de producto o permisos de implementación/CI/proveedores/lote6.

## Adenda de hallazgos de auditoría — lote 6, 2026-10-26

[AUD-01§11](audits/documentation-readiness.md#11-hallazgos-p0p3) registra AUD6-F01(P2/NON_BLOCKING), precisión de perspectiva lógica para oráculoIA sin contradicción normativa, yAUD6-F02(P3/NON_BLOCKING), ejemplo /api/v1/game/... inconsistente con catálogo /games. Solo correcciones propuestas, con fuentes/escenario/impacto/owner/verificación; no se editan09.01/10.01. La discrepancia de revisores sobre F01 se conserva con juicio coordinado y evidencia en el informe.

No se demuestra P0/P1 que bloqueePH01. [Condiciones§12](audits/documentation-readiness.md#12-condiciones-pendientes):L4-P01/PH05,retención/cleanupPH06,licencias/publicación,restauración/release y capacidades/recursos afectados conservan sus gates. No se cierran esos riesgos ni se aprueban propuestas por este informe; NORM-01–03/TK-N1-1.0.0 no se reabren. Auditoría no es certificación de producto o permiso de código. Dictamen formal y gate humano pendientes de revisión de AUD.

## Adenda de cierre aprobado del lote 6 — 2026-10-26

El propietario aceptó la auditoría final y READY_WITH_SCOPED_CONDITIONS sobre HEAD 6052689ff9b15a19214267d7cc6a614d17f8fcdd; PR #6 integrado por e81f01a0e94c3037e3b922894202b6eee80534e8. [Recibo](../PLANS.md#cierre-lote-6). Gate documentalL6 satisfecho; menciones previas de revisión pendiente son historia.

**AUD6-F01, AUD6-F02 y las seis condiciones AUD6-C01–C06 permanecen pendientes** con sus propietarios, escenarios, verificaciones y fases originales. No se aplica la precisión de09.01 o la corrección de10.01; no se resuelveL4-P01/retención/licencias/restore/capacidades/permisos por aprobar AUD. La preparación dePH00/PH01 no es ejecución o autorización de código. REG/TK-N1-1.0.0/14ACC/PRE09 yspecs aprobadas no cambian.
