package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//O(N) + O(N)
func middleNode2(head *ListNode) *ListNode {
    var stack []*ListNode
    for head != nil {
        stack = append(stack, head)
        head = head.Next
    }
    return stack[len(stack)/2]
}

//O(N) + O(1)
func middleNode(head *ListNode) *ListNode {
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
