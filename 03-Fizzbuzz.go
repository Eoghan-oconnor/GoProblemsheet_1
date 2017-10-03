//Eoghan O'Connor 2017
//Adapted from https://graysonkoonce.com/fizzbuzz-in-golang/


package main

import "fmt" // Formated I/O import 

func main() {  

	// for loop to print out numbers 1 to 100
    for i := 1; i <= 100; i++ {

		// if else if else code block to decide what to print out
		if i%15 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }// end if else if else code block
    }// end for loop
}//end func main 