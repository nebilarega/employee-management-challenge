package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	ID        int        `gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"unique"`
	Employees []Employee `json:"employees"`
	ImageName *string    `json:"imageName"`
	Count     int        `json:"count"`
}
