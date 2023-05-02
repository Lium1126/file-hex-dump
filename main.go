package main

import (
	"fmt"
	"os"

	"github.com/Lium1126/hexdump/internal"
	"github.com/Lium1126/hexdump/internal/pkg/logger"
	"github.com/alecthomas/kingpin/v2"
)

var (
	Cmd = kingpin.CommandLine

	debug = Cmd.Flag("debug", "Enable debug mode.").Short('d').Bool()
	fname = Cmd.Arg("src-file", "Source file").String()
)

func init() {
	Cmd.Name = "myhexdump"
	Cmd.Help = "Print HEX Dump of SHA256 from file content."
	Cmd.Version("0.0.1")
}

func main() {
	// CLI parse
	_, err := Cmd.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("failed to parce of command line: %v\n", err.Error())
		os.Exit(1)
	}

	// Logger Initialization
	if err := logger.InitZap(*debug); err != nil {
		fmt.Printf("failed to initiarize logger: %v\n", err.Error())
		os.Exit(1)
	}

	// File open
	f, err := os.Open(*fname)
	if err != nil {
		logger.LogErr("fail to open the file.", "error", err, "filename", *fname)
		return
	}
	defer func(*os.File) {
		if err := f.Close(); err != nil {
			logger.LogErr("failed to close the file.", "error", err)
		} else {
			logger.LogDebug("file close successfully.")
		}
	}(f)
	logger.LogDebug("file open successfully.")

	internal.Compute(f)
}
