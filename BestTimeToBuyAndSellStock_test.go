package leetcode

import "testing"
import "reflect"

func TestMaxProfit(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{7,1,5,3,6,4},5},
        {[]int{7,6,4,3,1},0},
        {[]int{},0},
        {[]int{1,4},3},
    }
    for _,c := range cases {
        actual:=maxProfit(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
