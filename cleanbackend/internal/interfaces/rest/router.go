// internal/interfaces/rest/router.go
package rest

import (
	"net/http"

	"github.com/ProlificLabs/captrivia/internal/application/usecases"
	"github.com/ProlificLabs/captrivia/internal/interfaces/storage"
	"github.com/gorilla/mux"
)

// NewRouter creates an HTTP router and registers the application's routes.
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Initialize our repositories and use cases.
	playerRepo := storage.NewPlayerRepository(/* pass in any required dependencies such as a database connection */)
	questionRepo := storage.NewQuestionRepository(/* pass in any required dependencies */)
	gameService := usecases.NewGameService(questionRepo)

	triviaUseCase := usecases.NewTriviaUseCase(questionRepo, gameService)
	playerUseCase := usecases.NewPlayerUseCase(playerRepo)

	// Player Routes
	playerHandler := NewPlayerHandler(playerUseCase)
	r.HandleFunc("/players", playerHandler.CreatePlayer).Methods(http.MethodPost)
	r.HandleFunc("/players/{playerID}", playerHandler.GetPlayer).Methods(http.MethodGet)
	r.HandleFunc("/players/{playerID}", playerHandler.UpdatePlayer).Methods(http.MethodPut)
	r.HandleFunc("/players/{playerID}", playerHandler.DeletePlayer).Methods(http.MethodDelete)

	// Trivia Game Routes
	triviaHandler := NewTriviaHandler(triviaUseCase)
	r.HandleFunc("/games/{playerID}/start", triviaHandler.StartGame).Methods(http.MethodPost)
	r.HandleFunc("/games/{playerID}/end", triviaHandler.EndGame).Methods(http.MethodPost)
	r.HandleFunc("/games/{playerID}/next", triviaHandler.GetNextQuestion).Methods(http.MethodGet)
	r.HandleFunc("/games/{playerID}/answer", triviaHandler.AnswerQuestion).Methods(http.MethodPost)

	// Add more routes as necessary for your application.

	return r
}