package app

import (
	"context"
	"errors"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/engine"
	"time"
)

var ErrBudget = errors.New("work budget exceeded")

func Compute(ctx context.Context, work int) (string, error) {
	if e := ctx.Err(); e != nil {
		return "", e
	}
	if work < 1 || work > 200 {
		return "", ErrBudget
	}
	value := "synthetic"
	for i := 0; i < work; i++ {
		timer := time.NewTimer(time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			return "", ctx.Err()
		case <-timer.C:
			value = engine.Digest(value)
		}
	}
	return value, nil
}
