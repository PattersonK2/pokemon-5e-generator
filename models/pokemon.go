package models

import (
	"fmt"
	"strings"
)

type Pokemon struct {
	Name            string     `json:"name"`
	Types           []string   `json:"types"`
	Size            string     `json:"size"`
	SR              string     `json:"sr"`
	MinLevel        int        `json:"minLevel"`
	EggGroups       []string   `json:"eggGroups"`
	GenderRateM     int        `json:"genderRateM"`
	GenderRateF     int        `json:"genderRateF"`
	EvolutionStage  string     `json:"evolutionStage"`
	AC              int        `json:"ac"`
	MinHP           int        `json:"minHP"`
	HitDice         string     `json:"hitDice"`
	Speeds          []string   `json:"speeds"`
	Stats           BaseStats  `json:"stats"`
	Skills          []string   `json:"skills"`
	SavingThrows    []string   `json:"savingThrows"`
	Vulnerabilities []string   `json:"vulnerabilities"`
	Resistances     []string   `json:"resistances"`
	Abilities       []string   `json:"abilities"`
	HiddenAbility   string     `json:"hiddenAbility"`
	LevelMoves      LevelMoves `json:"levelMoves"`
	TMs             []int      `json:"tms"`
	EggMoves        []string   `json:"eggMoves"`
}

type BaseStats struct {
	STR int `json:"str"`
	DEX int `json:"dex"`
	CON int `json:"con"`
	INT int `json:"int"`
	WIS int `json:"wis"`
	CHA int `json:"cha"`
}

type LevelMoves struct {
	StartingMoves []string `json:"startingMoves"`
	Level2        []string `json:"level2"`
	Level6        []string `json:"level6"`
	Level10       []string `json:"level10"`
	Level14       []string `json:"level14"`
	Level18       []string `json:"level18"`
}

func (p Pokemon) ToString() string {
	return fmt.Sprintf(
		`Name: %s
Types: %s
Size: %s
SR: %s
Min Level: %d
Egg Groups: %s
Gender Ratio: ♂ %d%% / ♀ %d%%
Evolution Stage: %s
AC: %d
Min HP: %d
Hit Dice: %s
Speeds: %s

Stats:
  STR: %d  DEX: %d  CON: %d
  INT: %d  WIS: %d  CHA: %d

Skills: %s
Saving Throws: %s
Vulnerabilities: %s
Resistances: %s
Abilities: %s
Hidden Ability: %s

Level-Up Moves:
	Starting: %s
	Level 2:  %s
	Level 6:  %s
	Level 10: %s
	Level 14: %s
	Level 18: %s

TMs: %v
Egg Moves: %s
`,
		p.Name,
		strings.Join(p.Types, ", "),
		p.Size,
		p.SR,
		p.MinLevel,
		strings.Join(p.EggGroups, ", "),
		p.GenderRateM, p.GenderRateF,
		p.EvolutionStage,
		p.AC,
		p.MinHP,
		p.HitDice,
		strings.Join(p.Speeds, ", "),

		p.Stats.STR, p.Stats.DEX, p.Stats.CON,
		p.Stats.INT, p.Stats.WIS, p.Stats.CHA,

		strings.Join(p.Skills, ", "),
		strings.Join(p.SavingThrows, ", "),
		strings.Join(p.Vulnerabilities, ", "),
		strings.Join(p.Resistances, ", "),
		strings.Join(p.Abilities, ", "),
		p.HiddenAbility,

		strings.Join(p.LevelMoves.StartingMoves, ", "),
		strings.Join(p.LevelMoves.Level2, ", "),
		strings.Join(p.LevelMoves.Level6, ", "),
		strings.Join(p.LevelMoves.Level10, ", "),
		strings.Join(p.LevelMoves.Level14, ", "),
		strings.Join(p.LevelMoves.Level18, ", "),

		p.TMs,
		strings.Join(p.EggMoves, ", "),
	)
}
