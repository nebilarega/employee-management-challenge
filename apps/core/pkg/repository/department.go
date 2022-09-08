package repository

import (
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/models"
	"gorm.io/gorm"
)

type DepartmentRepository struct {
	DB *gorm.DB
}

func ProvideDepartmentRepository(DB *gorm.DB) DepartmentRepository {
	return DepartmentRepository{DB: DB}
}

// Get ...
func (r *DepartmentRepository) Get(m *models.Department, ID int) error {
	return r.DB.Where("id = ?", ID).Preload("Employees").Take(&m).Error
}

// GetAll ...
func (r *DepartmentRepository) GetAll(p models.PaginationInput, searchTerm *string) ([]models.Department, int, error) {
	var result []models.Department

	dbOp := r.DB.Scopes(models.Paginate(&p)).Select("*, count(*) OVER() AS count")

	if searchTerm != nil {
		dbOp.Where("title ILIKE ?", "%"+*searchTerm+"%")
	}

	dbOp.Order("id ASC").Preload("Employees").Find(&result)

	var count int
	if len(result) > 0 {
		count = result[0].Count
	}

	if dbOp.Error != nil {
		return result, 0, dbOp.Error
	}

	return result, count, dbOp.Error
}

// Get ...
func (r *DepartmentRepository) GetByName(m *models.Department, name string) error {
	return r.DB.Where("title = ?", name).Take(&m).Error
}

// Save ...
func (r *DepartmentRepository) Save(m *models.Department) error {
	return r.DB.Create(&m).Error
}

// Update ...
func (r *DepartmentRepository) Update(m *models.Department) error {
	return r.DB.Updates(&m).Error
}

// Delete ...
func (r *DepartmentRepository) Delete(m *models.Department, ID int) error {
	return r.DB.Where("id = ?", ID).Delete(&m).Error
}
