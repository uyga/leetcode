package leetcode

import "testing"
import "reflect"

func TestIsReverse(t *testing.T) {
    cases := []struct {
        input []string
        expected bool
    }{
        {[]string{"reverse","esrever"},true},
        {[]string{"reverse","esreverqw"},false},
        {[]string{"reverse","reverse"},false},
        {[]string{"",""},true},
    }
    for _,c := range cases {
        actual:=isReverse(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
