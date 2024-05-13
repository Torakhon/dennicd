package entity

import "time"

type GetReqStr struct {
	Id            string
	IsActive      bool
	IsHardDeleted bool
}

type Status struct {
}

type Doctor struct {
	Id            string
	Order         int32
	FirstName     string
	LastName      string
	Gender        string
	BirthDate     string
	PhoneNumber   string
	Email         string
	Address       string
	City          string
	Country       string
	Salary        float32
	Bio           string
	StartWorkDate string
	EndWorkDate   string
	WorkYears     int32
	DepartmentId  string
	RoomNumber    int32
	Password      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
