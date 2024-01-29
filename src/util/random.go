package util

import (
	"math/rand"
	"time"
)

// generateRandomString generates a random string of the specified length
func GenerateRandomString(length int) string {
	s := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(s)
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}

	return string(b)
}
