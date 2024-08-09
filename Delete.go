package genericlist

// Delete - Delete the current node from the linked list.
func (list *LinkedList[T]) Delete() {
	list.lock.Lock()
	defer list.lock.Unlock()

	if list.curr == nil {
		return
	}

	// Update the previous node's next pointer if it exists
	if list.curr.prevPtr != nil {
		list.curr.prevPtr.nextPtr = list.curr.nextPtr
	} else {
		// The node to be deleted is the head
		list.head = list.curr.nextPtr
	}

	// Update the next node's previous pointer if it exists
	if list.curr.nextPtr != nil {
		list.curr.nextPtr.prevPtr = list.curr.prevPtr
	} else {
		// The node to be deleted is the tail
		list.tail = list.curr.prevPtr
	}

	// Move curr to a new valid position
	if list.curr.nextPtr != nil {
		list.curr = list.curr.nextPtr
	} else if list.curr.prevPtr != nil {
		list.curr = list.curr.prevPtr
	} else {
		// The list is now empty
		list.curr = nil
	}

	list.count-- // Decrement count
}
