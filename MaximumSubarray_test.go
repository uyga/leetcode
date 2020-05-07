package leetcode

import "testing"
import "reflect"

func TestMaxSubArray(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{-2,1,-3,4,-1,2,1,-5,4},6},
        {[]int{-2},-2},
        {[]int{},-2147483648},
    }
    for _,c := range cases {
        actual:=maxSubArray(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
