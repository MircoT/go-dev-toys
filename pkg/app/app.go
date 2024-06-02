package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MircoT/go-dev-toys/pkg/encdec/base64"
)

const (
	appName = "GoDevToys"
)

var (
	sections    = []string{"Encode/Decode"}
	subSections = map[string][]string{
		"Encode/Decode": {"Base64", "HTML", "URL"},
	}
)

func Run() error {
	a := app.New()
	w := a.NewWindow(appName)

	// var data = []string{"a", "string", "list"}
	// list := widget.NewList(
	// 	func() int {
	// 		return len(data)
	// 	},
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("template")
	// 	},
	// 	func(i widget.ListItemID, o fyne.CanvasObject) {
	// 		o.(*widget.Label).SetText(data[i])
	// 	})

	b64Input := widget.NewEntry()
	b64Input.SetPlaceHolder("source text")
	b64Input.MultiLine = true
	b64Input.Wrapping = fyne.TextWrapBreak

	b64Output := widget.NewEntry()
	b64Output.SetPlaceHolder("encoded text")
	b64Output.MultiLine = true
	b64Output.Wrapping = fyne.TextWrapBreak

	b64Input.OnSubmitted = func(newString string) {
		result, _ := base64.Encode(newString)
		b64Output.SetText(result)
	}
	b64Input.OnChanged = func(newString string) {
		result, _ := base64.Encode(newString)
		b64Output.SetText(result)
	}

	b64Output.OnSubmitted = func(newString string) {
		result, _ := base64.Decode(newString)
		b64Input.SetText(result)
	}
	b64Output.OnChanged = func(newString string) {
		result, _ := base64.Decode(newString)
		b64Input.SetText(result)
	}

	encDecB64 := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Decoded"), b64Input,
		widget.NewLabel("Encoded"), b64Output,
	)

	tabSections := container.NewAppTabs()
	tabSubSections := make(map[string]*container.AppTabs)

	for subSection, tabs := range subSections {
		tabSubSections[subSection] = container.NewAppTabs()
		tabSubSections[subSection].SetTabLocation(container.TabLocationLeading)

		for _, tab := range tabs {
			switch tab {
			case "Base64":
				tabSubSections[subSection].Append(container.NewTabItem(tab, encDecB64))
			default:
				tabSubSections[subSection].Append(container.NewTabItem(tab, widget.NewLabel(tab)))
			}
		}
	}

	// spew.Dump(tabSubSections)

	for _, section := range sections {
		tabSections.Append(
			container.NewTabItem(section, tabSubSections[section]),
		)
	}

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabSections.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabSections)

	w.Resize(fyne.NewSize(800, 600))

	w.ShowAndRun()

	return nil
}
