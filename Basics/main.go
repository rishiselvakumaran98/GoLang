package main

import "fmt"
import "unicode/utf8"

func main() {
	// when declaring a variable, you must specify the type
	var name string = "John"
	fmt.Println(name)
	// We can also specify the integer type
	var age int = 25
	fmt.Println(age)
	// the int has a size of 32 bits; where as int16 has a size of 16 bits
	var intNum int16 = 32767
	// if we do this the compiler will still run
	intNum += 1
	fmt.Println(intNum) // -32768

	// We can also use unsigned integers uints which can hold only positive numbers and have double the size of ints
	var uintNum uint16 = 65535
	// float64 can hold 64 bits of floating point numbers, the largest number it can hold is 1.797693134862315708145274237317043567981e+308
	var floatNum float64 = 3.14
	fmt.Println("This is float64 number: ", floatNum, "This is uint16 number: ", uintNum)

	// we can also cast int to float
	var tenDot1 float32 = 10.1
	var intNumRand int = 2
	var result = tenDot1 + float32(intNumRand)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println("Div Operation: ", (intNum1/intNum2))
	fmt.Println("Mod Operation: ", (intNum1%intNum2))

	var myString string = "Hello" + " " + "world"
	fmt.Println(myString)

	// In go we need len would give the number of bytes in the string in total 
	// so a special character like gamma 𝜸 will give 2 bytes

	// We can use runesCountInString to get total number of characters in a string
	fmt.Println("The number of characters in 'Hello World' are ", utf8.RuneCountInString(myString))

	// Go can set default values for data Types like int can be 0
	var defaultInt int
	fmt.Println("Default Value of int: ", defaultInt)

	var myBool bool = false
	fmt.Println("This is a bool: ", myBool)

	var intNum3 rune = 3
	fmt.Println("This is rune: ", intNum3)

	// These datatypes have 0 as default value
	// uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64, complex64, complex128, rune, byte
	// string is '' by default
	// bool is false by default

	// We can also use shorthand to declare variables
	myName := "Rishi"
	fmt.Println("This is my Name: ", myName)

	num1, num2 := 1, 2
	fmt.Println("This is num1: ", num1, "This is num2: ", num2)


	// Constants are used to store values that do not change
	const myConst string = "This is a constant"
	fmt.Println(myConst)

}