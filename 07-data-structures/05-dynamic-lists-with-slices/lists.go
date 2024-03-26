package main

import "fmt"

func main() {

	// this will create a slice and underlying array
	// if the slice by your actions grows beyond the capacity of the underlying array, a new array will be created
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])
	prices[1] = 9.99
	// you cannot change index which does not exists:
	// prices[2] = 5.99

	// we could use append to add a new element to the slice
	// this will copy the original array to a new array with a new capacity
	// and return new slice
	updatedPrices := append(prices, 5.99)
	fmt.Println("Updated Prices:", updatedPrices)
	// the original slice is not changed
	fmt.Println("Prices:", prices)


}
