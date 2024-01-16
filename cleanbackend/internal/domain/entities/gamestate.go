package entities

// GameState holds the information of a game in progress for a particular player.
type GameState struct {
	PlayerID     string // Link to the player entity
	CurrentScore int    // Current score in the game
	// Add additional fields needed to represent the game's state like CurrentQuestionID, TimeRemaining, etc.
}