package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var quote string
var name string

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("What is the quote?\t")
	quote, _ := in.ReadString('\n')

	fmt.Print("Who said it?\t")
	name, _ := in.ReadString('\n')

	fmt.Printf("%v says, \"%v\"\n", strings.Trim(name, "\n"), strings.Trim(quote, "\n"))
}
