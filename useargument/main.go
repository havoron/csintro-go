package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hi, ")
	fmt.Print(os.Args[1])
	fmt.Println("! How are you?")
}
