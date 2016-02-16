package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	var port string
	var path string
	if len(args) < 2 {
		port = "8080"
		path = "/var/www/html"
	} else {
		port = args[0]
		path = args[1]
	}
	log.Fatal(http.ListenAndServe(strings.Join([]string{":", port}, ""), http.FileServer(http.Dir(path))))
}
