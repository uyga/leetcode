package leetcode

import "testing"
import "reflect"

func TestIsPalindromeNumber(t *testing.T) {
    cases := []struct {
        input int
        expected bool
    }{
        {1221,true},
        {12223,false},
        {1,true},
        {0,true},
        {012223,false},
        {122210,false},
    }
    for _,c := range cases {
        actual:=isPalindromeNumber(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
