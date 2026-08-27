package reviewhdl

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"review/internal/domain"
	"review/internal/transport/http/dto/reviewdto"
	"review/internal/transport/http/httputil/helpers"
)

type Svc interface {
	GetAll(c context.Context) ([]domain.Review, error)
	GetAllByProductID(c context.Context, id uuid.UUID) ([]domain.Review, error)
	GetByID(c context.Context, id uuid.UUID) (*domain.Review, error)
	Create(c context.Context, review domain.Review) (domain.Review, error)
	Patch(c context.Context, review reviewdto.UpdateReviewRequest) (domain.Review, error)
	Delete(c context.Context, id uuid.UUID) error
}

type Handler struct {
	svc Svc
}

func NewHandler(svc Svc) *Handler {
	return &Handler{}
}

func (h *Handler) GetAll(c *echo.Context) ([]domain.Review, error) {
	return nil, helpers.InternalErr(c)
}

func (h *Handler) GetAllByProductID(c *echo.Context) ([]domain.Review, error) {
	return nil, helpers.InternalErr(c)
}

func (h *Handler) GetByID(c *echo.Context) (domain.Review, error) {
	return domain.Review{}, helpers.InternalErr(c)
}

func (h *Handler) Post(c *echo.Context) (domain.Review, error) {
	return domain.Review{}, helpers.InternalErr(c)
}

func (h *Handler) Update(c *echo.Context) (domain.Review, error) {
	return domain.Review{}, helpers.InternalErr(c)
}

func (h *Handler) Delete(c *echo.Context) error {
	return helpers.NoContent(c)
}
