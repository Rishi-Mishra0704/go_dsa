package main


type LinkedListOps[T any] interface {
	InsertAtBeginning(value T)
	InsertAtEnd(value T)
	Delete(value T) bool
	PrintListForward()
	PrintListBackward()
}
