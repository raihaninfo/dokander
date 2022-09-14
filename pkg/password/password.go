package password

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

var Salt string = "ThIsIsSecretFoRPasSwORForSecurity"

// Generate password hash if password and confirm password is same
func GenerateHash(password, confirmPassword string) (string, error) {
	var hashP string
	if password == confirmPassword {
		HashPassword, err := bcrypt.GenerateFromPassword([]byte(password+Salt), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		hashP = string(HashPassword)
	}
	return hashP, nil
}

// compare password from hash
func CompareHash(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+Salt))
	if err != nil {
		return false, err
	}
	return true, nil
}
