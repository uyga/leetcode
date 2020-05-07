package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    var hashA map[*ListNode]bool = make(map[*ListNode]bool)
    var hashB map[*ListNode]bool = make(map[*ListNode]bool)
    for headA != nil || headB != nil {
        hashA[headA]=true
        hashB[headB]=true
        _,okA:=hashA[headB]
        _,okB:=hashB[headA]
        if okA {
            return headB
        }
        if okB {
            return headA
        }
        if headA.Next != nil {
            headA = headA.Next
        }
        if headB.Next != nil {
            headB = headB.Next
        }
        if headA.Next == nil && headB.Next == nil {
            if headA == headB {
                return headA
            }
            break
        }
    }
    return nil
}
