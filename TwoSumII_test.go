package leetcode

import "testing"
import "reflect"

func TestTwoSumII(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected []int
    }{
        {[]int{2,7,11,15},9,[]int{1,2}},
    }
    for _,c := range cases {
        actual:=twoSumII(c.input1,c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
