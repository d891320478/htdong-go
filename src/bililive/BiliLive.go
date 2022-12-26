package bililive

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
	_ "github.com/Akegarasu/blivedm-go/utils"
)

const roomId = "222272"

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
