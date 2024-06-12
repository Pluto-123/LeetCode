package main

func hasCycle(head *ListNode) bool {
	mp := make(map[*ListNode]bool)
	p := head
	for p != nil {
		if mp[p] {
			return true
		} else {
			mp[p] = true
		}
		p = p.Next
	}
	return false
}
