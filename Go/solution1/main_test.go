package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func Test_parseBooks(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    []bookPrice
		wantErr bool
	}{
		{
			name: "First-dollars",
			args: []string{"command-name", "5$", "2$", "81$"},
			want: []bookPrice{
				{
					money:    5,
					currency: '$',
				},
				{
					money:    2,
					currency: '$',
				},
				{
					money:    81,
					currency: '$',
				},
			},
			wantErr: false,
		},
		{
			name: "Second-euros",
			args: []string{"command-name", "5€", "2€", "81€"},
			want: []bookPrice{
				{
					money:    5,
					currency: '€',
				},
				{
					money:    2,
					currency: '€',
				},
				{
					money:    81,
					currency: '€',
				},
			},
			wantErr: false,
		},
		{
			name:    "Third-err",
			args:    []string{"8#9$", "1asd$", "8531$"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBooks(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printTotal(t *testing.T) {
	type args struct {
		prices []bookPrice
	}
	tests := []struct {
		name  string
		args  []bookPrice
		want  uint
		want1 rune
	}{
		{
			name: "First",
			args: []bookPrice{
				{
					money:    5,
					currency: '€',
				},
				{
					money:    2,
					currency: '€',
				},
				{
					money:    81,
					currency: '€',
				},
			},
			want:  88,
			want1: '€',
		},
		{
			name: "Second",
			args: []bookPrice{
				{
					money:    25,
					currency: '$',
				},
				{
					money:    6758,
					currency: '$',
				},
				{
					money:    82,
					currency: '$',
				},
			},
			want:  6865,
			want1: '$',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := printTotal(tt.args)
			if got != tt.want {
				t.Errorf("printTotal() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("printTotal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func benchParseBooks(number int, b *testing.B) {
	var input []string
	for i := 0; i < number; i++ {
		price := rand.Uint64()
		priceString := fmt.Sprint(price)
		priceString += string('$')
		input = append(input, priceString)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseBooks(input)
	}
}
func BenchmarkParseBooks10(b *testing.B)    { benchParseBooks(10, b) }
func BenchmarkParseBooks100(b *testing.B)   { benchParseBooks(100, b) }
func BenchmarkParseBooks1000(b *testing.B)  { benchParseBooks(1000, b) }
func BenchmarkParseBooks10000(b *testing.B) { benchParseBooks(10000, b) }

func benchPrintTotal(number int, b *testing.B) {
	var input []bookPrice
	for i := 0; i < number; i++ {
		price := rand.Uint64()
		book := bookPrice{
			money:    uint(price),
			currency: '$',
		}
		input = append(input, book)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		printTotal(input)
	}
}

func BenchmarkPrintTotal10(b *testing.B)    { benchPrintTotal(10, b) }
func BenchmarkPrintTotal100(b *testing.B)   { benchPrintTotal(100, b) }
func BenchmarkPrintTotal1000(b *testing.B)  { benchPrintTotal(1000, b) }
func BenchmarkPrintTotal10000(b *testing.B) { benchPrintTotal(10000, b) }
