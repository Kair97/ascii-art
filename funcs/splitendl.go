package funcs

func SplitNewLine(s string) []string {
	var res []string
	temp := ""
	count := 0
	for i := 0; i < len(s); i++ {
		val := string(s[i])
		temp += val
		if i != len(s)-1 && val == "\\" && s[i+1] == 'n' {
			res = append(res, temp[:len(temp)-1])
			i++
			temp = ""
		}

		count++

	}

	if len(temp) != 0 {
		res = append(res, temp)
	}

	return res
}
