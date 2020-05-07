package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeLinkedList(head *ListNode) bool {
    var stack []int
    node := head
    for node != nil {
        stack = append(stack, node.Val)
        node = node.Next
    }
    node = head
    for node != nil {
        if len(stack) > 0 {
            if node.Val == stack[len(stack)-1] {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
        node = node.Next
    }
    return len(stack) == 0
}
