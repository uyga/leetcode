package leetcode

import "testing"
import "reflect"

func TestIsValid(t *testing.T) {
    cases := []struct {
        input string
        expected bool
    }{
        {"(",false},
        {"]",false},
        {"()",true},
        {"",true},
        {"(]",false},
        {"(([]))",true},
        {"()[]{}",true},
    }
    for _,c := range cases {
        actual:=isValid(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
