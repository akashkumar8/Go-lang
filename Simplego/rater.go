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

}