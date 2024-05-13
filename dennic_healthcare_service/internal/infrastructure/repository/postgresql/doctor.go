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
		"birth_date":      req.BirthDate,
		"phone_number":    req.PhoneNumber,
		"email":           req.Email,
		"password":        req.Password,
		"address":         req.Address,
		"city":            req.City,
		"country":         req.Country,
		"salary":          req.Salary,
		"biography":       req.Bio,
		"start_work_date": req.StartWorkDate,
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
	if endWorkYear.Valid {
		req.EndWorkDate = endWorkYear.Time.String()
	}
	if err != nil {
		return nil, h.db.Error(err)
	}

	return req, nil
}

func (h *DocTor) GetDoctorById(ctx context.Context, get *entity.GetReqStr) (*entity.Doctor, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctor, serviceNameDoctorRepoPrefix+"Get")
	span.SetAttributes(attribute.Key(get.Field).String(get.Value))

	defer span.End()

	var doctor entity.Doctor
	var startWorkYear, endWorkYear, updatedAt sql.NullTime
	queryBuilder := h.db.Sq.Builder.Select(h.docTorSelectQueryPrefix()).From(doctorTableName)

	if !get.IsActive {
		queryBuilder = queryBuilder.Where("deleted_at IS NULL")
	}

	queryBuilder = queryBuilder.Where(h.db.Sq.Equal(get.Field, get.Value))

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
	if endWorkYear.Valid {
		doctor.EndWorkDate = endWorkYear.Time.String()
	}

	if err != nil {
		return nil, err
	}

	return &doctor, nil
}

func (h *DocTor) GetAllDoctors(ctx context.Context, all *entity.GetAll) (*entity.ListDoctors, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctor, serviceNameDoctorRepoPrefix+"Get all")
	span.SetAttributes(attribute.Key(all.Field).String(all.Value))

	defer span.End()

	offset := all.Limit * (all.Page - 1)

	queryBuilder := h.db.Sq.Builder.Select(h.docTorSelectQueryPrefix()).From(doctorTableName)
	if all.Field != "" {
		queryBuilder = queryBuilder.Where(fmt.Sprintf(`%s ILIKE '%s'`, all.Field, all.Value+"%"))
	}
	if all.OrderBy != "" {
		queryBuilder = queryBuilder.OrderBy(all.OrderBy)
	}
	countBuilder := h.db.Sq.Builder.Select("count(*)").From(departmentTableName)
	if !all.IsActive {
		queryBuilder = queryBuilder.Where("deleted_at IS NULL")
		countBuilder = countBuilder.Where("deleted_at IS NULL")
	}
	queryBuilder = queryBuilder.Limit(uint64(all.Limit)).Offset(uint64(offset))
	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, err
	}
	var doctors entity.ListDoctors
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
		if endWorkYear.Valid {
			doctor.EndWorkDate = endWorkYear.Time.String()
		}

		doctors.Doctors = append(doctors.Doctors, doctor)
	}
	var count int64
	queryCount, _, err := countBuilder.ToSql()
	err = h.db.QueryRow(ctx, queryCount).Scan(&count)
	if err != nil {
		return nil, h.db.Error(err)
	}
	doctors.Count = count
	return &doctors, nil
}

func (h *DocTor) UpdateDoctor(ctx context.Context, update *entity.Doctor) (*entity.Doctor, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctor, serviceNameDoctorRepoPrefix+"Update")
	span.SetAttributes(attribute.Key("UpdateDoctor").String(update.Id))

	defer span.End()

	data := map[string]any{
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
		"updated_at":      time.Now().Add(time.Hour * 5),
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
	if endWorkYear.Valid {
		update.EndWorkDate = endWorkYear.Time.String()
	}

	if err != nil {
		return nil, h.db.Error(err)
	}

	return update, nil
}

func (h *DocTor) DeleteDoctor(ctx context.Context, del *entity.GetReqStr) (bool, error) {

	ctx, span := otlp.Start(ctx, serviceNameDoctor, serviceNameDoctorRepoPrefix+"Delete")
	span.SetAttributes(attribute.Key("DeleteDoctor").String(del.Value))

	defer span.End()

	data := map[string]any{
		"deleted_at": time.Now().Add(time.Hour * 5),
	}
	var args []interface{}
	var query string
	var err error
	if del.IsActive {
		query, args, err = h.db.Sq.Builder.Delete(h.tableName).From(h.tableName).
			Where(h.db.Sq.And(h.db.Sq.Equal(del.Field, del.Value), h.db.Sq.Equal("deleted_at", nil))).ToSql()
		if err != nil {
			return false, h.db.ErrSQLBuild(err, h.tableName+" delete")
		}
	} else {
		query, args, err = h.db.Sq.Builder.Update(h.tableName).SetMap(data).
			Where(h.db.Sq.And(h.db.Sq.Equal(del.Field, del.Value), h.db.Sq.Equal("deleted_at", nil))).ToSql()
		if err != nil {
			return false, h.db.ErrSQLBuild(err, h.tableName+" delete")
		}
	}
	resp, err := h.db.Exec(ctx, query, args...)
	if err != nil {
		return false, h.db.Error(err)
	}
	if resp.RowsAffected() > 0 {
		return true, nil
	}

	return false, nil
}

func (h *DocTor) ListDoctorsByDepartmentId(ctx context.Context, in *entity.GetReqStrDep) (doctors []*entity.Doctor, err error) {
	ctx, span := otlp.Start(ctx, serviceNameDoctor, serviceNameDoctorRepoPrefix+"Get all by department_id")
	span.SetAttributes(attribute.Key("ListDoctorsByDepartmentId").String(in.Field))

	defer span.End()

	offset := in.Limit * (in.Page - 1)

	queryBuilder := h.db.Sq.Builder.Select(h.docTorSelectQueryPrefix()).From(doctorTableName)
	if in.Field != "" {
		queryBuilder = queryBuilder.Where(fmt.Sprintf(`%s ILIKE '%s'`, in.Field, in.Value+"%"))
	}
	if in.OrderBy != "" {
		queryBuilder = queryBuilder.OrderBy(in.OrderBy)
	}
	if !in.IsActive {
		queryBuilder = queryBuilder.Where("deleted_at IS NULL")
	}

	queryBuilder = queryBuilder.Where(h.db.Sq.Equal("department_id", in.DepartmentId)).Limit(uint64(in.Limit)).Offset(uint64(offset))
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
		if endWorkYear.Valid {
			doctor.EndWorkDate = endWorkYear.Time.String()
		}

		doctors = append(doctors, &doctor)
	}
	return doctors, nil
}
