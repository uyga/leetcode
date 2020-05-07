package leetcode

import "testing"
import "reflect"

func TestCanConstruct(t *testing.T) {
    cases := []struct {
        input []string
        expected bool
    }{
        {[]string{"aan","aabbnn"},true},
        {[]string{"aan","opopopopq"},false},
    }
    for _,c := range cases {
        actual:=canConstruct(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
