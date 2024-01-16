// internal/application/ports/dtos.go
package ports

// QuestionDTO denotes the data transfer object for a trivia question.
type QuestionDTO struct {
	ID          string   `json:"id"`
	Text        string   `json:"text"`
	Options     []string `json:"options"`
	Difficulty  string   `json:"difficulty"` // Could be an enum later
}

// AnswerSubmissionDTO encapsulates a player's answer submission to a question.
type AnswerSubmissionDTO struct {
	PlayerID    string `json:"playerId"`
	QuestionID  string `json:"questionId"`
	AnswerIndex int    `json:"answerIndex"`
}

// TriviaGameResultDTO represents the result sent back after submitting an answer.
type TriviaGameResultDTO struct {
	Correct      bool `json:"correct"`
	UpdatedScore int  `json:"updatedScore"`
}

// Additional DTOs can be added here for other operations, such as starting a game, tracking progress, etc.