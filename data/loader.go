package data

import (
	_ "embed"
	"encoding/json"
	"github.com/PattersonK2/pokemon-5e-generator/models"
)

//go:embed files/pokemon.json
var pokemonJSON []byte

func LoadPokemonData() ([]models.Pokemon, error) {
	var pokedex []models.Pokemon
	err := json.Unmarshal(pokemonJSON, &pokedex)
	return pokedex, err
}
