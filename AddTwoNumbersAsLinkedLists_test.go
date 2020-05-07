package leetcode

import "testing"

func TestAddTwoNumbersAsLinkedLists(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 []int
        expected []int
    }{
        {[]int{1,2,3},[]int{4,5,6},[]int{5,7,9}},
        {[]int{1,2,3,4,5},[]int{1},[]int{2,2,3,4,5}},
        {[]int{5},[]int{5},[]int{0,1}},
        {[]int{0},[]int{5,5,5},[]int{5,5,5}},
        {[]int{},[]int{5,5,5},[]int{5,5,5}},
        {[]int{1,1,1,1},[]int{},[]int{1,1,1,1}},
    }
    for _,c := range cases {
        actual:=addTwoNumbersAsLinkedLists(arrayToLinkedList(c.input1),arrayToLinkedList(c.input2))
        if !linkedListsAreEqual(actual, arrayToLinkedList(c.expected)) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, linkedListToArray(actual))
        }
    }
}
