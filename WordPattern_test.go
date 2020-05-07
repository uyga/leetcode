package leetcode

import "testing"
import "reflect"

func TestWordPattern(t *testing.T) {
    cases := []struct {
        input []string
        expected bool
    }{
        {[]string{"abba","dog cat cat dog"},true},
        {[]string{"abba","dog cat cat dogaa"},false},
        {[]string{"","dog cat cat dogaa"},false},
        {[]string{"aaa",""},false},
        {[]string{"",""},false},
    }
    for _,c := range cases {
        actual:=wordPattern(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
