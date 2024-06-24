package objects

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MircoT/go-dev-toys/pkg/convert"
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

func MakeBytesConverter() *fyne.Container {
	bInput := widget.NewEntry()
	bInput.TextStyle.Monospace = true
	bInput.SetPlaceHolder("b")

	kbInput := widget.NewEntry()
	kbInput.TextStyle.Monospace = true
	kbInput.SetPlaceHolder("kb")

	mbInput := widget.NewEntry()
	mbInput.TextStyle.Monospace = true
	mbInput.SetPlaceHolder("mb")

	gbInput := widget.NewEntry()
	gbInput.TextStyle.Monospace = true
	gbInput.SetPlaceHolder("gb")

	tbInput := widget.NewEntry()
	tbInput.TextStyle.Monospace = true
	tbInput.SetPlaceHolder("tb")

	pbInput := widget.NewEntry()
	pbInput.TextStyle.Monospace = true
	pbInput.SetPlaceHolder("pb")

	kibInput := widget.NewEntry()
	kibInput.TextStyle.Monospace = true
	kibInput.SetPlaceHolder("kib")

	mibInput := widget.NewEntry()
	mibInput.TextStyle.Monospace = true
	mibInput.SetPlaceHolder("mib")

	gibInput := widget.NewEntry()
	gibInput.TextStyle.Monospace = true
	gibInput.SetPlaceHolder("gib")

	tibInput := widget.NewEntry()
	tibInput.TextStyle.Monospace = true
	tibInput.SetPlaceHolder("tib")

	pibInput := widget.NewEntry()
	pibInput.TextStyle.Monospace = true
	pibInput.SetPlaceHolder("pib")

	targetUnit := "kib"

	convertBtn := widget.NewButton("Convert", func() {
		curString := "0"

		switch targetUnit {
		case "b":
			curString = bInput.Text
		case "kb":
			curString = kbInput.Text
		case "mb":
			curString = mbInput.Text
		case "gb":
			curString = gbInput.Text
		case "tb":
			curString = tbInput.Text
		case "pb":
			curString = pbInput.Text
		case "kib":
			curString = kibInput.Text
		case "mib":
			curString = mibInput.Text
		case "gib":
			curString = gibInput.Text
		case "tib":
			curString = tibInput.Text
		case "pib":
			curString = pibInput.Text
		}

		uintVal, err := strconv.ParseUint(curString, 10, 0)
		if err == nil {
			values, err := convert.Bytes(uintVal, targetUnit)
			if err == nil {
				kbInput.SetText(fmt.Sprintf("%0.2f", values.KB))
				mbInput.SetText(fmt.Sprintf("%0.2f", values.MB))
				gbInput.SetText(fmt.Sprintf("%0.2f", values.GB))
				tbInput.SetText(fmt.Sprintf("%0.2f", values.TB))
				pbInput.SetText(fmt.Sprintf("%0.2f", values.PB))
				kibInput.SetText(fmt.Sprintf("%0.2f", values.KiB))
				mibInput.SetText(fmt.Sprintf("%0.2f", values.MiB))
				gibInput.SetText(fmt.Sprintf("%0.2f", values.GiB))
				tibInput.SetText(fmt.Sprintf("%0.2f", values.TiB))
				pibInput.SetText(fmt.Sprintf("%0.2f", values.PiB))
			}
		}
	})

	bInput.OnChanged = func(newString string) {
		targetUnit = "b"
	}

	kbInput.OnChanged = func(newString string) {
		targetUnit = "kb"
	}
	mbInput.OnChanged = func(newString string) {
		targetUnit = "mb"
	}
	gbInput.OnChanged = func(newString string) {
		targetUnit = "gb"
	}
	tbInput.OnChanged = func(newString string) {
		targetUnit = "tb"
	}
	pbInput.OnChanged = func(newString string) {
		targetUnit = "pb"
	}

	kibInput.OnChanged = func(newString string) {
		targetUnit = "kib"
	}
	mibInput.OnChanged = func(newString string) {
		targetUnit = "mib"
	}
	gibInput.OnChanged = func(newString string) {
		targetUnit = "gib"
	}
	tibInput.OnChanged = func(newString string) {
		targetUnit = "tib"
	}
	pibInput.OnChanged = func(newString string) {
		targetUnit = "pib"
	}

	formatBytesObj := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Bytes"), bInput,
		container.NewGridWithColumns(2,
			widget.NewLabel("KB"), widget.NewLabel("KiB"),
			kbInput, kibInput,
			widget.NewLabel("MB"), widget.NewLabel("MiB"),
			mbInput, mibInput,
			widget.NewLabel("GB"), widget.NewLabel("GiB"),
			gbInput, gibInput,
			widget.NewLabel("TB"), widget.NewLabel("TiB"),
			tbInput, tibInput,
			widget.NewLabel("PB"), widget.NewLabel("PiB"),
			pbInput, pibInput,
		),
		convertBtn,
	)

	return formatBytesObj
}
