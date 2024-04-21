package main

import "fmt"

func main(){
	// declare a pointer
	var p *int32 = new (int32) // to give this variable a address to store the int32 var we declare it with new keyword
	var i int32
	fmt.Printf("The value of p pointer is %v\n", *p)
	fmt.Printf("The value of i var is %v\n", i)
	i = 10
	// when we use & we are referencing the memory address of the variable not its Address
	p = &i
	fmt.Printf("The value of p pointer afterwards is %v\n", *p)
	// We can now change the value at the memory address of i
	*p = 1
	fmt.Printf("The value of i var is %v\n", i)

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// Slices /////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////////////
	
	var slices[]int32 = []int32{1,2,3}
	var slicesCopy = slices
	slicesCopy[2] = 4
	fmt.Printf("The value of slices is %v\n", slices) // [1 2 4]
	fmt.Printf("The value of slicesCopy is %v\n", slicesCopy) // [1 2 4]

	// Arrays

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("The memory location of the thing1 array is %p\n", &thing1)
	square(thing1)
	fmt.Printf("The value of thing1 using square is: %v\n", thing1)

	fmt.Printf("The memory location of the thing1 array address is %p\n", &thing1)
	square_pointer(&thing1)
	fmt.Printf("The value of thing1 using square_pointer is: %v\n", thing1)

}
// This is inefficient as we are allocating two arrays to do one operation which is 
// squaring the numbers in thing1 array
func square(thing2 [5]float64) [5]float64{
	fmt.Printf("The value of thing2 looks like this: %p\n", &thing2)
	for i := range(thing2){
		thing2[i] = thing2[i]*thing2[i]
	}
	return thing2
}
// Instead we can use pointers to square the values in thing1 
func square_pointer(thing1 *[5]float64)  {
	fmt.Printf("The value of thing2 address looks like this: %p\n", thing1)
	for i := range(thing1){
		thing1[i] = thing1[i]*thing1[i]
	}
}