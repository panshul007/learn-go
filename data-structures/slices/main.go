package main

import "fmt"

func main() {
  mySlice := []int{1, 3, 5, 7, 9, 11}
  fmt.Printf("%T\n", mySlice)
  fmt.Println(mySlice)

  mySlice1 := []string{"a", "b", "c", "g", "m", "z"}
  fmt.Println(mySlice1)
  fmt.Println(mySlice1[2:4]) //slicing a slice -> [c g] -> right inclusive, left exclusive
  fmt.Println(mySlice1[2])   //index access -> c
  fmt.Println("myString"[2]) //index access -> 83 which is ascii for `S`
  example1()
  example2()
  sliceofslice()
  doubleint()
}

func example1() {
  mySlice := make([]int, 0, 5)

  fmt.Println("---------------------")
  fmt.Println(mySlice)
  fmt.Println(len(mySlice))
  fmt.Println(cap(mySlice))
  fmt.Println("---------------------")

  for i := 0; i < 80; i++ {
    mySlice = append(mySlice, i)
    fmt.Println("Len: ", len(mySlice), " Capacity: ", cap(mySlice), " Value: ", mySlice[i])
  }
}

func example2() {
  mySlice := []int{1, 2, 3, 4, 5, 6, } // this will create len = capacity, so if u append, it has to create a new under lying array.
  // hence this has a performance hit.
  fmt.Println(mySlice)

  customerNumber := make([]int, 3) // in this case 3 will be both len and cap.
  fmt.Println(customerNumber)
}

func sliceofslice() {
  records := make([][]string, 0)

  //student 1
  student1 := make([]string, 4)
  student1[0] = "Foster"
  student1[1] = "Nathan"
  student1[2] = "100.00"
  student1[3] = "74.00"

  //store the record
  records = append(records, student1)

  //student 2
  student2 := make([]string, 4)
  student2[0] = "Foster"
  student2[1] = "Nathan"
  student2[2] = "100.00"
  student2[3] = "74.00"

  records = append(records, student2)

  fmt.Println(records)
}

func doubleint() {
  transactions := make([][]int, 0, 3)

  for i := 0; i < 3; i++ {
    transaction := make([]int, 0)
    for j := 0; j < 4; j++ {
      transaction = append(transaction, j)
    }
    transactions = append(transactions, transaction)
  }
  fmt.Println(transactions)
}
