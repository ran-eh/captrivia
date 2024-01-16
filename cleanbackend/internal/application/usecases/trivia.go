// internal/application/usecases/trivia.go
package usecases

import (
	"github.com/ProlificLabs/captrivia/internal/application/ports"
	"github.com/ProlificLabs/captrivia/internal/domain/entities"
	"github.com/ProlificLabs/captrivia/internal/domain/services"
)

// TriviaUseCase handles application logic and orchestrates the flow of the trivia game.
type TriviaUseCase struct {
	QuestionRepo ports.QuestionRepository
	GameService  *services.GameService
}

// NewTriviaUseCase creates a new TriviaUseCase with necessary dependencies.
func NewTriviaUseCase(questionRepo ports.QuestionRepository, gameService *services.GameService) *TriviaUseCase {
	return &TriviaUseCase{
		QuestionRepo: questionRepo,
		GameService:  gameService,
	}
}

// GetNextQuestion retrieves a new question for a player.
func (uc *TriviaUseCase) GetNextQuestion(playerID string) (*entities.Question, error) {
	// Implementation details go here.
	// E.g., Retrieve a random question from the QuestionRepo based on the player's progress.
	return &entities.Question{}, nil
}

// SubmitAnswer processes a player's answer and updates the game state.
func (uc *TriviaUseCase) SubmitAnswer(playerID string, questionID string, answerIndex int) (correct bool, updatedScore int, err error) {
	// Get the player's current state.
	player, err := uc.GameService.GetPlayer(playerID) // Not yet implemented
	if err != nil {
		return false, 0, err
	}

	// Retrieve the question.
	question, err := uc.QuestionRepo.GetQuestionByID(questionID)
	if err != nil {
		return false, 0, err
	}

	// Submit the player's answer using the domain service.
	err = uc.GameService.AnswerQuestion(player, question, answerIndex)
	if err != nil {
		return false, 0, err
	}

	// Return whether the answer was correct and the player's updated score.
	return answerIndex == question.Correct, player.Score, nil
}

// Implement additional use cases such as starting a new game, ending the current game, etc.