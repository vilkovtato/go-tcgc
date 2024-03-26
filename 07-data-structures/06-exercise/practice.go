package main

import "fmt"

type Product struct {
	title string
	id int
	price float64
}

func main() {

	// 1)
	hobbies := [3]string{"reading", "coding", "gaming"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	chosenHobbies := [2]string{hobbies[1], hobbies[2]}
	fmt.Println(chosenHobbies)

	// 3)
	hobbiesSlice := hobbies[0:2]
	fmt.Println(hobbiesSlice)
	//hobbiesSlice2 := hobbies[:2]
	//fmt.Println(hobbiesSlice2)

	// 4)
	hobbiesSlice = hobbiesSlice[1:3]
	fmt.Println(hobbiesSlice)

	// 5)
	myGoals := []string{"learn Go", "learn data structures"}
	fmt.Println(myGoals)

	// 6)
	myGoals[1] = "learn algorithms"
	myGoals = append(myGoals, "learn generics")
	fmt.Println(myGoals)

	// 7)
	products := []Product{Product{"Pen", 1, 1.99}, Product{"Pencil", 2, 0.99}}
	products = append(products, Product{"Eraser", 3, 0.49})
	fmt.Println(products)
}

