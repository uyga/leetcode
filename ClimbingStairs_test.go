package leetcode

import "testing"
import "reflect"

func TestClimbStairs(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {1,1},
        {2,2},
        {3,3},
        {4,5},
        {20,10946},
        {30,1346269},
    }
    for _,c := range cases {
        actual:=climbStairs(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
