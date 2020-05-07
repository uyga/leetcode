package leetcode

import "testing"
import "reflect"

func TestMaxCount(t *testing.T) {
    cases := []struct {
        input1 int
        input2 int
        input3 [][]int
        expected int
    }{
        {100,20,[][]int{[]int{2,2},[]int{3,3},[]int{1,1},[]int{20,10},[]int{15,1},[]int{1,15}},1},
    }
    for _,c := range cases {
        actual:=maxCount(c.input1,c.input2,c.input3)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.input3, c.expected, actual)
        }
    }
}
