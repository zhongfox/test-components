package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<body style='background-color:yellow;' ><h1>Hi, you've requested: %s</h1></body>", r.URL)
	})

	fmt.Println("web service try listen port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
