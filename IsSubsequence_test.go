package leetcode

import "testing"
import "reflect"

func TestIsSubsequence(t *testing.T) {
    cases := []struct {
        input []string
        expected bool
    }{
        {[]string{"abc","ahbgdc"},true},
        {[]string{"axc","ahbgdc"},false},
        {[]string{"ac","bhbacgdf"},true},
        {[]string{"ac","ac"},true},
        {[]string{"ac","jskhfkjhhfkjsdhfjsdhfjshkdfhjsdhfkjshdjfhkjdf"},false},
        {[]string{"ac",""},false},
        {[]string{"","ac"},true},
        {[]string{"",""},true},
    }
    for _,c := range cases {
        actual:=isSubsequence(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
