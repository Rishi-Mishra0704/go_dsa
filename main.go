package main


func main(){
// integer list  
intList := CreateIntList([]int{1, 2, 3, 4, 5})
PrintLinkedList(intList) // Output: 1 -> 2 -> 3 -> 4 -> 5 -> nil
// string list
  strList := CreateStringList([]string{"apple", "banana", "cherry"})
  PrintLinkedList(strList) // Output: apple -> banana -> cherry -> nil

}
