package postgresql

import (
	"Healthcare_Evrone/internal/entity"
	"Healthcare_Evrone/internal/pkg/otlp"
	"context"
	"database/sql"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"time"

	// "Healthcare_Evrone/internal/pkg/otlp"
	"Healthcare_Evrone/internal/pkg/postgres"
)

const (
	doctorTableName             = "doctors"
	serviceNameDoctor           = "doctors"
	serviceNameDoctorRepoPrefix = "doctors"
)

type DocTor struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewDoctorRepo(db *postgres.PostgresDB) *DocTor {
	return &DocTor{
		tableName: doctorTableName,
		db:        db,
	}
}

func (p *DocTor) docTorSelectQueryPrefix() string {
	return `id,
			doctor_order,
			first_name,
			last_name,
			gender,
			phone_number,
			email,
			password,
			address,
			city,
			country,
			salary,
			biography,
			start_work_date,
			end_work_date,
			work_years,
			department_id,
			room_number,
			created_at,
			updated_at
		`
}

func (h *DocTor) CreateDoctor(ctx context.Context, req *entity.Doctor) (*entity.Doctor, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctor, serviceNameDoctorRepoPrefix+"Create")
	span.SetAttributes(attribute.Key("CreateDoctor").String(req.Id))
	defer span.End()

	data := map[string]any{
		"id":              req.Id,
		"first_name":      req.FirstName,
		"last_name":       req.LastName,
		"gender":          req.Gender,
		"phone_number":    req.PhoneNumber,
		"email":           req.Email,
		"password":        req.Password,
		"address":         req.Address,
		"city":            req.City,
		"country":         req.Country,
		"salary":          req.Salary,
		"biography":       req.Bio,
		"start_work_date": req.StartWorkDate,
		"end_work_date":   req.EndWorkDate,
		"work_years":      req.WorkYears,
		"department_id":   req.DepartmentId,
		"room_number":     req.RoomNumber,
	}
	query, args, err := h.db.Sq.Builder.Insert(h.tableName).
		SetMap(data).Suffix(fmt.Sprintf("RETURNING %s", h.docTorSelectQueryPrefix())).ToSql()

	if err != nil {
		return nil, h.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", h.tableName, "create"))
	}
	var startWorkYear, endWorkYear, updatedAt sql.NullTime
	err = h.db.QueryRow(ctx, query, args...).Scan(
		&req.Id,
		&req.Order,
		&req.FirstName,
		&req.LastName,
		&req.Gender,
		&req.PhoneNumber,
		&req.Email,
		&req.Password,
		&req.Address,
		&req.City,
		&req.Country,
		&req.Salary,
		&req.Bio,
		&startWorkYear,
		&endWorkYear,
		&req.WorkYears,
		&req.DepartmentId,
		&req.RoomNumber,
		&req.CreatedAt,
		&updatedAt,
	)

	req.StartWorkDate = startWorkYear.Time.String()
	req.EndWorkDate = endWorkYear.Time.String()
	if err != nil {
		return nil, h.db.Error(err)
	}

	return req, nil
}

func (h *DocTor) GetDoctorById(ctx context.Context, get *entity.GetReqStr) (*entity.Doctor, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctor, serviceNameDoctorRepoPrefix+"Get")
	span.SetAttributes(attribute.Key("GetDoctorById").String(get.Id))

	defer span.End()

	var doctor entity.Doctor
	var startWorkYear, endWorkYear, updatedAt sql.NullTime
	queryBuilder := h.db.Sq.Builder.Select(h.docTorSelectQueryPrefix()).From(doctorTableName)

	if !get.IsActive {
		queryBuilder = queryBuilder.Where("deleted_at IS NULL")
	}

	queryBuilder = queryBuilder.Where(h.db.Sq.Equal("id", get.Id))

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	err = h.db.QueryRow(ctx, query, args...).Scan(
		&doctor.Id,
		&doctor.Order,
		&doctor.FirstName,
		&doctor.LastName,
		&doctor.Gender,
		&doctor.PhoneNumber,
		&doctor.Email,
		&doctor.Password,
		&doctor.Address,
		&doctor.City,
		&doctor.Country,
		&doctor.Salary,
		&doctor.Bio,
		&startWorkYear,
		&endWorkYear,
		&doctor.WorkYears,
		&doctor.DepartmentId,
		&doctor.RoomNumber,
		&doctor.CreatedAt,
		&updatedAt,
	)
	if updatedAt.Valid {
		doctor.UpdatedAt = updatedAt.Time
	}

	doctor.StartWorkDate = startWorkYear.Time.String()
	doctor.EndWorkDate = endWorkYear.Time.String()

	if err != nil {
		return nil, err
	}

	return &doctor, nil
}

func (h *DocTor) GetAllDoctors(ctx context.Context, page, limit int64, search string) (doctors []*entity.Doctor, err error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctor, serviceNameDoctorRepoPrefix+"Get all")
	span.SetAttributes(attribute.Key("GetAllDoctors").String(search))

	defer span.End()

	offset := limit * (page - 1)

	queryBuilder := h.db.Sq.Builder.Select(h.docTorSelectQueryPrefix()).From(doctorTableName)
	if search != "" {
		queryBuilder = queryBuilder.Where(fmt.Sprintf(`first_name ILIKE '%s' OR last_name ILIKE '%s'`, search+"%", search+"%"))
	}
	queryBuilder = queryBuilder.Limit(uint64(limit)).Offset(uint64(offset))
	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := h.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var doctor entity.Doctor
		var startWorkYear, endWorkYear, birthDate, updatedAt sql.NullTime
		err = rows.Scan(
			&doctor.Id,
			&doctor.Order,
			&doctor.FirstName,
			&doctor.LastName,
			&doctor.Gender,
			&doctor.PhoneNumber,
			&doctor.Email,
			&doctor.Password,
			&doctor.Address,
			&doctor.City,
			&doctor.Country,
			&doctor.Salary,
			&doctor.Bio,
			&startWorkYear,
			&endWorkYear,
			&doctor.WorkYears,
			&doctor.DepartmentId,
			&doctor.RoomNumber,
			&doctor.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		if updatedAt.Valid {
			doctor.UpdatedAt = updatedAt.Time
		}
		doctor.BirthDate = birthDate.Time.String()
		doctor.StartWorkDate = startWorkYear.Time.String()
		doctor.EndWorkDate = endWorkYear.Time.String()

		doctors = append(doctors, &doctor)
	}
	return doctors, nil
}

func (h *DocTor) UpdateDoctor(ctx context.Context, update *entity.Doctor) (*entity.Doctor, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctor, serviceNameDoctorRepoPrefix+"Update")
	span.SetAttributes(attribute.Key("UpdateDoctor").String(update.Id))

	defer span.End()

	data := map[string]any{
		"id":              update.Id,
		"doctor_order":    update.Order,
		"first_name":      update.FirstName,
		"last_name":       update.LastName,
		"gender":          update.Gender,
		"phone_number":    update.PhoneNumber,
		"email":           update.Email,
		"password":        update.Password,
		"address":         update.Address,
		"city":            update.City,
		"country":         update.Country,
		"salary":          update.Salary,
		"biography":       update.Bio,
		"start_work_date": update.StartWorkDate,
		"end_work_date":   update.EndWorkDate,
		"work_years":      update.WorkYears,
		"department_id":   update.DepartmentId,
		"room_number":     update.RoomNumber,
		"updated_at":      time.Now(),
	}

	query, args, err := h.db.Sq.Builder.Update(h.tableName).SetMap(data).
		Where(h.db.Sq.Equal("id", update.Id)).Suffix(fmt.Sprintf("RETURNING %s", h.docTorSelectQueryPrefix())).ToSql()

	if err != nil {

		return nil, h.db.ErrSQLBuild(err, h.tableName+" update")
	}
	var startWorkYear, endWorkYear sql.NullTime
	err = h.db.QueryRow(ctx, query, args...).Scan(
		&update.Id,
		&update.Order,
		&update.FirstName,
		&update.LastName,
		&update.Gender,
		&update.PhoneNumber,
		&update.Email,
		&update.Password,
		&update.Address,
		&update.City,
		&update.Country,
		&update.Salary,
		&update.Bio,
		&startWorkYear,
		&endWorkYear,
		&update.WorkYears,
		&update.DepartmentId,
		&update.RoomNumber,
		&update.CreatedAt,
		&update.UpdatedAt,
	)

	update.StartWorkDate = startWorkYear.Time.String()
	update.EndWorkDate = endWorkYear.Time.String()

	if err != nil {
		return nil, h.db.Error(err)
	}

	return update, nil
}

func (h *DocTor) DeleteDoctor(ctx context.Context, del *entity.GetReqStr) (bool, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctor, serviceNameDoctorRepoPrefix+"Delete")
	span.SetAttributes(attribute.Key("DeleteDoctor").String(del.Id))

	defer span.End()

	data := map[string]any{
		"deleted_at": time.Now(),
	}
	var args []interface{}
	var query string
	var err error
	if del.IsHardDeleted {
		query, args, err = h.db.Sq.Builder.Delete(h.tableName).From(h.tableName).Where(h.db.Sq.Equal("id", del.Id)).ToSql()
		if err != nil {
			return false, h.db.ErrSQLBuild(err, h.tableName+" delete")
		}
	} else {
		query, args, err = h.db.Sq.Builder.Update(h.tableName).SetMap(data).Where(h.db.Sq.Equal("id", del.Id)).ToSql()
		if err != nil {
			return false, h.db.ErrSQLBuild(err, h.tableName+" delete")
		}
	}
	_, err = h.db.Exec(ctx, query, args...)
	if err != nil {
		return false, h.db.Error(err)
	}
	return true, nil
}
