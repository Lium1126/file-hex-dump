package internal

import (
	"os"
	"sync"
)

/*
   routineProcess executes processing for the provided context.

   Parameters:
   - ctx: The context on which the processing is performed.

   Description:
   - This function executes the processing for the provided context.
     The specific actions or operations performed on the context depend on the implementation details of the function.
     It is assumed that the context contains relevant information or state necessary for the processing.

   Example usage:
   ctx := newContext("example")
   routineProcess(ctx)
   // Perform further operations based on the processed context.
*/
func routineProcess(ctx *context) {
	ctx.process()
	ctx.Unlock()
}

/*
   routineWrite continuously outputs the hash values stored in the provided contexts.
   The output process is a critical section, so it is passed to WriteC in a pre-locked state.
   Each context is processed in parallel, and after processing, the lock is released.
   This ensures that the hash values are output in a consistent order regardless of the processing time.

   Parameters:
   - WriteC: The channel from which the contexts with hash values are received.
   - f: The file to which the hash values are written.
   - wg: The WaitGroup used to synchronize the completion of goroutines.

   Description:
   - This function continuously receives contexts from the WriteC channel and outputs their hash values.
     The output is performed as a critical section, meaning that the file access is synchronized to maintain order.
     Each context is processed in parallel to maximize efficiency.
     After processing each context, the lock is released, allowing the next context to be processed.
     This ensures that the output order of the hash values remains consistent, regardless of the processing time for each context.

   Example usage:
   writeChannel := make(chan *context)
   file, _ := os.Create("output.txt")
   var wg sync.WaitGroup

   // Start goroutine for writing hash values
   go routineWrite(writeChannel, file, &wg)

   // Send contexts with hash values to the writeChannel
   for _, ctx := range contexts {
       writeChannel <- ctx
   }

   close(writeChannel)
   wg.Wait()
   // Output file now contains the hash values in a consistent order.
*/
func routineWrite(WriteC <-chan *context, f *os.File, wg *sync.WaitGroup) {
	for ctx := range WriteC {
		ctx.Lock()
		f.Write([]byte(ctx.hash + "\n"))
		ctx.Unlock()
		wg.Done()
	}
}
