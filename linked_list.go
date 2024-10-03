package main

import (
	"fmt"
)

// Node represents a node in the linked list

// Node represents a node in the singly linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	head *Node
}

// InsertAtBeginning inserts a new node at the beginning of the singly linked list
func (ll *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{
		value: value,
		next:  ll.head,
	}
	ll.head = newNode
}

// InsertAtEnd inserts a new node at the end of the singly linked list
func (ll *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{value: value}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Delete removes the node with the specified value
func (ll *LinkedList) Delete(value int) bool {
	if ll.head == nil {
		return false // List is empty
	}

	// If the node to be deleted is the head
	if ll.head.value == value {
		ll.head = ll.head.next
		return true
	}

	// Traverse to find the node to be deleted
	current := ll.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			return true
		}
		current = current.next
	}

	return false // Node not found
}

// PrintListForward prints all the values in the linked list (singly list, only forward)
func (ll *LinkedList) PrintListForward() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

// PrintListBackward is not applicable to singly linked lists
func (ll *LinkedList) PrintListBackward() {
	fmt.Println("PrintListBackward not supported in singly linked list.")
}
