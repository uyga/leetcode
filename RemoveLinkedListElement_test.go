package leetcode

import "testing"

func TestRemoveElementsFromLinkedList(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected []int
    }{
        {[]int{1,2,3,4,5,6},2,[]int{1,3,4,5,6}},
        {[]int{1,3,3,4,5,6},2,[]int{1,3,3,4,5,6}},
        {[]int{1,2,2,2,2,2},2,[]int{1}},
        {[]int{},2,[]int{}},
        {[]int{2},2,[]int{}},
    }
    for _,c := range cases {
        actual:=removeElementsFromLinkedList(arrayToLinkedList(c.input1),c.input2)
        if !linkedListsAreEqual(actual, arrayToLinkedList(c.expected)) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, linkedListToArray(actual))
        }
    }
}
