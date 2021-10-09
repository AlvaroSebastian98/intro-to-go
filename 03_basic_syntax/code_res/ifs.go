// Uncomment this entire file

package main

// import (
// 	"fmt",
// 	"errors"
// )

// func someFunction() error {
// 	return errors.New("some error")
// }

// func main() {

// 	var someVar = 9

// 	// Curly braces "{}" are required,
// 	// but parentheses are not
// 	if someVar > 10 {
// 		fmt.Println(someVar)
// 	}

// 	// ****************************

// 	if someVar > 100 {
// 		fmt.Println("Greater than 100")
// 	} else if someVar == 100 {
// 		fmt.Println("Equals 100")
// 	} else {
// 		fmt.Println("Less than 100")
// 	}

// 	// ****************************
// 	// Handling errors

// 	err := someFunction()
// 	// => If this function returns a value,
// 	// => it will be an  error of type Error

// 	// ****************************
// 	if err != nil {
// 	  fmt.Println(err.Error())
// 	}

// 	// if we define a variable within an "IF" blocked
// 	// is only accessible in this scope
// 	if err := someFunction(); err != nil {
// 		// err variable is accessible only inside "IF"
// 	  fmt.Println(err.Error())
// 	}

// 	// End of file curly brace
// }
