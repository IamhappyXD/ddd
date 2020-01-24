package student


func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	var prev *NodeL
	prev = nil
	var temp *NodeL
	temp = nil
	if l.Head.Data == data_ref {
		temp = l.Head.Next
		l.Head = nil
		l.Head = temp
	}
	it := l.Head

	for it != nil {
		if it.Data == data_ref && it == l.Tail {
			temp = l.Tail
			l.Tail = prev
			temp = nil
		}
		if it.Data == data_ref && it != l.Tail {
			temp = it
			prev.Next = it.Next
			temp = nil
			it = prev.Next
			continue
		}
		prev = it
		it = it.Next
	}
}

