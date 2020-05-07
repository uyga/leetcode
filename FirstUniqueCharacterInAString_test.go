package leetcode

import "testing"
import "reflect"

func TestFirstUniqChar(t *testing.T) {
    cases := []struct {
        input string
        expected int
    }{
        {"leetcode",0},
        {"loveleetcode",2},
        {"llaassmmdd",-1},
        {"llaassmmd",8},
    }
    for _,c := range cases {
        actual:=firstUniqChar(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
