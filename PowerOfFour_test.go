package leetcode

import "testing"
import "reflect"

func TestIsPowerOfFour(t *testing.T) {
    cases := []struct {
        input int
        expected bool
    }{
        {-1,false},
        {0,false},
        {1,true},
        {2,false},
        {4,true},
        {8,false},
        {16,true},
        {65536,true},
        {255,false},
        {128,false},
    }
    for _,c := range cases {
        actual:=isPowerOfFour(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
