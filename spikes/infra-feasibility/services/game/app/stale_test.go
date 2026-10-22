package app

import (
	"errors"
	"testing"
)

// Catches accepting a synthetic AI response computed against an older state.
func TestStaleAI(t *testing.T) {
	if !errors.Is(AcceptSyntheticResult(2, 1), ErrStale) {
		t.Fatal("stale result accepted")
	}
	if AcceptSyntheticResult(2, 2) != nil {
		t.Fatal("current result rejected")
	}
}
