package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "8080", "Default serve port")
	wd, _ := os.Getwd()

	flag.Parse()

	fs := http.FileServer(http.Dir(wd))
	http.Handle("/", fs)

	log.Println("Serving " + wd + " on " + "http://localhost:" + *port)
	http.ListenAndServe(":"+*port, nil)
}
