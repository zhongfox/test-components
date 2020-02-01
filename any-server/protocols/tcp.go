package protocols

import (
	"fmt"
	"io"
	"log"
	"net"
)

func StartTCPServer(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer lis.Close()
	log.Printf("tcp server listening at %v\n", lis.Addr())

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("error accepting: %v", err)
		}

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	defer conn.Close()
	client := conn.RemoteAddr().String()
	log.Printf("receiving call, %s", client)
	buf := make([]byte, 1024)

	for {
		reqLen, err := conn.Read(buf)
		if err == io.EOF {
			log.Printf("the client:%s is closed\n", client)
			return
		}
		if err != nil {
			log.Println("error to read message: ", err)
			return
		}

		if reqLen > 0 {
			in := string(buf[:reqLen])
			log.Println(reqLen)

			fmt.Printf("read:%s", in)
			_, err = conn.Write([]byte("hello: " + in))
		}
	}

}
