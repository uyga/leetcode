package leetcode

import "testing"
import "reflect"

func TestNumJewelsInStones(t *testing.T) {
    cases := []struct {
        input []string
        expected int
    }{
        {[]string{"aA","aAAbbbb"},3},
        {[]string{"z","ZZZZZZ"},0},
    }
    for _,c := range cases {
        actual:=numJewelsInStones(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
