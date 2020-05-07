package leetcode

import "testing"
import "reflect"

func TestIsAnagram(t *testing.T) {
    cases := []struct {
        input []string
        expected bool
    }{
        {[]string{"aa", "bb"},false},
        {[]string{"aaa", "bbb"},false},
        {[]string{"anagram", "anargam"},true},
        {[]string{"anagram", "nagaram"},true},
        {[]string{"anagram", "nagaramas"},false},
        {[]string{"anagram", "nagaran"},false},
        {[]string{"", ""},true},
        {[]string{"a", "a"},true},
        {[]string{"a", "b"},false},
        {[]string{"界世界", "世界界"},true},
    }
    for _,c := range cases {
        actual:=isAnagram(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
