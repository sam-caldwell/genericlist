package genericlist

import "sync"

// LinkedList - Wrapper for the linked list
type LinkedList[T any] struct {
	lock  sync.Mutex
	head  *LinkedListNode[T]
	curr  *LinkedListNode[T]
	tail  *LinkedListNode[T]
	count uint
}
