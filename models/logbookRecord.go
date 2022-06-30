package models

import "time"

/** Modelo 'LogBookRecord' (Registro en Cuaderno de Bitácora)
*
* Propósito: Representa el 'cuando' y el 'quien' registró la información en el sistema.
*	Se trata de una entidad fuerte.
*
* Propiedades:
*	- 'ReferenceId': De tipo 'string', representa el 'id' único de la entidad, u objeto, al que pertenece la información.
*	- 'RecordedAt': De tipo 'time.Time', representa el momento en el tiempo cuando tomó lugar el registro de la información.
		- Formato: "YY-MM-AA HH:MM:SS".
*	- 'RecordedByUserId': De tipo 'string', representa el identificador unívoco del usuario del sistema que registró la información.
*
* Comportamientos: Ninguno
*/

//* Definición  - Objeto 'LogbookRecord'
type LogbookRecord struct {
	ReferenceId      string    `json:"reference_id"`        //* Representa el 'id' único de la entidad, u objeto, referenciado.
	RecordedAt       time.Time `json:"recorded_at"`         //* Representa el momento en el tiempo cuando tomó lugar el registro de la información.
	RecordedByUserId string    `json:"recorded_by_user_id"` //* Representa el identificador unívoco del usuario del sistema que registró la información.
}
