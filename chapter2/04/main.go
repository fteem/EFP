package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var noun string
var verb string
var adjective string
var adverb string

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Press CTRL/CMD-C to exit...")
	for {
		fmt.Print("Enter a noun:")
		noun, _ := in.ReadString('\n')
		noun = strings.Trim(noun, "\n")

		fmt.Print("Enter a verb:")
		verb, _ := in.ReadString('\n')
		verb = strings.Trim(verb, "\n")

		fmt.Print("Enter an adjective:")
		adjective, _ := in.ReadString('\n')
		adjective = strings.Trim(adjective, "\n")

		fmt.Print("Enter an adverb:")
		adverb, _ := in.ReadString('\n')
		adverb = strings.Trim(adverb, "\n")

		fmt.Printf("Do you %s your %s %s %s? That's hilarious!\n\n", verb, adjective, noun, adverb)
	}
}
