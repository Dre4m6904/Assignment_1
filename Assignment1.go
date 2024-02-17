package main

import "fmt"

func main() {
	array := []int{5, 3, 9, 2, 8, 1, 4, 7}

	newFunction := appendFunc(PrintSlice, Sortarray, IncrementOdd, ReverseSlice)

	newFunction(array)
}

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(array []int) {
		dst(array)

		for _, fn := range src {
			fn(array)
			dst(array)
		}
	}
}

func Sortarray(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func IncrementOdd(array []int) {
	for i := 0; i < len(array); i++ {
		if i%2 != 0 {
			array[i]++
		}
	}
}

func PrintSlice(array []int) {
	fmt.Println("Slice:", array)
}

func ReverseSlice(array []int) {
	count := len(array) - 1
	for i := 0; i < count; i++ {
		array[i], array[count] = array[count], array[i]
		count--
	}
}
