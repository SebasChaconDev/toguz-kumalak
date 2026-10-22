# AGENTS — Toguz Kumalak

> Autorización posterior vigente: el propietario autoriza generar y publicar exclusivamente el lote 3 documental en docs/specs-lote-03: once specs EN_REVISIÓN, registros necesarios y PR contra main sin fusionar. Esta instrucción sustituye las paradas históricas de lotes anteriores; no autoriza producto, lote 4, infraestructura o pagos.

> Aprobación posterior del lote 2: el propietario declaró «APRUEBO LOTE 2» y autorizó fusionar PR #2 en main. Contenido aprobado: HEAD 598ac9a4fc6ce14cf512785fb1c4160fafb4eb9e; integrado por squash 1f401af8d76794981bbcfdf9a3aa1831cf2621f5 (2026-10-21T21:10:14Z). Este registro sustituye las menciones históricas a revisión pendiente o prohibición de esa fusión. No autoriza lote 3 ni implementación; el producto sigue sin verificar.

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: AGENTS · Versión: 0.4 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner: coordinación.

## Alcance autorizado

Autorización actual: «AUTORIZO LA GENERACIÓN Y PUBLICACIÓN DOCUMENTAL DEL LOTE 3», según solicitud íntegra adjunta del propietario. Crear únicamente las once specs 05.01–05.04, 07.01–07.04 y 06.01–06.03 del inventario; actualizar IDX/TRZ/CONF/PLANS y este registro operativo mínimo. Se permiten rama docs/specs-lote-03 basada en main, commits, push y PR, sin force push ni fusión. Las restricciones de cierre que siguen son historia de lotes 1 y 2. Al publicar, pedir «APRUEBO LOTE 3» y detenerse. Ningún permiso para implementar, instalar, probar producto, conectar proveedores, desplegar o generar lote 4.

El propietario aprobó el inventario y después declaró «APRUEBO LOTE 1», autorizando integrar PR #1 únicamente si su HEAD coincidía con e38cc84950858e072c1a3f06ffbf854cf7e544f5. La comprobación se cumplió y el squash en main es 329b7c1059e8a74f191c256870734e8ae8259966. También autorizó el lote 2: exactamente diez specs de los dominios 01, 02 y 03 del [inventario](docs/index.md#inventario), y las actualizaciones necesarias de índice, trazabilidad, riesgos y planes. La instrucción posterior «APRUEBO LOTE 2» autorizó fusionar el PR #2 en main, verificar la integración y actualizar los estados documentales. La fusión ya está realizada; solo procede cerrar estos registros y detenerse. No generar lote 3, implementar, instalar ni desplegar.

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
- Git y GitHub se limitan al cierre autorizado: PR #1 y PR #2 integrados, verificación de main y publicación de estados de aprobación. No force push, no eludir protecciones ni operaciones ajenas al alcance. Terminados esos registros, detenerse.
- No añadir un cuarto microservicio, motor TypeScript, infraestructura obligatoria Redis/Socket.IO ni costes sin nueva aprobación.
- Mantener MVP base y ampliado en el lanzamiento comprometido. Local no implica offline.
- No declarar FPS, fuerza de IA, latencia, costes, restauración o despliegue como medidos sin evidencia observada.

## Trabajo documental

- Sigue [11.01](docs/specs/11-orchestration/11.01-document-standard-and-change-control.md) y [11.02](docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md).
- Cita identificadores exactos `PRE-NN.Xn`; conserva el texto canónico del registro y la copia del maestro.
- Un autor por archivo; los revisores no escriben en archivos ajenos mientras su autor trabaja. El coordinador integra índices y trazabilidad.
- No crear archivos vacíos de lotes futuros para resolver enlaces. Enlazar su entrada existente en el índice y marcar `PENDIENTE_LOTE_N`.
- Registrar discrepancia, impacto, alternativas y aprobación requerida en conflictos; no convertir preferencias en requisitos.
- Ante discrepancia normativa del juego, preservar la decisión fuente y bloquear el cierre del algoritmo afectado hasta resolución humana.
- Comprobar rutas, IDs, enlaces, fences, cobertura y estado Git. Esas comprobaciones son documentales y no validan el producto.

## Parada obligatoria

El propietario aprobó expresamente el lote 2 y su integración. Las diez specs y sus registros están APROBADOS documentalmente; las pruebas de producto permanecen pendientes. Tras verificar main y publicar los estados, informar el cierre y detenerse. No iniciar lote 3 ni implementar código sin nuevas instrucciones.
