package main

import "sort"

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := &ListNode{}
	resp := res
	temp := make([]int, 0)
	for p := head; p != nil; p = p.Next {
		temp = append(temp, p.Val)
	}
	sort.Ints(temp)
	for i := 0; i < len(temp); i++ {
		resp.Val = temp[i]
		if i != len(temp)-1 {
			resp.Next = &ListNode{}
			resp = resp.Next
		}
	}
	return res
}
