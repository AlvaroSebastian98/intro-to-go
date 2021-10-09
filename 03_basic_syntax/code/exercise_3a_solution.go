package main

import "fmt"

func main() {

	sentence := "Lucas programa en Go"

	for i, value := range sentence {
		if i%2 != 0 {
			fmt.Println(value, string(value))
		}
	}

	// // If we only want to use the second argument we can
	// // put an underscore into any variable that we are not gonna use
	// for _, value := range sentence {
	// 	fmt.Println(string(value))
	// }

}
