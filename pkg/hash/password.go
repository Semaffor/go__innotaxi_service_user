package hash

import (
	"crypto/sha256"
	"fmt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type SHA256Hasher struct {
	salt string
}

func NewSHA256Hasher(salt string) *SHA256Hasher {
	return &SHA256Hasher{salt: salt}
}

// Hash creates SHA1 hash of given password.
func (h *SHA256Hasher) Hash(password string) (string, error) {
	hash := sha256.New()
	hash.BlockSize()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}
	hashed := fmt.Sprintf("%x", hash.Sum([]byte(h.salt)))

	return hashed[:40], nil
}
