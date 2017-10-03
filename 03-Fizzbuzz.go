// Eoghan O'Connor 2017
// This program was written with help from http://wiki.c2.com/?FizzBuzzTest

package main

import 
(
	// Formatted I/O import
	"fmt"
)
func main (){
	var i int = 1
	
	for i <= 100{
		

		//if else block to print fizz, buzz and fizzbuzz
		if (i % 3 == 0 && i % 5 == 0) {
            fmt.Println("Fizzbuzz")
		} else if (i % 3 == 0) {
            fmt.Println("Fizz")
		} else if (i % 5 == 0) {
            fmt.Println("Buzz")
		} else {
            fmt.Println(i)
        }
		i += 1
	}// end for 

}//end func main

