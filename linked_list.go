package main

import "fmt"

// Node represents a node in the linked list

// Node represents a node in the singly linked list
type Node[T any] struct {
	value T
	next  *Node[T]
}

// LinkedList represents a singly linked list
type LinkedList[T any] struct {
	head *Node[T]
}

// InsertAtBeginning inserts a new node at the beginning of the singly linked list
func (ll *LinkedList[T]) InsertAtBeginning(value T) {
	newNode := &Node[T]{
		value: value,
		next:  ll.head,
	}
	ll.head = newNode
}

// InsertAtEnd inserts a new node at the end of the singly linked list
func (ll *LinkedList[T]) InsertAtEnd(value T) {
	newNode := &Node[T]{value: value}

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
func (ll *LinkedList[T]) Delete(value T) bool {
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
func (ll *LinkedList[T]) PrintListForward() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

// PrintListBackward is not applicable to singly linked lists
func (ll *LinkedList[T]) PrintListBackward() {
	fmt.Println("PrintListBackward not supported in singly linked list.")
}
