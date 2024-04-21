package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = "rèsumè" 
	// we are indexing through the underlying byte-array in the string
	var indexed = myString[1]
	fmt.Printf("Indexed char = %v", indexed)
	for i, v := range myString{
		fmt.Println(i, v)
	}
	// Taking len of the string would be taking length of the bytes in the string
	// This is where we use Runes instead
	var rnString = []rune("rèsumè") 
	// we are indexing through the underlying byte-array in the string
	var ind_rn = rnString[1]
	fmt.Printf("Indexed char = %v", ind_rn)
	for i, v := range rnString{
		fmt.Println(i, v)
	}

	// Strings are immutable in GO
	var myRune = 'a'
	fmt.Printf("myRune: %v\n", myRune)
	// lets concate String; But the operation below is very inefficient
	var strSlice = []string{"r","i","s","h","i"}
	var catStr = ""
	for i := range strSlice{
		catStr += strSlice[i]
	}
	fmt.Printf("catStr: %v\n", catStr)
	// use stringbuilder for a more efficient string operation
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Printf("strBuilder: %v\n", strBuilder.String())

}