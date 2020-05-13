package leetcode

import "testing"
import "reflect"

func TestLemonadeChange(t *testing.T) {
    cases := []struct {
        input []int
        expected bool
    }{
        {[]int{5,5,5,10,20},true},
        {[]int{5,5,10},true},
        {[]int{10,10},false},
        {[]int{5,5,10,10,20},false},
        {[]int{5,5,10,20,5,5,5,5,5,5,5,5,5,10,5,5,20,5,20,5},true},
    }
    for _,c := range cases {
        actual:=lemonadeChange(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
