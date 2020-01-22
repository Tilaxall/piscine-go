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
	result := 0
	for l.Head != nil {
		result++
		l.Head = l.Head.Next
	}
	return result
}
