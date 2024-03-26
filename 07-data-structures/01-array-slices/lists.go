package main

import (
	"fmt"
)

func main() {

	productNames := [4]string{"apple"}
	prices := []float64{1.23, 2.34, 3.45, 4.56, 5.67}

	// get type of prices
	fmt.Printf("Type of variable1: %T\n", prices)

	productNames[2] = "banana"

	fmt.Println("Product Names:", productNames)
	fmt.Println("Prices:", prices)

	// create a slice from an array
	featuredPrices := prices[1:]
	fmt.Println("FeaturedPrices is slice from prices:", featuredPrices)

	// create a slice from a slice
	highlightedPrices := featuredPrices[:1]
	fmt.Println("HighlightedPrices is slice from featuredPrices:", highlightedPrices)

	// modify slice and see the effect on the original slice and the array
	highlightedPrices[0] = 199.99
	fmt.Println(highlightedPrices)
	fmt.Println(prices)

	// len and cap are the same because we are creating a slice from an array
	fmt.Println("len and cap of featuredPrices:", len(featuredPrices), cap(featuredPrices))

	// len and cap are different because we are creating a slice from a slice
	fmt.Println("len and cap of highlightedPrices:", len(highlightedPrices), cap(highlightedPrices))
	// because len=1 and cap=4, we can add 3 more elements to the slice
	// why this does work?
	// because the slice is created from a slice, and the original slice war created from an array
	// because the slice is window into the array, because its original slice is window to the array
	// so both slices are window to the underlying array
	// they can change the array, and the array can change them
	// so the highlightedPrices slice can have up to 4 elements, not 5 because the original slice was from 1 element of the array
	highlightedPrices = highlightedPrices[:4]
	fmt.Println("highlightedPrices after reslicing:", highlightedPrices)
	fmt.Println("len and cap of highlightedPrices after reslicing:", len(highlightedPrices), cap(highlightedPrices))

	// unpacking a slice
	newPrices := append(prices, highlightedPrices...)
	fmt.Println(newPrices)
}
