package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainLayout() fyne.CanvasObject {
	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Your Pokémon will appear here...")

	randomBtn := widget.NewButton("🎲 Random Pokémon", func() {
		output.SetText("Random Pokémon generated! (placeholder)")
	})

	buildBtn := widget.NewButton("🛠 Build Your Own", func() {
		output.SetText("Build mode activated! (placeholder)")
	})

	buttons := container.NewHBox(randomBtn, buildBtn)

	layout := container.NewVBox(
		widget.NewLabelWithStyle("Pokémon Builder & Randomizer", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		buttons,
		output,
	)

	return layout
}
