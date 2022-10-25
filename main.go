package main

import "fmt"

func fibonacci(x int) int {
	if x == 1 || x == 2 {
		return 1
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}
}

func salida() {
	x := 10
	i := 1
	for i <= x {
		fmt.Print(fibonacci(i), " ")
		i++
	}
}

func main() {
	salida()
}
