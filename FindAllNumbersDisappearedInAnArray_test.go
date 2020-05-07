package leetcode

import "testing"
import "reflect"

func TestFindDisappearedNumbers(t *testing.T) {
    cases := []struct {
        input []int
        expected []int
    }{
        {[]int{4,3,2,7,8,2,3,1},[]int{5,6}},
        {[]int{2,3,1},nil},
    }
    for _,c := range cases {
        actual:=findDisappearedNumbers(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
