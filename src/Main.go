package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func wait(w http.ResponseWriter, req *http.Request) {
	sleep, _ := strconv.ParseInt(req.URL.Query().Get("sleep"), 10, 64)
	time.Sleep(time.Duration(sleep) * time.Second)
	w.Write([]byte("wait"))
}

func checkBase64(ori string) []byte {
	bt, err := base64.StdEncoding.DecodeString(ori)
	if err != nil {
		return []byte(ori)
	}
	return bt
}

func Throwable() {
	err := recover()
	if err == nil {
		return
	}
	fmt.Println(err)
	fmt.Println(string(debug.Stack()))
}

func getServerFromSentinel() {
	sentinel := redis.NewSentinelClient(&redis.Options{
		Addr: "10.0.20.67:26379",
		// Password: "donghaotian",
	})
	addr, err := sentinel.GetMasterAddrByName(context.Background(), "main").Result()
	fmt.Println(addr)
	fmt.Println(err)
}

func main() {
	defer Throwable()
	// redisService.Put("sso1", "sso1", 1, time.Minute)
	// fmt.Println(redisService.Get("sso1"))
	// getServerFromSentinel()
	// smTest.Sm2WriteKeyFile()
	// smTest.Sm2Encrypt()
	// bililive.StartBiliHttp()
	var val []string = nil
	fmt.Println(len(val))
}

func runCmd(cmd string) string {
	cmd1 := exec.Command("/bin/bash", "-c", cmd)
	out1, err := cmd1.Output()
	if err != nil {
		fmt.Println(cmd)
		fmt.Println(err)
		return "0 a"
	} else {
		return string(out1)
	}
}
