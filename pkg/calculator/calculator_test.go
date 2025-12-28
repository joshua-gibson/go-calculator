package calculator

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	c := NewCalculator()

	testCases := []struct {
		Name        string
		A           float64
		B           float64
		Op          string
		Expected    float64
		ExpectError bool
	}{
		{Name: "Addition", A: 5, B: 10, Op: "+", Expected: 15, ExpectError: false},
		{Name: "Subtraction", A: 10, B: 5, Op: "-", Expected: 5, ExpectError: false},
		{Name: "Multiplication", A: 5, B: 10, Op: "*", Expected: 50, ExpectError: false},
		{Name: "Division", A: 10, B: 5, Op: "/", Expected: 2, ExpectError: false},
		{Name: "Division by zero", A: 10, B: 0, Op: "/", Expected: 0, ExpectError: true},
		{Name: "Unknown operator", A: 10, B: 5, Op: "%", Expected: 0, ExpectError: true},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := c.Calculate(tc.A, tc.Op, tc.B)

			if tc.ExpectError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error, but got: %v", err)
				}
				if result != tc.Expected {
					t.Errorf("Expected %f, but got %f", tc.Expected, result)
				}
			}
		})
	}
}
