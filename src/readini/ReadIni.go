package readini

import (
	"fmt"

	"github.com/larspensjo/config"
)

func ReadIni() {
	cof, _ := config.ReadDefault("/data/configproxy/1.ini")
	x, _ := cof.String("abc", "a")
	y, _ := cof.String("abc", "b")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(len(y))
}
