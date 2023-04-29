package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Lium1126/hexdump/internal/pkg/logger"
)

func fclose(f *os.File) {
	if err := f.Close(); err != nil {
		logger.LogErr("Failed to close the file.", "err", err)
	}
}

func main() {
	logger.InitZap()

	// fr_name is name of input file
	fr_name := "example.txt"
	// fw_name is name of output file
	fw_name := "output.txt"

	// open the input file
	fr, err := os.Open(fr_name)
	if err != nil {
		logger.LogErr("Cannot open the file.", "err", err)
		return
	}
	defer fclose(fr)
	logger.LogDebug("Input file open successfully.", "filename", fr_name)

	// open the output file
	fw, err := os.Create(fw_name)
	if err != nil {
		logger.LogErr("Cannot create the file.", "err", err)
		return
	}
	defer fclose(fw)
	logger.LogDebug("Output file open successfully.", "filename", fw_name)

	scanner := bufio.NewScanner(fr)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(time.Second * 1)
	}
}
