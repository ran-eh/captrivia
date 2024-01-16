// tests/e2e/game_test.go
package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ProlificLabs/captrivia/cmd/captrivia"
	"github.com/stretchr/testify/assert"
)

func TestGameFlow(t *testing.T) {
	// Initialize the application (this would normally be in TestMain or setup function).
	app, err := captrivia.Setup()
	assert.NoError(t, err)
	defer app.DB.Close()

	server := app.Server
	assert.NotNil(t, server)

	// Start a game for the player
	playerID := "test-player-id"
	req, err := http.NewRequest("POST", fmt.Sprintf("/players/%s/start", playerID), nil)
	assert.NoError(t, err)

	// Use a test http server.
	testServer := httptest.NewServer(app.Server.Handler)
	defer testServer.Close()

	client := testServer.Client()

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Retrieve the next question (let's assume this is how the API works).
	req, err = http.NewRequest("GET", fmt.Sprintf("/players/%s/question/next", playerID), nil)
	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var questionResp struct {
		QuestionID string   `json:"id"`
		Text       string   `json:"text"`
		Options    []string `json:"options"`
	}
	err = json.NewDecoder(resp.Body).Decode(&questionResp)
	assert.NoError(t, err)
	resp.Body.Close()

	// Submit an answer (again, assuming the business logic for the API route).
	answerIndex := 0 // Simplified, the answer index is hardcoded for this example.
	answerPayload := bytes.NewBuffer(nil)
	err = json.NewEncoder(answerPayload).Encode(map[string]interface{}{
		"answerIndex": answerIndex,
	})
	assert.NoError(t, err)

	req, err = http.NewRequest("POST", fmt.Sprintf("/players/%s/question/%s/answer", playerID, questionResp.QuestionID), answerPayload)
	assert.NoError(t, err)

	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var answerResp struct {
		Correct      bool `json:"correct"`
		UpdatedScore int  `json:"updatedScore"`
	}
	err = json.NewDecoder(resp.Body).Decode(&answerResp)
	assert.NoError(t, err)
	resp.Body.Close()

	// Add more assertions depending on how the game should behave, which may include checking the correctness of the answer, the score update, etc.

	// Continue the game flow by retrieving the next question or finishing the game.

	// Actual game tests would check various paths through the game, the persistence of game state between requests, and error handling.
}

// TestEndToEndGameFlow is the wrapper to run the game flow tests.
func TestEndToEndGameFlow(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping end-to-end test in short mode.")
	}

	t.Run("TestGameFlow", TestGameFlow)
	// Add more test scenarios as needed.
}
