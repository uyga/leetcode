package leetcode

import "testing"
import "reflect"

func TestIsHappy(t *testing.T) {
    cases := []struct {
        input int
        expected bool
    }{
        {119,false},
        {19,true},
        {18,false},
        {10,true},
        {5,false},
    }
    for _,c := range cases {
        actual:=isHappy(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
