package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	temp := make([]int, 0)
	p := head
	for p != nil {
		temp = append(temp, p.Val)
		p = p.Next
	}
	//p = head
	for i := 0; i < len(temp)/2; i++ {
		if temp[i] != temp[len(temp)-1-i] {
			return false
		}
	}
	return true
}
