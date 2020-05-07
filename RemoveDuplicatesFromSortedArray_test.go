package leetcode

import "testing"
import "reflect"

func TestRemoveDuplicates(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{1,1,1,1,1,1,1,1,1,2,2,2,2,3,4,5,5,5,5,5,5},5},
    }
    for _,c := range cases {
        actual:=removeDuplicates(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
