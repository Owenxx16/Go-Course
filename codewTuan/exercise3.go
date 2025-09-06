package main

import (
	"fmt"
)

func exercise3(a int, b int) {
	var sum int
	for i := a; i <= b; i++ {
		for j := 1; j <= 10; j++ {
			sum = i * j
			fmt.Printf("%d x %d = %d\n", i, j, sum)
		}
	}
}

// func main() {
// 	var a1 int
// 	var a2 int
// 	fmt.Println("Please enter number a1:")
// 	fmt.Scanln(&a1)
// 	fmt.Println("Please enter number a2:")
// 	fmt.Scanln(&a2)
// 	if a1 == 0 && a2 == 0 {
// 		fmt.Println("please enter a valid number")
// 	}
// 	if a1 < 0 || a2 < 0 {
// 		fmt.Println("please enter a positive number")
// 	}
// 	if a1 > a2 {
// 		fmt.Println("please enter a1 less than a2")
// 	}
// 	exercise3(a1, a2)
// 	fmt.Println("Multiplication table from", a1, "to", a2)
// 	fmt.Println("====================================")
// }
