package leetcode

import "testing"

func TestDeleteNode(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected []int
    }{
        {[]int{1,2,3,4,5,6},3,[]int{1,2,4,5,6}},
        {[]int{1,2,3,4,5,6},1,[]int{2,3,4,5,6}},
        {[]int{1,2,3,4,5,6},6,[]int{1,2,3,4,5}},
        {[]int{1,2,3,4,5,6},88,[]int{1,2,3,4,5,6}},
    }
    for _,c := range cases {
        l := arrayToLinkedList(c.input1)
        head := l
        var node *ListNode
        for head != nil {
            if head.Val == c.input2 {
                node = head
                break
            }
            head = head.Next
        }
        deleteNode(node)
        actual := l
        if !linkedListsAreEqual(actual, arrayToLinkedList(c.expected)) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, linkedListToArray(actual))
        }
    }
}
