package main

import (
  "net/http"
  "log"
  "bufio"
  "fmt"
  "os"
)

func main() {
  res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
  if err != nil {
    log.Fatalln(err)
  }

  words := make(map[string]string)

  sc := bufio.NewScanner(res.Body) // returns a pointer to the scanner with the body.
  sc.Split(bufio.ScanWords)        // splits the text in buffer at every word.

  for sc.Scan() { // starts scanning or reading the text in the scanner buffer.
    words[sc.Text()] = "" // puts every scanned word as text as a key in the map word, with a blank value.
  }
  if err := sc.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading input: ", err)
  }

  i := 0
  for k, _ := range words {
    fmt.Println(k)
    if i == 200 {
      break
    }
    i++
  }
}
