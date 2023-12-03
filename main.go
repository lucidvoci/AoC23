package main

import (
	"advent/days"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("Input an advent day as an argument.")
		return
	}

	switch args[1] {
	case "1":
		days.Day1()
	case "2":
		days.Day2()
	case "3":
		days.Day3()
	default:
		fmt.Println("Not an advent day.")
	}

}
