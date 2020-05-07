package leetcode

import "testing"
import "reflect"

func TestRotateArray(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected []int
    }{
        {[]int{1,2,3,4,5,6,7,8},4,[]int{5,6,7,8,1,2,3,4}},
    }
    for _,c := range cases {
        array:=c.input1
        rotateArray(array,c.input2)
        actual:=array
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
