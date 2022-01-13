package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "tail -n 1000 /data/logs/knowledgecenterdata.log_20220113_1 | grep 21009")
	output, _ := cmd.Output()
	fmt.Println(string(output))
}
