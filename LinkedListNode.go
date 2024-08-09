package genericlist

// LinkedListNode - Generic Linked list node
type LinkedListNode[T any] struct {
	data    T
	nextPtr *LinkedListNode[T]
	prevPtr *LinkedListNode[T]
}
