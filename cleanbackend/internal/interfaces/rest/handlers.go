// internal/interfaces/rest/handlers.go
package rest

import (
	"encoding/json"
	"net/http"

	"github.com/ProlificLabs/captrivia/internal/application/usecases"
	"github.com/gorilla/mux"
)

// TriviaHandler represents the HTTP handler for trivia game operations.
type TriviaHandler struct {
	TriviaUseCase *usecases.TriviaUseCase // Interface for the Trivia use case
}

// NewTriviaHandler creates a new handler for trivia related endpoints.
func NewTriviaHandler(triviaUseCase *usecases.TriviaUseCase) *TriviaHandler {
	return &TriviaHandler{
		TriviaUseCase: triviaUseCase,
	}
}

// GetNextQuestion is an HTTP handler for retrieving the next trivia question.
func (h *TriviaHandler) GetNextQuestion(w http.ResponseWriter, r *http.Request) {
	playerID := mux.Vars(r)["playerID"]
	question, err := h.TriviaUseCase.GetNextQuestion(playerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(question); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SubmitAnswer is an HTTP handler for submitting an answer to a trivia question.
func (h *TriviaHandler) SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID := vars["playerID"]
	questionID := vars["questionID"]

	var req struct {
		AnswerIndex int `json:"answerIndex"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	correct, updatedScore, err := h.TriviaUseCase.SubmitAnswer(playerID, questionID, req.AnswerIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		Correct      bool `json:"correct"`
		UpdatedScore int  `json:"updatedScore"`
	}{
		Correct:      correct,
		UpdatedScore: updatedScore,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RegisterTriviaRoutes sets up the routes for trivia game operations on the given router.
func RegisterTriviaRoutes(router *mux.Router, triviaUseCase *usecases.TriviaUseCase) {
	handler := NewTriviaHandler(triviaUseCase)

	router.HandleFunc("/player/{playerID}/question/next", handler.GetNextQuestion).Methods(http.MethodGet)
	router.HandleFunc("/player/{playerID}/question/{questionID}/answer", handler.SubmitAnswer).Methods(http.MethodPost)
}