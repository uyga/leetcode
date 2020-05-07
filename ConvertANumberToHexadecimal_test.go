package leetcode

import "testing"
import "reflect"

func TestToHex(t *testing.T) {
    cases := []struct {
        input int
        expected string
    }{
        {128,"80"},
        {12234,"2fca"},
        {1,"1"},
        {0,"0"},
        {-128,"ffffff80"},
    }
    for _,c := range cases {
        actual:=toHex(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
