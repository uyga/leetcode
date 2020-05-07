package leetcode

import "testing"
import "reflect"

func TestHasCycle(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected bool
    }{
        {[]int{1,2,3,4,5,6},2,true},
        {[]int{1,2,3,4,5,6},88,false},
        {[]int{1},0,true},
        {[]int{1},-1,false},
        {[]int{},-1,false},
    }
    for _,c := range cases {
        actual:=hasCycle(createCycledLinkedList(c.input1,c.input2))
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
