package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

func GenerateRandomSalt() []byte {
	var salt = make([]byte, 15)
	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func HashPassword(password string, salt []byte) (string, error) {
	if len(password) < 8 {
		return "", errors.New("şifre 8 karakterden kısa olamaz")
	}

	var passwordBytes = []byte(password)
	passwordBytes = append(passwordBytes, salt...)

	var sha512Hasher = sha512.New()
	sha512Hasher.Write(passwordBytes)

	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex, nil
}

func CheckPassword(hashedPassword, givenPassword string, salt []byte) bool {
	givenPasswordHash, _ := HashPassword(givenPassword, salt)
	return hashedPassword == givenPasswordHash
}
