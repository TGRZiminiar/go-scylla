package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tgrzimiar/go-scylla/config"
	"time"

	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type (
	server struct {
		app    *fiber.App
		db     *gocql.Session
		cfg    *config.Config
		logger *zap.Logger
	}
)

// func newMiddleware(cfg *config.Config) middlewarehandler.MiddlewareHandlerService {
// 	repo := middlewarerepository.NewMiddlewareRepository()
// 	usecase := middlewareusecase.NewMiddlewareUsecase(repo)
// 	return middlewarehandler.NewMiddlewareHandler(cfg, usecase)
// }

func (s *server) gracefulShutdown(pctx context.Context, quit <-chan os.Signal) {

	log.Printf("Starting service: %s", s.cfg.App.Name)

	<-quit

	log.Printf("Shutting down service: %s", s.cfg.App.Name)

	if err := s.app.Shutdown(); err != nil {
		log.Fatalf("Error: %v", err)
	}

}

func (s *server) httpListening() {
	if err := s.app.Listen(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}
}

func Start(pctx context.Context, cfg *config.Config, db *gocql.Session, logger *zap.Logger) {
	s := &server{
		db:  db,
		cfg: cfg,
		app: fiber.New(fiber.Config{
			AppName:      cfg.App.Name,
			BodyLimit:    10 * 1024 * 1024,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 20 * time.Second,
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
		logger: logger,
	}

	// Body Limit
	// app.Settings.MaxRequestBodySize = 10 * 1024 * 1024 // 10 MB

	switch s.cfg.App.Name {
	case "users":
		s.usersService()
	}

	// s.app.Static("/asset/")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go s.gracefulShutdown(pctx, quit)

	// Listening
	s.httpListening()

}
