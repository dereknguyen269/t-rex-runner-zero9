package main

import (
  "net/http"
  "fmt"
)

func main() {
  fmt.Println("T-rex is running on http://localhost:5000...")
  http.Handle("/", http.FileServer(http.Dir("./public")))
  http.ListenAndServe(":5000", nil)
}
