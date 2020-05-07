package leetcode

import "testing"
import "reflect"

func TestBinarySearchInArray(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected int
    }{
        {[]int{-1,0,3,5,9,12},9,4},
        {[]int{-1,0,3,5,9,12},2,-1},
        {[]int{-1,0,3,5,9,12},-1,0},
        {[]int{-1,0,3,5,9,12},12,5},
        {[]int{-1,0,3,5,9,12,20},5,3},
        {[]int{-1,0,3,5,9,12,20},28,-1},
        {[]int{},3,-1},
    }
    for _,c := range cases {
        actual:=binarySearchInArray(c.input1,c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
