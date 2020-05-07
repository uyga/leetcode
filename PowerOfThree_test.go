package leetcode

import "testing"
import "reflect"

func TestIsPowerOfThree(t *testing.T) {
    cases := []struct {
        input int
        expected bool
    }{
        {27,true},
        {45,false},
    }
    for _,c := range cases {
        actual:=isPowerOfThree(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
