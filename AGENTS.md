# AGENTS — Toguz Kumalak

ID: AGENTS · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner: coordinación.

## Alcance autorizado

El propietario aprobó el inventario con «APRUEBO INVENTARIO DOCUMENTAL» y ordenó generar únicamente el lote 1 y detenerse. Esta autorización permite los 31 Markdown del lote 1 del [inventario](docs/index.md#inventario). No autoriza los lotes 2–6, implementación, pruebas de producto ni infraestructura.

Lee primero el [estado y los gates](PLANS.md), el [registro](docs/decision-register.md), los [conflictos](docs/conflicts.md) y las decisiones pertinentes. El [maestro archivado](docs/sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) conserva la fuente íntegra.

### Autorización posterior de publicación — 2026-10-21

El propietario autorizó expresamente la inicialización/publicación en su cuenta personal de GitHub del repositorio público `toguz-kumalak`: crear `.gitignore`, configurar Git localmente, establecer `main`, hacer commits y push de documentación existente, y abrir un PR desde `docs/specs-lote-01` sin fusionarlo. Esta instrucción posterior sustituye las prohibiciones anteriores exclusivamente para esas operaciones de publicación. No aprueba el contenido EN_REVISIÓN, nuevos lotes, implementación, servicios, dependencias, licencia ni credenciales de producción. La autenticación debe verificarse y los archivos deben revisarse antes de publicar; no incluir secretos ni información privada.

## Autoridad

1. Instrucciones explícitas del propietario y aclaraciones aprobadas.
2. Registro de decisiones y ADR aprobadas por el propietario.
3. Especificaciones finales aprobadas.
4. Estas instrucciones operativas.
5. Planes y briefs de ejecución.

La aprobación de un inventario no aprueba el contenido todavía no revisado. Distingue estado documental `BORRADOR / EN_REVISIÓN / APROBADO` de naturaleza de una afirmación `APROBADO / POR_VERIFICAR / PROPUESTA`. Una decisión aprobada no acredita una integración probada.

## Restricciones actuales

- Crear o modificar únicamente rutas `.md` expresamente autorizadas. Los ejemplos permanecen dentro de bloques Markdown.
- No crear código, manifiestos, configuraciones, assets, secretos ni verificadores ejecutables en disco. Los análisis de texto pueden ejecutarse en memoria sin instalar herramientas.
- No instalar, iniciar servicios, ejecutar pruebas de producto, migrar, emitir tokens, conectar proveedores, provisionar ni desplegar.
- No crear ramas, commits, pushes o PR; no modificar GitHub ni el índice/configuración Git. Usar consultas Git de solo lectura sin locks opcionales.
- No añadir un cuarto microservicio, motor TypeScript, infraestructura obligatoria Redis/Socket.IO ni costes sin nueva aprobación.
- Mantener MVP base y ampliado en el lanzamiento comprometido. Local no implica offline.
- No declarar FPS, fuerza de IA, latencia, costes, restauración o despliegue como medidos sin evidencia observada.

## Trabajo documental

- Sigue [11.01](https://github.com/SebasChaconDev/toguz-kumalak/blob/docs/specs-lote-01/docs/specs/11-orchestration/11.01-document-standard-and-change-control.md) y [11.02](https://github.com/SebasChaconDev/toguz-kumalak/blob/docs/specs-lote-01/docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md).
- Cita identificadores exactos `PRE-NN.Xn`; conserva el texto canónico del registro y la copia del maestro.
- Un autor por archivo; los revisores no escriben en archivos ajenos mientras su autor trabaja. El coordinador integra índices y trazabilidad.
- No crear archivos vacíos de lotes futuros para resolver enlaces. Enlazar su entrada existente en el índice y marcar `PENDIENTE_LOTE_N`.
- Registrar discrepancia, impacto, alternativas y aprobación requerida en conflictos; no convertir preferencias en requisitos.
- Ante discrepancia normativa del juego, preservar la decisión fuente y bloquear el cierre del algoritmo afectado hasta resolución humana.
- Comprobar rutas, IDs, enlaces, fences, cobertura y estado Git. Esas comprobaciones son documentales y no validan el producto.

## Parada obligatoria

Al terminar el lote 1, informar archivos creados, decisiones cubiertas, verificaciones realizadas, dependencias, riesgos y pendientes. Dejar documentos `EN_REVISIÓN` y solicitar «APRUEBO LOTE 1». Detenerse hasta nueva instrucción; no comenzar automáticamente el lote 2.
