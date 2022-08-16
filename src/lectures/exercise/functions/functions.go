//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func displayMessage(message string) {
	fmt.Println(message)
}

func addThreeNumbers(numOne, numTwo, numThree int) int {
	return numOne + numTwo + numThree
}

func returnOne() int {
	return 1
}

func returnOneAndTwo() (int, int) {
	return 1, 2
}

func main() {
	greet("Red")
	displayMessage("Hello, I'm Red")

	one, two := returnOneAndTwo()
	four := addThreeNumbers(returnOne(), one, two)
	fmt.Println("The value is", four)

}
