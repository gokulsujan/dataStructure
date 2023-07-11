package main

import "fmt"

func fib(num int) int {
	if num == 1 {
		return 0
	} else if num == 2 {
		return 1
	} else {
		return fib(num-1) + fib(num-2)
	}
}

func main() {
	fmt.Println(fib(5))
}
