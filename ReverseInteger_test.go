package leetcode

import "testing"
import "reflect"

func TestReverseInteger(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {120,21},
    }
    for _,c := range cases {
        actual:=reverseInteger(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
