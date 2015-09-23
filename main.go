package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "9090", "Default serve port")
	wd, _ := os.Getwd()

	flag.Parse()

	fs := http.FileServer(http.Dir(wd))
	http.Handle("/", fs)

	log.Println("Serving from " + wd + " on " + *port)
	http.ListenAndServe(":"+*port, nil)
}
