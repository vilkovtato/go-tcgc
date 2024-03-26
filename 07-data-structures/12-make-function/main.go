package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	// this would create empty underlying array
	userNames := []string{}

	// later we would append some items which will lead to creation of new array because the first array was with capacity 0
	userNames = append(userNames, "Max")

	fmt.Println(userNames)

	// if we know at the beginning, that we will add some specific numbers of items we can use make function.
	// it non empty underlying array..
	userNames2 := make([]string, 2)

	userNames2[0] = "Manuel"

	fmt.Println(userNames2)

	// this returns slice of size 2, but underlying array will have maximum items 10
	// it non empty underlying array..
	userNames3 := make([]string, 2, 10)

	userNames3[0] = "Manuel"
	userNames3[1] = "Alexander"

	// now when using append function go doesnt have create new array, because array has enough memory to have 10 items
	// so 3 items will be enough
	userNames3 = append(userNames3, "Vilko")

	fmt.Println(userNames3)

	// the same holds for maps, when creating empty map, when adding item, go will have reallocate memory
	// so using make is the way to use it if we know at leas approximately how many items will there be
	// we will use only one argument - "intended length" and following code will not cause realocating memory
	courseRatings := make(map[string]float64, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.6
	courseRatings["angular"] = 4.5

	// we can use type alias
	courseRatings2 := make(floatMap, 3)
	courseRatings2["java"] = 4.7
    courseRatings2.output()

	// for loop for every collections
	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	// for map more appropriate name for the first value is "key"
	for key, value := range courseRatings {
		fmt.Println(key)
		fmt.Println(value)
	}
}
