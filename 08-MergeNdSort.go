//Eoghan O'Connor 2017
//Go Problem Sheet 1
//08-MergeNdSort
//Ref: https://golang.org/pkg/sort/#Ints
//Merge adapted from https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go

package main 

import (
	"fmt" //Formatted I/O import 
	"sort" // Sorting Slices and user-defined collections
)

func main(){

	//Declaring variables
	arrayA := [] int {1,3,5}
	arrayB := [] int {2,4,6}

	//Merging arrays 
	arrayC := append(arrayA,arrayB...)

	//Sorting arrays
	 sort.Ints(arrayC)

	//Printing out Arrays 
	fmt.Println("Array A: ",arrayA)
	fmt.Println("\nArray B: ",arrayB)
	fmt.Println("\nArray C (A and B sorted and merged): ",arrayC)




	
}