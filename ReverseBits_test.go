package leetcode

import "testing"
import "reflect"

func TestReverseBits(t *testing.T) {
    cases := []struct {
        input uint32
        expected uint32
    }{
        {43261596,0x39782940},
    }
    for _,c := range cases {
        actual:=reverseBits(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
