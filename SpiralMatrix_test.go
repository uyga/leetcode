package leetcode

import "testing"
import "reflect"

func TestSpiralOrder(t *testing.T) {
    cases := []struct {
        input [][]int
        expected []int
    }{
        {[][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}},[]int{1,2,3,6,9,8,7,4,5}},
        {[][]int{[]int{1,2,3,4},[]int{5,6,7,8},[]int{9,10,11,12}},[]int{1,2,3,4,8,12,11,10,9,5,6,7}},
        {[][]int{[]int{1,2,3,4}},[]int{1,2,3,4}},
        {[][]int{[]int{1},[]int{2},[]int{3},[]int{4}},[]int{1,2,3,4}},
        {[][]int{[]int{1}},[]int{1}},
    }
    for _,c := range cases {
        actual:=spiralOrder(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
