package leetcode

import "testing"
import "reflect"

func TestRemoveElement(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected int
    }{
        {[]int{0,1,2,2,3,0,4,2},2,5},
    }
    for _,c := range cases {
        actual:=removeElement(c.input1,c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
