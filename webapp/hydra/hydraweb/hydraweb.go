package hydraweb

import (
  "net/http"
  "github.com/panshul007/learn-go/webapp/hydra/hlogger"
  "fmt"
  //"time"
)

func Run() {

    http.HandleFunc("/", sroot)
    http.Handle("/testhandle", newHandler()) // this way we have limited the handler to the pattern
    http.HandleFunc("/testquery", queryTestHandler)
    http.ListenAndServe(":8080", nil)
   // to make it more flexible
   // http.ListenAndServe(":8080", newHandler())


  // Configuring the server
  /*
  server := &http.Server{
    Addr: ":8080",
    Handler: newHandler(),
    ReadTimeout: 5*time.Second,
    WriteTimeout: 5*time.Second,
  }
  server.ListenAndServe()
  */

}

func queryTestHandler(w http.ResponseWriter, r *http.Request) {
  q := r.URL.Query()
  message := fmt.Sprintf("Query map: %v \n", q)

  v1, v2 := q.Get("key1"), q.Get("key2")
  if v1 == v2 {
    message = message + fmt.Sprintf("V1 and V2 are equal %s \n", v1)
  } else {
    message = message + fmt.Sprintf("V1 is equal to %s, V2 is equal to %s \n", v1, v2)
  }
  fmt.Fprint(w, message)
}

func sroot(w http.ResponseWriter, r *http.Request) {
  logger := hlogger.GetInstance()
  fmt.Fprint(w, "Welcome to the Hydra software system.")
  logger.Println("Received an http Get request on root url.")
}