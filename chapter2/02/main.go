package main

import "fmt"

var name string

func main() {
	fmt.Printf("What is your name? ")
	fmt.Scanf("%s", &name)

	if len(name) == 0 {
		fmt.Println("Please provide your name.")
	} else {
		fmt.Printf("%v has %v characters.\n", name, len(name))
	}
}
