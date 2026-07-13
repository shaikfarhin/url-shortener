package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortCode(length int) string {

	rand.Seed(time.Now().UnixNano())

	short := make([]byte, length)

	for i := range short {
		short[i] = charset[rand.Intn(len(charset))]
	}

	return string(short)
}
