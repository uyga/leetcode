package leetcode

import "testing"
import "reflect"

func TestNumArraySumRange(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{0,2},1},
        {[]int{2,5},-1},
        {[]int{0,5},-3},
        {[]int{4,25},1137},
        {[]int{0,0},-2},
        {[]int{0,1},-2},
    }
    obj := NumArrayConstructor([]int{-2,0,3,-5,2,-1,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,15,16,17,80,888})
    for _,c := range cases {
        actual:=obj.SumRange(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
