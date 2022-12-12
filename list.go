package main

import (
	"fmt"
	"strings"
)

type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

var menu = []Menu{
	//itemNo  itemName   itemPrice
	{1, "suits", 20},
	{2, "Gala dresses", 25},
	{3, "Duves", 35.50},
	{4, "Sheets", 30},
	{5, "50 clothes", 60},
	{6, "underwear", 10},
	{7, "Shoes", 25},
	{8, "Jeans 2 pairs", 30},
	{9, "Sweaters 2 units", 30},
	{10, "Backpacks", 20},
	{11, "Strong Stain remover for white t-shirts", 30.00},
}

func printList() {
	fmt.Printf("%15s\n", "Washing Options")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "S.No.", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, element := range menu {
		fmt.Printf(" %-7d %.2f    %-4s\n", element.itemNo, element.itemPrice, element.itemName)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}
