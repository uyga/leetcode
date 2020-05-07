package leetcode

import "testing"
import "reflect"

func TestHammingWeight(t *testing.T) {
    cases := []struct {
        input uint32
        expected int
    }{
        {21,3},
    }
    for _,c := range cases {
        actual:=hammingWeight(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
