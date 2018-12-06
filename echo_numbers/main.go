package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	from, err := strconv.Atoi(strings.TrimSpace(os.Getenv("FROM")))
	if err != nil {
		fmt.Println(err)
		from = 1
	}

	to, err := strconv.Atoi(strings.TrimSpace(os.Getenv("TO")))
	if err != nil {
		fmt.Println(err)
		to = 100
	}

	for i := from; i < to; i++ {
		time.Sleep(time.Duration(100) * time.Millisecond)
		fmt.Println(i)
	}
}

type CMD struct {
	Command []string // cmd with args
	WorkDir string
}

func (c CMD) Run() (string, error) {
	cmdStr := strings.Join(c.Command, " ")
	fmt.Printf("[%s] Run CMD: %s\n", time.Now().Format("2006-01-02 15:04:05"), cmdStr)

	cmd := exec.Command(c.Command[0], c.Command[1:]...)
	if c.WorkDir != "" {
		cmd.Dir = c.WorkDir
	}

	data, err := cmd.CombinedOutput()
	result := string(data)
	if len(result) > 0 {
		fmt.Println(result)
	}

	return result, err
}
