//go:build app

package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/MircoT/go-dev-toys/pkg/app/objects"
)

const (
	appName = "GoDevToys"
)

var (
	sections    = []string{"Compressed text", "Encode/Decode", "Format", "Convert"}
	subSections = map[string][]string{
		"Encode/Decode":   {"Base64", "HTML", "URL", "JWT"},
		"Compressed text": {"Zip", "Gz", "Zstd"},
		"Format":          {"JSON"},
		"Convert":         {"Numbers"},
	}
)

func Run() error {
	a := app.New()
	w := a.NewWindow(appName)

	tabSections := container.NewAppTabs()
	tabSubSections := make(map[string]*container.AppTabs)

	for subSection, tabs := range subSections {
		tabSubSections[subSection] = container.NewAppTabs()
		tabSubSections[subSection].SetTabLocation(container.TabLocationLeading)

		for _, tab := range tabs {
			switch tab {
			case "Base64":
				tabSubSections[subSection].Append(
					container.NewTabItem(tab, objects.MakeEncDec(objects.ENCDECB64)),
				)
			case "HTML":
				tabSubSections[subSection].Append(
					container.NewTabItem(tab, objects.MakeEncDec(objects.ENCDECHTML)),
				)
			case "URL":
				tabSubSections[subSection].Append(
					container.NewTabItem(tab, objects.MakeEncDec(objects.ENCDECURL)),
				)
			case "JWT":
				tabSubSections[subSection].Append(
					container.NewTabItem(tab, objects.MakeEncDecJWT()),
				)
			case "Zip", "Gz", "Zstd":
				tabSubSections[subSection].Append(
					container.NewTabItem(tab, objects.MakeCText(w)),
				)
			case "JSON":
				tabSubSections[subSection].Append(
					container.NewTabItem(tab, objects.MakeFormatText(objects.FORMATJSON)),
				)
			case "Numbers":
				tabSubSections[subSection].Append(
					container.NewTabItem(tab, objects.MakeNumberFormat()),
				)
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
