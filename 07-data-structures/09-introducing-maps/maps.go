package main

import "fmt"

func main() {

	// map vs struct - the same type, you can add key/value
	// map vs array - its an array with custom index, which can be string/float/struct

	websites := map[string]string{
		"google": "https://www.google.com",
		"amazon web services": "https://aws.com",
	}

	fmt.Println(websites)

	// you can always add new key/value pair to map
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	// remove key/value pair from map
	delete(websites, "google")
	fmt.Println(websites)
}