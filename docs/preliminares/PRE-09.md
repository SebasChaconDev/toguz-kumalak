# PRE-09 — Testing, CI/CD y Definition of Done

ID: PRE-09 · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner: Q — calidad.

## Objetivo y contexto

Consolidar las 24 decisiones que definirán evidencia de calidad y cierre futuro. Cobertura, herramientas y confirmación de un agente no demuestran por sí solas corrección. Este lote verifica documentación; no ejecuta pruebas de producto ni configura CI/CD.

## Fuentes y decisiones

El [maestro archivado](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) conserva el Anexo A y el [registro canónico](../decision-register.md) sus decisiones exactas. Síntesis de naturaleza `APROBADO`:

- [PRE-09.A1](../decision-register.md#pre-09-a1): la pirámide priorizará riesgo con unitarias, SQL/contratos y E2E críticos.
- [PRE-09.A2](../decision-register.md#pre-09-a2): reglas, transacciones, autorización y contratos seguirán TDD; UI/3D requiere aceptación y pruebas visuales.
- [PRE-09.A3](../decision-register.md#pre-09-a3): objetivos de sentencias serán 90 % motor y 80 % lógica de servicios, además de rutas críticas.
- [PRE-09.A4](../decision-register.md#pre-09-a4): fixtures, semillas, casos normativos y factorías estarán aislados y serán reproducibles.
- [PRE-09.B1](../decision-register.md#pre-09-b1): corpus, propiedades, fuzzing y simulación comprobarán conservación de 162 piedras.
- [PRE-09.B2](../decision-register.md#pre-09-b2): IA necesita legalidad, límites, cancelación, semillas, terminales y diferencias medibles.
- [PRE-09.B3](../decision-register.md#pre-09-b3): integración CI usará PostgreSQL efímero real con Goose, sqlc, índices, restricciones y transacciones.
- [PRE-09.B4](../decision-register.md#pre-09-b4): duplicados, carreras, versiones antiguas y go test -race se cubrirán cuando aplique.
- [PRE-09.C1](../decision-register.md#pre-09-c1): compatibilidad HTTP/modelos y esquemas de eventos/consumidores se validará contra contratos.
- [PRE-09.C2](../decision-register.md#pre-09-c2): seguridad cubrirá JWKS, expiración, audiencia, rotación, privilegios internos y acceso cruzado.
- [PRE-09.C3](../decision-register.md#pre-09-c3): fallos de los tres servicios, Ably y PostgreSQL deben demostrar degradación segura/reintentos.
- [PRE-09.C4](../decision-register.md#pre-09-c4): Compose aislará integración; predominan proveedores simulados y se separa integración real.
- [PRE-09.D1](../decision-register.md#pre-09-d1): Vitest/React Testing Library cubrirán componentes y adaptadores visuales.
- [PRE-09.D2](../decision-register.md#pre-09-d2): Playwright abarcará Chromium, Firefox, WebKit, móvil y dos contextos de participantes.
- [PRE-09.D3](../decision-register.md#pre-09-d3): capturas 3D usan entorno fijo y revisión manual por variación GPU.
- [PRE-09.D4](../decision-register.md#pre-09-d4): WCAG 2.2 AA aplicable es objetivo junto a teclado/lector, movimiento reducido y HTML jugable.
- [PRE-09.E1](../decision-register.md#pre-09-e1): en implementación futura habrá ramas/PR, main protegida y checks obligatorios sin fusionar pruebas rotas.
- [PRE-09.E2](../decision-register.md#pre-09-e2): CI se escalonará entre PR, merge y pre-release; cambios compartidos activarán consumidores.
- [PRE-09.E3](../decision-register.md#pre-09-e3): local/CI, preview, integración sintética y producción permanecerán aislados.
- [PRE-09.E4](../decision-register.md#pre-09-e4): releases tendrán trazabilidad, smoke tests, rollback, expand/contract y recuperación de datos.
- [PRE-09.F1](../decision-register.md#pre-09-f1): P0/P1 bloquean release, con prioridad en corrupción, reglas e identidad.
- [PRE-09.F2](../decision-register.md#pre-09-f2): Web Vitals, p95, IA, FPS, descarga, reconexión y cuotas requieren benchmarks para umbrales.
- [PRE-09.F3](../decision-register.md#pre-09-f3): controles de secretos/dependencias, estático y govulncheck cubrirán el entorno multimódulo.
- [PRE-09.F4](../decision-register.md#pre-09-f4): DoD exige requisito, prueba observada, ausencia de bloqueantes, documentación y evidencia verificable.

## Agrupaciones de diseño y consecuencias

**Riesgo e invariantes.** Los tests futuros deben demostrar corrección normativa y consistencia, no copiar la implementación. Un porcentaje alto no exonera casos de tuzdyk, finales, repetición, conservación o concurrencia. El corpus depende de resolver discrepancias reglamentarias; no puede declarar válida una interpretación pendiente por estar automatizada.

**Fronteras y fallos.** PostgreSQL real contrasta transacciones, restricciones y migraciones que un doble no reproduce. Proveedores simulados permiten fallos controlados; la integración real queda separada con datos/credenciales de prueba. Las pruebas de seguridad deben negar acceso ajeno y privilegios incorrectos además de aceptar el caso feliz. Notificación caída con movimiento confirmado es un caso necesario de consistencia.

**Experiencia y release.** Tests visuales no sustituyen revisión multiplataforma, accesibilidad manual o mediciones GPU. La CI futura tiene capas para equilibrar tiempo y cobertura, pero un cambio del motor compartido requiere revisar consumidores Game y AI. Rollback de código no garantiza revertir datos; expand/contract y recuperación deben tratarse por separado.

## Sustituciones y matices

Ramas, PR y Actions son decisiones de trabajo futuro; no autorizan mutaciones Git o CI en esta misión. WCAG es objetivo de conformidad que exige evidencia, no etiqueta automática por tener HTML. Las cifras de cobertura son objetivos aprobados y no mediciones existentes. Los umbrales de rendimiento se fijarán a partir de benchmarks, sin declarar FPS, latencia o fuerza IA de antemano.

## Dependencias y fuentes canónicas

El [dominio 09 del índice](../index.md#spec-09-01) es `PENDIENTE_LOTE_5`. Consume riesgos/contratos de [PRE-03](PRE-03.md), [PRE-05](PRE-05.md), [PRE-06](PRE-06.md), [PRE-07](PRE-07.md) y [PRE-08](PRE-08.md), con visuales de [PRE-04](PRE-04.md) y releases de [PRE-10](PRE-10.md). Aplican [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [trazabilidad](../traceability.md).

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: corpus oficial, dispositivos/escenas, compatibilidad de herramientas, costes de CI, umbrales y comportamiento real ante fallos. El [registro de conflictos](../conflicts.md) conserva bloqueos. No se crean tests, workflows, Compose, escáneres ni ejecutables. La comprobación de Markdown no acredita funcionalidad ni seguridad de producto.

## Criterios futuros de aceptación y verificación

- **AC-01:** matriz riesgo → requisito → prueba → evidencia cubre casos positivos/negativos y propietarios, incluida cada ruta crítica aunque supere cobertura objetivo.
- **AC-02:** el plan separa unitarias, PostgreSQL real, contratos, E2E, visual/manual y proveedores reales/simulados con aislamiento explícito.
- **AC-03:** gates de release documentan P0/P1, evidencias requeridas, degradación, rollback y recuperación de datos sin asumir reversión automática.
- **AC-04:** ningún criterio aparece satisfecho sin resultado observado reproducible, entorno/versiones y vínculo al requisito en trazabilidad.

Esta consolidación documenta la estrategia; no informa tests pasando.

## Revisión humana

Revisar suficiencia de escenarios, evidencia y bloqueantes antes de aprobar estrategias detalladas. Mantener `EN_REVISIÓN`. El coordinador pedirá «APRUEBO LOTE 1» y detendrá el avance documental; implementar o probar requiere autorización independiente.
