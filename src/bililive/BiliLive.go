package bililive

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
)

const roomId = "222272"
const danmuFilePath = "/data/biliDanMu222272/%d-%s-%s.log"

func Register(channel chan int, total int) {
	limit := make(map[int]int)
	lock := new(sync.RWMutex)
	c := client.NewClient(roomId)
	//弹幕事件
	c.OnDanmaku(func(danmaku *message.Danmaku) {
		if danmaku.Type != message.EmoticonDanmaku {
			val, err := strconv.Atoi(danmaku.Content)
			if err == nil {
				fmt.Printf("[弹幕] %s：%s\n", danmaku.Sender.Uname, danmaku.Content)
				if val > 0 && val <= total {
					lock.Lock()
					limit[danmaku.Sender.Uid]++
					if limit[danmaku.Sender.Uid] <= 3 {
						channel <- val
					}
					lock.Unlock()
				}
			}
		}
	})
	err := c.Start()
	if err != nil {
		panic(err)
	}
}

func AllDanMu() {
	c := client.NewClient(roomId)
	//弹幕事件
	c.OnDanmaku(func(danmaku *message.Danmaku) {
		if danmaku.Type != message.EmoticonDanmaku {
			writeToFile(time.Unix(danmaku.Timestamp/1000, 0).Format("2006-01-02 15:04:05"), danmaku.Sender.Uname, danmaku.Content, danmaku.Sender.Uid)
		}
	})
	err := c.Start()
	if err != nil {
		panic(err)
	}
}

func writeToFile(tm, uname, content string, uid int) {
	now := time.Now()
	filePath := fmt.Sprintf(danmuFilePath, now.Year(), now.Format("01"), now.Format("02"))
	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(fmt.Sprintf("[%s] %s[%d]: %s\n", tm, uname, uid, content))
	write.Flush()
	file.Close()
}
