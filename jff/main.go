package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	input := os.Args[1:]

	banner := "standard.txt"
	os.WriteFile("out.txt", []byte{}, 0644)
	data, err := os.ReadFile(banner)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	dataStr := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	fixed := strings.ReplaceAll(input[0], "\\n", "\n")
	lines := strings.Split(fixed, "\n")
	// lines := funcs.SplitNewLine(input[0])
	f, err := os.OpenFile("out.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()
	// fmt.Println(lines)
	// fmt.Println("--------------------------------------")
	// for _, val := range lines {
	// 	fmt.Printf("%v <---\n", val)
	// }
	for i, line := range lines {

		if line == "" {
			if i == len(lines)-1 {
				continue
			}
			fmt.Println()
			continue
		}
		for i := 1; i < 9; i++ {

			for _, char := range line {
				k := int(char)

				startLine := (k - 32) * 9
				f.WriteString(dataStr[startLine+i])
				fmt.Print(dataStr[startLine+i])

			}
			f.WriteString("\n")
			fmt.Println()
		}
	}

}
