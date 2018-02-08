package interfaces

import (
	"github.com/Golang-Coach/server/models"
)

type IRepositoryStore interface {
	FindId(id string) (*models.RepositoryInfo, error)
	FindPackageWithinLimit(query string, skip int, limit int) (*[]models.RepositoryInfo, error)
}
