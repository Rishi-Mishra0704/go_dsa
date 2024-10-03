package main

import "fmt"

// DNode represents a node in a doubly linked list
type DNode struct {
	value int
	next  *DNode
	prev  *DNode
}

// DoublyLinkedList implements LinkedListOps for a doubly linked list
type DoublyLinkedList struct {
	head *DNode
	tail *DNode
}

// InsertAtBeginning adds a node to the front of the list
func (ll *DoublyLinkedList) InsertAtBeginning(value int) {
	newNode := &DNode{value: value}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head.prev = newNode
		ll.head = newNode
	}
}

// InsertAtEnd adds a node to the end of the list
func (ll *DoublyLinkedList) InsertAtEnd(value int) {
	newNode := &DNode{value: value}

	if ll.tail == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.prev = ll.tail
		ll.tail.next = newNode
		ll.tail = newNode
	}
}

// Delete removes a node by value
func (ll *DoublyLinkedList) Delete(value int) bool {
	if ll.head == nil {
		return false // List is empty
	}

	current := ll.head
	for current != nil {
		if current.value == value {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				ll.head = current.next // Node is the head
			}

			if current.next != nil {
				current.next.prev = current.prev
			} else {
				ll.tail = current.prev // Node is the tail
			}
			return true // Node deleted
		}
		current = current.next
	}
	return false // Value not found
}

// PrintListForward prints the list from head to tail
func (ll *DoublyLinkedList) PrintListForward() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

// PrintListBackward prints the list from tail to head
func (ll *DoublyLinkedList) PrintListBackward() {
	current := ll.tail
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.prev
	}
	fmt.Println("nil")
}

// NewDoublyLinkedList initializes a new doubly linked list
func NewDoublyLinkedList() LinkedListOps {
	return &DoublyLinkedList{}
}
