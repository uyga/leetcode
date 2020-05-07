package leetcode

import "testing"
import "reflect"

func TestMoveZeroes(t *testing.T) {
    cases := []struct {
        input []int
        expected []int
    }{
        {[]int{0,1,0,3,12},[]int{1,3,12,0,0}},
        {[]int{1,3,12,2,1},[]int{1,3,12,2,1}},
        {[]int{0,3,12,2,1},[]int{3,12,2,1,0}},
        {[]int{1,3,12,2,0},[]int{1,3,12,2,0}},
        {[]int{0},[]int{0}},
        {[]int{0,0,0,0,0,1},[]int{1,0,0,0,0,0}},
        {[]int{0,0,0,0,0,0},[]int{0,0,0,0,0,0}},
    }
    for _,c := range cases {
        nums := c.input
        moveZeroes(nums)
        actual := nums
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
