package leetcode

import "testing"
import "reflect"

func TestAddBinary(t *testing.T) {
    cases := []struct {
        input []string
        expected string
    }{
        {[]string{"101", "111"},"1100"},
        {[]string{"1010101", "11010101111111"},"11010111010100"},
        {[]string{"11", "1"},"100"},
        {[]string{"1", "1"},"10"},
        {[]string{"0", "1"},"1"},
        {[]string{"0" , "0"},"0"},
        {[]string{"1", "0"},"1"},
        {[]string{"1010", "11"},"1101"},
        {[]string{"1010", "1011"},"10101"},
        {[]string{"100101", "10101"},"111010"},
    }
    for _,c := range cases {
        actual:=addBinary(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
