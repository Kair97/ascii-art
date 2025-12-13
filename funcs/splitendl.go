package funcs

import "strings"

// func SplitNewLine(s string) []string {
// 	var res []string
// 	fmt.Println(s)
// 	temp := ""
// 	count := 0
// 	for i := 0; i < len(s); i++ {
// 		val := string(s[i])
// 		temp += val
// 		if i != len(s)-1 && val == "\\" && s[i+1] == 'n' {
// 			res = append(res, temp[:len(temp)-1])
// 			fmt.Printf("--> %v", temp)
// 			i++
// 			temp = ""
// 		}

// 		if val == " " {
// 			res = append(res, temp[:len(temp)-1])
// 			temp = ""

// 		}
// 		count++
// 		fmt.Printf("count is here --> %v\n", temp)
// 		fmt.Printf("res is here --> %v\n", res)

// 	}
// 	res = append(res, "i")
// 	fmt.Println(res)
// 	return res

// }

func SplitNewLine(s string) []string {

	return strings.Split(s, " ")

}
