package main

import "fmt"
import "net/http"
import "runtime"
import "flag"

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}

func main() {
  worker := flag.Int("worker", 1, "worker cound")
  flag.Parse()
  runtime.GOMAXPROCS(*worker)

  http.HandleFunc("/", hello)
  http.ListenAndServe(":5000", nil)
}
