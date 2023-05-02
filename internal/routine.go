package internal

import (
	"os"
	"sync"
)

func RoutineProcess(ctx *Context) {
	ctx.process()
	ctx.Unlock()
}

func RoutineWrite(WriteC <-chan *Context, f *os.File, wg *sync.WaitGroup) {
	for ctx := range WriteC {
		ctx.Lock()
		f.Write([]byte(ctx.hash + "\n"))
		ctx.Unlock()
		wg.Done()
	}
}
