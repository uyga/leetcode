package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var res,next *ListNode
    var cnt int = 0
    var arr []*ListNode
    res = head
    next = head
    if n > 0 && head != nil {
        for next != nil {
            arr = append(arr,next)
            cnt++
            next = next.Next
        }
        if cnt == 1 {
            res = nil
        } else if cnt-n+1 >= cnt {
            arr[cnt-n-1].Next = nil
        } else if cnt-n-1 < 0 {
            arr[0].Next = nil
            res = arr[1]
        } else {
            arr[cnt-n-1].Next = arr[cnt-n+1]
        }
    }
    return res
}
