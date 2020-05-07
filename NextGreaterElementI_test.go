package leetcode

import "testing"
import "reflect"

func TestNextGreaterElement(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 []int
        expected []int
    }{
        {[]int{4,1,2},[]int{1,3,4,2},[]int{-1,3,-1}},
        {[]int{2,4},[]int{1,2,3,4},[]int{3,-1}},
        {[]int{2,3,4},[]int{5,6,7,8},[]int{-1,-1,-1}},
    }
    for _,c := range cases {
        actual:=nextGreaterElement(c.input1,c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
