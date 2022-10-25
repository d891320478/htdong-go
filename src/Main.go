package main

import (
	"crypto/rand"
	"encoding/base64"
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
	f, _ := base64.StdEncoding.DecodeString("Myt+cnghYVQweT5VOzJNN1tOO2d9SEFYZFd2RUg2JExwQCB9flxkdD0wQllwW0JPZ1VCcCRlJ3NDaVl6LDcjOkMjcTBYXTNUI1pAcVEobzc+I1MkTT4lUE9oRCd6fE0mLSpJX3x6RDhPVWBNOXBwVihCdXB6NTIjXkxpO1whREA=")
	g := "3+~rx!aT0y>U;2M7[N;g}HAXdWvEH6$Lp@ }~\\dt=0BYp[BOgUBp$e'sCiYz,7#:C#q0X]3T#Z@qQ(o7>#S$M>%POhD'z|M&-*I_|zD8OU`M9ppV(Bupz52#^Li;\\!D@"
	fmt.Println(string(f) == g)
}
