package leetcode

import "testing"

func TestReverseLinkedList(t *testing.T) {
    cases := []struct {
        input []int
        expected []int
    }{
        {[]int{1,2,3,4,5},[]int{5,4,3,2,1}},
        {[]int{1},[]int{1}},
        {[]int{},[]int{}},
    }
    for _,c := range cases {
        actual:=reverseLinkedList(arrayToLinkedList(c.input))
        if !linkedListsAreEqual(actual, arrayToLinkedList(c.expected)) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, linkedListToArray(actual))
        }
    }
}
