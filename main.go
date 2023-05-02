package main

import (
	"bufio"
	"os"
	"sync"

	"github.com/Lium1126/hexdump/internal"
	"github.com/Lium1126/hexdump/internal/pkg/logger"
)

func fclose(f *os.File) {
	if err := f.Close(); err != nil {
		logger.LogErr("failed to close the file.", "error", err)
		return
	}
	logger.LogDebug("file close successfully.")
}

func main() {
	logger.InitZap()

	// fr_name is name of input file.
	fr_name := "example.txt"
	// fw_name is name of output file.
	fw_name := "output.txt"

	// open the input file
	fr, err := os.Open(fr_name)
	if err != nil {
		logger.LogErr("cannot open the file.", "error", err)
		return
	}
	defer fclose(fr)
	logger.LogDebug("input file open successfully.", "filename", fr_name)

	// open the output file
	fw, err := os.Create(fw_name)
	if err != nil {
		logger.LogErr("cannot create the file.", "error", err)
		return
	}
	defer fclose(fw)
	logger.LogDebug("output file open successfully.", "filename", fw_name)

	// processing
	var wg sync.WaitGroup
	var ctxs []*internal.Context

	scanner := bufio.NewScanner(fr)
	for scanner.Scan() {
		ctxs = append(ctxs, internal.NewContext(scanner.Text()))
	}

	wg.Add(len(ctxs))
	WriteC := make(chan *internal.Context, len(ctxs))
	go internal.RoutineWrite(WriteC, fw, &wg)

	for _, ctx := range ctxs {
		ctx.Lock()
		WriteC <- ctx
	}

	for _, ctx := range ctxs {
		go internal.RoutineProcess(ctx)
	}

	wg.Wait()
}
