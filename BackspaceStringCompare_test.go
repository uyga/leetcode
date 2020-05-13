package leetcode

import "testing"
import "reflect"

func TestBackspaceCompare(t *testing.T) {
    cases := []struct {
        input []string
        expected bool
    }{
        {[]string{"ab#c","ad#c"},true},
        {[]string{"ab##","c#d#"},true},
        {[]string{"a##c","#a#c"},true},
        {[]string{"a#c","b"},false},
    }
    for _,c := range cases {
        actual:=backspaceCompare(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
