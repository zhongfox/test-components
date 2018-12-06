package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := strings.TrimSpace(os.Getenv("SECOND"))
	if s == "" {
		s = "60"
	}

	t, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("start sleep %d second\n", t)
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Printf("finish sleep %s second\n", s)

	os.Exit(0)
}
