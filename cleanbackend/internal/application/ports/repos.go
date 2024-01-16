// internal/application/ports/repos.go
package ports

import (
	"github.com/ProlificLabs/captrivia/internal/domain/entities"
)

// QuestionRepository defines the interface for question data access.
type QuestionRepository interface {
	GetQuestionByID(questionID string) (*entities.Question, error)
	SaveQuestion(question *entities.Question) error
	GetNextQuestion(playerID string) (*entities.Question, error)
	// Add more methods as necessary, such as FindByDifficulty or a method to get a random question.
}

// PlayerRepository defines the interface for player data access.
type PlayerRepository interface {
	GetPlayerByID(playerID string) (*entities.Player, error)
	SavePlayer(player *entities.Player) error
	// Add more methods as necessary.
}

// ScoreRepository could be added to handle scores separately,
// which would be an interface with methods like SaveScore, GetTopScores, etc.