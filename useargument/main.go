package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("len(os.Args) =", len(os.Args))
	fmt.Print("Hi, ")
	fmt.Print(os.Args[1])
	fmt.Println("! How are you?")
}
