# Corpus TK-N1-1.0.0 — PH-01-T01

Oráculo de datos para el futuro motor Go. El autor M redacta los datos; la revisión independiente Q y sus hallazgos se conservan en la issue o PR. La instrucción de ejecución 001 autoriza estos dos archivos. [Expediente PH-00, issue #7](https://github.com/SebasChaconDev/toguz-kumalak/issues/7); [entrega T01, issue #8](https://github.com/SebasChaconDev/toguz-kumalak/issues/8).

## Alcance de la evidencia

El corpus contiene 14 CaseSpec, 66 CorpusEntry, 19 testigos aritméticos de transición y 14 entradas estructuralmente inválidas marcadas. Uno de los testigos reutiliza ACC09 para comprobar ACC14. Se transcriben los pares canónicos de 03.01 y se añaden las variantes concretas necesarias para los negativos de 09.01.

La validación del autor observa datos y aritmética. El corte de autoría fechado `2026-10-26T20:24:09Z`, incluido en `authoringSnapshot`, conserva `independentReview:NOT_RUN` como snapshot de autoría; no expresa el estado posterior de la revisión independiente. La revisión efectiva, los hallazgos y las correcciones posteriores se registran en la [issue #8](https://github.com/SebasChaconDev/toguz-kumalak/issues/8) y la PR. Motor, replay ejecutado, Game, PostgreSQL y concurrencia permanecen `NOT_RUN`. Los esperados no constituyen pruebas de producto superadas.

No existe aquí una trayectoria legal revisada de triple repetición. Los ledgers de una, dos, tres y cuatro apariciones son dobles de contrato declarados. La integración posterior deberá acreditar la historia legal confiable, la autorización, los receipts, los locks y la persistencia. No se aceptarán contadores del cliente como historia oficial.

## Formato y decisiones de representación

`normativeCorpusVersion:1` y `formatId:TK-N1-Corpus/1` fijan el formato de fixtures, compatible conceptualmente con `CaseSpec/1` y `CorpusEntry/1` de 09.01 §8. No aprueban DTO de producto ni las propuestas condicionales de L4-P01.

| Campo | Significado |
|---|---|
| `sources` | Ruta desde la raíz, versión documental, commit y SHA256 de los bytes exactos de `git show <commit>:<path>`. No hay conversión de saltos de línea del checkout. |
| `caseSpecs[].caseId` | Los 14 IDs canónicos únicos `F-Q1-N01` a `F-Q1-N14`; cada uno corresponde a una ACC. |
| `entries[].caseId` | FK al CaseSpec primario. Puede repetirse entre variantes. `normativeCaseIds` y `acceptanceIds` muestran la cobertura secundaria. |
| `entryId`, `fixtureId` | Identificadores de variante únicos globales. Los CaseSpec enlazan `positiveEntryIds`, `negativeEntryIds` y `fixtureIds`. |
| `negativeCases[].negativeCaseId` | ID único global para cada predicado negativo local; su `fixtureId` coincide con la entrada propietaria y `scope:LOCAL_FIXTURE_ASSERTION` delimita la aserción. Los CaseSpec enlazan estos IDs. |
| `initialState`, `action`, `expectedState` | Datos y salida esperada separados. En un rechazo no existe transición aceptada: `expectedState:null` y entrada conservada. |
| `expectedEvents`, `expectedOutcome` | Lista semántica del movimiento y terminal esperado. `null` significa no aplicable o no especificado para esa capa, nunca ejecución satisfactoria. |
| `inputValidity` | `INVALID_EXPECTED` identifica los 14 inputs malformados: no se les exige forma válida ni suma 162. Una acción ilegal sobre un estado válido conserva `VALID_COMPLETE_STRUCTURE`. |
| `verificationStatus` | Solo `PASS / FAIL / BLOCKED / NOT_RUN`. El PASS de datos del autor no cambia las ejecuciones de motor, Game, PostgreSQL, concurrencia o revisión independiente. |

El commit fuente es `a79655a7e5c459d763a4c600730e8e8acb921142`. Los hashes son de blobs Git, identificados por `GIT_BLOB_BYTES_NO_CHECKOUT_NEWLINE_CONVERSION`; no deben compararse con bytes CRLF sin verificar antes el convenio. Versiones: 03.01–03.04 v0.4; 09.01, PH01, ADR-0002 y 05.03 v0.2; AD-NORM-001 en el registro. Los localizadores distinguen transcripción de fixture y derivación manual desde reglas.

El estado usa `stateSchemaVersion:1`, reglas exactas `TK-N1-1.0.0` y revisión como cadena decimal. El tablero es `[white, black]`, con nueve enteros por fila; `pitIndex:0..8` corresponde a los números propios 1 a 9. Kazanes y slots de tuzdyk tienen ese mismo orden. Cada tuzdyk expresa propietario y localización en campo adversario, índice 0 a 7, con cero piedras. Si existen ambos, sus índices propios difieren.

El ciclo fijo es white0…white8 → black0…black8 → white0. Los kazanes no son destinos del ciclo. `terminal` es nulo o `{outcome, reason, winner, source}`; `expectedOutcome` copia ese terminal. Target y equal conservan el turno del actor porque preceden a la transferencia; atsyrau conserva al siguiente jugador agotado. El ganador se lee del terminal, nunca del turno. Un empate lleva ganador nulo.

Estado, comando, eventos, movimiento, replay y contexto tienen versiones distintas aunque aquí valgan 1. Revisión, `historyOrdinal`, `moveOrdinal` y aparición también son ejes distintos. HISTORY-MOVE-CLOSE exige dos ordinales, 71 y 72, para movimiento y cierre con una misma revisión 37, sin nueva aparición activa. Esos números son tokens sintéticos de contrato, no historia observada.

## Procedencia e historia

| Clasificación / alcance | Evidencia |
|---|---|
| `SOURCE_DIRECT / INITIAL` | NORM-INIT transcribe 18 × 9, blancas, kazanes vacíos y ausencia de tuzdyks. |
| `DERIVED_MANUAL / DERIVED_PATH_REVIEWED` | DER-SOW-01 y DER-SOW-CAPTURE parten, cada uno, de NORM-INIT y una sola jugada explícita revisada manualmente. REPLAY-SOW reutiliza el primer prefijo, sin afirmar replay ejecutado. |
| `SYNTHETIC_CONTRACT / UNPROVEN` | Todas las demás posiciones, transiciones y contratos. Verificar una transición no demuestra cómo alcanzar su entrada desde el inicio. |

`historyContext` declara `trustBoundary:TEST_DOUBLE_ONLY`, `officialLedger:false` y `completeForOfficialGame:false`. `trajectoryEvidence` distingue únicamente los prefijos anteriores. `candidateActiveOccurrenceDelta` describe el candidato puro. `officialOccurrenceDelta:0` indica que T01 no confirma ninguna operación en Game; una futura confirmación activa sí deberá registrar la aparición correspondiente a R-07.

Las etiquetas `synthetic-occurrence-N` no afirman movimientos intermedios ni ordinales oficiales. La reclamación supone como precondición la validación autoritativa del ledger por la integración futura. ACC14 inyecta un doble que devolvería tres apariciones; la rama terminal deberá evitar consultarlo. No se fabrica una partida activa históricamente repetida que ya debió terminar.

Los cierres administrativos no añaden apariciones de repetición, pero sí conservan su lifecycle e historial una vez. Los reintentos no duplican el cierre. Las carreras requieren dos conexiones PostgreSQL independientes y barreras; no se consideran probadas por estos datos.

En HISTORY-RESIGNATION y HISTORY-MUTUAL_AGREEMENT, `expectedApplicationEffect.retryReusesReceipt:true` exige que el reintento de la misma intención recupere el receipt previo, actualmente legible, con el mismo actor, recurso, commandId y fingerprint y con autorización actual sobre actor y recurso. La consulta del receipt precede a las comprobaciones de turno o fase para una intención nueva; el cierre previo no debe causar un segundo rechazo por esas precondiciones ya alteradas. Un fingerprint diferente bajo la misma identidad de comando produce COMMAND_ID_REUSED, sin nuevo efecto ni reemplazo del receipt. Otro actor o recurso no se trata como la misma intención ni recibe el receipt previo. La falta de legibilidad o autorización actual tampoco permite recuperarlo.

`retryReusesReceipt` es un booleano contractual cuando la fixture ensaya un reintento: `true` exige recuperación; `false` significaría que no se recupera. No representa «no aplicable». El campo se omite en HISTORY-PREVIEW, HISTORY-SIMULATION, HISTORY-INVALID_MOVE y HISTORY-HEARTBEAT porque esas entradas no ensayan reintentos. Todos estos `expectedApplicationEffect` son aceptación futura de contrato y permanecen `NOT_RUN`; no acreditan ejecución de Game o PostgreSQL.

## Eventos y testigos aritméticos

Los movimientos enumeran los eventos del convenio 03.03 con `{eventSchemaVersion:1, sequence, revision, type, data}`. La secuencia comienza en cero, sin huecos, y usa la revisión candidata de salida. `MoveStarted` conserva la cantidad original. `StoneSown` describe solo un depósito real: la piedra retenida no genera ese evento. En este formato, el ordinal de depósito empieza en 1 y es distinto de la secuencia del evento.

Un depósito en tuzdyk usa exclusivamente `TuzdykReceived`, conserva cero lógico y acredita al propietario. Captura sigue al depósito final; creación sigue a la retirada de tres; cambio de turno sigue a los efectos cuando R-05 lo permite; `GameEnded` es el último evento. Los contratos administrativos que no tienen lista exacta fijada en 03.03 usan `expectedEvents:null` y describen su aceptación en `expectedApplicationEffect`.

`manualWitness` aporta retención, depósitos explícitos con cantidades antes y después, efectos de captura y creación, y sumas completas. El validador inferior comprueba esa contabilidad; no elige destinos ni decide captura, tuzdyk o resultado. El revisor deberá cotejar esas decisiones con las reglas. Los estados intermedios con piedras en mano no son snapshots completos persistibles.

Ejemplos independientes:

- DER-SOW-CAPTURE conserva una en white9 y deposita ocho en campo negro. Black8 pasa de nueve a diez; se capturan las diez. Tablero 152 + kazanes 10 = 162.
- MULTI-TUZ-OWNER deposita 19 piedras desde white9=20. El tuzdyk blanco recibe dos, el negro una: K60/62 → K62/63. El recorrido vuelve al origen y lo deja con dos. Tablero 37 + kazanes 125 = 162. La entrada sigue UNPROVEN.
- TARGET-COMPLETE-SOW llega a Kwhite=82 en el primer depósito y completa los nueve restantes. Conserva 50 piedras en tablero y K82/30; solo después evalúa target, manteniendo el turno blanco.
- EQUAL81 deposita la última piedra en un tuzdyk negro: K81/80 → K81/81. No cambia el turno antes del empate. ACC07 es otro empate: atsyrau con K80/81 y una piedra remanente.

Los nombres de error están etiquetados como propuestas de las specs. REJECT-TUZ permite rechazo seguro por tuzdyk no movible o pozo vacío porque no se ha fijado la precedencia de esas guardas. La prohibición de crear un tuzdyk no rechaza la jugada: deja tres piedras normalmente.

## Trazabilidad

Cada fixture de esta tabla existe en `entries`. Los CaseSpec enlazan los IDs completos de variantes y predicados negativos. Hay 102 definiciones de predicados locales distribuidas entre 66 entradas; no se presentan como 102 escenarios independientes. `relatedFixtureReferences` enlaza variantes ejercitadas por otras entradas y no les atribuye ejecución dentro de la fixture que las referencia.

| ACC | Caso 09.01 | Fixtures |
|---|---|---|
| ACC-N1-01 | F-Q1-N01 | `NORM-INIT`, `DER-SOW-01`, `DER-SINGLE`, `INVALID-TOTAL161`, `INVALID-TOTAL163`, `INVALID-NEGATIVE`, `INVALID-FRACTION`, `INVALID-OVER162`, `INVALID-ROWS`, `INVALID-COLUMNS`, `INVALID-PLAYER`, `INVALID-STATEVERSION`, `INVALID-RULESVERSION`, `INVALID-TUZOWNFIELD`, `INVALID-TUZNINTH`, `INVALID-TUZNONZERO`, `INVALID-TUZHOMOLOGOUS`, `REJECT-EMPTY`, `REJECT-RIVAL`, `REJECT-INDEX-NEG`, `REJECT-INDEX-NINE`, `REJECT-WRONG-TURN`, `REJECT-TERMINAL`, `REJECT-TUZ`, `REPLAY-SOW` |
| ACC-N1-02 | F-Q1-N02 | `DER-SOW-01`, `DER-SOW-CAPTURE`, `TARGET-COMPLETE-SOW`, `RIVAL-ONE`, `RIVAL-FIVE`, `MULTI-TUZ-OWNER`, `REPLAY-SOW` |
| ACC-N1-03 | F-Q1-N03 | `DER-TUZ-CREATE`, `RIVAL-ONE`, `RIVAL-FIVE` |
| ACC-N1-04 | F-Q1-N04 | `DER-TUZ-NINTH`, `DER-TUZ-HOMOLOGOUS`, `ACC04-SECOND`, `INVALID-TUZOWNFIELD`, `INVALID-TUZNINTH`, `INVALID-TUZNONZERO`, `INVALID-TUZHOMOLOGOUS` |
| ACC-N1-05 | F-Q1-N05 | `DER-TUZ-DEPOSIT`, `MULTI-TUZ-OWNER`, `REJECT-TUZ` |
| ACC-N1-06 | F-Q1-N06 | `ACC06`, `TARGET-COMPLETE-SOW` |
| ACC-N1-07 | F-Q1-N07 | `ACC07` |
| ACC-N1-08 | F-Q1-N08 | `ACC08`, `ACC08-NEXT` |
| ACC-N1-09 | F-Q1-N09 | `ACC09`, `EQUAL81`, `TARGET-COMPLETE-SOW` |
| ACC-N1-10 | F-Q1-N10 | `REP-P-COUNT1`, `REP-P-COUNT2`, `REP-P-COUNT3`, `REP-P-COUNT4` |
| ACC-N1-11 | F-Q1-N11 | `CLAIM-WHITE`, `CLAIM-BLACK`, `CLAIM-REJECT-THIRD-PARTY`, `CLAIM-REJECT-STALE`, `CLAIM-REJECT-TWO`, `CLAIM-REJECT-TERMINAL`, `CLAIM-REJECT-INCOMPLETE`, `RACE-MOVE`, `RACE-CLAIM` |
| ACC-N1-12 | F-Q1-N12 | `CLAIM-REJECT-STALE`, `POSITION-TUZ-IDENTITY`, `POSITION-IDENTITY`, `POSITION-HASH-COLLISION` |
| ACC-N1-13 | F-Q1-N13 | `HISTORY-PREVIEW`, `HISTORY-SIMULATION`, `HISTORY-INVALID_MOVE`, `HISTORY-HEARTBEAT`, `HISTORY-DUPLICATE_COMMAND`, `HISTORY-RESIGNATION`, `HISTORY-MUTUAL_AGREEMENT`, `HISTORY-MOVE-CLOSE` |
| ACC-N1-14 | F-Q1-N14 | `ACC09`, `REJECT-TERMINAL`, `HISTORY-MOVE-CLOSE`, `TERMINAL-VS-REPETITION` |

## Validación reproducible

Desde la raíz, ejecutar este bloque completo en PowerShell con el Python ya disponible. Solo lee JSON y objetos Git; no instala ni crea un helper. El PASS que emite se limita a datos. La revisión normativa, el motor y las integraciones futuras requieren evidencia aparte. Los estados de revisión leídos por el comando pertenecen al corte de autoría, no a la revisión posterior registrada en la issue o PR.

```powershell
@'
import json, hashlib, subprocess, re
from pathlib import Path
c=json.loads(Path('modules/engine/testdata/tk-n1-1.0.0/corpus.json').read_text(encoding='utf-8'))
assert c['normativeCorpusVersion']==1 and c['rulesVersion']=='TK-N1-1.0.0'
specs={s['caseId']:s for s in c['caseSpecs']}
assert len(specs)==len(c['caseSpecs'])==14
assert set(specs)=={'F-Q1-N%02d'%n for n in range(1,15)}
assert {s['acceptanceId'] for s in specs.values()}=={'ACC-N1-%02d'%n for n in range(1,15)}
es=c['entries']; byid={e['entryId']:e for e in es}; byfixture={e['fixtureId']:e for e in es}
assert len(byid)==len(byfixture)==len(es)
neg=[n['negativeCaseId'] for e in es for n in e['negativeCases']]
assert len(neg)==len(set(neg))
sourceids={s['sourceId'] for s in c['sources']}
for s in c['sources']:
    b=subprocess.check_output(['git','show',s['gitCommit']+':'+s['path']])
    assert hashlib.sha256(b).hexdigest()==s['sha256']
    assert s['digestConvention']=='GIT_BLOB_BYTES_NO_CHECKOUT_NEWLINE_CONVERSION'
def errors(s):
    out=[];b=s['board'];k=s['kazans'];t=s['tuzdyks']
    if len(b)!=2 or any(len(row)!=9 for row in b):out.append('shape')
    vals=[n for row in b for n in row]+k
    if any(type(n) is not int or n<0 or n>162 for n in vals):out.append('range')
    if sum(vals)!=162:out.append('total')
    if s['rulesVersion']!=c['rulesVersion']:out.append('rules')
    if s['stateSchemaVersion']!=1:out.append('schema')
    if s['turn'] not in ['white','black']:out.append('turn')
    if len(k)!=2 or len(t)!=2:out.append('slots')
    indices=[]
    for owner,x in enumerate(t):
        if x:
            a=x['at'];i=a['pitIndex']
            if x['owner']!=['white','black'][owner] or a['player']!=['black','white'][owner] or not 0<=i<8:out.append('tuz')
            elif b[1-owner][i]!=0:out.append('tuz-nonzero')
            indices.append(i)
    if len(indices)==2 and indices[0]==indices[1]:out.append('tuz-homologous')
    return out
moves=0
for e in es:
    assert e['caseId'] in specs and e['rulesVersion']==c['rulesVersion']
    assert e['stateSchemaVersion']==e['moveRecordVersion']==1
    assert all(n in specs for n in e['normativeCaseIds'])
    assert all(a==specs[n]['acceptanceId'] for a,n in zip(e['acceptanceIds'],e['normativeCaseIds']))
    assert all(s['sourceId'] in sourceids for s in e['sourceReference'])
    for f in ['initialState','expectedState']:
        s=e[f]
        if s is None:continue
        bad=errors(s)
        if f=='initialState' and e['inputValidity']=='INVALID_EXPECTED':assert bad
        else:assert not bad,(e['entryId'],f,bad)
    for v in e.get('comparisonVariants',[]):
        if v['inputValidity']!='INVALID_EXPECTED':assert not errors(v['state'])
    for n in e['negativeCases']:
        assert n['fixtureId']==e['fixtureId'] and n['scope']=='LOCAL_FIXTURE_ASSERTION'
        if e.get('manualWitness') and e['manualWitness']['depositedCount']!=19:
            assert not re.search(r'19\s*(?:dep[oó]sit|deposit)',n['predicate'],re.I)
        if e.get('expectedRejection'):
            assert not re.search(r'remain\s*legal|salida\s+legal|jugada\s+legal',n['predicate'],re.I)
    for ref in e['relatedFixtureReferences']:
        assert ref['fixtureId'] in byfixture and ref['fixtureId']!=e['fixtureId']
        assert ref['relationship']=='SEPARATE_VARIANT_NOT_EXERCISED_BY_THIS_ENTRY'
    if 'beforeGameRevision' in e['action']:
        assert e['action']['beforeGameRevision']==e['initialState']['revision']
        assert e['action']['afterGameRevision']==e['expectedState']['revision']
    if e.get('expectedRejection'):assert e['expectedState'] is None and e['expectedRejection']['inputPreserved']
    ev=e['expectedEvents']
    if ev is not None:
        assert [x['sequence'] for x in ev]==list(range(len(ev)))
        assert all(x['eventSchemaVersion']==1 for x in ev)
        if ev:assert all(x['revision']==e['expectedState']['revision'] for x in ev)
    for f in ['independentReview','engineExecution','gamePostgresExecution','concurrencyExecution']:assert e['verificationStatus'][f]=='NOT_RUN'
    if 'manualWitness' not in e:continue
    moves+=1;w=e['manualWitness'];a=e['initialState'];z=e['expectedState']
    b=[r[:] for r in a['board']];k=a['kazans'][:];t=a['tuzdyks'][:]
    player=['white','black'].index(w['origin']['player']);i=w['origin']['pitIndex']
    assert b[player][i]==w['originBefore'];b[player][i]=w['retainedAtOrigin']
    assert w['depositedCount']==len(w['deposits'])==(1 if w['originBefore']==1 else w['originBefore']-1)
    assert [d['ordinal'] for d in w['deposits']]==list(range(1,w['depositedCount']+1))
    for d in w['deposits']:
        side=['white','black'].index(d['destination']['player']);i=d['destination']['pitIndex']
        assert b[side][i]==d['before']
        if d['kind']=='ordinary':assert d['after']==d['before']+1;b[side][i]=d['after']
        else:
            owner=['white','black'].index(d['owner'])
            assert d['before']==d['after']==0 and k[owner]==d['kazanBefore']
            assert d['kazanAfter']==d['kazanBefore']+1;k[owner]=d['kazanAfter']
    for effect in w['postDepositEffects']:
        owner=['white','black'].index(effect['owner'])
        if effect['kind']=='capture':
            side=['white','black'].index(effect['from']['player']);i=effect['from']['pitIndex']
            assert b[side][i]==effect['amount'] and k[owner]==effect['kazanBefore']
            assert effect['kazanAfter']==k[owner]+effect['amount']
            b[side][i]=0;k[owner]=effect['kazanAfter']
        else:t[owner]={'owner':effect['owner'],'at':effect['at']}
    assert b==z['board'] and k==z['kazans'] and t==z['tuzdyks'],e['entryId']
    assert int(z['revision'])==int(a['revision'])+1
    assert w['beforeBoardSum']==sum(map(sum,a['board'])) and w['beforeKazanSum']==sum(a['kazans'])
    assert w['afterBoardSum']==sum(map(sum,z['board'])) and w['afterKazanSum']==sum(z['kazans'])
    assert sum(map(sum,b))+sum(k)==162
for fixture in ['HISTORY-RESIGNATION','HISTORY-MUTUAL_AGREEMENT']:
    effect=byfixture[fixture]['expectedApplicationEffect']
    assert effect['retryReusesReceipt'] is True and effect['retryLifecycleDelta']==0
    assert effect['retryPreconditions']==dict(sameActor=True,sameResource=True,sameCommandId=True,sameFingerprint=True,receiptCurrentlyReadable=True,currentActorResourceAuthorization=True)
    assert effect['retryReceiptLookupBeforeNewIntentTurnPhaseChecks'] is True
    assert effect['retryNegativeExpectations']==dict(differentFingerprintWithSameActorResourceCommandId=dict(priorReceiptReturned=False,newEffect=False,receiptReplaced=False,code='COMMAND_ID_REUSED'),differentActorOrResource=dict(sameIntention=False,priorReceiptReturned=False),receiptUnreadableOrCurrentAuthorizationMissing=dict(priorReceiptReturned=False,newEffectFromRetry=False))
for fixture in ['HISTORY-PREVIEW','HISTORY-SIMULATION','HISTORY-INVALID_MOVE','HISTORY-HEARTBEAT']:
    assert 'retryReusesReceipt' not in byfixture[fixture]['expectedApplicationEffect']
for s in specs.values():
    assert all(i in byid for i in s['positiveEntryIds']+s['negativeEntryIds'])
    assert all(f in byfixture for f in s['fixtureIds'])
    assert all(n in neg for n in s['negativeCaseIds'])
print('PASS data-only:',len(specs),'ACC;',len(es),'entries;',moves,'witnesses;',len(neg),'negative predicates')
print('NOT_RUN: engine, replay, Game, PostgreSQL, concurrency; independentReview is the dated author snapshot')
'@ | python -
```

El revisor comprobará manualmente destinos y última piedra, captura, prohibiciones, propietarios, eventos y resultado/turno contra R-01 a R-08. La suma 162 y el testigo no sustituyen ese cotejo. Propiedades de aliasing, fuzzing, cobertura, replay real y carreras SQL siguen pendientes de implementación y autorización posteriores.
