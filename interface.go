package main

type LinkedListOps interface {
	InsertAtBeginning(value int)
	InsertAtEnd(value int)
	Delete(value int) bool
	PrintListForward()
	PrintListBackward()
}
