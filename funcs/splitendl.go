package funcs

import "fmt"

func SplitNewLine(s string) []string {
	var res []string
	temp := ""

	for i := 0; i < len(s); i++ {
		val := string(s[i])
		temp += val
		if i != len(s)-1 && val == "\\" && s[i+1] == 'n' {
			res = append(res, temp[:len(temp)-1])
			fmt.Printf("--> %v", temp)
			res = append(res, "\\n")
			i++
			temp = ""

		}
		if val == " " {

			res = append(res, temp[:len(temp)-1])
			temp = ""

		}

	}
	res = append(res, "i")
	fmt.Println(res)
	return res

}
