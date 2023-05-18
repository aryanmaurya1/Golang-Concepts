package main

import "fmt"

func sumInts(arr []int) int {
	var sum = 0
	for _, value := range arr {
		sum = sum + value
	}
	return sum
}

func sumFloats(arr []float32) float32 {
	var sum float32 = 0.0
	for _, value := range arr {
		sum = sum + value
	}
	return sum
}

// A function which can accept a slice of 'int'
// 'float32' and 'float64'
func sumIntsAndFloat[T int | float32 | float64](arr []T) T {
	var sum T
	for _, value := range arr {
		sum = sum + value
	}
	return sum
}

func main() {
	var ints = []int{1, 2, 3, 4, 5, 6}
	var floats = []float32{1.1, 2.3, 4.5, 7, 8, 9}

	fmt.Println(sumInts(ints))
	fmt.Println(sumFloats(floats))
	fmt.Println(sumIntsAndFloat(floats))
	fmt.Println(sumIntsAndFloat(ints))

	container := GenericContainer[int]{
		values: []int{1, 2, 3},
		keys:   []int{2, 3, 4},
	}

	container.DisplayGenericContainer()
	containerf := GenericContainer[float32]{
		values: []float32{1,2,3,4},
		keys: []float32{0,9,8,7},
	}
	containerf.DisplayGenericContainer()
}

// OUTPUT
// 21
// 31.9
// 31.9
// 21
// Keys :  [2 3 4]
// Values :  [1 2 3]
// Keys :  [0 9 8 7]
// Values :  [1 2 3 4]