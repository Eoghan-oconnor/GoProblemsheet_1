// Eoghan O'Connor 2017
// This program was written with help from http://wiki.c2.com/?FizzBuzzTest

package main

import 

	// Formatted I/O import
	"fmt"


func main (){
	var count int
	
	for count <= 100{
		count++

		if count % 3{
			fmt.Println("Fizz")
		}


		fmt.Println(count)
	}

}

