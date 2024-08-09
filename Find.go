package genericlist

// Find - Traverse the linked list to find a node that satisfies the given compare function.
func (list *LinkedList[T]) Find(data T, compare func(lhs, rhs T) bool) bool {
	list.lock.Lock()
	defer list.lock.Unlock()

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
