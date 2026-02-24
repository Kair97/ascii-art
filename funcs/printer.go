package funcs

import (
	"strings"
)

func PrintData(lines []string, dataStr []string) string {
	var builder strings.Builder

	for i, line := range lines {

		if line == "" {
			if i == len(lines)-1 {
				continue
			}
			builder.WriteString("\n")
			continue
		}
		for i := 1; i < 9; i++ {

			for _, char := range line {
				k := int(char)

				if k < 32 || k > 126 {
					continue
				}
				startLine := (k - 32) * 9
				builder.WriteString(dataStr[startLine+i])

			}
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
