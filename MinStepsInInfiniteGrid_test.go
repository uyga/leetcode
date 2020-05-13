package leetcode

import "testing"
import "reflect"

func TestCoverPoints(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 []int
        expected int
    }{
        {[]int{1,2,3},[]int{4,5,6},2},
        {[]int{121,23,434},[]int{324,234,521},509},
        {[]int{0,0,0,0,0},[]int{0,0,0,0,0},0},
    }
    for _,c := range cases {
        actual:=coverPoints(c.input1, c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
