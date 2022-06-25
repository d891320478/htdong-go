package pbkdf2Test

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func Test(key string, salt []byte) []byte {
	return pbkdf2.Key([]byte(key), salt, 100000, 128, sha256.New)
}
