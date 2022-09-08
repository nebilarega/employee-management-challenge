package repository

import (
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func ProvideEmployeeRepository(DB *gorm.DB) EmployeeRepository {
	return EmployeeRepository{DB: DB}
}

// Get ...
func (r *EmployeeRepository) Get(m *models.Employee, ID int) error {
	return r.DB.Where("id = ?", ID).Take(&m).Error
}

// GetAll ...
func (r *EmployeeRepository) GetAll(p models.PaginationInput, searchTerm *string) ([]models.Employee, int, error) {
	var result []models.Employee

	dbOp := r.DB.Scopes(models.Paginate(&p)).Select("*, count(*) OVER() AS count")

	if searchTerm != nil {
		dbOp.Where("first_name ILIKE ?", "%"+*searchTerm+"%")
	}

	dbOp.Order("id ASC").Find(&result)

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
func (r *EmployeeRepository) GetByName(m *models.Employee, name string) error {
	return r.DB.Where("title = ?", name).Take(&m).Error
}

// Save ...
func (r *EmployeeRepository) Save(m *models.Employee) error {
	return r.DB.Create(&m).Error
}

// Update ...
func (r *EmployeeRepository) Update(m *models.Employee) error {
	return r.DB.Updates(&m).Error
}

// Delete ...
func (r *EmployeeRepository) Delete(m *models.Employee, ID int) error {
	return r.DB.Where("id = ?", ID).Delete(&m).Error
}
