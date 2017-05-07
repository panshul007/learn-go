package main

import (
  "net/http"
  "fmt"
  "github.com/panshul007/learn-go/webapp/hydra/hlogger"
)

func main() {
  logger := hlogger.GetInstance()
  logger.Println("Starting Hydra web service")

  http.HandleFunc("/", sroot)
  http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
  logger := hlogger.GetInstance()
  fmt.Fprintf(w, "Welcome to the Hydra software system")

  logger.Println("Received an http Get request on root url")
}
