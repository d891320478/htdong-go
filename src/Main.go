package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aokoli/goutils"
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

const root_key_assembly_length = 128

func rt() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	rtkStr, _ := goutils.CryptoRandom(root_key_assembly_length, 0, 127, false, false)
	rtkFile, _ := os.OpenFile(dir+"/rtk", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer rtkFile.Close()
	write := bufio.NewWriter(rtkFile)
	write.WriteString(base64.StdEncoding.EncodeToString([]byte(rtkStr)))
	write.Flush()

	rtsStr, _ := goutils.CryptoRandom(root_key_assembly_length, 0, 127, false, false)
	rtsFile, _ := os.OpenFile(dir+"/rts", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer rtsFile.Close()
	write = bufio.NewWriter(rtsFile)
	write.WriteString(base64.StdEncoding.EncodeToString([]byte(rtsStr)))
	write.Flush()
}

func main() {
	// rt()
	fmt.Println(strings.Split("abc", "."))
}
