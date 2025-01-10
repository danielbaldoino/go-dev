package main

import (
	"fmt"
)

func isFibonacci(num int) bool {
	a, b := 0, 1

	if num == a || num == b {
		return true
	}

	for b <= num {
		a, b = b, a+b
		if b == num {
			return true
		}
	}

	return false
}

func main() {
	var num int
	fmt.Print("Informe um número: ")
	fmt.Scan(&num)

	if isFibonacci(num) {
		fmt.Printf("O número %d pertence à sequência de Fibonacci.\n", num)
	} else {
		fmt.Printf("O número %d NÃO pertence à sequência de Fibonacci.\n", num)
	}
}
