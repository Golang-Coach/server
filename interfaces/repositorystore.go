package interfaces

import (
	"github.com/Golang-Coach/server/models"
)

type IRepositoryStore interface {
	FindByFullName(owner string, name string) (*models.RepositoryInfo, error)
	FindPackageWithinLimit(query string, skip int, limit int) (*[]models.RepositoryInfo, error)
}
