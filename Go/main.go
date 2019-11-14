package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type bookPrice struct {
	money    uint
	currency string
}

func printTotal(prices []bookPrice) {
	var total uint
	for _, price := range prices {
		total += price.money
	}
	fmt.Printf("To pay: %d%s\n", total, prices[len(prices)-1].currency)
}

func parseBooks() ([]bookPrice, error) {
	books := make([]bookPrice, len(os.Args)-1)
	for _, arg := range os.Args[1:] {
		cost, err := strconv.Atoi(arg[:len(arg)-1])
		if err != nil {
			return nil, err
		}
		book := bookPrice{
			money:    uint(cost),
			currency: string(arg[len(arg)-1]),
		}
		books = append(books, book)
	}
	return books, nil
}

func main() {
	prices, err := parseBooks()
	if err != nil {
		log.Fatal(err)
	}
	printTotal(prices)
}
