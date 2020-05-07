package leetcode

import "testing"
import "reflect"

func TestAddDigitsUntilOne(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {38,2},
        {382435,7},
        {0,0},
        {7,7},
    }
    for _,c := range cases {
        actual:=addDigitsUntilOne(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
