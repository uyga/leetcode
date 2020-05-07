package leetcode

import "testing"
import "reflect"

func TestMySqrt(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {4,2},
        {8,2},
        {9,3},
        {1511615877,38879},
    }
    for _,c := range cases {
        actual:=mySqrt(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
