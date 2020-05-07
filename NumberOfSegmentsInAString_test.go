package leetcode

import "testing"
import "reflect"

func TestCountSegments(t *testing.T) {
    cases := []struct {
        input string
        expected int
    }{
        {"Hello, my name is John",5},
        {"    Hello,     my name is John     ",5},
        {"    ",0},
        {"     nnn",1},
        {"lll    ", 1},
        {"llsdasd", 1},
        {"l",1},
        {"",0},
    }
    for _,c := range cases {
        actual:=countSegments(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
