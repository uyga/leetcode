package leetcode

import "testing"
import "reflect"

func TestIsUgly(t *testing.T) {
    cases := []struct {
        input int
        expected bool
    }{
        {6,true},
        {7,false},
        {65536,true},
        {0,false},
    }
    for _,c := range cases {
        actual:=isUgly(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
