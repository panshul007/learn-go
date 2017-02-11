package main

import (
  "net/http"
  "log"
  "bufio"
  "fmt"
)

// map is built on hash tables which is built on arrays

// keys generate a hash code which decides which bucket to put the value into.

func main() {
  res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
  if err != nil {
    log.Fatalln(err)
  }

  // scan the page
  scanner := bufio.NewScanner(res.Body)
  defer res.Body.Close() // closes the response resource just before the method exists.

  // Set the split function for the scanning operation.
  scanner.Split(bufio.ScanWords)

  // Create slice to hold words
  buckets := make([]int, 200)

  //Loop over the words
  for scanner.Scan() {
    n := HashBucket(scanner.Text())
    buckets[n]++ // counts how many words will be in each bucket
  }

  fmt.Println(buckets[65:122])

  // gives a pretty bad distribution of words stored in a bucket. Some buckets have too many words and some have no words in them.
  // very bad utilization of memory, as empty buckets are also storing 200 units of memory.
}

// converts the first letter of word into integer.
func HashBucket(word string) int {
  return int(word[0])
}
