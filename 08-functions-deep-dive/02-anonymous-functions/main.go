package main

import "fmt"

// main is the entry point of the program.
func main() {
	numbers := []int{1, 2, 3}

	// anonymous function is used as a parameter to transformNumbers
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)

	// createTransformer returns a function that multiplies its input by the factor argument.
	double := createTransformer(2)
	triple := createTransformer(3)

	fmt.Println(double(2))
	fmt.Println(triple(2))
}

// transformNumbers applies a transformation function to each element in the numbers slice.
// It returns a new slice with the transformed values.
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// In this case, the anonymous function is a closure that encompasses the factor variable. 
// Even after createTransformer has finished executing, the returned function still has access to factor, allowing it to multiply its input by factor whenever it's called.
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}