package main

import (
	"fmt"
)

func exercise2() {
	exe := 0
	for i := 0; i <= 100; i++ {
		if i%2 == 1 {
			if exe > 0 {
				fmt.Print(",")
			}
			fmt.Printf("%d", i)
			exe++
			if exe == 3 {
				fmt.Println("\n")
				exe = 0
			}
		}
	}
}

// func main() {
// 	exercise2()
// }
