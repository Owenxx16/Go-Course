// Most of the time we don't need to think about the underlying array of a slice. We can create a new slice using the make function:

// // func make([]T, len, cap) []T
// mySlice := make([]int, 5, 10)

// // the capacity argument is usually omitted and defaults to the length
// mySlice := make([]int, 5)

//Make Exercise

package main

func getMessageCosts(messages []string) []float64 {
	// ?
	var costs = make([]float64, len(messages))
	for i, message := range messages {
		costs[i] = float64(len(message)) * 0.01
	}
	return costs
}

//Demo Append Function nhma x2 chỗ


// func Append(slice, data []byte) []byte {
// 	l := len(slice)
// 	if l + len(data) > cap(slice) {  // reallocate
// 			// Allocate double what's needed, for future growth.
// 			newSlice := make([]byte, (l+len(data))*2)
// 			// The copy function is predeclared and works for any slice type.
// 			copy(newSlice, slice)
// 			slice = newSlice
// 	}
// 	slice = slice[0:l+len(data)]
// 	copy(slice[l:], data)
// 	return slice
// }
