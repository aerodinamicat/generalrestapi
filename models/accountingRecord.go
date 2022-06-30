package models

/** Modelo 'Accounting record' (Registro contable)
*
* Propósito: Representa un movimiento contable de un 'accounting transaction', o asiento contable.
*	Se trata de una entidad débil en existencia, ya que depende de la entidad 'accounting transaction'.
*
* Propiedades:
*	- 'Id': De tipo 'string', para identificar de forma unívoca e inequívoca al usuario.
*	- 'AccountNumber': De tipo 'string', representa la cuenta contable del plan general contable del movimiento
*		contable. Es de tipo 'string' y no 'int', o numérico, ya que no vamos a operar con él.
*		- Rango: Oscila entre '[1000:9999]'. No todos los valores, pertenecientes al intervalo señalado, son usados.
*	- 'SubaccountNumber': De tipo 'string', representa la subcuenta contable del movimiento contable. Es de
*		tipo 'string' y no 'int', o numérico, ya que no vamos a operar con él.
*		- Rango: Oscila entre ']99999:999999]'. Todos los valores deberían ser usados consecutivamente. Puede variar según
			la magnitud del sistema, particularmente ahora está expresada en el orden de 'cientos de miles'
*	- 'Concept': De tipo 'string', representa una breve descripción de la razón del registro del movimiento contable.
*	- 'Nature': De tipo 'string', representa la naturaleza del movimiento contable.
*		- Valores admitidos:
*			- TYPE_CREDIT: 0 (para 'credit' o haber).
*			- TYPE_DEBIT: 1 (para 'debit' o debe).
*	- 'Amount': De tipo 'float64', representa la cantidad, o importe, del movimiento contable.
*		- Formato: "X.XX". Recomendado, aunque podrían añadirse valores a partir de la coma dependiendo del sistema.
*	- 'AccountingTransactionId': De tipo 'string', representa el 'id' único del asiento contable que referencia el movimiento contable.
*
* Comportamientos: Ninguno
*/

//* Definición - Constantes - Ámbito global
const (
	NATURE_CREDIT = "0" //* Representa la naturaleza del tipo 'credit', o haber.
	NATURE_DEBIT  = "1" //* Representa la naturaleza del tipo 'debit', o debe
)

//* Definición - Objeto 'AccountingRecord': representa un movimiento contable de un asiento contable.
type AccountingRecord struct {
	Id                      string  `json:"id"`                        //* Representa el identificador único.
	AccountNumber           uint16  `json:"account_number"`            //* Representa el número de la cuenta contable del plan general contable.
	SubaccountNumber        uint16  `json:"subaccount_number"`         //* Representa el número de la subcuenta contable del sistema de registro.
	Concept                 string  `json:"concept"`                   //* Representa el concepto, o breve descripción, de la razón para su registro.
	Nature                  string  `json:"type"`                      //* Representa la naturaleza del movimiento: 'debe' o 'haber'.
	Amount                  float64 `json:"amount"`                    //* Representa la cantidad, o importe.
	AccountingTransactionId string  `json:"accounting_transaction_id"` //* Representa el identificador único del asiento contable asociado.
}
