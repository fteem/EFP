package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("What is the first number? ")
	rawFirstNumber, _ := in.ReadString('\n')
	firstNumber, err := strconv.Atoi(strings.Trim(rawFirstNumber, "\n"))
	if err != nil || firstNumber < 0 {
		fmt.Println("Please enter a valid positive number.")
		os.Exit(1)
	}

	fmt.Print("What is the second number? ")
	rawSecondNumber, _ := in.ReadString('\n')
	secondNumber, err := strconv.Atoi(strings.Trim(rawSecondNumber, "\n"))
	if err != nil || secondNumber < 0 {
		fmt.Println("Please enter a valid positive number.")
		os.Exit(1)
	}

	fmt.Printf("%[1]v + %[2]v = %[3]v \n%[1]v - %[2]v = %[4]v \n%[1]v * %[2]v = %[5]v \n%[1]v / %[2]v = %[6]v\n", firstNumber, secondNumber, sum(firstNumber, secondNumber), subtract(firstNumber, secondNumber), multiply(firstNumber, secondNumber), divide(firstNumber, secondNumber))
}

func sum(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}
