// Package engine contains the pure TK-N1 state model. Structural validation
// never proves reachability, trusted history, or authority to continue a game.
package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Player uint8

const (
	White Player = iota
	Black
)

type PitRef struct {
	Player   Player `json:"player"`
	PitIndex uint8  `json:"pitIndex"`
}
type Tuzdyk struct {
	Owner Player `json:"owner"`
	At    PitRef `json:"at"`
}
type Position struct {
	RulesVersion string
	Board        [2][9]uint16
	Kazans       [2]uint16
	Tuzdyks      [2]*Tuzdyk
	Turn         Player
}
type State struct {
	StateSchemaVersion uint32
	Revision           uint64
	Position
	Terminal   *Terminal
	Repetition *RepetitionContext
}
type Outcome string
type TerminalReason string
type TerminalSource string
type Terminal struct {
	Outcome Outcome        `json:"outcome"`
	Reason  TerminalReason `json:"reason"`
	Winner  *Player        `json:"winner"`
	Source  TerminalSource `json:"source"`
}
type CoveredRange struct{ From, To uint64 }
type PositionOccurrence struct {
	Position       Position
	Count          uint64
	OccurrenceRefs []string
}

// RepetitionContext carries data only; complete is not a trust decision.
type RepetitionContext struct {
	ContextVersion uint32
	PolicyID       string
	Status         string
	LedgerRef      string
	CoveredRange   CoveredRange
	Positions      []PositionOccurrence
}
type ErrorCode string

const (
	InvalidStateShape       ErrorCode = "INVALID_STATE_SHAPE"
	InvalidPlayer           ErrorCode = "INVALID_PLAYER"
	InvalidPitIndex         ErrorCode = "INVALID_PIT_INDEX"
	InvalidStoneCount       ErrorCode = "INVALID_STONE_COUNT"
	StoneTotalMismatch      ErrorCode = "STONE_TOTAL_MISMATCH"
	InvalidTuzdyk           ErrorCode = "INVALID_TUZDYK"
	UnsupportedStateVersion ErrorCode = "UNSUPPORTED_STATE_VERSION"
	UnsupportedRulesVersion ErrorCode = "UNSUPPORTED_RULES_VERSION"
	InvalidRevision         ErrorCode = "INVALID_REVISION"
)

type RuleError struct {
	Code  ErrorCode
	Field string
}

func (e *RuleError) Error() string { return string(e.Code) + ": " + e.Field }

const rulesVersion = "TK-N1-1.0.0"

func failure(code ErrorCode, field string) error { return &RuleError{code, field} }

// NewInitialState creates independent local data, not a confirmed Game ledger.
func NewInitialState() State {
	s := State{StateSchemaVersion: 1, Position: Position{RulesVersion: rulesVersion, Turn: White}}
	for player := range s.Board {
		for pit := range s.Board[player] {
			s.Board[player][pit] = 9
		}
	}
	s.Repetition = &RepetitionContext{ContextVersion: 1, PolicyID: rulesVersion, Status: "complete", LedgerRef: "local-initial", Positions: []PositionOccurrence{{Position: s.Position, Count: 1, OccurrenceRefs: []string{"0"}}}}
	return s
}

// NewPitRef checks the signed external index before narrowing to uint8.
func NewPitRef(player Player, index int) (PitRef, error) {
	if player != White && player != Black {
		return PitRef{}, failure(InvalidPlayer, "player")
	}
	if index < 0 || index >= 9 {
		return PitRef{}, failure(InvalidPitIndex, "pitIndex")
	}
	return PitRef{player, uint8(index)}, nil
}

// NextPit is a geometric successor; it neither sows nor changes a state.
func NextPit(ref PitRef) (PitRef, error) {
	if _, err := NewPitRef(ref.Player, int(ref.PitIndex)); err != nil {
		return PitRef{}, err
	}
	if ref.PitIndex < 8 {
		return PitRef{ref.Player, ref.PitIndex + 1}, nil
	}
	return PitRef{1 - ref.Player, 0}, nil
}

func validatePosition(p Position, field string) error {
	if p.RulesVersion != rulesVersion {
		return failure(UnsupportedRulesVersion, field+".rulesVersion")
	}
	if p.Turn != White && p.Turn != Black {
		return failure(InvalidPlayer, field+".turn")
	}
	var total uint32
	for player, row := range p.Board {
		for pit, count := range row {
			if count > 162 {
				return failure(InvalidStoneCount, fmt.Sprintf("%s.board[%d][%d]", field, player, pit))
			}
			total += uint32(count)
		}
	}
	for player, count := range p.Kazans {
		if count > 162 {
			return failure(InvalidStoneCount, fmt.Sprintf("%s.kazans[%d]", field, player))
		}
		total += uint32(count)
	}
	if total != 162 {
		return failure(StoneTotalMismatch, field)
	}
	for owner, tuz := range p.Tuzdyks {
		if tuz != nil {
			if tuz.Owner != Player(owner) || tuz.At.Player != Player(1-owner) || tuz.At.PitIndex >= 8 {
				return failure(InvalidTuzdyk, field+".tuzdyks")
			}
			if p.Board[tuz.At.Player][tuz.At.PitIndex] != 0 {
				return failure(InvalidTuzdyk, field+".tuzdyks")
			}
		}
	}
	if p.Tuzdyks[0] != nil && p.Tuzdyks[1] != nil && p.Tuzdyks[0].At.PitIndex == p.Tuzdyks[1].At.PitIndex {
		return failure(InvalidTuzdyk, field+".tuzdyks")
	}
	return nil
}

// ValidateStructure is deliberately not ValidateHistory. A complete context or
// a terminal is checked as data; no outcome or repetition right is inferred.
func ValidateStructure(state State) error {
	if state.StateSchemaVersion != 1 {
		return failure(UnsupportedStateVersion, "stateSchemaVersion")
	}
	if err := validatePosition(state.Position, "state"); err != nil {
		return err
	}
	if terminal := state.Terminal; terminal != nil {
		win := terminal.Outcome == "white_win" || terminal.Outcome == "black_win"
		if win {
			want := White
			if terminal.Outcome == "black_win" {
				want = Black
			}
			if terminal.Winner == nil || *terminal.Winner != want {
				return failure(InvalidStateShape, "terminal.winner")
			}
		} else if (terminal.Outcome != "draw" && terminal.Outcome != "inconclusive") || terminal.Winner != nil {
			return failure(InvalidStateShape, "terminal.outcome")
		}
		valid := false
		switch terminal.Reason {
		case "target_reached":
			valid = win && terminal.Source == "rules"
		case "atsyrau":
			valid = (win || terminal.Outcome == "draw") && terminal.Source == "rules"
		case "equal_score", "repetition_claim":
			valid = terminal.Outcome == "draw" && terminal.Source == "rules"
		case "resignation":
			valid = win && terminal.Source == "application"
		case "mutual_agreement":
			valid = terminal.Outcome == "draw" && terminal.Source == "application"
		case "recovery_expired":
			valid = terminal.Outcome == "inconclusive" && terminal.Source == "application"
		}
		if !valid {
			return failure(InvalidStateShape, "terminal.reason")
		}
	}
	if context := state.Repetition; context != nil {
		if context.ContextVersion != 1 || context.PolicyID != rulesVersion || (context.Status != "complete" && context.Status != "incomplete") || context.LedgerRef == "" || !utf8.ValidString(context.LedgerRef) || context.CoveredRange.From > context.CoveredRange.To || context.Positions == nil {
			return failure(InvalidStateShape, "repetition")
		}
		for _, entry := range context.Positions {
			if err := validatePosition(entry.Position, "repetition.position"); err != nil {
				return err
			}
			if entry.Count == 0 || entry.OccurrenceRefs == nil {
				return failure(InvalidStateShape, "repetition.positions")
			}
			for _, ref := range entry.OccurrenceRefs {
				if ref == "" || !utf8.ValidString(ref) {
					return failure(InvalidStateShape, "repetition.occurrenceRefs")
				}
			}
		}
	}
	return nil
}

func clonePosition(p Position) Position {
	for i, tuz := range p.Tuzdyks {
		if tuz != nil {
			copy := *tuz
			p.Tuzdyks[i] = &copy
		}
	}
	return p
}

// CloneState owns every optional and slice in the returned value.
func CloneState(state State) State {
	state.Position = clonePosition(state.Position)
	if state.Terminal != nil {
		copy := *state.Terminal
		state.Terminal = &copy
		if copy.Winner != nil {
			winner := *copy.Winner
			copy.Winner = &winner
		}
	}
	if state.Repetition != nil {
		copy := *state.Repetition
		state.Repetition = &copy
		if copy.Positions != nil {
			copy.Positions = append(make([]PositionOccurrence, 0, len(copy.Positions)), copy.Positions...)
			for i := range copy.Positions {
				entry := &copy.Positions[i]
				entry.Position = clonePosition(entry.Position)
				if entry.OccurrenceRefs != nil {
					entry.OccurrenceRefs = append(make([]string, 0, len(entry.OccurrenceRefs)), entry.OccurrenceRefs...)
				}
			}
		}
	}
	return state
}

// Player's wire spelling is independent of presentation and user identity.
func (p Player) MarshalJSON() ([]byte, error) {
	if p == White {
		return []byte(`"white"`), nil
	}
	if p == Black {
		return []byte(`"black"`), nil
	}
	return nil, failure(InvalidPlayer, "player")
}

func positionJSON(p Position) map[string]any {
	return map[string]any{"rulesVersion": p.RulesVersion, "board": p.Board, "kazans": p.Kazans, "tuzdyks": p.Tuzdyks, "turn": p.Turn}
}

// EncodeState implements the local fixture codec, not a public Game DTO.
func EncodeState(state State) ([]byte, error) {
	if err := ValidateStructure(state); err != nil {
		return nil, err
	}
	m := positionJSON(state.Position)
	m["stateSchemaVersion"] = state.StateSchemaVersion
	m["revision"] = strconv.FormatUint(state.Revision, 10)
	m["terminal"] = state.Terminal
	if context := state.Repetition; context != nil {
		positions := make([]any, 0, len(context.Positions))
		for _, entry := range context.Positions {
			positions = append(positions, map[string]any{"position": positionJSON(entry.Position), "count": entry.Count, "occurrenceRefs": entry.OccurrenceRefs})
		}
		m["repetition"] = map[string]any{"contextVersion": context.ContextVersion, "policyId": context.PolicyID, "status": context.Status, "ledgerRef": context.LedgerRef, "coveredRange": map[string]string{"from": strconv.FormatUint(context.CoveredRange.From, 10), "to": strconv.FormatUint(context.CoveredRange.To, 10)}, "positions": positions}
	}
	return json.Marshal(m)
}

// readValue retains integer tokens exactly and rejects duplicate decoded keys
// at every level. Depth bounds malformed JSON, not the length of a game ledger.
func readValue(dec *json.Decoder, depth int) (any, error) {
	if depth > 64 {
		return nil, failure(InvalidStateShape, "state")
	}
	token, err := dec.Token()
	if err != nil {
		return nil, failure(InvalidStateShape, "state")
	}
	if delimiter, ok := token.(json.Delim); ok {
		switch delimiter {
		case '{':
			m := map[string]any{}
			for dec.More() {
				key, err := dec.Token()
				if err != nil {
					return nil, failure(InvalidStateShape, "state")
				}
				name, ok := key.(string)
				if !ok {
					return nil, failure(InvalidStateShape, "state")
				}
				if _, exists := m[name]; exists {
					return nil, failure(InvalidStateShape, "state")
				}
				v, err := readValue(dec, depth+1)
				if err != nil {
					return nil, err
				}
				m[name] = v
			}
			end, err := dec.Token()
			if err != nil || end != json.Delim('}') {
				return nil, failure(InvalidStateShape, "state")
			}
			return m, nil
		case '[':
			values := []any{}
			for dec.More() {
				v, err := readValue(dec, depth+1)
				if err != nil {
					return nil, err
				}
				values = append(values, v)
			}
			end, err := dec.Token()
			if err != nil || end != json.Delim(']') {
				return nil, failure(InvalidStateShape, "state")
			}
			return values, nil
		default:
			return nil, failure(InvalidStateShape, "state")
		}
	}
	return token, nil
}

// encoding/json replaces malformed UTF-16 escape sequences. Reject them before
// decoding, so identifiers cannot silently acquire replacement characters.
func validUnicode(data []byte) bool {
	if !utf8.Valid(data) {
		return false
	}
	for i := 0; i < len(data); i++ {
		if data[i] != '"' {
			continue
		}
		i++
		for i < len(data) && data[i] != '"' {
			if data[i] == '\\' {
				i++
				if i >= len(data) {
					return false
				}
				if data[i] == 'u' {
					if i+4 >= len(data) {
						return false
					}
					n, err := strconv.ParseUint(string(data[i+1:i+5]), 16, 16)
					if err != nil {
						return false
					}
					i += 4
					if n >= 0xD800 && n <= 0xDBFF {
						if i+6 >= len(data) || data[i+1] != '\\' || data[i+2] != 'u' {
							return false
						}
						low, err := strconv.ParseUint(string(data[i+3:i+7]), 16, 16)
						if err != nil || low < 0xDC00 || low > 0xDFFF {
							return false
						}
						i += 6
					} else if n >= 0xDC00 && n <= 0xDFFF {
						return false
					}
				}
			}
			i++
		}
	}
	return true
}

func object(value any, required, optional, field string) (map[string]any, error) {
	m, ok := value.(map[string]any)
	if !ok {
		return nil, failure(InvalidStateShape, field)
	}
	allowed := map[string]bool{}
	for _, name := range strings.Fields(required) {
		allowed[name] = true
		v, exists := m[name]
		if !exists || (v == nil && name != "terminal" && name != "winner" && name != "revision") {
			return nil, failure(InvalidStateShape, field)
		}
	}
	for _, name := range strings.Fields(optional) {
		allowed[name] = true
	}
	for name := range m {
		if !allowed[name] {
			return nil, failure(InvalidStateShape, field)
		}
	}
	return m, nil
}

func array(value any, length int, field string) ([]any, error) {
	items, ok := value.([]any)
	if !ok || (length >= 0 && len(items) != length) {
		return nil, failure(InvalidStateShape, field)
	}
	return items, nil
}

func unsigned(value any, max uint64, code ErrorCode, field string) (uint64, error) {
	number, ok := value.(json.Number)
	if !ok {
		return 0, failure(code, field)
	}
	n, err := strconv.ParseUint(string(number), 10, 64)
	if err != nil || n > max {
		return 0, failure(code, field)
	}
	return n, nil
}

func decimal(value any, field string) (uint64, error) {
	s, ok := value.(string)
	if !ok || s == "" || (len(s) > 1 && s[0] == '0') {
		return 0, failure(InvalidRevision, field)
	}
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0, failure(InvalidRevision, field)
		}
	}
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, failure(InvalidRevision, field)
	}
	return n, nil
}

func player(value any, field string) (Player, error) {
	if value == "white" {
		return White, nil
	}
	if value == "black" {
		return Black, nil
	}
	return 0, failure(InvalidPlayer, field)
}

const positionFields = "rulesVersion board kazans tuzdyks turn"

func decodePosition(m map[string]any, field string) (Position, error) {
	var p Position
	rules, ok := m["rulesVersion"].(string)
	if !ok || rules != rulesVersion {
		return p, failure(UnsupportedRulesVersion, field+".rulesVersion")
	}
	p.RulesVersion = rules
	turn, err := player(m["turn"], field+".turn")
	if err != nil {
		return p, err
	}
	p.Turn = turn
	rows, err := array(m["board"], 2, field+".board")
	if err != nil {
		return p, err
	}
	for side, row := range rows {
		pits, err := array(row, 9, field+".board")
		if err != nil {
			return p, err
		}
		for pit, value := range pits {
			n, err := unsigned(value, 162, InvalidStoneCount, field+".board")
			if err != nil {
				return p, err
			}
			p.Board[side][pit] = uint16(n)
		}
	}
	kazans, err := array(m["kazans"], 2, field+".kazans")
	if err != nil {
		return p, err
	}
	for side, value := range kazans {
		n, err := unsigned(value, 162, InvalidStoneCount, field+".kazans")
		if err != nil {
			return p, err
		}
		p.Kazans[side] = uint16(n)
	}
	tuzdyks, err := array(m["tuzdyks"], 2, field+".tuzdyks")
	if err != nil {
		return p, err
	}
	for slot, value := range tuzdyks {
		if value == nil {
			continue
		}
		tuz, err := object(value, "owner at", "", field+".tuzdyks")
		if err != nil {
			return p, err
		}
		at, err := object(tuz["at"], "player pitIndex", "", field+".tuzdyks.at")
		if err != nil {
			return p, err
		}
		owner, err := player(tuz["owner"], field+".tuzdyks.owner")
		if err != nil {
			return p, err
		}
		side, err := player(at["player"], field+".tuzdyks.at.player")
		if err != nil {
			return p, err
		}
		index, err := unsigned(at["pitIndex"], 8, InvalidTuzdyk, field+".tuzdyks.at.pitIndex")
		if err != nil {
			return p, err
		}
		p.Tuzdyks[slot] = &Tuzdyk{owner, PitRef{side, uint8(index)}}
	}
	return p, nil
}

func decodeTerminal(value any) (*Terminal, error) {
	if value == nil {
		return nil, nil
	}
	m, err := object(value, "outcome reason winner source", "", "terminal")
	if err != nil {
		return nil, err
	}
	outcome, o := m["outcome"].(string)
	reason, r := m["reason"].(string)
	source, s := m["source"].(string)
	if !o || !r || !s {
		return nil, failure(InvalidStateShape, "terminal")
	}
	terminal := &Terminal{Outcome: Outcome(outcome), Reason: TerminalReason(reason), Source: TerminalSource(source)}
	if m["winner"] != nil {
		p, err := player(m["winner"], "terminal.winner")
		if err != nil {
			return nil, err
		}
		terminal.Winner = &p
	}
	return terminal, nil
}

func decodeContext(value any) (*RepetitionContext, error) {
	if value == nil {
		return nil, nil
	}
	m, err := object(value, "contextVersion policyId status ledgerRef coveredRange positions", "", "repetition")
	if err != nil {
		return nil, err
	}
	version, err := unsigned(m["contextVersion"], 1, InvalidStateShape, "repetition.contextVersion")
	if err != nil || version != 1 {
		return nil, failure(InvalidStateShape, "repetition.contextVersion")
	}
	policy, p := m["policyId"].(string)
	status, s := m["status"].(string)
	ledger, l := m["ledgerRef"].(string)
	if !p || !s || !l {
		return nil, failure(InvalidStateShape, "repetition")
	}
	covered, err := object(m["coveredRange"], "from to", "", "repetition.coveredRange")
	if err != nil {
		return nil, err
	}
	from, err := decimal(covered["from"], "repetition.coveredRange.from")
	if err != nil {
		return nil, err
	}
	to, err := decimal(covered["to"], "repetition.coveredRange.to")
	if err != nil {
		return nil, err
	}
	items, err := array(m["positions"], -1, "repetition.positions")
	if err != nil {
		return nil, err
	}
	context := &RepetitionContext{ContextVersion: 1, PolicyID: policy, Status: status, LedgerRef: ledger, CoveredRange: CoveredRange{from, to}, Positions: make([]PositionOccurrence, 0, len(items))}
	for _, item := range items {
		entry, err := object(item, "position count occurrenceRefs", "", "repetition.positions")
		if err != nil {
			return nil, err
		}
		pos, err := object(entry["position"], positionFields, "", "repetition.position")
		if err != nil {
			return nil, err
		}
		position, err := decodePosition(pos, "repetition.position")
		if err != nil {
			return nil, err
		}
		count, err := unsigned(entry["count"], ^uint64(0), InvalidStateShape, "repetition.count")
		if err != nil {
			return nil, err
		}
		refs, err := array(entry["occurrenceRefs"], -1, "repetition.occurrenceRefs")
		if err != nil {
			return nil, err
		}
		references := make([]string, 0, len(refs))
		for _, ref := range refs {
			r, ok := ref.(string)
			if !ok {
				return nil, failure(InvalidStateShape, "repetition.occurrenceRefs")
			}
			references = append(references, r)
		}
		context.Positions = append(context.Positions, PositionOccurrence{position, count, references})
	}
	return context, nil
}

// DecodeState checks wire shapes and exact integers before constructing fixed
// arrays or narrowing numbers. No caller bytes or shared state are modified.
func DecodeState(data []byte) (State, error) {
	var state State
	if !validUnicode(data) {
		return state, failure(InvalidStateShape, "state")
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	value, err := readValue(dec, 0)
	if err != nil {
		return state, err
	}
	if _, err := dec.Token(); err != io.EOF {
		return state, failure(InvalidStateShape, "state")
	}
	m, err := object(value, positionFields+" stateSchemaVersion revision terminal", "repetition", "state")
	if err != nil {
		return state, err
	}
	version, err := unsigned(m["stateSchemaVersion"], 1, UnsupportedStateVersion, "stateSchemaVersion")
	if err != nil || version != 1 {
		return state, failure(UnsupportedStateVersion, "stateSchemaVersion")
	}
	state.StateSchemaVersion = 1
	state.Position, err = decodePosition(m, "state")
	if err != nil {
		return State{}, err
	}
	state.Revision, err = decimal(m["revision"], "revision")
	if err != nil {
		return State{}, err
	}
	state.Terminal, err = decodeTerminal(m["terminal"])
	if err != nil {
		return State{}, err
	}
	state.Repetition, err = decodeContext(m["repetition"])
	if err != nil {
		return State{}, err
	}
	if err = ValidateStructure(state); err != nil {
		return State{}, err
	}
	return state, nil
}
