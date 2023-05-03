// Package internal contains HEX dump processing for each line of the file.
package internal

import (
	"sync"
)

type context struct {
	sync.Mutex
	text string
	hash string
}

// newContext creates a context data with an internal value of 's'.
func newContext(s string) *context {
	return &context{
		Mutex: sync.Mutex{},
		text:  s,
		hash:  "",
	}
}

// process calculates the SHA256 checksum for the text stored in 'ctx'
// and stores the HEX dump of the checksum in 'ctx'.
func (ctx *context) process() {
	b := getSHA256Binary(ctx.text)
	ctx.hash = encodeHex(b)
}
