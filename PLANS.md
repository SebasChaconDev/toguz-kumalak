# PLANS — Lotes documentales, gates y estado

> **Estado vigente — 2026-10-26: LOTE_5_EN_REVISIÓN.** El propietario autoriza la generación/publicación documental completa de20rutas del lote5 en `docs/specs-lote-05`, con PR sin fusión. [Plan y alcance posterior](#ejecucion-lote-5) prevalece sobre cortes anteriores. Lotes1–4 APROBADOS/integrados; L4-P01 sigue PROPUESTA. No producto, CI ejecutable, instalaciones, recursos/despliegues, pagos o lote6.

> **Estado vigente — 2026-10-26: LOTE_4_APROBADO_DOCUMENTALMENTE.** PR #4 integrado en main por `d5dccfbe0a55c78ddc7c2c0396ead9fd2bfad910`, desde el HEAD aprobado `3ef3f7cbfd3bfc4aea837d9f1b0b414be064e8ee`. [Recibo de aprobación e integración](PLANS.md#cierre-lote-4) prevalece sobre los cortes históricos. L4-P01 sigue PROPUESTA pendiente de aprobación de contratos antes de implementación. Detenerse: no producto, despliegues o lote 5.

> **Estado vigente — 2026-10-26: LOTE_4_EN_REVISIÓN.** El propietario declaró «AUTORIZO LA GENERACIÓN Y PUBLICACIÓN DOCUMENTAL DEL LOTE 4». Esta autorización posterior permite exactamente ocho specs, los cuatro registros de soporte y PR desde `docs/specs-lote-04` contra main, sin fusión. [Plan vigente](#ejecucion-lote-4) prevalece sobre los cortes históricos; producto, infraestructura y lote 5 siguen fuera de alcance.

> **Estado vigente — 2026-10-25: READY_FOR_LOTE_4_DOCUMENTATION.** PR #3 integrado en main mediante c50ba9f; GO_REMOTE experimental cerrado y recursos remotos propios retirados. [Recibo verificado](PLANS.md#lote-3-integrado-verificado) prevalece sobre los cortes históricos siguientes. El lote 4 aún no se ha generado; el producto no está implementado.

> Autorización posterior vigente: el propietario autoriza generar y publicar exclusivamente el lote 3 documental en docs/specs-lote-03: once specs EN_REVISIÓN, registros necesarios y PR contra main sin fusionar. Esta instrucción sustituye las paradas históricas de lotes anteriores; no autoriza producto, lote 4, infraestructura o pagos.

> Aprobación posterior del lote 2: el propietario declaró «APRUEBO LOTE 2» y autorizó fusionar PR #2 en main. Contenido aprobado: HEAD 598ac9a4fc6ce14cf512785fb1c4160fafb4eb9e; integrado por squash 1f401af8d76794981bbcfdf9a3aa1831cf2621f5 (2026-10-21T21:10:14Z). Este registro sustituye las menciones históricas a revisión pendiente o prohibición de esa fusión. No autoriza lote 3 ni implementación; el producto sigue sin verificar.

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: PLANS · Versión: 0.11 · Fecha: 2026-10-26 · Estado documental: EN_REVISIÓN · Owner: C.

## Autorización y estado actual

- Gate 0: recibido del propietario el 2026-10-21: «APRUEBO INVENTARIO DOCUMENTAL».
- Historial: la restricción inicial de lote 1 fue satisfecha; la instrucción posterior «APRUEBO LOTE 1» autoriza su integración condicionada y la generación/publicación documental del lote 2.
- Inventario autorizado: [81 Markdown](docs/index.md#inventario),40 specs y 41 documentos de soporte.
- Permiso actual posterior: «AUTORIZO LA GENERACIÓN Y PUBLICACIÓN DOCUMENTAL DEL LOTE 3». Once specs inventariadas, cinco registros necesarios y PR desde docs/specs-lote-03 contra main sin fusionar; sin producto ni lote 4. La aprobación e integración de PR #2 permanecen como base verificada.
- Estado de ejecución documental: `LOTE_3_EN_REVISIÓN`. Lotes 1 y 2 y TK-N1-1.0.0 APROBADOS; once specs del lote 3 redactadas para PR sin fusionar y gate humano pendiente.
- Lotes 1 y 2 APROBADOS e integrados; 59 decisiones con spec aprobada. Lote 3 cubre 66 decisiones adicionales EN_REVISIÓN; 121 conservan spec futura. Generar/publicar no aprueba el contenido.
- Estado del producto: implementación no iniciada, proveedores no conectados, pruebas de producto no ejecutadas, despliegues no realizados.

La aprobación de inventario no aprueba anticipadamente el contenido. Un reporte emitido por el agente no sustituye aceptación humana. La instrucción más reciente del propietario prevalece si restringe una continuación.

## Fuentes y estado de partida

Se leyó el maestro íntegro:569 líneas,60.815 bytes, Anexo A con 246 decisiones únicas. No se encontró registro separado en repositorio o junto al maestro. La copia [SRC](docs/sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) conserva SHA-256:

```text
28C24C498829FCFAD410F07AAA6BD036D5A33382294401E11E86D486E4ECBBEF
```

ETAPA 0 e inicio de lote 1: árbol de trabajo vacío salvo.git, rama inicial master sin commits, sin remotos configurados. No había AGENTS aplicables en raíz o superiores inspeccionados. No se renombra la rama ni se crea historial. La existencia actual de documentación no cambia el estado del producto.

Jerarquía y requisitos: [REG](docs/decision-register.md), [CONF](docs/conflicts.md), [11.01](docs/specs/11-orchestration/11.01-document-standard-and-change-control.md).

## Secuencia documental autorizable

| Lote | Archivos nuevos | Contenido | Gate de revisión |
|---|---:|---|---|
|1|31|Fuentes, registro, preliminares, ocho ADR, raíces, índice/matriz/conflictos y 11.01–11.02|APRUEBO LOTE 1|
|2|10|Dominios 01,02 y 03; arquitectura, producto y motor|APRUEBO LOTE 2|
|3|11|05,07 y 06; servicios, identidad, datos y multiplayer|APRUEBO LOTE 3|
|4|8|04 y 08; visual, IA y aprendizaje|APRUEBO LOTE 4|
|5|20|09,10,11.03, ocho planes de fase y tres documentos harness|APRUEBO LOTE 5|
|6|1|Auditoría documental integral y reporte de pendientes|APRUEBO LOTE 6|

Las frases son gates de aceptación/revisión definidos por el maestro. Antes de cualquier escritura siguiente se comprobará la autorización humana efectiva, incluidas restricciones posteriores. Esta sesión termina tras entregar lote 1.

Los lotes 2–6 podrán actualizar índices, matriz, conflictos y documentos raíz pertinentes cuando estén autorizados; no agregar rutas fuera del inventario sin aprobación. Los cambios semánticos de documentos ya aprobados exigen procedimiento de cambio, no solo un chequeo editorial.

## Dependencias y paralelismo

[IDX](docs/index.md#inventario) contiene el manifiesto y las dependencias individuales de 40 specs. La numeración clasifica dominios; no obliga a implementarlos en orden numérico.

- Fijar fuente/IDs y normas comunes antes de integrar preliminares y ADR.
- Lote 2: arquitectura/producto en paralelo; motor depende del cotejo normativo y de su contrato de estado.
- Lote 3: límites 05.01 → identidad 07.01/07.02 → HTTP 05.02 → transacciones 05.03 → eventos 05.04 → perfiles/Elo07.03/07.04 → salas/reconciliación/recuperación 06.
- Lote 4: visual y aprendizaje en archivos independientes; consolidar equivalencia de eventos y accesibilidad.
- Lote 5: calidad y operación en paralelo; hitos/harness se cierran tras consolidarlas.
- Un autor por archivo, revisores sin escritura concurrente. Coordinación integra nombres, enlaces y matriz.

La capacidad operativa propuesta es coordinador más hasta tres especialistas activos, sin convertir ese número en requisito de producto. Los roles pueden agruparse, y no se delegan pasos dependientes de un gate no recibido.

## Hitos futuros — solo resumen de planificación

Las rutas PH-00–PH-07 se redactarán en lote 5; aquí se conserva el resumen exigido para orientar la documentación. No se ejecutan ni se simula su cierre.

| Hito previsto | Dependencias | Resultado demostrable futuro / aceptación | Evidencia y riesgos | Autorización |
|---|---|---|---|---|
|[PH-00](docs/index.md#phase-00) Factibilidad|01.02,05,10.01 y decisiones fuente|Go/Gin, motor compartido, Neon, Ably, comunicación protegida y transacción confirmables en cuatro proyectos|Builds/despliegues/logs/consumos; runtime beta, módulos y coste POR_VERIFICAR|Independiente para spike y cualquier recurso externo; pago siempre expreso|
|[PH-01](docs/index.md#phase-01) Motor|PH-00 y dominio 03 contrastado|Motor puro, corpus normativo, ilegales sin mutación y conservación 162|Unitarias/propiedades/fuzz/simulación observadas; RISK-01|Aprobación de implementación del hito|
|[PH-02](docs/index.md#phase-02) Contratos, datos e invitados|PH-01;05 y 07.01/07.02|Comandos autorizados/idempotentes y persistencia atómica con invitado verificable|DB real, contratos, concurrencia y acceso cruzado; RISK-05/07/09|Aprobación de implementación; migraciones/entornos según alcance|
|[PH-03](docs/index.md#phase-03) Cliente accesible y 3 D|PH-02;04|Partida mediante HTML y mejora 3 D progresiva sobre mismos estados/eventos|E 2 E/teclado/lector/visual/móvil; no claim FPS sin medir|Aprobación de implementación y revisiones visuales incrementales|
|[PH-04](docs/index.md#phase-04) Multiplayer privado|PH-03;06|Dos participantes, reconexión 24 h, duplicados seguros y reconciliación sin Ably|E 2 E dos contextos y fallos; cuotas, sesiones y expiración|Aprobación de implementación e integración externa|
|[PH-05](docs/index.md#phase-05) IA, práctica y tutorial|PH-04;08|IA fácil/normal legal, cancelación, respaldo; undo/redo y tutorial ES/EN en HTML/3D|Corpus/torneos de prueba/recursos/E2E; fuerza no acreditada antes de medir|Aprobación de implementación del hito|
|[PH-06](docs/index.md#phase-06) Cuentas, Elo y estadísticas|PH-05;07|OIDC, perfiles, historial, ranked y agregados atómicos, exportación/borrado|Pruebas de identidad/privacidad/ranking/concurrencia; política legal pendiente|Aprobación de implementación y configuración de proveedor|
|[PH-07](docs/index.md#phase-07) Hardening y portafolio|PH-06;09/10|DoD con evidencias, restauración ensayada, observabilidad y demo aceptada|Resultados reproducibles, revisión, rollback; licencia y coste|Aprobación de release/publicación/despliegue según acción|

La secuencia es propuesta documental, sin fechas ni costes garantizados. Mantiene IA básica, clasificación y estadísticas avanzadas en el lanzamiento comprometido. PH-00 no autoriza desplegar por estar descrito.

<a id="informe-lote-1"></a>

## Informe del lote 1

### Archivos y alcance

Se crean únicamente las 31 rutas siguientes; no había archivos de trabajo preexistentes que modificar. SRC se copia íntegramente, sin editar el original. No se generan stubs de otros lotes.

| ID | Archivo creado | Owner funcional | Estado |
|---|---|---|---|
| AGENTS | [AGENTS.md](AGENTS.md) | C | Nuevo; EN_REVISIÓN |
| GOALS | [GOALS.md](GOALS.md) | C | Nuevo; EN_REVISIÓN |
| PLANS | [PLANS.md](PLANS.md) | C | Nuevo; EN_REVISIÓN |
| PROMPTS | [PROMPTS.md](PROMPTS.md) | C | Nuevo; EN_REVISIÓN |
| README | [README.md](README.md) | C | Nuevo; EN_REVISIÓN |
| SRC | [docs/sources/TOGUZ_PROMPT_MAESTRO_CODEX.md](docs/sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) | C | Nuevo; EN_REVISIÓN |
| REG | [docs/decision-register.md](docs/decision-register.md) | C | Nuevo; EN_REVISIÓN |
| CONF | [docs/conflicts.md](docs/conflicts.md) | C | Nuevo; EN_REVISIÓN |
| IDX | [docs/index.md](docs/index.md) | C | Nuevo; EN_REVISIÓN |
| TRZ | [docs/traceability.md](docs/traceability.md) | C | Nuevo; EN_REVISIÓN |
| PRE-01 | [docs/preliminares/PRE-01.md](docs/preliminares/PRE-01.md) | A | Nuevo; EN_REVISIÓN |
| PRE-02 | [docs/preliminares/PRE-02.md](docs/preliminares/PRE-02.md) | C | Nuevo; EN_REVISIÓN |
| PRE-03 | [docs/preliminares/PRE-03.md](docs/preliminares/PRE-03.md) | M | Nuevo; EN_REVISIÓN |
| PRE-04 | [docs/preliminares/PRE-04.md](docs/preliminares/PRE-04.md) | V | Nuevo; EN_REVISIÓN |
| PRE-05 | [docs/preliminares/PRE-05.md](docs/preliminares/PRE-05.md) | S | Nuevo; EN_REVISIÓN |
| PRE-06 | [docs/preliminares/PRE-06.md](docs/preliminares/PRE-06.md) | S | Nuevo; EN_REVISIÓN |
| PRE-07 | [docs/preliminares/PRE-07.md](docs/preliminares/PRE-07.md) | S | Nuevo; EN_REVISIÓN |
| PRE-08 | [docs/preliminares/PRE-08.md](docs/preliminares/PRE-08.md) | I | Nuevo; EN_REVISIÓN |
| PRE-09 | [docs/preliminares/PRE-09.md](docs/preliminares/PRE-09.md) | Q | Nuevo; EN_REVISIÓN |
| PRE-10 | [docs/preliminares/PRE-10.md](docs/preliminares/PRE-10.md) | A | Nuevo; EN_REVISIÓN |
| PRE-11 | [docs/preliminares/PRE-11.md](docs/preliminares/PRE-11.md) | C | Nuevo; EN_REVISIÓN |
| ADR-0001 | [docs/adr/ADR-0001-three-go-services-and-data-ownership.md](docs/adr/ADR-0001-three-go-services-and-data-ownership.md) | A | Nuevo; EN_REVISIÓN |
| ADR-0002 | [docs/adr/ADR-0002-single-versioned-go-engine.md](docs/adr/ADR-0002-single-versioned-go-engine.md) | M | Nuevo; EN_REVISIÓN |
| ADR-0003 | [docs/adr/ADR-0003-rest-contracts-and-versioned-events.md](docs/adr/ADR-0003-rest-contracts-and-versioned-events.md) | S | Nuevo; EN_REVISIÓN |
| ADR-0004 | [docs/adr/ADR-0004-postgres-authority-and-postcommit-ably.md](docs/adr/ADR-0004-postgres-authority-and-postcommit-ably.md) | S | Nuevo; EN_REVISIÓN |
| ADR-0005 | [docs/adr/ADR-0005-session-and-service-trust-boundaries.md](docs/adr/ADR-0005-session-and-service-trust-boundaries.md) | S | Nuevo; EN_REVISIÓN |
| ADR-0006 | [docs/adr/ADR-0006-vercel-hobby-and-zero-cost-guardrails.md](docs/adr/ADR-0006-vercel-hobby-and-zero-cost-guardrails.md) | A | Nuevo; EN_REVISIÓN |
| ADR-0007 | [docs/adr/ADR-0007-3d-and-accessible-html-view.md](docs/adr/ADR-0007-3d-and-accessible-html-view.md) | V | Nuevo; EN_REVISIÓN |
| ADR-0008 | [docs/adr/ADR-0008-ranked-results-and-elo-ownership.md](docs/adr/ADR-0008-ranked-results-and-elo-ownership.md) | S | Nuevo; EN_REVISIÓN |
| 11.01 | [docs/specs/11-orchestration/11.01-document-standard-and-change-control.md](docs/specs/11-orchestration/11.01-document-standard-and-change-control.md) | C | Nuevo; EN_REVISIÓN |
| 11.02 | [docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md](docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md) | C | Nuevo; EN_REVISIÓN |

Los autores efectivos fueron: coordinador para diez archivos raíz/fuente/registro/índice/matriz/conflictos/plan/prompts; especialista `preliminares` para once PRE; especialista `adr` para ocho ADR; especialista `specs_gobierno` para 11.01–11.02. Se delegó conforme al §9 del maestro y PRE-11.C1–C 4. La coordinación revisa las 21 entregas delegadas y conserva responsabilidad de integración. Los roles funcionales del índice no se confunden con personas o agentes permanentes del producto.

### Decisiones referenciadas y cobertura

- Fuente:246 decisiones APROBADO en Anexo A, con conteos 17/18/22/25/24/20/24/24/24/24/24.
- Registro: texto literal con IDs y destino canónico.
- Once preliminares: contexto y consecuencias, sin fingir cuestionarios originales.
- Ocho ADR: decisiones existentes, alternativas y riesgos; documentos EN_REVISIÓN.
- Matriz:246 rutas canónicas;17 decisiones tienen spec redactada EN_REVISIÓN (13 en 11.01 y 4 en 11.02);229 esperan sus specs de lotes 2–5.
- No se utiliza 246/246 para afirmar que las 40 specs o el producto están completos.

### Verificaciones efectivamente realizadas

Estado: `AUDITORÍA_DOCUMENTAL_REALIZADA`; contenido pendiente de aprobación humana. Se ejecutaron análisis de texto en memoria con PowerShell/.NET y revisión cruzada, sin instalar herramientas ni guardar verificadores ejecutables.

| Comprobación observada | Resultado |
|---|---|
| Alcance de archivos | 31 de 31 rutas autorizadas; cero faltantes, cero inesperadas; todas `.md` |
| Fuente archivada | 60.815 bytes y SHA-256 idéntico al original; original sin editar |
| Registro literal | 246 decisiones, mismo orden y texto que Anexo A; cero diferencias |
| Preliminares | 246 IDs propios enlazados, conteos PRE 17/18/22/25/24/20/24/24/24/24/24; cero faltantes |
| Matriz | 246 filas únicas; todos los destinos coinciden con la asignación canónica autorizada |
| Estado de cobertura | 17 filas SPEC_EN_REVISIÓN y 229 PENDIENTE_LOTE; ninguna implica software completo |
| Inventario | 81 entradas y 40 especificaciones; solo el subconjunto de lote 1 existe |
| Dependencias | 40 nodos de specs, sin diferencias respecto al inventario ni ciclos de redacción |
| Enlaces locales | 1.830 enlaces comprobados; cero archivos o anchors faltantes |
| Estructura Markdown | Fences equilibrados, anchors explícitos sin duplicados y tablas con columnas coherentes |
| Metadata | 30 documentos generados EN_REVISIÓN; SRC conserva literalmente el estado histórico de la fuente |
| Specs 11.01 y 11.02 | 17 apartados cada una; catálogo de verificación documental para sus 13 y 4 decisiones |
| Espacios finales | Cero hallazgos en documentos generados; SRC se conserva sin normalización |
| Git | 31 archivos nuevos no seguidos, rama inicial master sin commits; diff --check sin salida |

El análisis de enlaces excluye ejemplos en fences y código inline. Las referencias externas se identifican por procedencia; no se volvió a comprobar su contenido en lote 1 ni se ejecutó una integración. No se utilizó renderizador de Markdown/Mermaid ni se acreditan capturas visuales. Git diff --check no cubre archivos no seguidos: se complementó con lectura y comprobaciones directas de los nuevos Markdown. Ningún resultado corresponde a pruebas de producto.

Revisión cruzada real: el autor de preliminares revisó las ocho ADR; el autor de ADR revisó los once PRE; el autor de gobierno revisó los diez archivos integrados por coordinación. Coordinación comprobó alcance, texto canónico, matriz, grafo, enlaces y formato de los 31 archivos, y revisó los límites críticos de autoridad, credenciales y Elo. No quedaron hallazgos semánticos bloqueantes del lote documental.

Correcciones integradas: referencias exactas y espaciado en prosa; Go/Gin frente a Node/Fastify en PRE-01; atributos de cookies en PRE-07; PRE-10 aclara que umbrales activan revisión y que los disparadores de degradación siguen por especificar. Esta última integración fue asumida por coordinación tras confirmar el cese de escritura del autor, sin coedición concurrente. Ninguna corrección cambia el texto aprobado del registro.

### Dependencias, contradicciones y riesgos pendientes

No hay contradicción con documentación previa, porque no existía. [CONF](docs/conflicts.md) mantiene 14 riesgos abiertos y 3 matices de fuente: IDs de ejemplo, realismo estilizado y ampliaciones de undo/tablas. Ningún riesgo técnico se marca resuelto por la generación documental.

Las dependencias futuras figuran como entradas planificadas del índice; no se crean archivos para hacerlas parecer satisfechas. El reglamento formal y la factibilidad Go/Vercel condicionan el cierre de las specs afectadas y la implementación futura.

### Criterios pendientes y gate

Pendiente: revisión humana de los 31 documentos, aceptación o correcciones del lote 1 y resolución posterior de los riesgos por sus owners. Ninguna ADR o spec generada se marca APROBADO automáticamente. Faltan los 50 archivos de lotes 2–6, fuera de la autorización actual.

Tras la auditoría de integración se detiene el trabajo y se solicita **«APRUEBO LOTE 1»**. La autorización de implementación, instalaciones, servicios, pruebas de producto, cambios Git o despliegues es independiente.

**NO se ha implementado, instalado, desplegado ni modificado código.** El único resultado de esta sesión son Markdown del lote 1.

## Autorización posterior de publicación en GitHub

El 2026-10-21 el propietario autorizó crear el repositorio público `toguz-kumalak` en la cuenta personal autenticada, configurar Git local, crear `.gitignore`, publicar la base existente y presentar las specs disponibles por PR sin fusionarlo. Esto constituye autorización independiente de Git/publicación; no es «APRUEBO LOTE 1» ni autoriza nuevos lotes o implementación. El informe del lote 1 anterior conserva el estado observado al generarlo.

La sesión de GitHub CLI activa se verificó por API como cuenta personal `SebasChaconDev`, con alcance `repo`. La credencial anterior de otra cuenta estaba inválida e inactiva y no se utiliza. La identidad de los commits se configura solo en este repositorio usando el alias público y correo noreply de GitHub, sin publicar el correo personal configurado globalmente.

Distribución preparada: commit inicial de 29 Markdown y `.gitignore` en `main`; rama `docs/specs-lote-01` con las specs existentes 11.01 y 11.02 y enlaces relativos integrados. En la base, sus referencias apuntan a la rama de revisión para evitar destinos locales ausentes. No se generan las specs de infraestructura del lote 2 ni se usa un título que las dé por existentes. Las dos ramas se publican juntas cuando sea posible y el PR queda abierto, sin merge.

La copia íntegra del maestro está en `docs/sources/TOGUZ_PROMPT_MAESTRO_CODEX.md`. No existe un archivo separado llamado `TOGUZ_REGISTRO_DECISIONES_PRE01_PRE11.md`; `docs/decision-register.md` ya conserva literalmente sus 246 decisiones fuente con procedencia verificable, por lo que no se crea un duplicado nominal. No hay ZIP que publicar, bases de datos, secretos o artefactos compilados en el inventario actual. Los directorios futuros `docs/phases/` y `harness/` permanecen pendientes, sin stubs.

Antes de publicar se aplica detección local de patrones de secretos y nombres sensibles, contraste de archivos/rutas y validación de enlaces/IDs. Un resultado sin hallazgos indica ausencia de patrones detectados, no una garantía universal. No se imprimen valores candidatos. La verificación remota de ramas, árboles, commits y PR se realiza antes de declarar publicación lograda en el informe al propietario. La elección de licencia sigue pendiente; no se añade licencia.

## Cierre verificado del lote 1 y autorización del lote 2

El propietario declaró «APRUEBO LOTE 1» y autorizó integrar PR #1 si HEAD=e38cc84950858e072c1a3f06ffbf854cf7e544f5. Se comprobaron 27 archivos: dos specs y 25 correcciones de referencias exclusivamente, sin diferencias adicionales. GitHub informó MERGEABLE/CLEAN, cero checks fallidos, sin protección ni reglas obligatorias de main. El squash es 329b7c1059e8a74f191c256870734e8ae8259966; GitHub confirmó ambas specs en main. No se usó force push ni se eludieron protecciones.

Los registros de aprobación se actualizan sin cambiar las 246 formulaciones fuente. Se autoriza después la rama docs/specs-lote-02 desde main actualizado, exactamente las diez rutas del inventario, auditoría documental y PR sin fusionar. La revisión final exige «APRUEBO LOTE 2»; no hay autorización para implementación ni lote 3. Los informes anteriores son históricos y conservan lo observado en su fecha.

<a id="informe-lote-2"></a>

## Lote 2 — Matriz de trabajo e informe de ejecución documental

Estado de ejecución: LOTE_2_EN_CURSO. Se parte de main 624518495e2a81aa5928576f769820345df6cdc0, que registra aprobación de lote 1 después del squash 329b7c1059e8a74f191c256870734e8ae8259966. Rama propia: `docs/specs-lote-02`. El propietario autoriza exactamente diez nuevas specs Markdown y actualización necesaria de registros. No están autorizados lote 3, implementación, instalaciones, pruebas de producto ni servicios externos.

### Matriz de asignación y fuentes

| ID | Ruta autorizada | Owner / autor efectivo | Dependencias de redacción | Decisiones canónicas |
|---|---|---|---|---|
| [01.01](docs/index.md#spec-01-01) | `docs/specs/01-infrastructure/01.01-stack-monorepo.md` | A / lote2_infra | Registro y norma documental | [PRE-01.A1](docs/decision-register.md#pre-01-a1), [PRE-01.A2](docs/decision-register.md#pre-01-a2), [PRE-01.A3](docs/decision-register.md#pre-01-a3), [PRE-01.A4](docs/decision-register.md#pre-01-a4), [PRE-01.C3](docs/decision-register.md#pre-01-c3) |
| [01.02](docs/index.md#spec-01-02) | `docs/specs/01-infrastructure/01.02-platform-feasibility.md` | A / lote2_infra | 01.01 | [PRE-01.B2](docs/decision-register.md#pre-01-b2), [PRE-01.C1](docs/decision-register.md#pre-01-c1), [PRE-01.C4](docs/decision-register.md#pre-01-c4), [PRE-05.F4](docs/decision-register.md#pre-05-f4) |
| [01.03](docs/index.md#spec-01-03) | `docs/specs/01-infrastructure/01.03-cost-capacity-evolution.md` | A / lote2_infra | 01.02 | [PRE-01.C2](docs/decision-register.md#pre-01-c2), [PRE-01.D3](docs/decision-register.md#pre-01-d3) |
| [02.01](docs/index.md#spec-02-01) | `docs/specs/02-product/02.01-scope-and-modes.md` | C / lote2_producto | Registro y norma documental | [PRE-02.A1](docs/decision-register.md#pre-02-a1), [PRE-02.A2](docs/decision-register.md#pre-02-a2), [PRE-02.A3](docs/decision-register.md#pre-02-a3), [PRE-02.A4](docs/decision-register.md#pre-02-a4), [PRE-02.E3](docs/decision-register.md#pre-02-e3) |
| [02.02](docs/index.md#spec-02-02) | `docs/specs/02-product/02.02-local-flows-and-i18n.md` | C / lote2_producto | 02.01 | [PRE-02.B3](docs/decision-register.md#pre-02-b3), [PRE-02.D3](docs/decision-register.md#pre-02-d3), [PRE-05.C4](docs/decision-register.md#pre-05-c4) |
| [02.03](docs/index.md#spec-02-03) | `docs/specs/02-product/02.03-accounts-history-and-licensing.md` | C / lote2_producto | 02.01 | [PRE-01.D4](docs/decision-register.md#pre-01-d4), [PRE-02.B1](docs/decision-register.md#pre-02-b1), [PRE-02.B2](docs/decision-register.md#pre-02-b2), [PRE-02.D2](docs/decision-register.md#pre-02-d2) |
| [03.01](docs/index.md#spec-03-01) | `docs/specs/03-game-engine/03.01-normative-rules-and-corpus.md` | M / lote2_normativa | 02.01 | [PRE-03.A3](docs/decision-register.md#pre-03-a3), [PRE-03.C1](docs/decision-register.md#pre-03-c1), [PRE-03.E1](docs/decision-register.md#pre-03-e1) |
| [03.02](docs/index.md#spec-03-02) | `docs/specs/03-game-engine/03.02-state-indices-and-invariants.md` | M / lote2_normativa | 01.01, 03.01 | [PRE-03.A4](docs/decision-register.md#pre-03-a4), [PRE-03.B1](docs/decision-register.md#pre-03-b1), [PRE-03.B4](docs/decision-register.md#pre-03-b4), [PRE-03.E2](docs/decision-register.md#pre-03-e2) |
| [03.03](docs/index.md#spec-03-03) | `docs/specs/03-game-engine/03.03-moves-and-semantic-events.md` | M / lote2_normativa | 03.02 | [PRE-02.C1](docs/decision-register.md#pre-02-c1), [PRE-03.B2](docs/decision-register.md#pre-03-b2), [PRE-03.B3](docs/decision-register.md#pre-03-b3), [PRE-03.B5](docs/decision-register.md#pre-03-b5), [PRE-03.C2](docs/decision-register.md#pre-03-c2), [PRE-03.C3](docs/decision-register.md#pre-03-c3) |
| [03.04](docs/index.md#spec-03-04) | `docs/specs/03-game-engine/03.04-outcomes-repetition-and-replay.md` | M / lote2_normativa | 03.03 | [PRE-02.C4](docs/decision-register.md#pre-02-c4), [PRE-02.D1](docs/decision-register.md#pre-02-d1), [PRE-03.A1](docs/decision-register.md#pre-03-a1), [PRE-03.A2](docs/decision-register.md#pre-03-a2), [PRE-03.C4](docs/decision-register.md#pre-03-c4), [PRE-03.E3](docs/decision-register.md#pre-03-e3) |

Son **42 decisiones con spec canónica en lote 2**: PRE-01=11, PRE-02=13, PRE-03=16, PRE-05=2. Las referencias cruzadas adicionales no trasladan propiedad canónica. Las 17 decisiones de gobierno de lote 1 tienen specs APROBADAS; las 42 nuevas quedarán EN_REVISIÓN y las otras 187 mantendrán specs futuras pendientes. Las 246 formulaciones fuente no se alteran.

### Propiedad y secuencia

Arquitectura y producto se redactan en paralelo en rutas independientes. La investigación normativa es solo lectura al principio. El autor del motor recibe después 01.01 y 02.01 entregadas, fuentes reglamentarias y restricciones; redacta 03.01 → 03.02 → 03.03 → 03.04 en orden, sin que otro autor cambie esos contratos. Coordinación escribe únicamente registros de integración, audita y comunica revisiones antes de corregir archivos de autores activos.

Las propuestas de diseño de contratos no son nuevas reglas oficiales. Si falta evidencia normativa, el apartado afectado se marca BLOQUEADO_PARA_APROBACIÓN_NORMATIVA; el documento puede entregarse EN_REVISIÓN con la discrepancia abierta, sin ejecutar algoritmos ni reemplazar decisiones.

### Dependencias concretas del lote 2

Flecha = requisito de redacción hacia consumidor; registros y norma 11.01/11.02 se consultan transversalmente, sin introducir aristas inversas.

```mermaid
flowchart LR
  I1["01.01 Monorepo"] --> I2["01.02 Factibilidad"]
  I2 --> I3["01.03 Coste y capacidad"]
  P1["02.01 Alcance"] --> P2["02.02 Local e idiomas"]
  P1 --> P3["02.03 Cuentas e historial"]
  P1 --> M1["03.01 Norma y corpus"]
  I1 --> M2["03.02 Estado"]
  M1 --> M2
  M2 --> M3["03.03 Movimientos y eventos"]
  M3 --> M4["03.04 Finales y replay"]
```

La falta de una interpretación normativa no se resuelve mediante el orden de redacción: el consumidor conserva el bloqueo y no puede aprobar el algoritmo afectado.

### Estado de auditoría y gate

### Comprobaciones documentales observadas en entrega inicial 9d06677

| Control | Resultado |
|---|---|
| Alcance físico | 42 archivos: 41 Markdown y .gitignore existente; exactamente diez specs nuevas y cuatro registros actualizados respecto a main 6245184. Sin rutas de lotes 3–6. |
| Inventario | 81 entradas y 40 specs conservadas; diez rutas nuevas coinciden exactamente con lote 2. |
| Estructura y estado | Diez specs v0.1 EN_REVISIÓN, 17 apartados numerados 1–17 cada una; los dos documentos de gobierno aprobados conservan sus 17. |
| Fuente canónica | 246 decisiones únicas, texto y orden exactos entre SRC/REG; 246 filas de trazabilidad únicas y 246 destinos originales intactos. |
| Cobertura | 17 decisiones con spec APROBADA + 42 con spec EN_REVISIÓN + 187 futuras = 246; 42 VP canónicos presentes y conectados con AC existentes. |
| Dependencias | 40 nodos comparados con inventario original, sin aristas cambiadas ni ciclos. |
| Enlaces/formato | 2413 vínculos locales comprobados; cero enlaces/anclas ausentes, IDs explícitos duplicados, fences desequilibrados o filas con columnas incoherentes. |
| Ejemplos | Tres bloques JSON completos analizados como datos; 14 estados numéricos de siete pares DER con dimensiones/rangos y suma 162 antes/después. No se ejecutó motor ni se probó alcanzabilidad. |
| Integridad/seguridad | Copia SRC conserva SHA256 original; solo Markdown autorizado modificado/creado; escaneo de patrones no detectó credenciales, correos privados ni archivos sensibles. .gitignore permanece sin cambios. |
| Producto | No implementado, compilado, probado, instalado, conectado ni desplegado. POC-01–11 y evidencias técnicas permanecen PENDIENTE. |

Los controles se ejecutaron en memoria sobre texto y Git; no se escribieron verificadores en el repositorio. Los procedimientos de producto descritos en las specs no se ejecutaron. El detalle de cobertura se mantiene en [TRZ](docs/traceability.md), los bloqueos en [CONF](docs/conflicts.md) y el inventario en [IDX](docs/index.md).

### Contribuciones y revisión observadas

- `lote2_infra`: autor único de 01.01–01.03; revisión independiente de producto y motor, sin editar rutas ajenas.
- `lote2_producto`: autor único de 02.01–02.03; revisión independiente de arquitectura y cotejo de cuotas oficiales.
- `lote2_normativa`: investigación normativa inicialmente solo lectura; autor único de 03.01–03.04 tras recibir dependencias estables. Leyó íntegramente N1; documentó acceso parcial/indexado a N2/AUX sin fingir vigencia.
- Coordinación: cierre condicionado del lote 1, matriz y briefs, lectura y revisión de contratos/fuentes, integración de cuatro registros, armonización de versiones en 02.02 después del cese del autor, auditoría y publicación. No se coeditaron contratos con autores activos.

Hallazgos corregidos: pipes de enums dentro de tablas; POC-06 distinguió secretos de servidor y tokens cliente previstos, conservando HttpOnly/access en memoria y prohibición de tokens en URL/logs; FAC-01 dejó claro que Game/AI, sin imponerlo a Identity, importan el motor; 02.02 alineó nombres propuestos de versiones con 03.02; 03.04 completó 81–81/equal_score, matriz resultado/causa/ganador/source y AC-09 tras hallazgo coincidente del coordinador y revisor independiente. La revisión posterior verificó esa corrección conservando los bloqueos. No se modificó ninguna decisión fuente ni se cerraron los 14 riesgos técnicos.

### Fuentes y bloqueos de aprobación

Consulta 2026-10-21: N1 [PDF inglés de dos páginas](https://mindsportsolympiad.com/wp-content/uploads/2021/07/Rules-toguz-en-official.pdf), leído completo; N2 [Asociación, Astana 2013, artículo 375](https://9kumalak.com/index.php/zhyldy-zhospar/375-ojyn-a-idasy-men-erezheleri), texto recuperado por búsqueda, apertura directa fallida; [AUX-509](https://9kumalak.com/index.php/zhyldy-zhospar/509-to-yz-mala-a-idasy-men-erezhesi), solo indexación. No se autentican firma, edición vigente o traducción oficial por el mero alojamiento.

[RISK-01](docs/conflicts.md#risk-01) conserva NORM-01 edición/vigencia, NORM-02 atsyrau/recogida/prioridad terminal y NORM-03 equivalencia/ventana/actor de reclamación como **BLOQUEADO_PARA_APROBACIÓN_NORMATIVA**. No se demostró contradicción contra PRE-03.A1/A3. La comparación con implementación independiente sigue PENDIENTE. Las fichas sintéticas de 03.01 son derivaciones manuales con alcanzabilidad UNPROVEN.

Infraestructura cita [Vercel Go](https://vercel.com/docs/functions/runtimes/go), [Hobby](https://vercel.com/docs/plans/hobby), [monorepos](https://vercel.com/docs/monorepos), [Neon precios](https://neon.com/pricing)/[pooling](https://neon.com/docs/connect/connection-pooling), [Ably Free](https://ably.com/docs/platform/pricing/free)/[límites](https://ably.com/docs/platform/pricing/limits). Son documentos oficiales fechados, no mediciones de cuenta. Go beta/empaquetado, capacidad gratuita, regiones y consumo siguen POR_VERIFICAR; POC-01–11 no se ejecutó. Licencia, procedencia de recursos, jurisdicción/retención y política de replay compartido permanecen pendientes de decisión/diseño.

### Recomendación para revisión humana

Revisar primero los tres bloqueos normativos y el convenio de índices/contexto/versiones; después contratos de transición/eventos/replay y la matriz de factibilidad/coste; finalmente escenarios de producto y permisos de historial. Aprobar el lote documental no equivale a aceptar algoritmos bloqueados ni autoriza código, proveedores, PoC o lote 3. Una resolución normativa futura debe identificar fuente, alcance y aprobación; se conservan los PRE originales.



La entrega se publica mediante PR hacia main con título `docs: lote 2 — architecture, product and Go game engine specifications`, sin fusionarlo. Se solicitará «APRUEBO LOTE 2» y se detendrá el trabajo; ese gate sigue PENDIENTE de respuesta humana explícita.


<a id="correccion-normativa-lote-2"></a>

## Corrección normativa posterior a auditoría externa — 2026-10-21

El propietario declara **NO APRUEBO TODAVÍA EL LOTE 2**. Autoriza investigar/corregir documentos existentes y actualizar PR#2, conservando PRE-01–PRE-11, rama y gate. Base de esta corrección: `9d06677409df185559edcd028b367edca41652c4`; PR abierto, sin revisiones/comentarios públicos en GitHub al inspeccionarlo (la auditoría fue comunicada por el propietario). No se presume que ausencia de comentario público invalide su revisión externa.

### Trabajo y propiedad

Coordinación fue único escritor de siete rutas:03.01,03.02,03.03,03.04,CONF,TRZ y PLANS. `normativa_edicion` investigó fuentes primarias y revisó interpretación en modo de solo lectura; `referencia_independiente` investigó y examinó software/test publicados y después revisó su documentación, sin instalar, clonar, ejecutar motores/pruebas o modificar archivos. Los cuatro contratos se entregaron estables antes de las revisiones. No se crearon documentos nuevos, ni cambió inventario/REG/SRC ni las seis specs01/02.

### Resultado normativo y alternativas sometidas al propietario

**No se consiguió una edición de juego completa con aprobación/fecha/vigencia verificables.** N1 se conserva como referencia documental íntegra identificada, no edición federativa adoptada; N2 y AUX-509 como textos distintos con acceso parcial. [03.01§3](docs/specs/03-game-engine/03.01-normative-rules-and-corpus.md#3-decisiones-fuente-y-autoridad) registra fuentes, fechas, autoridades observadas y límites, distinguiendo Toguz Kumalak de Toguz Korgool/Bestemshe.

Se examinó el [reglamento federativo del torneo escolar Almaty2023](https://mankala.cz/download/togyzqumalaqjun2023.pdf), fecha impresa2023-07-01,§VII pp.3/8; remite a reglas de Federación Mundial sin identificar edición de juego. La [noticia gubernamental del torneo2023-11-13](https://www.gov.kz/memleket/entities/vko-bilim/press/news/details/652688) declara uso de reglas aprobadas por Comité, sin acto ni texto. Ni fecha de torneo ni membrete prueban vigencia de N1/N2.

La [matriz NM-01–10](docs/specs/03-game-engine/03.01-normative-rules-and-corpus.md#matriz-normativa) relaciona fuente/cláusula/interpretación/algoritmo previsto/caso y estado. Avances: reclamante es un participante según N1; dos textos de ofertas difieren1/2; adicional y remanentes aparecen en N2/AUX-509; artículo didáctico KazAST2023 contempla continuidad al repoblar rival, sin valor de norma adoptada. Siguen abiertos:
- NORM-01: edición/adopción y vigencia.
- NORM-02: frontera de agotamiento, adicional/origen, repoblación/continuidad, recogida y prioridad terminal.
- NORM-03: campos de igualdad, historia contada, inicial/consecutividad, habilitación por turno y ventana de reclamación.
- NORM-04: texto/política de número y momento de ofertas, sin incorporar límite nuevo al MVP.

**Solicitud de decisión humana concreta:** A) identificar/aportar edición completa y aclaraciones vinculantes; o B) aprobar una adaptación del producto explícita/versionada que resuelva cada punto, diferenciada de norma federativa. La opciónB no se autoriza por pedir correcciones. No se escoge una interpretación por resultados de un motor ni por mayoría de páginas. Los PRE no se alteran; gate APRUEBO LOTE 2 continúa pendiente.

### Corpus y referencia independiente

[REF-01](docs/specs/03-game-engine/03.01-normative-rules-and-corpus.md#referencia-independiente): Mind-Sports-Games/strategygames, commit`16e9e824624c54e3963053c13f2901d809f12d2a`, fecha2026-10-06T13:28:24Z, versión declarada`10.2.1-s3-ps15`, MIT. Selección/inspección textual realizadas; comparación ejecutada PENDIENTE. Variante9×2/162 y tuz, desactiva repetición; reparto anticipado al agotarse rival no valida movimiento adicional ni cierre normativo. No se incorporó código externo.

[Corpus corregido](docs/specs/03-game-engine/03.01-normative-rules-and-corpus.md#corpus-corregido):
- Inicial publicadoN1 y dos casos anteriores de una jugada reclasificados con ruta manual.
- PATH-REF-01: secuencia publicada de seis jugadas, siete estados completos0–6 derivados manualmente, puntuaciones y tuz contrastados con expectativas textuales del test; alcanzabilidad documental condicionada a reglas comunes, sin replay observado.
- PUB-TEACH-A7: acción/captura publicada en artículo didáctico, salida completa derivada; no se afirmó transcripción de un diagrama no inspeccionado.
- ATS-DRAIN/DRAW/REFEED: material completo y alternativas de siembra/recogida, sin salida reglamentaria aceptada.
- SYN-END-82/81 y SYN-AGREEMENT: evaluación y consentimiento; REP-2/3/TURN/REV/UNAUTH y OFFER-POLICY: ledgers sintéticos/negativos de contrato, sin ciclo legal inventado.
- Queda pendiente una secuencia legal verificada de triple repetición y oráculos de atsyrau adoptados. No se declara corpus exhaustivo.

03.02–03.04 preservan invariantes y separan perfil de reglas, ledger/ocurrencias/caché, candidato de atsyrau, error de regla no resuelta y ausencia de transición oficial parcial. No se restringe reclamante al jugador de turno por intuición. REF-01 marca tuz1; adaptador conceptual la convierte a cero lógico. Aprobación de contratos propuestos aún pendiente.

### Verificación de esta corrección

Controles documentales realizados en memoria, sin motor ni pruebas de producto:
- Siete rutas modificadas respecto a9d06677; cero altas/bajas. Árbol conserva42 archivos (41Markdown y .gitignore existente); inventario81/40 y diez specs de lote2 sin cambio.
- Diez specs con17 apartados1–17 cada una; cuatro03.* v0.2 y seis01/02 v0.1, todas EN_REVISIÓN. Las dos specs de gobierno conservan aprobación.
- 246 textos fuente/REG idénticos y en orden;246 filas únicas y destinos canónicos originales. Cobertura17/42/187 inalterada;42VP enlazan AC existentes.
- 2448 vínculos locales sin destinos/anclas ausentes; tablas/fences/IDs explícitos sin errores. Tres ejemplos JSON completos analizados como datos.
-31 estados numéricos completos de corpus suman162:14 anteriores y17 añadidos. Esto comprueba aritmética, no prueba todos alcanzables; rutaPATH se revisó paso a paso por separado.
- CopiaSRC conserva SHA256 original28C24C498829FCFAD410F07AAA6BD036D5A33382294401E11E86D486E4ECBBEF. Índice y246 asignaciones no cambian; grafo40 nodos sin ciclos.
- Detección de patrones de credenciales/archivos sensibles sin hallazgos; git diff --check sin errores. No se crearon herramientas ejecutables, fixtures de producto, dependencias ni recursos externos.

### Revisión y parada de la corrección previa, anterior a AD-NORM-001

Revisión independiente real, posterior a entrega estable:
- normativa_edicion contrastó03.01§§3/5/11/16 y03.04§§7/8/16: sin falsas resoluciones o autoridad inflada. Señaló localizador textual dependiente de lengua enN5; coordinación lo cambió por tercer párrafo narrativo. Reconfirmó AUX-CONT punto9 como didáctico.
- referencia_independiente revisó commit/versión/licencia/código y las seis transicionesPATH, sumas y adaptación de tuz; sin hallazgos bloqueantes. No examinó de nuevo procedencia académica, cubierta por revisión normativa.
- Coordinación integró las revisiones, corrigió orden/continuidad de tablasAC y comprobó contratos, trazabilidad, estados y fuente. Ambas revisiones son documentales; ninguna acredita adopción normativa o ejecución.

Actualizar el mismo PR#2 con commit descriptivo y comprobación remota de HEAD/archivos; conservar título, base main y estado abierto sin fusionar. Entregar reglas referenciadas, límites, casos y preguntas. El propietario no ha emitido «APRUEBO LOTE 2» y esta investigación no lo sustituye. No continuar a implementación o lote 3.


<a id="adopcion-tk-n1-100"></a>

## Adenda vigente — adopción definitiva TK-N1-1.0.0, 2026-10-21

La instrucción expresa del propietario adopta íntegramente `TOGUZ_RESOLUCION_NORMATIVA_V1.md`:103 líneas,12.274 bytes, SHA256 `9330B3116357225429BC85466BA48E9937C7D434AFE117AA633CA2A04A61DDE5`. Se leyó completa. [REG/AD-NORM-001](docs/decision-register.md#ad-norm-001) conserva transcripción íntegra y autoridad. El aviso anterior de adjunto ausente quedó superado al proporcionarse su ruta; no se había editado nada durante esa espera.

**Reglas completas y aprobadas:** TK-N1-1.0.0 define atsyrau, triple repetición, reclamación y finalización sin bloqueos normativos pendientes. N1 es fuente contractual única; las convenciones digitales de la resolución se atribuyen al propietario. No se comprobó ni se afirma vigencia federativa/certificación de torneo. Quedan sustituidas las solicitudes previas de edición/adaptación y comparaciones para decidir comportamiento.

### Alcance y ejecución

Base de la actualización: `62fe22328c5c9c4d00dfb294b3bc3efc15708f06` en docs/specs-lote-02, PR#2 abierto. Coordinación escribe exclusivamente ocho Markdown existentes: cuatro specs03.*,REG,CONF,TRZ y PLANS. No nuevas rutas, cambios de arquitectura/funcionalidades ni reescritura de246 decisiones. Las cuatro specs pasan a v0.3 EN_REVISIÓN; la adenda/reglas están APROBADAS y la verificación de producto sigue PENDIENTE.

NORM-01/02/03: **RESUELTO_POR_DECISIÓN_DE_PRODUCTO** para este perfil. NORM-04 de fuentes excluidas: NO_APLICA_AL_PERFIL; no se importan límites1/2 ni ventana de reloj. La historia de investigación se conserva identificada como anterior/sustituida en CONF y Git. No hay nuevas elecciones normativas para TK-N1-1.0.0. Un perfil futuro requiere otra autorización expresa y no se agrega a los lotes siguientes.

No se consultaron reglamentos/motores externos en esta actualización ni se reabrió la decisión. PRE-09 conserva pruebas automatizadas de integridad/corrección: unitarias, propiedades, fuzzing, simulaciones, historia/replay, integración y E2E pertinentes; no ejecutadas por esta misión documental. Los fallos de pruebas futuras se corrigen contra este perfil, sin cambiarlo silenciosamente.

### Decisiones y casos integrados

| Grupo | Aplicación |
|---|---|
| R-01–R-04 | Inicio/ciclo, piedra única/N−1, captura par rival, tuz automático permitido y depósitos al propietario |
| R-05 | Completar efectos; target≥82 → equal81–81 → turno rival → atsyrau → aparición activa; terminal prevalece sobre repetición |
| R-06 | Siguiente sin piedras, resultado por su kazán, sin adicional/barrido/transferencia de remanentes; terminal conserva162 |
| R-07 | Clave exacta reglas/board/kazanes/tuz/turno; inicial cuenta1; confirmados completos activos;≥3 habilita a ambos participantes hasta siguiente movimiento/cierre, revisión atómica sin penalización por rechazo |
| R-08 |162, pureza, Game autoridad, Go único, versiones y replay con perfil original |

Los14ACC están copiados literalmente en [03.01§14](docs/specs/03-game-engine/03.01-normative-rules-and-corpus.md#14-criterios-de-aceptación) y relacionados con consumidores en [TRZ](docs/traceability.md#adenda-tk-n1-100). Estados completos: siete pares DER existentes más ACC04-SECOND (tuz previo), ACC06 (siguiente con60 pierde), ACC07 (siguiente con81 empata/K80–81), ACC08 (actor vacío no pierde prematuramente), ACC09 (82 precede) y REP-P/variaciones para contrato. Se revisa162 y clasificación, no se afirma que sintéticos sean partidas alcanzables.

ACC14 se ejercita como contrato del orden/ramas con detector controlado: tras terminal no se consulta/concede repetición. No se fabrica ledger que presente como activo un estado ya terminal. Las secuencias legales/historia de integración y resultados observados son obligaciones futuras de calidad, no pendientes reglamentarias.

### Verificación de adopción y revisión

Comprobaciones documentales en memoria:

- 42 archivos físicos (41 Markdown y .gitignore existente), ocho modificaciones y cero altas/bajas respecto a62fe223; inventario81 rutas/40specs conservado.
- Diez specs de lote2 con17 secciones1–17; cuatro motor v0.3, seis infraestructura/producto v0.1, todas EN_REVISIÓN. REG v0.3 y CONF/TRZ/PLANS v0.5 integran adendas, sin autoaprobar el lote.
- 246 decisiones fuente/REG únicas, literales y ordenadas;246 destinos canónicos sin cambios, cobertura17/42/187 y42VP→AC.
- Ocho reglas R-01–08 y14 filas ACC copiadas literalmente del archivo completo. Transcripción íntegra de la resolución en REG comprobada como texto, distinguiendo normalización de saltos de su SHA256 original.
- 2423 vínculos locales, sin enlaces/anclas ausentes, IDs explícitos duplicados, fences abiertos o columnas incoherentes. Tres bloques JSON completos analizados como datos; inicial con count1.
- 25 estados numéricos completos con dimensiones9/9/2, cantidades no negativas y total162; revisión manual de transiciones/propietarios, sin ejecutar motor ni declarar alcanzabilidad de sintéticos.
- Copia SRC mantiene SHA256 original28C24C498829FCFAD410F07AAA6BD036D5A33382294401E11E86D486E4ECBBEF. Resolución original mantiene9330B3116357225429BC85466BA48E9937C7D434AFE117AA633CA2A04A61DDE5.
- Grafo de40 specs sin ciclos, sin cambiar dependencias. Búsqueda de bloqueos normativos activos en03.* sin remanentes; NORMATIVE_RULE_UNRESOLVED queda solo para perfil desconocido/corrupto.
- Escaneo de patrones sensibles sin hallazgos y git diff --check sin errores. No código, archivos ejecutables, instalación, pruebas de producto, conexiones de proveedores ni despliegue.

Revisión real posterior a entrega estable, limitada al archivo adoptado y al repositorio, sin web ni fuentes/motores externos:
- review_tk_n1 contrastó cuatro specs,8R/14ACC y semántica completa; no encontró elecciones normativas activas ni certificación falsa.
- review_tk_acc contrastó14ACC,25estados/tuzyks, JSON inicial y las transiciones especiales; confirmó distinción sintético/alcanzable, confirmación oficial y rollback.
- Ambos señalaron falta de separador Markdown en tablaACC y «depósito en origen» en ejemplo de9piedras. Coordinación añadió separador, corrigió a «una piedra conservada en origen» y explicitó que StoneSown solo llega al origen si el recorrido retorna. Los controles afectados se repitieron.
- Los revisores no escribieron ni ejecutaron producto. Coordinación conserva autoría única de las ocho rutas; decisión APROBADA y documentación EN_REVISIÓN se mantienen separadas.

### Estado final y parada

Publicar commit y actualizar el PR#2 existente, preservando título/base/rama y sin fusionar. Informar ocho archivos, ocho reglas/catorceACC y resultados documentales. **El perfil está completo y sin bloqueos normativos; el lote2 espera aceptación documental.** No considerar «APRUEBO LOTE 2» recibido por adoptar reglas. Detenerse al actualizar el PR; sin código, proveedores, pruebas de producto, instalación, lote3 o despliegue.

<a id="cierre-lote-2"></a>

## Cierre vigente — lote 2 aprobado e integrado, 2026-10-21

Instrucción expresa del propietario: «APRUEBO LOTE 2. Fusiona el Pull Request #2 en main, verifica la integración y actualiza los estados documentales correspondientes. Detente al terminar. No inicies el lote 3 ni implementes código hasta recibir nuevas instrucciones».

Se comprobó PR #2 OPEN/MERGEABLE/CLEAN en HEAD `598ac9a4fc6ce14cf512785fb1c4160fafb4eb9e`, cuenta personal SebasChaconDev, árbol limpio y auditoría documental sin errores. Se fusionó por squash con coincidencia obligatoria de ese HEAD, sin eludir protecciones. Resultado: [PR #2](https://github.com/SebasChaconDev/toguz-kumalak/pull/2) MERGED, commit `1f401af8d76794981bbcfdf9a3aa1831cf2621f5`, fecha `2026-10-21T21:10:14Z`. GitHub no tenía comprobaciones CI configuradas; no se atribuyen pruebas de producto.

La actualización local de main fue fast-forward. La comparación completa entre el árbol integrado y el HEAD aprobado devolvió cero diferencias antes de registrar esta aprobación. La aprobación se refiere al contenido de ese HEAD; las ediciones posteriores son exclusivamente estados, evidencia de integración y permisos de continuidad.

Alcance de cierre: 19 Markdown existentes, cero altas/bajas: las diez specs de lote 2, AGENTS, GOALS, PROMPTS, README, PLANS, IDX, REG, CONF y TRZ. Las specs pasan a APROBADO con revisión menor de metadatos, sin cambiar su cuerpo aprobado; las menciones a EN_REVISIÓN o gates pendientes dentro del texto conservado describen fases anteriores y quedan sustituidas por el recibo superior. Los registros actualizan versiones/estado y evidencia. Las 246 decisiones originales, las ocho reglas adoptadas, los catorce ACC, el corpus y PRE-09 permanecen intactos.

Cobertura actual: 59 decisiones con spec APROBADA, cero con spec EN_REVISIÓN y 187 con spec futura; las 246 conservan destino canónico. Inventario de 81 entradas/40 specs, 41 Markdown físicos y .gitignore, sin rutas de lote 3. Las verificaciones documentales y la comprobación remota del cierre se realizan antes de declarar terminado; no son pruebas de implementación.

Controles de cierre ejecutados: 42 archivos físicos sin altas/bajas; 246 textos fuente/REG literales y ordenados; 246 destinos canónicos sin cambios; cobertura 59/0/187; 2423 enlaces locales sin errores de rutas/anclas, tablas, IDs o fences; 17 secciones en cada spec. Comparación de los diez cuerpos aprobados sin diferencias salvo recibo, versión y estado; transcripción íntegra de la resolución intacta. DAG de 40 specs sin ciclos ni dependencias modificadas. Fuente maestra, preliminares, ADR y gobierno 11.* sin cambios; SHA256 de SRC conservado. Escaneo de patrones sensibles sin hallazgos y diff documental sin errores. La publicación del recibo se verifica después contra main remoto, conservando PR #2 en estado MERGED.

Este cierre prevalece sobre informes históricos de espera o prohibición de fusionar PR #2. TK-N1-1.0.0 sigue completo sin bloqueos normativos pendientes. Los riesgos técnicos y las pruebas futuras permanecen POR_VERIFICAR. **Trabajo detenido: ninguna autorización de lote 3, código, instalaciones, pruebas de producto o despliegues.**

<a id="lote-3"></a>

## Lote 3 — plan de ejecución documental autorizado, 2026-10-21

Fuente de autorización: solicitud íntegra adjunta «TOGUZ KUMALAK — EJECUCIÓN DOCUMENTAL DEL LOTE 3». El propietario autoriza once specs y cinco registros existentes (AGENTS mínimo, PLANS, IDX, TRZ, CONF), rama `docs/specs-lote-03`, commits/push/PR sin fusión. Base verificada `main=aaeeace21860305091c0351bedb0aca0df386816`; PR #2 MERGED por squash `1f401af8d76794981bbcfdf9a3aa1831cf2621f5`, HEAD aprobado `598ac9a4fc6ce14cf512785fb1c4160fafb4eb9e`. Remoto correcto SebasChaconDev/toguz-kumalak, fetch y fast-forward sin cambios, árbol limpio, rama de lote 3 inexistente antes de crearla. No se sobrescribió trabajo ajeno.

Este plan es exclusivamente documental. Se adapta el flujo de planificación a PLANS, ruta expresamente autorizada; no se generan planes auxiliares, archivos de producto o pruebas ejecutables.

- [x] Inspección inicial, PR #2, origen/base, permiso y once rutas exactas.
- [x] Consolidar fuentes y contratos en orden topológico: 05.01 → 07.01 → 07.02 → 05.02 → 05.03 → 05.04 → 07.03/07.04 → 06.01 → 06.02 → 06.03, respetando dependencias adicionales del índice.
- [x] Redactar once specs v0.1 EN_REVISIÓN, diecisiete secciones y 66 decisiones canónicas: 05.01=5, 05.02=4, 05.03=8, 05.04=4, 07.01=8, 07.02=6, 07.03=7, 07.04=6, 06.01=6, 06.02=6, 06.03=6.
- [x] Integrar índice, trazabilidad y riesgos/propuestas sin cambiar REG, ADR o specs aprobadas.
- [x] Auditar estructura, fuentes, contratos y casos negativos con revisión independiente de autores; corregir por dueño de cada archivo.
- [x] Verificar diff, fuentes/246 decisiones, 14 ACC, enlaces, DAG, alcance y secretos; publicar rama y PR con título prescrito.
- [x] Confirmar HEAD/archivos remotos, entregar informe EN_REVISIÓN, solicitar «APRUEBO LOTE 3» y detenerse.

Propiedad: coordinación es autora única de los cinco registros; los especialistas redactan specs en rutas exclusivas. Tras cada entrega y cese explícito, coordinación puede integrar correcciones en la spec estable y consolidar contratos compartidos, sin escritores simultáneos. No se redacta una spec dependiente hasta recibir su entrada consolidada. Revisión independiente solo lectura. Las asignaciones y entregas reales se registran en el informe final; el plan no acredita participación ejecutada.

Invariantes: tres servicios Go; Identity y Game con SQL separado; AI sin DB inicial y sin autoridad; motor puro TK-N1-1.0.0 y sus 14 ACC intactos; OpenAPI3.1/REST como contrato; PostgreSQL commit antes de Ably; sin exactly-once, outbox/Redis/colas obligatorios ni cambios reglamentarios. PRE-09 se conserva. Producto/pruebas/factibilidad de infraestructura permanecen POR_VERIFICAR.

Las decisiones de detalle no fijadas por el propietario se rotulan PROPUESTA con un contrato concreto para revisión, impactos y casos de validación. Una contradicción real se registra con alternativas y autorización requerida; no se resuelve cambiando una fuente aprobada.


<a id="informe-lote-3"></a>

## Informe de lote 3 — entrega para revisión documental

### Inventario, cobertura y dependencias

Base main verificada: aaeeace21860305091c0351bedb0aca0df386816. Lotes1/2 aprobados e integrados; PR2 MERGED y perfil TK-N1-1.0.0/AD-NORM-001 preservados. Nueva rama docs/specs-lote-03, sin sobrescribir trabajo previo. Se crean exactamente estas once rutas Markdown v0.1 EN_REVISIÓN:

| ID | Documento | Decisiones canónicas |
|---|---|---|
| 05.01 | [docs/specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md](docs/specs/05-backend/05.01-service-boundaries-and-hexagonal-design.md) | 5 |
| 05.02 | [docs/specs/05-backend/05.02-http-openapi-and-errors.md](docs/specs/05-backend/05.02-http-openapi-and-errors.md) | 4 |
| 05.03 | [docs/specs/05-backend/05.03-transactions-data-and-migrations.md](docs/specs/05-backend/05.03-transactions-data-and-migrations.md) | 8 |
| 05.04 | [docs/specs/05-backend/05.04-event-schemas-and-postcommit-publication.md](docs/specs/05-backend/05.04-event-schemas-and-postcommit-publication.md) | 4 |
| 06.01 | [docs/specs/06-multiplayer/06.01-rooms-invitations-and-ranked-lifecycle.md](docs/specs/06-multiplayer/06.01-rooms-invitations-and-ranked-lifecycle.md) | 6 |
| 06.02 | [docs/specs/06-multiplayer/06.02-ably-notifications-and-reconciliation.md](docs/specs/06-multiplayer/06.02-ably-notifications-and-reconciliation.md) | 6 |
| 06.03 | [docs/specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md](docs/specs/06-multiplayer/06.03-recovery-multitab-and-partial-failures.md) | 6 |
| 07.01 | [docs/specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md](docs/specs/07-identity-data/07.01-oidc-guests-and-session-lifecycle.md) | 8 |
| 07.02 | [docs/specs/07-identity-data/07.02-tokens-authorization-and-service-identity.md](docs/specs/07-identity-data/07.02-tokens-authorization-and-service-identity.md) | 6 |
| 07.03 | [docs/specs/07-identity-data/07.03-profiles-history-and-data-rights.md](docs/specs/07-identity-data/07.03-profiles-history-and-data-rights.md) | 7 |
| 07.04 | [docs/specs/07-identity-data/07.04-elo-leaderboard-and-statistics.md](docs/specs/07-identity-data/07.04-elo-leaderboard-and-statistics.md) | 6 |

Únicos existentes modificados: AGENTS.md (permiso operativo mínimo), PLANS.md (plan/informe), docs/index.md (once entregas/estados), docs/traceability.md (66 filas/consumidores/casos) y docs/conflicts.md (riesgos/propuestas). No cambios en README/GOALS/PROMPTS/REG/SRC/ADR/specs aprobadas/.gitignore. Total PR: once altas y cinco modificaciones, cero bajas, solo Markdown.

Cobertura: 66 PRE únicos en lote3 (PRE01=4, PRE02=1, PRE05=20, PRE06=18, PRE07=23), con VP→AC exactos en sus specs y evidencia de producto PENDIENTE. Repositorio:59 decisiones SPEC_APROBADO,66 SPEC_EN_REVISIÓN,121 futuras;246 con dueño único sin reasignación. [Matriz completa](docs/traceability.md#lote-3-cobertura), [18 casos críticos](docs/traceability.md#lote-3-casos-criticos), [14 ACC consumidores](docs/traceability.md#lote-3-acc-consumidores).

Grafo de40 specs sin ciclos y sin cambiar ninguna dependencia inventariada. Orden aplicado:

```mermaid
flowchart LR
  A["05.01"] --> B["07.01"] --> C["07.02"] --> D["05.02"]
  D --> E["05.03"] --> F["05.04"]
  C --> F
  F --> G["07.03"]
  C --> G
  E --> H["07.04"]
  C --> H
  F --> I["06.01"]
  H --> I
  I --> J["06.02"]
  F --> J
  J --> K["06.03"]
  B --> K
  C --> K
```

El diagrama muestra el subgrafo del lote; entradas aprobadas01.02/02.03/03.04 y las demás aristas se conservan en [IDX](docs/index.md#dependencias-de-redacción). 07.03/07.04 se redactaron tras consolidar sus entradas, antes del dominio06; no se inventó un ciclo por una referencia de consulta.

### Contratos y decisiones técnicas para revisión

- Solo Identity/Game/AI Go/Gin, SQL exclusivo por dueño, AI sin persistencia inicial. Mismo motor Go puro/versión resuelta en Game y AI; frontend/Ably nunca adjudican resultados.
- REST/OpenAPI3.1.1 propuesto:61 operaciones v1, RFC9457, revisiones uint64 como strings, límites por operación, domain-json/1/JCS y binding S2S de operación/método/body/path/query. Generadores/validadores y releases compatibles siguen POR_VERIFICAR.
- PostgreSQL/pgx5/sqlc/Goose/Neon previstos: guardas humanas UUIDASC→roomsASC→gamesASC→ratingsASC→derivados, relojes frescos tras locks, recibos durables antes de stale/turno en retries y final/ambos Elo/agregados atómicos. Commit incierto se resuelve con la misma intención; ningún HTTP durante SQL/locks ni garantía distribuida Identity/Game.
- Eventos semánticos versionados frente a Notice mínimo privado. Commit conocido precede publicación acotada; fallo/crash puede perder aviso. Sin exactly-once, Redis/colas/outbox obligatorios; reconciliación REST y replay persistido oficial. Canales/capacidades Ably message-subscribe exactos, sin permisos publish/presence/history del navegador.
- OIDC Google Code+PKCE y credenciales propias ES256/JWKS; usuario/S2S/Ably separados. Access900s, guest sliding48h/horizonte firmado, refresh prepare/commit con proof/CSRF persistentes y recibo cifrado; revocación/rotación y ventanas residuales explícitas. Vínculo G→U existente solo historial terminado congelado, sin activas/recibos G/Elo retroactivo.
- API pública batch por playerIds O aliases registrados exactos, correlación por consulta y sin SQL cruzado/barrido del ranking. Perfil mínimo, alias/avatar propuestos, grants de replay terminado para guest/user y exportaciones propias per_owner sin snapshot entre bases.
- Borrado prepare/confirm instala ticket y nonce antes de revocar. Primera confirmación exige sesión vigente/10min; seguimiento30d sirve solo a la misma raíz/fase. Barrier ACK_REVOKED separado de cleanup con política aprobada/ACK_CLEANUP_COMPLETE; rival/finales/Elo íntegros y sin forfeit por borrar. Política jurídica/retención RISK10 pendiente, sin certificación de cumplimiento.
- Elo1200/K32/diez provisionales intactos; fórmula binary64/redondeo único orientado UUID, enteros auditables y ambas variaciones en una tx. Admisión3 inicios por par/24h es PROPUESTA previa a Start, nunca veto retroactivo. Métricas de captura/tuz por owner sin doble conteo, ranking vivo y rebuild de una sola política.
- Salas privadas de dos, pending30min inactividad frente a recuperación activa24h por participante; colores aleatorios una vez y rankedEligible inmutable. Custodia cifrada separada de receipts, entrada por código sin resolver público y revancha nueva sin reabrir terminal. ReadyConsent guarda evidencia verificada de ambos sin JWT; fase explícita en GameView. Fábrica pura r0 se adapta solo en metadata al Start oficial, con count1. Rematch mantiene oldGameRevision y usa RematchView/TTL10min con polling15s máximo40 lecturas/10min visible. Detalles de sincronización/actividad/multipestaña se concretan en06.02–06.03.

Ningún detalle técnico nuevo se atribuye a las246 decisiones aprobadas. Los parámetros/formatos nuevos se distinguen como PROPUESTA y su comportamiento como POR_VERIFICAR. PRE09 conserva pruebas automatizadas de integridad/corrección previstas. TK-N1-1.0.0 y sus14ACC siguen completos, sin bloqueos normativos pendientes.

### Revisión independiente y correcciones integradas

Asignación real: l3_backend redactó05.*, l3_identity07.01–07.03, l3_data07.04 y l3_multiplayer06.*; coordinación leyó íntegramente las entregas y consolidó contratos tras cada cese. l3_review_arch revisó independientemente las seis primeras specs y07.03; l3_backend revisó07.04 completa y sus cálculos/contratos. Revisión06 y auditoría final: l3_review_arch leyó íntegramente06.01–06.03 y comprobó cada integración focal posterior. l3_backend contrastó IDX/TRZ reales contra las fuentes:81 entradas,40 specs,80 aristas,246 propietarios,18 casos críticos y14ACC consumidores. Un autor por archivo, revisores solo lectura; ningún informe de autor se aceptó como prueba suficiente.

Hallazgos corregidos en nuevas specs, sin cambios semánticos a fuentes aprobadas: retención de cookies refresh/CSRF antes de perder respuesta; cursor ranking vivo sin snapshot inventado; pasos cleanup separados para no fijar policy=null; política LogoutAll explícita sin step-up indefinido; admisión por código sin roomId previo; custodia de invitación fuera de receipt y reseal al cambiar expiry; seat estable separado de color aleatorio; expiry24h solo activa y tiempo después de locks; history_entries padre común y eventos asociados una sola vez; parser uint64 antes de coerción numeric; alias exacto por API; proof delete antes de revocación. Revisión focal posterior confirmó coherencia. Las inconsistencias de borradores se corrigieron antes de publicar; no se alteró un requisito aprobado ni se solicitó cambiar el perfil normativo. Precisiones finales corregidas: RoomView con fase/adjunto separado, ReadyConsent y revisión inicial oficial; cap de todos los GET/POST sololectura de pantalla, activityVersion obligatoria por UUID; ActivityDecision histórica frente a proyección actual, estados recoverable/expiry_pending/ended, allowedActions vacío mientras expiry_pending, cuotas de actividad y detección de regresión contra ambos participantes. En TRZ se corrigió competencia sobre la misma expectedRevision y se añadió F-SAL-03/11 para30min. Revisiones focales posteriores sin remanentes.

### Auditoría observada y límites

Controles documentales finales observados sobre el árbol integrado:

- 53 archivos físicos:52 Markdown y .gitignore existente; once altas y cinco modificaciones autorizadas, sin archivos vacíos/ajenos ni rutas físicas de lotes4–6. Once specs nuevas EN_REVISIÓN con secciones1–17;23 specs físicas conservan17 secciones.
- 246 decisiones SRC/REG literales y ordenadas;246 propietarios únicos iguales al mapa inicial. Cobertura59/66/121 y66 VP→99 AC en las nuevas specs, con175 casos futuros previstos; no tests de producto ejecutados. Inventario81 entradas/40 specs,80 aristas idénticas y DAG sin ciclos.
- Matrices de18 casos críticos y14ACC consumidores revisadas. R-01–R-08,14ACC, AD-NORM-001, PRE09, fuentes/ADR/preliminares/specs aprobadas sin cambios frente a base.
- 3.172 enlaces locales comprobados: cero rutas/anclas ausentes, IDs explícitos duplicados, tablas incoherentes o fences abiertos. Cero restos PENDIENTE_LOTE_3 en nuevas specs; futuros lotes4/5 conservan sus marcas.
- 13 bloques JSON de las once specs analizados sin errores. Dos schemas Markdown05.04: tres ejemplos positivos aceptados (incluido aviso06.02) y dos mutaciones negativas rechazadas con PowerShell7.6.5 Test-Json en memoria. No ejecutados generadores/runtime Go/TS ni demostradas todas las restricciones semánticas por ese chequeo.
- Seis ejemplos aritméticos W/D/L y ambos órdenes de finales concurrentes1199/1201 revisados; ejemplo temporal06.03 conserva86.400s por fila y mínimo. Cálculo documental no prueba una implementación o alcanzabilidad de trayectorias.
- SRC conserva SHA25628C24C498829FCFAD410F07AAA6BD036D5A33382294401E11E86D486E4ECBBEF igual al original. Resolución original conserva9330B3116357225429BC85466BA48E9937C7D434AFE117AA633CA2A04A61DDE5; REG íntegro sin cambios. Comprobación Git de rutas protegidas sin diferencias.
- Escaneo de53 archivos sin coincidencias de patrones sensibles/archivos de secretos; diff --check sin errores. No equivale a auditoría de seguridad del producto. Sin hooks activos o CI local de producto ejecutada. l3_backend reprodujo independientemente estos conteos, enlaces, esquemas y escaneo final desde los archivos reales; no encontró nuevos hallazgos.

Las únicas verificaciones acreditadas son documentales: lectura humana asistida, análisis de texto/JSON/esquemas y aritmética en memoria. No se generaron archivos Go/TS/SQL ni se compiló o ejecutó producto, pruebas unitarias/integración/E2E, migraciones o servicios; los ejemplos permanecen dentro de bloques Markdown. No se instalaron dependencias, realizaron conexiones operativas a Google/Neon/Ably/Vercel/Stripe, activaron pagos o desplegaron recursos. Git/GitHub se utilizan únicamente para publicación documental autorizada.

### Riesgos, aprobación y publicación

Persisten RISK02/03/04/05/06/07/08/09/10/11/12/13/14 en los alcances técnicos ya registrados; redactar contratos no demuestra runtime Go/Vercel, empaquetado, cuotas/capacidad, Neon/pools, SDK, cookies/proveedores, seguridad, rendimiento, fuerza IA, restauración o licencias. RISK10 exige jurisdicción/política/versiones/alcance de retención y mínimos deportivos antes de completar limpieza; no inventar anonimización. Las propuestas numéricas, alias, plazos, cadencias y admisión requieren revisión del lote. [CONF](docs/conflicts.md#lote-3-riesgos) conserva dueños/impacto/evidencia pendiente. No hay elección reglamentaria nueva para TK-N1-1.0.0.

Publicación verificada: [PR #3](https://github.com/SebasChaconDev/toguz-kumalak/pull/3), título exacto prescrito, base main y head docs/specs-lote-03, OPEN y mergedAt=null. Commit inicial de las especificaciones: `69c1b31b92f14f47454f2d24761a05f6dc3b50a3`, igual a HEAD remoto comprobado al crear el PR. GitHub confirmó16 rutas exactas: once altas y cinco modificaciones Markdown; ningún código o lote4. Cuenta personal CLI SebasChaconDev; conector inspeccionado ligado a otra cuenta, sin cambiar sus autenticaciones. Árbol local limpio tras el primer push. Main permanece aaeeace21860305091c0351bedb0aca0df386816 y PR2 conserva su integración.

El PR no reporta comprobaciones CI; no se atribuyen pruebas de producto. Este recibo y checklist se publican en un segundo commit exclusivamente de PLANS, sin cambiar specs o cobertura. El SHA anterior identifica la entrega de contenido; el HEAD vigente tras el recibo se verifica en GitHub y en la entrega final. No fusión, force push, implementación, proveedores operativos, pagos, despliegues o lote4.

Estado final: LOTE_3_EN_REVISIÓN, pendiente del gate humano **«APRUEBO LOTE 3»**. Publicación no equivale a aprobación. No fusionar, generar lote4, implementar, probar producto o modificar infraestructura. Trabajo detenido para revisión del propietario.
<a id="preparacion-lote-4"></a>

## Cierre autorizado de lote3 y preparación de lote4 — 2026-10-25

Instrucción posterior literal del propietario: «Autorizo los pasos y requerimientos que sean necesarios para poder empezar a desarrollar lote 4. Aplica lo necesario y cuando todo esté listo para lote 4 avisa».

Autoridad: esta instrucción responde a los requisitos expuestos para continuar, incluidos cerrar el spike y el lote3. Coordinación la aplica como autorización para completar su revisión/corrección documental e integración. **No se atribuye la frase literal «APRUEBO LOTE 3»**, que el propietario no escribió. La autorización posterior sustituye la parada anterior de no fusionar exclusivamente para este cierre. Las propuestas técnicas conservan su naturaleza PROPUESTA; aprobar la documentación no reescribe246 PRE ni acredita implementación.

Estado de esta entrega previa a integración: **LOTE_3_APROBADO_DOCUMENTALMENTE / READY_PENDING_PR3_MERGE**. Contenido revisado9db0f3e9d4bf1fb642697103d9232bcfaef42205; dos precisiones menores05.02 corregidas: JoinRoomByCode usa K01-000000000000 conforme a06.01 y cierre funcional exige F-HTTP-01–16. Once specs v0.2 APROBADO; cuerpos restantes conservados. Revisión independiente lot3_readiness_review leyó las once completas: sin Critical/Important y solo esas dos Minor.

Spike: [commit publicado504caa5](https://github.com/SebasChaconDev/toguz-kumalak/commit/504caa5d8d4f538e131e1beffc8843cb8938772a), [auditoría y límites](https://github.com/SebasChaconDev/toguz-kumalak/blob/504caa5d8d4f538e131e1beffc8843cb8938772a/spikes/infra-feasibility/REMOTE_AUDIT.md#remote-close-20261025). GO_LOCAL original intacto y GO_REMOTE experimental con revisión independiente, C6 medido y retirada final. Ocho builds acumulados,0USD adicional contratado, sin upgrades. SQL80 técnico previsto superado:117 operaciones reconstruidas/instrumentadas declaradas, sin fingir contador wire exacto. Cookie/roles/OIDC/transacciones/Ably reales no equivalen a sesiones, motor completo o IA del producto. El código experimental permanece en su rama; **no se integra en main**.

Se reconfirmó GitHub personal SebasChaconDev, PR3 abierto/mergeable en su HEAD esperado, cero workflows y hooks; Vercel quedó sin proyectos/deployments. Publicar la rama del spike no activó un proyecto Vercel vinculado identificado. No se hicieron escrituras sobre Neon9 Stones ni las dos apps Ably preexistentes.

Trazabilidad al cerrar:125 decisiones con SPEC_APROBADO,0 SPEC_EN_REVISIÓN y121 futuras;246 propietarios canónicos intactos. Lote3 conserva66 VP→99 AC y175 pruebas futuras. PRE-09, TK-N1-1.0.0,14ACC, fuente y ADR aprobados sin cambios.

### Siguiente trabajo preparado, todavía no iniciado

Inventario exacto de lote4: ocho specs y55 decisiones, con rutas conservadas en [IDX](docs/index.md#inventario):

- 04.01 navegación/accesibilidad;04.02 tablero/assets/cámara/materiales;04.03 interacción/animación/audio;04.04 rendimiento/alternativa HTML.
- 08.01 búsqueda/heurística/límites;08.02 orquestación/respaldo/métricas;08.03 práctica reversible/escenarios;08.04 tutorial/ayudas/replay educativo.

Orden del DAG:04.01→04.02/04.03→04.04;08.01→08.02 y08.03→08.04, conservando todas las entradas cruzadas inventariadas. No se crean archivos vacíos. Redacción deberá concretar presupuestos IA todavía propuestos, contexto de repetición en transposiciones, respaldo legal y práctica reversible sin autorizar escenarios del cliente como estado oficial.

Política jurídica/retención, generadores/SDK, restauración, seguridad completa, p95/rendimiento/fuerza IA y pruebas PRE-09 siguen pendientes antes de implementar lo afectado; no bloquean redactar04/08. El cierre prepara exclusivamente la **siguiente etapa documental**. No se empieza lote4 ni código del producto en este turno. Después de verificar merge/main y publicar el recibo final, informar y detenerse.

Comprobaciones de este cierre antes de publicar: 18 modificaciones Markdown y cero altas/código; 246 PRE literales y propietarios intactos;125/0/121;11 specs APROBADO con17 secciones y13 ejemplos JSON válidos. Revisor independiente reprodujo40 nodos/80 aristas sin ciclos,66 VP→99 AC/175 casos futuros y3.195 enlaces sin destinos/anclas ausentes. Fuente conserva28C24C498829FCFAD410F07AAA6BD036D5A33382294401E11E86D486E4ECBBEF; AD-NORM-001 intacta; diff --check limpio. Son comprobaciones documentales, no pruebas del producto.

<a id="lote-3-integrado-verificado"></a>

## Recibo final — lote3 integrado, lote4 preparado

Resultado vigente: **READY_FOR_LOTE_4_DOCUMENTATION**. El [PR #3](https://github.com/SebasChaconDev/toguz-kumalak/pull/3) está MERGED desde2026-10-26T02:19:22Z, con HEAD revisado fa714d7ba0be34df790a926667492b246e691a3d y squash c50ba9f58edd386100364cb396bf69f15d2d3a20. Se usó coincidencia obligatoria de HEAD, sin bypass ni force push. GitHub no tenía CI configurada; no se atribuyen pruebas de producto.

Main local se actualizó mediante fast-forward. Comparación íntegra del árbol entre HEAD revisado y squash: cero diferencias antes de este recibo. Main conserva52 Markdown y .gitignore, sin Go/TS/SQL/YAML, código experimental ni archivos de lote4. Las246 PRE y sus propietarios, AD-NORM-001,14ACC, PRE-09 y cuerpos aprobados anteriores permanecen intactos. Cobertura125/0/121; once nuevas specs APROBADAS documentalmente.

Gate3.5 cerrado como GO_REMOTE acotado y publicado en504caa5d8d4f538e131e1beffc8843cb8938772a, exclusivamente rama spike/infra-feasibility-01. Todos sus recursos remotos propios y deployments se retiraron; credenciales de aplicación revocadas y caché DPAPI temporal eliminada. Se conservaron recursos previos del propietario. El informe registra SQL80 superado, mediciones de una muestra y demás límites sin convertirlos en garantías productivas.

El gate prepara **ocho especificaciones documentales**04.01–04.04/08.01–08.04 y55 decisiones del inventario. No autoriza saltar a implementación del producto ni presenta propuestas técnicas como nuevas PRE aprobadas. No quedan requisitos documentales o hallazgos técnicos Critical/Important que impidan iniciar esa redacción cuando el propietario lo indique.

Trabajo detenido después de publicar/verificar este recibo. No se inició lote4, no se implementó producto, no hay recursos remotos del spike activos ni planes de pago contratados.

Verificación final del recibo:52 Markdown versionados,246 PRE literales,125 decisiones con spec aprobada y121 futuras;3115 enlaces a archivos locales comprobados sin destinos ausentes, fences cerrados y sin patrones de credenciales observados. Cero diferencias en specs/fuentes/preliminares/ADR respecto al squash revisado. Cero archivos04/08 o código de producto. diff --check limpio; las comprobaciones de anchors/DAG/JSON del cierre independiente siguen aplicando y el nuevo anchor del recibo está presente.

<a id="ejecucion-lote-4"></a>

## Ejecución autorizada del lote 4 — 2026-10-26

Autorización literal: «AUTORIZO LA GENERACIÓN Y PUBLICACIÓN DOCUMENTAL DEL LOTE 4». El propietario exige ocho specs completas, 55 decisiones con sus dueños originales, 17 secciones por spec, revisión especializada e independiente, actualización necesaria de IDX/TRZ/CONF/PLANS y publicación en rama `docs/specs-lote-04` con PR contra main sin fusionar. No solicita otra auditoría preliminar ni autorizaciones intermedias. La respuesta posterior sobre la excepción Vercel/hasta ocho builds corresponde al spike ya cerrado: no reactiva recursos o despliegues en esta entrega documental.

Base verificada al preparar la rama: main local/remoto `74e8cf164dfb86f27883de5e966795da52eeeaa5`, fetch y fast-forward sin cambios; árbol versionado limpio. Se conservan las ramas y cachés locales ignoradas preexistentes. Se trabaja en la rama solicitada dentro del checkout existente, sin instalar o probar producto. Autenticación GitHub personal activa verificada; no se cambia la cuenta inactiva ni se renuevan credenciales de proveedores.

### Restricciones globales y foco de revisión

- Solo ocho altas Markdown del inventario y cuatro registros de soporte; ningún código, manifiesto, asset, script o archivo de lote 5. Ejemplos dentro de Markdown.
- Preservar literalmente 246 PRE, propietarios canónicos, AD-NORM-001, TK-N1-1.0.0 y 14 ACC; tres servicios Go y un único motor Go, sin reglas frontend.
- Specs nuevas EN_REVISIÓN; requisitos fuente APROBADO, diseños nuevos PROPUESTA y resultados no medidos POR_VERIFICAR. Evidencia del spike exclusivamente experimental.
- Cálculo histórico/TT, IA en ramas de práctica, resultados tardíos, efecto único de Game, fallback legal y recuperación tras crash deben tener contratos y casos negativos explícitos.
- HTML plenamente jugable, accesibilidad ES/EN, foco/intención preservados al cambiar renderer, mapeo estable de pozos y eventos confirmados; sin atribuir FPS, WCAG, fuerza o coste productivo medidos.
- Propuestas de enlace IA/simulación no modifican silenciosamente 05.02/07.02: identificar el delta y revisión humana requerida antes de implementarlo. El lote puede especificar su solución propuesta sin cambiar archivos aprobados.

### Plan de tareas y propiedad

Plan documental adaptado a las rutas expresamente autorizadas; este registro es el ledger, sin crear briefs/reportes auxiliares en disco. La solicitud explícita de equipos y 11.02 permiten redacción paralela de dominios independientes, con un escritor por archivo.

| Tarea | Autor / rutas propias | Entradas y salida verificable |
|---|---|---|
| L4-V | l4_visual_author:04.01–04.04 | PRE-04 y dueños IDX, ADR, 02/03/05/06/07 y 11.01/11.02; redactar 04.01→04.02/04.03→04.04, 25 VP y aceptación positiva/negativa |
| L4-I | l4_ai_author:08.01–08.04 | PRE-08, PRE-03 y dueños IDX, perfil normativo, 02/03/05/06/07 y estándares; 08.01→08.02,08.03→08.04; 30 VP y contratos de búsqueda/práctica |
| L4-C | Coordinación:IDX/TRZ/CONF/PLANS | Integrar estados, 55 VP→AC, diez recorridos UI/Game/AI, riesgos y evidencia documental; no escribir archivos de autores activos |
| L4-R | l4_independent_review:solo lectura | Revisar fuentes y ocho specs, integridad, contratos, casos negativos y diff completo; correcciones devueltas a sus autores |

- [x] Preparación: Git limpio, referencias actualizadas y rama exacta creada.
- [x] L4-V: cuatro specs visuales entregadas y escritura cesada; revisión independiente en curso.
- [x] L4-I: cuatro specs IA/educación entregadas y escritura cesada.
- [x] L4-C: registros y contratos cruzados integrados.
- [x] L4-R: revisión independiente y correcciones verificadas.
- [x] Cierre: ocho altas/cuatro modificaciones autorizadas, 55 decisiones, 17 secciones, DAG/enlaces/JSON/fences, fuentes y lotes previos intactos.
- [x] Publicación: commit/push/PR contra main verificados, sin fusión; entrega detenida para «APRUEBO LOTE 4».

La aprobación del lote y la implementación siguen pendientes. No se inicia lote 5 ni se conectan o despliegan proveedores.

<a id="lote-4-contratos-cruzados"></a>

### Matriz de integración UI/Game/AI para revisión del lote 4

Los recorridos siguientes son contratos y aceptación futura, no pruebas ejecutadas. Estado oficial siempre Game/PostgreSQL, reglas siempre motor Go; AI propone y HTML/Three presentan. `/api/v1`, commandVersion, revisión decimal uint64, rulesVersion TK-N1-1.0.0 y versiones de estado/eventos/historial son ejes diferentes. El versionado local de práctica y su generación no reemplazan una revisión oficial.

| Flujo | Productor → contrato → consumidor / autoridad | Criterio positivo y negativo futuros / dueño |
|---|---|---|
| L4-X01 | SelecciónUI→CommandEnvelope/receipt/GameView/eventos→HTML/3D; Game confirma | Confirmación separada y animación de lote completo; preview/Notice/timeout no confirman movimiento. 04.01/04.03 consumen05.02/05.04/06.02 |
| L4-X02 | Commit humano Game→grant/S2S→AI→revalidación Game→UI | Un candidato legal genera un efecto interno único; duplicado/obsoleto/grant repetido/versiones discordantes no mueven. 08.02 consume05.02/05.03/07.02 |
| L4-X03 | Fallo/deadline AI→respaldo Go en Game→respuesta/estado | Respaldo legal barato tras revalidación; terminal/base vieja/resultado AI tardío no inventan otro movimiento ni derrota. Dueño08.02 |
| L4-X04 | Escenario cliente→SimulationRequest/Go→authoritative=false→explicaciónHTML/3D | Estado hipotético validado no modifica historia oficial/Elo; tablero cliente en SubmitMove se rechaza. IA de rama requiere delta L4-P01 explícito, no API existente ficticia. Dueños08.03/08.04 y05.02 |
| L4-X05 | Historial local→cursor/contexto/generación→simulación Go→presentación | Undo/redo restaura contexto íntegro; historia truncada o AI de generación anterior no habilita continuación falsa. Dueño08.03 con03.04 |
| L4-X06 | Game/replay→header/checkpoint/ordinal/eventos→reproductor | Reproducción compatible y sin efectos deportivos; hueco/versión/corrupción requieren fallo explícito o snapshot verificado, nunca animación inventada. Dueños03.04/05.04/07.03 y08.04 |
| L4-X07 | Réplica e intención común→rendererHTML al fallarWebGL | Mismo contrato/jugabilidad/foco y acciones; transición no confirma ni reenvía con ID nuevo una mutación incierta. Dueños04.01/04.04 |
| L4-X08 | Pérdida de conexión→REST/receipt/RecoveryView→reconciliador | Revisión autorizada vigente reconstruye escena; epoch viejo/Notice/reintento no revierte estado ni adjudica expiración. Dueños06.02/06.03 y04.01/04.03 |
| L4-X09 | Locale/cámara→presentación de PitRef estable | ES/EN y orientación preservan posición/turno/acciones; no cambian PositionKey o reglas. Dueños02.02/03.02 y04.01/04.02 |
| L4-X10 | Modo/permiso sobre recurso oficial Game→acciones UI y autorización backend | Ranked conserva selección legal/preview básico y rechaza solicitudes de ayuda ligadas a ese recurso aunque el cliente cambie propósito o invoque la ruta directamente. Práctica no altera Elo. Una posición copiada como escenario sintético libre no tiene procedencia demostrable: no se promete detectarla o impedir toda ayuda externa. Dueños05.02/07.02/07.04 y08.04 |

Además del DAG canónico,08.02/08.03 requieren revisión conjunta de binding IA/práctica;04.01/04.04 con08.03/08.04 revisan equivalencia de acceso. Son controles de compatibilidad, no aristas inversas añadidas al inventario. PRE-09 se aplica ahora como requisito previsto aunque09.* siga PENDIENTE_LOTE_5.

<a id="informe-lote-4"></a>

## Informe del lote 4 — entrega documental para revisión

Estado de entrega: LOTE_4_EN_REVISIÓN; contenido y revisión documental completados. La publicación se verifica en el recibo posterior; el gate humano permanece pendiente. Los cortes intermedios siguientes conservan sus verificaciones y correcciones, sin atribuir pruebas de producto. La entrega se limita a ocho altas de04/08 y cuatro registros IDX/TRZ/CONF/PLANS, con aceptación humana y fusión pendientes.

Integración previa a revisión: coordinación solicitó alinear nombres practiceSessionId/branchGeneration/requestId/commandId entre04/08. Detectó en borradores educativos una raíz de captura de163piedras usada como positivo y una defensa cuyo texto citaba pozos vacíos; se devolvieron al autorI para separar el negativo y cotejar raíz/acción/objetivo por aritmética documental contra03.01. Esto no modifica el corpus aprobado ni acredita ejecución Go. Los resultados finales y la revisión independiente se registran después del cese de ambos autores.

Revisión independiente L4-V: lectura íntegra04.01–04.04 contra fuentes, cumplimiento y calidad conformes sin bloqueantes. Único P3: ambigüedad de «total3D» en presupuesto04.04. Su autor corrigió solo04.04: modelos/texturas≤8MiB, con JS≤9.5MiB, audio opt-in≤1MiB aparte y shell separado; casoF-PERF-01/AC-04 alineados. Cese de escritura confirmado; comprobación delimitada de la corrección incluida en revisión final. Son presupuestos propuestos y no bytes medidos de un producto.

### Inventario entregado y decisiones

Ambos autores entregaron y cesaron escritura:04.01–04.04, ownerV,25VP/30AC/29casos;08.01–08.04, ownerI,30VP/33AC/34casos. Las ocho rutas exactas están en [IDX](docs/index.md#spec-04-01) y el desglose individual en [TRZ](docs/traceability.md#trz-lote-4): **55VP→63AC y63 casos futuros**, no tests ejecutados. Son ocho altas Markdown y cuatro registros modificados. Cada spec v0.1/2026-10-26/EN_REVISIÓN conserva17 secciones, metodología ordenada, contratos/errores/versiones/autoridad, justificación, ejemplos, riesgos y aceptación positiva/negativa.

Cambios propuestos destacados: shell React/HTML accesible y controlador común3D/HTML; PitRef y162instancias; cámara/materiales/procedencia; estados de confirmación/eventos y animación omisible; perfiles y presupuestos de medición sin afirmar rendimiento. IA Negamax/alfabeta/iterativa con TT solo de orden (sin scores/bounds), contexto completo por rama, fácil150ms/12.000nodos/profundidad4/16MiB y normal700ms/100.000/profundidad10/32MiB, siempre dentro del sobre07.02 y sin fuerza demostrada. Cancelación cooperativa, efecto interno único y respaldo legal Game; práctica versionada con undo/redo/IA opcional; tutorial completo/opcional, objetivos verificables y ayudas fuera de ranked, con HTML/ES/EN equivalentes.

L4-P01 identifica un delta propuesto de05.02/05.03/07.02: solicitud/consulta/cancelación de IA de rama, recurso/grant discriminado, persistencia mínima no deportiva, recovery=fallback y metadatos pedagógicos de simulación. commandId deduplica negocio, requestId correlaciona transporte y branchGeneration invalida respuesta anterior. Esos documentos aprobados no se modifican; antes de implementar el delta requiere revisión/aprobación humana de contratos. No se añade cuarto servicio, motor frontend, DB AI, Redis o worker.

Correcciones de integración de ejemplos completadas porI: capture-01 usa exactamente el inicial9×9 del DER-SOW-CAPTURE;163queda negativo separado. Defense-01 es una raíz sintética nueva de162: A=[1,1,0,0,0,0,0,0,0],B=[0,0,0,0,0,0,0,0,1],K=[79,80],sin tuz,turnA;A1 evita captura inmediata, mientrasA2 dejaA1capturable porB9yB alcanza82. Los esperados se cotejaron aritméticamente; no se afirma ejecución Go ni alcanzabilidad histórica de la raíz sintética. Fuente03.01 intacta.

### Comprobación estructural previa a la revisión final

Análisis en memoria sobre archivos reales:60Markdown,31specs;246 PRE y propietarios conservados;125 SPEC_APROBADO/55 SPEC_EN_REVISIÓN/66 futuras. DAG40nodos/80aristas idéntico al aprobado, sin ciclos. Ocho specs con17secciones y55VP exactos;63AC identificados.18bloquesJSON completos parseados (16previos y2nuevos), fences cerrados/IDs explícitos únicos y cero destinos/anclas locales ausentes. SHA256 del maestro28C24C498829FCFAD410F07AAA6BD036D5A33382294401E11E86D486E4ECBBEF. Las rutas cambiadas son exactamente las doce autorizadas; registro, fuente, preliminares, ADR y specs anteriores no tienen diferencias frente a74e8cf1. Revisión semántica final y publicación todavía pendientes en este corte.

### Revisión independiente y corrección final

Lectura independiente íntegra de ocho specs y diff de cuatro registros: sin P0/P1 y con dosP2 que exigían corrección de coherencia antes de entregar. L4-R01: el resultado de IA de rama no distinguía opponent_move de hint y podía prescribir adoptar una jugada al recibir una pista; se exige separar ambos efectos y probar que hint no envía move ni modifica timeline/progreso. L4-R02: la garantía de bloqueo ranked excedía lo verificable al admitir escenarios synthetic sin procedencia; se delimita a UI/contexto y recursos oficiales identificados, declarando que copiar un tablero como escenario libre no tiene detección antitrampa garantizada. No se añade prohibición global de práctica concurrente ni decisión deportiva nueva.

Onda final1: correcciones devueltas al autorI en08.02/08.03/08.04; coordinación ajusta solo su matrizL4-X10 y registroCONF. El revisor comprobará únicamente ambos hallazgos, sus propagaciones y regresiones introducidas. Las verificaciones independientes ya reprodujeron55VP/63AC/63casos,246dueños,40nodos80aristas,18JSON y1.514 enlaces/anclas de las doce rutas, sin ausentes; fuentes y specs previas intactas. El dictamen previo era REQUIERE_CORRECCIÓN, no aprobación del lote, y no se declara resuelto hasta comprobar el cambio.

### Dictamen documental final y límites

Onda final1 cerrada: L4-R01 ADDRESSED y L4-R02 ADDRESSED, comprobados directamente por l4_independent_review sobre el diff de cinco rutas,43inserciones/24supresiones. Hint solo presenta ayuda/preview explícita sin adoptar una jugada ni cambiar timeline/progreso; mover después requiere otra intención humana. Control ranked limitado a recursos oficiales y UI del modo, sin procedencia sintética o antitrampa universal fingidas. Propagación en contratos, casos y AC verificada; sin regresiones introducidas. P3 visual también corregido/verificado. **CUMPLIMIENTO DOCUMENTAL=CONFORME; CALIDAD DOCUMENTAL=CONFORME**, sin hallazgos pendientes de esa revisión. El propietario conserva la aceptación: ocho specs EN_REVISIÓN, nunca APROBADO por dictamen de un agente.

Comprobaciones coordinadas:55VP con88vínculos a63AC definidos y63casos futuros;74tablas nuevas con columnas coherentes;18JSON válidos;60Markdown/31specs;246decisiones/propietarios yDAG40/80 intactos,0ciclos;3.652 enlaces locales y2.606fragmentos comprobados sin ausentes, fences abiertos o IDs explícitos duplicados. Escaneo limitado de60Markdown contra claves privadas, tokens GitHub/JWT/Ably y DSN con contraseña: cero coincidencias; no certifica seguridad universal o historia completa. Diff y staged diff --check limpios después de normalizar LF solo en las cuatro altas visuales; no configuración Git cambiada, verificadores auxiliares o archivos no autorizados.

Pendientes de implementación/revisión de propuestas: deltaL4-P01 de contratos, matriz de releases/generadores/SDK, presupuestos150/700ms y memoria/descarga, dispositivos/60–30FPS/WCAG, fuerza de IA, políticas de datos/retención, licencias y revisión cultural. No son elecciones reglamentarias. PRE-09 y14ACC conservan íntegros sus controles futuros de corrección/integridad, propiedades, fuzzing/race, contratos, integración, E2E, accesibilidad y rendimiento; el spike experimental no los reemplaza.

No se implementó producto, generaron assets o manifiestos, instalaron dependencias, ejecutaron tests/builds/producto, iniciaron Docker, modificaron proveedores, crearon recursos o desplegaron servicios. El código experimental sigue exclusivamente en su rama. Solo resta publicar/verificar la entrega documental y detenerse para **«APRUEBO LOTE 4»**, sin fusión ni lote5.

<a id="lote-4-publicado"></a>

### Recibo de publicación — PR #4, 2026-10-26

Publicación verificada: [PR #4](https://github.com/SebasChaconDev/toguz-kumalak/pull/4), título «docs: lote 4 — experiencia visual, IA y aprendizaje», estado OPEN, mergedAt=null, base main y head docs/specs-lote-04. Commit de contenido revisado: `a36619e90937fe534c1b7974ff5f0287cb900d03`, idéntico en local/remoto/PR al crear y consultar la entrega. GitHub confirmó doce rutas exactas, ocho ADDED y cuatro MODIFIED; cero código o cambios a specs/fuentes aprobadas. Cuenta CLI personal SebasChaconDev verificada; sin force push o bypass.

Main permanece `74e8cf164dfb86f27883de5e966795da52eeeaa5`; no se fusionó ningún PR. El árbol local quedó limpio y alineado después del primer push. GitHub no reporta checks CI configurados; no se atribuyen pruebas de producto. Este recibo se publica mediante un segundo commit solo de PLANS, conservando intacto el contenido revisado de las ocho specs. Su HEAD final se verifica en GitHub después del push, sin intentar incrustar el hash de un commit dentro de sí mismo.

Resultado de la entrega: **LOTE_4_EN_REVISIÓN**, con55 decisiones documentadas y gate **«APRUEBO LOTE 4»** pendiente del propietario. Generar/publicar no aprueba el contenido ni autoriza fusionar, implementar, desplegar o iniciar lote5. Trabajo detenido al concluir la verificación del recibo.

<a id="cierre-lote-4"></a>

## Cierre aprobado e integrado del lote 4 — 2026-10-26

Estado vigente: **LOTE_4_APROBADO_DOCUMENTALMENTE**. Fuente expresa: «APRUEBO LOTE 4 como documentación técnica.» El propietario autorizó verificar HEAD del PR #4, cambios posteriores y comprobaciones obligatorias; si coincidía, fusionar y actualizar únicamente estados documentales, trazabilidad y registros necesarios, conservando L4-P01 pendiente de aprobación de contratos. Esta instrucción posterior sustituye la prohibición de fusionar el PR4 solo para ese cierre; no permite código, despliegues o lote5.

### Condiciones e integración verificadas

- HEAD exacto `3ef3f7cbfd3bfc4aea837d9f1b0b414be064e8ee` confirmado en local, rama remota y PR. Dos commits de entrega, sin posteriores sin revisar; árbol versionado limpio. Main previo `74e8cf164dfb86f27883de5e966795da52eeeaa5`.
- GitHub: PR OPEN/MERGEABLE/CLEAN antes de integrar; cero reglas activas de checks obligatorios, check-runs, statuses registrados y workflows. El agregado statuses devolvía pending con total_count=0: no se presenta como fallo o como ejecución CI aprobada. No existían comprobaciones obligatorias fallidas identificadas.
- Fusión squash con coincidencia de HEAD exigida, sin admin/bypass/force push: PR #4 MERGED `2026-10-26T16:48:00Z`, commit `d5dccfbe0a55c78ddc7c2c0396ead9fd2bfad910`.
- Fetch y fast-forward de main completados. Comparación íntegra del árbol entre HEAD aprobado y squash: cero diferencias antes de actualizar estados. La integración contiene exactamente la entrega revisada.

### Estados y alcance del registro de cierre

Ocho specs04.01–04.04/08.01–08.04 pasan a v0.2/APROBADO documentalmente con recibo de autoridad y precedencia de cortes históricos. Sus cuerpos técnicos, ejemplos, parámetros y naturalezas PROPUESTA/POR_VERIFICAR se conservan. No se modifican semánticamente05.02/05.03/07.02 ni otras specs anteriores. IDX y55 filas TRZ reflejan180decisiones aprobadas, cero en revisión y66 futuras; REG incorpora solo AD-GATE-004, sin reescribir246PRE o AD-NORM-001. AGENTS/README y registros muestran el estado real y la parada.

**L4-P01 sigue PROPUESTA técnica pendiente de aprobación de contratos antes de implementarla**. La aprobación como documentación técnica no aprueba por inferencia rutas, recursos/grants, schemas, persistencia o permisos del delta. Las cifras y demás alternativas técnicas conservan su naturaleza; PRE-09,14ACC, un único motor Go y tres servicios permanecen intactos. No existe permiso para implementar o iniciar lote5.

El commit posterior de estados se limita a15Markdown: ocho specs por metadatos/recibo y siete registros AGENTS/README/PLANS/IDX/TRZ/CONF/REG. No hay altas, código, manifiestos, assets, verificadores en disco, instalaciones, pruebas de producto, Docker, mutaciones de proveedores o despliegues. Su hash final y la coincidencia local/remoto se informan tras el push; el squash anterior identifica inequívocamente la fusión del contenido aprobado.

La verificación documental del cierre comprueba secciones, cobertura, dueños, DAG, enlaces/fences/JSON, preservación íntegra de cuerpos técnicos y fuentes y estado Git. Es una comprobación documental, no una prueba de producto. Concluida su publicación y verificación, detenerse y comunicar el commit de integración y el estado final; lote5 continúa sin iniciar.

Resultado observado de esos controles:60Markdown/31specs;246 PRE/propietarios y cuerpo íntegro del registro previo, incluidas AD-NORM-001 y sus14ACC, conservados. Cobertura180 SPEC_APROBADO/0 SPEC_EN_REVISIÓN/66 futuras;40nodos80aristas sin cambios/ciclos;18JSON válidos;3.670 enlaces locales y2.624fragmentos sin ausentes, IDs explícitos únicos y fences cerrados. En las ocho specs, comparación exacta de§§2–17 contra el HEAD aprobado: cero diferencias; en§1 solo versión0.1→0.2 y estadoEN_REVISIÓN→APROBADO, además del recibo superior. L4-P01 conserva íntegro su cuerpo PROPUESTA. Son15 modificaciones Markdown y cero altas; las specs de lotes anteriores, maestro, preliminares y ADR no cambian. Diff --check sin errores; la publicación de estados se verifica por HEAD local/remoto y PR MERGED antes de informar el cierre.

<a id="ejecucion-lote-5"></a>

## Ejecución documental del lote 5 — 2026-10-26

Fuente: solicitud íntegra del propietario «TOGUZ KUMALAK — PROMPT MAESTRO DEL LOTE 5», leída en sus1.096líneas. Autorización literal: «AUTORIZO LA GENERACIÓN Y PUBLICACIÓN DOCUMENTAL COMPLETA DEL LOTE 5». Permite20altas Markdown inventariadas, coordinación especializada, controles documentales, commits/push y PR contra main, sin fusión. Se actualizan IDX/TRZ/CONF/PLANS y, por coherencia operativa/enlaces futuros, mínimamente AGENTS/PROMPTS/README. No se archiva otro archivo de fuente ni se crean briefs, reportes, verificadores o ejecutables auxiliares en disco.

### Base, administración y automatizaciones

Base local/remota verificada `b91ba98a380ea12b8b6ffc24155edd437e958ed0`; repositorio personal SebasChaconDev/toguz-kumalak. PR1–4 fusionados; PR4 squash `d5dccfbe0a55c78ddc7c2c0396ead9fd2bfad910` y recibo de estados posterior b91ba98. Ocho specs04/08 ya APROBADO, autorización humana AD-GATE-004 y cuerpos intactos: no existe regularización pendiente que requiera reeditarlos. Las menciones EN_REVISIÓN del lote4 en cortes históricos no revocan el recibo. Cobertura inicial180aprobadas/0enrevisión/66futuras.

Git inicialmente limpio, rama nueva `docs/specs-lote-05`, sin trabajo previo de lote5 identificado ni otras worktrees. Se conservaron ramas y artefactos ignorados del spike. GitHub informa cero workflows, hooks o deployments registrados al preparar el trabajo. Antes del push se revalidará también la conexión de despliegue; un mecanismo que pueda publicar servicios sin autorización detiene la publicación documental antes de activarlo. No se renuevan credenciales o modifican proveedores para comprobarlo.

### Restricciones globales

- Exactamente20altas: nueve specs09/10/11.03, ochoPH y tresharness. Estados EN_REVISIÓN; solo nueve specs llevan las17secciones de11.01. No crear AUD-01 del lote6.
- Conservar246 PRE/texto/dueños, AD-NORM-001, TK-N1-1.0.0 y14ACC; tres serviciosGo, motor único, PostgreSQL autoritativo, avisosAbly postcommit y HTML/3D equivalente. No modificar specs/ADR/fuentes aprobadas.
- L4-P01 sigue PROPUESTA: aprobar el lote5 no aprueba sus contratos. El procedimiento de cambio debe preceder a editar05.02/05.03/07.02 o implementar consecuencias.
- FuentesAPROBADO, diseño nuevoPROPUESTA, comportamientoPOR_VERIFICAR. Evidencia observada lleva procedencia/commit/entorno y limitaciones, sin crear un estado documental nuevo. Spike504caa5 demuestra solo experimento acotado; no se repite automáticamente.
- Solo Markdown y ejemplos en fences; no producto, workflows/scripts/Compose reales, dependencias, tests/builds de producto, Docker, cuentas, secretos, proveedores operativos, pagos o despliegues.
- Revisión independiente: pruebas positivas/negativas/evidencia por ruta crítica, guardas/límites compartidos, coherenciaQA/operaciones/fases/harness y separación de tres permisos: documentación, implementación y recursos/despliegues.

### Tareas, propietarios y dependencias

Este plan es el ledger de continuidad. El usuario autoriza equipos paralelos por dominios independientes; un escritor activo por archivo y coordinación exclusiva de registros. La revisión se hace sin escribir en archivos de sus autores.

| Tarea | Escritura exclusiva | Entradas/salida y orden |
|---|---|---|
| L5-Q | l5_quality_author:09.01–09.04 | PRE09/decisionesIDX33; contratos aprobados01–08 y estándar11.01/02.09.01+09.02→09.03→09.04;33VP→AC, cobertura90/80 y matriz14ACC/riesgo/prueba/evidencia |
| L5-A | l5_operations_author:10.01–10.04 | PRE10/decisionesIDX26; contratos01–08 y evidencia experimental.10.01→10.02/10.03→10.04;26VP→AC, regiones/secretos/cuotas/backup/retención/releases yrunbooks |
| L5-C1 | Coordinación:11.03 y siete registros | Después de09.04+10.04;7VP→AC, gates por evidencia, procedimiento L4-P01, fuentes intactas y trazabilidad66decisiones |
| L5-C2 | l5_delivery_author:PH00–07 yH-BUILD/H-CTX/H-LOG | Preparación en lectura entregada; escribir solo después de11.03 estable y después de ochoPH completar tresharness. Planes futuros con tareas/contratos/positivos/negativos/evidencia y permisos; sin entradas ficticias de implementación |
| L5-R | l5_independent_review:solo lectura | Fuentes directas y entregas estables; coherencia de20archivos,66VP,pruebas previstas, contratos, permisos, DAG y diff; autores corrigen sus rutas |

Foco de integración: QA/CI debe ejercer pruebas09.01/02; severidad/DoD09.04 y releases10.04 deben coincidir; métricas10.02 sustentan umbrales; backup10.03 requiere restauración observada; PH consumen todas sus dependencias y no anticipan Elo/contratos; harness aplica11.03. Sin exactly-once de notificación, rollback de datos automático, antitrampa de escenarios de procedencia desconocida o pruebas aprobadas por declaración de agente.

- [x] Base y rama verificadas; lote4 ya regularizado, sin cambios a sus specs.
- [x] L5-Q: cuatro specs de calidad entregadas y escritura cesada; revisión independiente conforme tras L5-R01/R02.
- [x] L5-A: cuatro specs operativas entregadas y escritura cesada; revisión independiente conforme.
- [x] L5-C1:11.03 después de09.04/10.04 estables, revisión independiente conforme; nueve specs y66filas integradas.
- [x] L5-C2:ochoPH y tresharness entregados EN_REVISIÓN,65tickets/54casos y12campos porPH; correcciones R03/R04 aplicadas, revalidación independiente conforme tras R03/R04.
- [x] L5-R:20documentos y deltas de siete registros CONFORMES; L5-R01–R04 ADDRESSED.
- [x] Cierre documental:20altas, nueve×17secciones,66VP,180/66/0,246fuentes/dueños y DAG40/80 sin cambios; enlaces/tablas/fences/JSON y secretos.
- [x] Publicación: automatizaciones revalidadas, commit/push/PR #5 con título exacto; OPEN/sin merge. Gate «APRUEBO LOTE 5» pendiente; detenerse tras verificar recibo.

<a id="lote-5-integracion"></a>

### Matriz de integración de calidad, operación y ejecución futura

La matriz identifica pruebas y evidencias **por producir bajo autorización futura**, no resultados del producto. Los IDs L5-X son controles de revisión del lote, sin nuevos propietarios PRE o aristas inversas del DAG.

| Control | Requisito/contrato y productores → consumidores | Positivo / negativo y evidencia futura exigida |
|---|---|---|
| L5-X01 | 03.01–03.04/R01–08/14ACC →09.01→PH01/09.04 | Corpus normativo y secuencias reproducibles,162/pureza/repetición reclamada; sintético válido no acredita alcanzabilidad, versión/historia inválida se rechaza |
| L5-X02 | 05.02/05.03 recibos/TX →09.01→PH02/PH04/10.03 | Retry mismoID recupera efecto único; hash distinto/stale/rollback/commitunknown/carrera/reinicio observables en PostgreSQL real, sin inferir rollback de timeout |
| L5-X03 | 07.04/05.03 Elo/lock guards →09.01→PH06 | Dos partidas del mismo jugador, ambos órdenes de cierre y doble finalización; conservación/ledger/secuencias, no Elo de invitado/IA/práctica |
| L5-X04 | 07.01/07.02 sesión/cookies/JWKS →09.01/09.02→PH02/PH06 | Sesión válida y refresh/rotación; emisor/aud/firma/CSRF/UUID cruzados y dependencia caída rechazan sin privilegio alternativo |
| L5-X05 | 08.01/08.02/07.02 búsqueda/grant →09.01→PH05 | Legalidad, perspectiva/signo,TT soloorden,contexto por rama,deadline/nodos/memoria; callback incierto/replay/IA tardía/fallback mantienen un efecto Game |
| L5-X06 | 04.* /08.04 acceso →09.02→PH03/PH05/09.04 | HTML y3D conservan funciones/foco/ESEN/teclado/lector; WebGL ausente durante unknown no confirma ni duplica movimiento; evidencia manual además de herramientas |
| L5-X07 | 05.04/06.02/06.03 avisos/REST →09.01/09.02→PH04/10.02/10.04 | Commit precede aviso; pérdida/desorden/duplicado/Ablycaído converge a REST autorizado, sin exactly-once prometido o cierre por timer cliente |
| L5-X08 | 08.03 rama/versiones →09.01/09.02→PH05 | Undo/redo/restauración conservan contexto; generación descarta respuestas aunque revisión coincida, nueva jugada descarta futuro solo si se acepta |
| L5-X09 | 08.04 hints y modo →09.01/09.02→PH05 | Hint no mueve ni cambia progreso; Game bloquea ayuda ligada a recurso ranked real. No prometer detectar tablero copiado como escenario sin procedencia |
| L5-X10 | 09.01/09.02 →09.03→10.04/PH07/harness | Matriz de checks por trigger/consumidor; motor dispara Game/AI. PR externo sin secretos, check obligatorio omitido/infra no disponible no equivale a PASS |
| L5-X11 | 01.03/07.02/10.02 cuotas →09.04/10.04/PH07 | Contadores persistidos y pruebas de frontera;50/75/90 activan revisión, sin pago/degradación nueva automática o límite global basado en RAM |
| L5-X12 | 05.03/07.02/10.03 recuperación →09.01/09.04/PH06/PH07 | Restore aislado valida consistencia e integridad con tiempos medidos; backup presente no basta, no resucitar refresh/grant ni duplicar resultado/Elo tras perder recibos posteriores |
| L5-X13 | 09.04+10.04 →11.03→PH/harness | DoD/evidencia/revisión y tres permisos independientes; compilación, mensaje de agente o aprobación documental no autoriza release/recursos |
| L5-X14 | 07.03/RISK10 y02.03/RISK13 →10.03/11.03→PH06/PH07 | Política/derechos aprobados antes de acciones afectadas; sin plazo jurídico, anonimización irreversible o licencia inventados. Pendiente humano no bloquea escribir docs independientes |
| L5-X15 | 10.02 observación →09.04/11.03/harness | Evidencia liga requisito/versión/commit/entorno/comando/salida; logs minimizados/cardinalidad acotada. No métricas sin muestra, SLO/RPO/RTO/fuerza o coste garantizados |

<a id="informe-lote-5"></a>

## Informe del lote 5 — entrega para revisión

Estado LOTE_5_EN_REVISIÓN, sin aprobación humana. Los20documentos están entregados EN_REVISIÓN: nueve specs09/10/11.03, ochoPH y tresharness. Cobertura de decisiones180SPEC_APROBADO/66SPEC_EN_REVISIÓN/0sin spec redactada, con246dueños intactos. Los20documentos y deltas de los siete registros tienen revisión independiente CONFORME; L5-R01–R04 ADDRESSED. El cierre final y el recibo Git se añaden después de observarlos, sin atribuir pruebas de producto.

### Entregas y revisión técnica de las nueve specs

L5-A entregó10.01–10.04 con26VP/31AC/40familias de casos y cese de escritura; revisión independiente íntegra CONFORME, sin hallazgos accionables. L5-Q entregó09.01–09.04 con33VP/35AC/61casos, incluidos14ACCnormativos. Dos observaciones de revisión se devolvieron a su autor: L5-R01/P2 distinguía apariciones de repetición frente a historia de cierre; L5-R02/P3 aclaraba coordinación de incidentes. Correcciones limitadas a09.01§5/F-Q1-N13 y09.04§7, verificadas ADDRESSED por el revisor sin cambios a fuentes aprobadas: cierre conserva lifecycle/history_ordinal y no añade aparición, A coordina operación y C gobierno/escalamiento/dossier.

Coordinación redactó11.03 solo después de09.04/10.04 estables:7VP/10AC/12casos y17secciones, triple permiso, cambio contractual L4-P01, hitos12campos y gate por pruebas observadas/revisión/decisión humana. Revisión independiente de11.03 CONFORME, sin nuevos hallazgos. Total intermedio de nueve specs:66VP→76AC y113casos futuros; no tests ejecutados. Los contratos QA-Evidence/1, AcceptanceDossier/1 y ReleaseManifest/1 permanecen PROPUESTA.

Verificación coordinada intermedia en memoria:69Markdown/40specs, nueve×17secciones,66VP exactos del inventario,246PRE/propietarios,40nodos80aristas sin cambios/ciclos,18JSON completos válidos y3.982enlaces locales sin ausentes. Es un corte previo a PH/harness, no se declara todavía80Markdown o20altas. REG/maestro/PRE/ADR/specs aprobadas sin cambios; no se implementa o ejecuta producto. El revisor usó fuentes directas y artefactos del spike por Git, sin cuentas, instalaciones o mutaciones.

### Entrega parcial de fases y revisión de dependencias

C2 entregó ochoPH estables y después inició los tresharness. Revisión independiente:65tickets únicos (5/8/8/8/8/9/9/10),12campos Milestone por fase y EN_REVISIÓN. DosP2 a corregir por su autor: L5-R03 añadeT04 a la parte de recuperación oficial dePH05-T06, cuyo efecto/orquestación se produce enT04; L5-R04 separa enPH06-T07 barrera/revocación autorizada sin policyVersion y cleanup condicionado a política. El contrato aprobado05.03/07.03 permite ACK_REVOKED sin política; raíz+barrier quedan BLOCKED_PENDING_DATA_POLICY y cleanup null/cambiado no se acepta ni equivale a supresión completa. Estas correcciones no cambian fuentes, cantidad de tickets o permisos; sus negativos y propagación al harness se verificarán al recibir la entrega estable.

### Observación de automatizaciones antes de publicar

Ámbito consultado: repositorio personal SebasChaconDev/toguz-kumalak y sesión web preexistente UltraSéneca de Vercel. GitHub al preparar la rama:0workflows,0hooks,0deployments. Consulta de solo lectura a Projects Vercel,2026-10-26: sesión autenticada Hobby, búsqueda vacía y mensaje «Deploy your first project», sin proyectos en ese ámbito. No se identificó proyecto Git conectado al repositorio o mecanismo de despliegue activable por este push. No es un inventario de cuentas/organizaciones ajenas.

No se inició sesión nueva, renovó tokenCLI, creó recurso, modificó protección/automatización o leyó un secreto para esa comprobación. Sin conector Vercel disponible se usó el navegador ya autorizado. Revalidar referencias/repositorio y ausencia de automatizaciones identificadas antes del push; si aparece un mecanismo concreto que active un despliegue no cubierto, detener publicación antes de activarlo.

### Entrega completa e inventario final del lote 5

C2 entregó y cesó escritura: ochoPH y tresharness v0.1,2026-10-26,OwnerC,EN_REVISIÓN. Revisor leyó íntegramente los3harness y revalidó L5-R03/R04: ADDRESSED, propagados a H-BUILD/H-CTX, sin nuevos hallazgos. Los20documentos son CONFORMES documentalmente; este resultado no es aprobación del propietario ni ejecución de producto.

| Documento nuevo | PRE/VP del propietario | Estado |
|---|---:|---|
| [09.01](docs/specs/09-quality/09.01-engine-data-contract-and-ai-verification.md) |16| EN_REVISIÓN |
| [09.02](docs/specs/09-quality/09.02-ui-e2e-accessibility-and-visual-verification.md) |6| EN_REVISIÓN |
| [09.03](docs/specs/09-quality/09.03-ci-security-and-release-controls.md) |6| EN_REVISIÓN |
| [09.04](docs/specs/09-quality/09.04-acceptance-performance-and-definition-of-done.md) |5| EN_REVISIÓN |
| [10.01](docs/specs/10-operations/10.01-deployments-origins-regions-and-secrets.md) |7| EN_REVISIÓN |
| [10.02](docs/specs/10-operations/10.02-observability-quotas-and-rate-limits.md) |11| EN_REVISIÓN |
| [10.03](docs/specs/10-operations/10.03-backups-restoration-and-retention.md) |5| EN_REVISIÓN |
| [10.04](docs/specs/10-operations/10.04-releases-incidents-and-portability.md) |3| EN_REVISIÓN |
| [11.03](docs/specs/11-orchestration/11.03-milestones-gates-and-evidence.md) |7| EN_REVISIÓN |
| [PH-00](docs/phases/phase-00-feasibility-spike.md) |consume specs, sin dueñoPRE nuevo| EN_REVISIÓN |
| [PH-01](docs/phases/phase-01-go-engine.md) |consume specs, sin dueñoPRE nuevo| EN_REVISIÓN |
| [PH-02](docs/phases/phase-02-contracts-persistence-guests.md) |consume specs, sin dueñoPRE nuevo| EN_REVISIÓN |
| [PH-03](docs/phases/phase-03-accessible-client-and-3d.md) |consume specs, sin dueñoPRE nuevo| EN_REVISIÓN |
| [PH-04](docs/phases/phase-04-private-multiplayer.md) |consume specs, sin dueñoPRE nuevo| EN_REVISIÓN |
| [PH-05](docs/phases/phase-05-ai-practice-tutorial.md) |consume specs, sin dueñoPRE nuevo| EN_REVISIÓN |
| [PH-06](docs/phases/phase-06-accounts-elo-statistics.md) |consume specs, sin dueñoPRE nuevo| EN_REVISIÓN |
| [PH-07](docs/phases/phase-07-hardening-and-portfolio.md) |consume specs, sin dueñoPRE nuevo| EN_REVISIÓN |
| [H-BUILD](harness/build/README.md) |consume11.03/QA/PH| EN_REVISIÓN |
| [H-CTX](harness/context/README.md) |consume fuentes/contratos por tarea| EN_REVISIÓN |
| [H-LOG](harness/build-log.md) |plantilla y ejemplo SYNTHETIC/NOT_RUN| EN_REVISIÓN |

### Matriz de dependencias y ejecución futura

Se conserva el DAG canónico del inventario,40nodos/80aristas/0ciclos; las referencias transversales de verificación no añaden aristas inversas. Esta tabla reproduce las entradas directas de las nueve specs, sin nuevos propietarios:

| Consumidor | Prerrequisitos canónicos |
|---|---|
|09.01|05.04,07.04,08.02|
|09.02|04.04,06.03,08.04|
|09.03|09.01,09.02|
|09.04|09.03|
|10.01|01.03,05.04,07.02|
|10.02|10.01,06.03,08.02|
|10.03|10.01,05.03,07.03|
|10.04|10.02,10.03|
|11.03|09.04,10.04|

PH00→PH01→PH02→PH03→PH04→PH05→PH06→PH07 es secuencia de hitos de implementación futura, separada del DAG de redacción. Leer contratos de una fase posterior no presupone implementarlos; interfaces/dobles deben declararse. Los planes contienen65tickets y54familias de casos futuros, no119pruebas ejecutadas. Doce campos Milestone porPH; cada ticket tiene alcance, dependencias, resultado y aceptación.

| Fase | Tickets / casos | Objetivo y frontera |
|---|---:|---|
|PH-00|5 /5|Conservar GO_LOCAL/GO_REMOTE experimental504caa5 y sus límites; no repetir spike por este plan.|
|PH-01|8 /4|Motor Go puro único, corpus14ACC/propiedades/fuzz y cobertura90%; sin regla deportivafrontend.|
|PH-02|8 /7|Contratos, PostgreSQL real, idempotencia/atomicidad/guest/S2S; sin atribuir OIDC/Elo aún.|
|PH-03|8 /6|Cliente HTML jugable primero y3D progresivo, mismas funciones/ESEN/foco y autoridadGo.|
|PH-04|8 /7|Privado casual/invitados, avisosAbly y recuperaciónREST; ranked/Elo se completan enPH06.|
|PH-05|9 /8|Negamax/TT/contexto/cancelación, práctica/hints/tutorial; deltaL4-P01 exige aprobación contractual específica.|
|PH-06|9 /8|OIDC/cuentas, Elo concurrente/estadísticas y derechos; barrier/revocación sinpolicy no se aplaza, cleanup condicionado.|
|PH-07|10 /9|Endurecimiento/DoD/restore/costes/licencias/cultura; demo/release solo tras permisos/evidencia.|

H-BUILD selecciona ticket, verifica tres permisos y rutas, resuelve fuentes/versiones, planea positivos/negativos, ejecuta solo lo autorizado, recoge QA-Evidence y recibe revisión/decisión humana. H-CTX contiene nueve paquetes: motor, backend, Identity, multijugador, frontend, IA, seguridad, operaciones yQA; fuentes directas/hash/docCommit separados delcodeBase. H-LOG es plantilla con ejemplo SYNTHETIC/NOT_RUN: no inventa implementación, agente, comando o autorización observados.

### Calidad, seguridad, operación y límites de aceptación

QA consolida14ACC, propiedades/fuzz/race, motor90% y lógica de servicios80% más rutas críticas, PostgreSQL real con barreras/carreras/rollback/commitunknown, contratos/error/recibos, Elo concurrente, JWKS/CSRF/S2S/recurso, IA histórica y cancelación. UI añade Vitest/RTL/Playwright, tres navegadores, ES/EN/HTML/3D, visual/manual teclado/lector y dispositivos/medición pertinentes. Pruebas positivas/negativas y artefactos exigidos se relacionan en15recorridosL5-X y113familias futuras de las nueve specs.

QA-Evidence/1 liga PRE/spec/AC/caso, commit/versiones/entorno, datos/seed, comando sanitizado, expected/observed, PASS/FAIL/BLOCKED/NOT_RUN y artefactos/hash/revisión. Las66VP tienen79vínculos a76AC definidos. Un check requerido omitido o sin evidencia suficiente no pasa; P0/P1/críticos bloquean. P2 exige propietario/plazo/mitigación/aceptación humana explícita y acotada, sin dispensar PRE o volver verde una prueba fallida.

CI/DevSecOps es diseño futuro: gates por trigger/consumidor, motor compartido dispara Game/AI, PR externos sin secretos, análisis/vulnerabilidades/compatibilidad y release por dossier. Hoy no se instala CI ni se ejecuta código o herramienta del producto. Operación especifica cuatro proyectos/tresGo,orígenes/regiones/secretos/cuotas/telemetría, backups y restore aislado, releases/incidentes/portabilidad. Rollback de código no revierte datos ni resucita sesiones/grants/efectos; sin verdad posterior a backup se mantiene cuarentena. A coordina operación, C gobierno y Q verifica recuperación.

Coste objetivo0USD se conserva sin garantía medida del producto. No hay recursos, builds o pagos nuevos.50/75/90% son umbrales candidatos de revisión; no upgrade/degradación automática. Toolchains, cuotas/capacidad, dispositivos/rendimiento/IA, SLO/RPO/RTO, destino/retención de backups y política de datos/licencias/cultura siguen pendientes de aprobación o medición en acciones afectadas. L4-P01 permanece PROPUESTA, también tras aprobar este lote; TK-N1-1.0.0 no tiene nueva elección reglamentaria pendiente.

### Revisión independiente y correcciones

Autores efectivos: l5_quality_author(09),l5_operations_author(10),coordinación(11.03 y siete registros),l5_delivery_author(PH/harness). Todos entregaron/cerraron su escritura antes de revisión final. l5_independent_review no escribió archivos; leyó fuentes directas, nueve specs, ochoPH y tresharness, revisó deltas de registros y revalidó las correcciones de los autores:

| Hallazgo | Corrección y fuente | Resultado |
|---|---|---|
|L5-R01/P2|09.01§5/F-Q1-N13: no incrementar apariciones no significa omitir lifecycle/history_ordinal del cierre; conserva contrato03.04/05.03.|ADDRESSED|
|L5-R02/P3|09.04§7: A operación/contención/timeline,C gobierno/escalamiento/dossier,Q verificación,compatible10.04.|ADDRESSED|
|L5-R03/P2|PH05-T06 dependeT03+T04+T05; recuperación oficial consume orquestación/efectoT04.|ADDRESSED|
|L5-R04/P2|PH06-T07/C07 yH-BUILD/H-CTX: barrierS2S sinpolicyVersion yACK_REVOKED durable, raíz+barrierBLOCKED_PENDING_DATA_POLICY; cleanup aprobado/non-null/inmutable,ACK_REVOKED≠COMPLETE.|ADDRESSED|

Resultado de los20documentos y deltas finales de siete registros: CONFORME documentalmente, sin hallazgos restantes. Revisor verificó de nuevo80Markdown/40specs,65tickets/54casosPH,solo20altas+7registros y fuentes protegidas intactas; cese sin escrituras. La conformidad no sustituye «APRUEBO LOTE 5» o pruebas de producto.

Los siete registros se modifican por necesidad concreta: IDX rutas/estados; TRZ66filas y evidencia; CONF propuestas/correcciones; PLANS autorización/ledger/informe/recibo; AGENTS límite vigente; PROMPTS contexto/permisos; README cantidades/gate y enlace a planes/harness. REG,maestro,PRE,ADR y31specs aprobadas no tienen cambios; no se regulariza otra vez el lote4 ni se cambian55decisiones o cuerpos aprobados.

### Controles documentales de cierre — corte previo a publicación

Análisis de texto en memoria, sin instalar ni escribir verificadores:80Markdown/40specs;20altas(9+8+3) y7registros modificados; nueve×17secciones y EN_REVISIÓN,66VP exactos/79vínculos→76AC;113familias en specs y54enPH.246textos/propietarios intactos;180SPEC_APROBADO+66SPEC_EN_REVISIÓN+0sin spec; DAG40/80 idéntico/sin ciclos;65tickets sin dependencias internas ausentes/cíclicas;12camposMilestone porPH;116tablas de altas coherentes.18bloquesJSON completos válidos, fences eIDs explícitos comprobados. Conteos finales de enlaces, scan acotado de secretos y diff/staged --check se añaden tras observarlos con registros cerrados.

SHA256 del maestro conservado:28C24C498829FCFAD410F07AAA6BD036D5A33382294401E11E86D486E4ECBBEF. Los scripts ejecutados solo analizan texto/Git; no son pruebas de producto, benchmarks o certificación de seguridad. AUD-01 no existe: se mantiene PENDIENTE_LOTE_6. Sin código/manifiestos/workflows/assets nuevos, dependencias, producto, proveedores mutables o despliegues.

Revalidación GitHub26oct:main=b91ba98a380ea12b8b6ffc24155edd437e958ed0,0workflows/0hooks/0deployments; rama remota del lote5 aún inexistente, sin force push. Observación Vercel autenticada Hobby/0proyectos en el ámbito indicado arriba. No se identifica automatismo de despliegue activable. La publicación será documental y el PR seguirá OPEN/sin merge; recibo con URL/HEAD observado se añade después. Recomendación: aceptar como documentación una vez revisado el contenido del PR; implementación y recursos continúan bajo permisos separados.

Cierre previo a commit observado:4.431enlaces locales/2.853fragmentos sin destino o ancla ausente,18JSON válidos, fences cerrados eIDs explícitos únicos. Scan acotado de80Markdown contra cinco clases de secretos (claves privadas, tokensGitHub/JWT/Ably yDSN con contraseña): cero coincidencias, sin garantía sobre toda clase de secreto o historia ajena. Diff --check limpio. Se normalizaron exclusivamente11altasPH/harness deCRLF→LF; lectura Unicode antes/después idéntica11/11, sin cambios semánticos o fuentes aprobadas. El primer control de finales de línea falló y se repitió tras esa corrección: ceroCR en20altas. La revisión independiente de los últimos registros es CONFORME; staged --check y recibo Git se verificarán antes de declarar publicación.

Staged diff --check observado sin errores sobre27rutas exactas(20A/7M). No hay core.hooksPath configurado ni hooks locales activos; no se han desactivado controles. Solo se publicará docs/specs-lote-05 y un PR contra main, sin fusionarlo. La revisión humana «APRUEBO LOTE 5» permanece pendiente.

<a id="lote-5-publicacion"></a>

### Recibo de publicación documental — 2026-10-26

[PR #5](https://github.com/SebasChaconDev/toguz-kumalak/pull/5) creado y adjunto al chat. Título exacto: «docs: lote 5 — quality, operations, milestones and implementation plans». Base main, rama docs/specs-lote-05, estado OPEN y mergedAt=null. Commit de contenido `b6b7ee355291df2167836d77185fd906fdccd23a`, fecha2026-10-26T18:03:35Z; HEAD local/remoto/PR coincidían al verificar la creación. GitHub confirma27rutas exactas:20ADDED y7MODIFIED, soloMarkdown y ningún cambio protegido. Descripción con inventario/cobertura/grafo/QA/DevSecOps/operaciones/PH/harness/responsables/auditoría/propuestas/gate.

Main permanece `b91ba98a380ea12b8b6ffc24155edd437e958ed0`; no se fusionó ningún PR o commit del lote5. Árbol local limpio y upstream sincronizado tras el primer push. PR statusCheckRollup=[]: no hay checksCI configurados, no pruebas de producto aprobadas por este hecho. Después del push, GitHub informa cero deployments registrados. El recibo se publica en un segundo commit limitado aPLANS; su HEAD se comprueba contra origin/PR y se comunica al propietario, sin autorreferencia circular en este archivo.

Se conservan20documentos EN_REVISIÓN y revisión independiente CONFORME/L5-R01–R04 ADDRESSED. Aprobar documentación no aprueba L4-P01 ni permite implementar/instalarCI/recursos/demo/release. No se inician lote6 o producto. Estado final solicitado: **LOTE_5_EN_REVISIÓN**, listo para revisión humana y gate **APRUEBO LOTE 5**; detenerse después de verificar publicación del recibo.
