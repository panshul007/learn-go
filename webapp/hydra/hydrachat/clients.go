package hydrachat

import (
  "bufio"
  "io"
)

type client struct {
  *bufio.Reader
  *bufio.Writer
  wc chan string // client write channel
}

// cn: client connection -> assuming to be TCP/IP connection
func StartClient(msgCh chan<- string, cn io.ReadWriteCloser, quit chan struct{}) (chan<- string, chan struct{}) {
  c := new(client)
  c.Reader = bufio.NewReader(cn)
  c.Writer = bufio.NewWriter(cn)
  c.wc = make(chan string)
  done := make(chan struct{}) // sending signal when client stops

  //setup the reader
  go func() {
    scanner := bufio.NewScanner(c.Reader) // init with tcp channel
    for scanner.Scan() { // read data as it comes over the tcp channel
      logger.Println(scanner.Text())
      msgCh <- scanner.Text()
    }
    done <- struct{}{} // loop exits when tcp connection breaks... done signal sent to close the client
  }()

  // setup the writer
  c.writeMonitor()

  go func() {
    select {
    case <-quit:
      cn.Close()
    case <-done:
    }
  }()

  return c.wc, done
}

func (c *client) writeMonitor() {
  go func() {
    for s:= range c.wc {
      c.WriteString(s + "\n")
      c.Flush()
    }
  }()
}