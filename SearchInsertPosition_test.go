package leetcode

import "testing"
import "reflect"

func TestSearchInsert(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected int
    }{
        {[]int{1,3,5,6},5,2},
        {[]int{2,3,5,6,9},7,4},
        {[]int{1,3,5,6},6,3},
        {[]int{1,3,5,6},2,1},
        {[]int{1,3,5,6},7,4},
        {[]int{1,3,5,6},0,0},
        {[]int{},0,0},
    }
    for _,c := range cases {
        actual:=searchInsert(c.input1,c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
