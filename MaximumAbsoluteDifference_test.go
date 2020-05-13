package leetcode

import "testing"
import "reflect"

func TestMaxAbsoluteDifference(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{1,3,-1},5},
        {[]int{55, -8, 43, 52, 8, 59, -91, -79, -18, -94},158},
    }
    for _,c := range cases {
        actual:=maxAbsoluteDifference(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
