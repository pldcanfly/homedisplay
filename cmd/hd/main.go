package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(container.NewVBox(
		widget.NewLabel("Date&Time"),
		widget.NewLabel("Roomtemperature"),
		widget.NewLabel("Outside Temperature"),
		widget.NewLabel("Ground Humidity"),
		widget.NewLabel("Washing Machine"),
		widget.NewLabel("Irrigation"),
		widget.NewLabel("Klimate"),
		widget.NewLabel("Currently Playing + Controls"),
	))

	w.ShowAndRun()
}
