package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	key := l
	/* indexlist := ListSize()
	if indexlist > pos {
		return nil
	} */
	if pos < 0 {
		return nil
	}
	for key != nil {
		if pos == 0 {
			return key
		}
		key, pos = key.Next, pos-1
	}
	return nil
}
