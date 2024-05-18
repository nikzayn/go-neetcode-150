package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// TC - O(N) | SC - O(1)
func ReverseLinkedListIterative(head *ListNode) *ListNode {
	var revHead *ListNode
	for head != nil {
		head.Next, revHead, head = revHead, head, head.Next
	}
	return revHead
}
