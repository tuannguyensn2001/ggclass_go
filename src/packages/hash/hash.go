package hash

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

func Hash(hash string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(hash), 14)

	result := string(bytes)

	return result, err
}

func Compare(hashString string, origin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(origin), []byte(hashString))
	return err == nil
}

func CompareWithContext(ctx context.Context, hashString string, origin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(origin), []byte(hashString))

	return err == nil
}
