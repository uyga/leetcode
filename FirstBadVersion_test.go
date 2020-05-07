package leetcode

import "testing"
import "reflect"

func TestFirstBadVersion(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {5,4},
        {4,4},
        {15,4},
    }
    for _,c := range cases {
        actual:=firstBadVersion(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
