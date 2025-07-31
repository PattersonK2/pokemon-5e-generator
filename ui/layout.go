package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PattersonK2/pokemon-5e-generator/core"
)

func MainLayout() fyne.CanvasObject {
	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Your Pokémon will appear here...")
	output.Wrapping = fyne.TextWrapWord

	randomBtn := widget.NewButton("🎲 Random Pokémon", func() {
		pokemon := core.GenerateRandomPokemon()
		buildPokemon := core.BuildPokemon(pokemon, 20)
		output.SetText(buildPokemon.ToString())
	})

	buildBtn := widget.NewButton("🛠 Build Your Own", func() {
		output.SetText("Build mode activated! (placeholder)")
	})

	header := widget.NewLabelWithStyle("Pokémon Builder & Randomizer", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	buttons := container.NewHBox(randomBtn, buildBtn)

	// Wrap buttons and header in a VBox
	top := container.NewVBox(header, buttons)

	// Return a Border layout: top is header + buttons, center is the growing output
	return container.NewBorder(top, nil, nil, nil, output)
}
