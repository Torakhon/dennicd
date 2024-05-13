package repository

import (
	"Healthcare_Evrone/internal/entity"
	"context"
)

type Doctor interface {
	CreateDoctor(ctx context.Context, doctor *entity.Doctor) (*entity.Doctor, error)
	GetDoctorById(ctx context.Context, get *entity.GetReqStr) (*entity.Doctor, error)
	GetAllDoctors(ctx context.Context, page, limit int64, search string) (doctors []*entity.Doctor, err error)
	UpdateDoctor(ctx context.Context, update *entity.Doctor) (*entity.Doctor, error)
	DeleteDoctor(ctx context.Context, del *entity.GetReqStr) (bool, error)
}
