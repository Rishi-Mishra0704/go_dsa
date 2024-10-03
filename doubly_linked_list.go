package main

import "fmt"

// DNode represents a node in a doubly linked list
type DNode[T any] struct {
  value T
  next  *DNode[T]
  prev  *DNode[T]
}

// DoublyLinkedList implements LinkedListOps for a doubly linked list
type DoublyLinkedList[T any] struct {
  head *DNode[T]
  tail *DNode[T]
}

// InsertAtBeginning adds a node to the front of the list
func (ll *DoublyLinkedList[T]) InsertAtBeginning(value T) {
  newNode := &DNode[T]{value: value}

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
func (ll *DoublyLinkedList[T]) InsertAtEnd(value T) {
  newNode := &DNode[T]{value: value}

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
func (ll *DoublyLinkedList[T]) Delete(value T) bool {
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
func (ll *DoublyLinkedList[T]) PrintListForward() {
  current := ll.head
    for current != nil {
    fmt.Printf("%v -> ", current.value)
    current = current.next
  }
  fmt.Println("nil")
}

// PrintListBackward prints the list from tail to head
func (ll *DoublyLinkedList[T]) PrintListBackward() {
  current := ll.tail
    for current != nil {
    fmt.Printf("%v -> ", current.value)
    current = current.prev
  }
  fmt.Println("nil")
}

// NewDoublyLinkedList initializes a new doubly linked list
func NewDoublyLinkedList[T any]() LinkedListOps[T] {
  return &DoublyLinkedList[T]{}
}


