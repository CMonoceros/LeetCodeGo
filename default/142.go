package main

// Linked List
// Two Pointers

// todo:
// 8ms
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fp, sp := head, head
	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
		if fp == sp {
			break
		}
	}
	if fp == nil || fp.Next == nil {
		return nil
	}
	sp = head
	for fp != sp {
		fp = fp.Next
		sp = sp.Next
	}
	return fp
}
