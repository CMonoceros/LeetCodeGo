package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// todo:1
// 20ms
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{0, nil}
	var temp = &ListNode{0, &ListNode{0, nil}}
	res = temp.Next
	var last = 0
	for l1 != nil || l2 != nil || last != 0 {
		num := last
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		last = num / 10
		if temp.Next == nil {
			temp.Next = &ListNode{num % 10, nil}
		} else {
			temp.Next.Val = num % 10
		}
		temp = temp.Next
	}
	return res
}
