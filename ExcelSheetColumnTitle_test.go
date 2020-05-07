package leetcode

import "testing"
import "reflect"

func TestConvertToTitle(t *testing.T) {
    cases := []struct {
        input int
        expected string
    }{
        {703,"AAA"},
        {702,"ZZ"},
        {701,"ZY"},
        {29,"AC"},
        {28,"AB"},
        {27,"AA"},
        {26,"Z"},
        {25,"Y"},
        {24,"X"},
        {2,"B"},
        {1,"A"},
    }
    for _,c := range cases {
        actual:=convertToTitle(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
