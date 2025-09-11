package main

import (
	"fmt"
	"strings"
)

func main() {
	Section("Strings, Runes && Bytes", runStrings)
}

func runStrings() {
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("Str[0]: %v, Type: %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Printf("[I V] = [%v %v]\n", i, v)
	}

	fmt.Printf("Len(myString) = %v\n", len(myString))

	var myRune = 'a'
	fmt.Printf("MyRune = %v\n", myRune)

	var strSlice = []string{"H","e","l","l","o",","," ","W","o","r","l","d","!"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("CatStr: %s\n", catStr)

	var sb strings.Builder
	for i := range strSlice {
		sb.WriteString(strSlice[i])
	}
	fmt.Printf("BuildStr: %s\n", sb.String())
}

 
