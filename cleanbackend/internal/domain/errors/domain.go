package errors

import "errors"

// Predefined domain errors
var (
	ErrQuestionNotFound      = errors.New("the question was not found")
	ErrInvalidAnswer         = errors.New("the provided answer is invalid")
	ErrPlayerAlreadyExists   = errors.New("player already exists")
	ErrPlayerNotFound        = errors.New("player not found")
	ErrGameAlreadyInProgress = errors.New("there is already a game in progress for the player")
	ErrNoGameInProgress      = errors.New("there is no game in progress to perform this operation")
	// Add more specific domain errors as needed
)