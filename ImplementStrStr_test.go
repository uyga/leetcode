package leetcode

import "testing"
import "reflect"

func TestStrStr(t *testing.T) {
    cases := []struct {
        input []string
        expected int
    }{
        {[]string{"hello", "ll"},2},
        {[]string{"aaaaaa", "bba"},-1},
        {[]string{"aaaaaa", ""},0},
        {[]string{"aaaaaa", "a"},0},
        {[]string{"", ""},0},
        {[]string{"", "hkshdfh 376428"},-1},
        {[]string{"hellollollollollollooooskjdhfjdhkjs jkshgjkhdskhfk fhkjsdhkjfhjshg jsdghjkdshgkjhdfkjghjkdhfgj dkjdghkfjdhgkdjhfgkjfh jfdhkjghhfgjhdfkghdfjkghdkjhgjkdfklwiur83795872bncnwjehfkjsh fdkhf dkjghskgks jsdhgkhdsg", "jsdhgkhdllljasg ksdfjgkadfhjgjkhdfg jhkfdgkahsfgklhdsfgdkjfg jdfhgjkdhlfgsdfhgjkdhfg jfghkjkhgsjkfghkdhfgjkadfhgjhkgahdfg"},-1},
    }
    for _,c := range cases {
        actual:=strStr(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
