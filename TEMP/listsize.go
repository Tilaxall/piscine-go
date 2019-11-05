package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	res := 0
	for l.Head != nil {
		res++
		l.Head = l.Head.Next
	}
	return res
}
