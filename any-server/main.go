package main

import (
	"any-server/protocols"
	"os"
	"strings"

)

func main() {
	protocol := strings.TrimSpace(os.Getenv("PROTOCOL"))
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "4000"
	}

	switch protocol {
	case "GRPC":
		protocols.StartGRPCServer(port)
	case "HTTP":
		protocols.StartHTTPServer(port)
	default:
		protocols.StartTCPServer(port)
	}
}
