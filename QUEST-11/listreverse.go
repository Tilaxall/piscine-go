package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	templist := &List{}
	for l.Head != nil {
		ListPushFront(templist, l.Head.Data)
		l.Head = l.Head.Next
	}
	l.Head, l.Tail = templist.Head, templist.Tail
}
func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = n, n
		return
	} else {
		n.Next, l.Head = l.Head, n
	}

}
