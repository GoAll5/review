package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v5"
	echoSwagger "github.com/swaggo/echo-swagger/v2"
	"log/slog"
	"review/internal/transport/http/echoserver"
)

type ReviewHandler interface {
	GetAll(c *echo.Context) error
	GetAllByProductID(c *echo.Context) error
	Update(c *echo.Context) error
	Post(c *echo.Context) error
	Delete(c *echo.Context) error
	GetByID(c *echo.Context) error
}

func setupHttp(
	ctx context.Context,
	log *slog.Logger,
	addr string,
	handler ReviewHandler) (*echoserver.Server, error) {
	op := "register protected routes"

	httpSrv := echoserver.New(ctx, log, addr)

	httpSrv.E.GET("/swagger/*", echoSwagger.WrapHandler)

	if err := registerReviewsRoutes(httpSrv, handler); err != nil {
		return httpSrv, fmt.Errorf("%s: %w", op, err)
	}

	return httpSrv, nil
}

func registerReviewsRoutes(httpSrv *echoserver.Server, reviewH ReviewHandler) error {
	reviews := httpSrv.E.Group("/reviews")
	reviews.GET("", reviewH.GetAll)
	reviews.GET("/products/:id", reviewH.GetAllByProductID)
	reviews.GET("/:id", reviewH.GetByID)
	reviews.PATCH("/:id", reviewH.Update)
	reviews.DELETE("/:id", reviewH.Delete)
	reviews.POST("", reviewH.Post)

	return nil
}
