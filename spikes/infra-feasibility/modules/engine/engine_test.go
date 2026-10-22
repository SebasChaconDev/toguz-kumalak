package engine

import "testing"

// Catches digest ignoring the payload; expected SHA-256 is an independent vector.
func TestDigest(t *testing.T) {
	if got := Digest("abc"); got != "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad" {
		t.Fatalf("digest = %q", got)
	}
	if Digest("abc") == Digest("abd") {
		t.Fatal("different inputs collide")
	}
}
