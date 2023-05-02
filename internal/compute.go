// Package internal contains HEX dump processing for each line of the file.
package internal

import (
	"bufio"
	"os"
	"sync"
)

// Compute computes the SHA256 checksum HEX dump in parallel for each line of the file
// pointed to by the file pointer 'f', and outputs the processing results in the original order
// of the lines.
func Compute(f *os.File) {
	var (
		wg   sync.WaitGroup
		ctxs []*context
	)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ctxs = append(ctxs, newContext(scanner.Text()))
	}

	wg.Add(len(ctxs))
	writeC := make(chan *context, len(ctxs))

	go routineWrite(writeC, &wg)

	for _, ctx := range ctxs {
		ctx.Lock()
		writeC <- ctx
	}

	for _, ctx := range ctxs {
		go routineProcess(ctx)
	}

	wg.Wait()
}
