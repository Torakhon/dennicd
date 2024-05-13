package usecase

import (
	"Healthcare_Evrone/internal/entity"
	"Healthcare_Evrone/internal/infrastructure/repository"
	"Healthcare_Evrone/internal/pkg/otlp"
	"context"
	"go.opentelemetry.io/otel/attribute"
	"time"
)

const (
	serviceNameSpecializationUseCase           = "specializationUseCase"
	serviceNameSpecializationUseCaseRepoPrefix = "specializationUseCase"
)

type SpecializationUsecase interface {
	CreateSpecialization(ctx context.Context, specialization *entity.Specialization) (*entity.Specialization, error)
	GetSpecializationById(ctx context.Context, in *entity.GetReqStr) (*entity.Specialization, error)
	GetAllSpecializations(ctx context.Context, page, limit int64, search string) ([]*entity.Specialization, error)
	UpdateSpecialization(ctx context.Context, in *entity.Specialization) (*entity.Specialization, error)
	DeleteSpecialization(ctx context.Context, in *entity.GetReqStr) (bool, error)
}

type newsSpecService struct {
	BaseUseCase
	repo       repository.SpecializationRepository
	ctxTimeout time.Duration
}

func NewSpecializationService(ctxTimeout time.Duration, repo repository.SpecializationRepository) newsSpecService {
	return newsSpecService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (n newsSpecService) CreateSpecialization(ctx context.Context, specialization *entity.Specialization) (*entity.Specialization, error) {
	ctx, cancel := context.WithTimeout(ctx, n.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameSpecializationUseCase, serviceNameSpecializationUseCaseRepoPrefix+"Create")
	span.SetAttributes(attribute.String("CreateSpecialization", specialization.ID))
	defer span.End()

	return n.repo.CreateSpecialization(ctx, specialization)
}

func (n newsSpecService) GetSpecializationById(ctx context.Context, in *entity.GetReqStr) (*entity.Specialization, error) {
	ctx, cancel := context.WithTimeout(ctx, n.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameSpecializationUseCase, serviceNameSpecializationUseCaseRepoPrefix+"Get")
	span.SetAttributes(attribute.String("GetSpecializationById", in.Id))

	defer span.End()

	return n.repo.GetSpecializationById(ctx, &entity.GetReqStr{
		Id:       in.Id,
		IsActive: in.IsActive,
	})
}

func (n newsSpecService) GetAllSpecializations(ctx context.Context, page, limit int64, search string) ([]*entity.Specialization, error) {
	ctx, cancel := context.WithTimeout(ctx, n.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameSpecializationUseCase, serviceNameSpecializationUseCaseRepoPrefix+"Get all")
	span.SetAttributes(attribute.String("GetAllSpecializations", search))

	defer span.End()

	return n.repo.GetAllSpecializations(ctx, page, limit, search)
}

func (n newsSpecService) UpdateSpecialization(ctx context.Context, in *entity.Specialization) (*entity.Specialization, error) {
	ctx, cancel := context.WithTimeout(ctx, n.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameSpecializationUseCase, serviceNameSpecializationUseCaseRepoPrefix+"Update")
	span.SetAttributes(attribute.String("UpdateSpecialization", in.ID))
	defer span.End()

	return n.repo.UpdateSpecialization(ctx, in)
}

func (n newsSpecService) DeleteSpecialization(ctx context.Context, in *entity.GetReqStr) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, n.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameSpecializationUseCase, serviceNameSpecializationUseCaseRepoPrefix+"Delete")
	span.SetAttributes(attribute.String("DeleteSpecialization", in.Id))

	defer span.End()

	return n.repo.DeleteSpecialization(ctx, in)
}
