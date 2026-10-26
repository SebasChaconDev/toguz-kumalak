# AGENTS — Toguz Kumalak

> **Estado vigente — 2026-10-26: LOTE_6_APROBADO_DOCUMENTALMENTE.** El propietario aprobó AUD-01 como auditoría documental final sobre HEAD 6052689ff9b15a19214267d7cc6a614d17f8fcdd; PR #6 integrado por e81f01a0e94c3037e3b922894202b6eee80534e8. [Recibo de cierre](PLANS.md#cierre-lote-6) prevalece sobre estados y paradas históricos. AUD6-F01, AUD6-F02 y AUD6-C01–C06 permanecen pendientes para sus fases. No se autoriza implementación, ejecución de PH00/PH01, workflows o despliegues.

> **Estado vigente — 2026-10-26: LOTE_5_APROBADO_DOCUMENTALMENTE.** PR #5 integrado en main por 6579ca930a9081f773b390ce61f6bc80cc6b1ae7, desde HEAD aprobado 9ff1a2857d0b91b9cf7925a4c36de30e377cef3d. [Recibo de aprobación e integración](PLANS.md#cierre-lote-5) prevalece sobre cortes históricos: veinte documentos del lote5 aprobados y246decisiones con especificación redactada y aprobada documentalmente. L4-P01, retención y demás propuestas siguen pendientes de aprobación específica. Detenerse: sin producto, workflows, proveedores/despliegues o lote6.

> **Autorización posterior vigente — 2026-10-26: LOTE_5_EN_REVISIÓN.** El propietario autoriza los20Markdown inventariados del lote5, siete registros mínimos y publicación desde `docs/specs-lote-05` con PR sin fusionar. [Alcance y ledger](PLANS.md#ejecucion-lote-5) prevalece sobre paradas anteriores. No producto, workflows ejecutables, instalaciones, tests de producto, proveedores/despliegues, pagos o lote6. L4-P01 sigue PROPUESTA pendiente de aprobación de contratos.

> **Estado vigente — 2026-10-26: LOTE_4_APROBADO_DOCUMENTALMENTE.** PR #4 integrado en main por `d5dccfbe0a55c78ddc7c2c0396ead9fd2bfad910`, desde el HEAD aprobado `3ef3f7cbfd3bfc4aea837d9f1b0b414be064e8ee`. [Recibo de aprobación e integración](PLANS.md#cierre-lote-4) prevalece sobre los cortes históricos. L4-P01 sigue PROPUESTA pendiente de aprobación de contratos antes de implementación. Detenerse: no producto, despliegues o lote 5.

> **Estado vigente — 2026-10-25: READY_FOR_LOTE_4_DOCUMENTATION.** PR #3 integrado en main mediante c50ba9f; GO_REMOTE experimental cerrado y recursos remotos propios retirados. [Recibo verificado](PLANS.md#lote-3-integrado-verificado) prevalece sobre los cortes históricos siguientes. El lote 4 aún no se ha generado; el producto no está implementado.

> Autorización posterior vigente: el propietario autoriza generar y publicar exclusivamente el lote 3 documental en docs/specs-lote-03: once specs EN_REVISIÓN, registros necesarios y PR contra main sin fusionar. Esta instrucción sustituye las paradas históricas de lotes anteriores; no autoriza producto, lote 4, infraestructura o pagos.

> Aprobación posterior del lote 2: el propietario declaró «APRUEBO LOTE 2» y autorizó fusionar PR #2 en main. Contenido aprobado: HEAD 598ac9a4fc6ce14cf512785fb1c4160fafb4eb9e; integrado por squash 1f401af8d76794981bbcfdf9a3aa1831cf2621f5 (2026-10-21T21:10:14Z). Este registro sustituye las menciones históricas a revisión pendiente o prohibición de esa fusión. No autoriza lote 3 ni implementación; el producto sigue sin verificar.

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: AGENTS · Versión: 0.9 · Fecha: 2026-10-26 · Estado documental: APROBADO · Owner: coordinación.

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
<a id="agentes-continuidad-lote-4"></a>

## Alcance posterior — preparación de lote4, 2026-10-25

La autorización posterior del propietario permite cerrar e integrar documentalmente lote3 y registrar GO_REMOTE experimental para preparar lote4. [PLANS](PLANS.md#preparacion-lote-4) conserva la cita íntegra y la interpretación operativa; no inventar una frase de aprobación literal ni una prueba de producto.

Se permiten únicamente las correcciones documentales menores, estados, trazabilidad, integración del PR3 y recibo de main necesarios para ese cierre. Código experimental queda en spike/infra-feasibility-01, no se integra en main. Concluida la verificación, detenerse y avisar: no generar las ocho specs del lote4 ni iniciar código de producción en este turno.

Al trabajar posteriormente, usar los permisos más recientes y preservar246 PRE, AD-NORM-001, tres microservicios Go y PRE-09. No reabrir reglas ni convertir propuestas en decisiones canónicas por cambiar su estado documental.

## Alcance posterior — cierre del lote4, 2026-10-26

El propietario declaró «APRUEBO LOTE 4 como documentación técnica.» y autorizó integrar PR #4 únicamente tras comprobar su HEAD `3ef3f7cbfd3bfc4aea837d9f1b0b414be064e8ee`. Condición cumplida; squash `d5dccfbe0a55c78ddc7c2c0396ead9fd2bfad910` integrado y árbol idéntico al revisado. Las ocho specs están APROBADAS documentalmente. Se permiten solo estados, trazabilidad y registros de este cierre; al verificarlos/publicarlos, detenerse.

L4-P01 sigue PROPUESTA, pendiente de aprobación de contratos antes de implementarla; la aprobación documental no modifica05.02/05.03/07.02 ni otras decisiones. No implementar producto, desplegar, iniciar lote5 o reabrir reglas. [Recibo y autoridad posterior](PLANS.md#cierre-lote-4) prevalece sobre las paradas históricas anteriores.

## Alcance posterior — cierre aprobado del lote 5, 2026-10-26

El propietario declaró «APRUEBO LOTE 5 como documentación técnica.» y autorizó verificar HEAD 9ff1a2857d0b91b9cf7925a4c36de30e377cef3d, ausencia de cambios posteriores sin revisar y comprobaciones obligatorias fallidas; si coincide, fusionar PR #5 en main y actualizar únicamente estados, trazabilidad y registros necesarios. Condición cumplida: squash 6579ca930a9081f773b390ce61f6bc80cc6b1ae7, 2026-10-26T18:25:50Z, árbol integrado idéntico al revisado y27rutas exactas(20altas/7registros). [Recibo](PLANS.md#cierre-lote-5) prevalece sobre todas las paradas históricas.

Los20documentos del lote5 pasan a v0.2/APROBADO documentalmente sin cambios a cuerpos técnicos; las40specs cubren246decisiones aprobadas, cero en revisión o sin redactar. REG añade solo AD-GATE-005;246PRE y adendas anteriores intactas. L4-P01, política de retención y demás propuestas técnicas mantienen aprobación específica pendiente. No implementar, instalar workflows, desplegar proveedores o iniciar lote6. Trabajo local ajeno se preserva en la carpeta original y no se publica con este cierre. Tras verificar main/estados y publicar el cierre, detenerse.

## Alcance posterior — cierre aprobado del lote 6, 2026-10-26

El propietario declaró «APRUEBO LOTE 6 como auditoría documental final, con resultado READY_WITH_SCOPED_CONDITIONS.» y autorizó fusionar PR #6 únicamente tras verificar HEAD 6052689ff9b15a19214267d7cc6a614d17f8fcdd, ausencia de cambios posteriores sin revisar y comprobaciones obligatorias fallidas. Condición cumplida: squash e81f01a0e94c3037e3b922894202b6eee80534e8,2026-10-26T19:39:26Z, árbol idéntico al aprobado e integración exacta deAUD-01 y cuatro registros. [Recibo](PLANS.md#cierre-lote-6).

Se permite solo cierre de estados/trazabilidad/recibos y estados operativos mínimosAGENTS/README; las40specs,REG,246PRE,ADR,TK-N1-1.0.0 y14ACC no cambian. AUD6-F01/F02 yAUD6-C01–C06 permanecen pendientes, no resueltos por aprobar la auditoría. L4-P01 yretención mantienen aprobación específica. PH00/PH01 están preparados para solicitar alcance posterior, sin ejecución autorizada. Tras verificar/publicar main, detenerse: no código,instalaciones,workflows,proveedores odespliegues.
