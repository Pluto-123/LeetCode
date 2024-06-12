package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	mp := make(map[*ListNode]bool)
	p := head
	for p != nil {
		if mp[p] {
			return p
		} else {
			mp[p] = true
		}
		p = p.Next
	}
	return nil
}
