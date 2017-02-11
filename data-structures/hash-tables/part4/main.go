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
  res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
  if err != nil {
    log.Fatalln(err)
  }

  // scan the page
  scanner := bufio.NewScanner(res.Body)
  defer res.Body.Close() // closes the response resource just before the method exists.

  // Set the split function for the scanning operation.
  scanner.Split(bufio.ScanWords)

  // Create slice of slice of string to hold slices of words
  buckets := make([][]string, 12)

  // initializing bucket slices with empty slices of string
  for i := 0; i < 12; i++ {
    buckets = append(buckets, []string{})
  }

  //Loop over the words
  for scanner.Scan() { // returns a bool true as long as there are more words to scan
    word := scanner.Text()
    n := HashBucket(word, 12)
    buckets[n] = append(buckets[n], word)
  }

  // Print length of each bucket
  for i := 0; i < 12; i++ {
    fmt.Println(i, " - ", len(buckets[i]))
  }

  // Print words in one of the buckets.
  // fmt.Println(buckets[6])

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
