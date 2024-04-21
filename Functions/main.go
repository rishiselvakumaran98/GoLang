package main

import (
	"fmt"
	"errors"
)

func main() {
	printMeStr := "Hello World!"
	printMe(printMeStr)

	q, r, err := intDivision(11,2)
	switch {
		case err != nil:
			fmt.Println(err.Error())
		case r == 0:
			fmt.Printf("The result of the integer division is %v", q)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v\n", q, r)
	}
	// conditional switch statement
	switch r {
	case 0:
		fmt.Println("the division was exact")
	case 1,2:
		fmt.Println("the division was close")
	default:
		fmt.Println("the division was not close")
	}
	
}

func printMe(printVar string){
	fmt.Println(printVar)
}

func intDivision(numerator int, denominator int) (int, int, error){
	var err error // this will be nil if there are no errors
	if denominator == 0{
		err = errors.New("Cannot divide by Zero")
		return 0, 0, err
	}
	var quotient int = numerator/denominator
	var remainder int = numerator%denominator
	return quotient, remainder, err
}