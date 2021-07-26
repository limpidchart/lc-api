package testutils

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
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

func RandomUUID(t *testing.T) uuid.UUID {
	t.Helper()

	u, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("unable to generate random UUID: %s", err)
	}

	return u
}
