package main

import (
	"ascii-art/funcs"
	"fmt"
	"os"
	"time"
)

func main() {
	before := time.Now()
	args := os.Args[1:]

	input, banner := funcs.GetInpAndBaner(args)
	fmt.Println(input)
	data, err := os.ReadFile("banners/" + banner)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	dataStr := funcs.GetData(data)
	lines := funcs.FixLines(input)

	fmt.Println(funcs.PrintData(lines, dataStr))
	fmt.Println(time.Since(before))

}
