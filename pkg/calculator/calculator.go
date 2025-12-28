package calculator

import (
	"fmt"
)

// Operation is a function that takes two float64s and returns a float64 and an error.
type Operation func(a, b float64) (float64, error)

// Calculator holds a map of operator strings to Operation functions.
type Calculator struct {
	operations map[string]Operation
}

// NewCalculator creates and initializes a new Calculator.
func NewCalculator() *Calculator {
	c := &Calculator{
		operations: make(map[string]Operation),
	}
	c.operations["+"] = add
	c.operations["-"] = subtract
	c.operations["*"] = multiply
	c.operations["/"] = divide
	return c
}

// Calculate performs the calculation for the given operator.
func (c *Calculator) Calculate(a float64, op string, b float64) (float64, error) {
	operation, ok := c.operations[op]
	if !ok {
		return 0, fmt.Errorf("unknown operator: %s", op)
	}
	return operation(a, b)
}

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
