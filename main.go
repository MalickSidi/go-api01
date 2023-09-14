package main

import (
  "net/http"
  "fmt"
)

func main() {
  server := &http.Server {

    Addr: ":3000",
    Handler: http.HandlerFunc(basicHandler),
  }
  err := server.ListenAndServe()
  if err != nil {
    fmt.Println("faild to listen to server", err)
  }
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Testing http, everything is alright?"))
}
