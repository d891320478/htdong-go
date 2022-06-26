package main

import (
	"fmt"
	"os"

	"github.com/htdong/gotest/src/encryption-algorithm/aesTest"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func main() {
	origin := "abcdefghijklmn"
	seed := "841daf3185844b9b98c3e4b5daa783be841daf3185844b9b98c3e4b5daa783be"
	ciphertext, _ := aesTest.AESGCMEncrypt(origin, seed)
	fmt.Println(ciphertext)
	originStr, _ := aesTest.AESGCMDecrypt(ciphertext, seed)
	fmt.Println(originStr)

	// sign := fmt.Sprintf("%x", smTest.SM3("123456"))
	// fmt.Println(sign)
	// fmt.Println(len(sign))

	// salt, _ := hex.DecodeString("841daf3185844b9b98c3e4b5daa783be")
	// pbkdf2 := fmt.Sprintf("%x", pbkdf2Test.Test("123456", salt))
	// fmt.Println(pbkdf2)
	// fmt.Println(len(pbkdf2))
	// fsnotifyTest.Test()
}
