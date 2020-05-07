package leetcode

import "testing"
import "reflect"

func TestCompress(t *testing.T) {
    cases := []struct {
        input []byte
        expected int
        output []byte
    }{
        {[]byte{'a','a','b','b','c','c','c'},6,[]byte{'a','2','b','2','c','3'}},
        {[]byte{'a'},1,[]byte{'a'}},
        {[]byte{'a','b'},2,[]byte{'a','b'}},
        {[]byte{'a','b','b'},3,[]byte{'a','b','2'}},
        {[]byte{'a','b','b','b','b','b','b','b','b','b','b','b','b','b'},4,[]byte{'a','b','1','3'}},
        {[]byte{'a','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b','b'},4,[]byte{'a','b','3','8'}},
        {[]byte{'a','b','b','b','b','b','b','b','b','b','b','b','b','b','c','c'},6,[]byte{'a','b','1','3','c','2'}},
        {[]byte{'a','b','b','b','b','b','b','b','b','b','b','b','b','b','c'},5,[]byte{'a','b','1','3','c'}},
    }
    for _,c := range cases {
        data:=c.input
        actual:=compress(data)
        if !reflect.DeepEqual(actual, c.expected) || !reflect.DeepEqual(data[:actual], c.output) {
            t.Errorf("Input %#v. Expected: %#v, %#v actual: %#v, %#v\n", c.input, c.expected, c.output, actual, data[:actual])
        }
    }
}
