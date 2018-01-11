package interfaces

import (
	"github.com/Golang-Coach/server/models"
)

type IDataStore interface {
	FindPackage(query interface{}) (*models.Repository, error)
	FindPackageWithinLimit(query string, skip int, limit int) (*[]models.Repository, error)
}

