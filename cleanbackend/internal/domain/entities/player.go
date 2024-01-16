package entities

// Player represents a participant in the CapTrivia game.
type Player struct {
	ID    string // Unique identifier for the player, e.g., a UUID
	Name  string // Player's display name
	Score int    // Player's current score
}