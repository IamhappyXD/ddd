package student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	link := &NodeI{}
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Data.(int) <= n2.Data.(int) {
			link =  n1
			link.Next = SortedListMerge(n1.Next, n2)
	} else {
			link = n2
			link.Next = SortedListMerge(n1, n2.Next)
	}
	
	return link
}
