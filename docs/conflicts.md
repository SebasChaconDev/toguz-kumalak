# CONF — Conflictos, riesgos y validaciones pendientes

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: CONF · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner: C; responsables técnicos por caso.

## Estado y reglas de uso

No existían documentos de producto que contradijeran el maestro: el árbol de trabajo estaba vacío en ETAPA 0 y al iniciar lote 1. La fuente aprobada y el inventario se preservan. Esta lista distingue riesgo de factibilidad, detalle de diseño pendiente y discrepancia textual ya matizada por el propio anexo. No acredita incidentes de un producto existente.

Un `POR_VERIFICAR` nunca se transforma en hecho probado por tener una ADR o una página oficial. Una `PROPUESTA` necesita aprobación si altera alcance, arquitectura, parámetros aprobados o rutas autorizadas. No se resuelven aquí reglas oficiales ni se ejecuta el spike futuro.

Referencias: [registro](decision-register.md), [trazabilidad](traceability.md), [inventario](index.md), [gates](../PLANS.md).

## Procedimiento de resolución

1. Abrir un ID estable con fuente PRE exacta, formulación original y evidencia disponible.
2. Clasificar: contradicción de requisitos, riesgo técnico o detalle pendiente.
3. Registrar impacto, propietario, documentos consumidores y evidencia futura necesaria.
4. Investigar en modo permitido. Una prueba de producto exige autorización independiente; la aprobación documental no la permite.
5. Si el resultado contradice una decisión, presentar alternativas e impacto al propietario. No cambiar el registro ni la pila por preferencia del agente.
6. Conservar aprobación humana textual y fecha; añadir ADR/cambio de requisitos cuando corresponda y actualizar matriz/consumidores.
7. Separar resolución documental de validación técnica. Una decisión puede seguir APROBADO y su implementación POR_VERIFICAR.

## Riesgos abiertos

| Caso | Naturaleza | Owner | Cierre documental / técnico |
|---|---|---|---|
| RISK-01 | POR_VERIFICAR | M | Fuente normativa y corpus / motor futuro |
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

Fuentes: [PRE-03.A1](decision-register.md#pre-03-a1), [A3](decision-register.md#pre-03-a3), [E1](decision-register.md#pre-03-e1), más jerarquía inmutable del anexo. Falta confirmar autoridad, vigencia y cláusulas de siembra, atsyrau, tuzdyk, repetición y finales. Se localizó en ETAPA 0 un [PDF alojado por Mind Sports Olympiad](https://mindsportsolympiad.com/wp-content/uploads/2021/07/Rules-toguz-en-official.pdf); su edición y adecuación normativa permanecen pendientes, sin afirmar un cotejo integral.

Impacto: bloquear el cierre de los algoritmos afectados en [03.01](index.md#spec-03-01), 03.02–03.04 y escenarios IA/tutorial. Evidencia futura: referencia normativa identificable, cláusula→regla→posición→resultado esperado y discrepancias sometidas al propietario. No aplicar automáticamente otra variante del juego.

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
