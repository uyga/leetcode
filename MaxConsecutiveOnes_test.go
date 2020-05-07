package leetcode

import "testing"
import "reflect"

func TestFindMaxConsecutiveOnes(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{1,1,0,1,1,1},3},
        {[]int{1},1},
        {[]int{0},0},
        {[]int{1,1,1,1,1,1,1,1},8},
        {[]int{0,0,0,0,0,0,0,0,0},0},
        {[]int{},0},
    }
    for _,c := range cases {
        actual:=findMaxConsecutiveOnes(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
