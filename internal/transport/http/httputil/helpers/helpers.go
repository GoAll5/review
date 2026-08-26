package helpers

import (
	"github.com/labstack/echo/v5"
	"net/http"
	"review/internal/domain"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func BadRequest(c *echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, ErrorResponse{Error: msg})
}

func Unauthorized(c *echo.Context, msg string) error {
	return c.JSON(http.StatusUnauthorized, ErrorResponse{Error: msg})
}

func InternalErr(c *echo.Context) error {
	return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "internal error"})
}

func NotFound(c *echo.Context) error {
	return c.JSON(http.StatusNotFound, ErrorResponse{Error: "not found"})
}

func Ok(c *echo.Context, review domain.Review) error {
	return c.JSON(http.StatusOK, review)
}

func NoContent(c *echo.Context) error {
	return c.JSON(http.StatusNoContent, ErrorResponse{Error: "no content"})
}
