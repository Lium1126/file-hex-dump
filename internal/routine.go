package internal

import (
	"os"
	"sync"
)

func RoutineProcess(ProcessC <-chan *Context) {
	for ctx := range ProcessC {
		ctx.process()
		ctx.Unlock()
	}
}

func RoutineWrite(WriteC <-chan *Context, f *os.File, wg *sync.WaitGroup) {
	for ctx := range WriteC {
		ctx.Lock()
		f.Write([]byte(ctx.hash + "\n"))
		ctx.Unlock()
		wg.Done()
	}
}
