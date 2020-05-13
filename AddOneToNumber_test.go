package leetcode

import "testing"
import "reflect"

func TestAddOneToNumber(t *testing.T) {
    cases := []struct {
        input []int
        expected []int
    }{
        {[]int{1,2,3},[]int{1,2,4}},
        {[]int{1,2,9},[]int{1,3,0}},
        {[]int{1,9,9},[]int{2,0,0}},
        {[]int{9},[]int{1,0}},
        {[]int{9,9,9},[]int{1,0,0,0}},
        {[]int{0,0,0,9,9,9},[]int{1,0,0,0}},
    }
    for _,c := range cases {
        actual:=addOneToNumber(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
