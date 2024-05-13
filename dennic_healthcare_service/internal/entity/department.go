package entity

import "time"

type Department struct {
	Id          string
	Order       int32
	Name        string
	Description string
	ImageUrl    string
	FloorNumber int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type GetReqInt struct {
	Id            int32
	IsActive      bool
	IsHardDeleted bool
}

type GetAll struct {
	Page   int64
	Limit  int64
	Search string
}

type ListDepartments struct {
	Count       int64
	Departments []Department
}
