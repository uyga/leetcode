package leetcode

import "testing"
import "reflect"

func TestTrailingZeroes(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {555,137},
        {0,0},
        {1,0},
        {53,12},
        {65536,16380},
    }
    for _,c := range cases {
        actual:=trailingZeroes(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
