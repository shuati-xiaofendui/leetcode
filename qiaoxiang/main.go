package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	cur := head

	for i := 1; i < k; i++ {
		cur = cur.Next
		if cur == nil {
			return head
		}
	}
	nextHead := cur.Next
	tail := reverseList(head, cur)
	tail.Next = reverseKGroup(nextHead, k)
	return cur
}

func reverseList(head, tail *ListNode) *ListNode {
	var pre, cur *ListNode
	for pre != tail {
		cur = head.Next
		head.Next = pre
		pre = head
		head = cur
	}
	for pre.Next != nil {
		pre = pre.Next
	}
	return pre
}
