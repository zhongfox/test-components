package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	sec := strings.TrimSpace(os.Getenv("SECOND"))
	if sec == "" {
		sec = "10"
	}

	s, err := strconv.ParseInt(sec, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("start sleep %d second\n", s)
	time.Sleep(time.Duration(s) * time.Second)
	fmt.Printf("finish sleep %d second\n", s)

	code := strings.TrimSpace(os.Getenv("EXIT_CODE"))
	if code == "" {
		code = "0"
	}
	c, err := strconv.Atoi(code)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("exit code is  %d\n", c)
	os.Exit(c)
}
