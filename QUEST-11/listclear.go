package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListClear(l *List) {
	if l.Head != nil || l.Tail != nil {
		l.Head = nil
		l.Tail = nil
	}
}
