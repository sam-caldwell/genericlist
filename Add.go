package genericlist

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
		list.curr.nextPtr = newNode

		newNode.prevPtr = list.curr
		if newNode.nextPtr != nil {
			newNode.nextPtr.prevPtr = newNode
		} else {
			list.tail = newNode // Update tail if the new node is the last one
		}
	}
	list.count++ // Increment count
}
