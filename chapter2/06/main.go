package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("What is your current age? ")
	rawAge, _ := in.ReadString('\n')
	age, err := strconv.Atoi(strings.Trim(rawAge, "\n"))
	if err != nil || age < 0 {
		fmt.Println("Please enter a valid positive number.")
		os.Exit(1)
	}

	fmt.Print("At what age would you like to retire? ")
	rawRetirement, _ := in.ReadString('\n')
	retirement, err := strconv.Atoi(strings.Trim(rawRetirement, "\n"))
	if err != nil || retirement < 0 {
		fmt.Println("Please enter a valid positive number.")
		os.Exit(1)
	}

	yearsUntilRetirement := untilRetirement(age, retirement)
	if yearsUntilRetirement < 1 {
		fmt.Println("Well, good news. You can already retire!")
	} else {
		fmt.Printf("You have %d years left until you can retire.\n")
		currentYear, retirementYear := years(age, retirement)
		fmt.Printf("It's %d, so you can retire in %d.\n", currentYear, retirementYear)
	}

}

func untilRetirement(currentAge, retirementAge int) int {
	return retirementAge - currentAge
}

func years(currentAge, retirementAge int) (int, int) {
	t := time.Now()
	currentYear := t.Year()
	retirementYear := currentYear + untilRetirement(currentAge, retirementAge)

	return currentYear, retirementYear
}
