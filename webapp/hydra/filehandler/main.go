package main

import (
  "os"
  "log"
  "io"
  "io/ioutil"
  "fmt"
  "bufio"
  "time"
)

func main() {

  // Open file for read only.
  f1, err := os.Open("test1.txt")
  PrintFatalError(err)
  defer f1.Close()

  // Creating a new file
  f2, err := os.Create("test2.txt")
  PrintFatalError(err)
  defer f2.Close()

  // Open file for read write
  f3, err := os.OpenFile("test3.txt", os.O_APPEND|os.O_RDWR, 0666)
  // 0666 => Owner: (read & write), Group: (read & write), and other: (read & write)
  PrintFatalError(err)
  defer f3.Close()

  // rename a file
  err = os.Rename("./test1.txt", "./test1Nex.txt")
  PrintFatalError(err)

  // move a file
  err = os.Rename("./test1.txt", "./testfolder/test1.txt")
  PrintFatalError(err)

  // copy a file
  CopyFile("testfile2.txt", "./testfolder/testfile3.txt")

  // delete a file
  err = os.Remove("test2.txt")
  PrintFatalError(err)

  // Read a file content
  bytes, err := ioutil.ReadFile("test3.txt")
  fmt.Println(string(bytes))

  // Read large files, line by line
  scanner := bufio.NewScanner(f3)
  count := 0
  for scanner.Scan() {
    count++
    fmt.Println("Found line: ", count, scanner.Text())
  }

  // buffered write, efficient store in memory, saves disk i/o
  writebuffer := bufio.NewWriter(f3)
  for i:=1; i <=5; i++ {
    writebuffer.WriteString(fmt.Sprintln("Added line: ", i))
  }
  writebuffer.Flush() // flush the contents of buffer to disk.

  GenerateFileStatusReport("testFile5.txt")
}

func PrintFatalError(err error) {
  if err != nil {
    log.Fatal("Error happened while processing file: ", err)
  }
}

func CopyFile(fname1, fname2 string) {
  fOld, err := os.Open(fname1)
  PrintFatalError(err)
  defer fOld.Close()

  fNew, err := os.Create(fname2)
  PrintFatalError(err)
  defer fNew.Close()

  // copy bytes from source to destination
  _, err = io.Copy(fNew, fOld)
  PrintFatalError(err)

  // flush file contents to disk
  err = fNew.Sync()
  PrintFatalError(err)
}

func GenerateFileStatusReport(filename string) {
  filestats, err := os.Stat(filename)
  PrintFatalError(err)

  filestats.Name()
  filestats.IsDir()
  filestats.Mode()
  filestats.ModTime()
  filestats.Size()
}

// should be run as a go routine.. to not block the main thread into an infinite loop.
func WatchFile(fname string) {
  filestat1, err := os.Stat(fname)
  PrintFatalError(err)
  for {
    time.Sleep(1 * time.Second)
    filestat2, err := os.Stat(fname)
    PrintFatalError(err)
    if filestat1.ModTime() != filestat2.ModTime() {
      fmt.Println("File was modified at: ", filestat2.ModTime())
      filestat1, err = os.Stat(fname)
      PrintFatalError(err)
    }
  }
}