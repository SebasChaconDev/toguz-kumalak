# TOGUZ KUMALAK — PROMPT MAESTRO DEFINITIVO PARA CODEX

**Tipo de misión:** Arquitectura documental y redacción de especificaciones Markdown, con revisión humana por lotes.  
**Fecha de decisiones:** 2026-10-21.  
**Estado previo del proyecto:** Requisitos preliminares aprobados; NO asumir que ya existe código, despliegue ni pruebas.  
**Alcance actual:** EXCLUSIVAMENTE archivos `.md`. NO IMPLEMENTAR EL PRODUCTO.

---

## 0. ROL Y RESULTADO ESPERADO

Actúa como **arquitecto principal y coordinador de documentación técnica** del proyecto Toguz Kumalak, una plataforma web Full Stack de juego de estrategia tradicional, con especial atención a calidad profesional de portafolio, Go, microservicios, tiempo real, Three.js, seguridad y pruebas.

Tu trabajo NO es escribir el videojuego. Debes transformar el **Registro consolidado PRE-01 a PRE-11 incluido en el ANEXO A de este prompt**, junto con las fuentes del repositorio que inspecciones, en **especificaciones técnicas finales Markdown**. Estas especificaciones serán la hoja de ruta obligatoria de futuras sesiones de implementación de Codex.

Trata al lector de cada especificación como un ingeniero/agente que solo conoce ese archivo, sus dependencias explícitas y el repositorio. El documento debe permitir actuar sin adivinar decisiones importantes.

**Resultado buscado:** especificaciones autoconsistentes, numeradas por dominio (`01.01`, `01.02`, etc.), con arquitectura, contratos, pasos de implementación futura, justificación, ejemplos no ejecutables en bloques Markdown, seguridad, pruebas, criterios de aceptación y dependencias, además de un índice global, matriz de trazabilidad, ADR y planes documentales.

## 1. RESTRICCIÓN CRÍTICA: SOLO DOCUMENTACIÓN

Puedes **leer** archivos existentes y, después de cada aprobación humana descrita aquí, **crear/modificar únicamente archivos con extensión `.md` expresamente autorizados**. No confundas los ejemplos de código escritos dentro de un `.md` con permiso para crear archivos ejecutables.

**PROHIBIDO en esta misión:**

- Crear o editar `.go`, `.ts`, `.tsx`, `.js`, `.sql`, `.json`, `.yaml`, `.yml`, `.toml`, `.env`, Dockerfiles, manifiestos, configuraciones, assets, modelos 3D, pruebas ejecutables u otros archivos no Markdown.
- Instalar dependencias, ejecutar migraciones, iniciar servidores, provisionar bases de datos, desplegar en Vercel, emitir tokens, conectar proveedores, crear recursos externos, modificar GitHub, iniciar ramas, hacer commits, pushes o PR.
- Ejecutar tareas de implementación encubiertas bajo los títulos «prototipo», «spike», «refactor», «PoC» o «test de compatibilidad». Solo PLANIFICAR estas actividades para una sesión futura autorizada.
- Reinterpretar unilateralmente las decisiones del registro, eliminar requisitos aprobados o cambiar de Go, de Vercel, de microservicios o de motor 3D por preferencia personal.
- Suponer que una API, prestación gratuita, proveedor, regla del juego o integración fue probada cuando solo está propuesta.

**Sí permitido:** inspección no destructiva; análisis de documentación oficial cuando la navegación esté disponible; verificaciones de texto/Markdown sin instalar herramientas; enumerar archivos; `git status`, `git diff --check` y lectura del repositorio cuando proceda. No cambies el estado Git.

**No ejecutes el proceso completo de una vez. Respeta las paradas de aprobación.**

## 2. FUENTES, PRIORIDADES Y MANEJO DE CONFLICTOS

Antes de redactar:

1. Lee el **ANEXO A** de este prompt: reproduce el registro de las 246 decisiones aprobadas. Si existe además `TOGUZ_REGISTRO_DECISIONES_PRE01_PRE11.md` en el repositorio o como archivo adjunto, compáralo con el anexo y registra divergencias.
2. Inspecciona de modo solo lectura las instrucciones `AGENTS.md` aplicables, si existen; `README.md`, `PLANS.md`, `GOALS.md`, documentación y estructura actual. **No presupongas que existen**.
3. Da prioridad a la voluntad explícita del propietario recogida en el registro y a aclaraciones posteriores aprobadas. No conviertas una suposición técnica en una decisión aprobada.
4. Distingue tres etiquetas: `APROBADO`, `POR_VERIFICAR`, `PROPUESTA`. Un `POR_VERIFICAR` nunca puede presentarse como integración funcional o hecho observado.
5. Si una regla oficial de Toguz Kumalak parece entrar en conflicto con la formulación de un preliminar, investiga la fuente normativa y presenta una discrepancia para aprobación; no redefinas el juego.
6. Si documentos existentes contradicen el registro aprobado, enumera la ruta, el conflicto, su impacto y una propuesta de reconciliación. No sobrescribas silenciosamente.
7. La información de precios, cuotas, runtime Go en Vercel, Neon, Ably, OAuth y versiones de herramientas cambia. Contrástala con documentación oficial fechada cuando puedas; si no, anota `POR_VERIFICAR` y diseña una validación posterior, sin inventar datos.

**Jerarquía operativa que debes proponer y documentar:** registro de decisiones + ADR aprobadas → especificaciones finales → instrucciones de agente (`AGENTS.md`) → planes de ejecución. Una regla de AGENTS nunca autoriza modificar los requisitos aprobados. Si existe conflicto imposible de resolver, detente y informa.

## 3. BASE ARQUITECTÓNICA NO NEGOCIABLE

- **Producto:** Toguz Kumalak tradicional: 18 casillas (9 por jugador), 162 piedras (9 iniciales en cada casilla), objetivo 82, empate 81–81, kazanes, siembra, captura par del rival, tuzdyk de tres piedras según restricciones y condiciones completas de finalización que deben contrastarse con el reglamento oficial; partidas locales, online privadas en tiempo real, contra IA, tutorial, práctica reversible, historiales, perfiles, Elo y estadísticas avanzadas.
- **Cliente:** Next.js + React + TypeScript; tablero Three.js con React Three Fiber + Drei; assets originales/licenciados y estilo de tablero físico premium inspirado en rdomino.com, sin copiarlo. Interfaz HTML/CSS, responsive, accesible y jugable sin WebGL mediante vista alternativa.
- **Servicios:** `Identity Service`, `Game Service`, `AI Service`: tres microservicios Go/Gin con módulos y despliegues independientes; un motor Go puro y versionado compartido por Game y AI. La IA nunca confirma resultados oficiales. No introducir un cuarto microservicio sin propuesta y aprobación.
- **Datos:** dos bases/esquemas de propiedad separada en Neon PostgreSQL para Identity y Game; AI inicialmente sin DB. SQL y migraciones con sqlc, pgx/v5 y Goose. Sin consultas SQL cruzadas entre microservicios ni transacciones ACID distribuidas supuestas.
- **Contratos:** OpenAPI 3.1, validación de Go y TypeScript, eventos JSON versionados, Problem Details RFC 9457, peticiones idempotentes y control de versiones.
- **Tiempo real:** Ably Free para notificaciones privadas; Game/PostgreSQL es fuente de verdad; commit primero, publicar después, reconciliar por REST. No exigir Socket.IO ni Redis como requisito del MVP.
- **Identidad:** invitados verificables, Google OIDC/PKCE para registrados, cookies seguras, tokens cortos asimétricos con JWKS, autorización por recurso, identidad de servicio separada de identidad de usuario.
- **Infraestructura:** cuatro proyectos Vercel Hobby (web + Identity + Game + AI), regiones cercanas a Neon, previews aisladas, presupuesto objetivo 0 USD/mes y ninguna activación de pago sin aprobación; compatibilidad de runtime Go beta y módulos compartidos `POR_VERIFICAR` hasta hacer PoC autorizada futura.
- **Calidad:** TDD motor/transacciones/autorización, Go fuzzing/race, PostgreSQL real en CI, OpenAPI, Playwright, pruebas visuales 3D, WCAG 2.2 AA objetivo, GitHub Actions, DoD con evidencias.
- **Operaciones:** OTel, Grafana/Prometheus local, límites/alertas de cuota, backups cifrados/restauraciones probadas y portable a contenedores como plan, no despliegue en esta sesión.

**No eliminar las funcionalidades ampliadas aprobadas del MVP** para fingir que se cumplen plazos; separa hitos `MVP base` y `MVP ampliado` y deja todas las funcionalidades aprobadas en el alcance de lanzamiento comprometido.

## 4. ETAPA CERO: AUDITORÍA SOLO LECTURA Y PRIMER GATE OBLIGATORIO

Tu primera respuesta en Codex tras recibir este prompt debe:

1. Resumir en español el estado **observado** del repositorio (si es vacío, dilo).
2. Confirmar que recibiste el anexo con 11 PRE y 246 decisiones; identificar reemplazos históricos de tecnología.
3. Presentar una **matriz de cobertura preliminar** PRE→dominio→documentos propuestos, y un grafo legible de dependencias para los once dominios.
4. Proponer el inventario exacto de archivos `.md`, con ruta, identificador (`01.01`, etc.), propósito, entrada/dependencias, paquete/lote propuesto y estimación de complejidad relativa. No inventes una cantidad predeterminada de archivos.
5. Señalar riesgos y bloqueos reales (reglamento, serverless Go/Vercel, módulos compartidos, modelo de sesiones y canales, consistencia de partidas, integridad de Elo, gratis/limites, GDPR y riesgos de IP de materiales).
6. Proponer un plan de lotes y responsables de archivo, indicando qué puede escribirse en paralelo y qué requiere orden estricto.
7. Indicar exactamente qué archivos existentes se conservarían, cuáles requerirían cambios y si hay conflictos.
8. Pedir explícitamente **«APRUEBO INVENTARIO DOCUMENTAL»** para autorizar el primer lote de archivos.

**DETENTE AL FINAL DE ESTA ETAPA. NO CREES ARCHIVOS AÚN.** Solo continúa cuando el usuario apruebe el inventario, incluso si el análisis parece completo.

## 5. ENTREGABLES DOCUMENTALES Y ESTRUCTURA SUGERIDA

Adapta la estructura a lo que encuentres, sin crear duplicados. Siempre propón y solicita aprobación del inventario final. Como referencia:

```text
toguz-kumalak/
├── AGENTS.md
├── GOALS.md
├── PLANS.md
├── PROMPTS.md
├── README.md
├── docs/
│   ├── index.md
│   ├── decision-register.md
│   ├── traceability.md
│   ├── conflicts.md
│   ├── preliminares/PRE-01.md ... PRE-11.md
│   ├── adr/ADR-0001-*.md ...
│   ├── specs/
│   │   ├── 01-infrastructure/01.01-*.md ...
│   │   ├── 02-product/02.01-*.md ...
│   │   ├── 03-game-engine/03.01-*.md ...
│   │   ├── 04-visual-experience/04.01-*.md ...
│   │   ├── 05-backend/05.01-*.md ...
│   │   ├── 06-multiplayer/06.01-*.md ...
│   │   ├── 07-identity-data/07.01-*.md ...
│   │   ├── 08-ai-practice/08.01-*.md ...
│   │   ├── 09-quality/09.01-*.md ...
│   │   ├── 10-operations/10.01-*.md ...
│   │   └── 11-orchestration/11.01-*.md ...
│   └── phases/phase-00-*.md ...
└── harness/
    ├── build/phase-*.md
    ├── context/README.md
    └── build-log.md
```

- `docs/preliminares/` documentará las decisiones existentes, no fingirá reproducir respuestas textuales fuera del registro.
- `docs/adr/` contendrá decisiones aprobadas y alternativas; no tiene autoridad para crear cambios sin aprobación.
- `docs/phases/` describirá **planes futuros** con dependencias y gates, no informes de ejecución realizada.
- `harness/build-log.md`, si se incluye, debe registrar solo trabajo documental realmente observado en esta misión; nunca afirmar que hay producto probado/desplegado.
- El conjunto exacto de rutas es deliberadamente **propuesto**; no lo escribas sin aprobación.

## 6. PLANTILLA DE CADA ESPECIFICACIÓN FINAL

Cada especificación técnica debe poseer, en la medida aplicable:

1. `ID`, título, dominio, versión, estado (`BORRADOR / EN_REVISIÓN / APROBADO`) y owner.
2. Objetivo concreto, alcance, exclusiones y descripción de valor funcional.
3. Decisiones fuente exactas (`PRE-xx.Cy`, ADR) y enlaces relativos válidos.
4. Precondiciones, documentos dependientes y elementos que consumen esta especificación.
5. Invariantes, reglas de dominio y entidades/estructuras tipadas que se afectan.
6. Requisitos funcionales y no funcionales, con prioridades para MVP base/ampliado.
7. Diseño arquitectónico, responsabilidades, límites de servicios y diagrama textual/Mermaid cuando ayude.
8. Contratos API/eventos/datos: campos, tipos, validación, errores, versionado, autenticación, idempotencia y paginación según proceda.
9. **Metodología de implementación futura PASO A PASO**, ordenada y suficientemente específica para un agente que no vio esta conversación.
10. **Justificación de cada decisión relevante**, alternativas descartadas, costes y compensaciones.
11. Ejemplo mínimo pertinente de Go/TypeScript/SQL/JSON/pseudocódigo **solo en bloque Markdown**; no presentar el ejemplo como código ya existente.
12. Casos límite y fallos: autorización, seguridad, concurrencia, reconexión, recuperación, rendimiento y accesibilidad según el módulo.
13. Estrategia de verificación: pruebas unitarias, integración, contratos, E2E, rendimiento/visual, datos de prueba y comando esperado **como documentación**, no ejecutado.
14. **Criterios de aceptación objetivos** (`AC-01`, etc.): entrada → resultado verificable → evidencia exigida; incluir casos negativos.
15. Observabilidad, métricas, despliegue, rollback, migración y guardrails si aplica.
16. Riesgos, dependencias externas y cuestiones `POR_VERIFICAR` con prueba futura propuesta.
17. Evidencias requeridas para cerrar implementación futura y links de trazabilidad.

Evita repetir párrafos enteros de otras specs. Enlaza a la fuente canónica. Secciones irrelevantes pueden marcarse `NO_APLICA` con motivo; no añadir relleno. Markdown debe ser válido, ejemplos coherentes y los links relativos deben apuntar a rutas existentes o planeadas identificadas.

## 7. DOMINIOS QUE DEBEN QUEDAR CUBIERTOS

**01 Infraestructura:** stack Go/TypeScript, 4 proyectos, monorepo pnpm+Go, runtime Vercel beta, módulos compartidos, costes y decisiones de arquitectura.  
**02 Producto:** modos, reglas de interfaz, MVP base/ampliado, invitados, rankings, tutorial, IA, criterios de alcance.  
**03 Motor del juego:** reglamento oficial, índices/orientación, estado `[2][9]`, siembra, captura, tuzdyk, repetición, atsyrau, finales, invariantes, movimiento inmutable, eventos semánticos, corpus y pruebas.  
**04 Visual/UX:** Three.js/R3F/Drei, cámara, PBR, modelado procedural/Blender, InstancedMesh, selección/animación, RDomino como inspiración sin copia, móvil, accesibilidad HTML y test visual.  
**05 Backend:** Go/Gin, hexagonal, REST/OpenAPI, RFC 9457, concurrencia, transacciones, pgx/sqlc/Goose, DB ownership, publicación de eventos y límites serverless.  
**06 Multiplayer:** invitaciones, salas, Ably, tokens privados, control de versiones, reconexión 24 h, clientes múltiples, caducidad, Elo/ranked, abuso y resiliencia.  
**07 Identity/Data:** Google OIDC/PKCE, invitados, cuentas, refresh/JWKS, RBAC, datos públicos, perfiles, Elo/estadísticas, exportación y eliminación distribuida.  
**08 IA/Practice:** Negamax/alpha-beta, heurística, tabla transposiciones, cómputo limitado Vercel, IA legal, dificultades, práctica undo/redo, escenarios, tutorial y fallback.  
**09 Quality:** TDD, Go fuzz/race, DB real, OpenAPI, Playwright, WCAG 2.2, cargas, seguridad/CI, bugs P0–P3, definition of done.  
**10 Operations:** Vercel Hobby/Go beta, Neon/Ably, regiones, backups/restauración, secretos, observabilidad, coste cero, previews, release, incidencias y portabilidad.  
**11 Orchestration:** AGENTS/GOALS/PLANS/PROMPTS, ADR, trazabilidad, lotes, gates, asignación de agentes, gestión de cambios y plan de hitos.

No omitas subdominios porque estén aprobados en otro PRE. Identifica propietario canónico y referencias cruzadas; evita definiciones contradictorias (p.ej. motor de reglas compartido vs autoridad de Game; partidas casuales vs ranked).

## 8. ORDEN DE TRABAJO Y GATES HUMANOS

Propón un DAG de dependencias; el siguiente es un orden ORIENTATIVO:

- **Etapa 0, solo lectura:** auditoría y propuesta de inventario. **GATE 0: «APRUEBO INVENTARIO DOCUMENTAL».**
- **Lote 1, bases documentales:** registro de decisiones, convenciones, índices, ADR iniciales, AGENTS/GOALS/PLANS/PROMPTS necesarios, matriz inicial y documentación de preliminares. **GATE 1: «APRUEBO LOTE 1».**
- **Lote 2, definición de producto y dominio:** documentos 01, 02 y 03, especialmente contrato de motor/reglamento. **GATE 2.**
- **Lote 3, servicios y contratos:** documentos 05, 07 y 06 con dependencias resueltas, separación de datos y OpenAPI/eventos. **GATE 3.**
- **Lote 4, presentación y aprendizaje:** documentos 04 y 08, usando contratos del motor y criterios de accesibilidad. Pueden elaborarse en paralelo si no comparten archivos. **GATE 4.**
- **Lote 5, calidad y operación:** documentos 09 y 10, y especificaciones de orquestación 11; planes futuros de hitos. **GATE 5.**
- **Lote 6, auditoría documental integral:** matriz completa, verificación de cobertura, compatibilidad semántica, links, ADR, formato, riesgos y requisitos huérfanos. Informe de documentación lista/pendiente; **sin implementación**.

Tras CADA LOTE: publica lista de archivos creados/modificados, decisiones referenciadas, dependencias, contradicciones, riesgos, verificaciones efectivamente realizadas, criterios pendientes y la fórmula exacta de aprobación del próximo lote; **DETENTE** hasta recibirla.

No interpretes una aprobación general del prompt como aprobación automática de todos los lotes.

## 9. ORQUESTACIÓN DE AGENTES

Si tu sesión de Codex dispone REALMENTE de delegación/subagentes, utiliza un coordinador principal y hasta un pequeño número de especialistas necesarios (motor Go, microservicios, UX/Three.js, QA/seguridad, infraestructura/documentación). No afirmes haber creado agentes si no existe la función.

Reglas:

- Delegar solo tareas independientes con brief explícito: archivos propios, fuentes, dependencias, salida, restricciones y pruebas documentales.
- Un único propietario **escribe** cada archivo en un lote; los revisores comentan sin colisionar con editores.
- El coordinador integra, resuelve duplicidades y asegura nombres/IDs/contratos comunes.
- No delegar pasos lineales dependientes que requieran aprobación previa.
- Contexto por agente: AGENTS raíz, decisiones PRE pertinentes, specs dependientes y tareas exactas; no cargar todo sin necesidad.
- Registrar en el reporte qué trabajo hizo cada agente y qué revisó el coordinador.
- Sin subagentes disponibles, ejecutar secuencialmente la MISMA metodología, sin inventar colaboración.

## 10. PRUEBAS DE CALIDAD DE LA DOCUMENTACIÓN

Después de cada lote y al final:

1. **Cobertura:** las 246 decisiones del anexo están enlazadas a una spec canónica o a una exclusión/deferimiento aprobado. No dejar decisiones huérfanas.
2. **Numeración:** unicidad de identificadores, nombres y secciones; sin enlaces rotos.
3. **Dependencias:** no circulares para ejecución; distinguir especificación conceptual y secuencia futura de desarrollo.
4. **Consistencia contractual:** Go/TypeScript/OpenAPI/eventos, tipos/errores/versiones y autorizaciones no se contradicen entre servicios.
5. **Seguridad/datos:** privacidad entre Identity/Game/AI, aislamiento de DB, idempotencia de resultados/Elo y límites de tokens claros.
6. **Reglamento:** parámetros y transiciones basados en documentación normativa; casos dudosos en conflicto explícito, no resueltos por invención.
7. **Realismo operativo:** no se declaran medidos los FPS, la fuerza de IA, las latencias, los costes, el despliegue, el backup o el runtime si todavía no se ejecutaron pruebas.
8. **Calidad de ejemplos:** ejemplos Go/SQL/TypeScript internamente coherentes con contratos, sin requerir implementación en esta fase.
9. **DoD:** criterios de aceptación medibles (positivos y negativos), método futuro de verificación y artefacto de evidencia.
10. **Auditoría de alcance:** `git status`/`git diff --check` cuando existan y sean aplicables; confirmar que solo hay cambios a `.md` aprobados. No utilizar instalaciones ni ejecutar código del producto.
11. **Paralelismo:** comprobar que ningún agente sobrescribió documentos ajenos o reemplazó sin aprobación una decisión.
12. **Accesibilidad/UX:** verificar requisitos sobre tableros HTML y Three.js sin duplicar reglas.

No marques como `APROBADA` una especificación que el propietario aún no revisó. La generación y validación documental pueden dejarla en `EN_REVISIÓN`.

## 11. PLAN DE IMPLEMENTACIÓN FUTURA — SOLO DOCUMENTADO

En `PLANS.md` o docs/phases, propone hitos verticales futuros (NO LOS EJECUTES):

1. Spike técnico autorizado posterior: Vercel Go/Gin, módulos Go compartidos, Neon pooled/direct, Ably, comunicación entre servicios, cuotas y latencia.
2. Motor Go puro y pruebas reglamentarias.
3. Contratos/API y persistencia, identidad invitada.
4. Cliente mínimo con tablero jugable y accesibilidad HTML; luego visual 3D en incrementos.
5. Salas privadas, presencia, reconexión y dos jugadores online.
6. IA básica, práctica reversible y tutorial.
7. Identidad OIDC, perfiles, partidas clasificatorias, Elo y estadísticas avanzadas.
8. Hardening, pruebas multidispositivo, backup/restauración y demo de portafolio.

Cada hito debe indicar dependencias, resultados demostrables, riesgos, pruebas, criterios de aceptación, documentación afectada y aprobación requerida. No prometer fechas ni costes no fundamentados.

## 12. FORMATO DE LA RESPUESTA INICIAL Y PARADA

Tu **PRIMERA respuesta** debe contener únicamente:

- Resumen observado del repositorio y fuentes leídas.
- Confirmación de recepción del registro de 246 decisiones y sustituciones clave.
- Matriz preliminar PRE → dominios → archivos propuestos.
- Inventario exacto y por lotes de `.md` a crear/modificar, con DAG de dependencias.
- Riesgos, discrepancias, supuestos y puntos `POR_VERIFICAR`.
- Declaración «NO se ha implementado, instalado, desplegado ni modificado código».
- Una petición concreta: **«Responde: APRUEBO INVENTARIO DOCUMENTAL para autorizar el primer lote».**

**DETENTE.** No generes archivos hasta esa aprobación. Las aprobaciones siguientes se solicitan al terminar cada lote. Todos los archivos creados deben ser Markdown. Cuando todo haya sido aprobado y auditado, concluye con un índice verificable y la declaración: **«Documentación final generada; implementación no iniciada».**

---

# ANEXO A — REGISTRO CONSOLIDADO DE DECISIONES (FUENTE DEL PROMPT)

El siguiente bloque forma parte de la fuente de requisitos, no es contenido opcional. Prioriza sus sustituciones y matices aprobados por el usuario.

# Toguz Kumalak — Registro consolidado de decisiones PRE-01 a PRE-11

**Estado:** APROBADO EN CONVERSACIÓN · **Fecha de consolidación:** 2026-10-21 · **Alcance:** documentación y futuras especificaciones del proyecto.

Este documento transcribe y consolida las decisiones explícitamente aprobadas por el propietario del proyecto en los once cuestionarios preliminares. **No acredita que exista software implementado, proveedores conectados, pruebas ejecutadas o despliegues realizados.** La aprobación se refiere a intención y requisitos.

## Cómo interpretar este registro

- Cada identificador `PRE-NN.X#` es una decisión aprobada. Los resúmenes están redactados para permitir generación de especificaciones técnicas sin necesitar el historial de chat.
- Las decisiones posteriores **sustituyen** tecnologías previas cuando se indica: Go/Gin sustituye Node/Fastify; Vercel sustituye Railway; Ably sustituye Socket.IO o servidores WebSocket propios como obligación; sqlc/pgx/Goose sustituyen Drizzle; microservicios Identity/Game/AI sustituyen al monolito inicial; 0 USD sustituye el techo de 20 USD.
- La opción 1 de cada pregunta fue aprobada salvo elecciones originales de PRE-01 y PRE-02 que aquí están reflejadas en su versión corregida y PRE-05.A1 (microservicios). No volver a aplicar opciones anteriores obsoletas.
- Lo **aprobado** no debe confundirse con lo **verificado**. Versiones, planes gratuitos, límites de proveedores, reglamento oficial y compatibilidad concreta Go/Vercel requieren contrastación técnica y deben documentarse como `POR_VERIFICAR` si todavía no se han probado.
- La primera intervención de Codex debe ser **solo documental**, con aprobación previa al listado de rutas y a cada lote posterior. No iniciar implementación.

## Jerarquía y reglas inmutables

1. Respeto del reglamento tradicional contrastado con fuente oficial. En caso de discrepancia, investigar y registrar sin cambiar unilateralmente el juego; conservar 18 pozos (9 por jugador), 9 piedras iniciales por pozo, 162 en total, kazanes, victoria con 82 piedras, empate 81–81, siembra, captura par en campo rival, tuzdyk al formar exactamente tres en un pozo rival permitido (sin noveno pozo ni pozo homólogo del tuzdyk contrario) y reglas oficiales de atsyrau/repetición/final; todo sujeto a verificación contra el reglamento formal.
2. Backend autoritativo y motor Go único de reglas. React/Three.js no decide jugadas legales, resultados ni Elo.
3. Tres microservicios Go independientes con propiedad de datos por dominio y contratos versionados; Identity, Game y AI, y un frontend Next.js separado.
4. Vercel Hobby como proveedor principal objetivo, Neon y Ably Free, respetando uso personal no comercial y sin cargos sin autorización; factibilidad técnica real no se presume.
5. Tablero 3D propio de calidad premium con inspiración en experiencia de rdomino.com, sin copiar propiedad intelectual o materiales ajenos.
6. Incluye IA básica, rankings y estadísticas avanzadas en el MVP ampliado, sin mover silenciosamente estas capacidades a pos-MVP.
7. Las decisiones se modifican solo por propuesta de cambio y aprobación humana documentada (ADR); Codex nunca las sustituye porque prefiera otra pila.

## Matriz integral de decisiones aprobadas

### PRE-01 — Infraestructura y arquitectura (17)
- **PRE-01.A1** — TypeScript para frontend y Go para backend; generar y verificar contratos tipados entre ambos lenguajes.
- **PRE-01.A2** — Next.js, React y TypeScript para la aplicación web.
- **PRE-01.A3** — Gin como framework HTTP principal de los microservicios Go.
- **PRE-01.A4** — Monorepositorio: pnpm para proyectos TypeScript y go.mod/go.work para proyectos Go; módulos compartidos con dependencias controladas.
- **PRE-01.B1** — Sustituida la elección inicial de Socket.IO por Ably Free para notificaciones distribuidas; no crear gateway Node.js ni exigir servidores WebSocket persistentes propios.
- **PRE-01.B2** — PostgreSQL como base de datos relacional; Neon Free es proveedor principal previsto.
- **PRE-01.B3** — Sustituir Drizzle por SQL tipado con sqlc + pgx/v5 y migraciones SQL con Goose, nativos de Go.
- **PRE-01.B4** — Guardar movimientos aceptados y estado actualizado de forma atómica y transaccional.
- **PRE-01.B5** — Permitir invitados que juegan de inmediato y cuentas opcionales para identidad persistente y clasificación.
- **PRE-01.C1** — Sustituir Railway por Vercel como proveedor principal; cuatro proyectos: frontend, Identity, Game y AI; factibilidad pendiente de pruebas de despliegue.
- **PRE-01.C2** — Sustituir el presupuesto máximo inicial de 20 USD por objetivo operativo de 0 USD/mes; prohibido habilitar recursos de pago sin permiso explícito.
- **PRE-01.C3** — Desarrollo local con servicios Go/Next.js nativos cuando sea práctico y PostgreSQL/infraestructura auxiliar en Docker Compose.
- **PRE-01.C4** — Separación entre desarrollo, previews/entorno de integración y producción; no compartir datos ni credenciales productivas.
- **PRE-01.D1** — GitHub Actions para tipado, formato/lint, tests, build, integración y controles de calidad.
- **PRE-01.D2** — Observabilidad avanzada con OpenTelemetry, Prometheus/Grafana principalmente localmente y telemetría productiva según cuotas verificadas.
- **PRE-01.D3** — Crecer desde una instancia por microservicio con medidas reales; Redis es opcional y se incorporará solo si se justifica por capacidad/coordinación, nunca preventivamente obligatorio.
- **PRE-01.D4** — Repositorio de portafolio público; elegir licencia expresamente antes de publicar derechos de reutilización.

### PRE-02 — Producto, funcionalidades y alcance (18)
- **PRE-02.A1** — Jugar localmente con dos personas en un mismo dispositivo y contra IA básica; incluir modo práctica separado.
- **PRE-02.A2** — Partidas online en salas privadas por enlace o código de invitación; sin emparejamiento abierto obligatorio.
- **PRE-02.A3** — Partidas online en tiempo real con ambos participantes conectados durante el juego ordinario y política de recuperación al desconectar.
- **PRE-02.A4** — Sin reloj competitivo ni control de tiempo por partida en el MVP.
- **PRE-02.B1** — Los invitados pueden jugar local/online; la cuenta registrada habilita historial personal persistente y clasificación.
- **PRE-02.B2** — Alias personalizable y un identificador temporal verificable para invitados.
- **PRE-02.B3** — Interfaz en español e inglés desde el MVP.
- **PRE-02.B4** — Tutorial interactivo opcional y consulta completa del reglamento.
- **PRE-02.C1** — Resaltar casillas legales, turno activo y vista previa básica de movimientos con datos autorizados por Go.
- **PRE-02.C2** — En partidas ordinarias no existe deshacer; se añadió expresamente un modo práctica reversible con deshacer/rehacer.
- **PRE-02.C3** — Reconexión automática y partida online recuperable inicialmente durante 24 h; vencido el plazo, se archiva como inconclusa.
- **PRE-02.C4** — Abandono voluntario y revancha consensuada; AMPLIADA en PRE-03 para incluir propuestas de tablas por mutuo acuerdo y reclamación reglamentaria por repetición, además de tablas por resultado.
- **PRE-02.D1** — Historial de partidas terminadas y reproducción básica de sus movimientos.
- **PRE-02.D2** — Perfil con alias, partidas, victorias, derrotas, empates e historial.
- **PRE-02.D3** — Guardado automático de la última partida local en el navegador y posibilidad de reanudarla; no se exige funcionamiento offline.
- **PRE-02.E1** — Three.js como motor visual principal del tablero tridimensional.
- **PRE-02.E2** — Estética de tablero físico premium/fotorrealista inspirada en rdomino.com, sin copiar archivos, diseños distintivos, modelos, sonidos ni código.
- **PRE-02.E3** — MVP ampliado: IA básica, clasificación competitiva y estadísticas avanzadas; torneos, chat, espectadores, IA avanzada, pagos y funciones sociales extensas quedan pospuestos.

### PRE-03 — Reglas, estados y motor Go (22)
- **PRE-03.A1** — Detectar triple repetición de una misma posición reglamentaria y permitir reclamación de tablas; no empate automático por tercera aparición salvo que la regla formal lo exija.
- **PRE-03.A2** — Permitir propuestas de tablas por mutuo acuerdo; modifica PRE-02 C4.
- **PRE-03.A3** — Tuzdyk se crea automáticamente al cumplir las reglas, con señalización, explicación y animación contextual.
- **PRE-03.A4** — Blancas empiezan; asignación online aleatoria de colores y elección de bando en local/IA.
- **PRE-03.B1** — Estado lógico con tablero fijo [2][9], kazanes, posición/propiedad de tuzdyks, turno y versión, con convenio documentado de índices y recorrido circular.
- **PRE-03.B2** — Motor puro determinista que recibe estado + movimiento y devuelve nuevo estado + eventos + error, sin mutar el original ni tocar persistencia/red.
- **PRE-03.B3** — Comandos autorizados por identidad de sesión con casilla origen, versión esperada e identificador idempotente; no recibir ni confiar en tablero completo de cliente.
- **PRE-03.B4** — Versionar formato de estado, reglamento y registros de movimientos para reproducción estable.
- **PRE-03.B5** — Go determina movimientos legales y previsualizaciones; frontend no duplica reglamento.
- **PRE-03.C1** — Representar tuzdyks y sus propietarios explícitamente; las piedras que caen allí incrementan inmediatamente el kazán correspondiente.
- **PRE-03.C2** — Generar eventos semánticos de siembra, captura, tuzdyk, turnos y fin para que Three.js pueda animarlos sin conocer reglas.
- **PRE-03.C3** — Rechazar movimientos ilegales con errores tipados y sin alterar el estado original.
- **PRE-03.C4** — Guardar estado de resultado separado de su causa (reglas, abandono, empate, desconexión/inconclusa, etc.).
- **PRE-03.D1** — Modo práctica con control libre y escenarios educativos; sin efectos competitivos.
- **PRE-03.D2** — Historial de práctica con deshacer/rehacer por estados; al jugar después de retroceder, descartar la rama posterior.
- **PRE-03.D3** — IA básica con Negamax/Minimax y poda alfa-beta, heurística explícita.
- **PRE-03.D4** — AI Service ejecuta Go en Vercel con límites estrictos de cálculo, cancelación y respuesta legal.
- **PRE-03.D5** — Dos dificultades: fácil y normal, diferenciadas y evaluadas.
- **PRE-03.E1** — Corpus de posiciones del reglamento formal y contraste independiente con implementaciones de referencia.
- **PRE-03.E2** — Pruebas unitarias, de propiedades, fuzzing, simulaciones e invariantes; en particular total de 162 piedras conservado.
- **PRE-03.E3** — Reproducción basada en movimientos/estados versionados y recálculo por motor; evitar persistir instrucciones 3D como fuente de verdad.
- **PRE-03.E4** — Benchmarks y objetivos cuantitativos después de obtener corrección funcional y mediciones reproducibles.

### PRE-04 — Experiencia visual, UI y Three.js (25)
- **PRE-04.A1** — Realismo físico estilizado de alta calidad, con acabados creíbles y estética premium, sin fotorrealismo extremo obligatorio.
- **PRE-04.A2** — Maderas nobles cálidas/oscuras, cavidades talladas, piedras naturales pulidas y detalles metálicos discretos.
- **PRE-04.A3** — Referencias culturales kazajas verificadas, aplicadas con respeto y sin ornamentos inventados.
- **PRE-04.A4** — Three.js representa tablero; menús, perfiles, HUD, estadísticas y controles accesibles en React/HTML/CSS.
- **PRE-04.B1** — Three.js + React Three Fiber + Drei.
- **PRE-04.B2** — Geometría procedural junto con Blender y exportación GLB/glTF para recursos detallados.
- **PRE-04.B3** — InstancedMesh para las 162 piedras, con disposición determinista dependiente del estado de juego.
- **PRE-04.B4** — Cámara oblicua ajustable con zoom y rotación controlados, cambios de orientación disponibles.
- **PRE-04.B5** — Materiales PBR, iluminación ambiental cuidada, sombras suaves y posprocesado moderado adaptativo.
- **PRE-04.C1** — Selección y confirmación separadas para evitar jugadas accidentales.
- **PRE-04.C2** — Animaciones de siembra secuenciales, acelerables y con opción de saltar sin alterar estado lógico.
- **PRE-04.C3** — Animaciones de Three.js consumen eventos semánticos del motor Go, no infieren reglas.
- **PRE-04.C4** — Sonidos de materiales y eventos opcionales, discretos, configurables, sin reproducción invasiva.
- **PRE-04.D1** — Portada, elección de modo, sala, tablero, resultados, tutorial, perfil, historial y clasificación como navegación clara por pantallas.
- **PRE-04.D2** — Móvil vertical y horizontal, adaptación automática de cámara/controles y sugerencia de horizontal cuando beneficie la legibilidad.
- **PRE-04.D3** — Orientación de cámara adaptada para juego local, giro opcional entre turnos y opción de cámara fija.
- **PRE-04.D4** — Accesibilidad de teclado, datos de casillas textuales, indicadores no basados solo en color y alternativa HTML interactiva.
- **PRE-04.D5** — Tutorial opcional con resaltado contextual de casillas y escenarios dentro del tablero 3D.
- **PRE-04.E1** — Objetivo 60 FPS escritorio y 30 FPS móvil compatible, contrastados con dispositivos de referencia; no afirmarlos como garantía.
- **PRE-04.E2** — Calidades automática, baja y alta con ajuste de resolución, sombras y efectos.
- **PRE-04.E3** — Carga progresiva/diferida de 3D, materiales reutilizados y recursos optimizados/comprimidos.
- **PRE-04.E4** — Vista alternativa HTML plenamente jugable si WebGL está ausente o falla.
- **PRE-04.F1** — Pruebas visuales automatizadas, referencia de escenas, revisiones multiplataforma y medición de rendimiento.
- **PRE-04.F2** — Recursos originales o debidamente licenciados, con inventario de procedencia; no copiar recursos de RDomino.
- **PRE-04.F3** — Aprobación visual incremental: geometría → materiales/cámara → interacción → animaciones → móvil → acabado.

### PRE-05 — Backend, contratos y persistencia (24)
- **PRE-05.A1** — SUSTITUIDO el monolito: exactamente tres microservicios Go/Gin independientes (Identity, Game, AI) y frontend Next.js; autonomía por dominio.
- **PRE-05.A2** — Arquitectura hexagonal pragmática dentro de cada microservicio: transporte, aplicación, dominio y adaptadores.
- **PRE-05.A3** — Dependencias inyectadas explícitamente mediante constructores e interfaces pequeñas cercanas al consumidor.
- **PRE-05.A4** — Cuatro proyectos Vercel, despliegues independientes desde el monorepositorio, módulo Go autónomo por servicio.
- **PRE-05.B1** — REST versionado bajo /api/v1, consultas de recursos y comandos explícitos.
- **PRE-05.B2** — OpenAPI 3.1 como fuente de verdad de HTTP, clientes/modelos generados en Go y TypeScript, versiones fijadas.
- **PRE-05.B3** — Errores estructurados compatibles con RFC 9457 (Problem Details), códigos de dominio, requestId y mensajes seguros.
- **PRE-05.B4** — APIs por recursos/dominios (games, players, practice, leaderboard), paginación y sin exponer tablas directamente.
- **PRE-05.C1** — Pipeline de jugada: autenticar → autorizar → versión → motor Go → persistencia → respuesta autoritativa.
- **PRE-05.C2** — Bloqueo transaccional por partida, comprobación de versión e identificador único idempotente.
- **PRE-05.C3** — Comando aceptado, movimiento, efectos y estado confirmados dentro de la misma transacción.
- **PRE-05.C4** — Simulación local/práctica mediante endpoints sin autoridad competitiva; sin motor paralelo TypeScript.
- **PRE-05.D1** — Propiedad de datos separada: Identity DB, Game DB, AI stateless; no consultas SQL cruzadas entre servicios.
- **PRE-05.D2** — Estado actual compacto versionado en JSONB validado por Go; movimientos normalizados y columnas de consulta relacionales.
- **PRE-05.D3** — pgxpool limitado, Neon pooled para aplicación cuando proceda y conexión directa para migraciones.
- **PRE-05.D4** — Goose con migraciones SQL versionadas/controladas, sin migrar automáticamente al arrancar Functions.
- **PRE-05.E1** — Eventos Ably ligeros con gameId, versión, tipo y datos mínimos; sin instrucciones 3D.
- **PRE-05.E2** — Confirmar PostgreSQL antes de publicar en Ably; reconciliar por REST tras reconexión/foco y verificaciones acotadas; documentar riesgo de aviso perdido sin outbox.
- **PRE-05.E3** — Tokens Ably temporales de suscripción limitados a canales de partida; publicación oficial desde Game.
- **PRE-05.E4** — JSON schemas versionados para eventos y validación en Go/TypeScript.
- **PRE-05.F1** — OpenTelemetry y trazas de API, servicios, SQL y Ably con referencias correlacionadas.
- **PRE-05.F2** — Pruebas con PostgreSQL real, HTTP/Gin, contratos, carreras/consistencia y doble de Ably.
- **PRE-05.F3** — Autorización server-side, límites de tasa, validación y antiabuso en salas, invitados e IA.
- **PRE-05.F4** — Prueba temprana obligatoria de Vercel Go/Gin + Neon + Ably + transacción antes de iniciar implementación completa; aún no ejecutada.

### PRE-06 — Salas, multijugador y ranking (20)
- **PRE-06.A1** — Invitaciones por token de alta entropía y código alternativo protegido con expiración y rate limits.
- **PRE-06.A2** — Salas pendientes caducan tras 30 min de inactividad (configurable), sin afectar política de partidas iniciadas.
- **PRE-06.A3** — Sala con máximo dos jugadores y sin espectadores en MVP.
- **PRE-06.A4** — Revancha genera nueva partida e identifica su relación con anterior; colores reasignados.
- **PRE-06.B1** — Selección visual inmediata, pero ejecución/animación definitiva solo tras confirmación autoritativa del servidor.
- **PRE-06.B2** — Eventos con ID y versión; detectar eventos perdidos/desordenados y volver a sincronizar por REST.
- **PRE-06.B3** — Estados visibles conectado/sincronizando/desconectado/recuperado, bloqueando acciones sin confirmar.
- **PRE-06.B4** — Un canal privado Ably por partida; cliente suscribe con token de capacidad restringida.
- **PRE-06.C1** — Reanudar exclusivamente con identidad del participante original, de invitado o registrada, y estado confirmado desde Game.
- **PRE-06.C2** — Persistir vencimiento de 24 horas y verificarlo en acciones/consultas; limpieza de datos en segundo plano no determina validez.
- **PRE-06.C3** — Varias pestañas pueden observar; transacciones e idempotencia impiden duplicar movimientos.
- **PRE-06.C4** — Tras reconexión, reconstruir tablero y mostrar último movimiento confirmado.
- **PRE-06.D1** — Solo partidas expresamente clasificatorias entre dos registrados afectan Elo; condición fijada al comienzo.
- **PRE-06.D2** — Elo básico documentado, independiente de estadísticas generales.
- **PRE-06.D3** — Medidas contra abuso de enfrentamientos repetidos/manipulación de Elo, con límites auditables.
- **PRE-06.D4** — Abandono voluntario es derrota; partida archivada por desconexión como inconclusa no modifica Elo.
- **PRE-06.E1** — Ante fallo Ably no perder movimientos PostgreSQL; informar interrupción y reconciliar por REST.
- **PRE-06.E2** — Game valida tokens firmados sin consultar Identity en cada jugada; caídas de Identity limitan altas/renovaciones.
- **PRE-06.E3** — E2E con dos clientes, desconexiones, duplicados, solicitudes simultáneas, múltiples pestañas y fallos parciales.
- **PRE-06.E4** — Medir p50/p95 de confirmación y resincronización antes de fijar SLOs definitivos.

### PRE-07 — Identidad, permisos, datos y Elo (24)
- **PRE-07.A1** — Registro con Google OpenID Connect / Authorization Code + PKCE, sin contraseñas propias en MVP.
- **PRE-07.A2** — Identity crea invitados con UUID, alias temporal y sesión verificable y recuperable.
- **PRE-07.A3** — Invitado puede vincular cuenta; partidas elegibles pasan a historial, no se convierten retroactivamente en Elo.
- **PRE-07.A4** — UUID inmutable para usuario, alias público único de registrados y alias temporales no necesariamente únicos para invitados.
- **PRE-07.B1** — Access tokens asimétricos cortos (JWKS público), audiencias/emisor validados en Game; Identity no consultado por movimiento.
- **PRE-07.B2** — Acceso aproximadamente 15 min, sesiones renovables con inactividad configurable y duración de invitados compatible con 24 h.
- **PRE-07.B3** — Refresh rotatorio, revocación al salir, registro de sesiones; tokens emitidos aún válidos hasta expiración corta salvo mecanismos extraordinarios.
- **PRE-07.B4** — Cookies HttpOnly/Secure/SameSite adecuadas, mismo origen, CSRF y verificación de origen; no guardar tokens sensibles en localStorage.
- **PRE-07.C1** — Roles mínimos guest/user/admin con autorización adicional por recurso.
- **PRE-07.C2** — Game verifica identidad, participante, turno, estado y permiso para cada operación.
- **PRE-07.C3** — Credenciales servicio-servicio de vida corta, audiencia y permisos propios, HTTPS y rotación.
- **PRE-07.C4** — Game referencia UUID y obtiene perfiles públicos por API de Identity en lotes; no accede a Identity DB.
- **PRE-07.D1** — Perfil público limita a alias, Elo y estadísticas; ocultar correo y datos privados.
- **PRE-07.D2** — Alias editable y avatar generado, sin subidas arbitrarias de imágenes en MVP.
- **PRE-07.D3** — Validación/moderación básica de alias, reservados y abuso.
- **PRE-07.D4** — Estadísticas/Elo públicos; historial detallado propio y acceso controlado a reproducción compartida.
- **PRE-07.E1** — Elo empieza 1200, K=32, primeras diez partidas clasificatorias provisionales; parámetros a simular/validar.
- **PRE-07.E2** — Estadísticas avanzadas: partidas, resultados, tasa victoria, evolución Elo, tiempo registrado por turno (sin reloj competitivo), capturas, tuzdyks, rendimiento por color/modalidad.
- **PRE-07.E3** — Ranking global paginado, búsqueda por alias, posición personal e identificación de provisionales.
- **PRE-07.E4** — Game actualiza resultado, Elo y agregados en transacción atómica al finalizar partida; reintentos idempotentes.
- **PRE-07.F1** — Identidades invitadas con caducidad y retención limitada, compatibles con recuperación 24 h, limpieza posterior documentada.
- **PRE-07.F2** — Exportación del propio perfil, historial y estadísticas a JSON estructurado.
- **PRE-07.F3** — Eliminación de cuenta distribuida e idempotente: bloquear, revocar, solicitar borrado/anonimización en Game, verificar y finalizar según privacidad aplicable.
- **PRE-07.F4** — Pruebas de tokens, cuentas, invitados, vinculación, accesos cruzados, borrado y manipulación de rankings.

### PRE-08 — IA, práctica y tutorial (24)
- **PRE-08.A1** — Turno IA se calcula automáticamente tras confirmar movimiento humano cuando toca a IA.
- **PRE-08.A2** — Dos dificultades por presupuesto/profundidad/heurística, demostrablemente diferenciadas.
- **PRE-08.A3** — Mostrar indicador calculando, dificultad y última jugada; no mostrar análisis profundo en partidas normales.
- **PRE-08.A4** — IA sigue reglas y juega hasta final reglamentario, sin rendiciones automáticas heurísticas.
- **PRE-08.B1** — Negamax alfa-beta con ordenación y profundización iterativa.
- **PRE-08.B2** — Evaluación explicable: diferencia de kazanes, tuzdyk, vulnerabilidad, movilidad, fase.
- **PRE-08.B3** — Tabla de transposiciones temporal y acotada, con tratamiento correcto del historial en triple repetición.
- **PRE-08.B4** — Límites de tiempo, profundidad, nodos y memoria; context.Context y mejor movimiento legal disponible.
- **PRE-08.C1** — Práctica controla ambos bandos, IA opcional y escenarios educativos verificados.
- **PRE-08.C2** — Deshacer/rehacer por movimiento; atajo «volver a mi turno» ante IA.
- **PRE-08.C3** — Catálogo inicial pequeño de ejercicios con captura, tuzdyk, defensa y final, validado.
- **PRE-08.C4** — Última sesión de práctica y historial reversible en almacenamiento local versionado.
- **PRE-08.D1** — Tutorial de lecciones progresivas: turnos, siembra, captura, tuzdyk, final y conceptos básicos.
- **PRE-08.D2** — Sugerencias opt-in con motivos verificables; deshabilitadas durante partidas clasificatorias.
- **PRE-08.D3** — Ejercicios con objetivos específicos, retroalimentación y distinción entre jugada legal y objetivo educativo.
- **PRE-08.D4** — Progreso local, reiniciable/repetible y traducido español/inglés.
- **PRE-08.E1** — Panel de práctica para deshacer/rehacer/reiniciar, consultar jugadas y escenarios, diferenciando partidas normales.
- **PRE-08.E2** — Previsualizaciones educativas 3D temporales sin alterar estado confirmado.
- **PRE-08.E3** — Historial paso a paso con anterior/siguiente/pausa/velocidad y reutilización de eventos semánticos.
- **PRE-08.E4** — Alternativa HTML sin WebGL permite mismos flujos de tutorial y práctica con un motor único Go.
- **PRE-08.F1** — Corpus reproducible apertura/medio/final, legalidad y condiciones terminales, cancelaciones y límites.
- **PRE-08.F2** — Torneos automatizados versus jugador aleatorio legal y referencias; medir fuerza relativa, tiempo y nodos sin Elo no verificado.
- **PRE-08.F3** — Si AI Service falla, Game verifica y puede escoger respaldo legal barato con motor compartido, sin corromper estado.
- **PRE-08.F4** — Medir duración, profundidad, nodos, recursos, cancelaciones y fallos sin datos privados innecesarios.

### PRE-09 — Testing, CI/CD y Definition of Done (24)
- **PRE-09.A1** — Pirámide basada en riesgos: unitarias motor/servicios, integración SQL/contratos y E2E críticos.
- **PRE-09.A2** — TDD para reglas, transacciones, autorización y contratos; aceptación y pruebas visuales para UI/3D.
- **PRE-09.A3** — Cobertura objetivo 90 % sentencias motor Go, 80 % lógica servicios; rutas críticas probadas aun con cobertura alta.
- **PRE-09.A4** — Fixtures y semillas reproducibles, casos reglamentarios y factorías aisladas para usuarios/partidas.
- **PRE-09.B1** — Corpus oficial, propiedades, fuzzing y simulación de partidas con conservación de 162 piedras.
- **PRE-09.B2** — IA: jugada legal, límites, cancelación, semillas fijas, estados terminales y diferencias medibles.
- **PRE-09.B3** — PostgreSQL real efímero en CI, Goose desde cero, consultas sqlc, índices, restricciones y transacciones.
- **PRE-09.B4** — Carreras de movimientos duplicados/concurrentes, versiones obsoletas y go test -race donde aplique.
- **PRE-09.C1** — Compatibilidad OpenAPI, modelos generados y validación de esquemas de eventos/consumidores.
- **PRE-09.C2** — Pruebas JWKS, caducidad, audiencias, rotación, privilegios de servicio y acceso cruzado.
- **PRE-09.C3** — Simulación de fallos de Identity, Game, AI, Ably y PostgreSQL; reintentos y degradación segura.
- **PRE-09.C4** — Docker Compose para integración aislada; proveedores simulados en la mayoría de tests, integración real aparte.
- **PRE-09.D1** — Vitest + React Testing Library para componentes y adaptadores visuales.
- **PRE-09.D2** — Playwright Chromium/Firefox/WebKit, móvil y dos participantes en contextos aislados.
- **PRE-09.D3** — Screenshots 3D con entorno fijo + revisión manual por variabilidad GPU.
- **PRE-09.D4** — Objetivo WCAG 2.2 AA aplicable, teclado/lector, movimiento reducido y alternativa HTML jugable.
- **PRE-09.E1** — Feature branches y PR, protecciones de main, check obligatorios; no fusionar tests rotos.
- **PRE-09.E2** — CI rápido por PR, completo antes de merge, pruebas intensivas escalonadas/pre-release; cambios compartidos activan consumidores.
- **PRE-09.E3** — Local/CI aislados, preview Vercel, integración sintética separada y producción independiente.
- **PRE-09.E4** — Releases trazables, smoke tests, rollback de código, migración expand/contract y recuperación de datos.
- **PRE-09.F1** — Incidencias P0-P3, P0/P1 bloquean releases, especialmente corrupción, reglas e identidad.
- **PRE-09.F2** — Medir Web Vitals, p95, IA, FPS, descarga 3D, reconexión y cuotas; umbrales con benchmarks.
- **PRE-09.F3** — Escáneres de secretos/dependencias, análisis estático, govulncheck y equivalente multimódulo.
- **PRE-09.F4** — Definition of Done: requisito + prueba observada + sin bloqueantes + docs + evidencia; confirmación de agente no basta.

### PRE-10 — Despliegue, operaciones y costes (24)
- **PRE-10.A1** — Cuatro proyectos Vercel vinculados al monorepo con builds independientes; CI verifica Go compartido y dependencias por proyecto.
- **PRE-10.A2** — Adaptador pequeño de http.Handler para Gin en Go runtime de Vercel, conservando arranque local.
- **PRE-10.A3** — Frontend y APIs Identity/Game bajo mismo origen vía rewrites; AI expone endpoint HTTPS protegido para llamadas servicio-servicio.
- **PRE-10.A4** — Regiones Vercel y Neon próximas, elegidas con mediciones, disponibilidad y privacidad.
- **PRE-10.B1** — Local, integración y producción con credenciales OAuth, Ably y bases independientes.
- **PRE-10.B2** — Secretos en Vercel por entorno, rotación, no exponer secretos/logs ni versionarlos.
- **PRE-10.B3** — Previews protegidas dentro de Hobby, datos sintéticos y OAuth de prueba aislado.
- **PRE-10.B4** — Límites compartidos/centralizados compatibles con serverless por IP/identidad/recurso, especial invitados e IA; no solo RAM de Function.
- **PRE-10.C1** — pgxpool con límites serverless, conexión agrupada donde corresponda, directas para migraciones.
- **PRE-10.C2** — PITR limitado del plan Neon Free + backups PostgreSQL externos cifrados y restauraciones de prueba; servicio externo/coste pendiente de verificación.
- **PRE-10.C3** — Goose migrations aprobadas, controladas y expand/contract, por esquema de su servicio.
- **PRE-10.C4** — Expiración lógica en solicitudes más limpieza diaria idempotente compatible con límites de cron Hobby.
- **PRE-10.D1** — OpenTelemetry Go/Next.js; collector + Prometheus/Grafana locales; producción ajustada a planes gratuitos verificados.
- **PRE-10.D2** — Propagar traceId/requestId/gameId entre microservicios, sin datos sensibles en trazas.
- **PRE-10.D3** — Métricas p50/p95, errores, invocaciones, login fallido, jugadas inválidas, reconexión, CPU IA y cuotas.
- **PRE-10.D4** — Alertas cuando sean posibles y chequeos sintéticos proporcionados, sin prometer 24/7 real no contratado.
- **PRE-10.E1** — Coste objetivo 0 USD/mes, ningún recurso de pago sin autorización expresa; Hobby uso personal no comercial.
- **PRE-10.E2** — Umbrales internos de revisión 50/75/90 % de cuotas, mediante datos que los proveedores ofrezcan.
- **PRE-10.E3** — Límites de CPU/búsquedas IA por identidad, cache segura donde se pruebe y medición de consumo.
- **PRE-10.E4** — Escalar solo tras pruebas; Redis, colas o réplicas se evalúan como cambios sujetos a aprobación.
- **PRE-10.F1** — Despliegues independientes compatibles, smoke tests, migraciones expand/contract y feature flags para riesgos.
- **PRE-10.F2** — Runbooks P0-P3, degradación controlada ante fallos de reglas, auth, datos, Ably, Elo, WebGL e IA.
- **PRE-10.F3** — RPO/RTO fundamentados en restauraciones ensayadas y capacidad real, no garantizar cero pérdida.
- **PRE-10.F4** — Arquitectura Go portable a contenedores mediante adaptadores, variables y Dockerfiles documentados (pero NO crearlos en la fase documental).

### PRE-11 — Documentación y orquestación Codex (24)
- **PRE-11.A1** — Specs Markdown por dominios con numeración incremental (01.01, 01.02, etc.), índices y referencias entre documentos.
- **PRE-11.A2** — Jerarquía: registro de decisiones aprobadas + ADR, specs detalladas, AGENTS.md operativo y planes por fase; conflictos explícitos.
- **PRE-11.A3** — AGENTS.md raíz conciso y específicos por servicio solo si necesarios.
- **PRE-11.A4** — ADR numerados con decisión, alternativas, motivo, consecuencias y estado.
- **PRE-11.B1** — Plantilla: ID, objetivo, requisitos, dependencias, pasos, porqués, ejemplos, contratos, pruebas y aceptación.
- **PRE-11.B2** — Detalle suficiente para implementar sin inventar arquitectura, sin describir código línea por línea.
- **PRE-11.B3** — Al menos un ejemplo relevante de Go/TS/SQL/JSON/pseudocódigo por especificación cuando aporte claridad, explicado.
- **PRE-11.B4** — Aceptación objetiva verificable y rastreable a casos de uso y fallos.
- **PRE-11.C1** — Coordinador de coherencia y especialistas delegados solo cuando haya trabajo independiente.
- **PRE-11.C2** — Especializaciones: arquitectura, motor Go, microservicios, Three.js, seguridad, QA y operaciones según necesidad.
- **PRE-11.C3** — Escritura paralela en ficheros independientes, única propiedad por fichero, integración/índices del coordinador.
- **PRE-11.C4** — Contexto acotado por agente: instrucciones globales, alcance, decisiones y specs dependientes relevantes.
- **PRE-11.D1** — Generación escalonada: auditoría, matriz decisiones, grafo dependencias, lotes documentales y revisión de coherencia.
- **PRE-11.D2** — Priorizar según dependencias técnicas, aunque los IDs clasifiquen por dominios.
- **PRE-11.D3** — Aprobación humana por lotes coherentes; informes de cambios/riesgos antes de seguir.
- **PRE-11.D4** — Auditoría automática de enlaces, IDs, Markdown, secciones, referencias y revisión semántica.
- **PRE-11.E1** — Primera ejecución redacta SOLO documentación .md; no código, tests ejecutables, dependencias, infra real ni despliegue.
- **PRE-11.E2** — Plan futuro de entregas verticales y spike temprano de Go/Vercel/Neon/Ably, documentado ahora, ejecutado solo más tarde y con aprobación.
- **PRE-11.E3** — Cierre de fase futuro condicionado a tests observados, aceptación, revisión, docs y evidencias.
- **PRE-11.E4** — Cambios relevantes requieren propuesta/impacto/ADR y aprobación; agentes no reescriben unilateralmente decisiones.
- **PRE-11.F1** — README profesional orientado a portafolio, arquitectura, desafíos, pruebas y demo prevista, sin fingir funcionalidades implementadas.
- **PRE-11.F2** — Matriz requisito → spec → fase → componentes → pruebas → evidencia → estado.
- **PRE-11.F3** — Codex en primera fase solo puede leer y generar .md autorizados; cambios en código, secretos, deps, despliegues y Git requieren autorización independiente.
- **PRE-11.F4** — Un prompt maestro orquesta inventario, lotes, verificaciones y gates de aprobación, con informe final de documentación creada/pending/bloqueos.

## Relaciones, ambigüedades y controles que Codex debe explicitar

- **Reglamento y fuente normativa:** resolver cualquier diferencia en convención de siembra, atsyrau, declaración de tuzdyk, tabla por repetición y estado terminal contra fuentes reglamentarias oficiales antes de especificar algoritmos. Registrar casos testables. No derivar el reglamento desde gráficos.
- **Local vs. offline:** las partidas locales son para dos usuarios en un dispositivo, y están previstas para ejecutarse mediante el motor Go; *offline* no fue aprobado como requisito. No prometer funcionamiento sin internet.
- **Vista HTML accesible:** debe utilizar contratos y estados del mismo motor Go. No implica un segundo motor en JavaScript ni necesariamente capacidad offline.
- **Ranking vs. invitado/IA:** solamente partidas etiquetadas como clasificatorias entre registrados afectan Elo. Partidas de práctica, invitados o IA no afectan Elo.
- **Tres microservicios vs. presupuesto:** desplegar cuatro proyectos Vercel y dos datos separados puede exceder cuotas según carga. La arquitectura permanece, pero anotar límites, riesgos y alternativas sin habilitar pago.
- **Token de usuario y servicio:** distinguir Google OIDC, sesión propia/JWKS de Identity, credenciales entre servicios y tokens de canal Ably, con sus autoridades, audiencias y permisos respectivos.
- **Notificaciones vs. datos:** una transacción de PostgreSQL confirmada es fuente de verdad aunque Ably falle. Sin outbox duradero en MVP, documentar retrasos de notificación y estrategias de resincronización/sondeo activo acotado.
- **Arquitectura serverless y archivos Go compartidos:** confirmar en una prueba posterior el empaquetado y la compilación de módulos privados compartidos dentro de raíces de proyecto Vercel. No declarar probada una configuración que solo exista en especificación.
- **MVP base y ampliado:** organizar entregas por hitos, pero el objetivo de lanzamiento aprobado incluye funcionalidades ampliadas. La fase documental no altera el compromiso de producto.
- **Copias de seguridad:** no prometer backup externo gratuito o RPO/RTO que no se haya verificado; documentar almacenamiento, seguridad, retención y prueba de restauración como decisiones operativas por validar.
- **Licencias:** repositorio público no significa software automáticamente de código abierto; se deberá escoger una licencia deliberadamente.

## Registro de aprobación

- PRE-01: 17 decisiones aprobadas, con sustituciones expresas posteriores en hosting, backend, tiempo real, SQL y presupuesto.
- PRE-02: 18 decisiones aprobadas, C2 extendida con modo práctica y C4 ampliada en PRE-03; E1=3, E2=3, E3=2.
- PRE-03: 22 decisiones, todas opción 1.
- PRE-04: 25 decisiones, todas opción 1.
- PRE-05: 24 decisiones; A1 sustituida por arquitectura de tres microservicios, restantes recomendaciones aprobadas.
- PRE-06: 20 decisiones, todas opción 1.
- PRE-07: 24 decisiones, todas opción 1.
- PRE-08: 24 decisiones, todas opción 1.
- PRE-09: 24 decisiones, todas opción 1.
- PRE-10: 24 decisiones, todas opción 1.
- PRE-11: 24 decisiones, todas opción 1.

**Total: 246 decisiones registradas.** Si detectas una contradicción de significado, crea una entrada `CONFLICTO` y no modifiques este registro sin aprobación del propietario.
