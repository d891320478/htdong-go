package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
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

// 生成随机内容的byte数组
func randomByteArray(length int) (rlt []byte, err error) {
	rlt = make([]byte, length)
	if _, err = io.ReadFull(rand.Reader, rlt); err != nil {
		return
	}
	return
}

func main() {
	mp := make(map[string]string)
	mp["a"] = "a"
	mp["b"] = "b"
	mp["c"] = "c"
	mp["f"] = "f"
	mp["d"] = "d"
	mp["e"] = "e"
	mp["bb"] = "bb"
	mp["aa"] = "aa"
	mp["Bcacefef"] = "adfasdf"
	fmt.Println(mp)
	for k, v := range mp {
		fmt.Printf("%s, %s  ", k, v)
	}
	fmt.Println()
	for k, v := range mp {
		fmt.Printf("%s, %s  ", k, v)
	}
	fmt.Println()

	workKeyByte, err := randomByteArray(32)
	if err != nil {
		return
	}
	fmt.Println(hex.EncodeToString(workKeyByte))

	origin := "utf-8"
	seed := "bc30af056a041c5608ebf2da933c8f4669b493bea264abd960b005f90c35d89b"
	ciphertext, _ := aesTest.AESGCMEncrypt(origin, seed)
	fmt.Println(ciphertext)
	originStr, _ := aesTest.AESGCMDecrypt("2beb3d440f9df09d9ec59ffc7dfc7086f661116c45bad575847e02afd17feae1", seed)
	fmt.Println(originStr)
}
