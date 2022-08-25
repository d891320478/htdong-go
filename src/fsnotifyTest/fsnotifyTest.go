package fsnotifyTest

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func Test() {
	// 创建文件/目录监听器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
	}
	defer watcher.Close()
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					continue
				}
				// 打印监听事件
				fmt.Println("file: ", event.Name)
				fmt.Println("event: ", event.Op)
			case err, ok := <-watcher.Errors:
				fmt.Println("error: ", ok, err)
			}
		}
	}()
	// 监听系统的根目录目录
	err = watcher.Add("./")
	if err != nil {
		fmt.Println(err)
	}
	err = watcher.Add("/data/configproxy")
	if err != nil {
		fmt.Println(err)
	}
	for true {
	}
}
