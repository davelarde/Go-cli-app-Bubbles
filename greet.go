package main

import (
	"fmt"
)

func greet(customerName string) {
	fmt.Printf("%52s %s\n", "Hello", customerName)
	fmt.Printf("%72s\n", "_/\\_ Welcome to Travel Bubble! _/\\_ ")
	fmt.Println()
}

func sayTata(customerName string) {
	fmt.Println()
	fmt.Printf("%17s", " ")
	fmt.Printf("_/\\_ Thank you %v for visiting our Bubble App! _/\\_\n", customerName)
	fmt.Printf("%55s\n", "We hope to see you again!")
	fmt.Printf("%51s\n", "Have a nice day! ")
}
