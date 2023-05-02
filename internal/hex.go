// Package internal contains HEX dump processing for each line of the file.
package internal

import "encoding/hex"

// encodeHex converts the given byte array into a hexadecimal string.
func encodeHex(b []byte) string {
	return hex.EncodeToString(b)
}
