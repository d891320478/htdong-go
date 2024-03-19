package cronTest

import (
	"fmt"
	"math/rand"

	"github.com/robfig/cron"
)

func Test1() {

	c := cron.New()

	c.AddFunc("*/30 * * * * *", func() {
		fmt.Println(rand.Int())
	})

	c.Start()

	select {}
}
