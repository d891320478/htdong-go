package main

import (
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

func main() {
	defer Throwable()
	// smTest.Sm2WriteKeyFile()
	// smTest.Sm2Encrypt()
	// var val []string = nil
	// fmt.Println(len(val))
	// cronTest.Test1()
}
