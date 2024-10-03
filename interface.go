package main


type ILinedList[T any] interface{
     InsertAtBeginning(value T)
      InsertAtEnd(value T)
      PrintListForward()
}
