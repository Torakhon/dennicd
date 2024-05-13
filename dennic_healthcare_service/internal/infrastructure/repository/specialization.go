package repository

import (
	"Healthcare_Evrone/internal/entity"
	"context"
)

type SpecializationRepository interface {
	CreateSpecialization(ctx context.Context, specialization *entity.Specialization) (*entity.Specialization, error)
	GetSpecializationById(ctx context.Context, in *entity.GetReqStr) (*entity.Specialization, error)
	GetAllSpecializations(ctx context.Context, page, limit int64, search string) ([]*entity.Specialization, error)
	UpdateSpecialization(ctx context.Context, in *entity.Specialization) (*entity.Specialization, error)
	DeleteSpecialization(ctx context.Context, in *entity.GetReqStr) (bool, error)
}
