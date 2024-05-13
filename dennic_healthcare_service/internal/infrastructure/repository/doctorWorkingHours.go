package repository

import (
	"Healthcare_Evrone/internal/entity"
	"context"
)

type DoctorWorkingHoursRepository interface {
	CreateDoctorWorkingHours(ctx context.Context, in *entity.DoctorWorkingHours) (*entity.DoctorWorkingHours, error)
	GetDoctorWorkingHoursById(ctx context.Context, in *entity.GetReqInt) (*entity.DoctorWorkingHours, error)
	GetAllDoctorWorkingHours(ctx context.Context, page, limit int64, search string) ([]*entity.DoctorWorkingHours, error)
	UpdateDoctorWorkingHours(ctx context.Context, in *entity.DoctorWorkingHours) (*entity.DoctorWorkingHours, error)
	DeleteDoctorWorkingHours(ctx context.Context, in *entity.GetReqInt) (bool, error)
}
