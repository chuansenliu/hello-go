package main

import (
	"fmt"
	"net/http"

	"github.com/eehut/hello-go/command"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world, %s!", request.URL.Path[1:])
}

// doHttp is command function
func doHttp(argv []string) int {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	return 0
}

func init() {
	command.AddCommand("http", "http test", "Usage:\n time ", doHttp)
}
