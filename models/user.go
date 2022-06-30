package models

type User struct {
	Id                       string `json:"id"`
	Email                    string `json:"email"`
	Password                 string `json:"password"`
	PasswordHint             string `json:"password_hint"`
	PasswordRecoveryQuestion string `json:"password_recovery_question"`
	PasswordRecoveryAnswer   string `json:"password_recovery_answer"`
}
