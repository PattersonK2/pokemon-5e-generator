package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PattersonK2/pokemon-5e-generator/core"
	"github.com/PattersonK2/pokemon-5e-generator/models"
)

func MainLayout() fyne.CanvasObject {
	allPokemon := core.GetPokedex()
	pokemonNames := make([]string, len(allPokemon))
	for i, p := range allPokemon {
		pokemonNames[i] = p.Name
	}

	currentState := models.StateRandom

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Your Pokémon will appear here...")
	output.Wrapping = fyne.TextWrapWord

	selectEntry := widget.NewSelectEntry(pokemonNames)
	selectEntry.SetPlaceHolder("Search for a Pokémon...")
	selectEntry.Hide()

	content := container.NewVBox()

	setState := func(state models.AppState) {
		currentState = state

		switch state {
		case models.StateRandom:
			selectEntry.Hide()
			pokemon := core.GenerateRandomPokemon()
			buildPokemon := core.BuildPokemon(pokemon, 20)
			output.SetText(buildPokemon.ToString())

		case models.StateBuild:
			selectEntry.Show()
			output.SetText("Build mode activated! Select a Pokémon from the dropdown.")

		case models.StateInit:
		default:
			output.SetText("Unknown state")
		}

		content.Refresh()
	}

	selectEntry.OnChanged = func(name string) {
		if currentState == models.StateBuild {
			for _, p := range allPokemon {
				if p.Name == name {
					output.SetText(p.ToString())
					break
				}
			}
		}
	}

	randomBtn := widget.NewButton("🎲 Random Pokémon", func() {
		setState(models.StateRandom)
	})

	buildBtn := widget.NewButton("🛠 Build Your Own", func() {
		setState(models.StateBuild)
	})

	header := widget.NewLabelWithStyle("Pokémon Builder & Randomizer", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	buttons := container.NewHBox(randomBtn, buildBtn)

	top := container.NewVBox(
		header,
		buttons,
		selectEntry,
	)

	return container.NewBorder(top, nil, nil, nil, output)
}
