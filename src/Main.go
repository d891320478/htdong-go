package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"runtime/debug"
)

func Throwable() {
	err := recover()
	if err == nil {
		return
	}
	fmt.Println(err)
	fmt.Println(string(debug.Stack()))
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	defer Throwable()
	// smTest.Sm2WriteKeyFile()
	// smTest.Sm2Encrypt()
	// var val []string = nil
	// fmt.Println(len(val))
	// cronTest.Test1()
}
