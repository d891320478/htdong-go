package aesTest

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

const nonce string = "f3c838dd7b6746a0b5d6df69"

func AESGCMEncrypt(origin, seed string) (ciphertext string, err error) {
	seedByte, err := hex.DecodeString(seed)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(seedByte)
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonceByte, err := hex.DecodeString(nonce)
	if err != nil {
		return
	}
	ciphertext = base64.RawURLEncoding.EncodeToString(aesgcm.Seal(nil, nonceByte, []byte(origin), nil))
	return
}

func AESGCMDecrypt(ciphertextStr, seed string) (origin string, err error) {
	ciphertext, err := base64.RawURLEncoding.DecodeString(ciphertextStr)
	if err != nil {
		return
	}
	seedByte, err := hex.DecodeString(seed)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(seedByte)
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceByte, err := hex.DecodeString(nonce)
	if err != nil {
		return
	}
	originByte, err := aesgcm.Open(nil, nonceByte, ciphertext, nil)
	if err != nil {
		return
	}
	origin = string(originByte)
	return
}
