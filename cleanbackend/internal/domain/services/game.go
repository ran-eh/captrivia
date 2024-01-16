// internal/domain/services/game.go
package services

import (
	"sync"

	"github.com/ProlificLabs/captrivia/internal/application/ports"
	"github.com/ProlificLabs/captrivia/internal/domain/entities"
	"github.com/ProlificLabs/captrivia/internal/domain/errors"
)

// GameService encapsulates business logic for the trivia game.
type GameService struct {
	questionRepo ports.QuestionRepository
	gameStates   sync.Map // The game state for each player (playerID is the key)
}

// NewGameService creates a new GameService instance with the necessary dependencies.
func NewGameService(questionRepo ports.QuestionRepository) *GameService {
	return &GameService{
		questionRepo: questionRepo,
	}
}

// StartGame initializes and starts a new game for the given player.
func (gs *GameService) StartGame(playerID string) error {
	// Check if there is an ongoing game for the player
	if _, ok := gs.gameStates.Load(playerID); ok {
		return errors.ErrGameAlreadyInProgress
	}

	// Initialize the game state
	gameState := &entities.GameState{
		CurrentScore: 0,
		// Initialize other relevant fields (e.g., current question, remaining time, etc.)
	}
	gs.gameStates.Store(playerID, gameState)

	// Set up any other game initialization specifics such as timers if using a timer service
	// ...

	return nil
}

// ProcessAnswer checks if the player's answer to a question is correct, updates the score and returns the result.
func (gs *GameService) ProcessAnswer(playerID string, questionID string, selectedAnswer int) (bool, error) {
	question, err := gs.questionRepo.GetQuestionByID(questionID)
	if err != nil {
		return false, errors.ErrQuestionNotFound
	}

	correct := question.CorrectAnswerIndex == selectedAnswer
	gs.updateScore(playerID, correct, question.Difficulty)

	return correct, nil
}

// updateScore updates the player's score based on the correctness of the answer and question difficulty.
func (gs *GameService) updateScore(playerID string, correct bool, difficulty string) {
	gameStateInterface, ok := gs.gameStates.Load(playerID)
	if !ok {
		return // Perhaps log that there was an attempt to update score for non-existent game state
	}

	gameState := gameStateInterface.(*entities.GameState)
	if correct {
		// Assuming a simple rule where correct answers increase score by difficulty multiplier
		difficultyMultiplier := gs.calculateScoreMultiplier(difficulty)
		gameState.CurrentScore += difficultyMultiplier
	}
	// You may wish to apply a penalty for incorrect answers, although that's not in this example

	// Save any updated game state back to the map
	gs.gameStates.Store(playerID, gameState)
}

// calculateScoreMultiplier returns a score multiplier based on the difficulty.
func (gs *GameService) calculateScoreMultiplier(difficulty string) int {
	switch difficulty {
	case "easy":
		return 1
	case "medium":
		return 2
	case "hard":
		return 3
	default:
		return 1
	}
}

// EndGame concludes the game for the player, cleaning up the state.
func (gs *GameService) EndGame(playerID string) (finalScore int, err error) {
	gameStateInterface, ok := gs.gameStates.Load(playerID)
	if !ok {
		return 0, errors.ErrNoGameInProgress
	}
	gameState := gameStateInterface.(*entities.GameState)

	// Clean up the game state
	gs.gameStates.Delete(playerID)

	// Return the final score (or perform any other end game logic required)
	return gameState.CurrentScore, nil
}