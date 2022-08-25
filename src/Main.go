package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
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
	//fsnotifyTest.Test()
	defer fmt.Println("exception")
	var b int32
	fmt.Scanf("%d", &b)
	a := 1 / b
	fmt.Println(a)
}
