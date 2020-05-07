package leetcode

import "testing"
import "reflect"

func TestFib(t *testing.T) {
    cases := []struct {
        input int
        expected int
    }{
        {3,2},
        {7,13},
        {0,0},
        {15,610},
        {20,6765},
        {30,832040},
    }
    for _,c := range cases {
        actual:=fib(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
