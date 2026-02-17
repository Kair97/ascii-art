package funcs

import "strings"

// func FileReader(s string) ([]byte, error) {

// 	data, err := os.ReadFile(s)
// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	return data, nil

// }

func GetData(data []byte) []string {

	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

}

func FixLines(input []string) []string {
	fixed := strings.ReplaceAll(input[0], "\\n", "\n")
	lines := strings.Split(fixed, "\n")
	// lines := funcs.SplitNewLine(input[0])
	return lines
}
