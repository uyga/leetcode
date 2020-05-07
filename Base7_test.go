package leetcode

import "testing"
import "reflect"

func TestConvertToBase7(t *testing.T) {
    cases := []struct {
        input int
        expected string
    }{
        {7,"10"},
        {0,"0"},
        {8,"11"},
        {100,"202"},
        {-7,"-10"},
        {1e+7,"150666343"},
    }
    for _,c := range cases {
        actual:=convertToBase7(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
