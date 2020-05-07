package leetcode

import "testing"
import "reflect"

func TestArrangeCoins(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {0,0},
        {1,1},
        {2,1},
        {3,2},
        {4,2},
        {5,2},
        {9,3},
        {10,4},
        {15,5},
        {65536,361},
    }
    for _,c := range cases {
        actual:=arrangeCoins(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
