package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"os/exec"
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
	f, _ := os.OpenFile("/Users/dht31261/Desktop/a.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer f.Close()
	out := exec.Command("/bin/bash", "-c", "cp /Users/dht31261/Desktop/a.txt /Users/dht31261/Desktop/com")
	a, err := out.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(a))
}
