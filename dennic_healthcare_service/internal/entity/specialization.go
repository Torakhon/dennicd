package entity

import "time"

type Specialization struct {
	ID           string
	Order        int32
	Name         string
	Description  string
	DepartmentId string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type ListSpecializations struct {
	Specializations []Specialization
}
