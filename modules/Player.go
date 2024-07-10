package modules

// Player represents a player in the arena.
type Player struct {
	ID       int
	Name     string
	Health   int
	Strength int
	Attack   int
}

// NewPlayer creates a new Player instance.
func NewPlayer(id int, name string, health int, strength int, attack int) *Player {
	return &Player{
		ID:       id,
		Name:     name,
		Health:   health,
		Strength: strength,
		Attack:   attack,
	}
}
