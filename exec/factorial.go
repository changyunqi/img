package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
	fmt.Println("Error Input.")
		os.Exit(0)
	}
	n, _ := strconv.ParseInt(os.Args[1], 10, 64)
	factorial := Computer(n)
	fmt.Println(factorial)
}

func Computer(n int64) (factorial int64) {

	if n == 1 {
		return 1
	} else if n > 1 {
		return n * Computer(n-1)
	} else {
		fmt.Println("不合理的输入！")
		return factorial
	}
	return factorial
}
