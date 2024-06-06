//go:build app

package objects

import (
	"fmt"
	"log/slog"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MircoT/go-dev-toys/pkg/dectext"
	"github.com/MircoT/go-dev-toys/pkg/dectext/gz"
	"github.com/MircoT/go-dev-toys/pkg/dectext/zip"
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

			targetFile := uri.URI().Path()

			typeStr, err := dectext.GetCompressType(targetFile)
			if err != nil {
				slog.Error(fmt.Errorf("cannot get file type of '%s': %w ", targetFile, err).Error())
			}

			data, err := os.ReadFile(targetFile)
			if err != nil {
				slog.Error(err.Error())
			}

			var result string

			switch typeStr {
			case "zstd", "zst":
				result, err = zstd.Decompress(data)
			case "zip":
				result, err = zip.Decompress(data)
			case "gz", "gzip":
				result, err = gz.Decompress(data)
			default:
				slog.Error(fmt.Errorf("%s is not a valid format", typeStr).Error())
			}

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
