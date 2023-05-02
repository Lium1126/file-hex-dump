package internal

import (
	"crypto/sha256"
)

func getSHA256Binary(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}
