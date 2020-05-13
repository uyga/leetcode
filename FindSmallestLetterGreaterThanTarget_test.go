package leetcode

import "testing"
import "reflect"

func TestNextGreatestLetter(t *testing.T) {
    cases := []struct {
        input1 []byte
        input2 byte
        expected byte
    }{
        {[]byte{'c', 'f', 'j'},'a','c'},
        {[]byte{'c', 'f', 'j'},'c','f'},
        {[]byte{'c', 'f', 'j'},'g','j'},
        {[]byte{'c', 'f', 'j'},'j','c'},
        {[]byte{'c', 'f', 'j'},'k','c'},
        {[]byte{'a', 'b', 'c'},'z','a'},
    }
    for _,c := range cases {
        actual:=nextGreatestLetter(c.input1,c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
