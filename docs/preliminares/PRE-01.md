# PRE-01 — Infraestructura y arquitectura

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: PRE-01 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner: A — arquitectura.

## Objetivo y contexto

Consolidar las elecciones de plataforma y sus sustituciones para que las futuras especificaciones partan de una arquitectura común. Este preliminar interpreta las 17 decisiones aprobadas; no acredita que un monorepositorio, runtime, integración o despliegue exista o funcione.

## Fuentes y decisiones

La fuente íntegra es el [maestro archivado](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md), Anexo A. El [registro canónico](../decision-register.md) conserva la formulación aprobada; estas implicaciones son una síntesis documental de naturaleza `APROBADO`:

- [PRE-01.A1](../decision-register.md#pre-01-a1): los contratos deben cruzar Go y TypeScript con generación y comprobación de tipos.
- [PRE-01.A2](../decision-register.md#pre-01-a2): la aplicación web utilizará Next.js, React y TypeScript.
- [PRE-01.A3](../decision-register.md#pre-01-a3): Gin será el transporte HTTP principal de los servicios Go.
- [PRE-01.A4](../decision-register.md#pre-01-a4): pnpm y módulos/workspace Go organizarán un monorepo con dependencias compartidas controladas.
- [PRE-01.B1](../decision-register.md#pre-01-b1): Ably Free reemplaza Socket.IO; no se necesita gateway Node.js ni WebSocket propio persistente.
- [PRE-01.B2](../decision-register.md#pre-01-b2): PostgreSQL será el almacén relacional; Neon Free es el proveedor previsto.
- [PRE-01.B3](../decision-register.md#pre-01-b3): sqlc, pgx/v5 y Goose sustituyen Drizzle para SQL y migraciones Go.
- [PRE-01.B4](../decision-register.md#pre-01-b4): aceptar un movimiento exige atomicidad entre su registro y el estado actualizado.
- [PRE-01.B5](../decision-register.md#pre-01-b5): el acceso inmediato de invitados convivirá con cuentas opcionales persistentes y clasificación.
- [PRE-01.C1](../decision-register.md#pre-01-c1): Vercel sustituye Railway mediante cuatro proyectos; la factibilidad sigue pendiente de prueba.
- [PRE-01.C2](../decision-register.md#pre-01-c2): el objetivo pasa de 20 a 0 USD/mes; ningún recurso de pago queda autorizado.
- [PRE-01.C3](../decision-register.md#pre-01-c3): el futuro entorno local combinará procesos nativos y PostgreSQL/auxiliares en Docker Compose.
- [PRE-01.C4](../decision-register.md#pre-01-c4): desarrollo, integración/previews y producción tendrán datos y credenciales separados.
- [PRE-01.D1](../decision-register.md#pre-01-d1): GitHub Actions ejecutará los futuros controles de tipado, estilo, tests, compilación e integración.
- [PRE-01.D2](../decision-register.md#pre-01-d2): OTel y Prometheus/Grafana cubrirán observabilidad, principalmente local y según cuotas productivas verificadas.
- [PRE-01.D3](../decision-register.md#pre-01-d3): el crecimiento partirá de una instancia por servicio; Redis requerirá justificación observada.
- [PRE-01.D4](../decision-register.md#pre-01-d4): el portafolio público necesita una elección expresa de licencia de reutilización.

## Agrupaciones de diseño y consecuencias

**Lenguajes, módulos y contratos.** El cliente consume datos tipados y no interpreta el reglamento. Los módulos compartidos deben permitir que Game y AI usen un único motor Go sin acoplar sus transportes o despliegues. La tipificación entre lenguajes reduce discrepancias, pero no reemplaza la validación en ejecución.

**Persistencia y distribución.** Identity y Game conservan propiedad separada de datos; AI comienza sin DB. La transacción de Game precede al aviso Ably. Un fallo de publicación no revierte una jugada confirmada; las especificaciones de servicios deberán definir reconciliación REST e idempotencia. No se supone atomicidad distribuida entre servicios.

**Capacidad y operación.** Cuatro proyectos Vercel significan web más tres servicios, no cuatro microservicios. El coste cero es objetivo de operación, no prueba de capacidad ilimitada. CI, observabilidad, aislamiento y licencia forman parte del diseño futuro, sin habilitar infraestructura ahora.

## Sustituciones y matices

El stack vigente reemplaza Node/Fastify por Go/Gin, Railway por Vercel, Socket.IO por Ably, Drizzle por sqlc/pgx/v5/Goose y el techo inicial de 20 USD por objetivo 0 USD. [PRE-05](PRE-05.md) precisa el reemplazo del monolito por tres servicios. No conservar alternativas históricas como requisitos acumulativos. El plan local describe dónde desarrollar: no aprueba juego offline. Redis queda opcional y sujeto a cambio aprobado si introduce capacidad, coordinación o coste nuevos.

## Dependencias y fuentes canónicas

Consumir [PRE-05](PRE-05.md) para límites de servicios y datos, [PRE-07](PRE-07.md) para identidad, [PRE-10](PRE-10.md) para operación y [PRE-11](PRE-11.md) para control documental. La definición detallada de infraestructura se mantiene en el [dominio 01 del índice](../index.md#spec-01-01), `PENDIENTE_LOTE_2`. Las reglas de elaboración son [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md) y [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md); la cobertura se integra en [trazabilidad](../traceability.md).

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: runtime Go beta de Vercel, empaquetado de módulos compartidos, integración Gin/Neon/Ably, cuotas gratuitas y disponibilidad de telemetría. Deben documentarse futuras pruebas y alternativas con impacto sin cambiar de proveedor ni añadir costes. Véase [conflictos](../conflicts.md). No se crean manifiestos, Dockerfiles, CI, credenciales, proveedores ni experimentos ejecutables en este lote.

## Criterios futuros de aceptación y verificación

- **AC-01:** un mapa de proyectos y dependencias identifica web, Identity, Game, AI y motor; la revisión demuestra ausencia de un cuarto servicio y de SQL cruzado.
- **AC-02:** un plan de factibilidad enumera compilación, adaptación HTTP, acceso Neon, publicación Ably y transacción; cada resultado futuro exige evidencia, no una declaración del agente.
- **AC-03:** una matriz por entorno identifica aislamiento, cuotas, límites y autorización necesaria antes de costes; la licencia figura como decisión explícita pendiente hasta aprobarse.
- **AC-04:** cada decisión anterior tiene destino en el índice y trazabilidad; un destino aún no redactado conserva su marca de lote pendiente.

Estos son criterios documentales y futuros. No se han ejecutado pruebas de producto ni mediciones de despliegue.

## Revisión humana

Revisar preservación de sustituciones, presupuesto, autonomía de servicios y riesgos de plataforma. El contenido permanece `EN_REVISIÓN`; aprobar el inventario no aprueba este texto. El cierre del lote corresponde al coordinador y requiere «APRUEBO LOTE 1» antes de avanzar.
