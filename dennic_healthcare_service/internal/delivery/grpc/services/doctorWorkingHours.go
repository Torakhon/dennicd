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

type doctorWorkingHoursServiceRPC struct {
	logger             *zap.Logger
	doctorWorkingHours usecase.DoctorWorkingHoursUseCase
}

const (
	serviceNameDoctorWorkingHoursDelivery           = "doctorWorkingHoursDelivery"
	serviceNameDoctorWorkingHoursDeliveryRepoPrefix = "doctorWorkingHoursDelivery"
)

func DoctorWorkingHoursServiceRPC(logger *zap.Logger, dwhUsecase usecase.DoctorWorkingHoursUseCase) pb.DoctorWorkingHoursServer {
	return &doctorWorkingHoursServiceRPC{
		logger,
		dwhUsecase,
	}
}

func (r doctorWorkingHoursServiceRPC) CreateDoctorWorkingHours(ctx context.Context, hours *pb.DoctorWorkingHours) (*pb.DoctorWorkingHours, error) {
	ctx, span := otlp.Start(ctx, serviceNameDoctorWorkingHoursDelivery, serviceNameDoctorWorkingHoursDeliveryRepoPrefix+"Create")
	span.SetAttributes(attribute.Key("CreateDoctorWorkingHours").String(string(hours.Id)))
	defer span.End()
	req := entity.DoctorWorkingHours{
		Id:         hours.Id,
		DoctorId:   hours.DoctorId,
		DayOfWeek:  hours.DayOfWeek,
		StartTime:  hours.StartTime,
		FinishTime: hours.FinishTime,
	}
	resp, err := r.doctorWorkingHours.CreateDoctorWorkingHours(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &pb.DoctorWorkingHours{
		Id:         resp.Id,
		DoctorId:   resp.DoctorId,
		DayOfWeek:  resp.DayOfWeek,
		StartTime:  resp.StartTime,
		FinishTime: resp.FinishTime,
		CreatedAt:  resp.CreatedAt.String(),
		UpdatedAt:  resp.UpdatedAt.String(),
		DeletedAt:  resp.DeletedAt.String(),
	}, nil
}

func (r doctorWorkingHoursServiceRPC) GetDoctorWorkingHoursById(ctx context.Context, in *pb.GetReqInt) (*pb.DoctorWorkingHours, error) {
	ctx, span := otlp.Start(ctx, serviceNameDoctorWorkingHoursDelivery, serviceNameDoctorWorkingHoursDeliveryRepoPrefix+"Get")
	span.SetAttributes(attribute.Key("GetDoctorWorkingHoursById").String(string(in.Id)))
	defer span.End()
	dwh, err := r.doctorWorkingHours.GetDoctorWorkingHoursById(ctx, &entity.GetReqInt{
		Id:       in.Id,
		IsActive: in.IsActive,
	})
	if err != nil {
		return nil, err
	}
	return &pb.DoctorWorkingHours{
		Id:         dwh.Id,
		DoctorId:   dwh.DoctorId,
		DayOfWeek:  dwh.DayOfWeek,
		StartTime:  dwh.StartTime,
		FinishTime: dwh.FinishTime,
		CreatedAt:  dwh.CreatedAt.String(),
		UpdatedAt:  dwh.UpdatedAt.String(),
		DeletedAt:  dwh.DeletedAt.String(),
	}, nil
}

func (r doctorWorkingHoursServiceRPC) GetAllDoctorWorkingHours(ctx context.Context, all *pb.GetAllDoctorWorkingHourS) (*pb.ListDoctorWorkingHours, error) {
	ctx, span := otlp.Start(ctx, serviceNameDoctorWorkingHoursDelivery, serviceNameDoctorWorkingHoursDeliveryRepoPrefix+"Get all")
	span.SetAttributes(attribute.Key("GetAllDoctorWorkingHours").String(all.Search))
	defer span.End()
	dwh, err := r.doctorWorkingHours.GetAllDoctorWorkingHours(ctx, all.Page, all.Limit, all.Search)
	if err != nil {
		return nil, err
	}
	var listDoctorWorkingHours pb.ListDoctorWorkingHours
	for _, d := range dwh {
		listDoctorWorkingHours.Dwh = append(listDoctorWorkingHours.Dwh, &pb.DoctorWorkingHours{
			Id:         d.Id,
			DoctorId:   d.DoctorId,
			DayOfWeek:  d.DayOfWeek,
			StartTime:  d.StartTime,
			FinishTime: d.FinishTime,
			CreatedAt:  d.CreatedAt.String(),
			UpdatedAt:  d.UpdatedAt.String(),
			DeletedAt:  d.DeletedAt.String(),
		})
	}
	return &listDoctorWorkingHours, nil
}

func (r doctorWorkingHoursServiceRPC) UpdateDoctorWorkingHours(ctx context.Context, hours *pb.DoctorWorkingHours) (*pb.DoctorWorkingHours, error) {
	ctx, span := otlp.Start(ctx, serviceNameDoctorWorkingHoursDelivery, serviceNameDoctorWorkingHoursDeliveryRepoPrefix+"Update")
	span.SetAttributes(attribute.Key("UpdateDoctorWorkingHours").String(string(hours.Id)))
	defer span.End()
	resp, err := r.doctorWorkingHours.UpdateDoctorWorkingHours(ctx, &entity.DoctorWorkingHours{
		Id:         hours.Id,
		DoctorId:   hours.DoctorId,
		DayOfWeek:  hours.DayOfWeek,
		StartTime:  hours.StartTime,
		FinishTime: hours.FinishTime,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  time.Time{},
	})
	if err != nil {
		return nil, err
	}
	return &pb.DoctorWorkingHours{
		Id:         resp.Id,
		DoctorId:   resp.DoctorId,
		DayOfWeek:  resp.DayOfWeek,
		StartTime:  resp.StartTime,
		FinishTime: resp.FinishTime,
		CreatedAt:  resp.CreatedAt.String(),
		UpdatedAt:  resp.UpdatedAt.String(),
		DeletedAt:  resp.DeletedAt.String(),
	}, nil
}

func (r doctorWorkingHoursServiceRPC) DeleteDoctorWorkingHours(ctx context.Context, reqInt *pb.GetReqInt) (*pb.StatusDoctorWorkingHours, error) {
	ctx, span := otlp.Start(ctx, serviceNameDoctorWorkingHoursDelivery, serviceNameDoctorWorkingHoursDeliveryRepoPrefix+"Delete")
	span.SetAttributes(attribute.Key("DeleteDoctorWorkingHours").String(string(reqInt.Id)))
	defer span.End()
	status, err := r.doctorWorkingHours.DeleteDoctorWorkingHours(ctx, &entity.GetReqInt{Id: reqInt.Id, IsHardDeleted: reqInt.IsHardDeleted})
	if err != nil {
		return nil, err
	}
	return &pb.StatusDoctorWorkingHours{Status: status}, nil
}