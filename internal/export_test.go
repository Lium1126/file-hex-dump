package internal

type Context = context

func NewHashSettedContext(t string, h string) *Context {
	return &Context{
		text: t,
		hash: h,
	}
}

var (
	NewContext                     = newContext
	Process         func(*context) = (*context).process
	EncodeHex                      = encodeHex
	GetSHA256Binary                = getSHA256Binary
)
