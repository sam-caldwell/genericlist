package genericlist

// Next - Move to the next node, return false if nil.
func (list *LinkedList[T]) Next() bool {
	list.lock.Lock()
	defer list.lock.Unlock()

	return list.moveNext()
}

func (list *LinkedList[T]) moveNext() bool {
	if list.curr.nextPtr != nil {
		list.curr = list.curr.nextPtr
		return true
	}
	return false
}
