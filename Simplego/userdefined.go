package main()

import "fmt"



func main() {

	fmt.Printf("Enter a number:")
	var input int
	fmt.Scanln(&input)
	fmt.Print(input)
	if(input %2== 0){
		fmt.Printf("is even")
	   } else {
		   fmt.Printf("is odd")
	   }
	
	}

}
