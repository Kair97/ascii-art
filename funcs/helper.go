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

func FixLines(input string) []string {
	fixed := strings.ReplaceAll(input, "\\n", "\n")
	lines := strings.Split(fixed, "\n")
	// lines := funcs.SplitNewLine(input[0])
	return lines
}

func GetInpAndBaner(args []string) (string, string) {

	if len(args) == 1 {
		return args[0], "standard.txt"
	} else if len(args) == 0 {
		return "", "standard.txt"
	} else {

		last := args[len(args)-1]
		if last == "standard.txt" || last == "thinkertoy.txt" || last == "shadow.txt" {
			return strings.Join(args[:len(args)-1], ""), last
		} else {
			return strings.Join(args, ""), "standard.txt"
		}

	}

}
