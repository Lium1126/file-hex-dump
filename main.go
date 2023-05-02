package main

import (
	"os"

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

	// open the input file
	fr, err := os.Open(fr_name)
	if err != nil {
		logger.LogErr("cannot open the file.", "error", err)
		return
	}
	defer fclose(fr)
	logger.LogDebug("input file open successfully.", "filename", fr_name)

	internal.Compute(fr)
}
