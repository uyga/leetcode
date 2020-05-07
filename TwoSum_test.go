package leetcode

import "testing"
import "reflect"

func TestTwoSum(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected []int
    }{
        {[]int{1,2,3,4,5,6,7,8,9,10},8,[]int{4,2}},
    }
    for _,c := range cases {
        actual:=twoSum(c.input1, c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
