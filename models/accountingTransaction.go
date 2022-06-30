package models

import (
	"fmt"
	"math"
	"time"
)

/** Modelo 'Accounting tranaction' (Asiento contable)
*
* Propósito: Representa un asiento contable. Es de tipo cuadrático y responde al paradigma "juego de suma ceros".
*	Se trata de una entidad fuerte.
*
* Propiedades:
*	- 'Id': De tipo 'string', para identificar de forma unívoca e inequívoca al usuario.
*	- 'DateOfRegistry': De tipo 'time.Time', representa el momento en el tiempo en el que se registra el asiento
*		contable, y los consecuentes movimientos contables.
*	- 'ZeroSum': De tipo 'bool', representa si la diferencia de la suma de los importes de los movimientos contables es
*		cero o no.
*
* Comportamientos:
*	- 'CalculateZeroSum(aRecords)': devuelve si el asiento contable esta cuadrado, o no.
*		- Parámetros de entrada:
*			- 'aRecords': De tipo 'map[int]*AccountingRecord', representa una colección de movimientos contables cuyo
*				'transactionId' corresponde a este 'Id'.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'bool', representa si la diferencia entre las sumas de los movimientos contables
*				cuya naturaleza 'debe' y 'haber' es cero (son iguales).
 */

//* Definición  - Objeto 'AccountingRecord': representa un movimiento contable de un asiento contable.
type AccountingTransaction struct {
	Id             string    `json:"id"`               //* Representa el identificador único.
	DateOfRegistry time.Time `json:"date_of_registry"` //* Representa el momento en el tiempo donde tomó lugar su registro.
	ZeroSum        bool      `json:"zero_sum"`         //* Representa si se trata de un asiento contable cuadrado, o no.
}

/** Definición - Método 'CalculateZeroSum': representa el nombre completo de la persona.
*		Concatena, si existen, sus nombres y apellidos.
*		- Parámetros de entrada:
*			- 'aRecords': De tipo '[]*AccountingRecord', representa una colección de movimientos con 'transactionId'='Id'.
*		- Parámetros de salida:
*			- 'Sin nombre': De tipo 'bool', representa si la diferencia entre las sumas de los movimientos contables
*				'debe' y 'haber' son iguales.
*			- 'Sin nombre' De tipo 'float32' representa la diferencia entre las sumas de los movimientos contables
*				'debe' y 'haber'.
 */
func (at *AccountingTransaction) CalculateZeroSum(aRecords []*AccountingRecord) (bool, float64) {
	var debitSum float64
	var creditSum float64

	for _, item := range aRecords {
		if item.Type == TYPE_CREDIT {
			creditSum += item.Amount
		}
		if item.Type == TYPE_DEBIT {
			debitSum += item.Amount
		}
	}

	at.ZeroSum = creditSum == debitSum
	fmt.Printf(
		"%f\t%f\t%f\t%f\t%f\n",
		debitSum-creditSum,
		(debitSum-creditSum)*100,
		math.Trunc(math.Round((debitSum-creditSum)*100)),
		float64(math.Trunc(math.Round((debitSum-creditSum)*100))),
		float64(math.Trunc(math.Round((debitSum-creditSum)*100)))/100)

	return creditSum == debitSum, float64(math.Trunc(math.Round((debitSum-creditSum)*100))) / 100
}
