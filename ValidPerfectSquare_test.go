package leetcode

import "testing"
import "reflect"

func TestIsPerfectSquare(t *testing.T) {
    cases := []struct {
        input int
        expected bool
    }{
        {25,true},
        {0,false},
        {1,true},
        {2,false},
        {4,true},
        {9,true},
        {14,false},
        {16,true},
    }
    for _,c := range cases {
        actual:=isPerfectSquare(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
