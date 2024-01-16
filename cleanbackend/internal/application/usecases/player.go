// internal/application/usecases/player.go
package usecases

import (
	"github.com/ProlificLabs/captrivia/internal/application/ports"
	"github.com/ProlificLabs/captrivia/internal/domain/entities"
	domainErrors "github.com/ProlificLabs/captrivia/internal/domain/errors"
	"github.com/google/uuid"
)

// PlayerUseCase represents the application's use case for player operations.
type PlayerUseCase struct {
	PlayerRepo ports.PlayerRepository
}

// NewPlayerUseCase creates a new instance of PlayerUseCase.
func NewPlayerUseCase(playerRepo ports.PlayerRepository) *PlayerUseCase {
	return &PlayerUseCase{
		PlayerRepo: playerRepo,
	}
}

// CreatePlayer creates a new player in the game.
func (uc *PlayerUseCase) CreatePlayer(name string) (*entities.Player, error) {
	newPlayerID := uuid.New().String()
	newPlayer := &entities.Player{
		ID:    newPlayerID,
		Name:  name,
		Score: 0,
	}
	err := uc.PlayerRepo.SavePlayer(newPlayer)
	if err != nil {
		return nil, err
	}
	return newPlayer, nil
}

// UpdatePlayerScore updates the score for a given player.
func (uc *PlayerUseCase) UpdatePlayerScore(playerID string, score int) (*entities.Player, error) {
	player, err := uc.PlayerRepo.GetPlayerByID(playerID)
	if err != nil {
		return nil, err
	}

	player.Score = score
	err = uc.PlayerRepo.SavePlayer(player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

// GetPlayer retrieves a player by their ID.
func (uc *PlayerUseCase) GetPlayer(playerID string) (*entities.Player, error) {
	player, err := uc.PlayerRepo.GetPlayerByID(playerID)
	if err != nil {
		return nil, err
	}
	return player, nil
}

// DeletePlayer deletes a player from the game.
func (uc *PlayerUseCase) DeletePlayer(playerID string) error {
	player, err := uc.GetPlayer(playerID)
	if err != nil {
		return err
	}

	if player == nil {
		return domainErrors.ErrPlayerNotFound
	}

	// Assuming PlayerRepo has a method DeletePlayerByID
	err = uc.PlayerRepo.DeletePlayerByID(playerID)
	if err != nil {
		return err
	}

	return nil
}

// Additional player-related use cases would be added here as needed.