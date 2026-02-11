package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/qndaa/pack-calculator/internal/repository"
	"github.com/qndaa/pack-calculator/internal/server"
	"github.com/qndaa/pack-calculator/internal/usecase"
)

type App struct {
	server *http.Server
}

func New() (*App, error) {
	cfg := NewConfig()

	packRepository, err := repository.NewPackRepository(cfg.repositoryConfig)
	if err != nil {
		return nil, err
	}

	calculator := usecase.NewCalculator(packRepository)
	handler := server.NewHandler(calculator)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	return &App{
		server: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}, nil
}

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("Starting server on :8080")
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return a.server.Shutdown(shutdownCtx)
}
