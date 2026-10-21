# ADR-0005 — Fronteras de confianza de sesión y servicios

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: ADR-0005 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner funcional: S (servicios y seguridad).

Dominio: 07/06, identidad y confianza · Lote: 1. Formato y cambios: [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md); propiedad de escritura: [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md).

Decisión base: APROBADO EN FUENTE. Validación técnica: POR_VERIFICAR. El documento no acredita OAuth, sesiones, rotación, JWKS ni capacidades Ably funcionando.

## Contexto y problema

Hay cuatro autoridades distintas: Google autentica una identidad externa, Identity mantiene la identidad/sesión de la plataforma, los servicios se autentican entre sí y Ably limita acceso a canales. Confundir credenciales permitiría usar una prueba de login como permiso de juego, dar privilegios de publicación al navegador o aceptar credenciales de servicio como si representaran a un participante.

Los invitados también necesitan identidad verificable y recuperable: conocer un alias o un enlace de sala no basta para retomar la partida de otro usuario. La autorización depende del recurso y de su estado, además del rol.

## Fuentes y trazabilidad

- [PRE-07.A1](../decision-register.md#pre-07-a1) fija Google OIDC/Authorization Code + PKCE; [PRE-07.A2](../decision-register.md#pre-07-a2) fija invitados con UUID y sesión verificable.
- [PRE-07.B1](../decision-register.md#pre-07-b1) define tokens cortos asimétricos/JWKS y validación local de Game; [PRE-07.B2](../decision-register.md#pre-07-b2) establece acceso aproximadamente 15 min y recuperación invitada compatible con 24 h.
- [PRE-07.B3](../decision-register.md#pre-07-b3) define refresh rotatorio, revocación y límite de access tokens ya emitidos; [PRE-07.B4](../decision-register.md#pre-07-b4) exige cookies seguras, mismo origen, CSRF y no guardar secretos en localStorage.
- [PRE-07.C1](../decision-register.md#pre-07-c1), [PRE-07.C2](../decision-register.md#pre-07-c2) y [PRE-07.C3](../decision-register.md#pre-07-c3) separan roles, recurso e identidad de servicio.
- [PRE-05.E3](../decision-register.md#pre-05-e3), [PRE-06.B4](../decision-register.md#pre-06-b4) y [PRE-06.E2](../decision-register.md#pre-06-e2) delimitan Ably y fallos de Identity.

El [maestro y sus controles de tokens](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) conserva la fuente. La [trazabilidad](../traceability.md) vincula estas decisiones con especificaciones y pruebas pendientes.

## Decisión delimitada

| Credencial/autoridad | Uso aprobado | Frontera que debe preservar |
| --- | --- | --- |
| Google OIDC | Registro/login externo por Authorization Code + PKCE | No sustituye la sesión propia ni autoriza por sí sola una jugada |
| Sesión y JWT de acceso propio de Identity | Representar al usuario registrado o invitado; firma asimétrica y JWKS público | Game valida emisor, audiencia, caducidad y firma, además del recurso |
| Credencial servicio-servicio | Llamada HTTPS con vida corta, audiencia y permisos propios | No representa automáticamente a un usuario ni hereda su rol |
| Token temporal Ably | Suscribir al canal privado de una partida con capacidad restringida | No permite publicar movimientos oficiales ni leer otras partidas |

Identity mantiene cuentas, UUID, invitados, sesiones y refresh. Game verifica el token propio sin consultar Identity en cada jugada y valida participante, turno, estado y permiso en cada operación. Los roles mínimos son guest/user/admin, sin que un rol sustituya la comprobación del recurso.

Los JWT propios de acceso son distintos de la credencial Google. Los claims, algoritmos admitidos, política de claves, publicación/cache JWKS y validación de rotación deberán detallarse sin mezclar autoridades. Esta ADR no elige algoritmo, tolerancia de reloj, duración exacta de cache ni formato de credenciales entre servicios.

El navegador usa cookies HttpOnly/Secure/SameSite adecuadas y controles de origen/CSRF según el flujo de mismo origen aprobado. No conserva tokens sensibles en localStorage. Web y APIs Identity/Game se exponen bajo mismo origen mediante rewrites; AI mantiene endpoint HTTPS protegido para servicios, conforme a [PRE-10.A3](../decision-register.md#pre-10-a3). Los atributos y alcance exactos de cookie quedan pendientes de diseño verificable, no se fijan arbitrariamente aquí.

## Alternativas y motivación

**Histórica descartada en fuente:** no se documenta una migración desde un sistema de contraseñas existente. PRE-07.A1 elige Google sin contraseñas propias en MVP; no se inventa una alternativa histórica ya implementada.

**PROPUESTA, no adoptadas:** una credencial universal para usuario/servicio/Ably, introspección obligatoria de Identity en cada movimiento o revocación central inmediata de todo token emitido. Cambiarían fronteras, disponibilidad o infraestructura y requerirían propuesta y aprobación. Guardar secretos en localStorage y publicar desde clientes contradice directamente las decisiones fuente.

La motivación trazable es la separación de identidad de servicio y usuario de PRE-07.C3, combinada con la autonomía de Game de PRE-07.B1. La disponibilidad durante una caída de Identity está limitada por la vigencia de tokens y material de validación; no es una garantía de juego indefinido sin Identity.

## Invariantes y casos límite

1. Una invitación permite solicitar acceso a una sala bajo sus reglas; no acredita ser el participante original. La recuperación exige su identidad de sesión según [PRE-06.C1](../decision-register.md#pre-06-c1).
2. Al salir se revoca/rota según el registro de sesiones; un access token ya emitido puede seguir válido hasta su expiración corta, salvo mecanismos extraordinarios. No se promete revocación instantánea global por cerrar sesión.
3. Una caída de Identity limita altas y renovaciones. El comportamiento ante claves desconocidas, expiradas o no disponibles debe fallar de forma segura según la política futura, sin aceptar tokens sin validar.
4. Vincular un invitado preserva la identidad/historial según reglas aprobadas y no convierte partidas previas en Elo, como documenta [ADR-0008](ADR-0008-ranked-results-and-elo-ownership.md).
5. El permiso de suscripción Ably se limita al canal autorizado. Game conserva la publicación oficial y REST vuelve a autorizar el acceso al estado.
6. Datos públicos no incluyen correo ni datos privados. Logs y trazas no contienen tokens, secretos o material de sesión, conforme a [PRE-07.D1](../decision-register.md#pre-07-d1) y [PRE-10.B2](../decision-register.md#pre-10-b2).

## Consecuencias, dependencias y riesgos

Se requieren gestión de sesiones, rotación de claves y secretos por entorno, más contratos explícitos para cada audiencia. Los fallos de cache JWKS, renovación o capacidad Ably tienen efectos distintos; no deben ocultarse como un único estado conectado. La compatibilidad de invitados con recuperación 24 h exige definir caducidad/retención y renovación conjuntamente, sin prolongar indiscriminadamente access tokens.

Depende de [07.01, ciclo de sesión](../index.md#spec-07-01), [07.02, tokens/autorización](../index.md#spec-07-02), [06.01, salas](../index.md#spec-06-01) y [05.04, permisos de eventos](../index.md#spec-05-04), PENDIENTE_LOTE_3; [09.01, verificación de identidad](../index.md#spec-09-01) y [10.01, origen/secretos](../index.md#spec-10-01), PENDIENTE_LOTE_5. [ADR-0001](ADR-0001-three-go-services-and-data-ownership.md) separa datos; [ADR-0003](ADR-0003-rest-contracts-and-versioned-events.md) separa contratos. Precios, soporte y configuración vigente de Google/Ably/Vercel permanecen POR_VERIFICAR mediante documentación oficial y pruebas futuras autorizadas.

## Validación y aceptación futura

Las pruebas deberán incluir firma, emisor, audiencia, expiración, rotación JWKS, refresh reutilizado, CSRF, origen, acceso cruzado y privilegios de servicio. La política de cache/rotación debe demostrar qué ocurre durante caídas de Identity. No se emiten credenciales ni se conecta un proveedor en este lote.

- **AC-01, positiva:** un participante con sesión válida actúa sobre su partida y Game verifica token/recurso sin consultar Identity por movimiento; existe evidencia de las comprobaciones.
- **AC-02, positiva:** un invitado recupera su partida dentro de su vigencia mediante la misma identidad verificable, con estado confirmado de Game.
- **AC-03, negativa:** token Google presentado como access token propio, audiencia errónea, identidad de servicio sin permiso o token Ably presentado al API son rechazados sin efectos.
- **AC-04, negativa:** un usuario válido ajeno a la partida o con turno incorrecto no mueve ni obtiene suscripción privada no autorizada.
- **AC-05, negativa:** una sesión cerrada no renueva; el comportamiento del access token anterior coincide con su expiración documentada, sin afirmar una revocación instantánea inexistente.

Trazabilidad de aceptación: AC-01/AC-03/AC-04 → [PRE-07.B1](../decision-register.md#pre-07-b1), [PRE-07.C2](../decision-register.md#pre-07-c2) y [PRE-07.C3](../decision-register.md#pre-07-c3), dueño [07.02](../index.md#spec-07-02); AC-02 → [PRE-06.C1](../decision-register.md#pre-06-c1), dueño [06.03](../index.md#spec-06-03); AC-05 → [PRE-07.B3](../decision-register.md#pre-07-b3), dueño [07.01](../index.md#spec-07-01); todos PENDIENTE_LOTE_3. Riesgos asociados: [RISK-05](../conflicts.md#risk-05) y [RISK-06](../conflicts.md#risk-06).

## Cambio y reversión documental

Cambiar proveedor, audiencia, modelo de sesión o garantía de revocación requiere propuesta con impacto de seguridad/disponibilidad/recuperación y aprobación. Registrar discrepancias en [conflictos](../conflicts.md). La reversión documental conserva fuentes y última versión aprobada; una futura rotación o revocación operativa necesita procedimiento propio y no se ejecuta aquí.
