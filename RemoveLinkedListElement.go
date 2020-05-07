package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElementsFromLinkedList(head *ListNode, val int) *ListNode {
    var res *ListNode
    if head != nil {
        for head.Next != nil && head.Val == val {
            head = head.Next
        }
        if head.Val != val {
            res = head
        }
        for head.Next != nil {
            if head.Next.Val == val {
                for head.Next != nil && head.Next.Val == val {
                    head.Next = head.Next.Next
                }
            }
            head = head.Next
            if head == nil {
                break
            }
        }
    }
    return res
}
