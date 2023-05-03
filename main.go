// Package main contains the main work of file HEX dump.
package main

import (
	"io"
	"log"
	"os"

	"github.com/Lium1126/hexdump/internal"
	"github.com/Lium1126/hexdump/internal/pkg/logger"
	"github.com/alecthomas/kingpin/v2"
)

func main() {
	// Application setting
	var (
		cmd   = kingpin.CommandLine
		debug = cmd.Flag("debug", "Enable debug mode.").Short('d').Bool()
		fname = cmd.Arg("src-file", "Source file").String()
	)

	cmd.Name, cmd.Help = "myhexdump", "Print HEX Dump of SHA256 from file content."
	cmd.Version("0.0.1")

	// CLI parse
	if _, err := cmd.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	// Logger Initialization
	if err := logger.InitZap(*debug); err != nil {
		log.Fatal(err)
	}

	// File open
	filePointer, err := os.Open(*fname)
	if err != nil {
		logger.LogErr("fail to open the file.", "error", err, "filename", *fname)

		return
	}

	defer func(*os.File) {
		if err := filePointer.Close(); err != nil {
			logger.LogErr("failed to close the file.", "error", err)
		} else {
			logger.LogDebug("file close successfully.")
		}
	}(filePointer)
	logger.LogDebug("file open successfully.")

	reader := io.Reader(filePointer)
	internal.Compute(&reader)
}
