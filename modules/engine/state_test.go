package engine

import (
	"bytes"
	"encoding/json"
	"errors"
	"math"
	"os"
	"reflect"
	"strings"
	"testing"
)

// Independent literal: a missing row, wrong starting colour or hidden shared
// context in the factory must fail these tests, not change the oracle.
func initialLiteral() State {
	return State{StateSchemaVersion: 1, Position: Position{
		RulesVersion: "TK-N1-1.0.0", Turn: White,
		Board: [2][9]uint16{{9, 9, 9, 9, 9, 9, 9, 9, 9}, {9, 9, 9, 9, 9, 9, 9, 9, 9}},
	}}
}

func requireCode(t *testing.T, err error, code ErrorCode) {
	t.Helper()
	var rule *RuleError
	if !errors.As(err, &rule) || rule.Code != code || rule.Field == "" {
		t.Fatalf("want typed %s with field; got %v", code, err)
	}
}

func TestInitialState(t *testing.T) {
	s := NewInitialState()
	if s.Position != initialLiteral().Position || s.StateSchemaVersion != 1 || s.Revision != 0 || s.Terminal != nil {
		t.Fatalf("incorrect initial state: %+v", s)
	}
	if err := ValidateStructure(s); err != nil {
		t.Fatal(err)
	}
	if s.Repetition == nil || s.Repetition.Status != "complete" || len(s.Repetition.Positions) != 1 || s.Repetition.Positions[0].Count != 1 {
		t.Fatal("initial context must describe exactly one local initial occurrence")
	}
	if s.Repetition.Positions[0].Position != s.Position || !reflect.DeepEqual(s.Repetition.Positions[0].OccurrenceRefs, []string{"0"}) {
		t.Fatal("initial occurrence must identify the initial position")
	}
	second := NewInitialState()
	s.Repetition.Positions[0].OccurrenceRefs[0] = "changed"
	if second.Repetition.Positions[0].OccurrenceRefs[0] != "0" {
		t.Fatal("factory shares mutable storage")
	}
}

func TestPitRefCycle(t *testing.T) {
	ref, err := NewPitRef(White, 0)
	if err != nil {
		t.Fatal(err)
	}
	seen := map[PitRef]bool{}
	for step := range 18 {
		wantPlayer, wantIndex := White, step
		if step >= 9 {
			wantPlayer, wantIndex = Black, step-9
		}
		if ref.Player != wantPlayer || int(ref.PitIndex) != wantIndex || seen[ref] {
			t.Fatalf("wrong reference at step %d: %+v", step, ref)
		}
		seen[ref] = true
		ref, err = NextPit(ref)
		if err != nil {
			t.Fatal(err)
		}
	}
	if ref != (PitRef{Player: White, PitIndex: 0}) {
		t.Fatal("cycle does not close")
	}
	for _, p := range []Player{White, Black} {
		for i := range 9 {
			r, err := NewPitRef(p, i)
			if err != nil || r != (PitRef{Player: p, PitIndex: uint8(i)}) {
				t.Fatalf("bad constructor %v %d: %v", p, i, err)
			}
		}
	}
}

func TestPitRefRejectsBeforeNarrowing(t *testing.T) {
	for _, index := range []int{-1, 9, 256, math.MaxInt, math.MinInt} {
		_, err := NewPitRef(White, index)
		requireCode(t, err, InvalidPitIndex)
	}
	_, err := NewPitRef(Player(255), 0)
	requireCode(t, err, InvalidPlayer)
	_, err = NextPit(PitRef{Player: Player(255)})
	requireCode(t, err, InvalidPlayer)
	_, err = NextPit(PitRef{Player: White, PitIndex: 255})
	requireCode(t, err, InvalidPitIndex)
}

func TestValidateState(t *testing.T) {
	for _, tc := range []struct {
		name   string
		change func(*State)
		code   ErrorCode
	}{
		{"schema", func(s *State) { s.StateSchemaVersion = 2 }, UnsupportedStateVersion},
		{"rules", func(s *State) { s.RulesVersion = "unknown" }, UnsupportedRulesVersion},
		{"player", func(s *State) { s.Turn = 255 }, InvalidPlayer},
		{"total161", func(s *State) { s.Board[0][0] = 8 }, StoneTotalMismatch},
		{"total163", func(s *State) { s.Kazans[0] = 1 }, StoneTotalMismatch},
		{"stone163", func(s *State) { s.Board[0][0] = 163 }, InvalidStoneCount},
		{"wide-sum", func(s *State) {
			for p := range 2 {
				for i := range 9 {
					s.Board[p][i] = math.MaxUint16
				}
			}
		}, InvalidStoneCount},
		{"kazan163", func(s *State) { s.Kazans[1] = 163 }, InvalidStoneCount},
		{"tuz-own-field", func(s *State) { s.Tuzdyks[0] = &Tuzdyk{Owner: White, At: PitRef{White, 0}} }, InvalidTuzdyk},
		{"tuz-owner", func(s *State) { s.Tuzdyks[0] = &Tuzdyk{Owner: Black, At: PitRef{Black, 0}} }, InvalidTuzdyk},
		{"tuz-ninth", func(s *State) { s.Tuzdyks[0] = &Tuzdyk{Owner: White, At: PitRef{Black, 8}} }, InvalidTuzdyk},
		{"tuz-overflow-index", func(s *State) { s.Tuzdyks[0] = &Tuzdyk{Owner: White, At: PitRef{Black, 255}} }, InvalidTuzdyk},
		{"tuz-nonzero", func(s *State) { s.Tuzdyks[0] = &Tuzdyk{Owner: White, At: PitRef{Black, 0}} }, InvalidTuzdyk},
		{"tuz-homologous", func(s *State) {
			s.Board[0][0] = 0
			s.Board[1][0] = 0
			s.Kazans[0] = 18
			s.Tuzdyks = [2]*Tuzdyk{{White, PitRef{Black, 0}}, {Black, PitRef{White, 0}}}
		}, InvalidTuzdyk},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := initialLiteral()
			tc.change(&s)
			before := CloneState(s)
			requireCode(t, ValidateStructure(s), tc.code)
			if !reflect.DeepEqual(before, s) {
				t.Fatal("validation mutates input")
			}
		})
	}
	s := initialLiteral()
	s.Board[0][1] = 0
	s.Board[1][0] = 0
	s.Kazans = [2]uint16{9, 9}
	s.Tuzdyks = [2]*Tuzdyk{{White, PitRef{Black, 0}}, {Black, PitRef{White, 1}}}
	if err := ValidateStructure(s); err != nil {
		t.Fatalf("allowed tuzdyk pair rejected: %v", err)
	}
	// Synthetic positions are structurally valid, not proof of a legal history.
	s = initialLiteral()
	s.Board[0][0] = 0
	s.Kazans[0] = 9
	if err := ValidateStructure(s); err != nil {
		t.Fatal(err)
	}
	if s.Repetition != nil {
		t.Fatal("validation fabricates history")
	}
}

func TestCloneState(t *testing.T) {
	s := initialLiteral()
	s.Board[1][0] = 0
	s.Kazans[0] = 9
	s.Tuzdyks[0] = &Tuzdyk{White, PitRef{Black, 0}}
	winner := White
	s.Terminal = &Terminal{Outcome: "white_win", Reason: "resignation", Winner: &winner, Source: "application"}
	s.Repetition = &RepetitionContext{ContextVersion: 1, PolicyID: "TK-N1-1.0.0", Status: "incomplete", LedgerRef: "synthetic", Positions: []PositionOccurrence{{Position: s.Position, Count: 1, OccurrenceRefs: []string{"0"}}}}
	cloned := CloneState(s)
	if !reflect.DeepEqual(s, cloned) {
		t.Fatal("clone changes values")
	}
	cloned.Tuzdyks[0].At.PitIndex = 2
	*cloned.Terminal.Winner = Black
	cloned.Repetition.Positions[0].OccurrenceRefs[0] = "changed"
	cloned.Repetition.Positions[0].Position.Tuzdyks[0].At.PitIndex = 3
	cloned.Board[0][0] = 0
	if s.Tuzdyks[0].At.PitIndex != 0 || *s.Terminal.Winner != White || s.Board[0][0] != 9 || s.Repetition.Positions[0].OccurrenceRefs[0] != "0" || s.Repetition.Positions[0].Position.Tuzdyks[0].At.PitIndex != 0 {
		t.Fatal("clone aliases source")
	}
	s.Repetition.Positions = append(s.Repetition.Positions, PositionOccurrence{})
	if len(cloned.Repetition.Positions) != 1 {
		t.Fatal("clone shares slice header")
	}
	// Empty and absent contexts remain distinguishable across a copy.
	s = initialLiteral()
	if !reflect.DeepEqual(s, CloneState(s)) {
		t.Fatal("nil context changed")
	}
	s.Repetition = &RepetitionContext{Positions: []PositionOccurrence{}}
	if CloneState(s).Repetition.Positions == nil {
		t.Fatal("empty slice collapsed to nil")
	}
}

type corpusEntry struct {
	FixtureID     string          `json:"fixtureId"`
	InputValidity string          `json:"inputValidity"`
	Initial       json.RawMessage `json:"initialState"`
	Expected      json.RawMessage `json:"expectedState"`
}

func corpus(t testing.TB) []corpusEntry {
	t.Helper()
	data, err := os.ReadFile("testdata/tk-n1-1.0.0/corpus.json")
	if err != nil {
		t.Fatal(err)
	}
	var c struct {
		Entries []corpusEntry `json:"entries"`
	}
	if err = json.Unmarshal(data, &c); err != nil {
		t.Fatal(err)
	}
	return c.Entries
}

func initialJSON(t testing.TB) []byte {
	t.Helper()
	for _, e := range corpus(t) {
		if e.FixtureID == "NORM-INIT" {
			return e.Expected
		}
	}
	t.Fatal("NORM-INIT missing")
	return nil
}

func TestStateCodecCorpus(t *testing.T) {
	for _, e := range corpus(t) {
		t.Run(e.FixtureID, func(t *testing.T) {
			for i, data := range []json.RawMessage{e.Initial, e.Expected} {
				if len(data) == 0 || bytes.Equal(data, []byte("null")) {
					continue
				}
				before := bytes.Clone(data)
				s, err := DecodeState(data)
				if i == 0 && e.InputValidity == "INVALID_EXPECTED" {
					if err == nil {
						t.Fatal("invalid corpus input accepted")
					}
					continue
				}
				if err != nil {
					t.Fatal(err)
				}
				if err = ValidateStructure(s); err != nil {
					t.Fatal(err)
				}
				encoded, err := EncodeState(s)
				if err != nil {
					t.Fatal(err)
				}
				again, err := DecodeState(encoded)
				if err != nil || !reflect.DeepEqual(s, again) {
					t.Fatalf("round trip changed state: %v", err)
				}
				if !bytes.Equal(data, before) {
					t.Fatal("codec mutates bytes")
				}
				if e.FixtureID == "NORM-INIT" && s.Position != initialLiteral().Position {
					t.Fatal("corpus initial disagrees with factory oracle")
				}
			}
		})
	}
}

func TestStateCodecAdversarial(t *testing.T) {
	base := initialJSON(t)
	for _, tc := range []struct {
		name   string
		change func(map[string]any)
		code   ErrorCode
	}{
		{"board8", func(m map[string]any) { m["board"].([]any)[0] = []any{9, 9} }, InvalidStateShape},
		{"board3", func(m map[string]any) { m["board"] = append(m["board"].([]any), []any{}) }, InvalidStateShape},
		{"kazans1", func(m map[string]any) { m["kazans"] = []int{0} }, InvalidStateShape},
		{"tuzdyks3", func(m map[string]any) { m["tuzdyks"] = []any{nil, nil, nil} }, InvalidStateShape},
		{"absent", func(m map[string]any) { delete(m, "turn") }, InvalidStateShape},
		{"null", func(m map[string]any) { m["board"] = nil }, InvalidStateShape},
		{"nullstone", func(m map[string]any) { m["board"].([]any)[0].([]any)[0] = nil }, InvalidStoneCount},
		{"negative", func(m map[string]any) { m["board"].([]any)[0].([]any)[0] = -1 }, InvalidStoneCount},
		{"fraction", func(m map[string]any) { m["board"].([]any)[0].([]any)[0] = 1.5 }, InvalidStoneCount},
		{"stringstone", func(m map[string]any) { m["board"].([]any)[0].([]any)[0] = "9" }, InvalidStoneCount},
		{"hugestone", func(m map[string]any) { m["board"].([]any)[0].([]any)[0] = json.Number("999999999999999999999") }, InvalidStoneCount},
		{"extra", func(m map[string]any) { m["secret-input"] = "not echoed" }, InvalidStateShape},
		{"nullschema", func(m map[string]any) { m["stateSchemaVersion"] = nil }, InvalidStateShape},
		{"float-schema", func(m map[string]any) { m["stateSchemaVersion"] = 1.5 }, UnsupportedStateVersion},
		{"numeric-player", func(m map[string]any) { m["turn"] = 0 }, InvalidPlayer},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var m map[string]any
			_ = json.Unmarshal(base, &m)
			tc.change(m)
			data, err := json.Marshal(m)
			if err != nil {
				t.Fatal(err)
			}
			saved := bytes.Clone(data)
			_, err = DecodeState(data)
			requireCode(t, err, tc.code)
			if !bytes.Equal(data, saved) {
				t.Fatal("reject mutates bytes")
			}
			if strings.Contains(err.Error(), "not echoed") {
				t.Fatal("error echoes data")
			}
		})
	}
	for name, data := range map[string][]byte{
		"duplicate":        append(append([]byte(`{"turn":"black",`), bytes.TrimSpace(base)[1:]...), byte(' ')),
		"nested duplicate": bytes.Replace(base, []byte(`"revision": "0"`), []byte(`"revision":"0","repetition":{"contextVersion":1,"contextVersion":1}`), 1),
		"trailing":         append(bytes.Clone(base), []byte(` {}`)...),
		"null state":       []byte(`null`), "nonobject": []byte(`[]`), "syntax": []byte(`{"board":`),
		"utf8": []byte{'"', 255, '"'}, "surrogate": []byte(`{"bad":"\ud800"}`),
		"depth": []byte(strings.Repeat("[", 100) + strings.Repeat("]", 100)),
	} {
		t.Run(name, func(t *testing.T) {
			_, err := DecodeState(data)
			if err == nil {
				t.Fatal("malformed JSON accepted")
			}
		})
	}
}

func TestStateCodecRevision(t *testing.T) {
	for _, n := range []uint64{0, 9007199254740993, math.MaxUint64} {
		s := initialLiteral()
		s.Revision = n
		b, err := EncodeState(s)
		if err != nil {
			t.Fatal(err)
		}
		var m map[string]json.RawMessage
		_ = json.Unmarshal(b, &m)
		if len(m["revision"]) < 2 || m["revision"][0] != '"' {
			t.Fatal("revision is not a JSON string")
		}
		got, err := DecodeState(b)
		if err != nil || got.Revision != n {
			t.Fatalf("lost uint64 precision %d: %v", n, err)
		}
	}
	for _, v := range []string{`"-1"`, `"18446744073709551616"`, `"1.0"`, `"1e2"`, `"01"`, `"+1"`, `""`, `null`, `0`, `1.5`, `" 1"`} {
		t.Run(v, func(t *testing.T) {
			var m map[string]json.RawMessage
			_ = json.Unmarshal(initialJSON(t), &m)
			m["revision"] = json.RawMessage(v)
			b, _ := json.Marshal(m)
			_, err := DecodeState(b)
			requireCode(t, err, InvalidRevision)
		})
	}
}

func TestStateCodecTerminalAndContext(t *testing.T) {
	for _, tc := range []Terminal{
		{Outcome: "white_win", Reason: "target_reached", Winner: playerPtr(White), Source: "rules"},
		{Outcome: "black_win", Reason: "atsyrau", Winner: playerPtr(Black), Source: "rules"},
		{Outcome: "draw", Reason: "atsyrau", Source: "rules"},
		{Outcome: "draw", Reason: "equal_score", Source: "rules"},
		{Outcome: "draw", Reason: "repetition_claim", Source: "rules"},
		{Outcome: "draw", Reason: "mutual_agreement", Source: "application"},
		{Outcome: "black_win", Reason: "resignation", Winner: playerPtr(Black), Source: "application"},
		{Outcome: "inconclusive", Reason: "recovery_expired", Source: "application"},
	} {
		s := NewInitialState()
		s.Terminal = &tc
		b, err := EncodeState(s)
		if err != nil {
			t.Fatal(err)
		}
		decoded, err := DecodeState(b)
		if err != nil || !reflect.DeepEqual(s, decoded) {
			t.Fatalf("typed data lost: %v", err)
		}
	}
	for _, tc := range []struct {
		name   string
		change func(*State)
	}{
		{"unknown outcome", func(s *State) { s.Terminal = &Terminal{Outcome: "arbitrary"} }},
		{"draw winner", func(s *State) {
			s.Terminal = &Terminal{Outcome: "draw", Reason: "atsyrau", Winner: playerPtr(White), Source: "rules"}
		}},
		{"win missing winner", func(s *State) {
			s.Terminal = &Terminal{Outcome: "white_win", Reason: "target_reached", Source: "rules"}
		}},
		{"winner mismatch", func(s *State) {
			s.Terminal = &Terminal{Outcome: "white_win", Reason: "target_reached", Winner: playerPtr(Black), Source: "rules"}
		}},
		{"source mismatch", func(s *State) { s.Terminal = &Terminal{Outcome: "draw", Reason: "mutual_agreement", Source: "rules"} }},
		{"context version", func(s *State) { s.Repetition.ContextVersion = 2 }},
		{"context policy", func(s *State) { s.Repetition.PolicyID = "unknown" }},
		{"context status", func(s *State) { s.Repetition.Status = "arbitrary" }},
		{"context position", func(s *State) { s.Repetition.Positions[0].Position.Board[0][0] = 8 }},
		{"zero count", func(s *State) { s.Repetition.Positions[0].Count = 0 }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := NewInitialState()
			if s.Repetition == nil {
				t.Fatal("factory context missing")
			}
			tc.change(&s)
			if _, err := EncodeState(s); err == nil {
				t.Fatal("encoder accepts invalid structure")
			}
			if ValidateStructure(s) == nil {
				t.Fatal("invalid structure accepted")
			}
		})
	}
	s := initialLiteral()
	s.Repetition = &RepetitionContext{ContextVersion: 1, PolicyID: "TK-N1-1.0.0", Status: "incomplete", LedgerRef: "synthetic", Positions: []PositionOccurrence{}}
	b, err := EncodeState(s)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := DecodeState(b)
	if err != nil || !reflect.DeepEqual(s, decoded) {
		t.Fatalf("incomplete context altered: %v", err)
	}
	// A count of three is data only: decoding must not adjudicate a draw.
	s = NewInitialState()
	s.Repetition.Positions[0].Count = 3
	b, err = EncodeState(s)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err = DecodeState(b)
	if err != nil || decoded.Terminal != nil {
		t.Fatal("codec adjudicates repetition")
	}
}

func playerPtr(p Player) *Player { return &p }

// These boundary cases exercise shared guards through the public codec. A
// missing nested shape/type/duplicate guard must not hide behind an invalid
// outer object or an already-invalid terminal.
func TestStateCodecNestedBoundaries(t *testing.T) {
	base, err := EncodeState(NewInitialState())
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name   string
		path   []string
		value  any
		remove bool
	}{
		{"context nonobject", []string{"repetition"}, []any{}, false},
		{"context null version", []string{"repetition", "contextVersion"}, nil, false},
		{"context bad policy type", []string{"repetition", "policyId"}, 42, false},
		{"context bad status type", []string{"repetition", "status"}, true, false},
		{"context bad ledger type", []string{"repetition", "ledgerRef"}, 42, false},
		{"context missing ledger", []string{"repetition", "ledgerRef"}, nil, true},
		{"context extra", []string{"repetition", "extra"}, "not echoed", false},
		{"range nonobject", []string{"repetition", "coveredRange"}, []any{}, false},
		{"range from", []string{"repetition", "coveredRange", "from"}, "-1", false},
		{"range to", []string{"repetition", "coveredRange", "to"}, "18446744073709551616", false},
		{"range reversed", []string{"repetition", "coveredRange", "from"}, "1", false},
		{"positions nonarray", []string{"repetition", "positions"}, map[string]any{}, false},
		{"positions null", []string{"repetition", "positions"}, nil, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var m map[string]any
			_ = json.Unmarshal(base, &m)
			target := m
			for _, key := range tc.path[:len(tc.path)-1] {
				target = target[key].(map[string]any)
			}
			key := tc.path[len(tc.path)-1]
			if tc.remove {
				delete(target, key)
			} else {
				target[key] = tc.value
			}
			data, _ := json.Marshal(m)
			if _, err := DecodeState(data); err == nil {
				t.Fatal("invalid nested value accepted")
			}
		})
	}
	for name, value := range map[string]any{
		"entry nonobject":    nil,
		"position nonobject": map[string]any{"position": []any{}, "count": 1, "occurrenceRefs": []string{"0"}},
		"position malformed": map[string]any{"position": map[string]any{}, "count": 1, "occurrenceRefs": []string{"0"}},
		"count overflow":     map[string]any{"count": json.Number("18446744073709551616")},
		"refs nonarray":      map[string]any{"occurrenceRefs": "0"},
		"refs nonstring":     map[string]any{"occurrenceRefs": []any{1}},
		"refs empty token":   map[string]any{"occurrenceRefs": []string{""}},
	} {
		t.Run(name, func(t *testing.T) {
			var m map[string]any
			_ = json.Unmarshal(base, &m)
			c := m["repetition"].(map[string]any)
			entry := c["positions"].([]any)[0]
			if patch, ok := value.(map[string]any); ok {
				for key, v := range patch {
					entry.(map[string]any)[key] = v
				}
			} else {
				entry = value
			}
			c["positions"] = []any{entry}
			data, _ := json.Marshal(m)
			if _, err := DecodeState(data); err == nil {
				t.Fatal("malformed occurrence accepted")
			}
		})
	}
	for name, replacement := range map[string]string{
		"nested duplicate":         `"contextVersion":1,"contextVersion":1`,
		"escaped nested duplicate": `"contextVersion":1,"\u0063ontextVersion":1`,
	} {
		t.Run(name, func(t *testing.T) {
			data := bytes.Replace(base, []byte(`"contextVersion":1`), []byte(replacement), 1)
			if _, err := DecodeState(data); err == nil {
				t.Fatal("duplicate nested key accepted")
			}
		})
	}
	for name, value := range map[string]any{
		"terminal nonobject":      []any{},
		"terminal missing winner": map[string]any{"outcome": "draw", "reason": "atsyrau", "source": "rules"},
		"terminal bad types":      map[string]any{"outcome": 1, "reason": "atsyrau", "source": "rules", "winner": nil},
		"terminal bad winner":     map[string]any{"outcome": "white_win", "reason": "atsyrau", "source": "rules", "winner": "red"},
	} {
		t.Run(name, func(t *testing.T) {
			var m map[string]any
			_ = json.Unmarshal(base, &m)
			m["terminal"] = value
			data, _ := json.Marshal(m)
			if _, err := DecodeState(data); err == nil {
				t.Fatal("malformed terminal accepted")
			}
		})
	}
	for _, escaped := range []string{`\udc00`, `\ud800x`, `\ud800\u0041`, `\ud800\uXXXX`, `\uXXXX`} {
		data := bytes.Replace(base, []byte(`"local-initial"`), []byte(`"`+escaped+`"`), 1)
		if _, err := DecodeState(data); err == nil {
			t.Fatalf("invalid Unicode escape accepted: %s", escaped)
		}
	}
	for _, escaped := range []string{`\ud83d\ude00`, `a\/b`, `literal-\\ud800`, `\u0061`} {
		data := bytes.Replace(base, []byte(`"local-initial"`), []byte(`"`+escaped+`"`), 1)
		s, err := DecodeState(data)
		if err != nil {
			t.Fatalf("valid escape rejected: %s: %v", escaped, err)
		}
		encoded, err := EncodeState(s)
		if err != nil {
			t.Fatal(err)
		}
		decoded, err := DecodeState(encoded)
		if err != nil || !reflect.DeepEqual(s, decoded) {
			t.Fatal("valid Unicode round trip changed data")
		}
	}
}

func FuzzDecodeState(f *testing.F) {
	f.Add(initialJSON(f))
	f.Add([]byte(`null`))
	f.Add([]byte(`{"revision":"18446744073709551616"}`))
	for _, e := range corpus(f) {
		if e.Initial != nil {
			f.Add([]byte(e.Initial))
		}
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		saved := bytes.Clone(data)
		s, err := DecodeState(data)
		if !bytes.Equal(saved, data) {
			t.Fatal("input mutated")
		}
		if err != nil {
			return
		}
		if err = ValidateStructure(s); err != nil {
			t.Fatal(err)
		}
		encoded, err := EncodeState(s)
		if err != nil {
			t.Fatal(err)
		}
		again, err := DecodeState(encoded)
		if err != nil || !reflect.DeepEqual(s, again) {
			t.Fatalf("round trip: %v", err)
		}
	})
}
