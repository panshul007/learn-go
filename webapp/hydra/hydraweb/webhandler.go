package hydraweb

import (
  "math/rand"
  "net/http"
  "fmt"
)

type testHandler struct {
  r int
}

func newHandler() testHandler {
  return testHandler{
    r: rand.Int(),
  }
}

func (h testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  switch r.URL.Path {
  case "/":
    fmt.Fprint(w, "Welcome to the Hydra  software system.")
  case "/testHandle":
    fmt.Fprint(w , "test handle object with random number: ", h.r)
  }
  fmt.Println(r.URL.Query())
}
