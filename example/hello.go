package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/weekface/pat"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
}

func myMiddleware(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello world")
}

func main() {
	m := pat.New()
	m.Use(http.HandlerFunc(myMiddleware))
	m.Get("/hello/:name", http.HandlerFunc(HelloServer))

	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
