package leetcode

import "testing"
import "reflect"

func TestIsOneBitCharacter(t *testing.T) {
    cases := []struct {
        input []int
        expected bool
    }{
        {[]int{1,1,0},true},
        {[]int{1,1,1,0},false},
        {[]int{0,0,0,0,0},true},
        {[]int{1,1,1,1,1,0},false},
        {[]int{1,1},false},
        {[]int{0,0},true},
        {[]int{0},true},
    }
    for _,c := range cases {
        actual:=isOneBitCharacter(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
