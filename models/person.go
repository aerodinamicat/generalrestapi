package models

import (
	"strings"
	"time"
)

type Person struct {
	Id            string    `json:"id"`
	FirstName     string    `json:"first_name"`
	SecondName    string    `json:"second_name"`
	FirstSurname  string    `json:"first_surname"`
	SecondSurname string    `json:"second_surname"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	Gender        string    `json:"gender"`
	Alias         string    `json:"alias"`

	CreatedAt        time.Time `json:"created_at"`
	CreatedByUserId  string    `json:"created_by_user_id"`
	ModifiedAt       time.Time `json:"modified_at"`
	ModifiedByUserId string    `json:"modified_by_user_id"`
}

func (person *Person) GetFullName() string {
	result := ""

	result = person.FirstName
	if person.SecondName != "" {
		result = result + " " + person.SecondName
	}

	result = result + " " + person.FirstSurname
	if person.SecondSurname != "" {
		result = result + " " + person.SecondSurname
	}

	return result
}

func (person *Person) GetFullNameToList(delimiterCharacter string) string {
	result := ""

	result = person.FirstSurname
	if person.SecondSurname != "" {
		result = result + " " + person.SecondSurname
	}

	if delimiterCharacter == "" {
		delimiterCharacter = ","
	}
	result = result + delimiterCharacter + " " + person.FirstName
	if person.SecondName != "" {
		result = result + " " + person.SecondName
	}

	return result
}

func (person *Person) GetInitials(withDelimiterCharacter bool) string {
	delimiterCharacter := "."

	result := ""
	personData := []string{person.FirstName, person.SecondName, person.FirstSurname, person.SecondSurname}
	for _, data := range personData {
		if data != "" {
			result = result + strings.ToUpper(data[0:1])
			if withDelimiterCharacter == true {
				result = result + delimiterCharacter
			}
		}
	}

	return result
}

func (person *Person) GetAge() int {
	return time.Now().Year() - person.DateOfBirth.Year()
}
