package leetcode

import "testing"
import "reflect"

func TestRob(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{1,2,3,1},4},
        {[]int{2,7,9,3,1},12},
        {[]int{2,7,9,3,1,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9},60},
    }
    for _,c := range cases {
        actual:=rob(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
