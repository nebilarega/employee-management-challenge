package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID           int       `gorm:"primaryKey"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Gender       string    `json:"gender"`
	PhoneNo      string    `json:"phoneNo" gorm:"size:100;not null;"`
	Email        string    `json:"email" gorm:"size:100;not null;unique"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	Country      string    `json:"country"`
	City         string    `json:"city"`
	SubCity      string    `json:"subCity"`
	Region       string    `json:"region"`
	Woreda       string    `json:"woreda"`
	Zone         string    `json:"zone"`
	Kebele       string    `json:"kebele"`
	HouseNo      string    `json:"houseNo"`
	DepartmentID int       `json:"departmentId"`
	ImageName    *string    `json:"imageName"`
	Count        int       `json:"count"`
}
