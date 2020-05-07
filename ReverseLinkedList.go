package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseLinkedList(head *ListNode) *ListNode {
    var curr,prev *ListNode
    prev = nil
    for head != nil {
        curr = head
        if head.Next == nil {
            head.Next = prev
            break
        }
        head = head.Next
        curr.Next = prev
        prev = curr
    }
    return head
}
