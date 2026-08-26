package reviewdto

import (
	"github.com/google/uuid"
	"review/internal/domain"
	"time"
)

type Review struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	UserID    uuid.UUID `json:"user_id"`
	Rate      int       `json:"rate"`
	Text      *string   `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r Review) ToDomain() domain.Review {
	return domain.Review{
		ID:        r.ID,
		ProductID: r.ProductID,
		UserID:    r.UserID,
		Rate:      r.Rate,
		Text:      r.Text,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func ReviewToResponse(r domain.Review) Review {
	return Review{
		ID:        r.ID,
		ProductID: r.ProductID,
		UserID:    r.UserID,
		Rate:      r.Rate,
		Text:      r.Text,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

type CreateReviewRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	Text      *string   `json:"text" validate:"omitempty"`
	Rate      int       `json:"rate" validate:"required,min=1,max=5"`
}

func (req CreateReviewRequest) ToDomain() domain.Review {
	return domain.Review{
		ProductID: req.ProductID,
		UserID:    req.UserID,
		Rate:      req.Rate,
		Text:      req.Text,
	}
}

type UpdateReviewRequest struct {
	ID   uuid.UUID `param:"id" validate:"required"`
	Text *string   `json:"text" validate:"omitempty"`
	Rate *int      `json:"rate" validate:"omitempty,min=1,max=5"`
}

//Достаете старый Review из базы данных по ID.
//Вызываете updatedReview := req.MergeTo(oldReview).
//Сохраняете updatedReview обратно в базу.

func (ur UpdateReviewRequest) ToMerge(current domain.Review) domain.Review {
	if ur.Text != nil {
		current.Text = ur.Text
	}
	if ur.Rate != nil {
		current.Rate = *ur.Rate
	}
	return current
}
