package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%d is even.\n", number)
		} else {
			fmt.Printf("%d is odd.\n", number)
		}
	}
}

// In the main.go file, define a function called 'main'.  Remember that the 'main' function will be called automatically.
// In the main function, create a slice of ints from 0 through 10
// Iterate through the slice with a for loop.  For each element, check to see if the number is even or odd.
// If the value is even, print out "even".  If it is odd, print out "odd"
// Run your code from the terminal by changing into your project directory then running 'go run main.go'.
// Your output should look like this:

// 0 is even
// 1 is odd
// 2 is even
// 3 is odd
// 4 is even
// 5 is odd
// 6 is even
// 7 is odd
// 8 is even
// 9 is odd
// 10 is even
