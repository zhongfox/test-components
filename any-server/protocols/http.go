package protocols

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func StartHTTPServer(port string) {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		name := v.Get("name")

		log.Printf("receiving call, %v", v)

		fmt.Fprintf(w, "hello %s", name)
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("http server listening at %v\n", lis.Addr())

	err = http.Serve(lis, nil)
	if err != nil {
		fmt.Println(err)
	}

}
