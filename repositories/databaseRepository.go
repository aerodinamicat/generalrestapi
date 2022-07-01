package repositories

import (
	"context"

	"github.com/aerodinamicat/generalrestapi/models"
)

type DatabaseRepository interface {
	CloseDatabaseConnection() error

	ListDatabasePersons(ctx context.Context, orderBy string, pageSize uint64, pageToken uint64) ([]*models.Person, error)
	GetDatabasePerson(ctx context.Context, personId string) (*models.Person, error)
	InsertDatabasePerson(ctx context.Context, person *models.Person) (*models.Person, error)
	UpdateDatabasePerson(ctx context.Context, person *models.Person) (*models.Person, error)
	DeleteDatabasePerson(ctx context.Context, personId string) error
}

var dbrImplementation DatabaseRepository

func SetDatabaseRepository(repository DatabaseRepository) {
	dbrImplementation = repository
}

func CloseDatabaseConnection() error {
	return dbrImplementation.CloseDatabaseConnection()
}

func ListDatabasePersons(ctx context.Context, orderBy string, pageSize uint64, pageToken uint64) ([]*models.Person, error) {
	return dbrImplementation.ListDatabasePersons(ctx, orderBy, pageSize, pageToken)
}
func GetDatabasePerson(ctx context.Context, personId string) (*models.Person, error) {
	return dbrImplementation.GetDatabasePerson(ctx, personId)
}

func InsertDatabasePerson(ctx context.Context, person *models.Person) (*models.Person, error) {
	return dbrImplementation.InsertDatabasePerson(ctx, person)
}

func UpdateDatabasePerson(ctx context.Context, person *models.Person) (*models.Person, error) {
	return dbrImplementation.UpdateDatabasePerson(ctx, person)
}

func DeleteDatabasePerson(ctx context.Context, personId string) error {
	return dbrImplementation.DeleteDatabasePerson(ctx, personId)
}
