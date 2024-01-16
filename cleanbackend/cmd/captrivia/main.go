// cmd/captrivia/main.go
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ProlificLabs/captrivia/internal/application/usecases"
	"github.com/ProlificLabs/captrivia/internal/domain/services"
	"github.com/ProlificLabs/captrivia/internal/infrastructure/config"
	"github.com/ProlificLabs/captrivia/internal/infrastructure/db"
	"github.com/ProlificLabs/captrivia/internal/infrastructure/logging"
	"github.com/ProlificLabs/captrivia/internal/infrastructure/web"
	"github.com/ProlificLabs/captrivia/internal/interfaces/rest"
	"github.com/ProlificLabs/captrivia/internal/interfaces/storage"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type App struct {
	Server    *http.Server
	DB        *sqlx.DB
	Logger    *zap.Logger
	WaitGroup sync.WaitGroup
}

// Setup initializes the components of the App and returns it.
func Setup() (*App, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	logger, err := logging.NewLogger()
	if err != nil {
		return nil, err
	}

	database, err := db.NewDatabase(cfg)
	if err != nil {
		logger.Fatal("failed to initialize database", zap.Error(err))
	}

	// Initialize repositories
	questionRepo := storage.NewQuestionRepo(database)
	playerRepo := storage.NewPlayerRepo(database)

	// Initialize domain services
	gameService := services.NewGameService(questionRepo)

	// Initialize use cases
	triviaUseCase := usecases.NewTriviaUseCase(questionRepo, gameService)
	playerUseCase := usecases.NewPlayerUseCase(playerRepo)

	// Setup HTTP router
	router := rest.NewRouter(triviaUseCase, playerUseCase)

	// Initialize HTTP server
	server := &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      web.SetupGlobalMiddleware(router),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return &App{Server: server, DB: database, Logger: logger}, nil
}

func main() {
	app, err := Setup()
	if err != nil {
		panic("Application setup failed: " + err.Error())
	}
	defer app.DB.Close()
	defer app.Logger.Sync()

	app.WaitGroup.Add(1)
	go func() {
		defer app.WaitGroup.Done()
		app.Logger.Info("Starting CapTrivia server", zap.String("address", app.Server.Addr))
		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("failed to listen and serve", zap.Error(err))
		}
	}()

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	app.Logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		app.Logger.Fatal("server shutdown failed", zap.Error(err))
	}

	app.Logger.Info("Server gracefully stopped")
	app.WaitGroup.Wait()
}