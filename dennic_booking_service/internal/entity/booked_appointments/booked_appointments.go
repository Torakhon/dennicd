package booked_appointments

import (
	"github.com/rickb777/date"
	"time"
)

type Appointment struct {
	Id              int64
	DepartmentId    string
	DoctorId        string
	PatientId       string
	AppointmentDate date.Date
	AppointmentTime time.Time
	Duration        int64
	Key             string
	ExpiresAt       time.Time
	PatientStatus   bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type AppointmentsType struct {
	Count        int64
	Appointments []*Appointment
}

type CreateAppointment struct {
	DepartmentId    string
	DoctorId        string
	PatientId       string
	AppointmentDate date.Date
	AppointmentTime time.Time
	Duration        int64
	Key             string
	ExpiresAt       time.Time
	PatientStatus   bool
}

type UpdateAppointment struct {
	Field           string
	Value           string
	AppointmentDate date.Date
	AppointmentTime time.Time
	Duration        int64
	Key             string
	ExpiresAt       time.Time
	PatientStatus   bool
}

type GetAllAppointment struct {
	Page         uint64
	Limit        uint64
	DeleteStatus bool
	Field        string
	Value        string
	OrderBy      string
}

type FieldValueReq struct {
	Field        string
	Value        string
	DeleteStatus bool
}

type StatusRes struct {
	Status bool
}
