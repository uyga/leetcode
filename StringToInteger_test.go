package leetcode

import "testing"
import "reflect"

func TestMyAtoi(t *testing.T) {
    cases := []struct {
        input string
        expected int
    }{
        {"1",1},
        {"42",42},
        {"1040",1040},
        {"   -42",-42},
        {" -  42",0},
        {"   34",34},
        {"   +34",34},
        {" ++ 34",0},
        {"4193 with words",4193},
        {"words and 987",0},
        {"-91283472332",-2147483648},
        {"",0},
        {"      ",0},
        {"10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459",2147483647},
        {"  0000000000012345678",12345678},
    }
    for _,c := range cases {
        actual:=myAtoi(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
