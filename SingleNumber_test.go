package leetcode

import "testing"
import "reflect"

func TestSingleNumber(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{2,1,1,2,3,3,5,4,4,6,6},5},
    }
    for _,c := range cases {
        actual:=singleNumber(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
