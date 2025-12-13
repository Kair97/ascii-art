package main

import (
	"ascii-art/funcs"
	"fmt"
	"os"
	"strings"
)

func main() {

	input := os.Args[1]
	fmt.Println(input)

	banner := "standard.txt"

	data, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Something went wrong!", err)
		return
	}
	lines := strings.Split(string(data), "\n")

	inputLines := funcs.SplitNewLine(input)

	fmt.Println(inputLines)
	for _, line := range inputLines {

		for i := 0; i < 8; i++ {

			for _, char := range line {
				k := int(char)
				if char == ' ' {
					k = 32
				}

				startLine := (k - 32) * 9

				fmt.Print(lines[startLine+i])

			}
			fmt.Println()
		}
	}

}
