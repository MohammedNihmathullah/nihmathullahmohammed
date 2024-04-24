package main

import "fmt"

func Addition(a, b int) int {
	return a + b
}

func main() {
	num1 := 10
	num2 := 20

	sum := Addition(num1, num2)

	fmt.Println(sum)
}
