package internal

import (
	"bufio"
	"os"
	"sync"
)

func Compute(fr *os.File, fw *os.File) {
	var wg sync.WaitGroup
	var ctxs []*context

	scanner := bufio.NewScanner(fr)
	for scanner.Scan() {
		ctxs = append(ctxs, newContext(scanner.Text()))
	}

	wg.Add(len(ctxs))
	WriteC := make(chan *context, len(ctxs))
	go RoutineWrite(WriteC, fw, &wg)

	for _, ctx := range ctxs {
		ctx.Lock()
		WriteC <- ctx
	}

	for _, ctx := range ctxs {
		go RoutineProcess(ctx)
	}

	wg.Wait()
}
