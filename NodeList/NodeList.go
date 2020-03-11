package NodeList

type ListNode struct {
	Val  int
	Next *ListNode
}
//循环
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	next := new(ListNode)
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}


