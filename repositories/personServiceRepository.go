package repositories

import (
	"context"

	"github.com/aerodinamicat/generalrestapi/models"
)

type PersonServiceRepository interface {
	ListPersons(ctx context.Context, pageSize uint, pageToken string) ([]*models.Person, error)
	GetPerson(ctx context.Context, personId string) (*models.Person, error)
	CreatePerson(ctx context.Context, person *models.Person) (*models.Person, error)
	UpdatePerson(ctx context.Context, person *models.Person) (*models.Person, error)
	DeletePerson(ctx context.Context, personId string) error
}

var psrImplementation PersonServiceRepository

func SetPersonServiceRepository(psr PersonServiceRepository) {
	psrImplementation = psr
}

func ListPersons(ctx context.Context, pageSize uint, pageToken string) ([]*models.Person, error) {
	return psrImplementation.ListPersons(ctx, pageSize, pageToken)
}
func GetPerson(ctx context.Context, personId string) (*models.Person, error) {
	return psrImplementation.GetPerson(ctx, personId)
}
func CreatePerson(ctx context.Context, person *models.Person) (*models.Person, error) {
	return psrImplementation.CreatePerson(ctx, person)
}
func UpdatePerson(ctx context.Context, person *models.Person) (*models.Person, error) {
	return psrImplementation.UpdatePerson(ctx, person)
}
func DeletePerson(ctx context.Context, personId string) error {
	return psrImplementation.DeletePerson(ctx, personId)
}
