package leetcode

import "testing"
import "reflect"

func TestToLowerCase(t *testing.T) {
    cases := []struct {
        input string
        expected string
    }{
        {"abc","abc"},
        {"Abc","abc"},
        {"ABC","abc"},
        {"abC","abc"},
        {"",""},
    }
    for _,c := range cases {
        actual:=toLowerCase(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
