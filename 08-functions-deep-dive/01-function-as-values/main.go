package main

import "fmt"

// transformer is a type that represents a function that takes an int as input and returns an int as output.
// It is used as a parameter type in the transformNumbers2 function to make the code more readable.
type transformer func(int) int

func main() {
	numbers := []int{1, 2, 3}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers2(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	bigNumbers := []int{100, 200, 300}
	doubled2 := transformNumbers(&numbers, getTransformerFunction(&numbers))
	tripled2 := transformNumbers2(&bigNumbers, getTransformerFunction(&bigNumbers))

	fmt.Println(doubled2)
	fmt.Println(tripled2)
}

// transformNumbers is a function that takes a pointer to a slice of integers and a transform function as arguments.
// It applies the transform function to each element in the slice and returns a new slice with the transformed values.
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	transformedNumbers := []int{}
	for _, number := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(number))
	}
	return transformedNumbers
}

// transformNumbers2 is a function that takes a pointer to a slice of integers and a transformer function as arguments.
// It applies the transformer function to each element in the slice and returns a new slice with the transformed values.
// the transformer type represents a function that takes an int as input and returns an int as output.
func transformNumbers2(numbers *[]int, transform transformer) []int {
	transformedNumbers := []int{}
	for _, number := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(number))
	}
	return transformedNumbers
}


// getTransformerFunction is a function that takes a pointer to a slice of integers as input and returns a transformer function.
func getTransformerFunction(numbers *[]int) transformer {
	if (*numbers)[0] == 1{
		return double
	}
	return triple
}

// double is a function that takes an integer as input and returns the double of that integer.
func double(n int) int {
	return n * 2
}

// triple is a function that takes an integer as input and returns the triple of that integer.
func triple(n int) int {
	return n * 3
}
