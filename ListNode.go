package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func arrayToLinkedList(data []int) *ListNode {
    var head *ListNode
    if len(data) > 0 {
        head = new(ListNode)
        head.Val = data[0]
        prev := head
        for i:=1;i<len(data);i++ {
            node := new(ListNode)
            node.Val = data[i]
            prev.Next = node
            prev = node
        }
    }
    return head
}

func linkedListToArray(head *ListNode) []int {
    var res []int
    for head != nil {
        res = append(res, head.Val)
        head = head.Next
    }
    return res
}

func linkedListsAreEqual(head1, head2 *ListNode) bool {
    for head1 != nil && head2 != nil {
        if head1 == nil || head2 == nil {
            return false
        }
        if head1.Val != head2.Val {
            return false
        }
        head1 = head1.Next
        head2 = head2.Next
    }
    return true
}

func createCycledLinkedList(data []int, cycleNodeIndex int) *ListNode {
    var head,cycleNodePointer *ListNode
    if len(data) > 0 {
        head = new(ListNode)
        head.Val = data[0]
        prev := head
        for i:=1;i<len(data);i++ {
            node := new(ListNode)
            node.Val = data[i]
            prev.Next = node
            prev = node
            if i == cycleNodeIndex {
                cycleNodePointer = prev
            }
        }
        if len(data) == 1 && cycleNodeIndex == 0 {
            prev.Next = prev
        } else {
            prev.Next = cycleNodePointer
        }
    }
    return head
}
