package genericlist

// Get - Return the data at the current node
func (list *LinkedList[T]) Get() T {
	if list.curr == nil {
		var empty T
		return empty
	}
	return list.curr.data
}
