// Package internal contains HEX dump processing for each line of the file.
package internal

import (
	"crypto/sha256"
)

// getSHA256Binary calculates the SHA256 checksum for the given string 's'.
func getSHA256Binary(s string) []byte {
	r := sha256.Sum256([]byte(s))

	return r[:]
}
