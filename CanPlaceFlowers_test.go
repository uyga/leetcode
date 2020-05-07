package leetcode

import "testing"
import "reflect"

func TestCanPlaceFlowers(t *testing.T) {
    cases := []struct {
        input1 []int
        input2 int
        expected bool
    }{
        {[]int{0,0,1,0},1,true},
        {[]int{0,0,1,0,0},2,true},
        {[]int{0,0,1,0,0},3,false},
    }
    for _,c := range cases {
        actual:=canPlaceFlowers(c.input1,c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v, %#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
