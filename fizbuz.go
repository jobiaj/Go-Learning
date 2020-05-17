package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 15

	/* check the boolean condition */
	if a%3 == 0 && a%5 == 0 {
		/* if condition is true then print the following */
		fmt.Printf("Fizz Buzz")
	} else if a%3 == 0 {
		/* if else if condition is true */
		fmt.Printf("Fizz")
	} else if a%5 == 0 {
		/* if else if condition is true  */
		fmt.Printf("Buzz")
	} else {
		/* if none of the conditions is true */
		fmt.Printf("Not valid")
	}
}
