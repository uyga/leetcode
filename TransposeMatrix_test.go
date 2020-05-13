package leetcode

import "testing"
import "reflect"

func TestTransposeMatrix(t *testing.T) {
    cases := []struct {
        input [][]int
        expected [][]int
    }{
        {[][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}},[][]int{[]int{1,4,7},[]int{2,5,8},[]int{3,6,9}}},
        {[][]int{[]int{1,2,3},[]int{4,5,6}},[][]int{[]int{1,4},[]int{2,5},[]int{3,6}}},
        {[][]int{[]int{1}},[][]int{[]int{1}}},
        {[][]int{[]int{1},[]int{2}},[][]int{[]int{1,2}}},
    }
    for _,c := range cases {
        actual:=transposeMatrix(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
