package objects

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MakeNumberConverter() *fyne.Container {
	decimalInput := widget.NewEntry()
	decimalInput.TextStyle.Monospace = true
	decimalInput.SetPlaceHolder("decimal number")

	hexInput := widget.NewEntry()
	hexInput.TextStyle.Monospace = true
	hexInput.SetPlaceHolder("hexadecimal number")

	octalInput := widget.NewEntry()
	octalInput.TextStyle.Monospace = true
	octalInput.SetPlaceHolder("octal number")

	binaryInput := widget.NewEntry()
	binaryInput.TextStyle.Monospace = true
	binaryInput.SetPlaceHolder("binary number")

	decimalInput.OnChanged = func(newString string) {
		intNum, err := strconv.ParseInt(newString, 10, 0)
		if err == nil {
			hexInput.SetText(fmt.Sprintf("%x", intNum))
			octalInput.SetText(fmt.Sprintf("%o", intNum))
			binaryInput.SetText(fmt.Sprintf("%b", intNum))
		}
	}
	hexInput.OnChanged = func(newString string) {
		intNum, err := strconv.ParseInt(newString, 16, 0)
		if err == nil {
			decimalInput.SetText(fmt.Sprintf("%d", intNum))
			octalInput.SetText(fmt.Sprintf("%o", intNum))
			binaryInput.SetText(fmt.Sprintf("%b", intNum))
		}
	}
	octalInput.OnChanged = func(newString string) {
		intNum, err := strconv.ParseInt(newString, 8, 0)
		if err == nil {
			decimalInput.SetText(fmt.Sprintf("%d", intNum))
			hexInput.SetText(fmt.Sprintf("%x", intNum))
			binaryInput.SetText(fmt.Sprintf("%b", intNum))
		}
	}
	binaryInput.OnChanged = func(newString string) {
		intNum, err := strconv.ParseInt(newString, 2, 0)
		if err == nil {
			decimalInput.SetText(fmt.Sprintf("%d", intNum))
			hexInput.SetText(fmt.Sprintf("%x", intNum))
			octalInput.SetText(fmt.Sprintf("%o", intNum))
		}
	}

	formatNumObj := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Decimal"), decimalInput,
		widget.NewLabel("Hexadecimal"), hexInput,
		widget.NewLabel("Octal"), octalInput,
		widget.NewLabel("Binary"), binaryInput,
	)

	return formatNumObj
}
