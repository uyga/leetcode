package leetcode

import "testing"
import "reflect"

func TestThirdMax(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{1,1,2},2},
        {[]int{3,2,1},1},
        {[]int{1,1,2,2,2,2,2,2,2,2,2,2},2},
        {[]int{1,2,3,4,5,6,7,8,8,8,8,8,9},7},
    }
    for _,c := range cases {
        actual:=thirdMax(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
