package leetcode

import "testing"
import "reflect"

func TestDivide(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{0,300},0},
        {[]int{10,3},3},
        {[]int{10,-3},-3},
        {[]int{-10,3},-3},
        {[]int{-10,-3},3},
        {[]int{-2147483647,1},-2147483647},
        {[]int{2147483647,1},2147483647},
        {[]int{-2147483648,-1},2147483647},
    }
    for _,c := range cases {
        actual:=divide(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
