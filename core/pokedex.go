package core

import (
	"log"
	"github.com/PattersonK2/pokemon-5e-generator/data"
	"github.com/PattersonK2/pokemon-5e-generator/models"
)

func GetPokedex() []models.Pokemon {
	pokedex, err := data.LoadPokemonData()
	if err != nil {
		log.Fatalf("Failed to load Pokémon data: %v", err)
	}

	if len(pokedex) == 0 {
		log.Fatal("Pokedex is empty")
	}

	return pokedex
}