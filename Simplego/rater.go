package main

import(
	"fmt"
	"build"
	"os"
	"strcopy"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	reader := bufio.NewReader(os.stdin)
	fmt.Println("Please enter your full name")
	name, _ := reader.ReadString(\n)
	
	
	reader = bufio.NewReader(os.stdin)
	fmt.Println("Please rate our Kloud learning Officials between 1 and 5: ")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)


	fmt.Printf(%v, %v, name,mynumRating)

}
