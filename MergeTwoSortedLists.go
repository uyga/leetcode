package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var res *ListNode
    if l1 == nil && l2 == nil {
        res = nil
    } else if l1 != nil && l2 == nil {
        res = l1
    } else if l1 == nil && l2 != nil {
        res = l2
    } else {
        if l1.Val < l2.Val {
            l1.Next = mergeTwoLists(l1.Next, l2)
            res = l1
        } else {
            l2.Next = mergeTwoLists(l1, l2.Next)
            res = l2
        }
    }
    return res
}
