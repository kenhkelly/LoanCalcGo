package main

import (
	"fmt"
	"math"
)

func calcLoan() {
	var (
		loan    float64
		apr     float64
		term    float64
		monthly float64
	)
	for {
		fmt.Print("Enter monthly amount: ")
		_, err := fmt.Scan(&monthly)
		if err != nil {
			fmt.Println("Invalid monthly input.")
			continue
		}
		break
	}
	for {
		fmt.Print("Enter annual interest rate (APR): ")
		_, err := fmt.Scan(&apr)
		if err != nil {
			fmt.Println("Invalid apr input.")
			continue
		}
		break
	}
	for {
		fmt.Print("Enter loan term in years: ")
		_, err := fmt.Scan(&term)
		if err != nil {
			fmt.Println("Invalid term input.")
			continue
		}
		break
	}

	//P = M / ( J / (1 - (1 + J)-N))
	j := (apr / 100 / 12)
	term = term * 12
	loan = monthly / (j / (1 - math.Pow(1+j, -term)))
	fmt.Printf("Loan is: %.2f", loan)
}
