package reviewsvc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"review/internal/domain"
	"review/internal/transport/http/dto/reviewdto"
)

//type Svc interface {
//	GetAll(c context.Context) ([]domain.Review, error)
//	GetAllByProductID(c context.Context, id uuid.UUID) ([]domain.Review, error)
//	GetByID(c context.Context, id uuid.UUID) (*domain.Review, error)
//	Create(c context.Context, review domain.Review) (domain.Review, error)
//	Update(c context.Context, review reviewdto.UpdateReviewRequest) (domain.Review, error)
//	Delete(c context.Context, id uuid.UUID) error
//}

type Repository interface {
	GetAll(c context.Context) ([]domain.Review, error)
	GetAllByProductID(c context.Context, id uuid.UUID) ([]domain.Review, error)
	GetByID(c context.Context, id uuid.UUID) (*domain.Review, error)
	Create(c context.Context, review domain.Review) (domain.Review, error)
	Update(c context.Context, review domain.Review) (domain.Review, error)
	Delete(c context.Context, id uuid.UUID) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) GetAll(c context.Context) ([]domain.Review, error) {
	return svc.repo.GetAll(c)
}

func (svc *Service) GetAllByProductID(c context.Context, id uuid.UUID) ([]domain.Review, error) {
	return svc.repo.GetAllByProductID(c, id)
}

func (svc *Service) GetByID(c context.Context, id uuid.UUID) (*domain.Review, error) {
	return svc.repo.GetByID(c, id)
}

func (svc *Service) Create(c context.Context, review domain.Review) (domain.Review, error) {
	return svc.repo.Create(c, review)
}

// TODO: op
func (svc *Service) Patch(c context.Context, review reviewdto.UpdateReviewRequest) (domain.Review, error) {
	currentReview, err := svc.repo.GetByID(c, review.ID)

	if err != nil {
		return domain.Review{}, err
	}
	if currentReview == nil {
		return domain.Review{}, fmt.Errorf("review with ID %d not found", review.ID)
	}
	//var baseReview domain.Review
	//    if currentReview != nil {
	//        baseReview = *currentReview
	//    }

	return svc.repo.Update(c, review.ToMerge(*currentReview))
}

func (svc *Service) Delete(c context.Context, id uuid.UUID) error {
	return svc.repo.Delete(c, id)
}
