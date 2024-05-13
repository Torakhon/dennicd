package repository

import (
	"Healthcare_Evrone/internal/entity"
	"context"
)

type DepartmentRepository interface {
	CreateDepartment(ctx context.Context, dep *entity.Department) (*entity.Department, error)
	GetDepartmentById(ctx context.Context, get *entity.GetReqStr) (*entity.Department, error)
	GetAllDepartments(ctx context.Context, page, limit int64, search string) ([]*entity.Department, error)
	UpdateDepartment(ctx context.Context, up *entity.Department) (*entity.Department, error)
	DeleteDepartment(ctx context.Context, get *entity.GetReqStr) (bool, error)
}
