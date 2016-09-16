package main

import (
  "fmt"
  "net/http"
  "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
  proc := sh("/usr/local/letsencrypt/letsencrypt.sh '%s' '%s' 2>&1", strings.Split(r.URL.Path, "/")[1], strings.Split(r.Header["Authorization"][0], " ")[1])
  fmt.Fprintln(w,proc.stderr + proc.stdout)
}

func main() {
  http.HandleFunc("/", handler)
  http.Handle("/.well-known/acme-challenge/", http.FileServer(http.Dir("/srv")))
  http.ListenAndServe(":8080", nil)
}
