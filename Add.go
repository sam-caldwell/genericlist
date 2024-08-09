package genericlist

// Add - insert record at current position
func (list *LinkedList[T]) Add(data T) {
	list.lock.Lock()
	defer list.lock.Unlock()

	newNode := &LinkedListNode[T]{data: data}

	if list.curr == nil { // List is empty
		list.head = newNode
		list.tail = newNode
		list.curr = newNode
	} else {
		newNode.nextPtr = list.curr.nextPtr
		if list.curr.nextPtr != nil {
			list.curr.nextPtr.prevPtr = newNode
		} else {
			list.tail = newNode // Update tail if the new node is the last one
		}
		list.curr.nextPtr = newNode
		newNode.prevPtr = list.curr
		list.curr = newNode // Move curr to the new node
	}
	list.count++ // Increment count
	list.sorted = false
}
