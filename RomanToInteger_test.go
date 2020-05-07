package leetcode

import "testing"
import "reflect"

func TestRomanToInt(t *testing.T) {
    cases := []struct {
        input string
        expected int
    }{
        {"D",500},
        {"III",3},
        {"VIII",8},
        {"XII",12},
        {"IV",4},
        {"IX",9},
        {"LVIII",58},
        {"MCMXCIV",1994},
        {"MDCXCV",1695},
    }
    for _,c := range cases {
        actual:=romanToInt(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
