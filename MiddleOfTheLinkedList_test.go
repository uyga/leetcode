package leetcode

import "testing"

func TestMiddleNode(t *testing.T) {
    cases := []struct {
        input []int
        expected []int
    }{
        {[]int{1,2,3,4,5},[]int{3,4,5}},
        {[]int{1,2,3,4,5,6},[]int{4,5,6}},
        {[]int{1,2},[]int{2}},
        {[]int{1},[]int{1}},
        {[]int{1,2,3,4,5,6,7,8,9,10,11,12,14,14},[]int{8,9,10,11,12,14,14}},
    }
    for _,c := range cases {
        actual:=middleNode(arrayToLinkedList(c.input))
        if !linkedListsAreEqual(actual, arrayToLinkedList(c.expected)) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, linkedListToArray(actual))
        }
    }
}
