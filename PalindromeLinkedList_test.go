package leetcode

import "testing"
import "reflect"

func TestIsPalindromeLinkedList(t *testing.T) {
    cases := []struct {
        input []int
        expected bool
    }{
        {[]int{1,1,1},true},
        {[]int{1,2},false},
        {[]int{1,2,2,1},true},
        {[]int{1,2,2,1,1},false},
        {[]int{},true},
        {[]int{1},true},
    }
    for _,c := range cases {
        actual:=isPalindromeLinkedList(arrayToLinkedList(c.input))
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
