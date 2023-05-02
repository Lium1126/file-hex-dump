package internal

import (
	"bufio"
	"os"
	"sync"
)

/*
   Compute computes the SHA256 checksum HEX dump in parallel for each line of the file
   pointed to by the file pointer 'fr', and outputs the processing results in the original order
   of the lines.

   Parameters:
     - fr: File pointer to the input file.

   Procedure:
     - For each line in the input file, the following processing is executed in parallel:
       - Computes the SHA256 checksum from the line data and obtains the HEX dump.
     - The processing results for each line are outputted in the original order of the lines.

   Notes:
     - This function utilizes parallel processing, with multiple threads executing concurrently.
     - Assumes the input file has UNIX-style LF (\n) line endings.

   Returns:
     - None
*/
func Compute(fr *os.File) {
	var wg sync.WaitGroup
	var ctxs []*context

	scanner := bufio.NewScanner(fr)
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
