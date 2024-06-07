package objects

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MircoT/go-dev-toys/pkg/format"
)

type formatTarget int

const (
	formatInputNumRows = 16
)

const (
	FORMATJSON formatTarget = iota
)

func MakeFormatText(target formatTarget) *fyne.Container {
	textInput := widget.NewEntry()

	switch target {
	case FORMATJSON:
		textInput.SetPlaceHolder("JSON")
	}

	textInput.MultiLine = true
	textInput.TextStyle.Monospace = true
	textInput.SetMinRowsVisible(formatInputNumRows)
	textInput.Wrapping = fyne.TextWrapBreak

	textOutput := widget.NewEntry()

	textOutput.SetPlaceHolder("Formatted text")

	textOutput.MultiLine = true
	textOutput.TextStyle.Monospace = true
	textOutput.SetMinRowsVisible(formatInputNumRows * 3)
	textOutput.Wrapping = fyne.TextWrapBreak

	textInput.OnSubmitted = func(newString string) {
		var (
			result string
			err    error
		)

		switch target {
		case FORMATJSON:
			result, err = format.JSON(newString)
		}

		if err == nil {
			textOutput.SetText(result)
		}
	}
	textInput.OnChanged = func(newString string) {
		var (
			result string
			err    error
		)

		switch target {
		case FORMATJSON:
			result, err = format.JSON(newString)
		}

		if err == nil {
			textOutput.SetText(result)
		}
	}

	encDecObj := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Decoded"), textInput,
		widget.NewLabel("Encoded"), textOutput,
	)

	return encDecObj
}
