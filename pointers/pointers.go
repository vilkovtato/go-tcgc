package main

import "fmt"


func main() {
	age := 32
	pointerAge := &age

	println("Age: ", age)
	println("Pointer Age: ", pointerAge)
	println("Pointer Age value: ", *pointerAge)

	changeAge(pointerAge)

	fmt.Println("Age: ", age)
}

func changeAge(val *int) {
	*val = *val - 18
}