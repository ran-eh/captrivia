// internal/interfaces/storage/question_repo.go
package storage

import (
	"github.com/ProlificLabs/captrivia/internal/application/ports"
	"github.com/ProlificLabs/captrivia/internal/domain/entities"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type QuestionRepo struct {
	DB *sqlx.DB
}

func NewQuestionRepo(db *sqlx.DB) ports.QuestionRepository {
	return &QuestionRepo{DB: db}
}

func (repo *QuestionRepo) GetQuestionByID(questionID string) (*entities.Question, error) {
	var question entities.Question
	query := `SELECT id, text, options, correct, difficulty FROM questions WHERE id=$1;`
	err := repo.DB.Get(&question, query, questionID)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (repo *QuestionRepo) SaveQuestion(question *entities.Question) error {
	query := `INSERT INTO questions (id, text, options, correct, difficulty) VALUES ($1, $2, $3, $4, $5);`
	_, err := repo.DB.Exec(query, question.ID, question.Text, pq.Array(question.Options), question.Correct, question.Difficulty)
	if err != nil {
		return err
	}
	return nil
}

// Additional methods such as UpdateQuestion or FindByDifficulty could be added here.