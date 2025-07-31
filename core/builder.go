package core

import (
	"github.com/PattersonK2/pokemon-5e-generator/models"
)

// BuildPokemon takes a base Pokemon and applies generation logic.
// For example: level scaling, assigning randomized traits, etc.
func BuildPokemon(base models.Pokemon, level int) models.Pokemon {
	// TODO: Implement logic to modify the base Pokémon.
	// This could involve:
	// - Selecting level-appropriate moves
	// - Adjusting stats
	// - Applying evolution logic
	// - Randomizing nature, ability, etc.

	// For now, just return the base
	return base
}
