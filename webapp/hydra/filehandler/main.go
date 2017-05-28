package main

import (
  "os"
)

func main() {

  // Open file for read only.
  f1, err := os.Open("test1.txt")
  println(err)
  defer f1.Close()

  // Creating a new file
  f2, err := os.Create("test2.txt")
  println(err)
  defer f2.Close()

  // Open file for read write
  f3, err := os.OpenFile("test3.txt", os.O_APPEND|os.O_RDWR, 0666)
  // 0666 => Owner: (read & write), Group: (read & write), and other: (read & write)
  println(err)
  defer f3.Close()

  // rename a file
  err = os.Rename("./test1.txt", "./test1Nex.txt")
  println(err)

  // move a file
  err = os.Rename("./test1.txt", "./testfolder/test1.txt")
  println(err)

  // copy a file

}