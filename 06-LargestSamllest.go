//Eoghan O'Connor 2017
// Go Worksheet 1
// 06-LasrgestSmartest

package main 

import (
	"fmt" // Formatted I/O import
)

func main (){

	//Declaring Variables and array of numbers 
	array := [7] int {2,4,7,9,78,9,3} 
	var largest int =0 
	var smallest int =0
	 

	//initializing variables 
	largest = array[0]
	smallest = array[0]


	// For loop with nested if statments to assign
	// largest and smallest
	for i := range array{

		if array[i] < smallest{
			smallest = array[i]
		} 
		if array[i] > largest{
			largest = array[i]
		}

	}

	fmt.Println("The smallest number is: ",smallest,"\nThe largest number is: ",largest)
	

}