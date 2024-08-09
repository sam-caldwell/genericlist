package genericlist

// Find - Perform a simple linear search from head to tail for a given item using the provided comparison function.
func (list *LinkedList[T]) Find(data T, compare func(lhs, rhs T) bool) bool {
	list.lock.Lock()
	defer list.lock.Unlock()

	//move to start
	list.curr = list.head

	for list.curr != nil {
		if compare(list.curr.data, data) {
			return true
		}
		if ok := list.moveNext(); !ok {
			return false
		}
	}
	return false
}
