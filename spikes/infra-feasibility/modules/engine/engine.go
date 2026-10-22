// Package engine is a throwaway deterministic digest, NOT a rules engine.
package engine

import (
	"crypto/sha256"
	"encoding/hex"
)

const Version = "spike-engine/v0.1.0"

func Digest(payload string) string {
	sum := sha256.Sum256([]byte(payload))
	return hex.EncodeToString(sum[:])
}
