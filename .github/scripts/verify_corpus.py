#!/usr/bin/env python3
"""CI-01: immutable approved corpus, data-only checks, no engine execution.

The static assertions below preserve the approved README validator verbatim
apart from strict JSON loading and safer git-show arguments. No Markdown code
is executed at runtime. Run from the repository root with full source history.
"""
import collections
import sys

if not __debug__:
    raise RuntimeError("Validation must not run with Python optimizations")

def unique_object(pairs):
    result = {}
    for key, value in pairs:
        if key in result:
            raise ValueError("Duplicate JSON object key")
        result[key] = value
    return result

import json, hashlib, subprocess, re
from pathlib import Path
c=json.loads(Path('modules/engine/testdata/tk-n1-1.0.0/corpus.json').read_text(encoding='utf-8'), object_pairs_hook=unique_object)
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
    assert re.fullmatch(r'[0-9a-f]{40}', s['gitCommit'])
    assert s['path'].startswith('docs/') and '..' not in s['path'].split('/') and ':' not in s['path']
    b=subprocess.check_output(['git','show','--end-of-options',s['gitCommit']+':'+s['path']])
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
# CI-01 strengthens (and never replaces) the assertions above.
assert len(es) == 66 and moves == 19 and len(neg) == 102
assert len(sourceids) == len(c['sources']) == 9
assert sum(e['inputValidity'] == 'INVALID_EXPECTED' for e in es) == 14
assert sum(len(e['expectedEvents'] or []) for e in es) == 122
for entry in es:
    assert len(entry['acceptanceIds']) == len(entry['normativeCaseIds'])
assert hashlib.sha256(Path('modules/engine/testdata/tk-n1-1.0.0/corpus.json').read_bytes()).hexdigest() == '5f26a6c6d64fb2b03db407737c6d96712ab17c5fc4d3f80ae676508ee743a66e'
assert hashlib.sha256(Path('modules/engine/testdata/tk-n1-1.0.0/README.md').read_bytes()).hexdigest() == 'fafe2de1550e66a898fc57732ea167a0982d3dd08d13c7ea6cf3a3da30a56e04'
print('PASS corpus-integrity: approved bytes, duplicate keys, counts, sources, IDs/FKs, states, witnesses, events and negatives')
