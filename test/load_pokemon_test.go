package test

import (
	"testing"

	"github.com/PattersonK2/pokemon-5e-generator/data"
)

func TestLoadPokemonData(t *testing.T) {
	pokedex, err := data.LoadPokemonData()
	if err != nil {
		t.Fatalf("LoadPokemonData failed: %v", err)
	}

	if len(pokedex) == 0 {
		t.Fatal("Pokedex is empty; expected at least one Pokémon")
	}

	// Basic check on first entry
	first := pokedex[0]
	if first.Name == "" {
		t.Error("First Pokémon has no name")
	}
	if len(first.Types) == 0 {
		t.Error("First Pokémon has no types")
	}
	if first.Stats.STR == 0 {
		t.Error("First Pokémon has uninitialized stats")
	}
}
