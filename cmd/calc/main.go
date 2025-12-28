package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/your-username/go-calculator/pkg/calculator"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: ./calc <valueA> <operator> <valueB>")
		os.Exit(1)
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid value: %s\n", os.Args[1])
		os.Exit(1)
	}

	op := os.Args[2]

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid value: %s\n", os.Args[3])
		os.Exit(1)
	}

	c := calculator.NewCalculator()
	result, err := c.Calculate(a, op, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
