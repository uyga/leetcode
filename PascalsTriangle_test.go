package leetcode

import "testing"
import "reflect"

func TestGenerate(t *testing.T) {
    cases := []struct {
        input int
        expected [][]int
    }{
        {10,[][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}, []int{1, 4, 6, 4, 1}, []int{1, 5, 10, 10, 5, 1}, []int{1, 6, 15, 20, 15, 6, 1}, []int{1, 7, 21, 35, 35, 21, 7, 1}, []int{1, 8, 28, 56, 70, 56, 28, 8, 1}, []int{1, 9, 36, 84, 126, 126, 84, 36, 9, 1}}},
    }
    for _,c := range cases {
        actual:=generate(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
