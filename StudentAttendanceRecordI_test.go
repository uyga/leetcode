package leetcode

import "testing"
import "reflect"

func TestCheckRecord(t *testing.T) {
    cases := []struct {
        input string
        expected bool
    }{
        {"PPALLP",true},
        {"PPALLL",false},
    }
    for _,c := range cases {
        actual:=checkRecord(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
