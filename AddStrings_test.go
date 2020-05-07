package leetcode

import "testing"
import "reflect"

func TestAddStrings(t *testing.T) {
    cases := []struct {
        input []string
        expected string
    }{
        {[]string{"123","1234"},"1357"},
        {[]string{"123456","789012356789"},"789012480245"},
        {[]string{"1","1"},"2"},
        {[]string{"9","9"},"18"},
        {[]string{"0","0"},"0"},
        {[]string{"",""},"0"},
    }
    for _,c := range cases {
        actual:=addStrings(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
