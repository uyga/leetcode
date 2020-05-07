package leetcode

import "testing"
import "reflect"

func TestMissingNumber(t *testing.T) {
    cases := []struct {
        input []int
        expected int
    }{
        {[]int{0,1,2,3,4,5,6,7,9},8},
    }
    for _,c := range cases {
        actual:=missingNumber(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
