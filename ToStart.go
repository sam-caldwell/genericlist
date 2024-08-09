package genericlist

// ToStart - Move to the start of the linked list.
func (list *LinkedList[T]) ToStart() {
	list.lock.Lock()
	defer list.lock.Unlock()
	list.curr = list.head
}
