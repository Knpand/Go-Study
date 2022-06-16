package modules

import(
	"golang.org/x/crypto/bcrypt"
)

func Hash_password(password string) string {
	new_password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(new_password)
}