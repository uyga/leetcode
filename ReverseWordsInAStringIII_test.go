package leetcode

import "testing"
import "reflect"

func TestReverseWords(t *testing.T) {
    cases := []struct {
        input string
        expected string
    }{
        {"Let's take LeetCode contest h mmmm i","s'teL ekat edoCteeL tsetnoc h mmmm i"},
    }
    for _,c := range cases {
        actual:=reverseWords(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
