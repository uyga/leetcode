package leetcode

import "testing"
import "reflect"

func TestFindComplement(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {5,2},
        {55,8},
    }
    for _,c := range cases {
        actual:=findComplement(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
