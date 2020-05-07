package leetcode

import "testing"
import "reflect"

func TestFindWords(t *testing.T) {
    cases := []struct {
        input []string
        expected []string
    }{
        {[]string{"Hello", "Alaska", "Dad", "Peace"},[]string{"Alaska", "Dad"}},
        {[]string{"", "", "", ""},[]string{"", "", "", ""}},
    }
    for _,c := range cases {
        actual:=findWords(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
