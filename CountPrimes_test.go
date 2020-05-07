package leetcode

import "testing"
import "reflect"

func TestCountPrimes(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {0,0},
        {100,25},
        {499979,41537},
    }
    for _,c := range cases {
        actual:=countPrimes(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
