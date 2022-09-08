package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/graphql/graph/generated"
	graph_models "github.com/tensorsystems/employee-management-challenge/apps/core/pkg/graphql/graph/model"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/models"
	deepCopy "github.com/ulule/deepcopier"
)

// CreateEmployee is the resolver for the createEmployee field.
func (r *mutationResolver) CreateEmployee(ctx context.Context, input graph_models.CreateEmployeeInput, pic *graphql.Upload) (*models.Employee, error) {
	var employee models.Employee
	deepCopy.Copy(&input).To(&employee)

	if pic != nil {
		err := WriteFile(pic.File, pic.Filename)
		if err != nil {
			return nil, err
		}

		employee.ImageName = &pic.Filename
	}

	if err := r.EmployeeRepository.Save(&employee); err != nil {
		return nil, err
	}

	return &employee, nil
}

// CreateDepartment is the resolver for the createDepartment field.
func (r *mutationResolver) CreateDepartment(ctx context.Context, input graph_models.CreateDepartmentInput, pic *graphql.Upload) (*models.Department, error) {
	var department models.Department
	deepCopy.Copy(&input).To(&department)

	if pic != nil {
		err := WriteFile(pic.File, pic.Filename)
		if err != nil {
			return nil, err
		}

		department.ImageName = &pic.Filename
	}

	if err := r.DepartmentRepository.Save(&department); err != nil {
		return nil, err
	}

	return &department, nil
}

// UpdateEmployee is the resolver for the updateEmployee field.
func (r *mutationResolver) UpdateEmployee(ctx context.Context, input graph_models.UpdateEmployeeInput, pic *graphql.Upload) (*models.Employee, error) {
	var employee models.Employee
	deepCopy.Copy(&input).To(&employee)

	if pic != nil {
		err := WriteFile(pic.File, pic.Filename)
		if err != nil {
			return nil, err
		}

		employee.ImageName = &pic.Filename
	}

	if err := r.EmployeeRepository.Update(&employee); err != nil {
		return nil, err
	}

	return &employee, nil
}

// UpdateDepartment is the resolver for the updateDepartment field.
func (r *mutationResolver) UpdateDepartment(ctx context.Context, input graph_models.UpdateDepartmentInput, pic *graphql.Upload) (*models.Department, error) {
	var department models.Department
	deepCopy.Copy(&input).To(&department)

	if pic != nil {
		err := WriteFile(pic.File, pic.Filename)
		if err != nil {
			return nil, err
		}

		department.ImageName = &pic.Filename
	}

	if err := r.DepartmentRepository.Update(&department); err != nil {
		return nil, err
	}

	return &department, nil
}

// DeleteEmployee is the resolver for the deleteEmployee field.
func (r *mutationResolver) DeleteEmployee(ctx context.Context, id int) (*models.Employee, error) {
	var employee models.Employee

	if err := r.EmployeeRepository.Delete(&employee, id); err != nil {
		return nil, err
	}

	return &employee, nil
}

// DeleteDepartment is the resolver for the deleteDepartment field.
func (r *mutationResolver) DeleteDepartment(ctx context.Context, id int) (*models.Department, error) {
	var department models.Department

	if err := r.DepartmentRepository.Delete(&department, id); err != nil {
		return nil, err
	}

	return &department, nil
}

// GetHealthCheck is the resolver for the getHealthCheck field.
func (r *queryResolver) GetHealthCheck(ctx context.Context) (*graph_models.HealthCheckReport, error) {
	panic(fmt.Errorf("not implemented: GetHealthCheck - getHealthCheck"))
}

// CurrentDateTime is the resolver for the currentDateTime field.
func (r *queryResolver) CurrentDateTime(ctx context.Context) (*time.Time, error) {
	t := time.Now()
	return &t, nil
}

// Employee is the resolver for the employee field.
func (r *queryResolver) Employee(ctx context.Context, id int) (*models.Employee, error) {
	var employee models.Employee

	if err := r.EmployeeRepository.Get(&employee, id); err != nil {
		return nil, err
	}

	return &employee, nil
}

// Employees is the resolver for the employees field.
func (r *queryResolver) Employees(ctx context.Context, searchTerm *string, page models.PaginationInput) (*graph_models.EmployeeConnection, error) {
	result, count, err := r.EmployeeRepository.GetAll(page, searchTerm)
	if err != nil {
		return nil, err
	}

	edges := make([]*graph_models.EmployeeEdge, len(result))

	for i, entity := range result {
		e := entity

		edges[i] = &graph_models.EmployeeEdge{
			Node: &e,
		}
	}

	pageInfo, totalCount := GetPageInfo(result, count, page)
	return &graph_models.EmployeeConnection{PageInfo: pageInfo, Edges: edges, TotalCount: totalCount}, nil
}

// Department is the resolver for the department field.
func (r *queryResolver) Department(ctx context.Context, id int) (*models.Department, error) {
	var department models.Department

	if err := r.DepartmentRepository.Get(&department, id); err != nil {
		return nil, err
	}

	return &department, nil
}

// Departments is the resolver for the departments field.
func (r *queryResolver) Departments(ctx context.Context, searchTerm *string, page models.PaginationInput) (*graph_models.DepartmentConnection, error) {
	result, count, err := r.DepartmentRepository.GetAll(page, searchTerm)
	if err != nil {
		return nil, err
	}

	edges := make([]*graph_models.DepartmentEdge, len(result))

	for i, entity := range result {
		e := entity

		edges[i] = &graph_models.DepartmentEdge{
			Node: &e,
		}
	}

	pageInfo, totalCount := GetPageInfo(result, count, page)
	return &graph_models.DepartmentConnection{PageInfo: pageInfo, Edges: edges, TotalCount: totalCount}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
