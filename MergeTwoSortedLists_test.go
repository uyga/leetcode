package leetcode

import "testing"

func TestMergeTwoLists(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 []int
        expected []int
    }{
        {[]int{1,2,3},[]int{4,5,6},[]int{1,2,3,4,5,6}},
        {[]int{4,5,6},[]int{1,2,3},[]int{1,2,3,4,5,6}},
        {[]int{1,2,3,7,8,9},[]int{4,5,6},[]int{1,2,3,4,5,6,7,8,9}},
        {[]int{1,5,7},[]int{2,3,4,6,8,9},[]int{1,2,3,4,5,6,7,8,9}},
    }
    for _,c := range cases {
        actual:=mergeTwoLists(arrayToLinkedList(c.input1),arrayToLinkedList(c.input2))
        if !linkedListsAreEqual(actual, arrayToLinkedList(c.expected)) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, linkedListToArray(actual))
        }
    }
}
