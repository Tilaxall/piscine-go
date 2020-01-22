package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	templist := l.Head
	if templist == nil {
		return nil
	}
	for templist != nil {
		if CompStr(templist.Data, ref) {
			return &templist.Data
		}
		templist = templist.Next
	}
	return &templist.Data
}
