package leetcode

import "testing"
import "reflect"

func TestFactorial(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {0,1},
        {5,120},
        {10,3628800},
        {13,6227020800},
    }
    for _,c := range cases {
        actual:=factorial(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
