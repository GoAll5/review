package echoserver

import (
	"context"
	"errors"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"log/slog"
	"net/http"
)

type Server struct {
	E   *echo.Echo
	srv *http.Server
	log *slog.Logger
}

func New(ctx context.Context, log *slog.Logger, addr string) *Server {
	e := echo.New()

	e.Use(middleware.RequestLogger()) // логирует каждый запрос

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"}, //сторонние сайты запрос которым разрешен
		AllowCredentials: true,
	}))
	e.Use(middleware.Recover()) // panic -> 500 + stack (в dev)
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure()) // базовые security headers
	//e.Use(middleware.Logger())
	e.Use(middleware.CSRF())

	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	return &Server{E: e, srv: srv, log: log}
}

// TODO: e.Start vs listenAndServe
func (s *Server) Start() error {
	op := "http starting server"
	s.log.Info(op, slog.String("addr", s.srv.Addr))
	err := s.srv.ListenAndServe() // s.E.Start()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("http server shutting down")
	return s.srv.Shutdown(ctx)
}
