/*
	nolint
*/

package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"crypto/sha1"
	"encoding/base64"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"time"

	graph_models "github.com/tensorsystems/employee-management-challenge/apps/core/pkg/graphql/graph/model"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/models"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	DepartmentRepository repository.DepartmentRepository
	EmployeeRepository   repository.EmployeeRepository
}

// WriteFile ...
func WriteFile(file io.Reader, fileName string) error {
	content, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return readErr
	}

	writeErr := ioutil.WriteFile("files/"+fileName, content, 0644)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

// RenameFile ...
func RenameFile(originalName string, newName string) error {
	return os.Rename("files/"+originalName, "files/"+newName)
}

// HashFileName ...
func HashFileName(name string) (fileName string, hashedFileName string, hash string, ext string) {
	s := strings.Split(name, ".")
	toHash := s[0] + time.Now().String()

	h := sha1.New()
	h.Write([]byte(toHash))

	fileName = s[0]
	hash = base64.URLEncoding.EncodeToString(h.Sum(nil))
	hashedFileName = s[0] + "_" + hash
	ext = s[1]

	return
}

// ConvertEntityToConnection ...
func GetPageInfo[R any](entities []R, count int, page models.PaginationInput) (*graph_models.PageInfo, int) {
	if len(entities) == 0 {
		return &graph_models.PageInfo{}, 0
	}

	pageInfo := graph_models.PageInfo{}
	totalPages := math.Ceil(float64(count) / float64(page.Size))
	pageInfo.TotalPages = int(totalPages)

	return &pageInfo, count
}
