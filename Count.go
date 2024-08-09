package genericlist

// Count - Return the node count of the linked list.
func (list *LinkedList[T]) Count() uint {
	list.lock.Lock()
	defer list.lock.Unlock()
	return list.count
}
