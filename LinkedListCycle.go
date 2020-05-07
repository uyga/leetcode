package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle2(head *ListNode) bool {
    var hash map[*ListNode]bool = make(map[*ListNode]bool)
    for head != nil {
        _, ok := hash[head]
        if ok {
            return true
        }
        hash[head] = true
        head = head.Next
    }
    return false
}

func hasCycle(head *ListNode) bool {
    var slow, fast *ListNode = head, head
    for slow != nil && fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
