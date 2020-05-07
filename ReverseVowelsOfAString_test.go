package leetcode

import "testing"
import "reflect"

func TestReverseVowels(t *testing.T) {
    cases := []struct {
        input string
        expected string
    }{
        {"hello","holle"},
        {"leetcode","leotcede"},
        {"qwerty","qwerty"},
        {"aeea","aeea"},
        {"aeae","eaea"},
        {"qew","qew"},
        {"q","q"},
        {"o","o"},
        {"",""},
    }
    for _,c := range cases {
        actual:=reverseVowels(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
