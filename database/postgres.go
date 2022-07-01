package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/aerodinamicat/generalrestapi/models"
)

type PostgresDatabaseRepository struct {
	db *sql.DB
}

func NewPostgresDatabaseRepository(databaseURL string) (*PostgresDatabaseRepository, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	return &PostgresDatabaseRepository{
		db: db,
	}, nil
}

func (pgr *PostgresDatabaseRepository) CloseDatabaseConnection() error {
	return pgr.db.Close()
}
func (pgr *PostgresDatabaseRepository) ListDatabasePersons(ctx context.Context, orderBy string, pageSize uint64, pageToken uint64) ([]*models.Person, error) {
	querySentence := `
		SELECT
		id,
		first_name,
		second_name,
		first_surname,
		second_surname,
		date_of_birth,
		gender,
		alias
		FROM persons ORDER_BY $1 LIMIT $2 OFFSET $3
	`
	resultQueryRows, err := pgr.db.QueryContext(ctx, querySentence, orderBy, pageSize, pageToken)
	defer func() {
		if err = resultQueryRows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var persons []*models.Person
	for resultQueryRows.Next() {
		var p = models.Person{}

		err = resultQueryRows.Scan(&p.Id, &p.FirstName, &p.SecondName, &p.FirstSurname, &p.SecondSurname, &p.DateOfBirth, &p.Gender, &p.Alias)
		if err == nil {
			persons = append(persons, &p)
		}
	}
	if err = resultQueryRows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}
func (pgr *PostgresDatabaseRepository) GetDatabasePerson(ctx context.Context, personId string) (*models.Person, error) {
	querySentence := `
		SELECT
		id,
		first_name,
		second_name,
		first_surname,
		second_surname,
		date_of_birth,
		gender,
		alias
		FROM persons WHERE id = $1
	`
	resultQueryRows, err := pgr.db.QueryContext(ctx, querySentence, personId)
	defer func() {
		if err = resultQueryRows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var p = models.Person{}
	for resultQueryRows.Next() {
		err = resultQueryRows.Scan(&p.Id, &p.FirstName, &p.SecondName, &p.FirstSurname, &p.SecondSurname, &p.DateOfBirth, &p.Gender, &p.Alias)
		if err != nil {
			return &p, nil
		}
	}

	return &p, nil
}
func (pgr *PostgresDatabaseRepository) InsertDatabasePerson(ctx context.Context, p *models.Person) (*models.Person, error) {
	/** A ésta altura, antes de agregar un nuevo registro a la DB, se podría consultar si existe ya un objeto en la DB con todas las
	* propiedades igual al objeto que se quiere añadir; y si es así, no realizar ninguna acción en la DB y avisar de un error de duplicidad.
	 */

	querySentence := `
		INSERT INTO persons
		id,
		first_name,
		second_name,
		first_surname,
		second_surname,
		date_of_birth,
		gender,
		alias
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := pgr.db.ExecContext(
		ctx,
		querySentence,
		p.Id,
		p.FirstName,
		p.SecondName,
		p.FirstSurname,
		p.SecondSurname,
		p.DateOfBirth.String(),
		p.Gender,
		p.Alias,
	)

	return p, err
}
func (pgr *PostgresDatabaseRepository) UpdateDatabasePerson(ctx context.Context, person *models.Person) (*models.Person, error) {
	/** A ésta altura, antes de actualizar el registro de la DB, se podría consultar el actual objeto en la DB a fin de evitar errores, como por
	* ejemplo, que no exista ya el registro en la DB, o que los nuevos datos lleven consigo un campo en blanco borrado accidentalmente.
	 */

	querySentence := `
		UPDATE persons SET
		first_name = $1,
		second_name = $2,
		first_surname = $3,
		second_surname = $4,
		date_of_birth = $5,
		gender = $6,
		alias = $7
		WHERE id = $8
	`
	_, err := pgr.db.ExecContext(
		ctx,
		querySentence,
		person.FirstName,
		person.SecondName,
		person.FirstSurname,
		person.SecondSurname,
		person.DateOfBirth.String(),
		person.Gender,
		person.Alias,
		person.Id,
	)

	return person, err
}
func (pgr *PostgresDatabaseRepository) DeleteDatabasePerson(ctx context.Context, personId string) error {
	/** A ésta altura, antes de borrar el registro de la DB, se podría consultar el actual objeto en la DB a fin de evitar errores, como por
	* ejemplo, que no exista ya el registro en la DB.
	 */
	querySentence := `
		DELETE FROM persons
		WHERE id = $1
	`
	_, err := pgr.db.ExecContext(ctx, querySentence, personId)

	return err
}
