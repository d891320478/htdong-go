package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"strconv"
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

func xor(k1 []byte, k2 []byte) []byte {
	l1 := len(k1)
	l2 := len(k2)
	l := l1
	if l < l2 {
		l = l2
	}
	rlt := make([]byte, l)
	for i := 0; i < l; i++ {
		var a byte = 0
		if i < l1 {
			a = k1[i]
		}
		var b byte = 0
		if i < l2 {
			b = k2[i]
		}
		rlt[i] = a ^ b
	}
	return rlt
}

func checkBase64(ori string) []byte {
	bt, err := base64.StdEncoding.DecodeString(ori)
	if err != nil {
		return []byte(ori)
	}
	return bt
}

const root_key_factor = "!,V>G)_K]/`Q#/\\wn/]>.if.H}\\|gw^*;BHxHR;o>*C0&XW{/zW5\"5(I0'>:(9XpWde^t[N3R7Fq'F;WM}*|k8o5kY2a9%l'#Y0zZJP6x{cf%5t^LP\\J4vy@&j<)a:%2"

func GetAllRootKeyFile() (val [][]byte) {

	keyFile, _ := os.Open("/data/configproxy/rtk")
	defer keyFile.Close()
	readKeyFile := bufio.NewReader(keyFile)
	readKey, _, err := readKeyFile.ReadLine()
	if err == io.EOF || len(readKey) == 0 {
		panic("root key file is empty")
	}

	val = make([][]byte, 1)
	val[0] = checkBase64(string(readKey))
	return
}

func GetAllRootKeyStr() (val [][]byte) {
	rtks := GetAllRootKeyFile()

	val = make([][]byte, len(rtks)+1)
	val[0] = []byte(root_key_factor)
	for i := 0; i < len(rtks); i++ {
		val[i+1] = rtks[i]
	}
	return
}

func GetSaltStr() (val []byte) {
	saltFile, _ := os.Open("/data/configproxy/rts")
	defer saltFile.Close()
	readSaltFile := bufio.NewReader(saltFile)
	readSalt, _, err := readSaltFile.ReadLine()
	if err == io.EOF || len(readSalt) == 0 {
		panic("root key salt is empty")
	}
	val = checkBase64(string(readSalt))
	return
}

func main() {
	rtkStr, _ := goutils.CryptoRandomAscii(128)
	rtkFile, err := os.OpenFile("rtk", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer rtkFile.Close()
	if err != nil {
		panic("write rtk file error." + err.Error())
	}
	write := bufio.NewWriter(rtkFile)
	write.WriteString(rtkStr)
	write.Flush()

	rtsStr, _ := goutils.CryptoRandomAscii(128)
	rtsFile, err := os.OpenFile("rts", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer rtsFile.Close()
	if err != nil {
		panic("write rts file error." + err.Error())
	}
	write = bufio.NewWriter(rtsFile)
	write.WriteString(rtsStr)
	write.Flush()
}
