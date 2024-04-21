package main

import "fmt"

// Arrays are of :
// 		1. Fixed length
// 		2. Same Type
// 		3. Indexable
// 		4. Contiguous in Memory

func main() {
	var intArr [3] int32 // 12 bytes of memory
	intArr[0] = 133
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	// Print the memory location of the elements
	fmt.Println(&intArr[0])

	// Initialize the array using this symbol
	initialArr  := [3]int32{1,2,3}
	fmt.Println(initialArr)

	coolArr  := [...]int32{1,2,3,4}
	fmt.Println(coolArr)

	// we can add values to an array using the slice
	var intSlice[]int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice)) // 3, 3
	intSlice = append(intSlice, 7) // arr -> [4,5,6,7, *, *]
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice)) // 4, 6
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32 {8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Create a custom array size using make
	var intSlice3 []int32 = make([]int32, 3, 8) // specify length, capacity of the slice
	fmt.Println(intSlice3)

	// Map
	var myMap map[string] uint8 = make(map[string]uint8) // key-> String; value -> uint8
	fmt.Println(myMap)

	// Be careful because map will always return something even if the key doesn't exist
	myMap2 := map[string] uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"]) // 23
	fmt.Println(myMap2["Jason"]) // 0
	// To check if a key exists 
	person := "Jason"
	age, exists := myMap2[person]
	if exists {
		fmt.Printf("%v age is %v\n", person, age)
	} else {
		fmt.Printf("%v doesn't exist, Sorry!\n", person)
	}

	// We can also delete entries in map
	delete(myMap2, "Sarah")
	fmt.Println(myMap2)

	// To iterate through an Array we can use the following syntax:
	for name, age := range myMap2{
		fmt.Printf("Name: %v, age: %v\n", name, age)
	}

	// Iterating through an Array
	for i,v := range intArr{
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// Loop iteration manually
	for i:=0; i < 10; i++{
		fmt.Println(i)
	}
	// Loop iteration using conditional
	j :=0
	for j<10{
		fmt.Println(j)
		j++ // or j = j + 1
	}

}