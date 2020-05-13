package leetcode

import "testing"
import "reflect"

func TestMaximum69Number(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {6666,9666},
        {9696,9996},
        {9,9},
        {6,9},
    }
    for _,c := range cases {
        actual:=maximum69Number(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
