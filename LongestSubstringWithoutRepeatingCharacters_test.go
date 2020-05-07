package leetcode

import "testing"
import "reflect"

func TestLengthOfLongestSubstring(t *testing.T) {
    cases := []struct {
        input string
        expected int
    }{
        {"abcabcbb",3},
        {"bbbbb",1},
        {"pwwkew",3},
        {"",0},
        {"vbnmvbnmkjhgfd",10},
        {"dvdf",3},
    }
    for _,c := range cases {
        actual:=lengthOfLongestSubstring(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
