package leetcode

import "testing"
import "reflect"

func TestIntersect(t *testing.T) {
    cases := []struct {
        input [][]int
        expected []int
    }{
        {[][]int{[]int{2,2,2,1,2,3,2,1,2,1,2,1,2,2,2,2},[]int{9,8,1,1,1,1,1,1,1,1,7,0,1,0,2,0,0,0,0,0,2,0,0,0,0}},[]int{1, 1, 1, 1, 2, 2}},
        {[][]int{[]int{2,2,2,1,2,3,2,1,2,1,2,1,2,2,2,2},[]int{9,8,7,0,0,5,0,0,0,0,0,9,0,0,0,0}},nil},
        {[][]int{[]int{},[]int{9,8,1,1,1,1,1,1,1,1,7,0,1,0,2,0,0,0,0,0,2,0,0,0,0}},nil},
        {[][]int{[]int{2,2,2,1,2,3,2,1,2,1,2,1,2,2,2,2},[]int{}},nil},
        {[][]int{[]int{1,2,3,4,5,6},[]int{1,2,3,4,5,6}},[]int{1, 2, 3, 4, 5, 6}},
        {[][]int{[]int{1,2,3,4,5,6},[]int{6,5,4,3,2,1}},[]int{6, 5, 4, 3, 2, 1}},
    }
    for _,c := range cases {
        actual:=intersect(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
