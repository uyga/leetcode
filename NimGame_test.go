package leetcode

import "testing"
import "reflect"

func TestCanWinNim(t *testing.T) {
    cases := []struct {
        input int
        expected bool
    }{
        {4,false},
        {14,true},
        {22,true},
        {30,true},
        {28,false},
        {100,false},
        {65536,false},
    }
    for _,c := range cases {
        actual:=canWinNim(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
