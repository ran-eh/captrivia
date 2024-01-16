// internal/interfaces/storage/player_repo.go
package storage

import (
	"github.com/ProlificLabs/captrivia/internal/application/ports"
	"github.com/ProlificLabs/captrivia/internal/domain/entities"
	"github.com/jmoiron/sqlx"
)

type PlayerRepo struct {
	DB *sqlx.DB
}

func NewPlayerRepo(db *sqlx.DB) ports.PlayerRepository {
	return &PlayerRepo{DB: db}
}

func (repo *PlayerRepo) GetPlayerByID(playerID string) (*entities.Player, error) {
	var player entities.Player
	query := `SELECT id, name, score FROM players WHERE id=$1;`
	err := repo.DB.Get(&player, query, playerID)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (repo *PlayerRepo) SavePlayer(player *entities.Player) error {
	query := `INSERT INTO players (id, name, score) VALUES ($1, $2, $3);`
	_, err := repo.DB.Exec(query, player.ID, player.Name, player.Score)
	if err != nil {
		return err
	}
	return nil
}

// Additional methods such as UpdatePlayer could be added here.