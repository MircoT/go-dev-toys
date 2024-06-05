//go:build app

package objects

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MircoT/go-dev-toys/pkg/encdec/base64"
	"github.com/MircoT/go-dev-toys/pkg/encdec/html"
	"github.com/MircoT/go-dev-toys/pkg/encdec/jwt"
	"github.com/MircoT/go-dev-toys/pkg/encdec/url"
)

const (
	encDecMinInputNumRows = 16
)

type encDecTarget int

const (
	ENCDECB64 encDecTarget = iota
	ENCDECHTML
	ENCDECURL
)

func MakeEncDec(target encDecTarget) *fyne.Container {
	textInput := widget.NewEntry()

	switch target {
	case ENCDECB64:
		textInput.SetPlaceHolder("source text")
	case ENCDECHTML:
		textInput.SetPlaceHolder("source HTML")
	case ENCDECURL:
		textInput.SetPlaceHolder("source URL")
	}

	textInput.MultiLine = true
	textInput.TextStyle.Monospace = true
	textInput.SetMinRowsVisible(encDecMinInputNumRows)
	textInput.Wrapping = fyne.TextWrapBreak

	textOutput := widget.NewEntry()

	switch target {
	case ENCDECB64:
		textInput.SetPlaceHolder("encoded text")
	case ENCDECHTML:
		textInput.SetPlaceHolder("encoded HTML")
	case ENCDECURL:
		textInput.SetPlaceHolder("encoded URL")
	}

	textOutput.MultiLine = true
	textOutput.TextStyle.Monospace = true
	textOutput.SetMinRowsVisible(encDecMinInputNumRows)
	textOutput.Wrapping = fyne.TextWrapBreak

	textInput.OnSubmitted = func(newString string) {
		var (
			result string
			err    error
		)

		switch target {
		case ENCDECB64:
			result, err = base64.Encode(newString)
		case ENCDECHTML:
			result, err = html.Encode(newString)
		case ENCDECURL:
			result, err = url.Encode(newString)
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
		case ENCDECB64:
			result, err = base64.Encode(newString)
		case ENCDECHTML:
			result, err = html.Encode(newString)
		case ENCDECURL:
			result, err = url.Encode(newString)
		}

		if err == nil {
			textOutput.SetText(result)
		}
	}

	textOutput.OnSubmitted = func(newString string) {
		var (
			result string
			err    error
		)

		switch target {
		case ENCDECB64:
			result, err = base64.Decode(newString)
		case ENCDECHTML:
			result, err = html.Decode(newString)
		case ENCDECURL:
			result, err = url.Decode(newString)
		}

		if err == nil {
			textInput.SetText(result)
		}
	}
	textOutput.OnChanged = func(newString string) {
		var (
			result string
			err    error
		)

		switch target {
		case ENCDECB64:
			result, err = base64.Decode(newString)
		case ENCDECHTML:
			result, err = html.Decode(newString)
		case ENCDECURL:
			result, err = url.Decode(newString)
		}

		if err == nil {
			textInput.SetText(result)
		}
	}

	encDecObj := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Decoded"), textInput,
		widget.NewLabel("Encoded"), textOutput,
	)

	return encDecObj
}

func MakeEncDecJWT() *fyne.Container {
	textInput := widget.NewEntry()

	textInput.SetPlaceHolder("JWT")
	textInput.SetMinRowsVisible(encDecMinInputNumRows)

	textInput.MultiLine = true
	textInput.TextStyle.Monospace = true
	textInput.Wrapping = fyne.TextWrapBreak

	textOutput := widget.NewEntry()

	textInput.SetPlaceHolder("")
	textOutput.SetMinRowsVisible(encDecMinInputNumRows + encDecMinInputNumRows)

	textOutput.MultiLine = true
	textOutput.TextStyle.Monospace = true
	textOutput.Wrapping = fyne.TextWrapBreak

	textInput.OnSubmitted = func(newString string) {
		var (
			result string
			err    error
		)

		result, err = jwt.Decode(newString)

		if err == nil {
			textOutput.SetText(result)
		}
	}
	textInput.OnChanged = func(newString string) {
		var (
			result string
			err    error
		)

		result, err = jwt.Decode(newString)

		if err == nil {
			textOutput.SetText(result)
		}
	}

	encDecObj := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Encoded"), textInput,
		widget.NewLabel("Decoded"), textOutput,
	)

	return encDecObj
}
