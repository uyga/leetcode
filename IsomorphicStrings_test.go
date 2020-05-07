package leetcode

import "testing"
import "reflect"

func TestIsIsomorphic(t *testing.T) {
    cases := []struct {
        input []string
        expected bool
    }{
        {[]string{"aabbnn", "llsspp"},true},
        {[]string{"aabbnn", "llsspa"},false},
    }
    for _,c := range cases {
        actual:=isIsomorphic(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
