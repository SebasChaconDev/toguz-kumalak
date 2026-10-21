# PRE-06 — Salas, multijugador y ranking

ID: PRE-06 · Versión: 0.1 · Fecha: 2026-10-21 · Estado documental: EN_REVISIÓN · Owner: S — servicios.

## Objetivo y contexto

Consolidar las 20 decisiones de acceso privado, sincronización, recuperación y elegibilidad competitiva. Una presencia en Ably no demuestra pertenencia a una partida ni confirma un movimiento. El online ordinario mantiene ambos jugadores conectados, con recuperación ante interrupciones.

## Fuentes y decisiones

Consultar el [maestro íntegro](../sources/TOGUZ_PROMPT_MAESTRO_CODEX.md) y el [registro canónico](../decision-register.md). Síntesis de naturaleza `APROBADO`:

- [PRE-06.A1](../decision-register.md#pre-06-a1): enlace de alta entropía y código alternativo se protegerán con expiración y límites de tasa.
- [PRE-06.A2](../decision-register.md#pre-06-a2): sala pendiente caduca inicialmente a los 30 minutos de inactividad, configurable y separada de partidas iniciadas.
- [PRE-06.A3](../decision-register.md#pre-06-a3): las salas admiten dos jugadores y carecen de espectadores en MVP.
- [PRE-06.A4](../decision-register.md#pre-06-a4): revancha crea otra partida relacionada y reasigna colores.
- [PRE-06.B1](../decision-register.md#pre-06-b1): selección visual puede ser inmediata; ejecución y animación definitiva esperan confirmación del servidor.
- [PRE-06.B2](../decision-register.md#pre-06-b2): ID/versión detectan huecos y desorden; REST restablece sincronía.
- [PRE-06.B3](../decision-register.md#pre-06-b3): conectado, sincronizando, desconectado y recuperado serán visibles y bloquearán acciones sin confirmar.
- [PRE-06.B4](../decision-register.md#pre-06-b4): cada partida usa canal Ably privado y capacidad de suscripción restringida.
- [PRE-06.C1](../decision-register.md#pre-06-c1): recuperar exige identidad del participante original y estado oficial de Game.
- [PRE-06.C2](../decision-register.md#pre-06-c2): vencimiento de 24 h persistido se comprueba al actuar/consultar, independientemente de limpieza posterior.
- [PRE-06.C3](../decision-register.md#pre-06-c3): varias pestañas pueden observar sin duplicar efectos gracias a transacciones e idempotencia.
- [PRE-06.C4](../decision-register.md#pre-06-c4): recuperar reconstruye el tablero y señala el último movimiento confirmado.
- [PRE-06.D1](../decision-register.md#pre-06-d1): solo dos registrados y condición ranked fijada al inicio habilitan efecto Elo.
- [PRE-06.D2](../decision-register.md#pre-06-d2): el cálculo Elo se documentará separado de las estadísticas generales.
- [PRE-06.D3](../decision-register.md#pre-06-d3): límites auditables deberán frenar repetición abusiva de enfrentamientos y manipulación de Elo.
- [PRE-06.D4](../decision-register.md#pre-06-d4): abandonar voluntariamente pierde; desconexión archivada como inconclusa no cambia Elo.
- [PRE-06.E1](../decision-register.md#pre-06-e1): caída Ably conserva movimientos PostgreSQL, informa interrupción y activa reconciliación REST.
- [PRE-06.E2](../decision-register.md#pre-06-e2): Game valida firma sin consultar Identity por jugada; altas/renovaciones sí dependen de Identity disponible.
- [PRE-06.E3](../decision-register.md#pre-06-e3): E2E futuro cubrirá dos clientes, cortes, duplicados, simultaneidad, pestañas y fallos parciales.
- [PRE-06.E4](../decision-register.md#pre-06-e4): p50/p95 de confirmación y resincronización precederán a SLO definitivos.

## Agrupaciones de diseño y consecuencias

**Acceso y ciclo de sala.** Conocer un código no sustituye la autorización del participante. La futura definición debe separar sala pendiente, partida iniciada y revancha; 30 minutos de inactividad de la primera no se aplica automáticamente a la segunda. Cambiar colores en revancha exige nueva partida, no reescribir la anterior.

**Sincronización autoritativa.** El cliente mantiene estado confirmado y selección/previsualización distinguidos. Los eventos son indicios versionados: duplicados, desorden o huecos no autorizan reconstrucciones arbitrarias. REST obtiene la verdad persistida. La vista no necesita recibir cada animación histórica para alcanzar el estado actual.

**Recuperación e integridad competitiva.** La identidad invitada debe seguir siendo recuperable durante el plazo compatible de 24 h. La validez temporal se determina en solicitudes, sin depender de un cron puntual. Varias pestañas comparten autoridad transaccional y no obtienen turnos extra. Una caída de red no se transforma automáticamente en derrota o Elo; abandono voluntario y archivo inconcluso conservan causas distintas.

## Sustituciones y matices

Ably sustituye Socket.IO como notificación privada, sin socket propio persistente ni Redis obligatorio. Recuperación de 24 h no define una modalidad de correspondencia ni introduce reloj competitivo. Las partidas con invitados, IA o práctica quedan fuera de Elo; registrar una cuenta después no vuelve clasificatoria una partida. Los límites antiabuso concretos deben justificarse y no inventarse como decisiones ya aprobadas.

## Dependencias y fuentes canónicas

El [dominio 06 del índice](../index.md#spec-06-01) es `PENDIENTE_LOTE_3`. Consume transacciones/eventos de [PRE-05](PRE-05.md), sesiones y Elo de [PRE-07](PRE-07.md), resultado/causa de [PRE-03](PRE-03.md) y UX de [PRE-04](PRE-04.md). Seguir [11.01](../specs/11-orchestration/11.01-document-standard-and-change-control.md), [11.02](../specs/11-orchestration/11.02-ownership-delegation-and-context.md) y [trazabilidad](../traceability.md).

## Límites, riesgos y cuestiones pendientes

`POR_VERIFICAR`: inicio/reinicio preciso de la ventana de recuperación, expiración frente a renovación de sesiones/canales, cuotas Ably, conciliación sin outbox y controles contra abuso. Son temas de [conflictos](../conflicts.md), no resultados medidos. No se emiten invitaciones/tokens ni se crean salas, proveedores, pruebas E2E o SLO observados aquí.

## Criterios futuros de aceptación y verificación

- **AC-01:** máquinas de estado de sala/partida distinguen caducidad pendiente, desconexión recuperable, archivo, abandono y revancha, con permisos por transición.
- **AC-02:** matriz de dos clientes/pestañas con duplicados y huecos termina en la misma versión oficial y nunca aplica dos veces una jugada.
- **AC-03:** un tercero con código, sesión ajena o canal no autorizado no recupera ni modifica una partida; la expiración se verifica aun si no pasó limpieza.
- **AC-04:** la matriz ranked × identidades × causa final demuestra cuándo Elo cambia y cuándo queda intacto; las metas de latencia esperan medición.

La comprobación de este lote es textual y no demuestra resiliencia de una integración.

## Revisión humana

Revisar política temporal, efecto competitivo de causas finales y experiencia de interrupción. Mantener `EN_REVISIÓN`; todo ajuste de alcance requiere aprobación. El cierre documental se someterá a «APRUEBO LOTE 1» y parada obligatoria.
