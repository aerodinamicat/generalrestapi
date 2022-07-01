package servers

import (
	"context"

	"github.com/aerodinamicat/generalrestapi/models"
	pbmodels "github.com/aerodinamicat/generalrestapi/pbmodels"
	"github.com/aerodinamicat/generalrestapi/repositories"
)

type PersonServer struct {
	repository repositories.PersonServiceRepository
	protoModel pbmodels.UnimplementedPersonServiceServer
}

func NewPersonServerInstance(repository repositories.PersonServiceRepository) *PersonServer {
	return &PersonServer{
		repository: repository,
	}
}

func (srv *PersonServer) ListPersons(ctx context.Context, pageSize uint, pageToken string) ([]*models.Person, error) {
	return nil, nil
}
func (srv *PersonServer) GetPerson(ctx context.Context, personId string) (*models.Person, error) {
	return nil, nil
}

func (srv *PersonServer) CreatePerson(ctx context.Context, person *models.Person) (*models.Person, error) {
	return nil, nil
}

func (srv *PersonServer) UpdatePerson(ctx context.Context, person *models.Person) (*models.Person, error) {
	return nil, nil
}

func (srv *PersonServer) DeletePerson(ctx context.Context, personId string) error {
	return nil
}
