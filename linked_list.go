package main

import "fmt"

// Node represents a node in the linked list
type Node[T any] struct {
    value T
    next  *Node[T]
}

// LinkedList represents the linked list
type LinkedList[T any] struct {
    head *Node[T]
}

// InsertAtBeginning inserts a new node at the beginning of the list
func (ll *LinkedList[T]) InsertAtBeginning(value T) {
    newNode := &Node[T]{
        value: value,
        next:  ll.head,
    }
    ll.head = newNode
}

// InsertAtEnd inserts a new node at the end of the list
func (ll *LinkedList[T]) InsertAtEnd(value T) {
    newNode := &Node[T]{
        value: value,
    }

    if ll.head == nil {
        ll.head = newNode
        return
    }

    last := ll.head
    for last.next != nil {
        last = last.next
    }
    last.next = newNode
}

// PrintList prints all the values in the linked list
func (ll *LinkedList[T]) PrintList() {
    current := ll.head
    for current != nil {
        fmt.Printf("%v -> ", current.value)
        current = current.next
    }
    fmt.Println("nil")
}

// CreateIntList initializes and prints an integer linked list
func CreateIntList(values []int) *LinkedList[int] {
    ll := &LinkedList[int]{}
    for _, value := range values {
        ll.InsertAtEnd(value)
    }
    return ll
}

// CreateStringList initializes and prints a string linked list
func CreateStringList(values []string) *LinkedList[string] {
    ll := &LinkedList[string]{}
    for _, value := range values {
        ll.InsertAtEnd(value)
    }
    return ll
}

// PrintLinkedList prints the given linked list
func PrintLinkedList[T any](ll *LinkedList[T]) {
    ll.PrintList()
}
