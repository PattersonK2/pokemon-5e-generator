package main

import (
	"github.com/PattersonK2/pokemon-5e-generator/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// Create the app
	myApp := app.New()
	myWindow := myApp.NewWindow("Pokémon Builder & Randomizer")

	// Set the UI layout
	myWindow.SetContent(ui.MainLayout())

	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.ShowAndRun()
}
