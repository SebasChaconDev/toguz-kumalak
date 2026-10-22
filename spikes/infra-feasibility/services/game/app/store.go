package app

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/engine"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

//go:embed migrations/*.sql
var Migrations embed.FS

var ErrForbidden = errors.New("forbidden")
var ErrConflict = errors.New("intent conflict")

type Intent struct {
	OperationID      string `json:"operationId"`
	ExpectedRevision int64  `json:"expectedRevision"`
	Payload          string `json:"payload"`
}
type State struct {
	Revision int64 `json:"revision"`
	Effects  int64 `json:"effects"`
}
type Notice struct {
	GameID   string
	Revision int64
}
type Store struct {
	Pool         *pgxpool.Pool
	BeforeCommit func() error
	Notify       func(context.Context, Notice) error
}

func (s *Store) Get(ctx context.Context, id, actor string) (State, error) {
	var st State
	var owner string
	e := s.Pool.QueryRow(ctx, "SELECT owner,revision,effects FROM synthetic_games WHERE id=$1", id).Scan(&owner, &st.Revision, &st.Effects)
	if e != nil {
		return st, e
	}
	if owner != actor {
		return State{}, ErrForbidden
	}
	return st, nil
}

func (s *Store) Apply(ctx context.Context, id, actor string, in Intent) (State, error) {
	tx, e := s.Pool.BeginTx(ctx, pgx.TxOptions{})
	if e != nil {
		return State{}, e
	}
	defer tx.Rollback(context.Background())
	var st State
	var owner string
	e = tx.QueryRow(ctx, "SELECT owner,revision,effects FROM synthetic_games WHERE id=$1 FOR UPDATE", id).Scan(&owner, &st.Revision, &st.Effects)
	if e != nil {
		return st, e
	}
	if owner != actor {
		return State{}, ErrForbidden
	}
	raw, _ := json.Marshal(in)
	hash := engine.Digest(string(raw))
	var previousHash string
	var previous State
	e = tx.QueryRow(ctx, "SELECT intent_hash,revision,effects FROM synthetic_receipts WHERE game_id=$1 AND actor=$2 AND operation_id=$3", id, actor, in.OperationID).Scan(&previousHash, &previous.Revision, &previous.Effects)
	if e == nil {
		if hash != previousHash {
			return State{}, ErrConflict
		}
		return previous, nil
	}
	if !errors.Is(e, pgx.ErrNoRows) {
		return State{}, e
	}
	if in.ExpectedRevision != st.Revision {
		return State{}, ErrStale
	}
	if st.Revision == 9223372036854775807 {
		return State{}, ErrConflict
	}
	st.Revision++
	st.Effects++
	_, e = tx.Exec(ctx, "UPDATE synthetic_games SET revision=$2,effects=$3 WHERE id=$1", id, st.Revision, st.Effects)
	if e != nil {
		return State{}, e
	}
	_, e = tx.Exec(ctx, "INSERT INTO synthetic_receipts(game_id,actor,operation_id,intent_hash,revision,effects) VALUES($1,$2,$3,$4,$5,$6)", id, actor, in.OperationID, hash, st.Revision, st.Effects)
	if e != nil {
		return State{}, e
	}
	if s.BeforeCommit != nil {
		if e = s.BeforeCommit(); e != nil {
			return State{}, e
		}
	}
	if e = tx.Commit(ctx); e != nil {
		return State{}, e
	}
	if s.Notify != nil {
		noticeCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
		defer cancel()
		_ = s.Notify(noticeCtx, Notice{id, st.Revision})
	}
	return st, nil
}
