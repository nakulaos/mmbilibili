package tools

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(userSalt, systemSalt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(userSalt+systemSalt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}
