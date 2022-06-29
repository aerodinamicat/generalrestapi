package models

import (
	"fmt"
	"testing"
)

type TestPerson struct {
	test           Person
	expectedResult string
}

func TestGetFullName(t *testing.T) {
	tables := []TestPerson{}

	for i := 0; i < 5; i++ {
		/*tables[i] = &TestPerson{
			test: &Person{
				FirstName:     fmt.Sprintf("Nombre1%d", i),
				SecondName:    fmt.Sprintf("Nombre2%d", i),
				FirstSurname:  fmt.Sprintf("Apellido1%d", i),
				SecondSurname: fmt.Sprintf("Apellido2%d", i),
			},
			expectedResult: fmt.Sprintf("Nombre1%d Nombre2%d Apellido1%d Apellido 2%d", i, i, i, i),
		}*/

		fn := fmt.Sprintf("Nombre1%d", i)
		sn := fmt.Sprintf("Nombre2%d", i)
		fs := fmt.Sprintf("Apellido1%d", i)
		ss := fmt.Sprintf("Apellido2%d", i)
		person := Person{
			FirstName:     fn,
			SecondName:    sn,
			FirstSurname:  fs,
			SecondSurname: ss,
		}

		er := fmt.Sprintf("Nombre1%d Nombre2%d Apellido1%d Apellido 2%d", i+1, i+1, i+1, i+1)

		tables[i] = TestPerson{
			test:           person,
			expectedResult: er,
		}
	}

	for _, item := range tables {
		objectTest := item.test.GetFullName()

		if item.expectedResult != objectTest {
			t.Errorf("GetFullName was incorrect, got %s expected '%s'.", objectTest, item.expectedResult)
		}
	}
}
