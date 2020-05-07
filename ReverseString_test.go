package leetcode

import "testing"
import "reflect"

func TestReverseString(t *testing.T) {
    cases := []struct {
        input []byte
        expected []byte
    }{
        {[]byte{'R','e','v','e','r','s','e'},[]byte{'e','s','r','e','v','e','R'}},
        {[]byte{'R','e','v','e','r','s','e','R'},[]byte{'R','e','s','r','e','v','e','R'}},
        {[]byte{'R'},[]byte{'R'}},
        {[]byte{},[]byte{}},
    }
    for _,c := range cases {
        data:=c.input
        reverseString(data)
        actual:=data
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
