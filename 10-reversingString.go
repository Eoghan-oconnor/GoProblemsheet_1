//Eoghan O'Connor 2017
//Go Problem Sheet 1
//10-reversingString.go
//Adapted from https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go

package main

import (
	"fmt"//Formatted I/O import
	
)

func main(){

	//Declaring string
	s := "We don't desevre dogs"

	//Printing out both strings
	fmt.Println("The original String is: ",s)
	fmt.Println("The Reserse of the Strong is: ",reversingString(s))
}

//Function to reserve string
func reversingString (s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)	
}