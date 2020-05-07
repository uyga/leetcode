package leetcode

import "testing"
import "reflect"

func TestMostCommonWord(t *testing.T) {
    cases := []struct {
        input1 string
        input2 []string
        expected string
    }{
        {"Bob hit a ball, the hit BALL flew far after it was hit.",[]string{"hit"},"ball"},
        {"a, a, a, a, b,b,b,c, c",[]string{"a"},"b"},
    }
    for _,c := range cases {
        actual:=mostCommonWord(c.input1, c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
