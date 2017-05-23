package main

import (
  "net/http"
  "fmt"
  "flag"
)

func main() {
  listenaddr := flag.String("listen", ":80", "Address on which IP Service listens")
  flag.Parse()

  fmt.Println("IP Service Started")
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, r.RemoteAddr)
  })
  http.ListenAndServe(*listenaddr, nil)
}
