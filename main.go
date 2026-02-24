package main

import (
	"ascii-art/funcs"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	input, banner := funcs.GetInpAndBaner(args)
	fmt.Println(input, banner)
	data, err := os.ReadFile("banners/" + banner)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	dataStr := funcs.GetData(data)
	lines := funcs.FixLines(input)

	fmt.Println(funcs.PrintData(lines, dataStr))

}
