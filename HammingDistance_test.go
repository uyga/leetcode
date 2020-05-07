package leetcode

import "testing"
import "reflect"

func TestHammingDistance(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{1,0},1},
        {[]int{1,5},1},
        {[]int{2,9},3},
        {[]int{100,65536},4},
    }
    for _,c := range cases {
        actual:=hammingDistance(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
