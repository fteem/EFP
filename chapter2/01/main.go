package main

import (
	"fmt"
	"strings"
)

var greetings = map[string]string{
	"ilija": "Hey man, how's it going?",
	"mike":  "Long time no see, eh?",
}
var name string

func main() {
	fmt.Printf("What is your name? ")
	fmt.Scanf("%s", &name)

	if greeting, exists := greetings[strings.ToLower(name)]; exists {
		fmt.Println(greeting)
	} else {
		fmt.Printf("Hello, %v, nice to meet you!\n", name)
	}
}
