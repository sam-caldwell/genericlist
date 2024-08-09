package genericlist

// Delete - Delete the specified node from the linked list.
func (list *LinkedList[T]) Delete() {
	list.lock.Lock()
	defer list.lock.Unlock()

	if list.curr == nil {
		return
	}

	// Update pointers
	if list.curr.prevPtr != nil {
		list.curr.prevPtr.nextPtr = list.curr.nextPtr
	} else {
		list.head = list.curr.nextPtr
	}

	if list.curr.nextPtr != nil {
		list.curr.nextPtr.prevPtr = list.curr.prevPtr
	} else {
		list.tail = list.curr.prevPtr
	}

	// Move curr to a new valid position
	if list.curr.nextPtr != nil {
		list.curr = list.curr.nextPtr
	} else if list.curr.prevPtr != nil {
		list.curr = list.curr.prevPtr
	} else {
		// List is now empty
		list.curr = nil
	}

	list.count-- // Decrement count
}
