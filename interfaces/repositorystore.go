package interfaces

import (
	"github.com/Golang-Coach/server/models"
)

type IRepositoryStore interface {
	FindPackage(query interface{}) (*models.RepositoryInfo, error)
	FindPackageWithinLimit(query string, skip int, limit int) (*[]models.RepositoryInfo, error)
}