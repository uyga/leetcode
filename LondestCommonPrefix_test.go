package leetcode

import "testing"
import "reflect"

func TestLongestCommonPrefix(t *testing.T) {
    cases := []struct {
        input []string
        expected string
    }{
        {[]string{"flower","flow","flight"},"fl"},
        {[]string{"flower","flow",""},""},
        {[]string{"flower","flow","flooooooo"},"flo"},
        {[]string{"dog","racecar","car"},""},
        {[]string{"","",""},""},
        {[]string{"","racecar","car"},""},
        {[]string{"o","o","o"},"o"},
    }
    for _,c := range cases {
        actual:=longestCommonPrefix(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
