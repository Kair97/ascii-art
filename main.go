package main

import (
	"ascii-art/funcs"
	"fmt"
	"os"
)

func main() {

	input := os.Args[1:]

	banner := "standard.txt"

	if (len(input) > 1) && (input[1] == "standard.txt" || input[1] == "thinkertoy.txt" || input[1] == "shadow.txt") {
		banner = input[1]
	}

	data, err := os.ReadFile(banner)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	dataStr := funcs.GetData(data)
	lines := funcs.FixLines(input)

	funcs.PrintData(lines, dataStr)

}
