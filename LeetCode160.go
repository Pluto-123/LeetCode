package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mp := make(map[*ListNode]bool)
	p := headA
	for p != nil {
		mp[p] = true
		p = p.Next
	}
	q := headB
	for q != nil {
		if mp[q] {
			return q
		}
		q = q.Next
	}
	return nil
}
