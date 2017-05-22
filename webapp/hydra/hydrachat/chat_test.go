package hydrachat

import (
  "testing"
  "time"
)

func TestRun(t *testing.T) {
  go func() {
    t.Log("Starting Hydra chat server...")
    if err := Run(":2300"); err != nil {
      t.Error("Could not start the chat server ", err)
      return
    } else {
      t.Log("Started Hydra chat server... ")
    }
  }()

  time.Sleep(10*time.Second)
}