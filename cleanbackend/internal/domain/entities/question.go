package entities

// Question represents a single trivia question with its possible answers.
type Question struct {
	ID       string   // Unique identifier for the question
	Text     string   // The question text
	Options  []string // Available answers
	CorrectAnswerIndex  int      // Index of the correct answer within the Options slice
	Difficulty string // Difficulty level of the question, e.g., "easy", "medium", "hard"
}