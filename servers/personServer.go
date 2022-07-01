package servers

import (
	"context"

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

func (srv *PersonServer) ListPersons(ctx context.Context, request *pbmodels.ListPersonsRequest) (*pbmodels.ListPersonsResponse, error) {
	return nil, nil
}
func (srv *PersonServer) GetPerson(ctx context.Context, request *pbmodels.GetPersonRequest) (*pbmodels.Person, error) {
	person, err := srv.repository.GetPerson(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &pbmodels.Person{
		Id:            person.Id,
		FirstName:     person.FirstName,
		SecondName:    person.SecondName,
		FirstSurname:  person.FirstSurname,
		SecondSurname: person.SecondSurname,
		DateOfBirth:   person.DateOfBirth.String(),
		Gender:        person.Gender,
		Alias:         person.Alias,
	}, nil
}

func (srv *PersonServer) CreatePerson(ctx context.Context, request *pbmodels.CreatePersonRequest) (*pbmodels.Person, error) {
	return nil, nil
}

func (srv *PersonServer) UpdatePerson(ctx context.Context, request *pbmodels.UpdatePersonRequest) (*pbmodels.Person, error) {
	return nil, nil
}

func (srv *PersonServer) DeletePerson(ctx context.Context, request *pbmodels.DeletePersonRequest) (*pbmodels.Empty, error) {
	return nil, nil
}
