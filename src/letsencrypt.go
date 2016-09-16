package main

import (
  "fmt"
  "net/http"
  "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
  proc := sh("/usr/local/letsencrypt/letsencrypt.sh %s %s 2>&1", r.URL.Path, strings.Split(r.Header["Authorization"][0], " ")[1])
  _ = proc.Err()
  fmt.Fprintln(w,proc.stderr + proc.stdout)
}

func main() {
  http.HandleFunc("/", handler)
  http.Handle("/.well-known/acme-challenge/", http.FileServer(http.Dir("/srv")))
  http.ListenAndServe(":8080", nil)
}
