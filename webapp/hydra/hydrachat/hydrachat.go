package hydrachat

import (
  "github.com/panshul007/learn-go/webapp/hydra/hlogger"
  "net"
  "os"
  "os/signal"
  "syscall"
  "fmt"
)

var logger = hlogger.GetInstance()

// Start hydra chat
func Run(connection string) error {
  l, err := net.Listen("tcp", connection)
  if err != nil {
    logger.Println("Error connecting to chat client", err)
    return err
  }
  r := CreateRoom("HydraChat")

  go func() {
    // Handle SIGINT and SIGTERM
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    <-ch // will block the goroutine till something is received in this channel

    l.Close()
    fmt.Println("Closing tcp connection")
    close(r.Quit)
    if r.ClCount() > 0 {
      <-r.Msgch
    }
    os.Exit(0)
  }()

  for {
    conn, err := l.Accept()
    if err != nil {
      logger.Println("Error accepting connection from chat client", err)
      break
    }
    go handleConnection(r, conn)
  }

  return err
}

func handleConnection(r *room, c net.Conn) {
  logger.Println("Received request from client", c.RemoteAddr())
  r.AddClient(c)
}