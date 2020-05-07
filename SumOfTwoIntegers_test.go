package leetcode

import "testing"
import "reflect"

func TestGetSum(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{126, 5443},5569},
        {[]int{1, 0},1},
        {[]int{-1, -100},-101},
        {[]int{128, 65536},65664},
    }
    for _,c := range cases {
        actual:=getSum(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
