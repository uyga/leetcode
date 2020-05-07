package leetcode

import "testing"
import "reflect"

func TestFindTheDifference(t *testing.T) {
    cases := []struct {
        input []string
        expected byte
    }{
        {[]string{"abcd","abcde"},'e'},
    }
    for _,c := range cases {
        actual:=findTheDifference(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
