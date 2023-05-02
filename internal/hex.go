package internal

import "encoding/hex"

/*
   encodeHex converts the given byte array into a hexadecimal string.

   Parameters:
   - b: The byte array to be converted.

   Returns:
   - The resulting hexadecimal string.

   Description:
   - This function takes a byte array as input and converts it into a hexadecimal string.
     Each byte in the array is converted into a two-character hexadecimal representation.
     The resulting string represents the hexadecimal values of the bytes in the same order as the input.

   Example usage:
   data := []byte{10, 20, 30}
   hexString := encodeHex(data)
   // hexString is now "0a141e"
*/
func encodeHex(b []byte) string {
	return hex.EncodeToString(b)
}
