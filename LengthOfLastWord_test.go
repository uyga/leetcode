package leetcode

import "testing"
import "reflect"

func TestLengthOfLastWord(t *testing.T) {
    cases := []struct {
        input string
        expected int
    }{
        {"p ",1},
        {"Hello World",5},
        {"HelloWorld",10},
        {"",0},
        {" ",0},
        {"Hello World    d d",1},
    }
    for _,c := range cases {
        actual:=lengthOfLastWord(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
