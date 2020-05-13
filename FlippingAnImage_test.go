package leetcode

import "testing"
import "reflect"

func TestFlipAndInvertImage(t *testing.T) {
    cases := []struct {
        input [][]int
        expected [][]int
    }{
        {[][]int{[]int{1,1,0},[]int{1,0,1},[]int{0,0,0}},[][]int{[]int{1,0,0},[]int{0,1,0},[]int{1,1,1}}},
        {[][]int{[]int{1,1,0,0},[]int{1,0,0,1},[]int{0,1,1,1},[]int{1,0,1,0}},[][]int{[]int{1,1,0,0},[]int{0,1,1,0},[]int{0,0,0,1},[]int{1,0,1,0}}},
        {[][]int{[]int{1}},[][]int{[]int{0}}},
    }
    for _,c := range cases {
        actual:=flipAndInvertImage(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
