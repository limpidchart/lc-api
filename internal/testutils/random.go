package testutils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(size int) string {
	//nolint: gosec
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	res := make([]byte, size)
	for i := range res {
		res[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(res)
}
