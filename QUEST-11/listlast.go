package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	list := l.Head
	for list.Next != nil {
		list = list.Next
	}
	return list.Data
}
