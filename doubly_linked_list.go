package main


type DNode[T any] struct {
  value T
  next *Node[T]
  prev *Node[T]

}

type DoublyLinkedList[T any] struct {
  head *DNode[T]
  tail *Node[T]

}

