# PRE-07 — Identidad, permisos, datos y Elo

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: PRE-07 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner: S — servicios.

## Objetivo y contexto

Consolidar las 24 decisiones que separan identidad, autorización, datos públicos, competición y privacidad. Una sesión verificable permite acceso, pero cada operación sigue necesitando permiso por recurso. Este documento no acredita una integración Google, gestión de claves o cumplimiento jurídico verificados.

## Fuentes y decisiones

El [maestro archivado](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) y el [registro canónico](../decision-register.md) preservan las formulaciones. Síntesis de naturaleza `APROBADO`:

- [PRE-07.A1](../decision-register.md#pre-07-a1): Google OIDC con Authorization Code/PKCE habilitará registro, sin contraseñas propias en MVP.
- [PRE-07.A2](../decision-register.md#pre-07-a2): Identity emitirá UUID, alias temporal y sesión invitada comprobable/recuperable.
- [PRE-07.A3](../decision-register.md#pre-07-a3): vincular invitado puede incorporar historial elegible, sin convertirlo retroactivamente en Elo.
- [PRE-07.A4](../decision-register.md#pre-07-a4): UUID es inmutable; alias registrado es único y el temporal invitado puede repetirse.
- [PRE-07.B1](../decision-register.md#pre-07-b1): access tokens cortos asimétricos/JWKS validarán emisor y audiencia en Game sin consulta por movimiento.
- [PRE-07.B2](../decision-register.md#pre-07-b2): acceso aproximado de 15 minutos y sesiones renovables deberán permitir recuperación invitada de 24 h.
- [PRE-07.B3](../decision-register.md#pre-07-b3): refresh rota y se revoca al salir; access tokens emitidos conservan su ventana corta salvo mecanismo extraordinario.
- [PRE-07.B4](../decision-register.md#pre-07-b4): cookies HttpOnly/Secure/SameSite adecuadas, mismo origen, defensa CSRF y verificación de origen excluyen tokens sensibles en localStorage.
- [PRE-07.C1](../decision-register.md#pre-07-c1): guest/user/admin son roles mínimos y requieren autorización adicional por recurso.
- [PRE-07.C2](../decision-register.md#pre-07-c2): Game revisará identidad, participación, turno, estado y permiso en cada operación.
- [PRE-07.C3](../decision-register.md#pre-07-c3): llamadas entre servicios usan credenciales propias cortas, audiencia/permisos, HTTPS y rotación.
- [PRE-07.C4](../decision-register.md#pre-07-c4): Game referencia UUID y solicita perfiles públicos por API agrupada, sin leer Identity DB.
- [PRE-07.D1](../decision-register.md#pre-07-d1): alias, Elo y estadísticas forman el perfil público; correo y privados quedan ocultos.
- [PRE-07.D2](../decision-register.md#pre-07-d2): se admite editar alias y avatar generado, sin cargas arbitrarias de imágenes.
- [PRE-07.D3](../decision-register.md#pre-07-d3): alias necesita validación, nombres reservados y moderación básica contra abuso.
- [PRE-07.D4](../decision-register.md#pre-07-d4): estadísticas/Elo son públicos, mientras historial propio y reproducción compartida requieren acceso controlado.
- [PRE-07.E1](../decision-register.md#pre-07-e1): Elo inicia en 1200 con K=32 y diez clasificatorias provisionales; parámetros requieren simulación/validación.
- [PRE-07.E2](../decision-register.md#pre-07-e2): estadísticas incluyen resultados, tasa, evolución Elo, tiempos por turno, capturas, tuzdyks y rendimiento por color/modalidad.
- [PRE-07.E3](../decision-register.md#pre-07-e3): ranking ofrecerá paginación, búsqueda por alias, posición personal y marca provisional.
- [PRE-07.E4](../decision-register.md#pre-07-e4): resultado, Elo y agregados se actualizan atómicamente e idempotentemente en Game.
- [PRE-07.F1](../decision-register.md#pre-07-f1): caducidad/retención invitada se limita sin impedir recuperación de 24 h; limpieza posterior se documenta.
- [PRE-07.F2](../decision-register.md#pre-07-f2): cada usuario podrá exportar perfil, historial y estadísticas propios en JSON estructurado.
- [PRE-07.F3](../decision-register.md#pre-07-f3): borrado distribuido bloquea/revoca, solicita borrado/anonimización Game, verifica y finaliza idempotentemente según privacidad aplicable.
- [PRE-07.F4](../decision-register.md#pre-07-f4): pruebas futuras cubren tokens, cuentas, invitados, vinculación, acceso cruzado, borrado y rankings manipulados.

## Agrupaciones de diseño y consecuencias

**Cuatro autoridades de credenciales.** Google acredita identidad externa; Identity gestiona la sesión propia/JWKS; las credenciales entre servicios habilitan acciones internas; Ably limita capacidades de canal. No son intercambiables. El cierre de sesión revoca renovación, pero no garantiza invalidación instantánea de cada access token ya emitido. Se deben especificar rotación, caché JWKS y comportamiento ante claves desconocidas sin inventar disponibilidad continua.

**Datos e integridad.** Game conserva partidas, resultados, Elo y agregados en su transacción; Identity posee cuentas y perfiles. La vista pública puede combinar ambos mediante API sin SQL cruzado. El alias visible no es identificador estable ni prueba de participación. La vinculación de invitado necesita definir colisiones y trazabilidad sin otorgar Elo retrospectivo.

**Privacidad y ciclo de vida.** Exportación y eliminación atraviesan fronteras de datos sin transacción ACID distribuida. El borrado requiere una secuencia recuperable con comprobación de finalización; una solicitud enviada no acredita borrado concluido. La retención invitada debe respetar recuperación, minimizar datos y explicitar tratamiento de partidas compartidas, backups y obligaciones aplicables antes de fijar plazos.

## Sustituciones y matices

No se introducen contraseñas propias, uploads de avatar o autenticación por alias. El acceso de unos 15 minutos es aproximado, no sustituto de la política completa de sesión. Tiempo por turno es estadística y no reloj competitivo. [PRE-06.D1](../decision-register.md#pre-06-d1) limita Elo a ranked entre dos registrados; práctica, IA, invitado y archivo inconcluso quedan fuera.

## Dependencias y fuentes canónicas

El [dominio 07 del índice](../index.md#spec-07-01) es `PENDIENTE_LOTE_3`. [PRE-05](PRE-05.md) gobierna contratos/propiedad, [PRE-06](PRE-06.md) elegibilidad/recuperación y [PRE-10](PRE-10.md) secretos/backups/entornos. La revisión se rige por [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [trazabilidad](../traceability.md).

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: compatibilidad OAuth/cookies entre entornos, recuperación invitada, rotación y revocación, parámetros Elo, política de moderación, retención y privacidad/GDPR aplicable. No se afirma cumplimiento legal ni seguridad probada. Registrar alternativas/impactos en [conflictos](../conflicts.md). No crear cuentas, tokens, claves, bases ni conexiones en este lote.

## Criterios futuros de aceptación y verificación

- **AC-01:** matriz de credenciales identifica emisor, audiencia, destinatario, vida, renovación, permisos y fallo seguro para las cuatro autoridades.
- **AC-02:** escenarios de acceso cruzado, alias duplicado, sesión expirada, rotación y vinculación no permiten apropiarse de una partida o dato privado.
- **AC-03:** resultados concurrentes/reintentados cambian Elo/agregados una sola vez; casos no elegibles mantienen Elo y el provisional se identifica.
- **AC-04:** exportación está limitada al titular y borrado distribuido documenta reintento, verificación, datos retenidos y evidencia antes de considerarse finalizado.

Son condiciones futuras de documentación y producto; aquí solo se revisa texto y cobertura.

## Revisión humana

Validar ciclo de sesiones, privacidad, recuperación y separación de estadísticas/competición. Mantener `EN_REVISIÓN`; parámetros aprobados no equivalen a calibración validada. El coordinador solicitará «APRUEBO LOTE 1» antes de detenerse.
