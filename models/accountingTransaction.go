package models

import (
	"math"
	"time"
)

/** Modelo 'Accounting tranaction' (Asiento contable)
*
* Propósito: Representa un asiento contable. Es de tipo cuadrático y responde al paradigma "juego de suma ceros".
*	Se trata de una entidad débil en existencia, ya que depende de la entidad 'company'.
*
* Propiedades:
*	- 'Id': De tipo 'string', para identificar de forma unívoca e inequívoca al usuario.
*	- 'DateOfRegistry': De tipo 'time.Time', representa el momento en el tiempo en el que se registra el asiento
*		contable, y los consecuentes movimientos contables.
*	- 'PostJournalPosition': De tipo 'string', representa el orden correlativo máximo en el momento que fué registrado.
*		Es de tipo 'string' y no 'int', o numérico, porque no vamos a operar con él. Únicamente llevaremos la cuenta
*	- 'CompanyId': De tipo 'string', representa  el 'id' de la empresa asociada al asiento, ya que se trata de una entidad débil
*		cero o no.
*
* Comportamientos:
*	- 'CalculateZeroSum(aRecords)': devuelve si el asiento contable esta cuadrado, o no.
*		- Parámetros de entrada:
*			- 'aRecords': De tipo '[]*AccountingRecord', representa una colección de movimientos contables cuyo
*				'transactionId' corresponde a este 'Id'.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'bool', representa si la diferencia entre las sumas de los movimientos contables
*				cuya naturaleza 'debe' y 'haber' es cero (son iguales).
 */

//* Definición  - Objeto 'AccountingRecord': representa un movimiento contable de un asiento contable.
type AccountingTransaction struct {
	Id                  string    `json:"id"`               //* Representa el identificador único.
	PostJournalPosition string    `json:"number"`           //* Representa la cantidad máxima, mas uno, de asientos en el momento que fué registrado.
	DateOfRegistry      time.Time `json:"date_of_registry"` //* Representa el momento en el tiempo donde tomó lugar su registro.
	CompanyId           string    `json:"company_id"`       //* Representa el identificado único de la empresa asociada.
}

/** Definición - Método 'CalculateZeroSum': representa el nombre completo de la persona.
*		Concatena, si existen, sus nombres y apellidos.
*		- Parámetros de entrada:
*			- 'aRecords': De tipo '[]*AccountingRecord', representa una colección de movimientos con 'transactionId'='Id'.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'bool', representa si la diferencia entre las sumas de los movimientos contables
*				'debe' y 'haber' son iguales.
*			- 'Sin nombre' De tipo 'float64' representa la diferencia entre las sumas de los movimientos contables
*				'debe' y 'haber'.
 */
func (at *AccountingTransaction) CalculateZeroSum(aRecords []*AccountingRecord) (bool, float64) {
	var debitSum float64
	var creditSum float64

	for _, item := range aRecords {
		if item.Nature == NATURE_CREDIT {
			creditSum += item.Amount
		}
		if item.Nature == NATURE_DEBIT {
			debitSum += item.Amount
		}
	}

	return creditSum == debitSum, float64(math.Trunc(math.Round((debitSum-creditSum)*100))) / 100
}
