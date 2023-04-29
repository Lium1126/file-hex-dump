package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// fr_name is name of input file
	fr_name := "example.txt"
	// fw_name is name of output file
	fw_name := "output.txt"

	fr, err := os.Open(fr_name)
	if err != nil {
		fmt.Printf("Cannot open the file: %s", err.Error())
		return
	}
	defer func() {
		err := fr.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fw, err := os.Create(fw_name)
	if err != nil {
		fmt.Printf("Cannot create the file: %s", err.Error())
		return
	}
	defer func() {
		err := fw.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	scanner := bufio.NewScanner(fr)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(time.Second * 1)
	}
}
