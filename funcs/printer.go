package funcs

import "fmt"

func PrintData(lines []string, dataStr []string) {
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
				fmt.Print(dataStr[startLine+i])

			}
			fmt.Println()
		}
	}
}
