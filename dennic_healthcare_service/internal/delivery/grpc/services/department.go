package services

import (
	pb "Healthcare_Evrone/genproto/healthcare-service"
	"Healthcare_Evrone/internal/entity"
	"Healthcare_Evrone/internal/pkg/otlp"
	"Healthcare_Evrone/internal/usecase"
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"time"
)

type departmentRPC struct {
	logger     *zap.Logger
	department usecase.DepartmentsUsecase
}

const (
	serviceNameDepartmentDelivery           = "DepartmentDelivery"
	serviceNameDepartmentDeliveryRepoPrefix = "DepartmentDelivery"
)

func DepartmentRPC(logget *zap.Logger, departmentUsecase usecase.DepartmentsUsecase) pb.DepartmentServiceServer {
	return &departmentRPC{
		logget,
		departmentUsecase,
	}

}

func (r departmentRPC) CreateDepartment(ctx context.Context, dep *pb.Department) (*pb.Department, error) {

	ctx, span := otlp.Start(ctx, serviceNameDepartmentDelivery, serviceNameDepartmentDeliveryRepoPrefix+"Create")
	span.SetAttributes(attribute.Key("CreateDepartment").String(dep.Id))
	defer span.End()

	req := entity.Department{
		Id:          dep.Id,
		Order:       dep.Order,
		Name:        dep.Name,
		Description: dep.Description,
		ImageUrl:    dep.ImageUrl,
		FloorNumber: dep.FloorNumber,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   time.Time{},
	}
	resp, err := r.department.CreateDepartment(ctx, &req)
	if err != nil {

		return nil, err
	}
	return &pb.Department{
		Id:          resp.Id,
		Order:       resp.Order,
		Name:        resp.Name,
		Description: resp.Description,
		ImageUrl:    resp.ImageUrl,
		FloorNumber: resp.FloorNumber,
		CreatedAt:   resp.CreatedAt.String(),
		UpdatedAt:   resp.UpdatedAt.String(),
		DeletedAt:   resp.DeletedAt.String(),
	}, nil

}

func (r departmentRPC) GetDepartmentById(ctx context.Context, get *pb.GetReqStrDepartment) (*pb.Department, error) {
	ctx, span := otlp.Start(ctx, serviceNameDepartmentDelivery, serviceNameDepartmentDeliveryRepoPrefix+"Get")
	span.SetAttributes(attribute.Key("GetDepartmentById").String(get.Id))
	defer span.End()
	resp, err := r.department.GetDepartmentById(ctx, &entity.GetReqStr{Id: get.Id, IsActive: get.IsActive, IsHardDeleted: false})
	if err != nil {
		return nil, err
	}
	return &pb.Department{
		Id:          resp.Id,
		Order:       resp.Order,
		Name:        resp.Name,
		Description: resp.Description,
		ImageUrl:    resp.ImageUrl,
		FloorNumber: resp.FloorNumber,
		CreatedAt:   resp.CreatedAt.String(),
		UpdatedAt:   resp.UpdatedAt.String(),
		DeletedAt:   resp.DeletedAt.String(),
	}, nil
}

func (r departmentRPC) GetAllDepartments(ctx context.Context, get *pb.GetAllDepartment) (*pb.ListDepartments, error) {

	ctx, span := otlp.Start(ctx, serviceNameDepartmentDelivery, serviceNameDepartmentDeliveryRepoPrefix+"Get all")
	span.SetAttributes(attribute.Key("GetAllDepartments").String(get.Search))
	defer span.End()

	resp, err := r.department.GetAllDepartments(ctx, get.Page, get.Limit, get.Search)
	if err != nil {
		return nil, err
	}

	var departments pb.ListDepartments

	for _, dep := range resp {
		departments.Departments = append(departments.Departments, &pb.Department{
			Id:          dep.Id,
			Order:       dep.Order,
			Name:        dep.Name,
			Description: dep.Description,
			ImageUrl:    dep.ImageUrl,
			FloorNumber: dep.FloorNumber,
			CreatedAt:   dep.CreatedAt.String(),
			UpdatedAt:   dep.UpdatedAt.String(),
			DeletedAt:   dep.DeletedAt.String(),
		})
	}

	return &departments, nil
}

func (r departmentRPC) UpdateDepartment(ctx context.Context, update *pb.Department) (*pb.Department, error) {

	ctx, span := otlp.Start(ctx, serviceNameDepartmentDelivery, serviceNameDepartmentDeliveryRepoPrefix+"Update")
	span.SetAttributes(attribute.Key("UpdateDepartment").String(update.Id))
	defer span.End()
	req := entity.Department{
		Id:          update.Id,
		Order:       update.Order,
		Name:        update.Name,
		Description: update.Description,
		ImageUrl:    update.ImageUrl,
		FloorNumber: update.FloorNumber,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   time.Time{},
	}
	resp, err := r.department.UpdateDepartment(ctx, &req)
	if err != nil {
		return nil, err
	}
	return &pb.Department{
		Id:          resp.Id,
		Order:       resp.Order,
		Name:        resp.Name,
		Description: resp.Description,
		ImageUrl:    resp.ImageUrl,
		FloorNumber: resp.FloorNumber,
		CreatedAt:   resp.CreatedAt.String(),
		UpdatedAt:   resp.UpdatedAt.String(),
		DeletedAt:   resp.DeletedAt.String(),
	}, nil
}

func (r departmentRPC) DeleteDepartment(ctx context.Context, del *pb.GetReqStrDepartment) (*pb.StatusDepartment, error) {
	ctx, span := otlp.Start(ctx, serviceNameDepartmentDelivery, serviceNameDepartmentDeliveryRepoPrefix+"Delete")
	span.SetAttributes(attribute.Key("DeleteDepartment").String(del.Id))
	defer span.End()
	status, err := r.department.DeleteDepartment(ctx, &entity.GetReqStr{Id: del.Id, IsHardDeleted: del.IsHardDeleted})
	if err != nil {
		r.logger.Error("deleted department error", zap.Error(err))
		return nil, err
	}
	return &pb.StatusDepartment{Status: status}, nil
}
