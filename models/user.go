package models

/** Modelo 'User' (Usuario)
*
* Propósito: Representa la información guardada sobre un usuario.
*	Se trata de una entidad débil en existencia, ya que depende de la entidad 'Person'.
*
* Propiedades:
*	- 'Id': De tipo 'string', para identificar de forma unívoca e inequívoca al usuario.
*	- 'Email': De tipo 'string', representa la dirección de  correo electrónico del usuario.
*	- 'Password': De tipo 'string', representa la contraseña, o clave secreta, que identifica la identidad del usuario.
*	- 'PasswordHint': De tipo 'string', representa un indicio, o pista, sobre la 'password' del usuario.
*	- 'PasswordQuestion': De tipo 'string', representa la pregunta a responder del proceso de 'recuperación de contraseña' del usuario.
*	- 'PasswordAnswer': De tipo '', representa la respuesta secreta para la pregunta del proceso de 'recuperación de contraseña' del usuario.
*	- 'PersonId': De tipo '', representa el 'id' de la persona asociada al usuario, ya que se trata de una entidad débil.
*
* Comportamientos: Ninguno
 */

//* Definición  - Objeto 'User': representa la información guardada sobre una persona
type User struct {
	Id                       string `json:"id"`                         //* Representa el identificador único.
	Email                    string `json:"email"`                      //* Representa la dirección de correo electrónico.
	Password                 string `json:"password"`                   //* Representa la contraseña, o palabra clave secreta.
	PasswordHint             string `json:"password_hint"`              //* Representa el indicio, o pista, de 'password'.
	PasswordRecoveryQuestion string `json:"password_recovery_question"` //* Representa la pregunta del proceso de recuperación de 'password'.
	PasswordRecoveryAnswer   string `json:"password_recovery_answer"`   //* Representa la respuesta a la pregunta del proceso de recuperación de 'password'.
	PersonId                 string `json:"person_id"`                  //* Representa el identificador único de la persona asociada.
}
