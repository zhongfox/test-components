package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strings.TrimSpace(os.Getenv("OUTPUTS"))
	if s == "" {
		return
	}

	outs := strings.Split(s, "\n")

	for _, out := range outs {
		if !strings.Contains(out, "=") {
			break
		}

		kv := strings.SplitN(out, "=", 2)
		fmt.Printf("[JOB_OUT] %s = %s\n", kv[0], kv[1])
	}
}
