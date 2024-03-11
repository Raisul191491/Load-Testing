package generator

import (
	"math/rand"
)

type Profile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Generator(cnt int) []Profile {
	var samples []Profile
	for i := 0; i < cnt; i++ {
		profile := Profile{
			Username: generateRandomString(10),
			Email:    generateRandomEmail(),
			Password: generateRandomString(8),
		}
		samples = append(samples, profile)
	}

	return samples
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func generateRandomEmail() string {
	return generateRandomString(8) + "@" + generateRandomString(5) + ".com"
}
