package main

import "fmt"

// PrintLinkedList prints the given linked list (either singly or doubly) using the interface
func PrintLinkedList(iLL LinkedListOps) {
	fmt.Println("Forward:")
	iLL.PrintListForward()
}

// Create and test the singly linked list
func CreateSinglyLinkedList() {
	list := &LinkedList{}
	list.InsertAtBeginning(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	fmt.Println("Singly Linked List")
	PrintLinkedList(list) // Calls the interface method
}

// Create and test the doubly linked list
func CreateDoublyLinkedList() {
	list := &DoublyLinkedList{}
	list.InsertAtBeginning(40)
	list.InsertAtEnd(50)
	list.InsertAtEnd(60)
	fmt.Println("Doubly Linked List")
	PrintLinkedList(list)    // Calls the interface method
	list.PrintListBackward() // Unique to doubly linked list
}

func main() {
	CreateSinglyLinkedList()
	CreateDoublyLinkedList()
}
