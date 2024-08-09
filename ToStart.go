package genericlist

// ToStart - Move to the start of the linked list.
func (list *LinkedList[T]) ToStart() *LinkedListNode[T] {
	list.lock.Lock()
	defer list.lock.Unlock()
	return list.head
}
