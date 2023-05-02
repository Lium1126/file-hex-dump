package internal

import (
	"sync"
)

type context struct {
	sync.Mutex
	text string
	hash string
}

func newContext(s string) *context {
	return &context{
		text: s,
	}
}

func (ctx *context) process() {
	b := getSHA256Binary(ctx.text)
	ctx.hash = encodeHex(b)
}
