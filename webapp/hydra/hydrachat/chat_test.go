package hydrachat

import (
  "testing"
  "time"
  "math/rand"
  "fmt"
  "net"
  "bufio"
  "strings"
)

func TestRun(t *testing.T) {
  if testing.Short() {
    t.Skip("Skipping test in short mode...")
  }
  go func() {
    t.Log("Starting Hydra chat server...")
    if err := Run(":2300"); err != nil {
      t.Error("Could not start the chat server ", err)
      return
    } else {
      t.Log("Started Hydra chat server... ")
    }
  }()

  time.Sleep(5*time.Second)

  rand.Seed(time.Now().UnixNano())
  name := fmt.Sprintf("Annonymous%d", rand.Intn(400))

  t.Logf("Hello %s, connecting to the hydra chat system... \n", name)
  conn, err := net.Dial("tcp", "127.0.0.1:2300")
  if err != nil {
    t.Fatal("Could not connect to hydra chat system", err)
  }
  t.Log("Connected to hydra chat system.")
  name += ": "
  defer conn.Close()
  msgCh := make(chan string)

  go func() {
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
      recvmsg := scanner.Text()
      sentmsg := <-msgCh
      if strings.Compare(recvmsg, sentmsg) != 0 {
        t.Errorf("Chat message %s does not match %s", recvmsg, sentmsg)
      }
    }
  }()

  for i:=0; i<= 10; i++ {
    msgbody := fmt.Sprintf("RandomMessage %d", rand.Intn(400))
    msg := name + msgbody
    _, err = fmt.Fprintf(conn, msg+"\n")
    if err != nil {
      t.Error(err)
      return
    }
    msgCh <- msg
  }
}