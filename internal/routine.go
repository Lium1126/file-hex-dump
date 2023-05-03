// Package internal contains HEX dump processing for each line of the file.
package internal

import (
	"fmt"
	"os"
	"sync"
)

// routineProcess executes processing for the provided context.
func routineProcess(ctx *context) {
	ctx.process()
	ctx.Unlock()
}

// routineWrite continuously outputs the hash values stored in the provided contexts.
// The output process is a critical section, so it is passed to WriteC in a pre-locked state.
// Each context is processed in parallel, and after processing, the lock is released.
// This ensures that the hash values are output in a consistent order regardless of the processing time.
func routineWrite(writeC <-chan *context, wg *sync.WaitGroup) {
	for ctx := range writeC {
		ctx.Lock()
		fmt.Fprintln(os.Stdout, ctx.hash)
		ctx.Unlock()
		wg.Done()
	}
}
