// Eoghan O'Connor
// Go Worksheet 1
// Palindrome tester 
// Adapted from http://www.golangpro.com/2016/01/check-string-palindrome-go.html

package main

import (
	"fmt" // Formatted I/O import 
	"strings" // Manipulating Strings import 
)

func main() {

	//Declaring Variables 
	var input string

	// User prompt to enter a string
	fmt.Println("Enter string you would like to check:")
	fmt.Scanf("%s\n", &input)
	// Makign all characters lower case 
	input = strings.ToLower(input)
	// Passing arguments to palindromeTester function
	fmt.Println(palindromeTester(input))
}

//Function to test if the string entered is a Palindrome
func palindromeTester(s string) string {
	// Calculates mid point of word and last index
	mid := len(s) / 2
	last := len(s) - 1
	// Used to ensure programme running correctly
	//fmt.Println(mid)
	//fmt.Println(last)
	// Compares letters in opposite positions 
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
		return "NO. It's not a Palimdrome."
		}
	}
	return "YES! You've entered a Palindrome"
}