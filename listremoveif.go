package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil {
		return
	}

	current := l.Head
	var previous *NodeL

	for current != nil {
		if current.Data == data_ref {
			
			if previous == nil {
				l.Head = current.Next
			} else {
				previous.Next = current.Next
			}

			if current.Next == nil {
				l.Tail = previous
			}

			current = current.Next
			continue
		}

		previous = current
		current = current.Next
	}
}
