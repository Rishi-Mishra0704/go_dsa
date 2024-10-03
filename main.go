package main

import "fmt"


func main(){
  // integer list  
  fmt.Println("Integer List:")
  intList := CreateIntList([]int{1, 2, 3, 4, 5})
  PrintLinkedList(intList) // Output: 1 -> 2 -> 3 -> 4 -> 5 -> nil
  // string list
  fmt.Println("String List:")
  strList := CreateStringList([]string{"apple", "banana", "cherry"})
  PrintLinkedList(strList) // Output: apple -> banana -> cherry -> nil

}
