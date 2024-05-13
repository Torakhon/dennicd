package services

import (
	pb "Healthcare_Evrone/genproto/healthcare-service"
	"Healthcare_Evrone/internal/entity"
	"Healthcare_Evrone/internal/pkg/otlp"
	"Healthcare_Evrone/internal/usecase"
	"context"
	"go.opentelemetry.io/otel/attribute"
	"time"

	"go.uber.org/zap"
)

type doctorRPC struct {
	logger *zap.Logger
	doctor usecase.DoctorUsecase
}

const (
	serviceNameDoctorDelivery           = "doctorDelivery"
	serviceNameDoctorDeliveryRepoPrefix = "doctorDelivery"
)

func DoctorRPC(logget *zap.Logger, doctorUsecase usecase.DoctorUsecase) pb.DoctorServiceServer {
	return &doctorRPC{
		logget,
		doctorUsecase,
	}

}

func (r doctorRPC) CreateDoctor(ctx context.Context, doctor *pb.Doctor) (*pb.Doctor, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctorDelivery, serviceNameDoctorDeliveryRepoPrefix+"Create")
	span.SetAttributes(attribute.Key("CreateDoctor").String(doctor.Id))
	defer span.End()

	req := entity.Doctor{
		Id:            doctor.Id,
		Order:         doctor.Order,
		FirstName:     doctor.FirstName,
		LastName:      doctor.LastName,
		Gender:        doctor.Gender,
		BirthDate:     doctor.BirthDate,
		PhoneNumber:   doctor.PhoneNumber,
		Email:         doctor.Email,
		Address:       doctor.Address,
		City:          doctor.City,
		Country:       doctor.Country,
		Salary:        doctor.Salary,
		Bio:           doctor.Bio,
		StartWorkDate: doctor.StartWorkDate,
		EndWorkDate:   doctor.EndWorkDate,
		WorkYears:     doctor.WorkYears,
		DepartmentId:  doctor.DepartmentId,
		RoomNumber:    doctor.RoomNumber,
		Password:      doctor.Password,
	}

	resp, err := r.doctor.CreateDoctor(ctx, &req)
	if err != nil {
		r.logger.Error("Failed to create doctor", zap.Error(err))
		return nil, err
	}

	return &pb.Doctor{
		Id:            resp.Id,
		Order:         resp.Order,
		FirstName:     resp.FirstName,
		LastName:      resp.LastName,
		Gender:        resp.Gender,
		BirthDate:     resp.BirthDate,
		PhoneNumber:   resp.PhoneNumber,
		Email:         resp.Email,
		Password:      resp.Password,
		Address:       resp.Address,
		City:          resp.City,
		Country:       resp.Country,
		Salary:        resp.Salary,
		Bio:           resp.Bio,
		StartWorkDate: resp.StartWorkDate,
		EndWorkDate:   resp.EndWorkDate,
		WorkYears:     resp.WorkYears,
		DepartmentId:  resp.DepartmentId,
		RoomNumber:    resp.RoomNumber,
		CreatedAt:     resp.CreatedAt.String(),
		UpdatedAt:     resp.UpdatedAt.String(),
		DeletedAt:     resp.DeletedAt.String(),
	}, nil
}

func (r doctorRPC) GetDoctorById(ctx context.Context, str *pb.GetReqStrDoctor) (*pb.Doctor, error) {
	ctx, span := otlp.Start(ctx, serviceNameDoctorDelivery, serviceNameDoctorDeliveryRepoPrefix+"Get")
	span.SetAttributes(attribute.Key("GetDoctorById").String(str.Id))
	defer span.End()
	doctor, err := r.doctor.GetDoctorById(ctx, &entity.GetReqStr{Id: str.Id, IsActive: str.IsActive})
	if err != nil {
		r.logger.Error("Failed to get doctor", zap.Error(err))
		return nil, err
	}
	return &pb.Doctor{
		Id:            doctor.Id,
		Order:         doctor.Order,
		FirstName:     doctor.FirstName,
		LastName:      doctor.LastName,
		Gender:        doctor.Gender,
		BirthDate:     doctor.BirthDate,
		PhoneNumber:   doctor.PhoneNumber,
		Email:         doctor.Email,
		Address:       doctor.Address,
		City:          doctor.City,
		Country:       doctor.Country,
		Salary:        doctor.Salary,
		Bio:           doctor.Bio,
		StartWorkDate: doctor.StartWorkDate,
		EndWorkDate:   doctor.EndWorkDate,
		WorkYears:     doctor.WorkYears,
		DepartmentId:  doctor.DepartmentId,
		RoomNumber:    doctor.RoomNumber,
		Password:      doctor.Password,
		CreatedAt:     doctor.CreatedAt.String(),
		UpdatedAt:     doctor.UpdatedAt.String(),
		DeletedAt:     doctor.DeletedAt.String(),
	}, nil
}

func (r doctorRPC) GetAllDoctors(ctx context.Context, all *pb.GetAllDoctorS) (*pb.ListDoctors, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctorDelivery, serviceNameDoctorDeliveryRepoPrefix+"Get all")
	span.SetAttributes(attribute.Key("GetAllDoctors").String(all.Search))
	defer span.End()
	resp, err := r.doctor.GetAllDoctors(ctx, all.Page, all.Limit, all.Search)
	if err != nil {
		r.logger.Error("Failed to get all doctors", zap.Error(err))
		return nil, err
	}

	var doctors pb.ListDoctors

	for _, doctor := range resp {
		doctors.Doctors = append(doctors.Doctors, &pb.Doctor{
			Id:            doctor.Id,
			Order:         doctor.Order,
			FirstName:     doctor.FirstName,
			LastName:      doctor.LastName,
			Gender:        doctor.Gender,
			BirthDate:     doctor.BirthDate,
			PhoneNumber:   doctor.PhoneNumber,
			Email:         doctor.Email,
			Address:       doctor.Address,
			City:          doctor.City,
			Country:       doctor.Country,
			Salary:        doctor.Salary,
			Bio:           doctor.Bio,
			StartWorkDate: doctor.StartWorkDate,
			EndWorkDate:   doctor.EndWorkDate,
			WorkYears:     doctor.WorkYears,
			DepartmentId:  doctor.DepartmentId,
			RoomNumber:    doctor.RoomNumber,
			Password:      doctor.Password,
			CreatedAt:     doctor.CreatedAt.String(),
			UpdatedAt:     doctor.UpdatedAt.String(),
			DeletedAt:     doctor.DeletedAt.String(),
		})
	}

	return &doctors, nil
}

func (r doctorRPC) UpdateDoctor(ctx context.Context, doctor *pb.Doctor) (*pb.Doctor, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctorDelivery, serviceNameDoctorDeliveryRepoPrefix+"Update")
	span.SetAttributes(attribute.Key("UpdateDoctor").String(doctor.Id))
	defer span.End()

	req := entity.Doctor{
		Id:            doctor.Id,
		Order:         doctor.Order,
		FirstName:     doctor.FirstName,
		LastName:      doctor.LastName,
		Gender:        doctor.Gender,
		BirthDate:     doctor.BirthDate,
		PhoneNumber:   doctor.PhoneNumber,
		Email:         doctor.Email,
		Address:       doctor.Address,
		City:          doctor.City,
		Country:       doctor.Country,
		Salary:        doctor.Salary,
		Bio:           doctor.Bio,
		StartWorkDate: doctor.StartWorkDate,
		EndWorkDate:   doctor.EndWorkDate,
		WorkYears:     doctor.WorkYears,
		DepartmentId:  doctor.DepartmentId,
		RoomNumber:    doctor.RoomNumber,
		Password:      doctor.Password,
		UpdatedAt:     time.Now(),
	}
	resp, err := r.doctor.UpdateDoctor(ctx, &req)
	if err != nil {
		r.logger.Error("Failed to update doctor", zap.Error(err))
		return nil, err
	}

	return &pb.Doctor{
		Id:            resp.Id,
		Order:         resp.Order,
		FirstName:     resp.FirstName,
		LastName:      resp.LastName,
		Gender:        resp.Gender,
		BirthDate:     resp.BirthDate,
		PhoneNumber:   resp.PhoneNumber,
		Email:         resp.Email,
		Password:      resp.Password,
		Address:       resp.Address,
		City:          resp.City,
		Country:       resp.Country,
		Salary:        resp.Salary,
		Bio:           resp.Bio,
		StartWorkDate: resp.StartWorkDate,
		EndWorkDate:   resp.EndWorkDate,
		WorkYears:     resp.WorkYears,
		DepartmentId:  resp.DepartmentId,
		RoomNumber:    resp.RoomNumber,
		CreatedAt:     resp.CreatedAt.String(),
		UpdatedAt:     resp.UpdatedAt.String(),
		DeletedAt:     resp.DeletedAt.String(),
	}, nil
}

func (r doctorRPC) DeleteDoctor(ctx context.Context, str *pb.GetReqStrDoctor) (*pb.StatusDoctor, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctorDelivery, serviceNameDoctorDeliveryRepoPrefix+"Delete")
	span.SetAttributes(attribute.Key("DeleteDoctor").String(str.Id))
	defer span.End()

	status, err := r.doctor.DeleteDoctor(ctx, &entity.GetReqStr{Id: str.Id, IsHardDeleted: str.IsHardDeleted})
	if err != nil {
		r.logger.Error("deleted doctor error", zap.Error(err))
		return nil, err
	}
	return &pb.StatusDoctor{Status: status}, nil
}