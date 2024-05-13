package usecase

import (
	"Healthcare_Evrone/internal/entity"
	"Healthcare_Evrone/internal/infrastructure/repository"
	"Healthcare_Evrone/internal/pkg/otlp"
	"go.opentelemetry.io/otel/attribute"

	// "Healthcare_Evrone/internal/pkg/otlp"
	"context"
	"time"
)

const (
	serviceNameDoctorUseCase           = "doctorUseCase"
	serviceNameDoctorUseCaseRepoPrefix = "doctorUseCase"
)

type DoctorUsecase interface {
	CreateDoctor(ctx context.Context, doctor *entity.Doctor) (*entity.Doctor, error)
	GetDoctorById(ctx context.Context, get *entity.GetReqStr) (*entity.Doctor, error)
	GetAllDoctors(ctx context.Context, page, limit int64, search string) (doctors []*entity.Doctor, err error)
	UpdateDoctor(ctx context.Context, update *entity.Doctor) (*entity.Doctor, error)
	DeleteDoctor(ctx context.Context, del *entity.GetReqStr) (bool, error)
}

type newsService struct {
	BaseUseCase
	repo       repository.Doctor
	ctxTimeout time.Duration
}

func NewDoctorService(ctxTimeout time.Duration, repo repository.Doctor) newsService {
	return newsService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u newsService) CreateDoctor(ctx context.Context, doctor *entity.Doctor) (*entity.Doctor, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameDoctorUseCase, serviceNameDoctorUseCaseRepoPrefix+"Create")
	span.SetAttributes(attribute.Key("CreateDoctor").String(doctor.Id))
	defer span.End()

	return u.repo.CreateDoctor(ctx, doctor)
}

func (u newsService) GetDoctorById(ctx context.Context, get *entity.GetReqStr) (*entity.Doctor, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)

	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameDoctorUseCase, serviceNameDoctorUseCaseRepoPrefix+"Get")
	span.SetAttributes(attribute.Key("GetDoctorById").String(get.Id))
	defer span.End()

	return u.repo.GetDoctorById(ctx, get)
}

func (u newsService) GetAllDoctors(ctx context.Context, page, limit int64, search string) (doctors []*entity.Doctor, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameDoctorUseCase, serviceNameDoctorUseCaseRepoPrefix+"Get all")
	span.SetAttributes(attribute.Key("GetAllDoctors").String(search))

	defer span.End()

	return u.repo.GetAllDoctors(ctx, page, limit, search)
}

func (u newsService) UpdateDoctor(ctx context.Context, update *entity.Doctor) (*entity.Doctor, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameDoctorUseCase, serviceNameDoctorUseCaseRepoPrefix+"Update")
	span.SetAttributes(attribute.Key("UpdateDoctor").String(update.Id))
	defer span.End()

	return u.repo.UpdateDoctor(ctx, update)
}

func (u newsService) DeleteDoctor(ctx context.Context, del *entity.GetReqStr) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameDoctorUseCase, serviceNameDoctorUseCaseRepoPrefix+"Delete")
	span.SetAttributes(attribute.Key("DeleteDoctor").String(del.Id))
	defer span.End()

	return u.repo.DeleteDoctor(ctx, del)
}
