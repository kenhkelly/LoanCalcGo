package main

import (
	"fmt"
)

const version string = "v0.1"

func main() {
	fmt.Println("AutoCalcGo! " + version + "\n\n")
	start()
}

func start() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("An error occured:", r)
			start()
		}
	}()
	fmt.Println(`
How would you like to calculate?

(1) Monthly payment
(2) Loan amount from monthly payment`)

	var i string
	_, err := fmt.Scan(&i)
	if err != nil {
		panic(err)
	}

	switch i {
	case "1", "month", "m":
		calcMonthly()
	case "2", "loan", "l":
		calcLoan()
	default:
		fmt.Printf("%s is not a valid selection.\n", i)
		start()
	}

	exit()
}

func exit() {
	fmt.Print("\n\nPerform another?")
	var y string
	_, err := fmt.Scan(&y)
	if err != nil {
		panic(err)
	}
	switch y {
	case "y", "ye", "yes":
		start()
	}
}
