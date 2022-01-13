import (
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "tail -n 1000 /data/logs/knowledgecenterdata.log_20220113_1 | grep 21009")
	output, err := cmd.Output()
	fmt.Println(string(output))
}