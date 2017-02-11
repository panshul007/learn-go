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
  buckets := make([]int, 12)

  //Loop over the words
  for scanner.Scan() {
    n := HashBucket(scanner.Text(), 12)
    buckets[n]++ // counts how many words will be in each bucket
  }

  fmt.Println(buckets)

  // now the words are divided among 12 buckets. the distribution is efficient but still not even. lowest 4k and highest 47k.
  // more efficient than the last process, but still memory is not evenly distributed.
}

// receives a total number of buckets to be used as 12
// converts the first letter of word into integer.
// bucket number = int of letter mod number of buckets.
func HashBucket(word string, buckets int) int {
  letter := int(word[0])
  bucket := letter % buckets
  return bucket
}
