//go:build app

package objects

import (
	"log/slog"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MircoT/go-dev-toys/pkg/dectext/zstd"
)

const (
	ctextMinInputNumRows = 16
)

func MakeCText(parent fyne.Window) *fyne.Container {
	textOutput := widget.NewEntry()
	textOutput.SetMinRowsVisible(ctextMinInputNumRows + ctextMinInputNumRows)

	textOutput.MultiLine = true
	textOutput.TextStyle.Monospace = true
	textOutput.Wrapping = fyne.TextWrapBreak

	openFile := widget.NewButton("Select file", func() {
		dialog := dialog.NewFileOpen(func(uri fyne.URIReadCloser, err error) {
			if err != nil {
				slog.Error(err.Error())
			}

			data, err := os.ReadFile(uri.URI().Path())
			if err != nil {
				slog.Error(err.Error())
			}

			result, err := zstd.Decompress(data)

			if err == nil {
				textOutput.SetText(result)
			} else {
				slog.Error(err.Error())
			}
		}, parent)

		dialog.Resize(fyne.Size{800, 600})

		dialog.Show()
	})

	encDecObj := container.New(layout.NewVBoxLayout(),
		openFile, textOutput,
	)

	return encDecObj
}
