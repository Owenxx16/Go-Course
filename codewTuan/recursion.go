package main

import "fmt"

// func countDown(number int) {

// 	fmt.Println(number)
// 	if number == 0 {
// 		fmt.Println("Happy New Year!")
// 	}

// 	if number > 0 {
// 		countDown(number - 1)

// 	}
// }

// func sumRecursion(n int) int {
// 	if n == 0 {
// 		fmt.Println("Wrong")
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	return n + sumRecursion(n-1)
// }

func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	// countDown(10)
	//sumRecursion(5)
	fmt.Println("Fibonacci of 5 is:", Fibonacci(5))
}
