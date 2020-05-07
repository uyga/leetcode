package leetcode

import "testing"
import "reflect"

func TestContainsDuplicate(t *testing.T) {
    cases := []struct {
        input []int
        expected bool
    }{
        {[]int{1,2,3,4,5,6},false},
        {[]int{1,2,3,4,5,6,1},true},
        {[]int{},false},
    }
    for _,c := range cases {
        actual:=containsDuplicate(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
