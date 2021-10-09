// Uncomment the entire file

package main

// import "fmt"

// func main() {

// 	var city string

// 	switch city {
// 	case "Des Moines":
// 		fmt.Println("You live in Iowa")

// 	// if city is "Minneapolis" OR "St Paul"
// 	case "Minneapolis,", "St Paul":
// 		fmt.Println("You live in Minnesota")

// 	case "Madison":
// 		fmt.Println("You live in Wisconsin")
// 	default:
// 		fmt.Println("You're not from around here.")
// 	}

// 	// ****************************
// 	// var i int

// 	// If we don't put the variable to compare in "switch",
// 	// the condition will evaluate in each "case"
// 	// switch {
// 	// case i > 10:
// 	// 	fmt.Println("Greater than 10")
// 	// case i < 10:
// 	// 	fmt.Println("Less than 10")
// 	// default:
// 	// 	fmt.Println("Is 10")
// 	// }

// 	// ****************************
// 	var i int = 9

// 	// If we write "fallthrough" go will continue evaluating
// 	// the next case no matter if the first case meets the condition
// 	switch {
// 	case i != 10:
// 		fmt.Println("Does not equal 10")
// 		fallthrough
// 	case i < 10:
// 		fmt.Println("Less than 10")
// 	case i > 10:
// 		fmt.Println("Greater than 10")
// 	default:
// 		fmt.Println("Is 10")
// 	}
// }
