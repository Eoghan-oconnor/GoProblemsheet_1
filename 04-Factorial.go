//Eoghan O'Connor 2017
// Go Work Sheet 1 Problem 4.

package main

import (
	"fmt"      // Formatted I/O import 
	"math/big" // Math Import
	"strconv"  // String Conversion Import 
)

func digitSum(input int) int {

	factorial := big.NewInt(1) 
	var factorialSum, num int  
	

	// Calculating the Factorial of the user inputted number.
	for i := 1; i <= input; i++ {
		factorial.Mul(factorial, big.NewInt(int64(i))) 
	}
	fmt.Println(input,"!  : ",factorial)//Outputting the factorial


	//calculate sum of factorial nums
	for i := range factorial.String() {
			num, _ = strconv.Atoi(string(factorial.String()[i])) 
			factorialSum += num                                 
	}
	return factorialSum
	
}//digitSum

func main() {
	//take user input from console
	var input int
	fmt.Print("Please enter a number(1-100): " )
	fmt.Scan(&input)

	fmt.Println("Input : ",input)
	fmt.Println("Sum of Factorial nums : ",digitSum(input))
}