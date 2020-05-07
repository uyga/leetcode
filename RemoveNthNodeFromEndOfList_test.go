package leetcode

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected []int
    }{
        {[]int{1,2,3,4,5,6,7,8,9},3,[]int{1,2,3,4,5,6,8,9}},
        {[]int{1,2,3,4,5,6,7,8,9},1,[]int{1,2,3,4,5,6,7,8}},
        {[]int{1,2,3,4,5,6,7,8,9},9,[]int{2,3,4,5,6,7,8,9}},
    }
    for _,c := range cases {
        actual:=removeNthFromEnd(arrayToLinkedList(c.input1),c.input2)
        if !linkedListsAreEqual(actual, arrayToLinkedList(c.expected)) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, linkedListToArray(actual))
        }
    }
}
