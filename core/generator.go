package core

import (
	"math/rand"
	"time"
	"log"

	"github.com/PattersonK2/pokemon-5e-generator/models"
	"github.com/PattersonK2/pokemon-5e-generator/data"
)

// Create a package-level random source and rand instance
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
func RandomInt(min, max int) int {
	return rnd.Intn(max-min+1) + min
}

func GenerateRandomPokemon() models.Pokemon {
	pokedex, err := data.LoadPokemonData()
	if err != nil {
		log.Fatalf("Failed to load Pokémon data: %v", err)
	}

	if len(pokedex) == 0 {
		log.Fatal("Pokedex is empty")
	}

	first := pokedex[0]
	//take random pokemon
	//randomPokemon := pokedex[RandomInt(0, len(pokedex-1))]
	return first
}
