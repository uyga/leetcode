package leetcode

import "testing"
import "reflect"

func TestIsPowerOfTwo(t *testing.T) {
    cases := []struct {
        input int
        expected bool
    }{
        {-16,false},
        {1,true},
        {2,true},
        {3,false},
        {4,true},
        {7,false},
        {8,true},
        {16,true},
        {32,true},
        {64,true},
        {218,false},
    }
    for _,c := range cases {
        actual:=isPowerOfTwo(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
