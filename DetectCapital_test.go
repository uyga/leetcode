package leetcode

import "testing"
import "reflect"

func TestDetectCapitalUse(t *testing.T) {
    cases := []struct {
        input string
        expected bool
    }{
        {"USA",true},
        {"leetcode",true},
        {"FlaG",false},
        {"Flag",true},
        {"fLag",false},
        {"flaG",false},
        {"",true},
    }
    for _,c := range cases {
        actual:=detectCapitalUse(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
