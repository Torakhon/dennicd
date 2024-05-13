package postgresql

import (
	"Healthcare_Evrone/internal/entity"
	"Healthcare_Evrone/internal/pkg/otlp"
	"Healthcare_Evrone/internal/pkg/postgres"
	"context"
	"database/sql"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"time"
)

const (
	doctorServicesWorkingHoursTableName = "doctor_service"
	serviceNameDoctorServices           = "doctors_service"
	serviceNameDoctorServicesRepoPrefix = "doctors_service"
)

type Ds struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewDoctorServicesRepo(db *postgres.PostgresDB) *Ds {
	return &Ds{
		tableName: doctorServicesWorkingHoursTableName,
		db:        db,
	}
}

func (d Ds) doctorServicesSelectQueryPrefix() string {
	return `
			id,
			Dservice_order,
			doctor_id,
			specialization_id,
			online_price,
			offline_price,
			created_at,
			updated_at,
			deleted_at
		`
}

func (d Ds) CreateDoctorServices(ctx context.Context, in *entity.DoctorServices) (*entity.DoctorServices, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctorServices, serviceNameDoctorServicesRepoPrefix+"Create")
	span.SetAttributes(attribute.Key("CreateDoctorServices").String(in.Id))
	defer span.End()

	data := map[string]any{
		"id":                in.Id,
		"doctor_id":         in.DoctorId,
		"specialization_id": in.SpecializationId,
		"online_price":      in.OnlinePrice,
		"offline_price":     in.OfflinePrice,
	}
	query, args, err := d.db.Sq.Builder.Insert(d.tableName).
		SetMap(data).Suffix(fmt.Sprintf("RETURNING %s", d.doctorServicesSelectQueryPrefix())).ToSql()

	if err != nil {
		return nil, d.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", d.tableName, "create"))
	}
	var updatedAt, deletedAt sql.NullTime
	err = d.db.QueryRow(ctx, query, args...).Scan(
		&in.Id,
		&in.Order,
		&in.DoctorId,
		&in.SpecializationId,
		&in.OnlinePrice,
		&in.OfflinePrice,
		&in.CreatedAt,
		&updatedAt,
		&deletedAt,
	)
	if updatedAt.Valid {
		in.UpdatedAt = updatedAt.Time
	}
	if deletedAt.Valid {
		in.DeletedAt = deletedAt.Time
	}

	if err != nil {
		return nil, d.db.Error(err)
	}
	return in, nil
}

func (d Ds) GetDoctorServiceByID(ctx context.Context, in *entity.GetReqStr) (*entity.DoctorServices, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctorServices, serviceNameDoctorServicesRepoPrefix+"Get")
	defer span.End()

	var doctorService entity.DoctorServices
	queryBuilder := d.db.Sq.Builder.Select(d.doctorServicesSelectQueryPrefix()).From(d.tableName)
	if !in.IsActive {
		queryBuilder = queryBuilder.Where("deleted_at IS NULL")
	}
	queryBuilder = queryBuilder.Where(d.db.Sq.Equal("id", in.Id))
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var updatedAt, deletedAt sql.NullTime
	err = d.db.QueryRow(ctx, query, args...).Scan(
		&doctorService.Id,
		&doctorService.Order,
		&doctorService.DoctorId,
		&doctorService.SpecializationId,
		&doctorService.OnlinePrice,
		&doctorService.OfflinePrice,
		&doctorService.CreatedAt,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		return nil, err
	}
	if updatedAt.Valid {
		doctorService.UpdatedAt = updatedAt.Time
	}
	if deletedAt.Valid {
		doctorService.DeletedAt = deletedAt.Time
	}
	return &doctorService, nil
}

func (d Ds) GetAllDoctorServices(ctx context.Context, page, limit int64, search string) ([]*entity.DoctorServices, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctorServices, serviceNameDoctorServicesRepoPrefix+"Get all")
	defer span.End()

	offset := limit * (page - 1)

	queryBuilder := d.db.Sq.Builder.Select(d.doctorServicesSelectQueryPrefix()).From(d.tableName)
	//if search != "" {
	//	queryBuilder = queryBuilder.Where(fmt.Sprintf(`online_price ILIKE %f OR offline_price ILIKE %s`, search+"%", search+"%"))
	//}
	queryBuilder = queryBuilder.Limit(uint64(limit)).Offset(uint64(offset))
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var doctorServices []*entity.DoctorServices
	rows, err := d.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var dService entity.DoctorServices
		var deletedAt, updatedAt sql.NullTime
		err = rows.Scan(
			&dService.Id,
			&dService.Order,
			&dService.DoctorId,
			&dService.SpecializationId,
			&dService.OnlinePrice,
			&dService.OfflinePrice,
			&dService.CreatedAt,
			&updatedAt,
			&deletedAt,
		)
		if updatedAt.Valid {
			dService.UpdatedAt = updatedAt.Time
		}
		if deletedAt.Valid {
			dService.DeletedAt = deletedAt.Time
		}
		if err != nil {
			return nil, err
		}
		doctorServices = append(doctorServices, &dService)
	}
	return doctorServices, nil
}

func (d Ds) UpdateDoctorServices(ctx context.Context, services *entity.DoctorServices) (*entity.DoctorServices, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctorServices, serviceNameDoctorServicesRepoPrefix+"Update")
	defer span.End()

	data := map[string]any{
		"doctor_id":         services.DoctorId,
		"specialization_id": services.SpecializationId,
		"online_price":      services.OnlinePrice,
		"offline_price":     services.OfflinePrice,
		"updated_at":        time.Now(),
	}
	query, args, err := d.db.Sq.Builder.Update(d.tableName).
		SetMap(data).Where(d.db.Sq.Equal("id", services.Id)).
		Suffix(fmt.Sprintf("RETURNING %s", d.doctorServicesSelectQueryPrefix())).ToSql()

	if err != nil {
		return nil, d.db.ErrSQLBuild(err, d.tableName+" update")
	}
	var deletedAt, updatedAt sql.NullTime
	err = d.db.QueryRow(ctx, query, args...).Scan(
		&services.Id,
		&services.Order,
		&services.DoctorId,
		&services.SpecializationId,
		&services.OnlinePrice,
		&services.OfflinePrice,
		&services.CreatedAt,
		&updatedAt,
		&deletedAt,
	)
	if updatedAt.Valid {
		services.UpdatedAt = updatedAt.Time
	}
	if deletedAt.Valid {
		services.DeletedAt = deletedAt.Time
	}
	if err != nil {
		return nil, d.db.Error(err)
	}
	return services, nil
}

func (d Ds) DeleteDoctorService(ctx context.Context, in *entity.GetReqStr) (bool, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctorServices, serviceNameDoctorServicesRepoPrefix+"Delete")
	defer span.End()

	data := map[string]any{
		"deleted_at": time.Now(),
	}

	var args []interface{}
	var query string
	var err error
	if in.IsHardDeleted {
		query, args, err = d.db.Sq.Builder.Delete(d.tableName).From(d.tableName).Where(d.db.Sq.Equal("id", in.Id)).ToSql()
		if err != nil {
			return false, d.db.ErrSQLBuild(err, d.tableName+" delete")
		}
	} else {
		query, args, err = d.db.Sq.Builder.Update(d.tableName).SetMap(data).Where(d.db.Sq.Equal("id", in.Id)).ToSql()
		if err != nil {
			return false, d.db.ErrSQLBuild(err, d.tableName+" delete")
		}
	}
	_, err = d.db.Exec(ctx, query, args...)
	if err != nil {
		return false, d.db.Error(err)
	}
	return true, nil
}
