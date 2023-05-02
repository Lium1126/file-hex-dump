package internal

import (
	"crypto/sha256"
)

/*
   getSHA256Binary calculates the SHA256 checksum for the given string s.

   Parameters:
   - s: The string for which the SHA256 checksum is calculated.

   Returns:
   - The resulting SHA256 checksum as a binary byte array.

   Description:
   - This function calculates the SHA256 checksum for the provided string s.
     The SHA256 algorithm produces a 256-bit hash value, which is returned as a binary byte array.
     Each byte in the array represents 8 bits of the hash value.

   Example usage:
   input := "example"
   checksum := getSHA256Binary(input)
   // Use the resulting checksum for further processing or comparison.
*/
func getSHA256Binary(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}
