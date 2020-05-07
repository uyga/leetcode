package leetcode

import "testing"

func TestDeleteDuplicatesFromLinkedList(t *testing.T) {
    cases := []struct {
        input []int
        expected []int
    }{
        {[]int{1,2,3,3,3,3,4,5},[]int{1,2,3,4,5}},
        {[]int{1,2,3,3,3,3},[]int{1,2,3}},
        {[]int{3,3,3,3,4,5},[]int{3,4,5}},
        {[]int{1,2,3},[]int{1,2,3}},
        {[]int{3,3,3,3},[]int{3}},
        {[]int{},[]int{}},
    }
    for _,c := range cases {
        actual:=deleteDuplicatesFromLinkedList(arrayToLinkedList(c.input))
        if !linkedListsAreEqual(actual, arrayToLinkedList(c.expected)) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, linkedListToArray(actual))
        }
    }
}
