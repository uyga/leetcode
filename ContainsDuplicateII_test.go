package leetcode

import "testing"
import "reflect"

func TestContainsNearbyDuplicate(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected bool
    }{
        {[]int{1,2,3,1},3,true},
        {[]int{1,2,3,1,2,3},2,false},
    }
    for _,c := range cases {
        actual:=containsNearbyDuplicate(c.input1, c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v, %#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
