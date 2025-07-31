// test 1
package main

// func concat(strs ...string) string {
// 	final := ""
// 	// strs is just a slice of strings
// 	for i := 0; i < len(strs); i++ {
// 		final += strs[i]
// 	}
// 	return final
// }

// func main() {
// 	final := concat("Hello ", "there ", "friend!")
// 	fmt.Println(final)
// 	// Output: Hello there friend!
// }

// func printStrings(strings ...string) {
// 	for i := 0; i < len(strings); i++ {
// 		fmt.Println(strings[i])
// 	}
// }

// func main() {
// 	names := []string{"bob", "sue", "alice"}
// 	printStrings(names...)
// }

//Assignment: Variadic Functions
func sum(nums ...int) int {
	// ?
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}
