package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition3(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	ptr := head
	ans := &ListNode{}
	ansp := ans
	for ptr != nil {
		if ptr.Val < x {
			ansp.Val = ptr.Val
			ansp.Next = &ListNode{}
			ansp = ansp.Next
		}
		ptr = ptr.Next
	}
	ptr = head
	for ptr != nil {
		if ptr.Val >= x {
			ansp.Val = ptr.Val
			ansp.Next = &ListNode{}
			ansp = ansp.Next
		}
		ptr = ptr.Next
	}

	for p1 := ans; p1 != nil; p1 = p1.Next {
		if p1.Next.Next == nil {
			p1.Next = nil
		}
	}

	return ans
}

func main() {
	partition3(nil, 2)
}
