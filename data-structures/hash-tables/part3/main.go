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

  // now the words are divided among 12 buckets. the distribution is efficient, lowest 11k and highest 32k
  // much more efficient than the last process.
}

// receives a total number of buckets to be used as 12
// calculates a sum integer values of each letter in the word
// bucket number = above sum mod number of buckets.
func HashBucket(word string, buckets int) int {
  var sum int
  for _, v := range word {
    sum += int(v)
  }
  return sum % buckets
}
