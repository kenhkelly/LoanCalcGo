package main

import (
	"fmt"
	"math"
)

func calcMonthly() {
	var (
		loan    float64
		apr     float64
		term    float64
		monthly float64
	)
	for {
		fmt.Print("Enter loan amount: ")
		_, err := fmt.Scan(&loan)
		if err != nil {
			fmt.Println("Invalid loan input.")
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

	//M = P * ( J / (1 - (1 + J)-N))
	j := (apr / 100 / 12)
	term = term * 12
	monthly = loan * (j / (1 - math.Pow(1+j, -term)))
	fmt.Printf("Monthly payment is: %.2f", monthly)
}
