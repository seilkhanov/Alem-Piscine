package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	for l.Head != nil {
		if CompStr(l.Head.Data, ref) {
			return &l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return nil
}
