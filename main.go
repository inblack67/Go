package main

import (
	"flag"
	"log"
	"net/http"
)

func main (){
	port := flag.String("p", "5000", "port")
	dir:= flag.String("d", "static", "dir")
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("Serving %s on port %s \n", *dir, *port)
	log.Fatal(http.ListenAndServe(":"+ *port, nil))
}
