// tests/integration/api_test.go
package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ProlificLabs/captrivia/cmd/captrivia"

	"go.uber.org/zap"
)

// setupTestApp initializes the application and the test server.
func setupTestApp() (*captrivia.App, *httptest.Server, *http.Client) {
	app, err := captrivia.Setup()
	if err != nil {
		zap.L().Fatal("Failed to set up the application for testing.", zap.Error(err))
	}

	// Create a test server.
	server := httptest.NewServer(app.Server.Handler)

	// Create a client with the test server URL.
	client := server.Client()

	return app, server, client
}

func TestMain(m *testing.M) {
	// Setup the application and the test server.
	var err error
	app, server, client = setupTestApp()
	defer server.Close()
	defer app.DB.Close()

	// Run the tests.
	exitCode := m.Run()

	// Perform any cleanup and exit.
	server.Close()
	app.DB.Close()
	os.Exit(exitCode)
}

func TestSubmitAnswerEndpoint(t *testing.T) {
	// Assume 'playerID' and 'questionID' are valid identifiers within your test fixtures.
	playerID, questionID := "test-player-id", "test-question-id"

	// Create JSON payload for the request.
	payload := map[string]interface{}{
		"answerIndex": 1,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}
	
	// Build the request URL.
	reqURL := fmt.Sprintf("%s/player/%s/question/%s/answer", server.URL, playerID, questionID)
	
	// Create the POST request.
	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Set required headers.
	req.Header.Set("Content-Type", "application/json")

	// Send the request.
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read and unmarshal response body.
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var result struct {
		Correct      bool `json:"correct"`
		UpdatedScore int  `json:"updatedScore"`
	}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Test assertions.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got: %d", resp.StatusCode)
	}

	// Here you will need to implement further assertions depending on your application's logic.
	// For example, was the answer correct? If so, the response should indicate it.
	// Was the score updated as expected?

	// Below are just example assertions, please adjust them according to your actual logic.
	expectedCorrect := false
	if result.Correct != expectedCorrect {
		t.Errorf("Expected result.Correct to be %v, got: %v", expectedCorrect, result.Correct)
	}

	expectedScore := 0
	if result.UpdatedScore != expectedScore {
		t.Errorf("Expected result.UpdatedScore to be %d, got: %d", expectedScore, result.UpdatedScore)
	}
}