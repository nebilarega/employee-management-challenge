package models

import (
	"gorm.io/gorm"
)

// PaginationInput ...
type PaginationInput struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// Paginate ...
func Paginate(p *PaginationInput) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := p.Page
		if page == 0 {
			page = 1
		}

		pageSize := p.Size
		switch {
		case pageSize > 1000:
			pageSize = 1000
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
