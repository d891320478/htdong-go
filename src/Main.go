package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/aokoli/goutils"
	"github.com/htdong/gotest/src/bililive"
	"github.com/redis/go-redis/v9"
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

func createRt() {
	rtkStr, _ := goutils.CryptoRandomAscii(128)
	rtkFile, err := os.OpenFile("rtk", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer rtkFile.Close()
	if err != nil {
		panic("write rtk file error." + err.Error())
	}
	write := bufio.NewWriter(rtkFile)
	write.WriteString(base64.StdEncoding.EncodeToString([]byte(rtkStr)))
	write.Flush()

	rtsStr, _ := goutils.CryptoRandomAscii(128)
	rtsFile, err := os.OpenFile("rts", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer rtsFile.Close()
	if err != nil {
		panic("write rts file error." + err.Error())
	}
	write = bufio.NewWriter(rtsFile)
	write.WriteString(base64.StdEncoding.EncodeToString([]byte(rtsStr)))
	write.Flush()
}

func Throwable() {
	err := recover()
	if err == nil {
		return
	}
	fmt.Println(err)
	fmt.Println(string(debug.Stack()))
}

const song_list_file = "list.txt"

func biliToupiao() {
	if !PathExists(song_list_file) {
		file, _ := os.Create(song_list_file)
		defer file.Close()
		w := bufio.NewWriter(file)
		w.WriteString("")
		w.Flush()
		file.Close()
	}
	stdinReader := bufio.NewReader(os.Stdin)
	fmt.Println("歌单放到list.txt，保存，然后按回车。。。。。。")
	stdinReader.ReadString('\n')
	// 读取歌单内容，生成编号，初始化票数
	mp := make(map[int]int)

	f, _ := os.Open("list.txt")
	defer f.Close()
	a, _ := io.ReadAll(f)
	a = append(a, 13)
	var list []string
	// list = append(list, "规则：征集时请把想听的歌打在弹幕 选曲没任何限制0w0")
	// list = append(list, "投票时每个人最多三票")
	// list = append(list, "会唱的里面得票最高的今天唱 不会唱的里面得票最高的下周唱")

	tmp := make([]byte, 0)
	for _, v := range a {
		if v == byte(10) || v == byte(13) {
			ss := string(tmp)
			fmt.Println(tmp)
			fmt.Printf("ss = %s\n", ss)
			if len(strings.TrimSpace(ss)) > 0 {
				list = append(list, strings.TrimSpace(ss))
			}
			tmp = make([]byte, 0)
		} else {
			tmp = append(tmp, v)
		}
	}
	total := len(list)

	for i := 0; i < total; i++ {
		mp[i] = 0
	}
	fmt.Println(total)
	fmt.Println(list)
	fmt.Println(mp)
	// 回写文件
	writeToListFile(mp, list, total)
	// 发请求start
	http.Get("http://47.97.10.207:9961/htdong/liveVote/startVote")
	time.Sleep(3 * time.Second)
	resp, _ := http.Get(fmt.Sprintf("http://47.97.10.207:9961/startLive/startGetDanMu?total=%d", total))
	for {
		fmt.Print("startGetDanMu ")
		fmt.Println(resp.StatusCode)
		if resp.StatusCode == 200 {
			break
		}
		time.Sleep(1 * time.Second)
		resp, _ = http.Get(fmt.Sprintf("http://47.97.10.207:9961/startLive/startGetDanMu?total=%d", total))
	}
	mp = make(map[int]int)
	for {
		resp, _ := http.Get("http://47.97.10.207:9961/startLive/getCountRlt")
		if resp.StatusCode == 200 {
			defer resp.Body.Close()
			jsonStr, _ := io.ReadAll(resp.Body)
			json.Unmarshal(jsonStr, &mp)
			fmt.Println(mp)
			writeToListFile(mp, list, total)
		} else {
			fmt.Println(resp.StatusCode)
		}
		time.Sleep(3 * time.Second)
	}
	// var channel chan int = make(chan int)

	// bililive.Register(channel, total)
	// for {
	// 	val := <-channel
	// 	mp[val-1]++
	// 	writeToListFile(mp, list, total)
	// }
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
	// biliToupiao()
	// bililive.StartBiliHttp()
	bililive.AllDanMu()
}

func writeToListFile(mp map[int]int, list []string, total int) {
	file, _ := os.OpenFile(song_list_file, os.O_WRONLY, os.ModeAppend)
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("规则：征集时请把想听的歌打在弹幕 选曲没任何限制0w0\r\n")
	write.WriteString("投票时每个人最多三票\r\n")
	write.WriteString("从会唱的里面选两首的票最高的今天唱/吹\r\n")
	write.WriteString("不会的里最高的下周唱（不含卡祖笛曲）\r\n")
	write.WriteString("想听卡祖笛版也可以，在歌名前面加上卡祖笛三字。下周翻唱\r\n")
	for i := 0; i < total; i++ {
		val := strconv.Itoa(i+1) + ". " + strings.TrimSpace(list[i]) + "   " + strconv.Itoa(mp[i]) + " 票\r\n"
		write.WriteString(val)
	}
	write.Flush()
}
