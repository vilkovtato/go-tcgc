package main

import "fmt"

func main() {
	// Calculate the sum of numbers using the sumup function
	sum := sumup(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(sum)

	numbers := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

	// unpack the numbers slice and pass it to the sumup function
	anotherSum := sumup(1, numbers...)

	fmt.Println(anotherSum)

}

// sumup calculates the sum of numbers provided as arguments.
// It takes a startingValue as the first argument, followed by a variadic parameter numbers.
// It returns the sum of all the numbers.
func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, value := range numbers {
		sum += value
	}

	return sum
}
