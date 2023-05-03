// Package internal contains HEX dump processing for each line of the file.
package internal

import (
	"bufio"
	"io"
	"sync"
)

// Compute computes the SHA256 checksum HEX dump in parallel for each line of the file
// pointed to by the file pointer 'f', and outputs the processing results in the original order
// of the lines.
func Compute(file *io.Reader) {
	var (
		waitGroup sync.WaitGroup
		ctxs      []*context
	)

	scanner := bufio.NewScanner(*file)
	for scanner.Scan() {
		ctxs = append(ctxs, newContext(scanner.Text()))
	}

	waitGroup.Add(len(ctxs))
	writeC := make(chan *context, len(ctxs))

	go routineWrite(writeC, &waitGroup)

	for _, ctx := range ctxs {
		ctx.Lock()
		writeC <- ctx
	}

	for _, ctx := range ctxs {
		go routineProcess(ctx)
	}

	waitGroup.Wait()
}
