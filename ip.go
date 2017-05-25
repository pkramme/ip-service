package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	listenaddr := flag.String("listen", ":80", "Address on which IP Service listens")
	loglocation := flag.String("log", "./ipservice.log", "Location of the logfile")
	flag.Parse()

	var f *os.File
	var err error
	if *loglocation != "disabled" {
		f, err = os.OpenFile(*loglocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.SetOutput(f)
		log.Println("IP Service Started")
	}
	fmt.Println("IP Service Started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.RemoteAddr)
		log.Printf("Request from %s", r.RemoteAddr)
	})
	log.Fatal(http.ListenAndServe(*listenaddr, nil))
}
