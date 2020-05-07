package leetcode

import "testing"
import "reflect"

func TestGetRow(t *testing.T) {
    cases := []struct {
        input int
        expected []int
    }{
        {9,[]int{1, 9, 36, 84, 126, 126, 84, 36, 9, 1}},
    }
    for _,c := range cases {
        actual:=getRow(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
