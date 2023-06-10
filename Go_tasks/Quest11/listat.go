package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}

	for i := 0; i < pos; i++ {
		if l.Next == nil {
			return nil
		}

		l = l.Next
	}
	return l
}
