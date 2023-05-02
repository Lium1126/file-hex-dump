package internal

import (
	"sync"
)

type context struct {
	sync.Mutex
	text string
	hash string
}

/*
   newContext creates a context data with an internal value of s.

   Parameters:
   - s: The value to be stored internally in the context.

   Returns:
   - A pointer to the newly created context data.

   Description:
   - This function creates a context data with an internal value of s.
     The context data can be used to store additional information or state related to the provided value.

   Example usage:
   ctx := newContext("example")
   // Use ctx for further processing or to access the internal value.
*/
func newContext(s string) *context {
	return &context{
		text: s,
	}
}

/*
   process calculates the SHA256 checksum for the text stored in ctx
   and stores the HEX dump of the checksum in ctx.

   Parameters:
   - ctx: The context containing the text for which the SHA256 checksum is calculated.

   Description:
   - This function calculates the SHA256 checksum for the text stored in the provided context, ctx.
     The SHA256 algorithm produces a 256-bit hash value, which is then converted into a HEX dump.
     The resulting HEX dump is stored back in the ctx.

   Example usage:
   ctx := newContext("example")
   ctx.process()
   // The HEX dump of the SHA256 checksum is now stored in ctx for further use.
*/
func (ctx *context) process() {
	b := getSHA256Binary(ctx.text)
	ctx.hash = encodeHex(b)
}
