package leetcode

import "testing"
import "reflect"

func TestTitleToNumber(t *testing.T) {
    cases := []struct {
        input string
        expected int
    }{
        {"A",1},
        {"Z",26},
        {"ZY",701},
        {"ZYA",18227},
        {"AAB",704},
    }
    for _,c := range cases {
        actual:=titleToNumber(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
