package internal

import (
	"os"
	"sync"
)

func RoutineProcess(ctx *context) {
	ctx.process()
	ctx.Unlock()
}

func RoutineWrite(WriteC <-chan *context, f *os.File, wg *sync.WaitGroup) {
	for ctx := range WriteC {
		ctx.Lock()
		f.Write([]byte(ctx.hash + "\n"))
		ctx.Unlock()
		wg.Done()
	}
}
