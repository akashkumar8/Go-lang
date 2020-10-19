package main

import "fmt"

func main() {

	var a int = 10

	if a%2 == 0 {

		fmt.Printf("a is even\n")
	} else {

		fmt.Printf("a is odd\n")
	}
	fmt.Printf("value of a is : %d\n", a)
}
