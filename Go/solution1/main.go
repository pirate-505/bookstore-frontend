package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type bookPrice struct {
	money    uint
	currency rune
}

func parseBooks(args []string) ([]bookPrice, error) {
	// books := make([]bookPrice, len(args)-1)
	var books []bookPrice
	for _, arg := range args[1:] {
		currency, _ := utf8.DecodeLastRuneInString(arg)
		arg = strings.TrimSuffix(arg, string(currency))
		cost, err := strconv.Atoi(arg[:len(arg)])
		if err != nil {
			return nil, err
		}

		book := bookPrice{
			money:    uint(cost),
			currency: currency,
		}
		books = append(books, book)
	}
	return books, nil
}

func printTotal(prices []bookPrice) (uint, rune) {
	var total uint
	for _, price := range prices {
		total += price.money
	}
	return total, prices[0].currency
}

func main() {
	prices, err := parseBooks(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	total, currency := printTotal(prices)
	fmt.Printf("To pay: %d%v\n", total, string(currency))
}
