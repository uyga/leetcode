package leetcode

import "testing"
import "reflect"

func TestProductExceptSelf(t *testing.T) {
    cases := []struct {
        input []int
        expected []int
    }{
        {[]int{1,2,3,4,5},[]int{120, 60, 40, 30, 24}},
        {[]int{5,4,3,2,1},[]int{24, 30, 40, 60, 120}},
        {[]int{1},[]int{1}},
    }
    for _,c := range cases {
        actual:=productExceptSelf(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
