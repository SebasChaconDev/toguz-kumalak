# ADR-0006 — Cuatro proyectos Vercel Hobby y objetivo de coste cero

> Aprobación posterior: el propietario declaró «APRUEBO LOTE 1» el 2026-10-21 para el contenido revisado en HEAD e38cc84950858e072c1a3f06ffbf854cf7e544f5. PR #1 integrado por squash en main: 329b7c1059e8a74f191c256870734e8ae8259966. Este registro prevalece sobre menciones históricas de revisión pendiente; no acredita producto probado. Los permisos actuales se consultan en AGENTS y PLANS.

ID: ADR-0006 · Versión: 0.2 · Fecha: 2026-10-21 · Estado documental: APROBADO · Owner funcional: A (arquitectura e infraestructura).

Dominio: 01/10, plataforma y coste · Lote: 1. Formato y cambios: [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md); propiedad de escritura: [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md).

Decisión base: APROBADO EN FUENTE. Validación técnica: POR_VERIFICAR. Vercel/Neon/Ably son proveedores previstos; no hay despliegue, precio, cuota ni compatibilidad demostrados por esta ADR.

## Contexto y problema

La arquitectura aprobada utiliza frontend y tres servicios Go independientes, datos separados y notificaciones privadas. Debe mantenerse en cuatro proyectos Vercel con objetivo operativo 0 USD/mes. El soporte Go beta, el empaquetado de módulos compartidos, límites serverless y cuotas gratuitas pueden condicionar la factibilidad; no permiten sustituir la arquitectura ni activar pagos unilateralmente.

## Fuentes y trazabilidad

- [PRE-01.C1](../decision-register.md#pre-01-c1) sustituye Railway por Vercel y exige cuatro proyectos; [PRE-01.C2](../decision-register.md#pre-01-c2) sustituye 20 USD por objetivo 0 USD/mes y permiso explícito para pagos.
- [PRE-05.A4](../decision-register.md#pre-05-a4) exige despliegues y módulos independientes; [PRE-05.F4](../decision-register.md#pre-05-f4) exige prueba temprana futura de Go/Gin, Neon, Ably y transacción.
- [PRE-10.A1](../decision-register.md#pre-10-a1) describe monorepo y builds; [PRE-10.A2](../decision-register.md#pre-10-a2) describe adaptador de http.Handler para Gin; [PRE-10.A4](../decision-register.md#pre-10-a4) condiciona regiones a disponibilidad, privacidad y medidas.
- [PRE-10.E1](../decision-register.md#pre-10-e1) conserva objetivo cero y ámbito personal no comercial de Hobby; [PRE-10.E2](../decision-register.md#pre-10-e2) fija umbrales internos 50/75/90 %; [PRE-10.E4](../decision-register.md#pre-10-e4) sujeta escalado a pruebas/aprobación.
- [PRE-10.F4](../decision-register.md#pre-10-f4) conserva portabilidad a contenedores; [PRE-11.E2](../decision-register.md#pre-11-e2) limita el spike actual a planificación documental.

Ver [maestro, base arquitectónica y controles de coste/runtime](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) y [trazabilidad](../traceability.md). Las condiciones de proveedores son información pendiente de contraste oficial fechado, no hechos verificados en este lote.

## Decisión delimitada

Se mantienen cuatro proyectos Vercel Hobby desde el monorepo: web Next.js, Identity, Game y AI. Cada servicio Go/Gin tiene módulo y despliegue propio; el motor compartido debe poder compilarse y empaquetarse sin convertirlos en un solo servicio. Se prevé adaptador pequeño de `http.Handler` conservando arranque local.

El runtime Go de Vercel se trata como beta y su compatibilidad real queda POR_VERIFICAR. No se declara viable el despliegue hasta observar una prueba futura autorizada. Tampoco se asume que los límites actuales de Functions, builds o duración permitan cualquier búsqueda IA.

El objetivo es 0 USD/mes; no es coste observado ni garantía de alojamiento perpetuamente gratuito. Ningún recurso de pago, upgrade o proveedor facturable se activa sin autorización expresa. Si las cuotas gratuitas no alcanzan, se presenta evidencia y alternativas al propietario; no se elimina el MVP ampliado ni se cambia de proveedor por decisión del agente.

## Guardrails y límites

1. Revisar consumos al 50/75/90 % de cuotas mediante los datos disponibles del proveedor. Son umbrales internos de revisión aprobados, no cuotas comerciales afirmadas ni acciones automáticas de compra.
2. Separar local, integración/previews y producción, con credenciales y datos independientes. Previews protegidas, datos sintéticos y OAuth de prueba siguen [PRE-10.B1](../decision-register.md#pre-10-b1) y [PRE-10.B3](../decision-register.md#pre-10-b3); disponibilidad efectiva en Hobby sigue POR_VERIFICAR.
3. Limitar conexiones pgxpool y usar acceso agrupado o directo según aplicación/migraciones; no migrar al arrancar Functions. No se fijan tamaños de pool antes de verificar límites reales.
4. Proteger CPU/búsquedas de IA y límites por identidad/IP/recurso compatibles con serverless, sin confiar solo en RAM de una Function, según [PRE-10.B4](../decision-register.md#pre-10-b4) y [PRE-10.E3](../decision-register.md#pre-10-e3). El mecanismo y valores quedan pendientes, sin imponer Redis.
5. Seleccionar regiones cercanas a Neon mediante disponibilidad, privacidad y mediciones; no se elige una región ni se promete latencia aquí.
6. Backups externos cifrados, restauraciones y RPO/RTO necesitan validación. No se supone almacenamiento externo gratis ni restauración ensayada, según [PRE-10.C2](../decision-register.md#pre-10-c2) y [PRE-10.F3](../decision-register.md#pre-10-f3).

## Alternativas y motivación

**Históricas descartadas en fuente:** Railway como proveedor principal y presupuesto máximo de 20 USD se sustituyeron expresamente. No son permisos residuales para pagar o desplegar fuera de Vercel.

**PROPUESTA, no adoptadas:** plan de pago, otro hosting principal, colas/Redis o un despliegue directo en contenedores. La portabilidad a contenedores es un plan aprobado; ejecutarla como cambio de hosting exige una nueva autorización. No hay un cuarto microservicio adicional aprobado para resolver límites.

La motivación trazable es la elección expresa de proveedor/presupuesto y la independencia de PRE-05.A4. El coste de coordinación y el riesgo de cuotas compartidas deben hacerse visibles sin prometer que todo cabe. Los límites serverless no autorizan recortar las funcionalidades ampliadas del lanzamiento comprometido.

## Dependencias y validación futura

Depende de [01.01, monorepo](../index.md#spec-01-01), [01.02, factibilidad](../index.md#spec-01-02) y [01.03, coste/capacidad](../index.md#spec-01-03), PENDIENTE_LOTE_2; [10.01, despliegue](../index.md#spec-10-01), [10.02, cuotas](../index.md#spec-10-02), [10.03, restauración](../index.md#spec-10-03) y [09.03, controles de release](../index.md#spec-09-03), PENDIENTE_LOTE_5. [ADR-0001](ADR-0001-three-go-services-and-data-ownership.md) y [ADR-0002](ADR-0002-single-versioned-go-engine.md) fijan la forma que debe validarse, no una configuración ejecutada.

La futura prueba temprana requiere autorización independiente para implementación/proveedores. Deberá observar build y empaquetado del motor común, adaptador Gin, cuatro proyectos independientes, conexión Neon, transacción de Game y publicación Ably posterior al commit. Se recogerán límites y documentación oficial fechada, consumo, secretos por entorno, fallo de publicación y rollback compatible. Un éxito aislado no demuestra capacidad de producción ni coste mensual cero.

## Aceptación futura

- **AC-01, positiva:** bajo autorización posterior, cuatro proyectos compilan y despliegan independientemente con el motor Go compartido y sin acceso cruzado de datos; la evidencia identifica runtime, revisiones y entornos.
- **AC-02, positiva:** un movimiento de prueba sintético se confirma en PostgreSQL antes de notificar, con aislamiento de datos y secretos.
- **AC-03, negativa:** falta de cuota o incompatibilidad Go bloquea la declaración de factibilidad y genera conflicto con evidencia; no activa pago ni sustituye stack automáticamente.
- **AC-04, negativa:** una preview no recibe credenciales o datos productivos ni se considera segura únicamente por ser una URL difícil de adivinar.
- **AC-05, negativa:** una estimación, un dashboard parcial o una prueba breve no se presenta como coste cero mensual, RPO/RTO o latencia garantizados.

Estas aceptaciones son futuras y permanecen POR_VERIFICAR.

Trazabilidad de aceptación: AC-01/AC-02 → [PRE-05.F4](../decision-register.md#pre-05-f4), dueño [01.02](../index.md#spec-01-02), PENDIENTE_LOTE_2; AC-03 → [PRE-01.C2](../decision-register.md#pre-01-c2), dueño [01.03](../index.md#spec-01-03), PENDIENTE_LOTE_2; AC-04 → [PRE-10.B3](../decision-register.md#pre-10-b3), dueño [10.01](../index.md#spec-10-01), PENDIENTE_LOTE_5; AC-05 → [PRE-10.E1](../decision-register.md#pre-10-e1) y [PRE-10.F3](../decision-register.md#pre-10-f3), dueños [10.02](../index.md#spec-10-02)/[10.03](../index.md#spec-10-03), PENDIENTE_LOTE_5. Riesgos asociados: [RISK-02](../conflicts.md#risk-02), [RISK-03](../conflicts.md#risk-03) y [RISK-11](../conflicts.md#risk-11).

## Cambio y reversión documental

Un cambio de proveedor/plan, coste, número de servicios o despliegues necesita propuesta con límites observados, alternativas e impacto de privacidad/operación. Registrar en [conflictos](../conflicts.md), obtener aprobación y actualizar ADR/trazabilidad. Una redacción rechazada se revierte documentalmente a la última aprobada preservando la evidencia. No se ejecutan despliegue, compra, migración ni rollback de infraestructura en esta fase.
