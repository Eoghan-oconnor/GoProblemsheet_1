//Eoghan O'Connor 2017
//Go Problem Sheet 1
// 09-NewtonsSqRoot.go

package main

import (
	"fmt" //Formatted I/O import
	"math" //Math import 
)

func newtonSquareRoot(z float64, x float64)float64{
	if z ==0{
		return 0
	}

	return z - (((z*z) - x) / (2*z))// newtons formula for square root	
}

func main() {
	x := 99.0 // Number to find the square root of
	z := 1.0  // Initial guess

	// Run code while z doesn't equal the guess
	for z = 1.0; z != newtonSquareRoot(z, x); z = newtonSquareRoot(z, x) {
		fmt.Printf("The current guess is: %2.6f\n", z) // print guess
	}

	// Print the approximation according to Netwon's Method for sqaure root
	fmt.Printf("\n%2f is an approximation for the square root of %2f\n", z, x)

	// Print the square root according to math.Sqrt()
	fmt.Printf("The square root using math.Sqrt is %2f\n", math.Sqrt(x))
}