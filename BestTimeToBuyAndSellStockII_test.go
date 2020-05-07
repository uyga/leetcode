package leetcode

import "testing"
import "reflect"

func TestMaxProfitII(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{7,1,5,3,6,4},7},
        {[]int{1,2,3,4,5},4},
        {[]int{5,4,3,2,1},0},
    }
    for _,c := range cases {
        actual:=maxProfitII(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
