package main

import "fmt"

func main(){
	// Generics are a way to allow a function to receive multiple types of input 
	var intSlices = []int{1,2,3}
	var float32Slices = []float32{1,2,3}


	fmt.Println("sum of IntSlices -> ", sumSlices(intSlices))
	fmt.Println("sum of Float32Slices -> ", sumSlices(float32Slices))
	fmt.Println("length of Float32Slices -> ", findLength(float32Slices))
}

func sumSlices[T int | float32 | float64](slice [] T) T {
	var sum T
	for _, v := range(slice){
		sum += v
	}
	return sum
}

// We can use any type

func findLength[T any](slice [] T) int {
	return len(slice)
}