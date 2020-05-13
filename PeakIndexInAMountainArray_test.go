package leetcode

import "testing"
import "reflect"

func TestPeakIndexInMountainArray(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{0,1,2,3,4,5,6,7,100,8,7,6,5,4,34,3,1},8},
        {[]int{0,1,0},1},
        {[]int{0,2,1,0},1},
    }
    for _,c := range cases {
        actual:=peakIndexInMountainArray(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
