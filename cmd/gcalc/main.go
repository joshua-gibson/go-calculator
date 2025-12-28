package main

import (
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/your-username/go-calculator/pkg/calculator"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")

	calc := calculator.NewCalculator()
	var currentOp string
	var valueA float64
	isNewNumber := true

	output := widget.NewLabel("0")
	output.Alignment = fyne.TextAlignTrailing
	output.TextStyle.Bold = true

	setDisplayText := func(text string) {
		output.SetText(text)
		isNewNumber = true
	}

	numberButton := func(n int) *widget.Button {
		return widget.NewButton(strconv.Itoa(n), func() {
			if isNewNumber {
				output.SetText(strconv.Itoa(n))
				isNewNumber = false
			} else {
				output.SetText(output.Text + strconv.Itoa(n))
			}
		})
	}

	opButton := func(op string) *widget.Button {
		return widget.NewButton(op, func() {
			var err error
			valueA, err = strconv.ParseFloat(output.Text, 64)
			if err != nil {
				setDisplayText("Error")
				return
			}
			currentOp = op
			isNewNumber = true
		})
	}

	equalsButton := widget.NewButton("=", func() {
		valueB, err := strconv.ParseFloat(output.Text, 64)
		if err != nil {
			setDisplayText("Error")
			return
		}
		result, err := calc.Calculate(valueA, currentOp, valueB)
		if err != nil {
			setDisplayText("Error")
		} else {
			setDisplayText(strconv.FormatFloat(result, 'f', -1, 64))
		}
	})

	clearButton := widget.NewButton("C", func() {
		setDisplayText("0")
		valueA = 0
		currentOp = ""
	})

	w.SetContent(container.NewVBox(
		output,
		container.NewGridWithColumns(4,
			numberButton(7), numberButton(8), numberButton(9), opButton("/"),
			numberButton(4), numberButton(5), numberButton(6), opButton("*"),
			numberButton(1), numberButton(2), numberButton(3), opButton("-"),
			clearButton, numberButton(0), equalsButton, opButton("+"),
		),
	))

	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
