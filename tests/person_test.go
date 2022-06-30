package tests

import (
	"testing"
	"time"

	"github.com/aerodinamicat/generalrestapi/models"
)

var testPersons []models.Person
var expectedResultsFromGetFullName []string
var expectedResultsFromGetFullNameToList []string
var expectedResultsFromGetInitialsWithoutDelimiter []string
var expectedResultsFromGetInitialsWithDelimiter []string
var expectedResultsFromGetAge []int

func FillTestPersonVariables() {
	testPersons = []models.Person{
		{
			FirstName:     "Aaaa",
			SecondName:    "Bbbb",
			FirstSurname:  "Cccc",
			SecondSurname: "Dddd",
			DateOfBirth:   time.Date(1980, 1, 1, 0, 10, 10, 0, time.Local),
		},
		{
			FirstName:     "Eeee",
			SecondName:    "Ffff",
			FirstSurname:  "Gggg",
			SecondSurname: "Hhhh",
			DateOfBirth:   time.Date(1990, 2, 2, 5, 20, 20, 0, time.Local),
		},
		{
			FirstName:     "Iiii",
			SecondName:    "Jjjj",
			FirstSurname:  "Kkkk",
			SecondSurname: "Llll",
			DateOfBirth:   time.Date(2000, 3, 3, 10, 30, 30, 0, time.Local),
		},
		{
			FirstName:     "Mmmm",
			SecondName:    "Nnnn",
			FirstSurname:  "Oooo",
			SecondSurname: "Pppp",
			DateOfBirth:   time.Date(2010, 4, 4, 15, 40, 40, 0, time.Local),
		},
		{
			FirstName:     "Qqqq",
			SecondName:    "Rrrr",
			FirstSurname:  "Ssss",
			SecondSurname: "Tttt",
			DateOfBirth:   time.Date(2020, 5, 5, 20, 50, 50, 0, time.Local),
		},
	}
	expectedResultsFromGetFullName = []string{
		"Aaaa Bbbb Cccc Dddd",
		"Eeee Ffff Gggg Hhhh",
		"Iiii Jjjj Kkkk Llll",
		"Mmmm Nnnn Oooo Pppp",
		"Qqqq Rrrr Ssss Tttt",
	}
	expectedResultsFromGetFullNameToList = []string{
		"Cccc Dddd, Aaaa Bbbb",
		"Gggg Hhhh, Eeee Ffff",
		"Kkkk Llll, Iiii Jjjj",
		"Oooo Pppp, Mmmm Nnnn",
		"Ssss Tttt, Qqqq Rrrr",
	}
	expectedResultsFromGetInitialsWithDelimiter = []string{"A.B.C.D.", "E.F.G.H.", "I.J.K.L.", "M.N.O.P.", "Q.R.S.T."}
	expectedResultsFromGetInitialsWithoutDelimiter = []string{"ABCD", "EFGH", "IJKL", "MNOP", "QRST"}
	expectedResultsFromGetAge = []int{42, 32, 22, 12, 2}
}

func TestGetFullName(t *testing.T) {
	var testResult string
	var expectedResult string

	FillTestPersonVariables()

	for index, item := range testPersons {
		testResult = item.GetFullName()
		expectedResult = expectedResultsFromGetFullName[index]

		if expectedResult != testResult {
			t.Errorf("GetFullName was incorrect, got '%s' expected '%s'.", testResult, expectedResult)
		}
	}
}

func TestGetFullNameToList(t *testing.T) {
	var testResult string
	var expectedResult string

	FillTestPersonVariables()

	for index, item := range testPersons {
		testResult = item.GetFullNameToList(",")
		expectedResult = expectedResultsFromGetFullNameToList[index]

		if expectedResult != testResult {
			t.Errorf("GetFullName was incorrect, got '%s' expected '%s'.", testResult, expectedResult)
		}
	}
}

func TestGetInitials(t *testing.T) {
	var testResult string
	var expectedResult string

	FillTestPersonVariables()

	for index, item := range testPersons {
		testResult = item.GetInitials(true)
		expectedResult = expectedResultsFromGetInitialsWithDelimiter[index]

		if expectedResult != testResult {
			t.Errorf("GetFullName was incorrect, got '%s' expected '%s'.", testResult, expectedResult)
		}

		testResult = item.GetInitials(false)
		expectedResult = expectedResultsFromGetInitialsWithoutDelimiter[index]

		if expectedResult != testResult {
			t.Errorf("GetFullName was incorrect, got '%s' expected '%s'.", testResult, expectedResult)
		}
	}
}

func TestGetAge(t *testing.T) {
	var testResult int
	var expectedResult int

	FillTestPersonVariables()

	for index, item := range testPersons {
		testResult = item.GetAge()
		expectedResult = expectedResultsFromGetAge[index]

		if expectedResult != testResult {
			t.Errorf("GetFullName was incorrect, got '%d' expected '%d'.", testResult, expectedResult)
		}
	}
}
