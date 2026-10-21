# PRE-11 — Documentación y orquestación Codex

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: PRE-11 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner: C — coordinación.

## Objetivo y contexto

Consolidar las 24 decisiones que gobiernan documentación, trazabilidad y trabajo delegado. La autorización vigente permite únicamente los Markdown del lote 1 y exige parada al terminar. Aprobar el inventario autoriza rutas, no aprueba el contenido todavía no revisado ni habilita implementación.

## Fuentes y decisiones

El [maestro archivado](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md), Anexo A, y el [registro canónico](../decision-register.md) conservan las decisiones. Síntesis de naturaleza `APROBADO`:

- [PRE-11.A1](../decision-register.md#pre-11-a1): las specs se numeran por dominio con índices y referencias explícitas.
- [PRE-11.A2](../decision-register.md#pre-11-a2): decisiones/ADR aprobadas preceden specs, instrucciones operativas y planes; todo conflicto se explicita.
- [PRE-11.A3](../decision-register.md#pre-11-a3): AGENTS raíz será conciso y solo habrá instrucciones por servicio cuando sean necesarias.
- [PRE-11.A4](../decision-register.md#pre-11-a4): ADR numerados expondrán decisión, alternativas, motivo, consecuencias y estado.
- [PRE-11.B1](../decision-register.md#pre-11-b1): la plantilla integra requisitos, dependencias, pasos, porqués, contratos, ejemplos, pruebas y aceptación.
- [PRE-11.B2](../decision-register.md#pre-11-b2): el detalle permitirá implementar sin inventar arquitectura y sin descripción línea por línea.
- [PRE-11.B3](../decision-register.md#pre-11-b3): cada especificación incluirá ejemplo pertinente explicado cuando aporte claridad.
- [PRE-11.B4](../decision-register.md#pre-11-b4): la aceptación debe ser objetiva, verificable y vinculada a usos/fallos.
- [PRE-11.C1](../decision-register.md#pre-11-c1): coordinación asegura coherencia y delega solo trabajo independiente.
- [PRE-11.C2](../decision-register.md#pre-11-c2): arquitectura, motor, servicios, Three.js, seguridad, QA y operaciones se especializan según necesidad.
- [PRE-11.C3](../decision-register.md#pre-11-c3): cada archivo tiene un escritor; el coordinador integra índices y escritura paralela independiente.
- [PRE-11.C4](../decision-register.md#pre-11-c4): cada agente recibe instrucciones, alcance y decisiones/dependencias acotadas.
- [PRE-11.D1](../decision-register.md#pre-11-d1): auditoría, matriz, DAG, lotes y revisión de coherencia organizan la generación.
- [PRE-11.D2](../decision-register.md#pre-11-d2): el orden sigue dependencias técnicas, no necesariamente el orden numérico de dominios.
- [PRE-11.D3](../decision-register.md#pre-11-d3): cada lote requiere informe de cambios/riesgos y aprobación humana antes de seguir.
- [PRE-11.D4](../decision-register.md#pre-11-d4): enlaces, IDs, formato, secciones y referencias se auditan junto a revisión semántica.
- [PRE-11.E1](../decision-register.md#pre-11-e1): la primera ejecución produce solo Markdown, sin código, tests ejecutables, dependencias, infraestructura o despliegue.
- [PRE-11.E2](../decision-register.md#pre-11-e2): entregas verticales y prueba temprana Go/Vercel/Neon/Ably se planifican ahora y ejecutan después con autorización.
- [PRE-11.E3](../decision-register.md#pre-11-e3): cierre futuro de fase exige tests observados, aceptación, revisión, documentación y evidencias.
- [PRE-11.E4](../decision-register.md#pre-11-e4): cambios relevantes requieren propuesta, impacto, ADR y aprobación antes de alterar decisiones.
- [PRE-11.F1](../decision-register.md#pre-11-f1): README de portafolio mostrará arquitectura, retos, pruebas y demo prevista con estado real.
- [PRE-11.F2](../decision-register.md#pre-11-f2): trazabilidad relacionará requisito, spec, fase, componente, prueba, evidencia y estado.
- [PRE-11.F3](../decision-register.md#pre-11-f3): código, secretos, dependencias, despliegue y mutaciones Git necesitan autorización independiente de esta fase Markdown.
- [PRE-11.F4](../decision-register.md#pre-11-f4): el maestro gobierna inventario, lotes, checks y gates con informe final de creado/pendiente/bloqueado.

## Agrupaciones de diseño y consecuencias

**Autoridad y estado.** Una decisión fuente aprobada conserva naturaleza `APROBADO`; el documento que la sintetiza permanece `EN_REVISIÓN`. `POR_VERIFICAR` identifica incertidumbre de reglas/integración/capacidad y `PROPUESTA` identifica alternativas sin aprobación. Ninguna de estas etiquetas es resultado de prueba. Las instrucciones explícitas del propietario prevalecen y cualquier cambio normativo o arquitectónico necesita resolución identificable.

**Propiedad y contexto.** Delegar exige brief con rutas propias, dependencias, fuentes y restricciones. Los revisores comentan sin editar archivos cuyo autor sigue trabajando; coordinación integra índices, enlaces y trazabilidad. Las dependencias lineales y aprobaciones no se paralelizan. No basta repetir todos los documentos a cada agente: el contexto acotado conserva las decisiones pertinentes y la fuente canónica.

**Gates y evidencia.** Cada lote debe informar cambios reales, cobertura, comprobaciones, riesgos, dependencias y pendientes. Un enlace al índice con `PENDIENTE_LOTE_N` representa trabajo futuro y evita crear archivos vacíos para aparentar integridad. Las fases de implementación se planifican; no se registran como ejecutadas. La auditoría de enlaces no demuestra que un contrato o algoritmo funcione.

## Sustituciones y matices

No se simulan cuestionarios originales ni respuestas no conservadas: estos PRE sintetizan decisiones existentes. ADR no es permiso para cambiar una decisión del propietario. Los ejemplos de specs permanecen dentro de Markdown y solo se incluyen cuando aclaran; esta consolidación no necesita fabricar código de producto. Aprobar una fase documental no autoriza ramas, commits, PR, CI, instalación o proveedor alguno.

## Dependencias y fuentes canónicas

Las normas actuales son [11.01 — estándar y cambios](../specs/11-orchestration/11.01-document-standard-and-change-control.md) y [11.02 — ownership y contexto](../specs/11-orchestration/11.02-ownership-delegation-and-context.md). Consultar [AGENTS](../../AGENTS.md), [PLANS](../../PLANS.md), [índice](../index.md), [conflictos](../conflicts.md) y [trazabilidad](../traceability.md). Los restantes documentos del [dominio 11](../index.md#spec-11-01) y las fases conservan su lote según el inventario; no se crean para completar enlaces. Todos los PRE anteriores aportan requisitos a esta disciplina transversal.

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: completitud de cobertura y coherencia semántica final después de integrar autores; validez de enlaces/rutas debe revisarse con los archivos reales. Riesgos de contexto perdido, documentación que aparenta implementación, aprobación ambigua y requisitos huérfanos deben registrarse en [conflictos](../conflicts.md). Este autor solo escribe los once PRE asignados; coordinación mantiene registro, índices y estado global.

## Criterios futuros de aceptación y verificación

- **AC-01:** cada decisión tiene ID exacto y enlace al registro; la matriz identifica destino, estado, dependencia y evidencia pendiente sin duplicar autoridad normativa.
- **AC-02:** el reporte de lote corresponde a rutas realmente creadas/modificadas y distingue auditoría textual de pruebas de producto.
- **AC-03:** cada conflicto contiene discrepancia, impacto, alternativas y aprobación requerida; los algoritmos normativos afectados quedan bloqueados hasta resolución humana.
- **AC-04:** tras entregar lote 1 todos sus documentos siguen `EN_REVISIÓN`, se solicita literalmente «APRUEBO LOTE 1» y no comienza lote 2 automáticamente.

Las comprobaciones textuales efectivas se reportan al coordinador después de integrar; no se atribuyen aquí resultados aún no observados.

## Revisión humana

Revisar jerarquía, propiedad por archivo, control de cambios y carácter documental de la entrega. El propietario aprobará o devolverá observaciones mediante el gate del lote; este texto permanece `EN_REVISIÓN` hasta revisión explícita. La parada obligatoria no habilita un lote siguiente por silencio.
