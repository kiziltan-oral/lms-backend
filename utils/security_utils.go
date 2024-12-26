package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"
)

func ComputeSHA256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func GenerateRandomNumeric(length int) string {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	charRunes := []rune("0123456789")
	randomRunes := make([]rune, length)
	for i := range randomRunes {
		randomRunes[i] = charRunes[randomGenerator.Intn(len(charRunes))]
	}
	return string(randomRunes)
}

func GenerateRandomAlphaNumeric(length int) string {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	charRunes := []rune("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789")
	randomRunes := make([]rune, length)
	for i := range randomRunes {
		randomRunes[i] = charRunes[randomGenerator.Intn(len(charRunes))]
	}
	return string(randomRunes)
}
