package internal

import (
	"sync"
)

type Context struct {
	sync.Mutex
	text string
	hash string
}

func NewContext(s string) *Context {
	return &Context{
		text: s,
	}
}

func (ctx *Context) process() {
	b := getSHA256Binary(ctx.text)
	ctx.hash = encodeHex(b)
}
