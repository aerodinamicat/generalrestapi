package models

import (
	"strings"
	"time"
)

/** Modelo 'Person' (Persona)
*
* Propósito: Representa la información guardada sobre una persona.
*	Se trata de una entidad fuerte.
*
* Propiedades:
*	- 'Id': De tipo 'string', para identificar de forma unívoca e inequívoca.
*	- 'FirstName': De tipo 'string', representa el nombre, o primer nombre, de la persona.
*	- 'SecondName': De tipo 'string', representa el sobrenombre, o segundo nombre (en caso de que el nombre sea compuesto), de la persona.
*	- 'FirstSurname': De tipo 'string', representa el apellido paterno, o primer apellido, de la persona.
*	- 'SecondSurname': De tipo 'string', representa el apellido materno, o segundo apellido (en el caso de que exista), de la persona.
*	- 'DateOfBirth': De tipo 'time.Time', representa el momento en el tiempo cuando tomó lugar el nacimiento de la persona.
*		- Formato: "YY-MM-AA HH:MM:SS".
*	- 'Gender': De tipo 'string', representa el género biológico binario de la persona.
*		- Valores admitidos:
*			- GENDER_FEMALE: 0 (Femenino).
*			- GENDER_MALE: 1 (Masculino).
*	- 'Alias': De tipo 'string', representa un pseudónimo, una forma mas popular o corta de nombrar a la persona.
*
* Comportamientos:
*	- 'GetFullName()': devuelve el nombre completo de la persona. Concatena, si existen, sus nombres y apellidos.
*		- Parámetros de entrada: Ninguno.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'string', representa el nombre completo de la persona.
*				- Formato: "FirstName+' '+SecondName+' '+FirstSurname'+' '+SecondSurname"
*	- 'GetFullNameToList(delimiterCharacter)': devuelve el nombre completo de la persona. Concatena todos sus nombres y apellidos.
*		- Parámetros de entrada:
*			- 'delimiterCharacter': De tipo 'string', representa el caracter que separa el nombre de los apellidos.
*				- Predeterminado: ','.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'string', representa el nombre completo de la persona.
*				- Formato: "FirstSurname+' '+SecondSurname+delimiterCharacter+' '+FirstName+' '+SecondName"
*	- 'GetInitials(withDelimiterCharacter)': devuelve el acrónimo del nombre completo de la persona. Concatena todos sus nombres y apellidos.
*		- Parámetros de entrada:
*			- 'withDelimiterCharacter': De tipo 'bool', representa si usar, o no, un carácter que separe cada letra del acrónimo resultante.
*				- Predeterminado: '.'.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'string' representa el acrónimo del nombre completo de la persona.
*	- 'GetAge()': devuelve la duración de tiempo, en años, que han transcurrido desde el momento de nacimiento de la persona hasta la actualidad.
*		- Parámetros de entrada: Ninguno.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'int', representa la diferencia, en años, del momento de nacimiento de la persona y el de la actualidad.
 */

//* Definición - Constantes - Ámbito global
const (
	GENDER_FEMALE = "0" //* Representa el género 'female', o femenino.
	GENDER_MALE   = "1" //* Representa el género 'male', o masculino
)

//* Definición  - Objeto 'Person': representa la información guardada sobre una persona
type Person struct {
	Id            string    `json:"id"`             //* Representa el identificador único.
	FirstName     string    `json:"first_name"`     //* Representa el nombre, o primer nombre.
	SecondName    string    `json:"second_name"`    //* Representa el sobrenombre, o segundo nombre (en caso de que el nombre sea compuesto).
	FirstSurname  string    `json:"first_surname"`  //* Representa el apellido paterno, o primer apellido.
	SecondSurname string    `json:"second_surname"` //* Representa el apellido materno, o segundo apellido (en el caso de que exista).
	DateOfBirth   time.Time `json:"date_of_birth"`  //* Representa el momento en el tiempo cuando tomó lugar el nacimiento.
	Gender        string    `json:"gender"`         //* Representa el género biológico binario.
	Alias         string    `json:"alias"`          //* Representa un pseudónimo, una forma mas popular o corta de nombrar.
}

/** Definición - Método 'GetFullName': representa el nombre completo de la persona. Concatena, si existen, sus nombres y apellidos.
*		- Parámetros de entrada: Ninguno.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'string', representa el nombre completo de la persona.
*				- Formato: "FirstName+' '+SecondName+' '+FirstSurname'+' '+SecondSurname"
 */
func (p *Person) GetFullName() string {
	result := ""

	result = p.FirstName
	if p.SecondName != "" {
		result = result + " " + p.SecondName
	}

	result = result + " " + p.FirstSurname
	if p.SecondSurname != "" {
		result = result + " " + p.SecondSurname
	}

	return result
}

/** Definición - Método 'GetFullNameToList': representa el nombre completo de la persona. Concatena, si existen, sus nombres y apellidos.
*		- Parámetros de entrada:
*			- 'delimiterCharacter': De tipo 'string', representa el caracter que separa el nombre de los apellidos.
*				- Predeterminado: ','.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'string', representa el nombre completo de la persona.
*				- Formato: "FirstSurname+' '+SecondSurname+delimiterCharacter+' '+FirstName+' '+SecondName"
 */
func (p *Person) GetFullNameToList(delimiterCharacter string) string {
	result := ""

	result = p.FirstSurname
	if p.SecondSurname != "" {
		result = result + " " + p.SecondSurname
	}

	if delimiterCharacter == "" {
		delimiterCharacter = ","
	}
	result = result + delimiterCharacter + " " + p.FirstName
	if p.SecondName != "" {
		result = result + " " + p.SecondName
	}

	return result
}

/** Definición - Método 'GetInitials': representa el acrónimo del nombre completo de la persona. Concatena, si existen, sus nombres y apellidos, con o sin carácter delimitador
*		- Parámetros de entrada:
*			- 'withDelimiterCharacter': De tipo 'bool', representa si usar, o no, un carácter que separe cada letra del acrónimo resultante.
*				- Predeterminado: '.'.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'string' representa el acrónimo del nombre completo de la persona.
 */
func (p *Person) GetInitials(withDelimiterCharacter bool) string {
	delimiterCharacter := "."

	result := ""
	personData := []string{p.FirstName, p.SecondName, p.FirstSurname, p.SecondSurname}
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

/** Definición - Método 'GetAge': representa la cantidad de tiempo, en años, que han transcurrido desde el momento de nacimiento de la persona hasta la actualidad.
*		- Parámetros de entrada: Ninguno.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'int', representa la diferencia, en años, del momento de nacimiento de la persona y el de la actualidad.
 */
func (p *Person) GetAge() int {
	return time.Now().Year() - p.DateOfBirth.Year()
}
