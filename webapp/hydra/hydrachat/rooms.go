package hydrachat

import (
  "sync"
  "io"
  "fmt"
)

type room struct {
  name string
  Msgch chan string // channel to receive messages
  clients map[chan<- string]struct{} // set of send only channels
  Quit chan struct{}
  *sync.RWMutex
}

func CreateRoom(name string) *room {
  r := &room{
    name: name,
    Msgch: make(chan string),
    RWMutex: new(sync.RWMutex),
    clients: make(map[chan<- string]struct{}),
    Quit: make(chan struct{}),
  }
  r.Run()
  return r
}

// Fanout pattern: Run(), broadcastMsg()
func (r *room) Run() {
  logger.Println("Starting chat room", r.name)
  go func() {  // this will exit when close(r.Msgch) is called in RemoveClient method.
    for msg := range r.Msgch {
      r.broadcastMsg(msg)
    }
  }()
}

func (r *room) broadcastMsg(msg string) {
  r.RLock()
  defer r.RUnlock()
  fmt.Println("Received message: ", msg)
  for wc, _ := range r.clients {
    go func(wc chan<- string) {
      wc <- msg
    }(wc)
  }
}

// Clients joining the chat room, need to be registered to the list of clients listening to the chat room.
func (r *room) AddClient(c io.ReadWriteCloser) {
  r.Lock()
  wc, done := StartClient(r.Msgch, c, r.Quit)
  r.clients[wc] = struct{}{}
  r.Unlock()

  // remove client when done is signalled
  go func() {
    <- done
    r.RemoveClient(wc)
  }()
}

func (r *room) ClCount() int {
  return len(r.clients)
}

func (r *room) RemoveClient(wc chan<- string) {
  logger.Println("Remove client ")
  r.Lock()
  close(wc)
  delete(r.clients, wc)
  r.Unlock()
  select {
  case <-r.Quit:
    if len(r.clients) == 0 {
      close(r.Msgch)
    }
  default: // empty default so that code does not block when there is nothing in r.Quit channel.
  }
}