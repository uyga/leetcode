package leetcode

import "testing"
import "reflect"

func TestIsPalindrome(t *testing.T) {
    cases := []struct {
        input string
        expected bool
    }{
        {"0P",false},
        {"                    A man, a plan, a canal: Panama",true},
        {"race a car",false},
        {"",true},
        {"a",true},
    }
    for _,c := range cases {
        actual:=isPalindrome(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
