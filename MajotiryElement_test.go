package leetcode

import "testing"
import "reflect"

func TestMajorityElement(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{3,3,4},3},
        {[]int{3,2,3},3},
        {[]int{2,2,1,1,1,2,2},2},
        {[]int{3,2,3,2},3},
        {[]int{3,2,3,2,1,1},1},
        {[]int{1,1},1},
        {[]int{1},1},
        {[]int{},0},
    }
    for _,c := range cases {
        actual:=majorityElement(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
