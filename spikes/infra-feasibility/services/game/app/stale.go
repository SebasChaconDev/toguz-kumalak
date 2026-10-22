package app

import "errors"

var ErrStale = errors.New("stale revision")

// AcceptSyntheticResult is an experimental revision guard, not a game command.
func AcceptSyntheticResult(current, base int64) error {
	if current < 0 || base != current {
		return ErrStale
	}
	return nil
}
