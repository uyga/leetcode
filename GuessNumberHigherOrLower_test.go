package leetcode

import "testing"
import "reflect"

func TestGuessNumber(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {3,1},
        {10,10},
    }
    for _,c := range cases {
        actual:=guessNumber(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
