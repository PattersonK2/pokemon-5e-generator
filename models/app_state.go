package models

type AppState int

const (
	StateInit AppState = iota
	StateRandom
	StateBuild
)
