package leetcode

import "testing"
import "reflect"

func TestFizzBuzz(t *testing.T) {
    cases := []struct {
        input int
        expected []string
    }{
        {15,[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
    }
    for _,c := range cases {
        actual:=fizzBuzz(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
