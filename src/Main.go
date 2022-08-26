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
	_, err := os.Create("/Users/dht31261/Desktop/a.txt")
	fmt.Println(err)
	fmt.Println("-----")
	_, err = os.Create("/Users/dht31261/Desktop/a.txt")
	fmt.Println(err)
	//fsnotifyTest.Test()
	// defer fmt.Println("exception")
	// defer func() {
	// 	err := recover()
	// 	fmt.Println(err)
	// }()
	// var b int32
	// fmt.Scanf("%d", &b)
	// a := 1 / b
	// fmt.Println(a)
}
