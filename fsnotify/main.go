package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func main() {
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
				fmt.Println("event: ", event)
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
	err = watcher.Add("/home/admin/conf")
	if err != nil {
		fmt.Println(err)
	}
	for true {
	}
}
