package repository

import (
	"Healthcare_Evrone/internal/entity"
	"context"
)

type DoctorServices interface {
	CreateDoctorServices(ctx context.Context, in *entity.DoctorServices) (*entity.DoctorServices, error)
	GetDoctorServiceByID(ctx context.Context, in *entity.GetReqStr) (*entity.DoctorServices, error)
	GetAllDoctorServices(ctx context.Context, page, limit int64, search string) ([]*entity.DoctorServices, error)
	UpdateDoctorServices(ctx context.Context, in *entity.DoctorServices) (*entity.DoctorServices, error)
	DeleteDoctorService(ctx context.Context, in *entity.GetReqStr) (bool, error)
}
