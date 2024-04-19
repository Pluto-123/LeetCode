package main

// fail
func reorderList(head *ListNode) {
	// 先把尾部的节点截取出来
	p1 := head
	p2 := head.Next
	for {
		if p2.Next == nil {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	p1.Next = nil

	// 插入
	p2.Next = head.Next
	head.Next = p2
}
