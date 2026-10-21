# PRE-10 — Despliegue, operaciones y costes

ID: PRE-10 · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner: A — operaciones.

## Objetivo y contexto

Consolidar las 24 decisiones operativas sin presentar planes gratuitos, recuperación o despliegue como capacidades observadas. El objetivo es operar web y tres servicios independientes con datos protegidos y coste objetivo cero, sujeto a factibilidad futura.

## Fuentes y decisiones

El [maestro archivado](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) y el [registro canónico](../decision-register.md) conservan las fuentes. Síntesis de naturaleza `APROBADO`:

- [PRE-10.A1](../decision-register.md#pre-10-a1): cuatro proyectos Vercel tendrán builds independientes y CI verificará dependencias/motor Go compartido.
- [PRE-10.A2](../decision-register.md#pre-10-a2): un adaptador pequeño http.Handler integrará Gin en runtime Go conservando ejecución local.
- [PRE-10.A3](../decision-register.md#pre-10-a3): rewrites reunirán web/Identity/Game bajo mismo origen; AI tendrá HTTPS protegido entre servicios.
- [PRE-10.A4](../decision-register.md#pre-10-a4): regiones cercanas Vercel/Neon se elegirán por medición, disponibilidad y privacidad.
- [PRE-10.B1](../decision-register.md#pre-10-b1): local, integración y producción tendrán OAuth, Ably y bases independientes.
- [PRE-10.B2](../decision-register.md#pre-10-b2): secretos por entorno en Vercel requieren rotación y exclusión de código/logs públicos.
- [PRE-10.B3](../decision-register.md#pre-10-b3): previews protegidas dentro de Hobby usarán datos sintéticos y OAuth aislado.
- [PRE-10.B4](../decision-register.md#pre-10-b4): límites serverless se compartirán/centralizarán por IP, identidad y recurso, sin depender solo de RAM.
- [PRE-10.C1](../decision-register.md#pre-10-c1): pgxpool acotado y conexiones agrupadas/directas se ajustarán a aplicación y migración.
- [PRE-10.C2](../decision-register.md#pre-10-c2): PITR limitado Neon Free se complementará con backup externo cifrado/restaurado; proveedor/coste requieren verificación.
- [PRE-10.C3](../decision-register.md#pre-10-c3): migraciones Goose aprobadas/controladas usarán expand/contract y propiedad de esquema por servicio.
- [PRE-10.C4](../decision-register.md#pre-10-c4): validez expira en solicitudes; limpieza diaria idempotente se ajustará a cron Hobby.
- [PRE-10.D1](../decision-register.md#pre-10-d1): OTel Go/Next.js y collector/Prometheus/Grafana local se complementan con producción gratuita verificada.
- [PRE-10.D2](../decision-register.md#pre-10-d2): traceId, requestId y gameId correlacionarán servicios sin datos sensibles innecesarios.
- [PRE-10.D3](../decision-register.md#pre-10-d3): métricas cubren p50/p95, errores, invocaciones, login, jugadas inválidas, reconexión, CPU IA y cuotas.
- [PRE-10.D4](../decision-register.md#pre-10-d4): alertas/sintéticos serán proporcionados a prestaciones reales, sin prometer atención 24/7 no contratada.
- [PRE-10.E1](../decision-register.md#pre-10-e1): coste objetivo 0 USD/mes y uso Hobby personal no comercial exigen permiso para cualquier pago.
- [PRE-10.E2](../decision-register.md#pre-10-e2): revisión interna al 50/75/90 % depende de datos de cuotas accesibles por proveedor.
- [PRE-10.E3](../decision-register.md#pre-10-e3): CPU/búsquedas IA se limitarán por identidad y solo se cacheará con seguridad probada/consumo medido.
- [PRE-10.E4](../decision-register.md#pre-10-e4): escalar requiere pruebas y aprobación de cambios como Redis, colas o réplicas.
- [PRE-10.F1](../decision-register.md#pre-10-f1): releases independientes mantendrán compatibilidad, smoke tests, expand/contract y flags para riesgos.
- [PRE-10.F2](../decision-register.md#pre-10-f2): runbooks P0–P3 cubrirán reglas, auth, datos, Ably, Elo, WebGL e IA con degradación controlada.
- [PRE-10.F3](../decision-register.md#pre-10-f3): RPO/RTO se fundamentarán en restauraciones ensayadas, sin garantía de pérdida cero.
- [PRE-10.F4](../decision-register.md#pre-10-f4): adaptadores/variables conservarán portabilidad Go a contenedores; Dockerfiles se documentan y no se crean ahora.

## Agrupaciones de diseño y consecuencias

**Plataforma y límites.** La independencia de builds requiere verificar empaquetado del motor compartido en cada raíz de proyecto. El mismo origen facilita cookies, pero rewrites no reemplazan autorización ni control de origen. La arquitectura de tres servicios permanece aun cuando las cuotas requieran revisar carga; no se cambia a otro proveedor ni se activa pago para resolver la incertidumbre silenciosamente.

**Estado durable y recuperación.** Las Functions no deben migrar al arrancar ni depender de memoria local para límites globales. Las solicitudes determinan validez temporal; el cron solo limpia después. Backups requieren destino, cifrado, retención, permisos, integridad y restauración ensayada. Un backup existente no demuestra capacidad de restauración; un rollback de código tampoco restaura datos borrados o migrados.

**Observación y coste.** Correlación sin secretos permite analizar fallos entre servicios. Los umbrales internos de cuota activan revisión; la degradación proporcionada y sus disparadores quedan por especificar y validar, sin garantizar disponibilidad del proveedor. OTel y métricas deben tener presupuesto de cardinalidad/retención; Prometheus/Grafana local no implica stack productivo gratuito. La IA necesita límites antes de multiplicar coste de cómputo.

## Sustituciones y matices

[PRE-01.C1](../decision-register.md#pre-01-c1) y [PRE-01.C2](../decision-register.md#pre-01-c2) sustituyen Railway y 20 USD por Vercel/0 USD objetivo. Ninguna cuota, precio, cron, región, backup externo gratuito o prestación Hobby se considera verificada aquí. Redis, colas y réplicas son alternativas sujetas a cambio aprobado, no infraestructura obligatoria del MVP. Portabilidad a contenedores es plan, sin archivos o despliegue creados.

## Dependencias y fuentes canónicas

El [dominio 10 del índice](../index.md#spec-10-01) es `PENDIENTE_LOTE_5`. Consume arquitectura de [PRE-01](PRE-01.md), persistencia de [PRE-05](PRE-05.md), seguridad/privacidad de [PRE-07](PRE-07.md), recursos IA de [PRE-08](PRE-08.md) y release/DoD de [PRE-09](PRE-09.md). Aplicar [11.01](https://github.com/SebasChaconDev/toguz-kumalak/blob/docs/specs-lote-01/docs/specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](https://github.com/SebasChaconDev/toguz-kumalak/blob/docs/specs-lote-01/docs/specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [trazabilidad](../traceability.md).

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: runtime Go beta, módulos compartidos, cuotas/protección de previews/cron, compatibilidad de mismo origen, pooling, observabilidad, regiones, backup externo y RPO/RTO. El [registro de conflictos](../conflicts.md) debe contener impactos y alternativas. No se instalan herramientas, conectan proveedores, provisionan datos, ejecutan restauraciones ni publican recursos; no hay coste observado de producción.

## Criterios futuros de aceptación y verificación

- **AC-01:** una matriz por proyecto/entorno identifica secretos, datos, región, límites, dependencias y gates; ninguna preview accede a datos productivos.
- **AC-02:** el plan de plataforma exige evidencia de build independiente, adaptación Gin, transacción y aviso/reconciliación antes de implementación completa.
- **AC-03:** runbooks distinguen expiración, limpieza, backup, restauración, rollback y migración; RPO/RTO solo se publican con ensayo medido.
- **AC-04:** presupuesto/cuotas y degradación quedan documentados con fecha/fuente a verificar, sin pagos automáticos ni nueva infraestructura asumida.

Estos criterios describen validación futura, no operaciones realizadas.

## Revisión humana

Revisar coste cero, uso previsto de Hobby, mecanismos pendientes de límites/backup y evidencia de recuperación requerida. Mantener `EN_REVISIÓN`. El coordinador pedirá «APRUEBO LOTE 1» y detendrá el avance sin desplegar.
