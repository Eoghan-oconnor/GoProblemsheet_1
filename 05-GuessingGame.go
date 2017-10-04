//Eoghan O'Connor 2017
// Go Work Sheet 1 Problem 5.

package main 

import (
	"fmt" // Formatted I/O import 
	"math/rand" // Random Number import.
	"time" // 
)

//This function was adapted form adapted from http://golangcookbook.blogspot.ie/2012/11/generate-random-number-in-given-range.html

func randomNumGen(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func main(){

	//Declaring Variables and calling function
	var input int
	randomNum := randomNumGen(1,100)
	var prevousNum int
	var found bool = false
	var count int = 0

	//Heading
	fmt.Println("I want to play a Game")
	fmt.Println("=========================")

//fmt.Println(randomNum) used cto make sure the number was actually random

	//Loop to allow user to play the game 
	for found != true{

		//user prompt to enter a number 
		fmt.Println("Please Enter a Number between 1-100: ")
		fmt.Scan(&input)
		count++

		if input == prevousNum{
			count--
			fmt.Println("Why Pick The Same Number Twice?")
			fmt.Println("\nTry again")
		}

		//if else if code block to sort number 
		if input > randomNum{
			fmt.Println("Number was too large")
		}else if input < randomNum{
			fmt.Println("Number was too small")
		}else{
			fmt.Println("\nWell Done you guessed correctly")
			fmt.Println("And it only took you",count,"tries")
			found = true // ends loop
		}

		prevousNum = input



	}

}
