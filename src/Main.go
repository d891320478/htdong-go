package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
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

func checkStatus(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("success"))
}

func wait(w http.ResponseWriter, req *http.Request) {
	sleep, _ := strconv.ParseInt(req.URL.Query().Get("sleep"), 10, 64)
	time.Sleep(time.Duration(sleep) * time.Second)
	w.Write([]byte("wait"))
}

func main() {
	mp := make(map[string]string)
	mp["a"] = "b"
	_, ok := mp["x"]
	fmt.Println(ok)
}
