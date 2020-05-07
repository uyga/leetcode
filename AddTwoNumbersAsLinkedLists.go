package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbersAsLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var sum, rem int
    var prevnode,head *ListNode
    for l1 != nil || l2 != nil || rem != 0 {
        sum = rem
        rem = 0
        if l1 != nil {
            sum+=l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum+=l2.Val
            l2 = l2.Next
        }
        if sum > 9 {
            rem = sum/10
            sum = sum%10
        }
        node := new(ListNode)
        node.Val = sum
        if prevnode != nil {
            prevnode.Next = node
        }
        if head == nil {
            head = node
        }
        prevnode = node
    }
    return head
}
