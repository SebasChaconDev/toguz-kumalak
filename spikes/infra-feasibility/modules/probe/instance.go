package probe

import (
	"crypto/rand"
	"encoding/hex"
)

var remoteInstance = func() string { b := make([]byte, 16); _, _ = rand.Read(b); return hex.EncodeToString(b) }()
