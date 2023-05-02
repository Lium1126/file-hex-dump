package internal

import "encoding/hex"

func encodeHex(b []byte) string {
	return hex.EncodeToString(b)
}
