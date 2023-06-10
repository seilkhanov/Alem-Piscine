package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	previous := l.Head
	current := l.Head

	for current != nil {
		if current.Data == data_ref {
			previous.Next = current.Next
		} else {
			previous = current
		}
		current = current.Next
	}

	if previous != nil {
		l.Tail = previous
	}
}
