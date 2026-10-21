# ADR-0001 — Tres servicios Go y propiedad de datos

ID: ADR-0001 · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner funcional: A (arquitectura).

Dominio: 01/05/07, arquitectura y propiedad de datos · Lote: 1. Formato y cambios: [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md); propiedad de escritura: [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md).

Decisión base: APROBADO EN FUENTE. Validación técnica: POR_VERIFICAR. Esta ADR documenta decisiones preliminares aprobadas; su contenido aún requiere revisión del propietario y no acredita implementación.

## Contexto y problema

La plataforma reúne identidad, partidas autoritativas e IA. Es necesario fijar quién puede cambiar cada dato antes de definir APIs, persistencia o despliegues. La distribución de servicios no permite asumir transacciones ACID que abarquen sus bases, ni compartir acceso SQL por conveniencia. El lanzamiento conserva MVP base y ampliado; la separación por dominios no reduce ese alcance.

## Fuentes y trazabilidad

- [PRE-05.A1](../decision-register.md#pre-05-a1) sustituye el monolito por exactamente Identity, Game y AI, independientes y en Go/Gin; [PRE-05.A2](../decision-register.md#pre-05-a2) delimita arquitectura hexagonal pragmática.
- [PRE-05.A4](../decision-register.md#pre-05-a4) establece módulos Go autónomos y cuatro proyectos; [PRE-05.D1](../decision-register.md#pre-05-d1) separa Identity DB y Game DB y deja AI sin persistencia inicial.
- [PRE-07.C4](../decision-register.md#pre-07-c4) exige obtener perfiles por API, sin acceder a Identity DB; [PRE-07.E4](../decision-register.md#pre-07-e4) asigna a Game resultado, Elo y agregados atómicos.
- [PRE-01.B3](../decision-register.md#pre-01-b3) sustituye Drizzle por sqlc, pgx/v5 y Goose; [PRE-11.E4](../decision-register.md#pre-11-e4) controla cambios arquitectónicos.

La copia íntegra se conserva en el [maestro, base arquitectónica y anexo A](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md). La [matriz de trazabilidad](../traceability.md) relaciona estas fuentes con futuras especificaciones y evidencias.

## Decisión delimitada

Se conservan exactamente tres microservicios Go/Gin con módulos y despliegues independientes. El frontend Next.js es un cuarto proyecto de despliegue, no un cuarto microservicio de dominio.

| Unidad | Responsabilidad y escritura autorizada | Límite |
| --- | --- | --- |
| Identity | Cuentas, invitados, sesiones, credenciales y perfiles de identidad en sus datos | No confirma movimientos, resultados ni actualiza Elo de Game |
| Game | Partidas, movimientos, resultados, Elo, estadísticas y agregados competitivos en sus datos | No consulta SQL de Identity ni delega su autoridad a AI |
| AI | Cálculo acotado de movimientos mediante el motor Go compartido | Inicialmente sin DB; una recomendación no es un movimiento confirmado |
| Web | Presentación, entrada y consumo de contratos | No escribe directamente bases ni decide reglas o resultados oficiales |

Se usan dos bases o esquemas de propiedad separada en Neon PostgreSQL, sin acceso SQL cruzado. La elección exacta entre bases o esquemas y los privilegios concretos quedan POR_VERIFICAR en su especificación. Cada servicio controla las migraciones de sus datos. Se conserva sqlc + pgx/v5 y Goose como herramientas previstas, sin ejecutar ni generar archivos técnicos en este lote.

Los servicios separan transporte, aplicación, dominio y adaptadores. Las interfaces pequeñas cercanas al consumidor y la inyección por constructores siguen [PRE-05.A3](../decision-register.md#pre-05-a3). Compartir el motor puro no habilita compartir repositorios SQL ni contratos internos de persistencia.

## Alternativas y motivación

**Histórica descartada en fuente:** el monolito de PRE-05.A1 fue sustituido expresamente; Drizzle fue sustituido por herramientas nativas de Go. No se presentan como opciones actuales.

**PROPUESTA, no adoptadas:** añadir un servicio de rankings, permitir un esquema SQL compartido sin propietarios o dar persistencia propia a AI. Requerirían justificar necesidad, coste, coordinación y revisión del propietario. No hay evidencia de que se hayan evaluado o aprobado.

La motivación trazable es la autonomía por dominio indicada por PRE-05.A1 y la propiedad separada de PRE-05.D1. Como consecuencia arquitectónica, los cambios de identidad y los de partidas se coordinan mediante contratos y estados observables; no se documenta una atomicidad distribuida inexistente.

## Contratos e invariantes

1. Game referencia usuarios por UUID y consulta perfiles públicos en lotes a Identity. Un correo o una credencial no se replica por conveniencia en contratos de partidas.
2. Game verifica la recomendación de AI contra su estado vigente y el motor, y conserva el control de versión y la confirmación transaccional.
3. Un resultado y su actualización competitiva pertenecen a una transacción de Game, conforme a [ADR-0008](ADR-0008-ranked-results-and-elo-ownership.md).
4. La eliminación de cuenta cruza propietarios mediante un flujo idempotente y verificable; [PRE-07.F3](../decision-register.md#pre-07-f3) no autoriza una transacción SQL entre bases.
5. Los contratos HTTP/eventos se versionan como establece [ADR-0003](ADR-0003-rest-contracts-and-versioned-events.md). El motor y el despliegue tienen límites propios en [ADR-0002](ADR-0002-single-versioned-go-engine.md) y [ADR-0006](ADR-0006-vercel-hobby-and-zero-cost-guardrails.md).

## Consecuencias, costes y dependencias

La separación exige coordinar contratos y compatibilidad de despliegues, autorización entre servicios y fallos parciales. Los perfiles o procesos distribuidos pueden no reflejar simultáneamente cada cambio; deben describirse sus estados y recuperación. No hay un coste operativo medido ni garantía de que las cuotas gratuitas soporten la carga prevista.

Depende de [01.01, stack y monorepo](../index.md#spec-01-01), PENDIENTE_LOTE_2; [05.01, límites de servicios](../index.md#spec-05-01), [05.03, datos/transacciones](../index.md#spec-05-03) y [07.03, perfiles y derechos de datos](../index.md#spec-07-03), PENDIENTE_LOTE_3. La eliminación distribuida necesita política de privacidad aplicable y evidencia futura. El empaquetado de módulos compartidos en Vercel continúa POR_VERIFICAR.

## Riesgos y validación futura

En una sesión de implementación autorizada deberán verificarse el aislamiento efectivo de credenciales, la ausencia de SQL cruzado, compatibilidad de contratos y fallos parciales de Identity/AI. Las carreras de finalización y los cambios concurrentes de rankings necesitan PostgreSQL real y evidencias observadas. No se ejecutan aquí esas pruebas.

## Aceptación futura

- **AC-01, positiva:** dada una consulta de perfiles desde Game, se obtiene únicamente el perfil público por API de Identity y queda evidencia de que Game carece de permisos SQL sobre sus datos.
- **AC-02, positiva:** dado un movimiento recomendado por AI, Game lo valida y confirma por su propia transacción; la evidencia muestra que AI no escribió el resultado.
- **AC-03, negativa:** un intento de lectura SQL cruzada o una credencial de AI intentando modificar partidas se rechaza sin cambiar datos.
- **AC-04, negativa:** una caída de Identity durante una operación distribuida no se presenta como commit atómico de ambos servicios; el estado pendiente y su recuperación son verificables.

Estos criterios están pendientes y no equivalen a verificaciones documentales de este lote.

Trazabilidad de aceptación: AC-01/AC-03 → [PRE-05.D1](../decision-register.md#pre-05-d1) y [PRE-07.C4](../decision-register.md#pre-07-c4), dueño canónico futuro [05.01](../index.md#spec-05-01), PENDIENTE_LOTE_3; AC-02 → [PRE-08.F3](../decision-register.md#pre-08-f3), dueño [08.02](../index.md#spec-08-02), PENDIENTE_LOTE_4; AC-04 → [PRE-07.F3](../decision-register.md#pre-07-f3), dueño [07.03](../index.md#spec-07-03), PENDIENTE_LOTE_3. Riesgo asociado: [RISK-09](../conflicts.md#risk-09).

## Cambio y reversión documental

Un cuarto servicio, una base compartida sin propiedad separada o trasladar Elo a Identity requiere propuesta de cambio, impacto en fuentes/contratos/costes y aprobación expresa. Registrar la discrepancia en [conflictos](../conflicts.md) y actualizar la trazabilidad antes de sustituir esta decisión. Una revisión rechazada se revierte documentalmente a la última versión aprobada, conservando el historial y la fuente; no se autoriza rollback de sistemas ni migraciones en esta fase.
