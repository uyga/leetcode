package leetcode

import "testing"
import "reflect"

func TestMinCostClimbingStairs(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{10, 15, 20},15},
        {[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},6},
        {[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1,1,1,1,1,2,3,4,5,6,7,8,9,100,1,2,3,1},36},
        {[]int{0,0,0,0,0,0},0},
    }
    for _,c := range cases {
        actual:=minCostClimbingStairs(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
