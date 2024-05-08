package main

func removeElements(head *ListNode, val int) *ListNode {

	xshHead := &ListNode{Next: head}
	for p := xshHead; p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return xshHead.Next
}
